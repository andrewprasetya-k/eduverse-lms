import { computed, ref } from 'vue'
import { getNotificationUnreadCount } from '../services/notifications'

const unreadCount = ref(0)
const loading = ref(false)
const error = ref('')

function normalizeCount(value: number) {
  return Math.max(0, Number.isFinite(value) ? value : 0)
}

export function useNotificationUnreadCount() {
  async function refresh() {
    loading.value = true
    error.value = ''

    try {
      const response = await getNotificationUnreadCount()
      unreadCount.value = normalizeCount(response.unreadCount)
    } catch {
      error.value = 'Jumlah notifikasi belum bisa dimuat.'
    } finally {
      loading.value = false
    }
  }

  function set(value: number) {
    unreadCount.value = normalizeCount(value)
  }

  function decrement(step = 1) {
    unreadCount.value = normalizeCount(unreadCount.value - step)
  }

  function clear() {
    unreadCount.value = 0
  }

  const badgeLabel = computed(() => {
    if (unreadCount.value <= 0) return ''
    return unreadCount.value > 99 ? '99+' : String(unreadCount.value)
  })

  return {
    unreadCount,
    loading,
    error,
    badgeLabel,
    refresh,
    set,
    decrement,
    clear,
  }
}
