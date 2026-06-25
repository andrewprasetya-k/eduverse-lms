import { api } from './api'
import type {
  StudentGlobalMaterialNotesResponse,
  SaveStudentMaterialNotePayload,
  StudentMaterialNoteResponse,
  StudentSubjectMaterialNotesResponse,
} from '../types/studentNotes'

export async function getStudentMaterialNote(materialId: string) {
  const { data } = await api.get<StudentMaterialNoteResponse>(
    `/notes/material/${materialId}`,
  )
  return data
}

export async function saveStudentMaterialNote(
  materialId: string,
  payload: SaveStudentMaterialNotePayload,
) {
  const { data } = await api.put<StudentMaterialNoteResponse>(
    `/notes/material/${materialId}`,
    payload,
  )
  return data
}

export async function deleteStudentMaterialNote(materialId: string) {
  const { data } = await api.delete<{ message: string }>(
    `/notes/material/${materialId}`,
  )
  return data
}

export async function getStudentSubjectClassNotes(subjectClassId: string) {
  const { data } = await api.get<StudentSubjectMaterialNotesResponse>(
    `/notes/subject-class/${subjectClassId}`,
  )
  return {
    notes: data.notes ?? [],
  }
}

export async function getStudentNotes() {
  const { data } = await api.get<StudentGlobalMaterialNotesResponse>('/notes')
  return {
    notes: data.notes ?? [],
  }
}
