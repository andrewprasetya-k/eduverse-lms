export interface StudentClassEnrollment {
  enrollmentId: string
  schoolId: string
  schoolUserId: string
  classId: string
  classTitle?: string
  classCode?: string
  role: 'student' | 'teacher' | string
  joinedAt: string
}
