import { api } from './api'
import type { AssignmentItem, AssignmentListResponse } from '../types/assignment'

export async function getSubjectAssignments(subjectClassId: string, page = 1, limit = 20) {
  const { data } = await api.get<AssignmentListResponse>(
    `/assignments/subject-class/${subjectClassId}`,
    {
      params: { page, limit },
    }
  )
  return data
}

export async function getSubjectAssignmentDetail(subjectClassId: string, assignmentId: string) {
  const response = await getSubjectAssignments(subjectClassId, 1, 100)
  const assignment =
    (response.data.data as AssignmentItem[]).find((item) => item.assignmentId === assignmentId) ??
    null

  return {
    subjectClass: response.subjectClass,
    assignment,
  }
}
