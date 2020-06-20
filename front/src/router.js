import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from "./components/HelloWorld";
import Followers from "./components/Followers"

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/followers',
            name: 'followers',
            component: Followers
        },
        {
            path: '/',
            name: 'index',
            component: HelloWorld
        }
    ]
})
