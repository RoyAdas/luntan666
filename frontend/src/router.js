import { createRouter, createWebHistory } from 'vue-router'
import PostList from './components/PostList.vue'
import PostDetail from './components/PostDetail.vue'
import NewPost from './components/NewPost.vue'
import UserProfile from './components/UserProfile.vue'

const routes = [
  { path: '/', component: PostList },
  { path: '/post/:id', component: PostDetail },
  { path: '/new', component: NewPost },
  { path: '/profile', component: UserProfile }
]

export default createRouter({
  history: createWebHistory(),
  routes
})
