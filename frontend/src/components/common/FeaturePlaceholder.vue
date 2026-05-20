<script setup lang="ts">
import { PhSparkle } from '@phosphor-icons/vue'

defineProps<{
  eyebrow: string
  title: string
  description: string
}>()
</script>

<template>
  <section class="flex min-h-screen flex-1 items-center justify-center px-5 py-8">
    <article class="soft-card w-full max-w-2xl rounded-[28px] p-8">
      <div
        class="mb-5 flex h-12 w-12 items-center justify-center rounded-2xl bg-[#eef2ff] text-[#4f46e5]"
      >
        <PhSparkle :size="24" weight="duotone" />
      </div>
      <p class="text-sm font-medium text-[#4f46e5]">{{ eyebrow }}</p>
      <h1 class="mt-3 text-3xl font-medium text-[#171322]">Fitur sedang dikembangkan</h1>
      <h2 class="mt-2 text-xl font-medium text-[#3f3a4a]">{{ title }}</h2>
      <p class="mt-4 max-w-xl text-sm leading-6 text-[#6b6475]">
        {{ description }}
      </p>
    </article>
  </section>
</template>
