import { api } from "./api";
import type {
  CreateSuperAdminSchoolPayload,
  SuperAdminSchoolSummary,
  SuperAdminSchoolsResponse,
} from "../types/superAdminSchool";

export async function getSuperAdminSchools(params: {
  page?: number;
  limit?: number;
  search?: string;
  status?: "active" | "deleted" | "all";
}) {
  const { data } = await api.get<SuperAdminSchoolsResponse>("/schools", {
    params: {
      page: params.page ?? 1,
      limit: params.limit ?? 50,
      search: params.search || undefined,
      status: params.status ?? "all",
      sortBy: "createdAt",
      order: "desc",
    },
  });
  return data;
}

export async function getSuperAdminSchoolSummary() {
  const { data } = await api.get<SuperAdminSchoolSummary>("/schools/summary");
  return data;
}

export async function createSuperAdminSchool(payload: CreateSuperAdminSchoolPayload) {
  const { data } = await api.post("/schools", payload);
  return data;
}
