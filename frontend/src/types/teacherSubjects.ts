export interface TeacherSubjectClass {
  subjectClassId: string
  classId: string
  className: string
  classCode?: string
  subjectId: string
  subjectName: string
  subjectCode?: string
  studentCount: number
  materialCount: number
  assignmentCount: number
  pendingSubmissions: number
}

export interface TeacherSubjectClassesResponse {
  data: TeacherSubjectClass[]
}
