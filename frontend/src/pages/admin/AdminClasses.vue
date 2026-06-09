<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import {
  PhBookOpen,
  PhCalendarBlank,
  PhChalkboardTeacher,
  PhPlusCircle,
  PhWarningCircle,
} from "@phosphor-icons/vue";
import { useAuthStore } from "../../stores/auth";
import { useToastStore } from "../../stores/toast";
import { getAcademicYearsBySchool, getTermsByAcademicYear } from "../../services/adminAcademic";
import { createAdminClass, getAdminClasses } from "../../services/adminClass";
import type { AcademicYearItem, TermItem } from "../../types/adminAcademic";
import type { AdminClassItem } from "../../types/adminClass";
import { formatDateTime } from "../../utils/date";

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

const academicYears = ref<AcademicYearItem[]>([]);
const terms = ref<TermItem[]>([]);
const classes = ref<AdminClassItem[]>([]);
const selectedAcademicYearId = ref("");
const selectedTermId = ref("");

const yearsLoading = ref(false);
const termsLoading = ref(false);
const classesLoading = ref(false);
const yearsError = ref("");
const termsError = ref("");
const classesError = ref("");
const isCreating = ref(false);

const classForm = ref({
  classCode: "",
  classTitle: "",
  classDesc: "",
});

const selectedAcademicYear = computed(
  () =>
    academicYears.value.find((year) => year.academicYearId === selectedAcademicYearId.value) ??
    null,
);

const selectedTerm = computed(
  () => terms.value.find((term) => term.termId === selectedTermId.value) ?? null,
);

async function loadAcademicYears() {
  if (!currentSchool.value.hasContext) return;
  yearsLoading.value = true;
  yearsError.value = "";

  try {
    const data = await getAcademicYearsBySchool(currentSchool.value.schoolCode);
    academicYears.value = data.data ?? [];
    const defaultYear =
      academicYears.value.find((year) => year.isActive) ?? academicYears.value[0] ?? null;
    selectedAcademicYearId.value = defaultYear?.academicYearId ?? "";
  } catch {
    yearsError.value = "Tahun ajaran belum bisa dimuat.";
  } finally {
    yearsLoading.value = false;
  }
}

async function loadTerms(selectDefault = false) {
  terms.value = [];
  termsError.value = "";
  selectedTermId.value = selectDefault ? "" : selectedTermId.value;

  if (!selectedAcademicYearId.value) return;

  termsLoading.value = true;
  try {
    const data = await getTermsByAcademicYear(selectedAcademicYearId.value);
    terms.value = data ?? [];

    const selectedStillValid = terms.value.some((term) => term.termId === selectedTermId.value);
    if (selectDefault || !selectedStillValid) {
      const defaultTerm = terms.value.find((term) => term.isActive) ?? terms.value[0] ?? null;
      selectedTermId.value = defaultTerm?.termId ?? "";
    }
  } catch {
    termsError.value = "Semester belum bisa dimuat.";
  } finally {
    termsLoading.value = false;
  }
}

async function loadClasses() {
  classes.value = [];
  classesError.value = "";

  if (!currentSchool.value.hasContext || !selectedTermId.value) return;

  classesLoading.value = true;
  try {
    const data = await getAdminClasses({
      schoolCode: currentSchool.value.schoolCode,
      termId: selectedTermId.value,
      page: 1,
      limit: 50,
    });
    classes.value = data.data?.data ?? [];
  } catch {
    classesError.value = "Daftar kelas belum bisa dimuat.";
  } finally {
    classesLoading.value = false;
  }
}

async function handleAcademicYearChange() {
  await loadTerms(true);
  await loadClasses();
}

async function handleTermChange() {
  await loadClasses();
}

async function submitClass() {
  if (!currentSchool.value.schoolId || !currentSchool.value.schoolCode) {
    toast.error("Context sekolah aktif belum tersedia.");
    return;
  }
  if (!selectedTermId.value) {
    toast.error("Pilih semester terlebih dahulu.");
    return;
  }
  if (!classForm.value.classCode.trim() || !classForm.value.classTitle.trim()) {
    toast.error("Kode dan nama kelas wajib diisi.");
    return;
  }

  isCreating.value = true;
  try {
    await createAdminClass({
      schoolId: currentSchool.value.schoolId,
      termId: selectedTermId.value,
      classCode: classForm.value.classCode.trim(),
      classTitle: classForm.value.classTitle.trim(),
      classDesc: classForm.value.classDesc.trim(),
    });
    classForm.value = { classCode: "", classTitle: "", classDesc: "" };
    toast.success("Kelas berhasil dibuat.");
    await loadClasses();
  } catch {
    toast.error("Kelas belum bisa dibuat.");
  } finally {
    isCreating.value = false;
  }
}

onMounted(async () => {
  if (!currentSchool.value.hasContext) return;
  await loadAcademicYears();
  await loadTerms(true);
  await loadClasses();
});
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-6 sm:px-8 lg:px-10">
    <section class="mx-auto flex max-w-6xl flex-col gap-6">
      <header class="soft-card rounded-[22px] p-5">
        <div class="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
          <div>
            <p class="text-[11px] font-medium uppercase text-[#9CA3AF]">Admin sekolah</p>
            <h1 class="mt-2 text-2xl font-medium text-[#111827]">Kelas</h1>
            <p class="mt-2 max-w-3xl text-sm leading-6 text-[#6B7280]">
              Buat dan lihat kelas berdasarkan tahun ajaran serta semester yang dipilih sebelum penempatan kelas dan penugasan mengajar.
            </p>
          </div>
          <div class="flex flex-wrap gap-2 text-xs">
            <span class="rounded-lg bg-[#EEF2FF] px-3 py-1.5 font-medium text-[#4F46E5]">
              {{ currentSchool.schoolName || "Sekolah belum tersedia" }}
            </span>
            <span class="rounded-lg bg-[#F9FAFB] px-3 py-1.5 font-medium text-[#6B7280]">
              {{ currentSchool.schoolCode || "Kode sekolah belum tersedia" }}
            </span>
          </div>
        </div>

        <div
          v-if="!currentSchool.hasContext"
          class="mt-4 rounded-[10px] border border-[#FECACA] bg-[#FEF2F2] px-4 py-3 text-sm text-[#DC2626]"
        >
          Context sekolah aktif belum tersedia. Pastikan akun admin memiliki membership sekolah.
        </div>

      </header>

      <section class="grid gap-5 lg:grid-cols-[minmax(0,0.9fr)_minmax(0,1.1fr)]">
        <article class="bg-white border border-[#EBEBEB] rounded-[18px] p-5">
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-[11px] font-medium uppercase text-[#9CA3AF]">Periode akademik</p>
              <h2 class="mt-2 text-base font-medium text-[#111827]">Pilih tahun ajaran dan semester</h2>
            </div>
            <PhCalendarBlank :size="22" class="text-[#4F46E5]" weight="duotone" />
          </div>

          <div class="mt-5 space-y-4">
            <label class="block text-sm font-medium text-[#374151]">
              Tahun ajaran
              <select
                v-model="selectedAcademicYearId"
                class="mt-2 w-full rounded-2xl border border-[#EBEBEB] bg-white px-4 py-3 text-sm text-[#111827] outline-none transition focus:border-[#4F46E5]"
                :disabled="yearsLoading || academicYears.length === 0"
                @change="handleAcademicYearChange"
              >
                <option value="" disabled>Pilih tahun ajaran</option>
                <option
                  v-for="year in academicYears"
                  :key="year.academicYearId"
                  :value="year.academicYearId"
                >
                  {{ year.academicYearName }}{{ year.isActive ? " - Aktif" : "" }}
                </option>
              </select>
            </label>

            <label class="block text-sm font-medium text-[#374151]">
              Semester
              <select
                v-model="selectedTermId"
                class="mt-2 w-full rounded-2xl border border-[#EBEBEB] bg-white px-4 py-3 text-sm text-[#111827] outline-none transition focus:border-[#4F46E5]"
                :disabled="termsLoading || terms.length === 0"
                @change="handleTermChange"
              >
                <option value="" disabled>Pilih semester</option>
                <option v-for="term in terms" :key="term.termId" :value="term.termId">
                  {{ term.termName }}{{ term.isActive ? " - Aktif" : "" }}
                </option>
              </select>
            </label>
          </div>

          <div class="mt-5 space-y-2 text-sm">
            <p v-if="yearsLoading" class="text-[#6B7280]">Memuat tahun ajaran...</p>
            <p v-else-if="yearsError" class="text-[#DC2626]">{{ yearsError }}</p>
            <p v-else-if="academicYears.length === 0" class="text-[#6B7280]">
              Belum ada tahun ajaran. Buat tahun ajaran di Struktur Akademik terlebih dahulu.
            </p>

            <p v-if="termsLoading" class="text-[#6B7280]">Memuat semester...</p>
            <p v-else-if="termsError" class="text-[#DC2626]">{{ termsError }}</p>
            <p v-else-if="selectedAcademicYearId && terms.length === 0" class="text-[#6B7280]">
              Belum ada semester untuk tahun ajaran ini.
            </p>
          </div>

          <div class="mt-5 rounded-[18px] bg-[#FBFAF8] p-4">
            <p class="text-[11px] font-medium uppercase text-[#9CA3AF]">Context aktif</p>
            <div class="mt-3 space-y-2 text-sm text-[#374151]">
              <p>
                Tahun ajaran:
                <span class="font-medium text-[#111827]">
                  {{ selectedAcademicYear?.academicYearName || "Belum dipilih" }}
                </span>
              </p>
              <p>
                Semester:
                <span class="font-medium text-[#111827]">
                  {{ selectedTerm?.termName || "Belum dipilih" }}
                </span>
              </p>
            </div>
          </div>
        </article>

        <article class="bg-white border border-[#EBEBEB] rounded-[18px] p-5">
          <div class="flex items-start justify-between gap-4">
            <div>
              <p class="text-[11px] font-medium uppercase text-[#9CA3AF]">Create class</p>
              <h2 class="mt-2 text-base font-medium text-[#111827]">Tambah kelas</h2>
            </div>
            <PhPlusCircle :size="22" class="text-[#059669]" weight="duotone" />
          </div>

          <form class="mt-5 grid gap-3 sm:grid-cols-2" @submit.prevent="submitClass">
            <input
              v-model="classForm.classCode"
              type="text"
              placeholder="Kode kelas, contoh: X-IPA-1"
              class="rounded-2xl border border-[#EBEBEB] bg-white px-4 py-3 text-sm text-[#111827] outline-none transition placeholder:text-[#9CA3AF] focus:border-[#4F46E5]"
            >
            <input
              v-model="classForm.classTitle"
              type="text"
              placeholder="Nama kelas"
              class="rounded-2xl border border-[#EBEBEB] bg-white px-4 py-3 text-sm text-[#111827] outline-none transition placeholder:text-[#9CA3AF] focus:border-[#4F46E5]"
            >
            <textarea
              v-model="classForm.classDesc"
              rows="3"
              placeholder="Deskripsi singkat, opsional"
              class="sm:col-span-2 rounded-2xl border border-[#EBEBEB] bg-white px-4 py-3 text-sm text-[#111827] outline-none transition placeholder:text-[#9CA3AF] focus:border-[#4F46E5]"
            />
            <button
              type="submit"
              class="sm:col-span-2 inline-flex items-center justify-center gap-2 rounded-2xl bg-[#111827] px-4 py-3 text-sm font-medium text-white transition hover:bg-[#374151] disabled:cursor-not-allowed disabled:opacity-60"
              :disabled="isCreating || !currentSchool.hasContext || !selectedTermId"
            >
              <PhChalkboardTeacher :size="18" weight="duotone" />
              Buat kelas
            </button>
          </form>
        </article>
      </section>

      <section class="bg-white border border-[#EBEBEB] rounded-[18px] p-5">
        <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
          <div>
            <p class="text-[11px] font-medium uppercase text-[#9CA3AF]">Kelas</p>
            <h2 class="mt-2 text-base font-medium text-[#111827]">Daftar kelas</h2>
            <p class="mt-1 text-sm text-[#6B7280]">
              Ditampilkan berdasarkan sekolah dan semester yang sedang dipilih.
            </p>
          </div>
          <div class="inline-flex items-center gap-2 rounded-lg bg-[#EEF2FF] px-3 py-2 text-xs font-medium text-[#4F46E5]">
            <PhBookOpen :size="16" weight="duotone" />
            {{ classes.length }} kelas
          </div>
        </div>

        <div class="mt-5">
          <div v-if="classesLoading" class="rounded-[18px] bg-[#FBFAF8] p-5 text-sm text-[#6B7280]">
            Memuat daftar kelas...
          </div>

          <div
            v-else-if="classesError"
            class="flex items-start gap-3 rounded-[18px] border border-[#FECACA] bg-[#FEF2F2] p-5 text-sm text-[#DC2626]"
          >
            <PhWarningCircle :size="20" weight="duotone" />
            <p>{{ classesError }}</p>
          </div>

          <div
            v-else-if="!selectedTermId"
            class="rounded-[18px] bg-[#FBFAF8] p-5 text-sm text-[#6B7280]"
          >
            Pilih semester untuk melihat daftar kelas.
          </div>

          <div
            v-else-if="classes.length === 0"
            class="rounded-[18px] bg-[#FBFAF8] p-5 text-sm text-[#6B7280]"
          >
            Belum ada kelas untuk semester ini.
          </div>

          <div v-else class="grid gap-3 md:grid-cols-2">
            <article
              v-for="classItem in classes"
              :key="classItem.classId"
              class="rounded-[18px] border border-[#EBEBEB] bg-[#FBFAF8] p-4"
            >
              <div class="flex items-start justify-between gap-3">
                <div>
                  <div class="flex flex-wrap items-center gap-2">
                    <h3 class="text-sm font-medium text-[#111827]">{{ classItem.classTitle }}</h3>
                    <span class="rounded-lg bg-white px-2 py-1 text-[11px] font-medium text-[#6B7280]">
                      {{ classItem.classCode }}
                    </span>
                  </div>
                  <p class="mt-1 text-sm leading-6 text-[#6B7280]">
                    {{ classItem.classDesc || "Deskripsi belum ditambahkan" }}
                  </p>
                </div>
                <span
                  class="shrink-0 rounded-lg px-2 py-1 text-[11px] font-medium"
                  :class="classItem.isActive ? 'bg-[#ECFDF5] text-[#059669]' : 'bg-[#F9FAFB] text-[#6B7280]'"
                >
                  {{ classItem.isActive ? "Aktif" : "Nonaktif" }}
                </span>
              </div>

              <div class="mt-4 grid gap-2 text-xs text-[#6B7280] sm:grid-cols-2">
                <p>
                  Semester:
                  <span class="font-medium text-[#374151]">{{ classItem.termName || selectedTerm?.termName }}</span>
                </p>
                <p>
                  Tahun:
                  <span class="font-medium text-[#374151]">
                    {{ classItem.academicYearName || selectedAcademicYear?.academicYearName }}
                  </span>
                </p>
                <p>
                  Dibuat:
                  <span class="font-medium text-[#374151]">{{ formatDateTime(classItem.createdAt) }}</span>
                </p>
                <p>
                  Oleh:
                  <span class="font-medium text-[#374151]">{{ classItem.creatorName || "Pembuat tidak tersedia" }}</span>
                </p>
              </div>
            </article>
          </div>
        </div>
      </section>
    </section>
  </main>
</template>
