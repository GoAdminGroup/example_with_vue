<template>
    <div id="app">
        <div class="wrapper">
            <va-content-wrap></va-content-wrap>
        </div>
    </div>
</template>

<script>
    import VAContentWrap from 'ContentWrap.vue'
    import store from './vuex/store.js'

    export default {
        name: 'app',
        data() {
            return {}
        },
        created() {
        },
        components: {
            'va-content-wrap': VAContentWrap,
        },
        store
    }
</script>

<style>

</style>
