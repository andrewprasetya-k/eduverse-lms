import { api } from './api'
import type { TeacherSubjectClassesResponse } from '../types/teacherSubjects'

export async function getMyTeachingSubjectClasses() {
  const { data } = await api.get<TeacherSubjectClassesResponse>('/subject-classes/my-teaching')
  return data.data ?? []
}

export async function getMyTeachingSubjectClassById(subjectClassId: string) {
  const subjects = await getMyTeachingSubjectClasses()
  return subjects.find((subject) => subject.subjectClassId === subjectClassId) ?? null
}
