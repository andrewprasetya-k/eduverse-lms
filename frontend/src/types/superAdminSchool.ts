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

export interface SuperAdminSchoolBootstrapPayload {
  school: CreateSuperAdminSchoolPayload;
  adminUser:
    | {
        mode: "new";
        fullName: string;
        email: string;
        password: string;
      }
    | {
        mode: "existing";
        userId: string;
      };
}

export interface SuperAdminSchoolBootstrapResponse {
  school: {
    schoolId: string;
    schoolName: string;
    schoolCode: string;
  };
  adminUser: {
    userId: string;
    fullName: string;
    email: string;
    isActive: boolean;
  };
  schoolUserId: string;
  assignedRoles: string[];
}
