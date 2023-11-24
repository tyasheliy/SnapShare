import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import NotFoundView from '../views/NotFoundView.vue'
import ProfileView from '../views/ProfileView.vue'
import RegisterView from '../views/RegisterView.vue'
import ShareView from '../views/ShareView.vue'
import GetView from '../views/GetView.vue'
import AboutView from '../views/AboutView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfileView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: "/share",
    name: "share",
    component: ShareView
  },
  {
    path: "/get/:id",
    name: "get",
    component: GetView
  },
  {
    path: "/about",
    name: "about",
    component: AboutView
  },
  { path: '/:pathMatch(.*)*', component: NotFoundView }
]

const router = createRouter({ history: createWebHistory(process.env.BASE_URL), routes })

export default router
