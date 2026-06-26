import { api } from "./api";
import type {
  AdminSchoolMemberImportCommitPayload,
  AdminSchoolMemberImportCommitResponse,
  AdminSchoolMemberImportPreviewResponse,
} from "../types/adminSchoolMemberImport";

export async function previewSchoolMemberImport(file: File) {
  const formData = new FormData();
  formData.append("file", file);

  const { data } = await api.post<AdminSchoolMemberImportPreviewResponse>(
    "/admin/school-members/import/preview",
    formData,
    {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    },
  );
  return data;
}

export async function commitSchoolMemberImport(
  payload: AdminSchoolMemberImportCommitPayload,
) {
  const { data } = await api.post<AdminSchoolMemberImportCommitResponse>(
    "/admin/school-members/import/commit",
    payload,
  );
  return data;
}
