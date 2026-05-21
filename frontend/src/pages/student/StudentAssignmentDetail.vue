<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import {
  PhArrowLeft,
  PhCalendarBlank,
  PhClipboardText,
  PhFileText,
  PhWarningCircle,
} from '@phosphor-icons/vue'
import { getSubjectAssignmentDetail } from '../../services/assignment'
import type { AssignmentItem, SubjectClassHeader } from '../../types/assignment'

const route = useRoute()
const subjectClassId = computed(() => String(route.params.sclId ?? ''))
const assignmentId = computed(() => String(route.params.asgId ?? ''))
const assignment = ref<AssignmentItem | null>(null)
const subjectClass = ref<SubjectClassHeader | null>(null)
const isLoading = ref(true)
const errorMessage = ref('')
const didLoad = ref(false)

async function loadAssignment() {
  if (!subjectClassId.value || !assignmentId.value) {
    isLoading.value = false
    errorMessage.value = 'Konteks tugas tidak lengkap.'
    return
  }

  isLoading.value = true
  errorMessage.value = ''
  didLoad.value = false

  try {
    const data = await getSubjectAssignmentDetail(subjectClassId.value, assignmentId.value)
    subjectClass.value = data.subjectClass
    assignment.value = data.assignment
    didLoad.value = true
  } catch {
    errorMessage.value = 'Detail tugas belum bisa dimuat. Periksa koneksi atau coba lagi nanti.'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadAssignment)
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-6 sm:px-8 lg:px-10">
    <RouterLink
      class="mb-6 inline-flex items-center gap-2 rounded-md bg-white px-4 py-2 text-sm font-medium text-[#4f46e5] transition hover:bg-[#eef2ff]"
      :to="`/student/subjects/${subjectClassId}`"
    >
      <PhArrowLeft :size="18" />
      Kembali ke subject
    </RouterLink>

    <section v-if="isLoading" class="max-w-3xl space-y-3">
      <div class="h-40 animate-pulse rounded-3xl border border-[#ebe7df] bg-white" />
      <div class="h-28 animate-pulse rounded-3xl border border-[#ebe7df] bg-white" />
    </section>

    <section v-else-if="errorMessage" class="soft-card max-w-3xl rounded-3xl p-6">
      <div class="mb-4 flex h-11 w-11 items-center justify-center rounded-2xl bg-[#fff1f0] text-[#f2756a]">
        <PhWarningCircle :size="24" weight="duotone" />
      </div>
      <p class="text-sm font-medium text-[#171322]">Tidak bisa memuat tugas</p>
      <p class="mt-2 text-sm leading-6 text-[#7a7385]">{{ errorMessage }}</p>
      <button
        class="mt-5 rounded-2xl bg-[#4f46e5] px-4 py-2 text-sm font-medium text-white"
        type="button"
        @click="loadAssignment"
      >
        Coba lagi
      </button>
    </section>

    <section v-else-if="didLoad && !assignment" class="soft-card max-w-3xl rounded-3xl p-6">
      <div class="mb-4 flex h-11 w-11 items-center justify-center rounded-2xl bg-[#eef2ff] text-[#4f46e5]">
        <PhClipboardText :size="24" weight="duotone" />
      </div>
      <p class="text-sm font-medium text-[#171322]">Tugas tidak ditemukan</p>
      <p class="mt-2 text-sm leading-6 text-[#7a7385]">
        Assignment ID ini tidak ditemukan pada subject class yang sedang dibuka.
      </p>
    </section>

    <section v-else-if="assignment" class="max-w-3xl space-y-4">
      <article class="soft-card rounded-3xl p-6">
        <div class="mb-6 flex items-start gap-4">
          <div
            class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-[#eef2ff] text-[#4f46e5]"
          >
            <PhClipboardText :size="24" weight="duotone" />
          </div>
          <div class="min-w-0">
            <p class="text-sm text-[#7a7385]">
              {{ subjectClass?.subjectName || subjectClass?.subjectCode || 'Subject assignment' }}
            </p>
            <h1 class="mt-2 text-3xl font-medium tracking-normal text-[#171322]">
              {{ assignment.assignmentTitle }}
            </h1>
            <p v-if="assignment.categoryName" class="mt-2 text-sm text-[#4f46e5]">
              {{ assignment.categoryName }}
            </p>
          </div>
        </div>

        <div class="grid gap-3 sm:grid-cols-3">
          <div class="rounded-2xl bg-[#fbfaf8] p-4">
            <div class="mb-2 flex items-center gap-2 text-[#4f46e5]">
              <PhCalendarBlank :size="17" />
              <p class="text-xs font-medium">Deadline</p>
            </div>
            <p class="text-sm text-[#3f3a4a]">
              {{ assignment.deadline || 'Belum tersedia' }}
            </p>
          </div>
          <div class="rounded-2xl bg-[#fbfaf8] p-4">
            <p class="text-xs font-medium text-[#7a7385]">Late submission</p>
            <p class="mt-2 text-sm text-[#3f3a4a]">
              {{ assignment.allowLateSubmission ? 'Diizinkan' : 'Tidak diizinkan' }}
            </p>
          </div>
          <div class="rounded-2xl bg-[#fbfaf8] p-4">
            <p class="text-xs font-medium text-[#7a7385]">Dibuat</p>
            <p class="mt-2 text-sm text-[#3f3a4a]">
              {{ assignment.createdAt || 'Belum tersedia' }}
            </p>
          </div>
        </div>

        <div class="mt-6 rounded-2xl bg-white p-4">
          <p class="text-sm font-medium text-[#171322]">Deskripsi</p>
          <p
            v-if="assignment.assignmentDescription"
            class="mt-3 whitespace-pre-line text-sm leading-6 text-[#6b6475]"
          >
            {{ assignment.assignmentDescription }}
          </p>
          <p v-else class="mt-3 text-sm leading-6 text-[#7a7385]">
            Deskripsi tugas belum tersedia.
          </p>
        </div>
      </article>

      <article v-if="assignment.attachments?.length" class="rounded-3xl border border-[#ebe7df] bg-white p-5">
        <p class="text-sm font-medium text-[#171322]">Lampiran</p>
        <div class="mt-3 space-y-2">
          <a
            v-for="attachment in assignment.attachments"
            :key="attachment.mediaId"
            class="flex items-center gap-3 rounded-2xl bg-[#fbfaf8] px-4 py-3 text-sm text-[#4a4356]"
            :href="attachment.fileUrl"
            rel="noreferrer"
            target="_blank"
          >
            <PhFileText :size="18" class="text-[#4f46e5]" />
            <span class="truncate">{{ attachment.mediaName || 'Lampiran tugas' }}</span>
          </a>
        </div>
      </article>

      <article class="rounded-3xl border border-[#ebe7df] bg-white p-5">
        <p class="text-sm font-medium text-[#171322]">Pengumpulan tugas</p>
        <p class="mt-2 text-sm leading-6 text-[#7a7385]">
          Pengumpulan tugas akan tersedia pada tahap berikutnya.
        </p>
        <button
          class="mt-4 rounded-2xl bg-[#d8d5dd] px-4 py-2 text-sm font-medium text-white"
          disabled
          type="button"
        >
          Submit belum tersedia
        </button>
      </article>
    </section>
  </main>
</template>
