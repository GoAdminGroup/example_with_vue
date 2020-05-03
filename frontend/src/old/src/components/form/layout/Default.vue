<template>
    <div class="box-body">
        <div class="fields-group">
            <goadmin-form-components></goadmin-form-components>
        </div>
    </div>
</template>

<script>
    import Components from "@/components/form/layout/Components";

    export default {
        name: 'Form',
        components: {
            Components,
        },
        data() {
            return {
                tabHeaders: [],
                tabContentList: [],
                contentList: [],
                layout: {},
                url: "",
                infoUrl: "",
                method: "",
                hiddenFields: [],
                formStyle: "",
            }
        },
    }
</script>

<style scoped>

</style>
