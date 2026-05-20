import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import type { RoleName } from '../types/auth'
import AuthLayout from '../layouts/AuthLayout.vue'
import StudentLayout from '../layouts/StudentLayout.vue'
import TeacherLayout from '../layouts/TeacherLayout.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import SuperAdminLayout from '../layouts/SuperAdminLayout.vue'
import LoginPage from '../pages/auth/LoginPage.vue'
import UnauthorizedPage from '../pages/auth/UnauthorizedPage.vue'
import StudentDashboard from '../pages/student/StudentDashboard.vue'
import TeacherDashboard from '../pages/teacher/TeacherDashboard.vue'
import AdminDashboard from '../pages/admin/AdminDashboard.vue'
import SuperAdminDashboard from '../pages/superadmin/SuperAdminDashboard.vue'

export const dashboardByRole: Record<RoleName, string> = {
  super_admin: '/superadmin/dashboard',
  admin: '/admin/dashboard',
  teacher: '/teacher/dashboard',
  student: '/student/dashboard',
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    {
      path: '/',
      component: AuthLayout,
      children: [
        {
          path: 'login',
          name: 'login',
          component: LoginPage,
        },
        {
          path: 'unauthorized',
          name: 'unauthorized',
          component: UnauthorizedPage,
        },
      ],
    },
    {
      path: '/student',
      component: StudentLayout,
      meta: { requiresAuth: true, roles: ['student'] },
      children: [
        {
          path: 'dashboard',
          name: 'student-dashboard',
          component: StudentDashboard,
        },
      ],
    },
    {
      path: '/teacher',
      component: TeacherLayout,
      meta: { requiresAuth: true, roles: ['teacher'] },
      children: [
        {
          path: 'dashboard',
          name: 'teacher-dashboard',
          component: TeacherDashboard,
        },
      ],
    },
    {
      path: '/admin',
      component: AdminLayout,
      meta: { requiresAuth: true, roles: ['admin'] },
      children: [
        {
          path: 'dashboard',
          name: 'admin-dashboard',
          component: AdminDashboard,
        },
      ],
    },
    {
      path: '/superadmin',
      component: SuperAdminLayout,
      meta: { requiresAuth: true, roles: ['super_admin'] },
      children: [
        {
          path: 'dashboard',
          name: 'superadmin-dashboard',
          component: SuperAdminDashboard,
        },
      ],
    },
  ],
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  auth.restoreSession()

  if (to.name === 'login' && auth.isAuthenticated) {
    const role = auth.primaryRole()
    return role ? dashboardByRole[role] : '/unauthorized'
  }

  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    return { name: 'login', query: { redirect: to.fullPath } }
  }

  const requiredRoles = to.matched.flatMap((record) => record.meta.roles ?? [])
  if (requiredRoles.length > 0 && !auth.hasAnyRole(requiredRoles)) {
    return { name: 'unauthorized' }
  }

  return true
})

export default router
