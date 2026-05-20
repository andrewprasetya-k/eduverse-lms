import { api } from './api'
import type {
  EnrollmentClass,
  NotificationListResponse,
  StudentDashboardSummary,
  UnreadCountResponse,
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
  const { data } = await api.get<NotificationListResponse>('/notifications/', {
    params: { page: 1, limit: 5 },
  })
  return data
}

export async function getUnreadNotificationCount() {
  const { data } = await api.get<UnreadCountResponse>('/notifications/unread-count')
  return data
}
