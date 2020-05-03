import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import GoAdmin from '@/components/GoAdmin'
import Form from '@/components/Form'
import Datatable from '@/components/DataTable'
import Layout from "@/components/Layout";

Vue.use(Router);

export default new Router({
    mode: "history",
    base: "/admin/vue",
    routes: [
        {
            path: '/hello',
            name: 'HelloWorld',
            component: HelloWorld
        }, {
            path: '/goadmin',
            name: 'GoAdmin',
            component: GoAdmin
        }, {
            path: "/",
            component: Layout,
            children: [
                {
                    path: "/table",
                    name: "Table",
                    component: Datatable,
                    meta: {
                        title: "表格",
                        description: "表格",
                    },
                }, {
                    path: '/form',
                    name: 'Form',
                    component: Form,
                    meta: {
                        title: "表格",
                        description: "表格",
                    },
                }
            ],
        },
    ]
})
