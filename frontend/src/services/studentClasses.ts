import { api } from './api'
import type { StudentClassEnrollment } from '../types/studentClasses'

export async function getStudentClasses(schoolUserId: string) {
  const { data } = await api.get<StudentClassEnrollment[]>(`/enrollments/member/${schoolUserId}`)
  return data
}
