<template>
    <div class="vue-form">
        <div class="box box-">
            <div class="box-header with-border">
                <h3 class="box-title">新建</h3>
                <div class="box-tools">
                    <div class="btn-group pull-right" style="margin-right: 10px">
                        <a href="{{ infoUrl }}}" class="btn btn-sm btn-default form-history-back"><i class="fa fa-arrow-left"></i>&nbsp;返回</a>
                    </div>
                </div>
            </div>
            <div class="box-body">
                <form v-bind:action="url" v-bind:method="method" accept-charset="UTF-8" class="form-horizontal" pjax-container
                      v-bind:style="formStyle">
                    <div v-bind:class="formWrapperClass">

                        <!--                {{if ne (len .TabHeaders) 0}}-->
                        <!--                {{ template "form_layout_tab" . }}-->
                        <!--                {{else if ne (len .ContentList) 0}}-->
                        <!--                {{ template "form_layout_two_col" . }}-->
                        <!--                {{else if .Layout.Flow}}-->
                        <!--                {{ template "form_layout_flow" . }}-->
                        <!--                {{else}}-->
                        <!--                {{ template "form_layout_default" . }}-->
                        <!--                {{end}}-->

                    </div>

                    <!--            {{if ne .OperationFooter ""}}-->
                    <!--            <div class="box-footer">-->
                    <!--                {{.OperationFooter}}-->
                    <!--            </div>-->
                    <!--            {{end}}-->

                    <!--            {{range $key, $value := .HiddenFields}}-->
                    <!--            <input type="hidden" name="{{$key}}" value='{{$value}}'>-->
                    <!--            {{end}}-->
                    <div class="box-footer">
                        <div class="col-md-2 "></div>
                        <div class="col-md-8 ">
                            <div class="btn-group pull-right">
                                <button type="submit" class="btn  btn-primary" data-loading-text="&nbsp;保存">
                                    保存
                                </button>
                            </div>
                            <label class="pull-right" style="margin: 5px 10px 0 0;">
                                <div class="icheckbox_minimal-blue" aria-checked="false" aria-disabled="false" style="position: relative;"><input type="checkbox" class="continue_new" style="position: absolute; opacity: 0;"><ins class="iCheck-helper" style="position: absolute; top: 0%; left: 0%; display: block; width: 100%; height: 100%; margin: 0px; padding: 0px; background: rgb(255, 255, 255); border: 0px; opacity: 0;"></ins></div> 继续新增
                            </label>
                            <div class="btn-group pull-left">
                                <button type="reset" class="btn  btn-warning" data-loading-text="&nbsp;保存">
                                    重置
                                </button>
                            </div>
                            <script>
                                let previous_url_goadmin = $('input[name="__go_admin_previous_"]').attr("value");
                                $('.continue_new').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
                                    if (this.checked) {
                                        $('input[name="__go_admin_previous_"]').val(location.href)
                                    } else {
                                        $('input[name="__go_admin_previous_"]').val(previous_url_goadmin)
                                    }
                                });
                            </script>
                        </div>
                    </div>
                </form>
            </div>
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
        created: function () {
            console.log('created', this.msg);
            this.getPanelData()
        },
        computed: {
            formStyle: function () {
                let s = {
                    'background-color': 'white',
                };
                if (this.tabHeaders.length > 0) {
                    s["padding"] = "0"
                }
                return s
            },
            formWrapperClass: function () {
                if (this.tabHeaders.length > 0) {
                    return {"row": true}
                }
                return {"box-body": true}
            }
        },
        methods: {
            getPanelData() {
                this.$ajax({
                    method: 'get',
                    url: '/admin/api/update/form/manager?__goadmin_edit_pk=2',
                }).then((res) => {
                    console.log(res.data);
                    if (res.data.code === 200) {
                        let panel = res.data.data.panel;
                        this.tabHeaders = panel.group_field_headers;
                        this.tabContentList = panel.group_field_list;
                        this.url = res.data.data.urls.edit;
                        this.infoUrl = res.data.data.urls.info;
                        // this.thead = panel.thead;
                        // this.tbody = panel.info_list;
                        // this.footer = res.data.data.footer;

                        this.$store.commit('setTitle', panel.title);
                        this.$store.commit('setDescription', panel.description);
                    }
                }).catch((err) => {
                    console.log(err)
                })
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
