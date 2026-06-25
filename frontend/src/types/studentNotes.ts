export interface StudentMaterialNote {
  noteId: string
  materialId: string
  content: string
  createdAt: string
  updatedAt: string
}

export interface StudentMaterialNoteResponse {
  note: StudentMaterialNote | null
}

export interface SaveStudentMaterialNotePayload {
  content: string
}

export interface StudentSubjectMaterialNote extends StudentMaterialNote {
  materialTitle: string
}

export interface StudentSubjectMaterialNotesResponse {
  notes: StudentSubjectMaterialNote[]
}

export interface StudentGlobalMaterialNote extends StudentSubjectMaterialNote {
  materialType: string
  subjectClassId: string
  subjectId: string
  subjectName: string
  subjectCode: string
  classId: string
  className: string
  classCode: string
}

export interface StudentGlobalMaterialNotesResponse {
  notes: StudentGlobalMaterialNote[]
}
