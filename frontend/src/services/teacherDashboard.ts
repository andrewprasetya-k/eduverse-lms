import { api } from './api'
import type { TeacherDashboardSummary } from '../types/teacherDashboard'

export async function getTeacherDashboard(schoolUserId: string) {
  const { data } = await api.get<TeacherDashboardSummary>(`/dashboard/teacher/${schoolUserId}`)
  return data
}
