<template>
    <div class="box-body">
        <div class="fields-group">
            {{ template "form_components_layout" . }}
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Form',
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
