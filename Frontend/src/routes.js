import { createRouter, createWebHistory } from 'vue-router'
import EmailList from './components/EmailList.vue'
import EmailDetail from './components/EmailDetail.vue'

const routes = [
  { path: '/emails', component: EmailList },
  { path: '/email/:id', component: EmailDetail, props: true }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
