export type RoleName = 'student' | 'teacher' | 'admin' | 'super_admin'

export interface UserInfo {
  id: string
  fullName: string
  email: string
}

export interface SchoolInfo {
  id: string
  code: string
  name: string
}

export interface MembershipInfo {
  schoolUserId: string
  school: SchoolInfo
  roles: RoleName[]
  isDefault: boolean
}

export interface DefaultContext {
  schoolId: string
  schoolUserId: string
  roles: RoleName[]
}

export interface LoginResponse {
  token: string
  user: UserInfo
  memberships: MembershipInfo[]
  globalRoles: RoleName[]
  defaultContext?: DefaultContext
}

export interface LoginPayload {
  email: string
  password: string
}
