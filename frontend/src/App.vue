<template>
  <div style="max-width:900px;margin:20px auto;font-family:Arial,Helvetica;">
    <header style="display:flex;gap:20px;align-items:center;">
      <h1 style="margin:0">论坛666</h1>
      <nav style="margin-left:auto;">
        <router-link to="/">首页</router-link> |
        <router-link to="/new">发帖</router-link> |
        <router-link to="/profile">用户</router-link>
      </nav>
    </header>
    <hr/>
    <router-view/>
  </div>
</template>

<script>
export default { name: 'App' }
</script>
