export interface CreateAssignmentPayload {
  schoolId: string
  subjectClassId: string
  categoryId: string
  assignmentTitle: string
  assignmentDescription: string
  deadline?: string
  allowLateSubmission: boolean
  mediaIds: string[]
}

export interface AssignmentCategory {
  asc_id: string
  asc_name: string
}

export interface SchoolCategoriesResponse {
  schoolId: string
  schoolName: string
  data: AssignmentCategory[]
}
