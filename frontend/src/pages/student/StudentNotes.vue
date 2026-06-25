<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import {
  PhArrowRight,
  PhBookOpen,
  PhNotebook,
  PhWarningCircle,
} from "@phosphor-icons/vue";
import { getStudentNotes } from "../../services/studentNotes";
import type { StudentGlobalMaterialNote } from "../../types/studentNotes";
import { formatDateTime } from "../../utils/date";

interface NoteGroup {
  key: string;
  className: string;
  classCode: string;
  subjectName: string;
  subjectCode: string;
  latestUpdatedAt: string;
  notes: StudentGlobalMaterialNote[];
}

const notes = ref<StudentGlobalMaterialNote[]>([]);
const isLoading = ref(true);
const errorMessage = ref("");

const groupedNotes = computed<NoteGroup[]>(() => {
  const groups = new Map<string, NoteGroup>();

  for (const note of notes.value) {
    const key = `${note.classId}:${note.subjectId}`;
    const existing = groups.get(key);

    if (existing) {
      existing.notes.push(note);
      if (getTime(note.updatedAt) > getTime(existing.latestUpdatedAt)) {
        existing.latestUpdatedAt = note.updatedAt;
      }
      continue;
    }

    groups.set(key, {
      key,
      className: note.className,
      classCode: note.classCode,
      subjectName: note.subjectName,
      subjectCode: note.subjectCode,
      latestUpdatedAt: note.updatedAt,
      notes: [note],
    });
  }

  return [...groups.values()]
    .map((group) => ({
      ...group,
      notes: [...group.notes].sort(
        (a, b) => getTime(b.updatedAt) - getTime(a.updatedAt),
      ),
    }))
    .sort(
      (a, b) => getTime(b.latestUpdatedAt) - getTime(a.latestUpdatedAt),
    );
});

async function loadNotes() {
  isLoading.value = true;
  errorMessage.value = "";

  try {
    const response = await getStudentNotes();
    notes.value = response.notes;
  } catch {
    errorMessage.value =
      "Catatan belum bisa dimuat. Periksa koneksi atau coba lagi nanti.";
  } finally {
    isLoading.value = false;
  }
}

function getTime(value?: string | null) {
  if (!value) return 0;
  const time = new Date(value).getTime();
  return Number.isNaN(time) ? 0 : time;
}

onMounted(loadNotes);
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-6 sm:px-8 lg:px-10">
    <header class="mb-6">
      <p class="text-sm font-medium text-[#4f46e5]">Ruang belajar pribadi</p>
      <h1 class="mt-2 text-3xl font-medium text-[#171322]">Catatan Saya</h1>
      <p class="mt-2 max-w-2xl text-sm leading-6 text-[#6b6475]">
        Kumpulan catatan dari materi yang masih dapat kamu akses di sekolah
        aktif.
      </p>
    </header>

    <section v-if="isLoading" class="space-y-5">
      <article
        v-for="item in 3"
        :key="item"
        class="animate-pulse rounded-[18px] border border-[#ebe7df] bg-white p-5"
      >
        <div class="h-5 w-48 rounded bg-[#f1efeb]" />
        <div class="mt-3 h-4 w-32 rounded bg-[#f1efeb]" />
        <div class="mt-5 grid gap-3 lg:grid-cols-2">
          <div class="h-40 rounded-xl bg-[#f8f7f4]" />
          <div class="h-40 rounded-xl bg-[#f8f7f4]" />
        </div>
      </article>
    </section>

    <section
      v-else-if="errorMessage"
      class="rounded-[18px] border border-[#f1d6d3] bg-white p-6"
    >
      <div class="flex items-start gap-3">
        <PhWarningCircle
          :size="24"
          class="mt-0.5 shrink-0 text-[#dc2626]"
          weight="duotone"
        />
        <div>
          <h2 class="text-base font-medium text-[#171322]">
            Catatan tidak dapat dimuat
          </h2>
          <p class="mt-1 text-sm leading-6 text-[#6b6475]">
            {{ errorMessage }}
          </p>
          <button
            class="mt-4 rounded-lg bg-[#4f46e5] px-4 py-2 text-sm font-medium text-white transition hover:bg-[#4338ca]"
            type="button"
            @click="loadNotes"
          >
            Coba lagi
          </button>
        </div>
      </div>
    </section>

    <section
      v-else-if="notes.length === 0"
      class="rounded-[18px] border border-[#ebe7df] bg-white p-8 text-center"
    >
      <div
        class="mx-auto flex h-12 w-12 items-center justify-center rounded-2xl bg-[#eef2ff] text-[#4f46e5]"
      >
        <PhNotebook :size="24" weight="duotone" />
      </div>
      <h2 class="mt-4 text-base font-medium text-[#171322]">
        Belum ada catatan
      </h2>
      <p class="mx-auto mt-2 max-w-lg text-sm leading-6 text-[#6b6475]">
        Belum ada catatan. Buka materi dan tulis catatan untuk mulai membangun
        ruang belajarmu.
      </p>
      <RouterLink
        class="mt-5 inline-flex items-center gap-2 rounded-lg bg-[#4f46e5] px-4 py-2 text-sm font-medium text-white transition hover:bg-[#4338ca]"
        to="/student/subjects"
      >
        Buka mata pelajaran
        <PhArrowRight :size="16" />
      </RouterLink>
    </section>

    <section v-else class="space-y-5">
      <section
        v-for="group in groupedNotes"
        :key="group.key"
        class="border-b border-[#ebe7df] pb-6 last:border-b-0"
      >
        <header
          class="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between"
        >
          <div>
            <div class="flex flex-wrap items-center gap-2">
              <h2 class="text-base font-medium text-[#171322]">
                {{ group.subjectName || "Mata pelajaran" }}
              </h2>
              <span
                v-if="group.subjectCode"
                class="rounded-full bg-[#eef2ff] px-2.5 py-1 text-xs font-medium text-[#4f46e5]"
              >
                {{ group.subjectCode }}
              </span>
            </div>
            <p class="mt-1 text-sm text-[#6b6475]">
              {{ group.className || "Kelas" }}
              <span v-if="group.classCode">· {{ group.classCode }}</span>
            </p>
          </div>
          <p class="text-xs text-[#9ca3af]">
            {{ group.notes.length }} catatan
          </p>
        </header>

        <div class="mt-4 grid gap-3 lg:grid-cols-2">
          <article
            v-for="note in group.notes"
            :key="note.noteId"
            class="flex min-w-0 flex-col rounded-xl border border-[#ebe7df] bg-white p-4"
          >
            <div class="flex items-start gap-3">
              <div
                class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-white text-[#4f46e5]"
              >
                <PhBookOpen :size="20" weight="duotone" />
              </div>
              <div class="min-w-0">
                <h3 class="text-sm font-medium text-[#171322]">
                  {{ note.materialTitle }}
                </h3>
                <p class="mt-1 text-xs text-[#9ca3af]">
                  Diperbarui {{ formatDateTime(note.updatedAt) }}
                </p>
              </div>
            </div>

            <p
              class="mt-4 line-clamp-5 whitespace-pre-line break-words text-sm leading-6 text-[#4f4858]"
            >
              {{ note.content }}
            </p>

            <div class="mt-auto flex flex-wrap gap-2 pt-5">
              <RouterLink
                class="rounded-lg border border-[#ddd8e4] px-3 py-2 text-xs font-medium text-[#6b6475] transition hover:bg-white"
                :to="`/student/subjects/${note.subjectClassId}/materials/${note.materialId}`"
              >
                Lihat materi
              </RouterLink>
              <RouterLink
                class="inline-flex items-center gap-2 rounded-lg bg-[#4f46e5] px-3 py-2 text-xs font-medium text-white transition hover:bg-[#4338ca]"
                :to="`/student/subjects/${note.subjectClassId}/materials/${note.materialId}/note`"
              >
                Buka catatan
                <PhArrowRight :size="14" />
              </RouterLink>
            </div>
          </article>
        </div>
      </section>
    </section>
  </main>
</template>
