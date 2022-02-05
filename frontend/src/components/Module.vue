<template>
    <div  @click="toggleModule">
        <h2>{{ info.modulename }}</h2>
        <p>Module Code: {{ info.modulecode }}</p>
        <p>Synopsis: {{ info.synopsis }}</p>
        <p>Learning Objectives: {{ info.learningobjective }}</p>
    </div>
</template>

<script>
export default {
    props: ['info'],
    data() {
        return {
            showModule: false,
        }
    },
    methods: {
        toggleModule() {
            this.showModule = !this.showModule
            console.log("toggled modal")
        }
    }
}
</script>


<style scoped>
</style>
