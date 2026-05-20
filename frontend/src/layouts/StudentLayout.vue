<script setup lang="ts">
import SlimSidebar from '../components/layout/Sidebar.vue'
</script>

<template>
  <div class="min-h-screen bg-[#f8f7f4] text-[#2f2b3a]">
    <div class="mx-auto flex min-h-screen max-w-360">
      <SlimSidebar class="sticky top-0 h-screen shrink-0" />
      <RouterView />
    </div>
  </div>
</template>
