import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('./views/Home')
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('./views/Register')
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('./views/Login')
    },
    {
      path: '/logout',
      name: 'Logout',
      component: () => import('./views/Logout')
    },
    {
      path: '/user',
      name: 'User',
      component: () => import('./views/User')
    },
    {
      path: '/adminupdatebus/:busID',
      name: 'AdminUpdateBus',
      component: () => import('./views/AdminUpdateBus')
    },
    {
      path: '/adminaddbus',
      name: 'AdminAddBus',
      component: () => import('./views/AdminUpdateBus')
    },
    {
      path: '/updateorders',
      name: 'UpdateOrders',
      component: () => import('./views/UpdateOrders')
    },
    {
      path: '/codeslist',
      name: 'CodesList',
      component: () => import('./views/CodesList')
    },
    {
      path: '/orders',
      name: 'Orders',
      component: () => import('./views/Orders')
    },
    {
      path: '/comment/:orderID',
      name: 'Comment',
      component: () => import('./views/Comment')
    },
    {
      path: '/purchase/:busID',
      name: 'Purchase',
      component: () => import('./views/Purchase')
    },
    {
      path: '/favorites',
      name: 'Favorites',
      component: () => import('./views/Favorites')
    },
    {
      path: '/bus/:busID',
      name: 'BusInfo',
      component: () => import('./views/BusInfo')
    },
    {
      path: '/404',
      name: 'not-found',
      component: () => import('./views/NotFound')
    },
    {
      path: '*',
      redirect: '/404'
    }
  ]
})
