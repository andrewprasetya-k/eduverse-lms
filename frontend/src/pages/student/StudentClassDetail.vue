<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { PhBookOpen } from '@phosphor-icons/vue'

const route = useRoute()
const classId = computed(() => String(route.params.classId ?? ''))
</script>

<template>
  <main class="min-h-screen flex-1 px-5 py-6 sm:px-8 lg:px-10">
    <section class="soft-card max-w-3xl rounded-[28px] p-8">
      <div class="mb-5 flex h-12 w-12 items-center justify-center rounded-2xl bg-[#eef2ff] text-[#4f46e5]">
        <PhBookOpen :size="24" weight="duotone" />
      </div>
      <p class="text-sm font-medium text-[#4f46e5]">Class detail</p>
      <h1 class="mt-3 text-3xl font-medium text-[#171322]">Detail kelas sedang dikembangkan</h1>
      <p class="mt-4 text-sm leading-6 text-[#6b6475]">
        Halaman kelas ini nantinya akan berisi materi, tugas, feed, nilai, chat, dan notes sesuai
        konteks kelas. Untuk tahap ini, halaman hanya menyiapkan route detail kelas.
      </p>
      <p class="mt-5 rounded-2xl bg-[#fbfaf8] px-4 py-3 text-xs text-[#7a7385]">
        Class ID: {{ classId }}
      </p>
    </section>
  </main>
</template>
