<template>
  <div>
    <div v-if="!logged">
      <h3>登录</h3>
      <input v-model="username" placeholder="用户名" /><br/>
      <input v-model="password" type="password" placeholder="密码" /><br/>
      <button @click="login">登录</button>
    </div>
    <div v-else>
      <p>当前: {{ profile.username }} (角色: {{ profile.role }})</p>
      <h3>修改资料</h3>
      <input v-model="profile.username" placeholder="新用户名" /><br/>
      <input v-model="newPassword" type="password" placeholder="新密码" /><br/>
      <button @click="update">保存</button>
      <button @click="logout" style="margin-left:8px">登出</button>
    </div>
  </div>
</template>

<script>
import { login, getUser, updateUser } from '../api'
export default {
  data() { return {
    username:'', password:'', logged:false, profile:{}, newPassword:''
  }},
  async created() {
    const token = localStorage.getItem('token')
    if (token) {
      try {
        const res = await getUser()
        this.profile = res.data
        this.logged = true
      } catch (e) {
        localStorage.removeItem('token')
        localStorage.removeItem('role')
      }
    }
  },
  methods: {
    async login() {
      try {
        const res = await login(this.username, this.password)
        localStorage.setItem('token', res.data.token)
        localStorage.setItem('role', res.data.role)
        const u = await getUser()
        this.profile = u.data
        this.logged = true
      } catch (e) {
        alert('登录失败')
      }
    },
    async update() {
      await updateUser({ username: this.profile.username, password: this.newPassword })
      alert('已保存')
    },
    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('role')
      this.logged = false
      this.profile = {}
    }
  }
}
</script>
