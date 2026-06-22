<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import {
  PhArrowRight,
  PhCheckCircle,
  PhClipboardText,
  PhClock,
  PhWarningCircle,
} from "@phosphor-icons/vue";
import { getSubjectClassSubmissions } from "../../services/teacherAssignment";
import { getMyTeachingSubjectClasses } from "../../services/teacherSubjects";
import type {
  TeacherSubmissionGroup,
  TeacherSubjectClassSubmissionsResponse,
} from "../../types/teacherAssignment";
import type { TeacherSubjectClass } from "../../types/teacherSubjects";
import { formatDate } from "../../utils/date";

type InboxFilter = "all" | "pending" | "graded";

interface SubmissionInboxItem {
  subjectClassId: string;
  subjectName: string;
  subjectCode?: string;
  className: string;
  classCode?: string;
  assignment: TeacherSubmissionGroup["assignment"];
  submissionCount: number;
  gradedCount: number;
  pendingCount: number;
  lateCount: number;
}

const loading = ref(false);
const errorMessage = ref("");
const subjects = ref<TeacherSubjectClass[]>([]);
const inboxItems = ref<SubmissionInboxItem[]>([]);
const activeFilter = ref<InboxFilter>("all");

const summary = computed(() =>
  inboxItems.value.reduce(
    (total, item) => ({
      submissions: total.submissions + item.submissionCount,
      pending: total.pending + item.pendingCount,
      graded: total.graded + item.gradedCount,
      late: total.late + item.lateCount,
    }),
    { submissions: 0, pending: 0, graded: 0, late: 0 },
  ),
);

const filterTabs = computed(() => [
  { id: "all" as const, label: "Semua", count: inboxItems.value.length },
  {
    id: "pending" as const,
    label: "Perlu review",
    count: inboxItems.value.filter((item) => item.pendingCount > 0).length,
  },
  {
    id: "graded" as const,
    label: "Sudah dinilai",
    count: inboxItems.value.filter(
      (item) => item.submissionCount > 0 && item.pendingCount === 0,
    ).length,
  },
]);

const filteredItems = computed(() => {
  const items = inboxItems.value.filter((item) => {
    if (activeFilter.value === "pending") return item.pendingCount > 0;
    if (activeFilter.value === "graded") {
      return item.submissionCount > 0 && item.pendingCount === 0;
    }
    return true;
  });

  return [...items].sort(compareInboxItems);
});

function compareInboxItems(a: SubmissionInboxItem, b: SubmissionInboxItem) {
  const pendingDiff = Number(b.pendingCount > 0) - Number(a.pendingCount > 0);
  if (pendingDiff !== 0) return pendingDiff;

  const aDeadline = getDeadlineTime(a.assignment.deadline);
  const bDeadline = getDeadlineTime(b.assignment.deadline);
  if (aDeadline !== bDeadline) return aDeadline - bDeadline;

  return (a.assignment.assignmentTitle || "").localeCompare(
    b.assignment.assignmentTitle || "",
  );
}

function getDeadlineTime(deadline?: string) {
  if (!deadline) return Number.MAX_SAFE_INTEGER;
  const value = new Date(deadline).getTime();
  return Number.isNaN(value) ? Number.MAX_SAFE_INTEGER : value;
}

function buildInboxItem(
  subject: TeacherSubjectClass,
  response: TeacherSubjectClassSubmissionsResponse,
  group: TeacherSubmissionGroup,
): SubmissionInboxItem {
  const lateCount = group.submissions.filter((submission) => submission.isLate).length;
  return {
    subjectClassId: subject.subjectClassId,
    subjectName:
      response.subjectClass.subjectName || subject.subjectName || "Subject",
    subjectCode: response.subjectClass.subjectCode || subject.subjectCode,
    className: subject.className,
    classCode: subject.classCode,
    assignment: group.assignment,
    submissionCount: group.submissionCount,
    gradedCount: group.gradedCount,
    pendingCount: group.pendingCount,
    lateCount,
  };
}

async function loadInbox() {
  loading.value = true;
  errorMessage.value = "";
  subjects.value = [];
  inboxItems.value = [];

  try {
    const subjectList = await getMyTeachingSubjectClasses();
    subjects.value = subjectList;

    if (subjectList.length === 0) return;

    const responses = await Promise.all(
      subjectList.map(async (subject) => ({
        subject,
        data: await getSubjectClassSubmissions(subject.subjectClassId),
      })),
    );

    inboxItems.value = responses.flatMap(({ subject, data }) =>
      (data.assignments ?? [])
        .filter((group) => group.submissionCount > 0 || group.submissions.length > 0)
        .map((group) => buildInboxItem(subject, data, group)),
    );
  } catch {
    errorMessage.value =
      "Inbox pengumpulan belum bisa dimuat. Coba lagi beberapa saat.";
  } finally {
    loading.value = false;
  }
}

onMounted(loadInbox);
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-5 sm:px-6 lg:px-8">
    <section class="flex w-full max-w-none flex-col gap-5">
      <header
        class="rounded-[22px] bg-[#f0e9dd] px-5 py-5 shadow-sm ring-1 ring-black/5 md:px-6"
      >
        <p class="text-sm font-medium text-[#8a6d3b]">Pengumpulan siswa</p>
        <h1 class="mt-3 text-3xl font-medium text-[#171322] md:text-4xl">
          Inbox pengumpulan
        </h1>
        <p class="mt-3 max-w-2xl text-sm leading-6 text-[#6b6475]">
          Pantau pengumpulan dari semua subject yang diajar. Proses nilai dan
          feedback tetap dilakukan dari halaman review tugas.
        </p>
      </header>

      <section v-if="loading" class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
        <p class="text-sm text-[#6b6475]">Memuat inbox pengumpulan...</p>
      </section>

      <section
        v-else-if="errorMessage"
        class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5"
      >
        <div
          class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
        >
          <div class="flex items-start gap-3">
            <PhWarningCircle
              :size="24"
              class="mt-0.5 text-[#e58f86]"
              weight="duotone"
            />
            <div>
              <h2 class="text-lg font-medium text-[#171322]">
                Gagal memuat inbox
              </h2>
              <p class="mt-2 text-sm leading-6 text-[#6b6475]">
                {{ errorMessage }}
              </p>
            </div>
          </div>
          <button
            type="button"
            class="rounded-2xl bg-[#171322] px-4 py-3 text-sm font-medium text-white"
            @click="loadInbox"
          >
            Coba lagi
          </button>
        </div>
      </section>

      <template v-else>
        <section class="grid gap-4 md:grid-cols-4">
          <article class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhClipboardText :size="24" class="text-[#7aa7d9]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Total submission</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">
              {{ summary.submissions }}
            </p>
          </article>
          <article class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhWarningCircle :size="24" class="text-[#e58f86]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Perlu review</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">
              {{ summary.pending }}
            </p>
          </article>
          <article class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhCheckCircle :size="24" class="text-[#74bfa5]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Sudah dinilai</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">
              {{ summary.graded }}
            </p>
          </article>
          <article class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <PhClock :size="24" class="text-[#b889c9]" weight="duotone" />
            <p class="mt-4 text-sm text-[#8a8494]">Terlambat</p>
            <p class="mt-1 text-2xl font-medium text-[#171322]">
              {{ summary.late }}
            </p>
          </article>
        </section>

        <section class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5">
          <div class="flex flex-col gap-4 border-b border-[#ece8df] pb-4 lg:flex-row lg:items-end lg:justify-between">
            <div>
              <p class="text-sm font-medium text-[#171322]">
                Daftar assignment dengan pengumpulan
              </p>
              <p class="mt-1 text-sm text-[#8a8494]">
                {{ subjects.length }} subject diajar dalam school aktif.
              </p>
            </div>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="tab in filterTabs"
                :key="tab.id"
                type="button"
                class="rounded-2xl px-4 py-2.5 text-sm font-medium transition"
                :class="
                  activeFilter === tab.id
                    ? 'bg-[#171322] text-white'
                    : 'bg-[#faf8f4] text-[#6b6475] hover:bg-[#f0e9dd] hover:text-[#171322]'
                "
                @click="activeFilter = tab.id"
              >
                {{ tab.label }}
                <span class="ml-2 opacity-70">{{ tab.count }}</span>
              </button>
            </div>
          </div>

          <div v-if="subjects.length === 0" class="py-10 text-center">
            <PhClipboardText
              :size="34"
              class="mx-auto text-[#b5afbf]"
              weight="duotone"
            />
            <h2 class="mt-3 text-lg font-medium text-[#171322]">
              Belum ada subject yang diajar
            </h2>
            <p class="mx-auto mt-2 max-w-xl text-sm leading-6 text-[#6b6475]">
              Pengumpulan akan tampil setelah admin menugaskan teacher ke
              subject class dan siswa mulai mengumpulkan tugas.
            </p>
          </div>

          <div v-else-if="filteredItems.length === 0" class="py-10 text-center">
            <PhCheckCircle
              :size="34"
              class="mx-auto text-[#b5afbf]"
              weight="duotone"
            />
            <h2 class="mt-3 text-lg font-medium text-[#171322]">
              Belum ada pengumpulan
            </h2>
            <p class="mx-auto mt-2 max-w-xl text-sm leading-6 text-[#6b6475]">
              Tidak ada assignment yang sesuai dengan filter saat ini.
            </p>
          </div>

          <div v-else class="space-y-3 pt-5">
            <article
              v-for="item in filteredItems"
              :key="`${item.subjectClassId}-${item.assignment.assignmentId}`"
              class="rounded-[18px] bg-[#faf8f4] p-5 ring-1 ring-black/5"
            >
              <div class="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
                <div class="min-w-0">
                  <div class="flex flex-wrap gap-2 text-xs font-medium">
                    <span class="rounded-2xl bg-white px-3 py-1.5 text-[#4f46e5]">
                      {{ item.subjectName }}
                    </span>
                    <span
                      v-if="item.subjectCode"
                      class="rounded-2xl bg-white px-3 py-1.5 text-[#6b6475]"
                    >
                      {{ item.subjectCode }}
                    </span>
                    <span class="rounded-2xl bg-white px-3 py-1.5 text-[#6b6475]">
                      {{ item.className || item.classCode || "Kelas" }}
                    </span>
                  </div>

                  <h2 class="mt-4 text-lg font-medium text-[#171322]">
                    {{ item.assignment.assignmentTitle }}
                  </h2>
                  <p
                    v-if="item.assignment.categoryName || item.assignment.deadline"
                    class="mt-2 text-sm text-[#6b6475]"
                  >
                    <span v-if="item.assignment.categoryName">
                      {{ item.assignment.categoryName }}
                    </span>
                    <span v-if="item.assignment.categoryName && item.assignment.deadline">
                      ·
                    </span>
                    <span v-if="item.assignment.deadline">
                      Deadline {{ formatDate(item.assignment.deadline) }}
                    </span>
                  </p>
                </div>

                <RouterLink
                  :to="{
                    name: 'teacher-assignment-review',
                    params: { assignmentId: item.assignment.assignmentId },
                  }"
                  class="inline-flex items-center justify-center gap-2 rounded-2xl bg-[#171322] px-4 py-3 text-sm font-medium text-white transition hover:bg-[#2f2b3a]"
                >
                  Review pengumpulan
                  <PhArrowRight :size="16" />
                </RouterLink>
              </div>

              <div class="mt-5 grid gap-3 sm:grid-cols-4">
                <div class="rounded-2xl bg-white p-4">
                  <p class="text-xs text-[#8a8494]">Submission</p>
                  <p class="mt-1 text-xl font-medium text-[#171322]">
                    {{ item.submissionCount }}
                  </p>
                </div>
                <div class="rounded-2xl bg-white p-4">
                  <p class="text-xs text-[#8a8494]">Perlu review</p>
                  <p class="mt-1 text-xl font-medium text-[#171322]">
                    {{ item.pendingCount }}
                  </p>
                </div>
                <div class="rounded-2xl bg-white p-4">
                  <p class="text-xs text-[#8a8494]">Sudah dinilai</p>
                  <p class="mt-1 text-xl font-medium text-[#171322]">
                    {{ item.gradedCount }}
                  </p>
                </div>
                <div class="rounded-2xl bg-white p-4">
                  <p class="text-xs text-[#8a8494]">Terlambat</p>
                  <p class="mt-1 text-xl font-medium text-[#171322]">
                    {{ item.lateCount }}
                  </p>
                </div>
              </div>
            </article>
          </div>
        </section>
      </template>
    </section>
  </main>
</template>
