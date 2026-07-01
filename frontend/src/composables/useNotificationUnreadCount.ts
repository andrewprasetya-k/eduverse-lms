import { computed, ref } from 'vue'
import { getNotificationUnreadCount } from '../services/notifications'

const unreadCount = ref(0)
const loading = ref(false)
const error = ref('')
let initialized = false
let lifecycleStarted = false
let refreshPromise: Promise<void> | null = null

function normalizeCount(value: number) {
  return Math.max(0, Number.isFinite(value) ? value : 0)
}

async function refreshUnreadCount() {
  if (refreshPromise) return refreshPromise

  loading.value = true
  error.value = ''

  refreshPromise = getNotificationUnreadCount()
    .then((response) => {
      unreadCount.value = normalizeCount(response.unreadCount)
    })
    .catch(() => {
      error.value = 'Jumlah notifikasi belum bisa dimuat.'
    })
    .finally(() => {
      loading.value = false
      refreshPromise = null
    })

  return refreshPromise
}

function startLifecycleRefresh() {
  if (lifecycleStarted || typeof document === 'undefined') return
  lifecycleStarted = true

  document.addEventListener('visibilitychange', () => {
    if (document.visibilityState === 'visible') {
      void refreshUnreadCount()
    }
  })
}

export function useNotificationUnreadCount() {
  startLifecycleRefresh()

  if (!initialized) {
    initialized = true
    void refreshUnreadCount()
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
    refresh: refreshUnreadCount,
    set,
    decrement,
    clear,
  }
}
