<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { RouterLink, useRoute } from "vue-router";
import {
  PhArrowLeft,
  PhBookOpen,
  PhClipboardText,
  PhFileText,
  PhPaperPlaneTilt,
  PhUsersThree,
  PhWarningCircle,
  PhCalendarBlank,
} from "@phosphor-icons/vue";
import { getMyTeachingSubjectClassById } from "../../services/teacherSubjects";
import { getSubjectAssignments } from "../../services/assignment";
import { getSubjectMaterials } from "../../services/teacherMaterial";
import type { TeacherSubjectClass } from "../../types/teacherSubjects";
import type { AssignmentItem } from "../../types/assignment";
import type { MaterialItem } from "../../types/teacherMaterial";
import { getSubjectColor } from "../../utils/color";

const route = useRoute();
const subjectClassId = computed(() =>
  String(route.params.subjectClassId ?? ""),
);

const subject = ref<TeacherSubjectClass | null>(null);
const materials = ref<MaterialItem[]>([]);
const assignments = ref<AssignmentItem[]>([]);
const loading = ref(false);
const errorMessage = ref("");

async function loadData() {
  loading.value = true;
  errorMessage.value = "";
  try {
    const [subjectData, materialData, assignmentData] = await Promise.all([
      getMyTeachingSubjectClassById(subjectClassId.value),
      getSubjectMaterials(subjectClassId.value),
      getSubjectAssignments(subjectClassId.value)
    ]);
    subject.value = subjectData;
    materials.value = materialData.data.data;
    assignments.value = assignmentData.data.data as AssignmentItem[];
  } catch {
    errorMessage.value =
      "Workspace subject belum bisa dimuat sepenuhnya. Coba lagi beberapa saat.";
  } finally {
    loading.value = false;
  }
}

function formatDate(dateStr?: string) {
  if (!dateStr) return "-";
  return new Date(dateStr).toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "short",
    year: "numeric"
  });
}

onMounted(loadData);
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-8 md:px-8 lg:px-10">
    <section class="mx-auto flex max-w-6xl flex-col gap-6">
      <RouterLink
        to="/teacher/subjects"
        class="inline-flex items-center gap-2 self-start text-sm font-medium text-[#6b6475] transition hover:text-[#171322]"
      >
        <PhArrowLeft :size="18" />
        Kembali ke subjects
      </RouterLink>

      <header
        class="rounded-4xl bg-white p-6 shadow-sm ring-1 ring-black/5 md:p-8"
      >
        <div
          class="mb-5 flex h-14 w-14 items-center justify-center rounded-2xl text-white"
          :style="{ backgroundColor: getSubjectColor(subjectClassId) }"
        >
          <PhBookOpen :size="28" weight="duotone" />
        </div>
        <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
          <div>
            <p class="text-sm font-medium text-[#7b61a8]">Subject workspace</p>
            <h1 class="mt-3 text-3xl font-medium text-[#171322]">
              {{
                subject?.subjectName ??
                (loading ? "Memuat subject..." : "Workspace subject")
              }}
            </h1>
          </div>
          <RouterLink
            v-if="subject"
            :to="`/teacher/subjects/${subjectClassId}/create`"
            class="inline-flex items-center gap-2 px-5 py-3 text-sm font-medium text-white bg-[#171322] rounded-2xl hover:bg-[#2f2b3a] transition"
          >
            <PhPaperPlaneTilt :size="18" weight="bold" />
            Tambah Konten
          </RouterLink>
        </div>
        <p class="mt-3 max-w-2xl text-sm leading-6 text-[#6b6475]">
          <span v-if="subject">
            {{ subject.className }} menjadi konteks class untuk subject ini.
            Material dan tugas berikutnya akan dibuat di level subject class.
          </span>
          <span v-else>
            Detail subject class mengambil data dari endpoint current teacher
            agar guru hanya melihat subject yang dia ampu.
          </span>
        </p>
      </header>

      <section
        v-if="errorMessage"
        class="rounded-[28px] bg-white p-6 shadow-sm ring-1 ring-black/5"
      >
        <div class="flex items-start gap-3">
          <PhWarningCircle
            :size="24"
            class="mt-0.5 text-[#e58f86]"
            weight="duotone"
          />
          <div>
            <h2 class="text-lg font-medium text-[#171322]">
              Gagal memuat workspace
            </h2>
            <p class="mt-2 text-sm leading-6 text-[#6b6475]">
              {{ errorMessage }}
            </p>
          </div>
        </div>
      </section>

      <section
        v-else-if="!loading && !subject"
        class="rounded-[28px] bg-white p-6 shadow-sm ring-1 ring-black/5"
      >
        <h2 class="text-lg font-medium text-[#171322]">
          Subject tidak ditemukan
        </h2>
        <p class="mt-2 text-sm leading-6 text-[#6b6475]">
          Subject class ini tidak tersedia untuk akun guru pada school aktif.
        </p>
      </section>

      <template v-if="subject">
        <section class="grid gap-4 md:grid-cols-4">
          <article class="rounded-3xl bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhUsersThree :size="24" class="text-[#74bfa5]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Siswa</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">{{ subject.studentCount }}</p>
          </article>
          <article class="rounded-3xl bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhFileText :size="24" class="text-[#7aa7d9]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Materi</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">{{ subject.materialCount }}</p>
          </article>
          <article class="rounded-3xl bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhClipboardText :size="24" class="text-[#e58f86]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Tugas</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">{{ subject.assignmentCount }}</p>
          </article>
          <article class="rounded-3xl bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhWarningCircle :size="24" class="text-[#b889c9]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Perlu review</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">{{ subject.pendingSubmissions }}</p>
          </article>
        </section>

        <div class="grid gap-6 lg:grid-cols-2">
          <!-- Materials List -->
          <section class="bg-white rounded-4xl p-6 border border-[#EBEBEB] shadow-sm">
            <div class="flex items-center justify-between mb-6">
              <h2 class="text-xl font-semibold text-[#171322] flex items-center gap-2">
                <PhFileText :size="24" class="text-[#7aa7d9]" weight="duotone" />
                Materi
              </h2>
              <span class="text-xs font-medium text-[#6B7280] bg-[#F3F4F6] px-2.5 py-1 rounded-full">
                {{ materials.length }} Total
              </span>
            </div>

            <div v-if="materials.length === 0" class="py-12 flex flex-col items-center justify-center text-center">
              <div class="w-16 h-16 bg-[#F9FAFB] rounded-full flex items-center justify-center mb-4">
                <PhFileText :size="32" class="text-[#D1D5DB]" />
              </div>
              <p class="text-sm text-[#6B7280]">Belum ada materi yang diterbitkan.</p>
            </div>

            <div v-else class="space-y-3">
              <div v-for="mat in materials" :key="mat.materialId" class="group flex items-center justify-between p-4 bg-[#F9FAFB] rounded-2xl border border-transparent hover:border-[#4F46E5]/10 hover:bg-white transition shadow-sm hover:shadow-md">
                <div class="flex items-center gap-4 overflow-hidden">
                  <div class="w-10 h-10 shrink-0 bg-white rounded-xl flex items-center justify-center border border-[#EBEBEB] text-[#7aa7d9]">
                    <PhFileText :size="20" weight="duotone" />
                  </div>
                  <div class="min-w-0">
                    <h3 class="text-sm font-semibold text-[#171322] truncate">{{ mat.materialTitle }}</h3>
                    <p class="text-[11px] text-[#8a8494] mt-0.5 uppercase">{{ mat.materialType }} • {{ formatDate(mat.createdAt) }}</p>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <!-- Assignments List -->
          <section class="bg-white rounded-4xl p-6 border border-[#EBEBEB] shadow-sm">
            <div class="flex items-center justify-between mb-6">
              <h2 class="text-xl font-semibold text-[#171322] flex items-center gap-2">
                <PhClipboardText :size="24" class="text-[#e58f86]" weight="duotone" />
                Tugas
              </h2>
              <span class="text-xs font-medium text-[#6B7280] bg-[#F3F4F6] px-2.5 py-1 rounded-full">
                {{ assignments.length }} Total
              </span>
            </div>

            <div v-if="assignments.length === 0" class="py-12 flex flex-col items-center justify-center text-center">
              <div class="w-16 h-16 bg-[#F9FAFB] rounded-full flex items-center justify-center mb-4">
                <PhClipboardText :size="32" class="text-[#D1D5DB]" />
              </div>
              <p class="text-sm text-[#6B7280]">Belum ada tugas yang diberikan.</p>
            </div>

            <div v-else class="space-y-3">
              <div v-for="asg in assignments" :key="asg.assignmentId" class="group p-4 bg-[#F9FAFB] rounded-2xl border border-transparent hover:border-[#4F46E5]/10 hover:bg-white transition shadow-sm hover:shadow-md">
                <div class="flex items-start justify-between gap-4">
                  <div class="flex items-center gap-4 overflow-hidden">
                    <div class="w-10 h-10 shrink-0 bg-white rounded-xl flex items-center justify-center border border-[#EBEBEB] text-[#e58f86]">
                      <PhClipboardText :size="20" weight="duotone" />
                    </div>
                    <div class="min-w-0">
                      <h3 class="text-sm font-semibold text-[#171322] truncate">{{ asg.assignmentTitle }}</h3>
                      <div class="flex items-center gap-3 mt-1">
                        <span class="flex items-center gap-1 text-[10px] text-[#8a8494] font-medium uppercase">
                          <PhCalendarBlank :size="12" />
                          {{ formatDate(asg.deadline) }}
                        </span>
                      </div>
                    </div>
                  </div>
                  <RouterLink 
                    :to="`/teacher/assignments/${asg.assignmentId}/review`"
                    class="shrink-0 px-3 py-1.5 text-[11px] font-bold text-[#4F46E5] bg-[#EEF2FF] rounded-lg hover:bg-[#4F46E5] hover:text-white transition"
                  >
                    Review
                  </RouterLink>
                </div>
              </div>
            </div>
          </section>
        </div>
      </template>
    </section>
  </main>
</template>
