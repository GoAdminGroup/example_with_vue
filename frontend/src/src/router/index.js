import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import GoAdmin from '@/components/GoAdmin'

Vue.use(Router);

export default new Router({
    mode: "history",
    base: "/admin/vue",
    routes: [
        {
            path: '/',
            name: 'HelloWorld',
            component: HelloWorld
        }, {
            path: '/goadmin',
            name: 'GoAdmin',
            component: GoAdmin
        }
    ]
})
