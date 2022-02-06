<template>
    <div style="width: 70%" class="tabs">
        <ul class="tabs_header">
            <li 
            v-for="title in tabTitles" 
            :key="title" 
            :class="{ selected: title == selectedTitle }"
            @click="selectedTitle = title">
            {{ title }}
            </li>
        </ul>

        <slot />
    </div>
</template>

<script>
import { ref, provide } from 'vue'
export default{
    setup(props, { slots }) {
        const tabTitles = ref(slots.default().map((tab) => tab.props.title))
        const selectedTitle = ref(tabTitles.value[0])

        provide("selectedTitle", selectedTitle)
        return {
            selectedTitle, tabTitles
        }
    },props: {
        tabs: {
            type: Array,
            required: true,
        },
        selected: {
            type: String,
            required: true,
        }
    },
}
</script>

<style scoped>
    .tabNav{
        display: flex;
        flex-direction: column;
        justify-content: center;  
    }

    .tabs {
        max-width: 400px;
        margin: 0 auto;
    }

    .tabs_header {
        margin-bottom: 10px;
        list-style: none;
        padding: 0;
        display: flex;
        justify-content: center;
    }

    .tabs_header li {
        width: 130px;
        text-align: center;
        padding: 10px 20px;
        margin-right: 10px;
        background-color: #ddd;
        border-radius: 5px;
        cursor: pointer;
        transition: 0.4s all ease-out;
    }

    .tabs_header li.selected {
        background-color: #333;
        color: #cc6148;
    }
</style>
