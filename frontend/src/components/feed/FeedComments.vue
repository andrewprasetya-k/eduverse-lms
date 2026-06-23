<script setup lang="ts">
import { computed, ref } from 'vue'
import {
  PhArrowClockwise,
  PhChatCircleText,
  PhPaperPlaneTilt,
  PhTrash,
} from '@phosphor-icons/vue'
import {
  createFeedComment,
  deleteFeedComment,
  getFeedComments,
} from '../../services/feed'
import type { FeedComment, FeedPost } from '../../types/feed'
import { formatDateTime } from '../../utils/date'

const props = defineProps<{
  post: FeedPost
}>()

const emit = defineEmits<{
  (event: 'comment-count-change', feedId: string, count: number): void
}>()

const isExpanded = ref(false)
const hasLoaded = ref(false)
const isLoading = ref(false)
const isSubmitting = ref(false)
const errorMessage = ref('')
const commentText = ref('')
const comments = ref<FeedComment[]>([])
const deletingCommentIds = ref<Set<string>>(new Set())

const visibleCommentCount = computed(() => props.post.commentCount ?? comments.value.length)

function getCommentErrorMessage(error: unknown, fallback: string) {
  if (typeof error === 'object' && error !== null && 'response' in error) {
    const response = (error as {
      response?: {
        status?: number
        data?: { error?: unknown; message?: unknown }
      }
    }).response

    if (response?.status === 403) {
      return 'Anda tidak memiliki akses untuk komentar ini.'
    }

    if (response?.status === 404) {
      return 'Komentar atau feed tidak ditemukan.'
    }

    if (typeof response?.data?.error === 'string') {
      return response.data.error
    }

    if (typeof response?.data?.message === 'string') {
      return response.data.message
    }
  }

  return fallback
}

async function loadComments() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    comments.value = await getFeedComments(props.post.feedId)
    hasLoaded.value = true
    emit('comment-count-change', props.post.feedId, comments.value.length)
  } catch (error) {
    errorMessage.value = getCommentErrorMessage(error, 'Komentar belum bisa dimuat.')
  } finally {
    isLoading.value = false
  }
}

async function toggleComments() {
  isExpanded.value = !isExpanded.value

  if (isExpanded.value && !hasLoaded.value) {
    await loadComments()
  }
}

async function submitComment() {
  const trimmed = commentText.value.trim()
  if (!trimmed || isSubmitting.value) {
    return
  }

  isSubmitting.value = true
  errorMessage.value = ''

  try {
    await createFeedComment(props.post.feedId, trimmed)
    commentText.value = ''
    await loadComments()
  } catch (error) {
    errorMessage.value = getCommentErrorMessage(error, 'Komentar belum bisa dikirim.')
  } finally {
    isSubmitting.value = false
  }
}

async function removeComment(comment: FeedComment) {
  if (!comment.isMine || deletingCommentIds.value.has(comment.commentId)) {
    return
  }

  deletingCommentIds.value = new Set([...deletingCommentIds.value, comment.commentId])
  errorMessage.value = ''

  try {
    await deleteFeedComment(comment.commentId)
    comments.value = comments.value.filter((item) => item.commentId !== comment.commentId)
    emit('comment-count-change', props.post.feedId, comments.value.length)
  } catch (error) {
    errorMessage.value = getCommentErrorMessage(error, 'Komentar belum bisa dihapus.')
  } finally {
    const nextDeletingIds = new Set(deletingCommentIds.value)
    nextDeletingIds.delete(comment.commentId)
    deletingCommentIds.value = nextDeletingIds
  }
}
</script>

<template>
  <div class="mt-4 border-t border-[#ebe7df] pt-3">
    <button
      class="inline-flex items-center gap-2 rounded-2xl px-3 py-2 text-xs font-medium text-[#4f46e5] transition hover:bg-[#eef2ff] focus:outline-none focus:ring-2 focus:ring-[#4f46e5]/25"
      type="button"
      @click="toggleComments"
    >
      <PhChatCircleText :size="16" weight="duotone" />
      {{ isExpanded ? 'Sembunyikan komentar' : `Lihat komentar${visibleCommentCount ? ` (${visibleCommentCount})` : ''}` }}
    </button>

    <div v-if="isExpanded" class="mt-3 space-y-3 rounded-2xl bg-white/70 p-3">
      <div v-if="isLoading" class="rounded-2xl bg-[#fbfaf8] p-3">
        <p class="text-xs text-[#7a7385]">Memuat komentar...</p>
      </div>

      <div
        v-else-if="errorMessage"
        class="rounded-2xl bg-[#fff7ed] p-3"
      >
        <p class="text-xs leading-5 text-[#9a3412]">{{ errorMessage }}</p>
        <button
          class="mt-3 inline-flex items-center gap-2 rounded-2xl border border-[#fed7aa] px-3 py-2 text-xs font-medium text-[#9a3412] transition hover:bg-[#ffedd5]"
          type="button"
          @click="loadComments"
        >
          <PhArrowClockwise :size="14" />
          Coba lagi
        </button>
      </div>

      <div v-else class="space-y-3">
        <div
          v-if="comments.length === 0"
          class="rounded-2xl bg-[#fbfaf8] p-3"
        >
          <p class="text-xs text-[#7a7385]">Belum ada komentar.</p>
        </div>

        <div
          v-for="comment in comments"
          :key="comment.commentId"
          class="rounded-2xl bg-[#fbfaf8] p-3"
        >
          <div class="flex items-start justify-between gap-3">
            <div class="min-w-0">
              <p class="truncate text-xs font-medium text-[#171322]">
                {{ comment.creatorName || 'Pengirim tidak tersedia' }}
              </p>
              <p class="mt-0.5 text-[11px] text-[#a09aa8]">
                {{ formatDateTime(comment.createdAt) }}
              </p>
            </div>
            <button
              v-if="comment.isMine"
              class="inline-flex shrink-0 items-center gap-1 rounded-xl px-2 py-1 text-[11px] font-medium text-[#b42318] transition hover:bg-[#fff1f0] disabled:cursor-not-allowed disabled:opacity-60"
              type="button"
              :disabled="deletingCommentIds.has(comment.commentId)"
              @click="removeComment(comment)"
            >
              <PhTrash :size="13" />
              Hapus
            </button>
          </div>
          <p class="mt-2 whitespace-pre-line break-words text-xs leading-5 text-[#4a4356]">
            {{ comment.content }}
          </p>
        </div>
      </div>

      <form class="space-y-2" @submit.prevent="submitComment">
        <label class="sr-only" :for="`feed-comment-${post.feedId}`">Tulis komentar</label>
        <textarea
          :id="`feed-comment-${post.feedId}`"
          v-model="commentText"
          class="min-h-20 w-full resize-y rounded-2xl border border-[#ebe7df] bg-white px-3 py-2 text-xs leading-5 text-[#171322] outline-none transition placeholder:text-[#a09aa8] focus:border-[#4f46e5]"
          maxlength="800"
          placeholder="Tulis komentar singkat..."
        />
        <div class="flex items-center justify-between gap-3">
          <p class="text-[11px] text-[#8b8592]">Komentar hanya untuk feed kelas.</p>
          <button
            class="inline-flex items-center gap-2 rounded-2xl bg-[#4f46e5] px-3 py-2 text-xs font-medium text-white transition hover:bg-[#4338ca] disabled:cursor-not-allowed disabled:opacity-60"
            type="submit"
            :disabled="!commentText.trim() || isSubmitting"
          >
            <PhPaperPlaneTilt :size="14" weight="duotone" />
            {{ isSubmitting ? 'Mengirim...' : 'Kirim' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
