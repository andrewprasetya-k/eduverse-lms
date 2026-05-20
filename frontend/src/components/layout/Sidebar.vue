<script setup lang="ts">
import {
  PhBookOpen,
  PhCalendarBlank,
  PhChartBar,
  PhChatCircle,
  PhHouse,
  PhNotebook,
  PhSignOut,
} from "@phosphor-icons/vue";
import { useAuthStore } from "../../stores/auth";

const auth = useAuthStore();

const items = [
  { label: "Dashboard", icon: PhHouse, active: true },
  { label: "Classes", icon: PhBookOpen },
  { label: "Assignments", icon: PhCalendarBlank },
  { label: "Messages", icon: PhChatCircle, hasDot: true },
  { label: "Notes", icon: PhNotebook },
  { label: "Grades", icon: PhChartBar },
];
</script>

<template>
  <aside
    class="flex h-full w-16 flex-col items-center border-r border-[#ebe7df] bg-white/95 px-0 py-4"
  >
    <div
      class="mb-3 flex h-9 w-9 items-center justify-center rounded-xl bg-[#4f46e5] text-[13px] font-medium text-white"
    >
      Ev
    </div>

    <nav
      class="flex flex-1 flex-col items-center gap-1"
      aria-label="Student navigation"
    >
      <button
        v-for="item in items"
        :key="item.label"
        :title="item.label"
        class="relative flex h-10 w-10 items-center justify-center rounded-xl text-[#a3a1aa] transition hover:bg-[#f3f1ec] hover:text-[#3f3a4a]"
        :class="item.active ? 'bg-[#eef2ff] text-[#4f46e5]' : ''"
        type="button"
      >
        <component :is="item.icon" :size="20" weight="regular" />
        <span
          v-if="item.hasDot"
          class="absolute right-2 top-2 h-1.5 w-1.5 rounded-full border border-white bg-[#4f46e5]"
        />
      </button>
    </nav>

    <button
      title="Logout"
      class="mb-3 flex h-10 w-10 items-center justify-center rounded-xl text-[#a3a1aa] transition hover:bg-[#f3f1ec] hover:text-[#3f3a4a]"
      type="button"
      @click="auth.logout()"
    >
      <PhSignOut :size="19" />
    </button>

    <div
      class="flex h-8 w-8 items-center justify-center rounded-full bg-[#4f46e5] text-[11px] font-medium text-white"
    >
      {{ auth.user?.fullName?.slice(0, 2).toUpperCase() || "EV" }}
    </div>
  </aside>
</template>
