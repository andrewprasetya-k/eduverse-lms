export interface SuperAdminSchoolItem {
  schoolId: string;
  schoolName: string;
  schoolCode: string;
  schoolLogo?: string;
  schoolAddress: string;
  schoolEmail: string;
  schoolPhone: string;
  schoolWebsite?: string;
  isDeleted: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface SuperAdminSchoolsResponse {
  data: SuperAdminSchoolItem[];
  totalItems: number;
  page: number;
  limit: number;
  totalPages: number;
}

export interface SuperAdminSchoolSummary {
  totalActive: number;
  totalDeleted: number;
  totalSchools: number;
}

export interface CreateSuperAdminSchoolPayload {
  schoolName: string;
  schoolCode?: string;
  schoolAddress: string;
  schoolEmail: string;
  schoolPhone: string;
  schoolWebsite?: string;
}
