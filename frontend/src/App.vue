<template>
  <div class="topnav">
    <router-link to="/"><h1>Edu-fi Module Dashboard</h1></router-link>
  </div>
  <router-view/>
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
  color: #2c3e50;
}
.topnav{
  background-color: #333;
  overflow: hidden;
  display: flexbox;
}
.topnav h1{
  background-color: #333;
  text-align: left;
  padding: 0 16px ;
  text-decoration: none;
}
</style>
