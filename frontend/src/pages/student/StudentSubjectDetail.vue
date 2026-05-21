<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import {
  PhBookOpen,
  PhChatCircleText,
  PhClipboardText,
  PhFileText,
  PhNotebook,
  PhWarningCircle,
} from '@phosphor-icons/vue'
import { getSubjectMaterials } from '../../services/classWorkspace'
import type { MaterialItem } from '../../types/classWorkspace'

const route = useRoute()
const subjectClassId = computed(() => String(route.params.sclId ?? ''))
const activeTab = ref('materials')
const materials = ref<MaterialItem[]>([])
const subjectTitle = ref(String(route.query.title ?? 'Detail Mata Pelajaran'))
const teacherName = ref('')
const isLoading = ref(true)
const errorMessage = ref('')

const tabs = [
  {
    key: 'materials',
    label: 'Materi',
    icon: PhBookOpen,
  },
  {
    key: 'assignments',
    label: 'Tugas',
    icon: PhClipboardText,
  },
  {
    key: 'notes',
    label: 'Catatan',
    icon: PhNotebook,
  },
]

const currentTab = computed(() => tabs.find((tab) => tab.key === activeTab.value) ?? tabs[0])

async function loadSubject() {
  if (!subjectClassId.value) {
    isLoading.value = false
    errorMessage.value = 'Subject class ID tidak tersedia.'
    return
  }

  isLoading.value = true
  errorMessage.value = ''

  try {
    const data = await getSubjectMaterials(subjectClassId.value)
    materials.value = data.materials
    subjectTitle.value =
      data.subjectClass.subjectName || data.subjectClass.subjectCode || subjectTitle.value
    teacherName.value = data.subjectClass.teacherName || ''
  } catch {
    errorMessage.value =
      'Detail mata pelajaran belum bisa dimuat. Periksa koneksi atau coba lagi nanti.'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadSubject)
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-6 sm:px-8 lg:px-10">
    <header class="mb-6 flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
      <div>
        <p class="text-sm text-[#7a7385]">Subject workspace</p>
        <h1 class="mt-2 text-3xl font-medium tracking-normal text-[#171322]">
          {{ subjectTitle }}
        </h1>
        <p class="mt-3 max-w-2xl text-sm leading-6 text-[#7a7385]">
          Materi dan tugas berada di konteks subject. Feed kelas tersedia terpisah sebagai komunikasi
          level class.
        </p>
      </div>

      <button
        class="inline-flex items-center gap-2 rounded-2xl border border-[#ebe7df] bg-white px-4 py-2 text-sm font-medium text-[#3f3a4a]"
        type="button"
      >
        <PhChatCircleText :size="18" />
        Chat subject sedang dikembangkan
      </button>
    </header>

    <section class="soft-card overflow-hidden rounded-md">
      <div class="border-b border-[#ebe7df] bg-white/80 px-4 py-3">
        <div class="flex flex-wrap gap-2">
          <button
            v-for="tab in tabs"
            :key="tab.key"
            class="flex h-10 items-center gap-2 px-4 text-sm transition"
            :class="
              activeTab === tab.key
                ? 'border-b-2 border-[#4f46e5] bg-[#eef2ff]/30 font-medium text-[#4f46e5]'
                : 'text-[#7a7385] hover:bg-[#f8f7f4] hover:text-[#3f3a4a]'
            "
            type="button"
            @click="activeTab = tab.key"
          >
            <component :is="tab.icon" :size="17" />
            {{ tab.label }}
          </button>
        </div>
      </div>

      <div class="grid gap-6 p-6 lg:grid-cols-[1fr_260px]">
        <article class="rounded-md bg-[#fbfaf8] p-6">
          <div
            class="mb-5 flex h-12 w-12 items-center justify-center rounded-md bg-[#eef2ff] text-[#4f46e5]"
          >
            <component :is="currentTab.icon" :size="24" weight="duotone" />
          </div>

          <template v-if="activeTab === 'materials'">
            <p class="text-sm font-medium text-[#4f46e5]">Materi</p>
            <h2 class="mt-3 text-2xl font-medium text-[#171322]">Materi subject</h2>

            <div v-if="isLoading" class="mt-6 space-y-3">
              <div
                v-for="item in 3"
                :key="item"
                class="h-20 animate-pulse rounded-md bg-white"
              />
            </div>

            <div v-else-if="errorMessage" class="mt-6 rounded-md bg-white p-4">
              <div class="flex items-start gap-3">
                <PhWarningCircle
                  :size="22"
                  class="mt-0.5 text-[#f2756a]"
                  weight="duotone"
                />
                <div>
                  <p class="text-sm font-medium text-[#171322]">Tidak bisa memuat materi</p>
                  <p class="mt-1 text-sm leading-6 text-[#7a7385]">{{ errorMessage }}</p>
                  <button
                    class="mt-4 rounded-md bg-[#4f46e5] px-4 py-2 text-sm font-medium text-white"
                    type="button"
                    @click="loadSubject"
                  >
                    Coba lagi
                  </button>
                </div>
              </div>
            </div>

            <div v-else-if="materials.length === 0" class="mt-6 rounded-md bg-white p-4">
              <p class="text-sm font-medium text-[#171322]">Belum ada materi</p>
              <p class="mt-2 text-sm leading-6 text-[#7a7385]">
                Materi akan tampil setelah guru menambahkan konten pada subject ini.
              </p>
            </div>

            <div v-else class="mt-6 space-y-3">
              <article
                v-for="material in materials"
                :key="material.materialId"
                class="rounded-md bg-white p-4"
              >
                <div class="flex items-start gap-3">
                  <div
                    class="flex h-10 w-10 shrink-0 items-center justify-center rounded-md bg-[#eef2ff] text-[#4f46e5]"
                  >
                    <PhFileText :size="20" weight="duotone" />
                  </div>
                  <div class="min-w-0 flex-1">
                    <div class="flex flex-wrap items-center gap-2">
                      <p class="text-sm font-medium text-[#171322]">
                        {{ material.materialTitle }}
                      </p>
                      <span
                        class="rounded-full bg-[#f3ecff] px-2 py-0.5 text-[10px] uppercase tracking-wide text-[#9d5bd2]"
                      >
                        {{ material.materialType }}
                      </span>
                    </div>
                    <p
                      v-if="material.materialDesc"
                      class="mt-2 line-clamp-2 text-sm leading-6 text-[#6b6475]"
                    >
                      {{ material.materialDesc }}
                    </p>
                    <p class="mt-2 text-xs text-[#a09aa8]">
                      {{ material.creatorName || 'Creator tidak tersedia' }} ·
                      {{ material.createdAt }}
                    </p>
                  </div>
                </div>
              </article>
            </div>
          </template>

          <template v-else-if="activeTab === 'assignments'">
            <p class="text-sm font-medium text-[#4f46e5]">Tugas</p>
            <h2 class="mt-3 text-2xl font-medium text-[#171322]">
              Tugas subject akan tampil di sini
            </h2>
            <p class="mt-4 max-w-xl text-sm leading-6 text-[#6b6475]">
              Integrasi daftar tugas subject belum diaktifkan pada tahap ini.
            </p>
          </template>

          <template v-else>
            <p class="text-sm font-medium text-[#4f46e5]">Catatan</p>
            <h2 class="mt-3 text-2xl font-medium text-[#171322]">
              Catatan pribadi akan tersedia di halaman materi
            </h2>
            <p class="mt-4 max-w-xl text-sm leading-6 text-[#6b6475]">
              Notes adalah fitur future per-material. Belum ada autosave atau data catatan yang
              difake di tahap ini.
            </p>
          </template>
        </article>

        <aside class="rounded-md border border-[#ebe7df] bg-white p-5">
          <p class="text-sm font-medium text-[#171322]">Subject context</p>
          <p class="mt-2 text-xs leading-5 text-[#8b8592]">
            {{ teacherName || 'Guru belum tersedia' }}
          </p>
          <p class="mt-5 rounded-md bg-[#fbfaf8] px-4 py-3 text-xs text-[#7a7385]">
            SubjectClass ID: {{ subjectClassId }}
          </p>
        </aside>
      </div>
    </section>
  </main>
</template>
