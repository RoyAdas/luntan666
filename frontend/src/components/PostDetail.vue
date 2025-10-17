<template>
  <div v-if="post">
    <h2>{{ post.title || '（无标题）' }}</h2>
    <p>{{ post.content }}</p>
    <p>作者: {{ post.author }} | 点赞: {{ post.likes }}</p>
    <div style="margin-top:10px">
      <button @click="like">点赞</button>
      <button v-if="isAdmin" @click="del" style="margin-left:8px">删除（管理员）</button>
    </div>
  </div>
</template>

<script>
import { fetchPost, likePost, deletePost } from '../api'
export default {
  data() { return { post: null } },
  computed: {
    isAdmin() { return localStorage.getItem('role') === 'admin' }
  },
  async created() {
    const id = this.$route.params.id
    const res = await fetchPost(id)
    this.post = res.data
  },
  methods: {
    async like() {
      await likePost(this.post.id)
      const res = await fetchPost(this.post.id)
      this.post = res.data
    },
    async del() {
      if (!confirm('确定删除此帖子？')) return
      await deletePost(this.post.id)
      this.$router.push('/')
    }
  }
}
</script>
