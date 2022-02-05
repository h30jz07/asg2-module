<template>
    <div class="list">
        <h1>{{ header }}</h1><br>
        <input type="text" v-model="search" placeholder="Search Modules"/>
        <div v-if="modules.length">
            <div v-for="mod in filteredModules" :key="mod.moduleCode" class="module">
                <router-link :to="{ name: 'ModuleDetails', params: { code: mod.modulecode } }">
                    <Module :info="mod"/>
                </router-link>
            </div>
        </div>
        <div v-else>
            <p>Loading modules...</p>
        </div>
    </div>
</template>

<script>
import Module from './Module.vue'

export default {
    components: { Module },
    props: ['header'],
    data() {
        return {
            modules: [],
            search: ''
        }
    },
    mounted() {
        //GetAllModules() from 3.4
        fetch(`http://${process.env.VUE_APP_URL}:${process.env.VUE_APP_BACKEND_PORT}/module/v1/list`)
            .then(resp => resp.json())
            .then(data => this.modules = data)
            .catch(err => console.log(err.message))
        console.log(`http://${process.env.VUE_APP_URL}:${process.env.VUE_APP_BACKEND_PORT}/module/v1/list`)
    },
    computed: {
        filteredModules: function(){
            return this.modules.filter(module => {
                return module.modulename.toLowerCase().match(this.search.toLowerCase())
            })
        }
    }
}
</script>


<style>
    .module{
        background: #f4f4f4;
        padding: 20px;
        border-radius: 10px;
        margin: 10px auto;
        max-width: 600px;
        cursor: pointer;
        color: #444;
    }
    .module:hover {
        background: #ddd;
    }
    .module a{
        text-decoration: none;
    }
    input {
        border: 1px solid grey;
        border-radius: 5px;
        height: 20px;
        width: 90%;
        padding: 2px 23px 2px 30px;
        outline: 0;
        background-color: #f5f5f5;
    }
</style>
