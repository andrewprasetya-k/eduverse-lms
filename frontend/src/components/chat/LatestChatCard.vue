<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { RouterLink } from "vue-router";
import {
  PhArrowRight,
  PhChatCircleText,
  PhWarningCircle,
} from "@phosphor-icons/vue";
import { getChatRooms } from "../../services/chat";
import type { ChatRoom } from "../../types/chat";

const props = withDefaults(
  defineProps<{
    to: string;
    limit?: number;
  }>(),
  {
    limit: 4,
  },
);

const rooms = ref<ChatRoom[]>([]);
const isLoading = ref(false);
const hasError = ref(false);

const visibleRooms = computed(() => rooms.value.slice(0, props.limit));

onMounted(loadLatestChats);

async function loadLatestChats() {
  isLoading.value = true;
  hasError.value = false;
  try {
    rooms.value = await getChatRooms();
  } catch {
    rooms.value = [];
    hasError.value = true;
  } finally {
    isLoading.value = false;
  }
}

function roomDisplayName(room: ChatRoom) {
  if (room.roomRefType === "school") return "Ruang Sekolah";
  if (room.roomType === "dm") {
    return room.dmTargetName || room.dmTargetEmail || "Pesan Langsung";
  }
  return room.roomName || "Ruang Grup";
}

function roomPreview(room: ChatRoom) {
  if (!room.lastMessage?.content) return "Belum ada pesan.";
  const sender = room.lastMessage.senderName || "Pengguna";
  return `${sender}: ${room.lastMessage.content}`;
}

function formatTime(value?: string | null) {
  if (!value) return "";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "";
  return new Intl.DateTimeFormat("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}
</script>

<template>
  <article class="rounded-xl border border-[#ebe7df] bg-white p-4 sm:p-5">
    <div class="mb-4 flex items-center justify-between gap-3">
      <div class="min-w-0">
        <p class="text-sm font-semibold text-[#171322]">Chat terbaru</p>
        <p class="mt-1 text-xs leading-5 text-[#8b8592]">
          Ruang percakapan dengan aktivitas terakhir.
        </p>
      </div>
      <RouterLink
        :to="to"
        class="inline-flex shrink-0 items-center gap-1 text-xs font-semibold text-[#4f46e5] transition hover:text-[#4338ca]"
      >
        Buka chat
        <PhArrowRight :size="14" />
      </RouterLink>
    </div>

    <div v-if="isLoading" class="space-y-2">
      <div
        v-for="item in 3"
        :key="item"
        class="h-14 animate-pulse rounded-lg bg-[#f3f4f6]"
      />
    </div>

    <div
      v-else-if="hasError"
      class="flex gap-3 rounded-lg border border-[#f1d6d3] bg-[#fffafa] p-3"
    >
      <PhWarningCircle :size="18" class="mt-0.5 shrink-0 text-[#dc2626]" />
      <p class="text-xs leading-5 text-[#7a7385]">
        Ringkasan chat belum bisa dimuat. Halaman lain tetap dapat digunakan.
      </p>
    </div>

    <div v-else-if="visibleRooms.length > 0" class="space-y-2">
      <RouterLink
        v-for="room in visibleRooms"
        :key="room.roomId"
        :to="to"
        class="flex min-w-0 items-center gap-3 rounded-lg border border-[#ebe7df] bg-[#fbfaf8] p-3 transition hover:border-[#c7d2fe] hover:bg-white"
      >
        <span
          class="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg bg-[#eef2ff] text-[#4f46e5]"
        >
          <PhChatCircleText :size="18" weight="duotone" />
        </span>
        <span class="min-w-0 flex-1">
          <span
            class="block truncate text-sm text-[#171322]"
            :class="room.unreadCount > 0 ? 'font-bold' : 'font-semibold'"
          >
            {{ roomDisplayName(room) }}
          </span>
          <span
            class="mt-0.5 block truncate text-xs"
            :class="
              room.unreadCount > 0
                ? 'font-semibold text-[#3f3a4a]'
                : 'text-[#7a7385]'
            "
          >
            {{ roomPreview(room) }}
          </span>
        </span>
        <span class="flex shrink-0 flex-col items-end gap-1">
          <span class="text-[11px] text-[#9ca3af]">
            {{ formatTime(room.lastMessageAt) }}
          </span>
          <span
            v-if="room.unreadCount > 0"
            class="rounded-full bg-[#4f46e5] px-2 py-0.5 text-[10px] font-semibold text-white"
          >
            {{ room.unreadCount }}
          </span>
        </span>
      </RouterLink>
    </div>

    <p
      v-else
      class="rounded-lg border border-[#ebe7df] bg-[#fbfaf8] p-4 text-sm leading-6 text-[#7a7385]"
    >
      Belum ada aktivitas chat terbaru.
    </p>
  </article>
</template>
