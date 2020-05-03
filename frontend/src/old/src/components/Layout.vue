<template>
    <div class="app-main-container">
        <section class="content-header">
            <h1>
                {{ title }}
                <small>{{ description }}</small>
            </h1>
            <ol class="breadcrumb" style="margin-right: 30px;">
                <li><a href="/admin"><i class="fa fa-dashboard"></i> 首页</a></li>
                <li>表格</li>
            </ol>
        </section>
        <section class="content">
            <router-view :key="key" style="min-height: 78vh;"/>
        </section>
    </div>
</template>


<script>

    export default {
        name: 'Layout',
        data() {
            return {}
        },
        computed: {
            title() {
                return this.$store.getters.title;
            },
            description() {
                return this.$store.getters.description;
            }
        },
    }
</script>