<template>
    <div v-if="resourceFound">
        <div :key="index" v-for="(entry, index) in content">
            <div v-html="processEntry(entry)"></div>
            <div class="read-full-link" v-if="!entry.complete">
                <a href="#" @click="$event.preventDefault(); openEntry(entry)">{{ $t('preview.readMore') }}</a>
            </div>
            <hr v-if="index + 1 < content.length" />
        </div>
    </div>
    <div v-else>
        <h1>{{ errorMessage }}</h1>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
    name: 'ResourceView',
    props: {
        path: {
            type: String,
            required: true
        }
    },
    data: () => ({
        content: [],
        paginate: ['articles'],
        resourceFound: true
    }),
    computed: {
        ...mapGetters({
            language: 'getLanguage'
        }),
        errorMessage () {
            return this.formatString(this.$t('error.notFound'), this.path)
        },
        variables () {
            const self = this
            return {
                email: atob(self.config.ui.email),
                apiUrl: self.config.url.api,
                staticUrl: self.config.url.static
            }
        }
    },
    watch: {
        language () {
            this.loadArticle()
        },
        path () {
            this.loadArticle()
        }
    },
    methods: {
        async loadArticle () {
            this.resourceFound = true
            try {
                this.content = await this.$store.dispatch('renderResource', this.path)
            } catch (err) {
                this.resourceFound = false
            }
        },
        openEntry (entry) {
            // Cut off the language from URL
            const url = entry.url.substring(entry.url.indexOf('/') + 1)
            this.$router.push(`/article/${encodeURIComponent(url)}`)
        },
        processEntry (entry) {
            let text = entry.content
            for (let key of Object.keys(this.variables)) {
                text = text.replace(new RegExp(`%${key}%`, 'g'), this.variables[key])
            }
            const base64Regex = /%base64:[A-Za-z0-9+/=]+%/g
            const matched = text.match(base64Regex) || []
            for (let match of matched) {
                const b64 = match.substring(8, match.length - 1)
                text = text.replace(match, atob(b64))
            }
            return text
        }
    },
    created () {
        this.loadArticle()
    }
}
</script>

<style scoped>
.read-full-link {
    text-align: end;
}
</style>
