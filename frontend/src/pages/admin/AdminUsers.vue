<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import {
  PhDownloadSimple,
  PhFileCsv,
  PhMagnifyingGlass,
  PhPlusCircle,
  PhShieldCheck,
  PhTrash,
  PhUploadSimple,
  PhUsers,
  PhWarningCircle,
} from "@phosphor-icons/vue";
import { useAuthStore } from "../../stores/auth";
import { useToastStore } from "../../stores/toast";
import {
  getRoles,
  syncUserRoles,
} from "../../services/adminUser";
import {
  createAdminSchoolMember,
  getAdminSchoolMembers,
  removeAdminSchoolMember,
} from "../../services/adminSchoolMember";
import {
  commitSchoolMemberImport,
  previewSchoolMemberImport,
} from "../../services/adminSchoolMemberImport";
import type {
  RoleItem,
} from "../../types/adminUser";
import type {
  AdminSchoolMemberCreatePayload,
  AdminSchoolMemberItem,
} from "../../types/adminSchoolMember";
import type {
  AdminSchoolMemberImportCommitResponse,
  AdminSchoolMemberImportPreviewResponse,
} from "../../types/adminSchoolMemberImport";
import { formatDateTime } from "../../utils/date";

const allowedRoleNames = ["student", "teacher", "admin"];
const auth = useAuthStore();
const toast = useToastStore();

const currentSchool = computed(() => {
  const activeId = auth.activeSchoolId ?? auth.defaultContext?.schoolId ?? null;
  const current =
    auth.memberships.find((membership) => membership.school.id === activeId) ??
    auth.memberships.find((membership) => membership.isDefault) ??
    auth.memberships[0] ??
    null;

  return {
    schoolId: activeId ?? "",
    schoolCode: current?.school.code ?? "",
    schoolName: current?.school.name ?? "",
    hasContext: Boolean(activeId && current?.school.code),
  };
});

const members = ref<AdminSchoolMemberItem[]>([]);
const roles = ref<RoleItem[]>([]);
const memberRoleDrafts = ref<Record<string, string>>({});

const membersLoading = ref(false);
const rolesLoading = ref(false);
const savingRolesSchoolUserId = ref("");
const importPreviewLoading = ref(false);
const importCommitLoading = ref(false);
const isCreatingMember = ref(false);
const removingSchoolUserId = ref("");

const membersError = ref("");
const rolesError = ref("");
const importError = ref("");

const memberSearch = ref("");
const importFile = ref<File | null>(null);
const importDefaultPassword = ref("");
const importPreview = ref<AdminSchoolMemberImportPreviewResponse | null>(null);
const importResult = ref<AdminSchoolMemberImportCommitResponse | null>(null);
const manualForm = ref<AdminSchoolMemberCreatePayload>({
  fullName: "",
  email: "",
  password: "",
  role: "student",
  classCode: "",
});

const allowedRoles = computed(() =>
  roles.value.filter((role) =>
    allowedRoleNames.includes(normalizeRoleName(role.roleName)),
  ),
);

function normalizeRoleName(roleName: string) {
  return roleName.trim().toLowerCase();
}

function roleLabel(roleName: string) {
  const normalized = normalizeRoleName(roleName);
  if (normalized === "student") return "Siswa";
  if (normalized === "teacher") return "Guru";
  if (normalized === "admin") return "Admin sekolah";
  return roleName;
}

function rolePriority(roleName: string) {
  const normalized = normalizeRoleName(roleName);
  if (normalized === "admin") return 0;
  if (normalized === "teacher") return 1;
  if (normalized === "student") return 2;
  return 99;
}

function initializeRoleDrafts() {
  const roleByName = new Map(
    allowedRoles.value.map((role) => [
      normalizeRoleName(role.roleName),
      role.roleId,
    ]),
  );
  const nextDrafts: Record<string, string> = {};

  for (const member of members.value) {
    const selectedRole =
      member.roles
        ?.filter((roleName) => roleByName.has(normalizeRoleName(roleName)))
        .sort((a, b) => rolePriority(a) - rolePriority(b))[0] ?? "";
    nextDrafts[member.schoolUserId] = selectedRole
      ? (roleByName.get(normalizeRoleName(selectedRole)) ?? "")
      : "";
  }

  memberRoleDrafts.value = nextDrafts;
}

function primaryRoleName(member: AdminSchoolMemberItem) {
  return (
    member.roles
      ?.filter((roleName) =>
        allowedRoleNames.includes(normalizeRoleName(roleName)),
      )
      .sort((a, b) => rolePriority(a) - rolePriority(b))[0] ?? ""
  );
}

function hasMultipleAllowedRoles(member: AdminSchoolMemberItem) {
  const uniqueRoles = new Set(
    member.roles
      ?.map((roleName) => normalizeRoleName(roleName))
      .filter((roleName) => allowedRoleNames.includes(roleName)) ?? [],
  );
  return uniqueRoles.size > 1;
}

async function loadRoles() {
  rolesLoading.value = true;
  rolesError.value = "";
  try {
    const data = await getRoles();
    roles.value = data ?? [];
  } catch {
    rolesError.value = "Daftar peran belum bisa dimuat.";
  } finally {
    rolesLoading.value = false;
  }
}

async function loadMembers() {
  if (!currentSchool.value.hasContext) return;

  membersLoading.value = true;
  membersError.value = "";
  try {
    const data = await getAdminSchoolMembers({
      page: 1,
      limit: 50,
      search: memberSearch.value.trim(),
    });
    members.value = data.data ?? [];
    initializeRoleDrafts();
  } catch {
    membersError.value = "Warga sekolah belum bisa dimuat.";
  } finally {
    membersLoading.value = false;
  }
}

async function syncRoleForMember(schoolUserId: string, roleId: string) {
  if (!roleId) {
    toast.error("Pilih satu peran.");
    return;
  }

  savingRolesSchoolUserId.value = schoolUserId;
  try {
    await syncUserRoles(schoolUserId, { roleIds: [roleId] });
    toast.success("Peran warga sekolah berhasil diperbarui.");
    await loadMembers();
  } catch {
    toast.error("Peran warga sekolah belum bisa diperbarui.");
  } finally {
    savingRolesSchoolUserId.value = "";
  }
}

function setRoleDraft(schoolUserId: string, roleId: string) {
  memberRoleDrafts.value = {
    ...memberRoleDrafts.value,
    [schoolUserId]: roleId,
  };
}

function getApiErrorMessage(error: unknown, fallback: string) {
  if (typeof error === "object" && error !== null && "response" in error) {
    const response = (
      error as {
        response?: { data?: { error?: unknown; message?: unknown } | string };
      }
    ).response;
    if (typeof response?.data === "string") return response.data;
    if (typeof response?.data?.error === "string") return response.data.error;
    if (typeof response?.data?.message === "string")
      return response.data.message;
  }
  return fallback;
}

function downloadTemplate() {
  const csv = "fullName,email,role,classCode\nBudi Santoso,budi@siswa.sch.id,student,X-IPA-1\nSiti Rahma,siti@guru.sch.id,teacher,\nAdmin Sekolah,admin@sekolah.sch.id,admin,\n";
  const blob = new Blob([csv], { type: "text/csv;charset=utf-8" });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = "template-import-warga-sekolah.csv";
  link.click();
  URL.revokeObjectURL(url);
}

function resetImportState() {
  importFile.value = null;
  importPreview.value = null;
  importResult.value = null;
  importError.value = "";
}

async function handleImportFileChange(event: Event) {
  const target = event.target as HTMLInputElement;
  const file = target.files?.[0] ?? null;
  importFile.value = file;
  importPreview.value = null;
  importResult.value = null;
  importError.value = "";
  if (!file) return;

  importPreviewLoading.value = true;
  try {
    importPreview.value = await previewSchoolMemberImport(file);
  } catch (error) {
    importError.value = getApiErrorMessage(
      error,
      "File import belum bisa divalidasi. Pastikan format CSV sesuai template.",
    );
  } finally {
    importPreviewLoading.value = false;
  }
}

async function submitImportCommit() {
  if (!importPreview.value || importPreview.value.rows.length === 0) {
    toast.error("Preview import belum tersedia.");
    return;
  }
  if (importPreview.value.invalidCount > 0) {
    toast.error("Perbaiki baris yang tidak valid sebelum import.");
    return;
  }
  if (!importDefaultPassword.value.trim()) {
    toast.error("Password awal wajib diisi.");
    return;
  }

  importCommitLoading.value = true;
  importError.value = "";
  importResult.value = null;
  try {
    importResult.value = await commitSchoolMemberImport({
      defaultPassword: importDefaultPassword.value,
      rows: importPreview.value.rows,
    });
    toast.success("Import warga sekolah selesai.");
    memberSearch.value = "";
    await loadMembers();
  } catch (error) {
    if (
      typeof error === "object" &&
      error !== null &&
      "response" in error &&
      (error as { response?: { data?: AdminSchoolMemberImportCommitResponse } })
        .response?.data?.results
    ) {
      importResult.value = (
        error as { response: { data: AdminSchoolMemberImportCommitResponse } }
      ).response.data;
    }
    importError.value = getApiErrorMessage(
      error,
      "Import belum bisa diproses. Pastikan semua baris valid dan password awal terisi.",
    );
  } finally {
    importCommitLoading.value = false;
  }
}

function resetManualForm() {
  manualForm.value = {
    fullName: "",
    email: "",
    password: "",
    role: "student",
    classCode: "",
  };
}

async function submitManualMember() {
  const payload: AdminSchoolMemberCreatePayload = {
    fullName: manualForm.value.fullName.trim(),
    email: manualForm.value.email.trim(),
    password: manualForm.value.password,
    role: manualForm.value.role,
    classCode:
      manualForm.value.role === "student"
        ? manualForm.value.classCode?.trim() || undefined
        : undefined,
  };
  if (!payload.fullName || !payload.email || !payload.password || !payload.role) {
    toast.error("Nama, email, password awal, dan peran wajib diisi.");
    return;
  }

  isCreatingMember.value = true;
  try {
    await createAdminSchoolMember(payload);
    toast.success("Warga sekolah berhasil ditambahkan.");
    resetManualForm();
    memberSearch.value = "";
    await loadMembers();
  } catch (error) {
    toast.error(
      getApiErrorMessage(
        error,
        "Warga sekolah belum bisa ditambahkan. Pastikan data valid.",
      ),
    );
  } finally {
    isCreatingMember.value = false;
  }
}

async function removeMember(member: AdminSchoolMemberItem) {
  const confirmed = window.confirm(
    "Akun global tidak akan dihapus. Warga ini hanya dikeluarkan dari sekolah aktif. Lanjutkan?",
  );
  if (!confirmed) return;

  removingSchoolUserId.value = member.schoolUserId;
  try {
    await removeAdminSchoolMember(member.schoolUserId);
    toast.success("Warga sekolah berhasil dihapus dari sekolah aktif.");
    await loadMembers();
  } catch (error) {
    toast.error(
      getApiErrorMessage(
        error,
        "Warga sekolah belum bisa dihapus dari sekolah aktif.",
      ),
    );
  } finally {
    removingSchoolUserId.value = "";
  }
}

onMounted(async () => {
  if (!currentSchool.value.hasContext) return;
  await loadRoles();
  await loadMembers();
});
</script>

<template>
  <main class="min-h-screen min-w-0 flex-1 overflow-x-hidden bg-[#f8f7f4]">
    <header class="border-b border-[#ebe7df] bg-white">
      <div
        class="flex min-w-0 flex-col gap-3 px-5 py-5 sm:px-6 lg:flex-row lg:items-end lg:justify-between lg:px-8"
      >
        <div class="min-w-0">
          <h1 class="mt-1 text-2xl font-semibold text-[#171322] sm:text-3xl">
            Warga Sekolah
          </h1>
          <p class="mt-2 max-w-3xl text-sm leading-6 text-[#6b7280]">
            Kelola warga pada sekolah aktif dan import data siswa, guru, atau
            admin sekolah dari template CSV.
          </p>
        </div>
        <div class="flex min-w-0 flex-wrap gap-2 text-xs">
          <span
            class="max-w-full truncate rounded-lg bg-[#fff4ee] px-3 py-2 font-medium text-[#ea580c]"
          >
            {{ currentSchool.schoolName || "Sekolah belum tersedia" }}
          </span>
          <span
            class="rounded-lg bg-[#f3f4f6] px-3 py-2 font-medium text-[#6b7280]"
          >
            {{ currentSchool.schoolCode || "Kode belum tersedia" }}
          </span>
        </div>
      </div>
    </header>

    <section class="px-5 py-5 sm:px-6 lg:px-8">
      <div
        v-if="!currentSchool.hasContext"
        class="mb-5 flex items-start gap-3 rounded-xl border border-[#fecaca] bg-[#fef2f2] p-4 text-sm leading-6 text-[#dc2626]"
      >
        <PhWarningCircle :size="20" class="mt-0.5 shrink-0" weight="duotone" />
        <p>
          Konteks sekolah aktif belum tersedia. Pastikan akun admin terhubung
          dengan sekolah.
        </p>
      </div>

      <div class="grid min-w-0 gap-5 lg:grid-cols-[minmax(0,1fr)_360px]">
        <section
          class="order-2 min-w-0 rounded-xl border border-[#ebe7df] bg-white lg:order-1"
        >
          <div
            class="flex flex-col gap-4 border-b border-[#ebe7df] px-4 py-4 sm:px-5"
          >
            <div
              class="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between"
            >
              <div>
                <p
                  class="text-[10px] font-medium uppercase tracking-[0.08em] text-[#9ca3af]"
                >
                  Warga sekolah aktif
                </p>
                <h2 class="mt-1 text-base font-semibold text-[#171322]">
                  Daftar pengguna sekolah
                </h2>
                <p class="mt-1 text-sm text-[#6b7280]">
                  Kelola peran utama setiap pengguna pada sekolah aktif.
                </p>
              </div>
              <span
                class="inline-flex shrink-0 items-center gap-2 self-start rounded-lg bg-[#eef2ff] px-3 py-2 text-xs font-medium text-[#4f46e5]"
              >
                <PhUsers :size="16" weight="duotone" />
                {{ members.length }} warga
              </span>
            </div>

            <form
              class="flex min-w-0 flex-col gap-2 sm:flex-row"
              @submit.prevent="loadMembers"
            >
              <input
                v-model="memberSearch"
                type="search"
                placeholder="Cari nama atau email warga sekolah"
                class="min-w-0 flex-1 rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white"
              />
              <button
                type="submit"
                class="inline-flex shrink-0 items-center justify-center gap-2 rounded-lg bg-[#171322] px-4 py-2.5 text-sm font-medium text-white transition hover:bg-[#374151] disabled:opacity-60"
                :disabled="membersLoading || !currentSchool.hasContext"
              >
                <PhMagnifyingGlass :size="17" weight="duotone" />
                Cari
              </button>
            </form>
          </div>

          <div class="p-4 sm:p-5">
            <div v-if="rolesLoading || membersLoading" class="space-y-3">
              <div
                v-for="item in 3"
                :key="item"
                class="h-28 animate-pulse rounded-lg bg-[#fbfaf8]"
              />
            </div>

            <div
              v-else-if="rolesError || membersError"
              class="rounded-lg border border-[#fecaca] bg-[#fef2f2] p-5 text-center"
            >
              <PhWarningCircle
                :size="26"
                class="mx-auto text-[#dc2626]"
                weight="duotone"
              />
              <h3 class="mt-3 text-sm font-semibold text-[#171322]">
                Warga sekolah belum bisa dimuat
              </h3>
              <p class="mt-2 text-sm leading-6 text-[#6b7280]">
                {{ rolesError || membersError }}
              </p>
              <button
                type="button"
                class="mt-4 rounded-lg bg-[#171322] px-4 py-2 text-sm font-medium text-white transition hover:bg-[#374151]"
                @click="rolesError ? loadRoles() : loadMembers()"
              >
                Coba lagi
              </button>
            </div>

            <div
              v-else-if="members.length === 0"
              class="rounded-lg bg-[#fbfaf8] px-5 py-10 text-center"
            >
              <PhUsers
                :size="28"
                class="mx-auto text-[#9ca3af]"
                weight="duotone"
              />
              <h3 class="mt-3 text-sm font-semibold text-[#171322]">
                Belum ada warga sekolah
              </h3>
              <p class="mt-2 text-sm leading-6 text-[#6b7280]">
                Import warga sekolah dari panel kanan untuk memulai.
              </p>
            </div>

            <div v-else class="divide-y divide-[#ebe7df]">
              <article
                v-for="member in members"
                :key="member.schoolUserId"
                class="min-w-0 py-4 first:pt-0 last:pb-0"
              >
                <div
                  class="grid min-w-0 gap-4 xl:grid-cols-[minmax(0,1fr)_280px]"
                >
                  <div class="flex min-w-0 items-start gap-3">
                    <div
                      class="flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-[#ea580c] text-xs font-semibold text-white"
                    >
                      {{ (member.fullName || "W").charAt(0).toUpperCase() }}
                    </div>
                    <div class="min-w-0">
                      <div class="flex min-w-0 flex-wrap items-center gap-2">
                        <h3
                          class="min-w-0 wrap-break-word text-sm font-semibold text-[#171322]"
                        >
                          {{ member.fullName || "Nama tidak tersedia" }}
                        </h3>
                        <span
                          v-if="primaryRoleName(member)"
                          class="rounded-lg bg-[#eef2ff] px-2 py-1 text-[11px] font-medium text-[#4f46e5]"
                        >
                          {{ roleLabel(primaryRoleName(member)) }}
                        </span>
                      </div>
                      <p class="mt-1 break-all text-xs text-[#6b7280]">
                        {{ member.email || "Email tidak tersedia" }}
                      </p>
                      <p class="mt-2 text-[11px] text-[#9ca3af]">
                        Bergabung {{ formatDateTime(member.createdAt) }}
                      </p>
                      <p
                        v-if="member.classCodes?.length"
                        class="mt-2 text-[11px] font-medium text-[#6b7280]"
                      >
                        Kelas: {{ member.classCodes.join(", ") }}
                      </p>
                      <p
                        v-if="hasMultipleAllowedRoles(member)"
                        class="mt-2 rounded-lg border border-[#fde68a] bg-[#fff7ed] px-3 py-2 text-xs leading-5 text-[#92400e]"
                      >
                        Data lama memiliki lebih dari satu peran. Saat
                        diperbarui, satu peran utama akan disimpan.
                      </p>
                    </div>
                  </div>

                  <div class="min-w-0 rounded-lg bg-[#fbfaf8] p-3">
                    <label class="block text-xs font-medium text-[#6b7280]">
                      Peran sekolah
                      <select
                        class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-white px-3 py-2.5 text-sm text-[#171322] outline-none transition focus:border-[#4f46e5]"
                        :value="memberRoleDrafts[member.schoolUserId] ?? ''"
                        :disabled="rolesLoading || allowedRoles.length === 0"
                        @change="
                          setRoleDraft(
                            member.schoolUserId,
                            ($event.target as HTMLSelectElement).value,
                          )
                        "
                      >
                        <option value="" disabled>Pilih peran</option>
                        <option
                          v-for="role in allowedRoles"
                          :key="role.roleId"
                          :value="role.roleId"
                        >
                          {{ roleLabel(role.roleName) }}
                        </option>
                      </select>
                    </label>
                    <button
                      type="button"
                      class="mt-2.5 inline-flex w-full items-center justify-center gap-2 rounded-lg bg-[#171322] px-3 py-2.5 text-sm font-medium text-white transition hover:bg-[#374151] disabled:opacity-60"
                      :disabled="
                        savingRolesSchoolUserId === member.schoolUserId ||
                        !memberRoleDrafts[member.schoolUserId]
                      "
                      @click="
                        syncRoleForMember(
                          member.schoolUserId,
                          memberRoleDrafts[member.schoolUserId] ?? '',
                        )
                      "
                    >
                      <PhShieldCheck :size="17" weight="duotone" />
                      {{
                        savingRolesSchoolUserId === member.schoolUserId
                          ? "Menyimpan..."
                          : "Simpan peran"
                      }}
                    </button>
                    <button
                      type="button"
                      class="mt-2 inline-flex w-full items-center justify-center gap-2 rounded-lg border border-[#fecaca] bg-white px-3 py-2.5 text-sm font-medium text-[#dc2626] transition hover:bg-[#fef2f2] disabled:opacity-60"
                      :disabled="removingSchoolUserId === member.schoolUserId"
                      @click="removeMember(member)"
                    >
                      <PhTrash :size="17" weight="duotone" />
                      {{
                        removingSchoolUserId === member.schoolUserId
                          ? "Menghapus..."
                          : "Hapus dari sekolah"
                      }}
                    </button>
                  </div>
                </div>
              </article>
            </div>
          </div>
        </section>

        <aside class="order-1 min-w-0 lg:order-2">
          <section
            class="rounded-xl border border-[#ebe7df] bg-white p-5 lg:sticky lg:top-6"
          >
            <div class="flex items-start justify-between gap-3">
              <div>
                <p
                  class="text-[10px] font-medium uppercase tracking-[0.08em] text-[#9ca3af]"
                >
                  Tambah warga sekolah
                </p>
                <h2 class="mt-1 text-base font-semibold text-[#171322]">
                  Manual atau import CSV
                </h2>
                <p class="mt-1 text-xs leading-5 text-[#6b7280]">
                  Tambahkan warga ke sekolah aktif. Akun global yang sudah ada
                  dipakai ulang berdasarkan email tanpa membuka daftar pengguna
                  platform.
                </p>
              </div>
              <PhPlusCircle
                :size="21"
                class="text-[#ea580c]"
                weight="duotone"
              />
            </div>

            <form class="mt-5 space-y-3" @submit.prevent="submitManualMember">
              <label class="block text-xs font-medium text-[#6b7280]">
                Nama lengkap
                <input
                  v-model="manualForm.fullName"
                  type="text"
                  placeholder="Nama warga sekolah"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white"
                />
              </label>
              <label class="block text-xs font-medium text-[#6b7280]">
                Email
                <input
                  v-model="manualForm.email"
                  type="email"
                  placeholder="email@sekolah.sch.id"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white"
                />
              </label>
              <label class="block text-xs font-medium text-[#6b7280]">
                Password awal
                <input
                  v-model="manualForm.password"
                  type="password"
                  placeholder="Minimal 6 karakter"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white"
                />
              </label>
              <label class="block text-xs font-medium text-[#6b7280]">
                Peran
                <select
                  v-model="manualForm.role"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition focus:border-[#4f46e5] focus:bg-white"
                >
                  <option value="student">Siswa</option>
                  <option value="teacher">Guru</option>
                  <option value="admin">Admin sekolah</option>
                </select>
              </label>
              <label class="block text-xs font-medium text-[#6b7280]">
                Kode kelas
                <input
                  v-model="manualForm.classCode"
                  type="text"
                  placeholder="Opsional untuk siswa"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white disabled:cursor-not-allowed disabled:opacity-60"
                  :disabled="manualForm.role !== 'student'"
                />
              </label>
              <p class="rounded-lg border border-[#fed7aa] bg-[#fff7ed] px-3 py-2 text-xs leading-5 text-[#92400e]">
                Password awal digunakan untuk akun baru. Pengguna dapat
                mengganti password setelah login.
              </p>
              <button
                type="submit"
                class="inline-flex w-full items-center justify-center gap-2 rounded-lg bg-[#171322] px-4 py-2.5 text-sm font-medium text-white transition hover:bg-[#374151] disabled:opacity-60"
                :disabled="isCreatingMember"
              >
                <PhPlusCircle :size="17" weight="duotone" />
                {{ isCreatingMember ? "Menambahkan..." : "Tambah warga" }}
              </button>
            </form>

            <div class="mt-6 border-t border-[#ebe7df] pt-5">
              <div class="flex items-start gap-3">
                <PhFileCsv
                  :size="21"
                  class="mt-0.5 text-[#ea580c]"
                  weight="duotone"
                />
                <div>
                  <h3 class="text-sm font-semibold text-[#171322]">
                    Import warga sekolah
                  </h3>
                  <p class="mt-1 text-xs leading-5 text-[#6b7280]">
                    Upload template CSV untuk menambahkan banyak warga sekaligus.
                  </p>
                </div>
              </div>

              <div class="mt-4 space-y-4">
              <button
                type="button"
                class="inline-flex w-full items-center justify-center gap-2 rounded-lg border border-[#ebe7df] bg-white px-4 py-2.5 text-sm font-medium text-[#171322] transition hover:border-[#ea580c] hover:text-[#ea580c]"
                @click="downloadTemplate"
              >
                <PhDownloadSimple :size="17" weight="duotone" />
                Download template CSV
              </button>

              <label class="block text-xs font-medium text-[#6b7280]">
                File CSV
                <input
                  type="file"
                  accept=".csv,text/csv"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition file:mr-3 file:rounded-md file:border-0 file:bg-[#fff4ee] file:px-3 file:py-1.5 file:text-xs file:font-semibold file:text-[#ea580c] focus:border-[#4f46e5] focus:bg-white"
                  @change="handleImportFileChange"
                />
              </label>

              <label class="block text-xs font-medium text-[#6b7280]">
                Password awal untuk akun baru
                <input
                  v-model="importDefaultPassword"
                  type="password"
                  placeholder="Minimal 6 karakter"
                  class="mt-2 w-full rounded-lg border border-[#ebe7df] bg-[#fbfaf8] px-3.5 py-2.5 text-sm text-[#171322] outline-none transition placeholder:text-[#9ca3af] focus:border-[#4f46e5] focus:bg-white"
                />
              </label>

              <p class="rounded-lg border border-[#fed7aa] bg-[#fff7ed] px-3 py-2 text-xs leading-5 text-[#92400e]">
                Password awal hanya dipakai untuk akun baru. Pengguna dapat
                mengganti password setelah login.
              </p>

              <button
                type="button"
                class="inline-flex w-full items-center justify-center gap-2 rounded-lg bg-[#ea580c] px-4 py-2.5 text-sm font-medium text-white transition hover:bg-[#c2410c] disabled:opacity-60"
                :disabled="
                  importCommitLoading ||
                  importPreviewLoading ||
                  !importPreview ||
                  importPreview.invalidCount > 0
                "
                @click="submitImportCommit"
              >
                <PhUploadSimple :size="17" weight="duotone" />
                {{ importCommitLoading ? "Mengimport..." : "Import warga" }}
              </button>

              <button
                v-if="importPreview || importResult || importFile"
                type="button"
                class="inline-flex w-full items-center justify-center rounded-lg border border-[#ebe7df] bg-white px-4 py-2.5 text-sm font-medium text-[#171322] transition hover:bg-[#fbfaf8]"
                :disabled="importCommitLoading || importPreviewLoading"
                @click="resetImportState"
              >
                Reset import
              </button>
            </div>
            </div>

            <div class="mt-4 space-y-3">
              <p
                v-if="importError"
                class="rounded-lg bg-[#fef2f2] px-3 py-2 text-xs leading-5 text-[#dc2626]"
              >
                {{ importError }}
              </p>
              <p
                v-else-if="importPreviewLoading"
                class="rounded-lg bg-[#fbfaf8] px-3 py-3 text-xs leading-5 text-[#6b7280]"
              >
                Memvalidasi file import...
              </p>
              <p
                v-else-if="!importPreview"
                class="rounded-lg bg-[#fbfaf8] px-3 py-3 text-xs leading-5 text-[#6b7280]"
              >
                Pilih file CSV untuk melihat preview validasi.
              </p>

              <div
                v-if="importPreview"
                class="rounded-lg border border-[#ebe7df] bg-[#fbfaf8] p-3"
              >
                <div class="flex flex-wrap gap-2 text-xs">
                  <span class="rounded-lg bg-[#ecfdf3] px-2.5 py-1 font-semibold text-[#027a48]">
                    {{ importPreview.validCount }} valid
                  </span>
                  <span class="rounded-lg bg-[#fef2f2] px-2.5 py-1 font-semibold text-[#dc2626]">
                    {{ importPreview.invalidCount }} invalid
                  </span>
                </div>

                <div class="mt-3 max-h-72 space-y-2 overflow-y-auto pr-1">
                  <article
                    v-for="row in importPreview.rows"
                    :key="row.rowNumber"
                    class="rounded-lg border bg-white p-3"
                    :class="
                      row.status === 'valid'
                        ? 'border-[#bbf7d0]'
                        : 'border-[#fecaca]'
                    "
                  >
                    <div class="flex items-start justify-between gap-2">
                      <div class="min-w-0">
                        <p class="text-xs font-semibold text-[#171322]">
                          Baris {{ row.rowNumber }} · {{ row.fullName || "Nama kosong" }}
                        </p>
                        <p class="mt-1 break-all text-xs text-[#6b7280]">
                          {{ row.email || "Email kosong" }}
                        </p>
                        <p class="mt-1 text-xs text-[#6b7280]">
                          {{ roleLabel(row.role) }}
                          <span v-if="row.classCode"> · {{ row.classCode }}</span>
                        </p>
                      </div>
                      <span
                        class="shrink-0 rounded-full px-2 py-1 text-[10px] font-semibold"
                        :class="
                          row.status === 'valid'
                            ? 'bg-[#ecfdf3] text-[#027a48]'
                            : 'bg-[#fef2f2] text-[#dc2626]'
                        "
                      >
                        {{ row.status === "valid" ? "Valid" : "Invalid" }}
                      </span>
                    </div>
                    <ul
                      v-if="row.errors.length > 0"
                      class="mt-2 list-disc space-y-1 pl-4 text-xs leading-5 text-[#dc2626]"
                    >
                      <li v-for="error in row.errors" :key="error">
                        {{ error }}
                      </li>
                    </ul>
                  </article>
                </div>
              </div>

              <div
                v-if="importResult"
                class="rounded-lg border border-[#bbf7d0] bg-[#f0fdf4] p-3 text-xs leading-5 text-[#166534]"
              >
                <p class="font-semibold">Import selesai</p>
                <p class="mt-1">
                  {{ importResult.importedCount }} diproses,
                  {{ importResult.skippedCount }} dilewati,
                  {{ importResult.failedCount }} gagal.
                </p>
              </div>
            </div>
          </section>
        </aside>
      </div>
    </section>
  </main>
</template>
