export interface CreateMaterialPayload {
  schoolId: string
  subjectClassId: string
  materialTitle: string
  materialDesc?: string
  materialType: 'video' | 'pdf' | 'ppt' | 'other'
  mediaIds: string[]
}

export interface MaterialItem {
  materialId: string
  materialTitle: string
  materialType: string
  createdAt: string
}

export interface MaterialListResponse {
  subjectClass: {
    subjectClassId: string
    subjectName: string
  }
  data: {
    data: MaterialItem[]
  }
}
