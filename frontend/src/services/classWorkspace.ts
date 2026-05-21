import { api } from './api'
import type {
  MaterialItem,
  MaterialListWithSubjectResponse,
  SubjectClassesResponse,
} from '../types/classWorkspace'

export async function getSubjectClassesByClass(classId: string) {
  const { data } = await api.get<SubjectClassesResponse>(`/subject-classes/class/${classId}`)
  return data
}

export async function getMaterialsBySubjectClass(subjectClassId: string) {
  const { data } = await api.get<MaterialListWithSubjectResponse>('/materials', {
    params: { subjectClassId, page: 1, limit: 20 },
  })
  return data
}

export async function getSubjectMaterials(subjectClassId: string) {
  const response = await getMaterialsBySubjectClass(subjectClassId)
  return {
    subjectClass: response.subjectClass,
    materials: response.data.data || [],
    pagination: {
      totalItems: response.data.totalItems,
      page: response.data.page,
      limit: response.data.limit,
      totalPages: response.data.totalPages,
    },
  }
}

export async function getMaterialById(materialId: string) {
  const { data } = await api.get<MaterialItem>(`/materials/${materialId}`)
  return data
}

export async function getClassMaterials(classId: string) {
  const subjectClasses = await getSubjectClassesByClass(classId)
  const materialResponses = await Promise.all(
    subjectClasses.subjects.map((subjectClass) =>
      getMaterialsBySubjectClass(subjectClass.subjectClassId),
    ),
  )

  const materials: MaterialItem[] = materialResponses.flatMap((response) => 
    response.data.data || []
  )
  
  return {
    classInfo: subjectClasses.class,
    subjects: subjectClasses.subjects,
    materials,
  }
}
