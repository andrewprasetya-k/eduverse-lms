<script setup lang="ts">
import {
  PhBookOpen,
  PhCalendarBlank,
  PhChartBar,
  PhChatCircle,
  PhHouse,
  PhNotebook,
  PhUserCircle,
} from '@phosphor-icons/vue'
import SlimSidebar from '../components/layout/Sidebar.vue'

const items = [
  { label: 'Dashboard', icon: PhHouse, to: '/student/dashboard' },
  { label: 'Classes', icon: PhBookOpen, to: '/student/classes' },
  { label: 'Assignments', icon: PhCalendarBlank, to: '/student/assignments' },
  { label: 'Chat', icon: PhChatCircle, to: '/student/chat', hasDot: true },
  { label: 'Notes', icon: PhNotebook, to: '/student/notes' },
  { label: 'Grades', icon: PhChartBar, to: '/student/grades' },
  { label: 'Profile', icon: PhUserCircle, to: '/student/profile' },
]
</script>

<template>
  <div class="min-h-screen bg-[#f8f7f4] text-[#2f2b3a]">
    <div class="mx-auto flex min-h-screen max-w-360">
      <SlimSidebar class="sticky top-0 h-screen shrink-0" label="Student navigation" :items="items" />
      <RouterView />
    </div>
  </div>
</template>
