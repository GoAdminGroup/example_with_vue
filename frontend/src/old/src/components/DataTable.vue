<template>
    <div class="box box-">
        <div class="box-header with-border">
            <div class="pull-right">
                <div class="dropdown pull-right column-selector" style="margin-right: 10px">
                    <button type="button" class="btn btn-sm btn-instagram dropdown-toggle"
                            data-toggle="dropdown">
                        <i class="fa fa-table"></i>
                        <span class="caret"></span>
                    </button>
                    <ul class="dropdown-menu" role="menu"
                        style="padding: 10px;max-height: 400px;overflow: scroll;">
                        <li>
                            <ul style="padding: 0;">
                                <li class="checkbox icheck" style="margin: 0;" v-for="item in thead">
                                    <label style="width: 100%;padding: 3px;">
                                        <input type="checkbox" class="column-select-item" data-id="{{item.field}}"
                                               style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;{{item.head}}
                                    </label>
                                </li>
                            </ul>
                        </li>
                        <li class="divider">
                        </li>
                        <li class="text-right">
                            <button class="btn btn-sm btn-default column-select-all">全部</button>&nbsp;&nbsp;
                            <button class="btn btn-sm btn-primary column-select-submit">提交</button>
                        </li>
                    </ul>
                </div>
                <div class="btn-group pull-right" style="margin-right: 10px">
                </div>
            </div>
            <span><a class="btn btn-sm btn-primary grid-refresh"><i class="fa fa-refresh"></i> 刷新</a></span>
            <script>
                let toastMsg = '刷新成功 !';
                $('.grid-refresh').on('click', function () {
                    $.pjax.reload('#pjax-container');
                    toastr.success(toastMsg);
                });
            </script>
        </div>
        <div class="box-body" style="overflow-x: scroll;overflow-y: hidden;padding:0;">
            <table class="table table-hover" style="min-width: 1000px;table-layout: auto;">
                <tbody>
                <tr>
                    <th style="text-align: center;">
                        <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
                    </th>
                    <th v-for="item in thead">
                        {{ item.head }}
                    </th>
                </tr>
                <tr v-for="(row, index) in tbody">
                    <td style="text-align: center;">
                        <input type="checkbox" class="grid-row-checkbox" data-id="{{index}}"
                               style="position: absolute; opacity: 0;">
                    </td>
                    <td v-for="item in thead" v-html="row[item.field].content"></td>
                </tr>
                </tbody>
            </table>
            <style>
                table tbody tr td {
                    word-wrap: break-word;
                    word-break: break-all;
                }
            </style>
        </div>
        <div class="box-footer clearfix" v-html="footer">
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Table',
        data() {
            return {
                footer: '',
                thead: [],
                tbody: []
            }
        },
        created: function () {
            console.log('created');
            this.getPanelData();
        },
        methods: {
            getPanelData() {
                this.$ajax({
                    method: 'get',
                    url: '/admin/api/list/manager',
                }).then((res) => {
                    console.log(res.data);
                    if (res.data.code === 200) {
                        let panel = res.data.data.panel;
                        this.thead = panel.thead;
                        this.tbody = panel.info_list;
                        this.footer = res.data.data.footer;

                        this.$store.commit('setTitle', panel.title);
                        this.$store.commit('setDescription', panel.description);
                    }
                }).catch((err) => {
                    console.log(err)
                })
            }
        }
    }

    window.selectedRows = function () {
        let selected = [];
        $('.grid-row-checkbox:checked').each(function () {
            selected.push($(this).data('id'));
        });
        return selected;
    };
    const selectedAllFieldsRows = function () {
        let selected = [];
        $('.column-select-item:checked').each(function () {
            selected.push($(this).data('id'));
        });
        return selected;
    };
    const pjaxContainer = "#pjax-container";
    const noAnimation = "__go_admin_no_animation_";

    function iCheck(el) {
        el.iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
            if (this.checked) {
                $(this).closest('tr').css('background-color', "#ffffd5");
            } else {
                $(this).closest('tr').css('background-color', '');
            }
        });
    }

    $(function () {
        $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
            if (this.checked) {
                $('.grid-row-checkbox').iCheck('check');
            } else {
                $('.grid-row-checkbox').iCheck('uncheck');
            }
        });
        let items = $('.column-select-item');
        iCheck(items);
        iCheck($('.grid-row-checkbox'));
        let columns = getQueryVariable("__columns");
        if (columns === -1) {
            items.iCheck('check');
        } else {
            let columnsArr = columns.split(",");
            for (let i = 0; i < columnsArr.length; i++) {
                for (let j = 0; j < items.length; j++) {
                    if (decodeURI(columnsArr[i]) === $(items[j]).attr("data-id")) {
                        $(items[j]).iCheck('check');
                    }
                }
            }
        }
        let lastTd = $("table tr:last td:last div");
        if (lastTd.hasClass("dropdown")) {
            let popUpHeight = $("table tr:last td:last div ul").height();
            let trs = $("table tr");
            let totalHeight = 0;
            for (let i = 1; i < trs.length - 1; i++) {
                totalHeight += $(trs[i]).height();
            }
            if (popUpHeight > totalHeight) {
                let h = popUpHeight + 16;
                $("table tbody").append("<tr style='height:" + h + "px;'></tr>");
            }
            trs = $("table tr");
            for (let i = trs.length - 1; i > 1; i--) {
                let td = $(trs[i]).find("td:last-child div");
                let combineHeight = $(trs[i]).height() / 2 - 20;
                for (let j = i + 1; j < trs.length; j++) {
                    combineHeight += $(trs[j]).height();
                }
                if (combineHeight < popUpHeight) {
                    td.removeClass("dropdown");
                    td.addClass("dropup");
                }
            }
        }
        let sort = getQueryVariable("__sort");
        let sort_type = getQueryVariable("__sort_type");
        if (sort !== -1 && sort_type !== -1) {
            let sortFa = $('#sort-' + sort);
            if (sort_type === 'asc') {
                sortFa.attr('href', '?__sort=' + sort + "&__sort_type=desc")
            } else {
                sortFa.attr('href', '?__sort=' + sort + "&__sort_type=asc")
            }
            sortFa.removeClass('fa-sort');
            sortFa.addClass('fa-sort-amount-' + sort_type);
        }
    });
    $('.column-select-all').on('click', function () {
        if ($(this).data('check') === '') {
            $('.column-select-item').iCheck('check');
            $(this).data('check', 'true')
        } else {
            $('.column-select-item').iCheck('uncheck');
            $(this).data('check', '')
        }
    });
    $('.column-select-submit').on('click', function () {
        let param = new Map();
        param.set('__columns', selectedAllFieldsRows().join(','));
        param.set(noAnimation, 'true');
        $.pjax({
            url: addParameterToURL(param),
            container: pjaxContainer
        });
        toastr.success('加载成功 !');
    });

    function getQueryVariable(variable) {
        let query = window.location.search.substring(1);
        let vars = query.split("&");
        for (let i = 0; i < vars.length; i++) {
            let pair = vars[i].split("=");
            if (pair[0] === variable) {
                return pair[1];
            }
        }
        return -1;
    }

    function addParameterToURL(params) {
        let newUrl = location.href.replace("#", "");
        for (let [field, value] of params) {
            if (getQueryVariable(field) !== -1) {
                newUrl = replaceParamVal(newUrl, field, value);
            } else {
                if (newUrl.indexOf("?") > 0) {
                    newUrl = newUrl + "&" + field + "=" + value;
                } else {
                    newUrl = newUrl + "?" + field + "=" + value;
                }
            }
        }
        return newUrl
    }

    function replaceParamVal(oUrl, paramName, replaceWith) {
        let re = eval('/(' + paramName + '=)([^&]*)/gi');
        return oUrl.replace(re, paramName + '=' + replaceWith);
    }

    $(function () {
        $('.editable-td-select').editable({
            "type": "select",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
        });
        $('.editable-td-text').editable({
            emptytext: "<i class=\"fa fa-pencil\"><\/i>",
            type: "text"
        });
        $('.editable-td-datetime').editable({
            "type": "combodate",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
            "format": "YYYY-MM-DD HH:mm:ss",
            "viewformat": "YYYY-MM-DD HH:mm:ss",
            "template": "YYYY-MM-DD HH:mm:ss",
            "combodate": {"maxYear": 2035}
        });
        $('.editable-td-date').editable({
            "type": "combodate",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
            "format": "YYYY-MM-DD",
            "viewformat": "YYYY-MM-DD",
            "template": "YYYY-MM-DD",
            "combodate": {"maxYear": 2035}
        });
        $('.editable-td-year').editable({
            "type": "combodate",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
            "format": "YYYY",
            "viewformat": "YYYY",
            "template": "YYYY",
            "combodate": {"maxYear": 2035}
        });
        $('.editable-td-month').editable({
            "type": "combodate",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
            "format": "MM",
            "viewformat": "MM",
            "template": "MM",
            "combodate": {"maxYear": 2035}
        });
        $('.editable-td-day').editable({
            "type": "combodate",
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
            "format": "DD",
            "viewformat": "DD",
            "template": "DD",
            "combodate": {"maxYear": 2035}
        });
        $('.editable-td-textarea').editable({
            "type": "textarea",
            "rows": 10,
            "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
        });
        $(".info_edit_switch").bootstrapSwitch({
            onSwitchChange: function (event, state) {
                let obejct = $(event.target);
                let val = "";
                if (state) {
                    val = obejct.closest('.bootstrap-switch').next().val();
                } else {
                    val = obejct.closest('.bootstrap-switch').next().next().val()
                }
                $.ajax({
                    method: 'post',
                    url: obejct.data("updateurl"),
                    data: {
                        name: obejct.data("field"),
                        value: val,
                        pk: obejct.data("pk")
                    },
                    success: function (data) {
                        if (typeof (data) === "string") {
                            data = JSON.parse(data);
                        }
                        if (data.code !== 200) {
                            swal(data.msg, '', 'error');
                        }
                    },
                    error: function (data) {
                        if (data.responseText !== "") {
                            swal(data.responseJSON.msg, '', 'error');
                        } else {
                            swal("错误", '', 'error');
                        }
                    },
                });
            }
        })
    });

    let gridPerPaper = $('.grid-per-pager');
    gridPerPaper.on('change', function () {
        $.pjax({url: this.value, container: '#pjax-container'});
    });
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
