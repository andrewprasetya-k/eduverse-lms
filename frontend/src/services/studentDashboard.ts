import { api } from './api'
import type {
  EnrollmentClass,
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
