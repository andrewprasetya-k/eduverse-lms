import { api } from './api'
import type {
  EnrollmentClass,
  NotificationListResponse,
  StudentDashboardSummary,
} from '../types/dashboard'

export async function getStudentDashboard(userId: string) {
  const { data } = await api.get<StudentDashboardSummary>(`/dashboard/student/${userId}`)
  return data
}

export async function getMemberClasses(schoolUserId: string) {
  const { data } = await api.get<EnrollmentClass[]>(`/enrollments/member/${schoolUserId}`)
  return data
}

export async function getRecentNotifications() {
  const { data } = await api.get<NotificationListResponse>('/notifications', {
    params: { page: 1, limit: 5 },
  })
  return data
}

export async function markNotificationAsRead(notificationId: string) {
  const { data } = await api.patch<{ message: string }>(`/notifications/read/${notificationId}`)
  return data
}

export async function markAllNotificationsAsRead() {
  const { data } = await api.patch<{ message: string }>('/notifications/read-all')
  return data
}
