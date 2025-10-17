<template>
  <div>
    <div style="display:flex;gap:8px;align-items:center;">
      <input v-model="search" placeholder="搜索标题或内容" />
      <button @click="load">搜索</button>
    </div>
    <ul style="list-style:none;padding:0">
      <li v-for="p in posts" :key="p.id" style="padding:8px;border-bottom:1px solid #eee">
        <router-link :to="'/post/' + p.id"><strong>{{p.title || '（无标题）'}}</strong></router-link>
        <div style="font-size:0.9em;color:#555">{{ p.content }}</div>
        <div style="margin-top:6px">
          <button @click="like(p.id)">点赞 ({{p.likes}})</button>
          <span style="margin-left:8px;color:#888">作者: {{p.author}}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import { fetchPosts, likePost } from '../api'

export default {
  data() {
    return { posts: [], search: '' }
  },
  created() { this.load() },
  methods: {
    async load() {
      const res = await fetchPosts(this.search)
      this.posts = res.data
    },
    async like(id) {
      await likePost(id)
      this.load()
    }
  }
}
</script>
