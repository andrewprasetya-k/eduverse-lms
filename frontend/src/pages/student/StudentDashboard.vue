<script setup lang="ts">
import { PhBell, PhCaretLeft, PhCaretRight, PhChatCircleText } from '@phosphor-icons/vue'
import { useAuthStore } from '../../stores/auth'

const auth = useAuthStore()

const classes = [
  {
    name: 'Bahasa Indonesia',
    color: '#4f8ef7',
    deadline: 'Tenggat 15 Sept',
    detail: '23:59 - Teks Eksplanasi',
    progress: 50,
  },
  {
    name: 'Matematika',
    color: '#f2756a',
    deadline: 'Tenggat hari ini',
    detail: '22:00 - Tugas 1',
    progress: 50,
  },
  {
    name: 'PPKN',
    color: '#c673d8',
    deadline: 'Tenggat besok',
    detail: '08:00 - Laporan',
    progress: 70,
  },
  {
    name: 'Fisika',
    color: '#f0a05a',
    deadline: 'Tidak ada tugas aktif',
    detail: '',
    progress: 30,
  },
]

const chats = [
  {
    name: 'Bahasa Indonesia',
    avatar: 'BI',
    color: '#4f8ef7',
    time: '10 mnt',
    preview: 'Aisya: nanti apakah boleh minta tolong bawa laptop?',
    unread: 2,
  },
  {
    name: 'Matematika',
    avatar: 'MTK',
    color: '#f2756a',
    time: '1 jam',
    preview: 'El: btw ngingetin besok ada tugas 3 sudah ada tuh',
    unread: 5,
  },
  {
    name: 'PPKN',
    avatar: 'PKN',
    color: '#c673d8',
    time: 'Kemarin',
    preview: 'Pak Hendra: Laporan dikumpul paling lambat besok ya',
    unread: 0,
  },
  {
    name: 'Riana & Siti',
    avatar: 'RS',
    color: '#4f46e5',
    time: 'Kemarin',
    preview: 'Siti: oke ketemu di kantin jam 12 ya guys',
    unread: 0,
  },
]

const calendarDays = [
  '31',
  '1',
  '2',
  '3',
  '4',
  '5',
  '6',
  '7',
  '8',
  '9',
  '10',
  '11',
  '12',
  '13',
  '14',
  '15',
  '16',
  '17',
  '18',
  '19',
  '20',
]
</script>

<template>
  <main class="grid min-h-screen flex-1 grid-cols-1 overflow-hidden lg:grid-cols-[1fr_320px]">
    <section class="flex flex-col gap-6 px-5 py-6 sm:px-8 lg:px-10">
      <header class="flex flex-col gap-2">
        <p class="text-sm text-[#7a7385]">
          {{ auth.defaultContext ? auth.memberships[0]?.school.name : 'Eduverse Academy' }}
        </p>
        <h1 class="text-2xl font-medium tracking-normal text-[#171322]">
          Selamat pagi, {{ auth.user?.fullName?.split(' ')[0] || 'Siswa' }}
        </h1>
        <p class="text-sm text-[#7a7385]">Rabu, 20 Mei 2026 - Semester Genap - Minggu ke-14</p>
      </header>

      <section class="grid gap-3 sm:grid-cols-2 xl:grid-cols-3">
        <article
          v-for="item in classes"
          :key="item.name"
          class="overflow-hidden rounded-[18px] border border-[#ebe7df] bg-white transition hover:-translate-y-0.5 hover:shadow-[0_18px_40px_rgba(66,55,40,0.08)]"
        >
          <div
            class="flex h-24 flex-col justify-end px-4 pb-4 text-white"
            :style="{ backgroundColor: item.color }"
          >
            <h2 class="text-base font-medium">{{ item.name }}</h2>
          </div>
          <div class="space-y-3 px-4 py-4">
            <p class="min-h-9 text-xs leading-5 text-[#7a7385]">
              <strong class="font-medium text-[#3f3a4a]">{{ item.deadline }}</strong>
              <span v-if="item.detail"> - {{ item.detail }}</span>
            </p>
            <div class="flex items-center gap-2">
              <div class="h-1 flex-1 overflow-hidden rounded-full bg-[#f0ede8]">
                <div
                  class="h-full rounded-full"
                  :style="{ width: `${item.progress}%`, backgroundColor: item.color }"
                />
              </div>
              <span class="w-8 text-right text-[11px] text-[#9a95a3]">{{ item.progress }}%</span>
            </div>
          </div>
        </article>
      </section>

      <section class="grid gap-4 xl:grid-cols-[1.1fr_0.9fr]">
        <article class="soft-card rounded-[22px] p-5">
          <div class="mb-4 flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-[#171322]">Aktivitas minggu ini</p>
              <p class="text-xs text-[#8b8592]">Ringkasan awal untuk workspace siswa</p>
            </div>
            <PhBell :size="20" class="text-[#4f46e5]" />
          </div>
          <div class="grid gap-3 sm:grid-cols-3">
            <div class="rounded-2xl bg-[#eef2ff] p-4">
              <p class="text-2xl font-medium text-[#4f46e5]">4</p>
              <p class="mt-1 text-xs text-[#6b6475]">kelas aktif</p>
            </div>
            <div class="rounded-2xl bg-[#fff1f0] p-4">
              <p class="text-2xl font-medium text-[#f2756a]">2</p>
              <p class="mt-1 text-xs text-[#6b6475]">tugas dekat deadline</p>
            </div>
            <div class="rounded-2xl bg-[#f3ecff] p-4">
              <p class="text-2xl font-medium text-[#9d5bd2]">76%</p>
              <p class="mt-1 text-xs text-[#6b6475]">progress materi</p>
            </div>
          </div>
        </article>

        <article class="soft-card rounded-[22px] p-5">
          <p class="text-sm font-medium text-[#171322]">Hari ini</p>
          <div class="mt-4 space-y-3 text-sm text-[#6b6475]">
            <div class="rounded-2xl bg-[#fbfaf8] p-3">Matematika - Tugas 1 perlu dikumpulkan</div>
            <div class="rounded-2xl bg-[#fbfaf8] p-3">Bahasa Indonesia - Materi baru tersedia</div>
          </div>
        </article>
      </section>
    </section>

    <aside class="border-l border-[#ebe7df] bg-white/95">
      <div class="flex border-b border-[#ebe7df] px-5">
        <button
          class="flex items-center gap-2 border-b-2 border-[#4f46e5] px-1 py-4 text-sm font-medium text-[#4f46e5]"
          type="button"
        >
          <PhChatCircleText :size="18" />
          Chat
        </button>
        <button class="px-5 py-4 text-sm text-[#a09aa8]" type="button">Notifikasi</button>
      </div>

      <div class="space-y-1 p-4">
        <article
          v-for="chat in chats"
          :key="chat.name"
          class="flex gap-3 rounded-2xl p-3 transition hover:bg-[#f8f7f4]"
          :class="chat.unread ? 'bg-[#f5f7ff]' : ''"
        >
          <div
            class="flex h-9 w-9 shrink-0 items-center justify-center rounded-full text-[11px] font-medium text-white"
            :style="{ backgroundColor: chat.color }"
          >
            {{ chat.avatar }}
          </div>
          <div class="min-w-0 flex-1">
            <div class="flex items-baseline justify-between gap-2">
              <p class="truncate text-sm font-medium text-[#171322]">{{ chat.name }}</p>
              <span class="shrink-0 text-[10px] text-[#a09aa8]">{{ chat.time }}</span>
            </div>
            <p class="truncate text-xs text-[#7a7385]">{{ chat.preview }}</p>
            <span
              v-if="chat.unread"
              class="mt-1 inline-flex rounded-full bg-[#4f46e5] px-2 py-0.5 text-[10px] font-medium text-white"
            >
              {{ chat.unread }}
            </span>
          </div>
        </article>
      </div>

      <section class="border-t border-[#ebe7df] p-5">
        <div class="mb-4 flex items-center justify-between">
          <p class="text-sm font-medium text-[#171322]">Mei 2026</p>
          <div class="flex gap-1">
            <button class="rounded-lg border border-[#ebe7df] p-1.5 text-[#7a7385]" type="button">
              <PhCaretLeft :size="14" />
            </button>
            <button class="rounded-lg border border-[#ebe7df] p-1.5 text-[#7a7385]" type="button">
              <PhCaretRight :size="14" />
            </button>
          </div>
        </div>

        <div class="grid grid-cols-7 gap-1 text-center">
          <span
            v-for="day in ['Min', 'Sen', 'Sel', 'Rab', 'Kam', 'Jum', 'Sab']"
            :key="day"
            class="py-1 text-[10px] text-[#a09aa8]"
          >
            {{ day }}
          </span>
          <span
            v-for="day in calendarDays"
            :key="day"
            class="relative rounded-lg py-1.5 text-xs text-[#4a4356]"
            :class="day === '20' ? 'bg-[#4f46e5] font-medium text-white' : ''"
          >
            {{ day }}
            <span
              v-if="['21', '15'].includes(day)"
              class="absolute bottom-0.5 left-1/2 h-1 w-1 -translate-x-1/2 rounded-full bg-[#f0a05a]"
            />
          </span>
        </div>
      </section>
    </aside>
  </main>
</template>
