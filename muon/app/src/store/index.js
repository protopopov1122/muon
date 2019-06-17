import Configuration from '@/config'
import axios from 'axios'

const state = {
    apiUrl: Configuration.url.api,
    staticUrl: Configuration.url.static,
    layout: [],
    resourceLoading: 0,
    language: localStorage.getItem('language') || 'en'
}

const getters = {
    getApiUrl (state) {
        return state.apiUrl
    },
    getStaticUrl (state) {
        return state.staticUrl
    },
    getAppMenu (state) {
        return state.layout.menu
    },
    getAppHome (state) {
        return state.layout.home
    },
    getLanguage (state) {
        return state.language
    },
    isResourceLoading (state) {
        return state.resourceLoading > 0
    }
}

const mutations = {
    setLayout (state, layout) {
        state.layout = layout
    },
    setLanguage (state, lang) {
        state.language = lang
        localStorage.setItem('language', lang)
    },
    setResourceLoading (state, loading) {
        if (loading) {
            state.resourceLoading++
        } else {
            state.resourceLoading--
        }
    }
}

const actions = {
    async renderResource (store, path) {
        const api = store.getters['getApiUrl']
        store.commit('setResourceLoading', true)
        const response = await axios.get(`${api}/render/${store.getters['getLanguage']}/${path}`).finally(() => {
            store.commit('setResourceLoading', false)
        })
        if (response.data.error) {
            throw response.data.error
        } else {
            return response.data.result
        }
    },
    async loadResource (store, path) {
        const url = store.getters['getStaticUrl']
        const response = await axios.get(`${url}/${path}`)
        return response.data
    },
    async loadLayout (store) {
        const layout = await store.dispatch('loadResource', 'layout.json')
        store.commit('setLayout', layout)
    }
}

export default {
    state,
    getters,
    mutations,
    actions
}
