<script setup lang="ts">
import {
  PhArrowRight,
  PhBookOpen,
  PhCalendarBlank,
  PhChalkboardTeacher,
  PhClipboardText,
  PhHouse,
  PhLinkSimple,
  PhStudent,
  PhUsers,
} from "@phosphor-icons/vue";

const summaryCards = [
  {
    label: "Struktur Akademik",
    value: "Fondasi",
    helper: "Tahun ajaran, semester, mata pelajaran, dan kategori tugas.",
    icon: PhCalendarBlank,
    tone: "bg-[#fff4ee] text-[#ea580c]",
    to: "/admin/academic-years",
  },
  {
    label: "Kelas",
    value: "Ruang belajar",
    helper: "Kelas dibuat sesuai periode akademik yang dipilih.",
    icon: PhBookOpen,
    tone: "bg-[#eef2ff] text-[#4f46e5]",
    to: "/admin/classes",
  },
  {
    label: "Warga Sekolah",
    value: "Akun & peran",
    helper: "Kelola member sekolah dan peran di konteks sekolah.",
    icon: PhUsers,
    tone: "bg-[#ecfdf5] text-[#059669]",
    to: "/admin/users",
  },
  {
    label: "Penempatan Kelas",
    value: "Akses kelas",
    helper: "Tempatkan student atau teacher ke kelas aktif.",
    icon: PhStudent,
    tone: "bg-[#fff7ed] text-[#ea580c]",
    to: "/admin/enrollments",
  },
];

const quickActions = [
  {
    title: "Atur Struktur Akademik",
    description: "Siapkan tahun ajaran, semester, subject, dan kategori tugas.",
    icon: PhCalendarBlank,
    tone: "bg-[#fff4ee] text-[#ea580c]",
    to: "/admin/academic-years",
  },
  {
    title: "Buat Kelas",
    description: "Tambahkan kelas untuk periode akademik yang sudah aktif.",
    icon: PhBookOpen,
    tone: "bg-[#eef2ff] text-[#4f46e5]",
    to: "/admin/classes",
  },
  {
    title: "Kelola Warga Sekolah",
    description: "Tambahkan akun global ke sekolah dan atur perannya.",
    icon: PhUsers,
    tone: "bg-[#ecfdf5] text-[#059669]",
    to: "/admin/users",
  },
  {
    title: "Tempatkan ke Kelas",
    description: "Masukkan student atau teacher ke kelas yang sesuai.",
    icon: PhClipboardText,
    tone: "bg-[#fff7ed] text-[#ea580c]",
    to: "/admin/enrollments",
  },
];

const setupFlow = [
  {
    title: "Struktur Akademik",
    description: "Buat periode akademik dan data subject lebih dulu.",
    to: "/admin/academic-years",
  },
  {
    title: "Kelas",
    description: "Buat kelas di semester yang akan digunakan.",
    to: "/admin/classes",
  },
  {
    title: "Warga Sekolah",
    description: "Pastikan student, teacher, dan admin sudah menjadi member.",
    to: "/admin/users",
  },
  {
    title: "Penempatan Kelas",
    description: "Tempatkan student dan teacher ke kelas aktif.",
    to: "/admin/enrollments",
  },
  {
    title: "Penugasan Mengajar",
    description: "Hubungkan teacher, kelas, dan mata pelajaran.",
    to: "/admin/subject-classes",
  },
];

const managementLinks = [
  {
    label: "Penugasan Mengajar",
    description: "Buat ruang mengajar teacher untuk subject di kelas.",
    icon: PhChalkboardTeacher,
    tone: "bg-[#f3f4f6] text-[#6b7280]",
    to: "/admin/subject-classes",
  },
  {
    label: "Warga Sekolah",
    description: "Periksa membership dan peran sebelum penempatan kelas.",
    icon: PhUsers,
    tone: "bg-[#ecfdf5] text-[#059669]",
    to: "/admin/users",
  },
  {
    label: "Penempatan Kelas",
    description: "Keluarkan atau tempatkan ulang member tanpa menghapus histori.",
    icon: PhStudent,
    tone: "bg-[#fff7ed] text-[#ea580c]",
    to: "/admin/enrollments",
  },
];
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-5 sm:px-6 lg:px-8">
    <section class="grid w-full max-w-none gap-5 xl:grid-cols-[minmax(0,1fr)_320px]">
      <div class="flex min-w-0 flex-col gap-5">
        <header
          class="rounded-[22px] bg-white p-5 shadow-sm ring-1 ring-black/5 md:p-6"
        >
          <div class="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
            <div>
              <div class="flex flex-wrap items-center gap-2">
                <p class="text-sm font-medium text-[#ea580c]">
                  Admin sekolah
                </p>
                <span
                  class="rounded-full bg-[#fff4ee] px-2.5 py-1 text-xs font-medium text-[#ea580c]"
                >
                  Setup MVP
                </span>
              </div>
              <h1 class="mt-3 text-2xl font-medium text-[#171322] md:text-3xl">
                Dashboard admin sekolah
              </h1>
              <p class="mt-3 max-w-3xl text-sm leading-6 text-[#6b6475]">
                Pantau alur setup sekolah dari struktur akademik sampai
                penugasan mengajar. Kartu di halaman ini adalah panduan
                operasional dan tidak menampilkan statistik live.
              </p>
            </div>
            <RouterLink
              to="/admin/academic-years"
              class="inline-flex items-center justify-center gap-2 rounded-2xl bg-[#ea580c] px-4 py-2.5 text-sm font-medium text-white transition hover:bg-[#c2410c]"
            >
              Mulai setup
              <PhArrowRight :size="16" weight="bold" />
            </RouterLink>
          </div>
        </header>

        <section class="grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
          <RouterLink
            v-for="item in summaryCards"
            :key="item.to"
            :to="item.to"
            class="rounded-[18px] border border-[#ebe7df] bg-white p-4 shadow-sm transition hover:border-[#ea580c] hover:shadow-md"
          >
            <div
              class="flex h-10 w-10 items-center justify-center rounded-xl"
              :class="item.tone"
            >
              <component :is="item.icon" :size="20" weight="duotone" />
            </div>
            <p class="mt-4 text-xs text-[#6b6475]">{{ item.label }}</p>
            <p class="mt-1 text-lg font-medium text-[#171322]">
              {{ item.value }}
            </p>
            <p class="mt-2 text-xs leading-5 text-[#7a7385]">
              {{ item.helper }}
            </p>
          </RouterLink>
        </section>

        <section>
          <p class="mb-3 text-xs font-medium uppercase tracking-wide text-[#9ca3af]">
            Aksi cepat
          </p>
          <div class="grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
            <RouterLink
              v-for="item in quickActions"
              :key="item.to"
              :to="item.to"
              class="flex min-h-36 flex-col rounded-[18px] border border-[#ebe7df] bg-white p-4 shadow-sm transition hover:-translate-y-0.5 hover:border-[#ea580c] hover:shadow-md"
            >
              <div
                class="flex h-10 w-10 items-center justify-center rounded-xl"
                :class="item.tone"
              >
                <component :is="item.icon" :size="20" weight="duotone" />
              </div>
              <p class="mt-4 text-sm font-medium text-[#171322]">
                {{ item.title }}
              </p>
              <p class="mt-2 text-xs leading-5 text-[#7a7385]">
                {{ item.description }}
              </p>
              <p class="mt-auto pt-4 text-xs font-medium text-[#ea580c]">
                Buka halaman
              </p>
            </RouterLink>
          </div>
        </section>

        <section class="grid gap-5 lg:grid-cols-2">
          <article class="rounded-[20px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <div class="flex items-center justify-between gap-4">
              <div>
                <p class="text-sm font-medium text-[#171322]">Alur setup</p>
                <p class="mt-1 text-xs text-[#7a7385]">
                  Urutan kerja yang disarankan untuk admin sekolah.
                </p>
              </div>
              <PhLinkSimple :size="20" class="text-[#ea580c]" weight="duotone" />
            </div>

            <div class="mt-4 divide-y divide-[#f3f4f6]">
              <RouterLink
                v-for="(item, index) in setupFlow"
                :key="item.to"
                :to="item.to"
                class="flex items-start gap-3 py-3"
              >
                <span
                  class="mt-0.5 flex h-7 w-7 shrink-0 items-center justify-center rounded-lg bg-[#fff4ee] text-xs font-medium text-[#ea580c]"
                >
                  {{ index + 1 }}
                </span>
                <span class="min-w-0">
                  <span class="block text-sm font-medium text-[#171322]">
                    {{ item.title }}
                  </span>
                  <span class="mt-1 block text-xs leading-5 text-[#7a7385]">
                    {{ item.description }}
                  </span>
                </span>
              </RouterLink>
            </div>
          </article>

          <article class="rounded-[20px] bg-white p-5 shadow-sm ring-1 ring-black/5">
            <div class="flex items-center justify-between gap-4">
              <div>
                <p class="text-sm font-medium text-[#171322]">
                  Manajemen aktif
                </p>
                <p class="mt-1 text-xs text-[#7a7385]">
                  Area yang biasanya perlu dicek setelah setup awal.
                </p>
              </div>
              <PhHouse :size="20" class="text-[#4f46e5]" weight="duotone" />
            </div>

            <div class="mt-4 space-y-2">
              <RouterLink
                v-for="item in managementLinks"
                :key="item.to"
                :to="item.to"
                class="flex items-center gap-3 rounded-2xl p-3 transition hover:bg-[#f9fafb]"
              >
                <span
                  class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl"
                  :class="item.tone"
                >
                  <component :is="item.icon" :size="20" weight="duotone" />
                </span>
                <span class="min-w-0 flex-1">
                  <span class="block text-sm font-medium text-[#171322]">
                    {{ item.label }}
                  </span>
                  <span class="mt-1 block text-xs leading-5 text-[#7a7385]">
                    {{ item.description }}
                  </span>
                </span>
                <PhArrowRight :size="16" class="shrink-0 text-[#9ca3af]" />
              </RouterLink>
            </div>
          </article>
        </section>
      </div>

      <aside class="flex min-w-0 flex-col gap-5">
        <section class="rounded-[20px] bg-white p-5 shadow-sm ring-1 ring-black/5">
          <p class="text-sm font-medium text-[#171322]">Manajemen sekolah</p>
          <p class="mt-1 text-xs leading-5 text-[#7a7385]">
            Modul utama mengikuti alur setup MVP dan mengarah ke halaman admin
            yang sudah tersedia.
          </p>

          <div class="mt-4 space-y-2">
            <RouterLink
              v-for="item in setupFlow"
              :key="`module-${item.to}`"
              :to="item.to"
              class="block rounded-2xl px-3 py-2.5 transition hover:bg-[#f9fafb]"
            >
              <p class="text-sm font-medium text-[#171322]">{{ item.title }}</p>
              <p class="mt-1 text-xs leading-5 text-[#9ca3af]">
                {{ item.description }}
              </p>
            </RouterLink>
          </div>
        </section>

        <section class="rounded-[20px] bg-white p-5 shadow-sm ring-1 ring-black/5">
          <p class="text-sm font-medium text-[#171322]">Catatan MVP</p>
          <div class="mt-4 space-y-3">
            <div class="flex gap-3">
              <span class="mt-1.5 h-2 w-2 shrink-0 rounded-full bg-[#059669]" />
              <p class="text-xs leading-5 text-[#6b6475]">
                Data live dan statistik operasional belum ditarik di dashboard
                ini agar tidak menampilkan angka yang menyesatkan.
              </p>
            </div>
            <div class="flex gap-3">
              <span class="mt-1.5 h-2 w-2 shrink-0 rounded-full bg-[#4f46e5]" />
              <p class="text-xs leading-5 text-[#6b6475]">
                Gunakan halaman Penempatan Kelas dan Penugasan Mengajar untuk
                membuka akses teacher dan student ke ruang belajar.
              </p>
            </div>
            <div class="flex gap-3">
              <span class="mt-1.5 h-2 w-2 shrink-0 rounded-full bg-[#ea580c]" />
              <p class="text-xs leading-5 text-[#6b6475]">
                Histori akademik tidak dikelola dari dashboard; tindakan utama
                tetap berada di halaman setup masing-masing.
              </p>
            </div>
          </div>
        </section>
      </aside>
    </section>
  </main>
</template>
