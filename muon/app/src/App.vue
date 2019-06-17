<template>
  <div id="app" class="app-background">
    <app-header class="app-header"></app-header>
    <div class="app-pane">
      <app-content class="app-contents"></app-content>
    </div>
    <app-footer class="app-footer"></app-footer>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import AppHeader from '@/components/Header'
import AppContent from '@/components/Content'
import AppFooter from '@/components/Footer'

export default {
  name: 'app',
  components: {
    AppHeader,
    AppContent,
    AppFooter
  },
  computed: {
    ...mapGetters({
      appHome: 'getAppHome',
      language: 'getLanguage'
    })
  },
  created () {
    this.$store.dispatch('loadLayout').then(() => {
      if (this.$route.fullPath === '/') {
        this.$router.replace(`/article/${encodeURIComponent(this.appHome)}`)
      }
    })
    this.$i18n.locale = this.language
    document.title = this.$t('title')
    const firstVisit = window.localStorage.getItem('language') === null
    if (firstVisit) {
        window.localStorage.setItem('language', this.language)
        const msg = this.$t('message.firstVisit')
        if (msg) {
          setTimeout(() => alert(msg), 0)
        }
    }
  }
}
</script>

<style scoped>
#app {
  display: flex;
  flex-direction: column;
  position: absolute;
  left: 0px;
  top: 0px;
  min-width: 100vw;
  min-height: 100vh;
}

.app-header {
  min-height: fit-content;
  flex-grow: 0;
}

.app-pane {
  flex-grow: 1;
  display: flex;
  justify-content: center;
}

.app-menu {
  width: 100vw;
}

.app-contents {
  flex-grow: 1;
  max-width: 1300px;
  margin: 0px 15px 0px 15px;
}

.app-footer {
  max-height: 30px;
  flex-grow: 0;
}
</style>
