import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'
import DashboardLayout from '@/layouts/DashboardLayout'
import SignInView from '@/views/SignInView'
import UpdatePasswordView from '@/views/UpdatePasswordView'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect (to) {
      if (store.state.data.isAuthenticated) {
        return {
          name: 'UpdatePassword'
        }
      }

      return {
        name: 'SignIn'
      }
    }
  },
  {
    path: '/auth/sign-in',
    name: 'SignIn',
    component: SignInView
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: {
      requiresAuth: true
    },
    children: [
      {
        path: '/update-password',
        name: 'UpdatePassword',
        component: UpdatePasswordView
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.state.data.isAuthenticated) {
      return next({
        name: 'SignIn'
      })
    }
  }

  return next()
})

export default router
