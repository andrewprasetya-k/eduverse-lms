<script setup lang="ts">
import {
  PhBookOpen,
  PhCalendarBlank,
  PhChartBar,
  PhChatCircle,
  PhHouse,
  PhMegaphone,
  PhNotebook,
} from '@phosphor-icons/vue'
import SlimSidebar from '../components/layout/Sidebar.vue'

const items = [
  { label: 'Dashboard', icon: PhHouse, to: '/student/dashboard' },
  { label: 'Mata Pelajaran', icon: PhBookOpen, to: '/student/subjects' },
  { label: 'Feed', icon: PhMegaphone, to: '/student/feed' },
  { label: 'Tugas', icon: PhCalendarBlank, to: '/student/assignments' },
  { label: 'Nilai', icon: PhChartBar, to: '/student/grades' },
  { label: 'Chat', icon: PhChatCircle, to: '/student/chat' },
  { label: 'Catatan', icon: PhNotebook, to: '/student/notes' },
]
</script>

<template>
  <div class="min-h-screen bg-[#f8f7f4] text-[#2f2b3a]">
    <div class="mx-auto flex min-h-screen max-w-360">
      <SlimSidebar
        class="sticky top-0 h-screen shrink-0"
        label="Navigasi siswa"
        :items="items"
        profile-to="/student/profile"
      />
      <RouterView />
    </div>
  </div>
</template>
