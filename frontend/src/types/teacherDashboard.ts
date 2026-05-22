export interface TeacherClassPerformance {
  classId: string
  className: string
  subjectName: string
  averageScore: number
  submissionRate: number
  totalStudents: number
}

export interface TeacherDashboardSummary {
  pendingReviews: number
  totalStudents: number
  submissionRate: number
  classPerformance: TeacherClassPerformance[]
}
