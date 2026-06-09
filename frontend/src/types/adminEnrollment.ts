export type ClassEnrollmentRole = "student" | "teacher";

export interface EnrollmentMemberItem {
  enrollmentId: string;
  schoolUserId: string;
  userFullName?: string;
  userEmail?: string;
  role: ClassEnrollmentRole;
  joinedAt: string;
}

export interface EnrollmentClassHeader {
  classId: string;
  classTitle: string;
  classCode: string;
}

export interface EnrollmentMembersPaginatedData {
  data: EnrollmentMemberItem[];
  totalItems: number;
  page: number;
  limit: number;
  totalPages: number;
}

export interface ClassEnrollmentsResponse {
  class: EnrollmentClassHeader;
  members: EnrollmentMembersPaginatedData;
}

export interface CreateEnrollmentPayload {
  schoolId: string;
  schoolUserIds: string[];
  classId: string;
  role: ClassEnrollmentRole;
}
