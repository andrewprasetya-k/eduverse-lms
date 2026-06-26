import { api } from "./api";
import type {
  AdminSchoolMemberCreatePayload,
  AdminSchoolMemberItem,
  AdminSchoolMemberListResponse,
} from "../types/adminSchoolMember";

export async function getAdminSchoolMembers(params: {
  page?: number;
  limit?: number;
  search?: string;
  role?: string;
  includeDeleted?: boolean;
}) {
  const { data } = await api.get<AdminSchoolMemberListResponse>(
    "/admin/school-members",
    {
      params: {
        page: params.page ?? 1,
        limit: params.limit ?? 50,
        search: params.search || undefined,
        role: params.role || undefined,
        includeDeleted: params.includeDeleted ? "true" : undefined,
      },
    },
  );
  return data;
}

export async function createAdminSchoolMember(
  payload: AdminSchoolMemberCreatePayload,
) {
  const { data } = await api.post<AdminSchoolMemberItem>(
    "/admin/school-members",
    payload,
  );
  return data;
}

export async function removeAdminSchoolMember(schoolUserId: string) {
  const { data } = await api.delete(`/admin/school-members/${schoolUserId}`);
  return data;
}
