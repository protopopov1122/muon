<template>
    <div class="header theme-background theme-text-color">
        <div class="title-container">
            <div class="title">{{ $t('header.title') }}</div>
            <div>
                <a href="#" class="language-link theme-text-color" :key="lang" v-for="lang of languages" @click="$event.preventDefault(); setLanguage(lang)">{{ lang }}</a>
            </div>
        </div>
        <div class="subtitle cursive-font">{{ $t('header.subtitle') }}</div>
        <app-menu></app-menu>
    </div>
</template>

<script>
import AppMenu from '@/components/Menu'
export default {
    name: 'AppHeader',
    components: {
        AppMenu
    },
    computed: {
        languages () {
            return this.config.ui.languages
        }
    },
    methods: {
        setLanguage (lang) {
            this.$i18n.locale = lang
            document.title = this.$t('title')
            this.$store.commit('setLanguage', lang)
        }
    }
}
</script>

<style scoped>
.header {
    padding: 10px 25px 0px 25px;
    display: flex;
    flex-direction: column;
}

.title-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.title {
    display: inline-block;
    width: fit-content;
    font-size: x-large;
    font-weight: bold;
}

.subtitle {
    display: inline-block;
    width: fit-content;
    font-size: larger;
}

.language-link {
    text-transform: uppercase;
    margin-left: 8px;
}
</style>
