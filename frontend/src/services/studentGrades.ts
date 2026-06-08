import { api } from './api'
import type { MyGradebookResponse } from '../types/studentGrades'

export async function getMyGradebookByClass(classId: string) {
  const { data } = await api.get<MyGradebookResponse>(`/grades/my-grades/${classId}`)
  return data
}
