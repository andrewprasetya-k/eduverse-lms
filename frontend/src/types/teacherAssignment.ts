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

export interface TeacherSubmissionAttachment {
  mediaId: string
  mediaName: string
  fileUrl?: string
  mimeType?: string
  fileSize?: number
}

export interface TeacherSubmissionAssessment {
  score: number
  feedback: string
  assessor?: string
  assessorName?: string
  assessedAt?: string
}

export interface TeacherSubmission {
  submissionId: string
  studentName: string
  submittedAt: string
  isLate: boolean
  attachments?: TeacherSubmissionAttachment[]
  assessment?: TeacherSubmissionAssessment
}

export interface AssignmentWithSubmissionsResponse {
  assignment: {
    assignmentId: string
    assignmentTitle: string
    subjectName?: string
    categoryName?: string
    deadline?: string
  }
  submissions: TeacherSubmission[]
}

export interface TeacherAssignmentHeader {
  assignmentId: string
  assignmentTitle: string
  subjectName?: string
  categoryName?: string
  deadline?: string
}

export interface TeacherSubjectClassHeader {
  subjectClassId: string
  subjectCode: string
  subjectName?: string
  teacherId: string
  teacherName?: string
}

export interface TeacherSubmissionGroup {
  assignment: TeacherAssignmentHeader
  submissionCount: number
  gradedCount: number
  pendingCount: number
  submissions: TeacherSubmission[]
}

export interface TeacherSubmissionSummary {
  assignmentCount: number
  submissionCount: number
  gradedCount: number
  pendingCount: number
  lateCount: number
}

export interface TeacherSubjectClassSubmissionsResponse {
  subjectClass: TeacherSubjectClassHeader
  assignments: TeacherSubmissionGroup[]
  summary: TeacherSubmissionSummary
}
