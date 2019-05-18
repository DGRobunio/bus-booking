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
      path: '/user',
      name: 'User',
      component: () => import('./views/User')
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
