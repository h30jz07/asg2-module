<template>
  <List :header="listHeader"/>
</template>

<script>
import List from '../components/List.vue'

export default {
  name: 'App',
  data() {
    return {
      listHeader: "View all modules"
    }
  },
  components: { List },
  methods: {
    search() {
      console.log(this.$refs.query)
    }
  }
}
</script>
<style scoped>

h1 {
  border-bottom: 1px solid #ddd;
  display: inline-block;
  padding-bottom: 10px;
}
</style>
