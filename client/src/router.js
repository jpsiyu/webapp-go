import VueRouter from 'vue-router'
import App from "./base/App"
import Miss from './base/Miss'

const routes = [
    { path: "/", component: App },
    { path: "*", component: Miss },
]

const createRouter = () => {
    return new VueRouter({
        mode: "history",
        routes,
    })
}

export { createRouter }