import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
const state = {
    title: "",
    description: ""
};
const getters = {
    title(state) {
        console.log("getters title", state.title);
        return state.title
    },
    description() {
        console.log("getters description", state.description);
        return state.description
    }
};
const mutations = {
    setTitle(state, title) {
        console.log("setTitle", title);
        state.title = title;
    },
    setDescription(state, description) {
        console.log("setDescription", description);
        state.description = description;
    },
};
const store = new Vuex.Store({
    state,
    getters,
    mutations
});
export default store;
