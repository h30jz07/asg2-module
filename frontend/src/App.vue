<template>
  <div >
    <router-link to="/" class="topnav"><h1>Edu-fi Module Dashboard</h1></router-link>
  <router-view/>
  </div>
</template>

<script>
</script>

<style>
h1 {
  color: #cc6148;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
.topnav{
  background-color: #333;
  overflow: hidden;
  display: flexbox;
  text-decoration: none;
}
.topnav h1{
  background-color: #333;
  text-align: left;
  padding: 10px 16px ;
}
</style>
