import Vuex from 'vuex'

const person = {
    state: {},
    mutations: {},
}

const behavior = {
    state: {},
    mutations: {},
}

const createStore = () => {
    return new Vuex.Store({
        modules: {
            person,
            behavior,
        }
    })
}

export { createStore }