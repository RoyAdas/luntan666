<template>
  <div>
    <h2>发帖</h2>
    <input v-model="title" placeholder="标题" /><br/><br/>
    <textarea v-model="content" rows="6" style="width:100%" placeholder="内容"></textarea><br/>
    <button @click="publish">发布</button>
  </div>
</template>

<script>
import { createPost } from '../api'
export default {
  data() { return { title:'', content:'' } },
  methods: {
    async publish() {
      await createPost({ title: this.title, content: this.content })
      this.$router.push('/')
    }
  }
}
</script>
