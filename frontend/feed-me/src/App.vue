<template>
    <div>
        <sandwich></sandwich>
    </div>
</template>

<script>
// Imports
// import addsandwich from './components/addSandwich.vue';
import sandwich from './components/sandwich.vue'

export default {
    components: {
        'sandwich': sandwich
    },
    data () {
        return {

        }
    },
    methods: {

    }
}
</script>

<style>
body{
    margin: 0;
    font-family: 'Nunito SemiBold';
}
</style>
