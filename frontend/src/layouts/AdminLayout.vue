<script setup lang="ts">
import {
  PhBookOpen,
  PhCalendarBlank,
  PhHouse,
  PhStudent,
  PhUsers,
  PhUserCircle,
} from '@phosphor-icons/vue'
import SlimSidebar from '../components/layout/Sidebar.vue'

const items = [
  { label: 'Dashboard', icon: PhHouse, to: '/admin/dashboard' },
  { label: 'Classes', icon: PhBookOpen, to: '/admin/classes' },
  { label: 'Users', icon: PhUsers, to: '/admin/users' },
  { label: 'Enrollments', icon: PhStudent, to: '/admin/enrollments' },
  { label: 'Academic Years', icon: PhCalendarBlank, to: '/admin/academic-years' },
  { label: 'Profile', icon: PhUserCircle, to: '/admin/profile' },
]
</script>

<template>
  <div class="min-h-screen bg-[#f8f7f4] text-[#2f2b3a]">
    <div class="mx-auto flex min-h-screen max-w-360">
      <SlimSidebar class="sticky top-0 h-screen shrink-0" label="Admin navigation" :items="items" />
      <RouterView />
    </div>
  </div>
</template>
