export interface StudentDeadline {
  assignmentId: string
  assignmentTitle: string
  subjectName: string
  deadline: string
  isSubmitted: boolean
}

export interface StudentDashboardSummary {
  pendingAssignments: number
  upcomingDeadlines: StudentDeadline[]
  averageScore: number
  completedMaterials: number
  totalMaterials: number
}

export interface EnrollmentClass {
  classId?: string
  classTitle?: string
  classCode?: string
  subjectName?: string
}

export interface NotificationItem {
  notificationId: string
  type: string
  title: string
  message: string
  link?: string
  isRead: boolean
  createdAt: string
}

export interface NotificationListResponse {
  data: NotificationItem[]
  unreadCount: number
  totalItems: number
  page: number
  limit: number
  totalPages: number
}

export interface UnreadCountResponse {
  unreadCount: number
}
