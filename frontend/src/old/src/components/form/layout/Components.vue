<template>
    <div>

        <div v-if="contentList.length > 0">
            <div v-bind:class="colClass" v-for="content in contentList">
                <div class="box-body">
                    <div class="fields-group" v-for="data in content">
                        <div class='form-group divider' v-if="data.divider && data.divider_title !== ''">
                            <div class="col-sm-{{ data.head_width }} control-label divider-title">
                                {{ data.divider_title }}
                            </div>
                        </div>
                        <div class='col-sm-12 pb-3' v-if="data.divider">
                            <hr>
                        </div>
                        <input type="hidden" name="{{data.field}}" value='{{data.value}}' v-if="data.hide">
                        <div class="form-group" v-bind:style="formGroupStyle(data.width)">
                            <label v-if="data.head !== ''"
                                   for="{{data.field}}"
                                   v-bind:class="formGroupClass(data.head_width, data.must)">
                                {{data.head}}
                            </label>
                            <div class="col-sm-{{data.input_width}}">
                                {{template "form_components" $data}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div v-else-if="tabHeaders.length > 0">
            <div class="tab-pane {{if eq $key 0}}active{{end}}" v-for="(key, content) in tabContentList"
                 id="tab-form-{{key}}}">
                <div class="fields-group" v-for="data in content">
                    <div class='form-group divider' v-if="data.divider && data.divider_title !== ''">
                        <div class="col-sm-{{ data.head_width }} control-label divider-title">
                            {{ data.divider_title }}
                        </div>
                    </div>
                    <div class='col-sm-12 pb-3' v-if="data.divider">
                        <hr>
                    </div>
                    <input type="hidden" name="{{data.field}}" value='{{data.value}}' v-if="data.hide">
                    <div class="form-group" v-bind:style="formGroupStyle(data.width)">
                        <label v-if="data.head !== ''"
                               for="{{data.field}}"
                               v-bind:class="formGroupClass(data.head_width, data.must)">
                            {{data.head}}
                        </label>
                        <div class="col-sm-{{data.input_width}}">
                            {{template "form_components" $data}}
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div v-else-if="layout.flow">
            <div class="fields-group" v-for="data in content">
                <div class='form-group divider' v-if="data.divider && data.divider_title !== ''">
                    <div class="col-sm-{{ data.head_width }} control-label divider-title">
                        {{ data.divider_title }}
                    </div>
                </div>
                <div class='col-sm-12 pb-3' v-if="data.divider">
                    <hr>
                </div>
                <input type="hidden" name="{{data.field}}" value='{{data.value}}' v-if="data.hide">
                <div class="form-group" v-bind:style="formGroupStyle(data.width)">
                    <label v-if="data.head !== ''"
                           for="{{data.field}}"
                           v-bind:class="formGroupClass(data.head_width, data.must)">
                        {{data.head}}
                    </label>
                    <div class="col-sm-{{data.input_width}}">
                        {{template "form_components" $data}}
                    </div>
                </div>
            </div>
        </div>

        <div v-else="layout.flow">
            <div class="fields-group" v-for="data in content">
                <div class='form-group divider' v-if="data.divider && data.divider_title !== ''">
                    <div class="col-sm-{{ data.head_width }} control-label divider-title">
                        {{ data.divider_title }}
                    </div>
                </div>
                <div class='col-sm-12 pb-3' v-if="data.divider">
                    <hr>
                </div>
                <input type="hidden" name="{{data.field}}" value='{{data.value}}' v-if="data.hide">
                <div class="form-group" v-bind:style="formGroupStyle(data.width)">
                    <label v-if="data.head !== ''"
                           for="{{data.field}}"
                           v-bind:class="formGroupClass(data.head_width, data.must)">
                        {{data.head}}
                    </label>
                    <div class="col-sm-{{data.input_width}}">
                        {{template "form_components" $data}}
                    </div>
                </div>
            </div>
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
        computed: {
            colClass() {
                // col-md-{{divide 12 (len $.ContentList)}}
                return "";
            },
            formGroupStyle(width) {
                // {{data.width !== 0 ? 'width: ' + data.width + 'px;': ''}}
                return ""
            },
            formGroupClass(head_width, must) {
                // col-sm-{{data.head_width}} {{if $data.Must}}asterisk{{end}} control-label
                return {}
            }
        }
    }
</script>

<style scoped>

</style>
