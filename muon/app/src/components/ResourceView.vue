<template>
    <div v-if="resourceFound">
        <div v-if="header">
            <div v-html="header.content"></div>
            <hr />
        </div>
        <div :key="index + currentPage * pageSize" v-for="(entry, index) in currentPageContent">
            <div v-html="processEntry(entry)"></div>
            <div class="read-full-link" v-if="!entry.complete">
                <a href="#" @click="$event.preventDefault(); openEntry(entry)">{{ $t('preview.readMore') }}</a>
            </div>
            <hr v-if="index + 1 < currentPageContent.length" />
        </div>
        <div v-if="footer">
            <hr />
            <div v-html="footer.content"></div>
        </div>
        <div class="pagination" v-if="pageLinks.length > 1">
            <span v-if="currentPage > 0">
                <a href="#" @click="$event.preventDefault(); changePage(currentPage - 1)">
                    Prev
                </a>
                <a href="#" @click="$event.preventDefault(); changePage(0)" v-if="pageLinks[0].page > 0">
                    First
                </a>
            </span>
            <span v-for="pageLink in pageLinks" :key="pageLink.text">
                <a href="#" v-if="pageLink.page !== null" @click="$event.preventDefault(); changePage(pageLink.page)">
                    {{ pageLink.text }}
                </a>
                <span v-else>
                    {{ pageLink.text }}
                </span>
            </span>
            <span v-if="currentPage + 1 < pageCount">
                <a href="#" @click="$event.preventDefault(); changePage(pageCount - 1)" v-if="pageLinks[pageLinks.length - 1].page + 1 < pageCount">
                    Last
                </a>
                <a href="#" @click="$event.preventDefault(); changePage(currentPage + 1)">
                    Next
                </a>
            </span>
        </div>
    </div>
    <div v-else>
        <h1>{{ errorMessage }}</h1>
    </div>
</template>

<script>
import Configuration from '@/config'
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
        resourceFound: true,
        currentPage: 0
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
        },
        filteredContent () {
            return this.content.filter(entry => !entry.fixed)
        },
        currentPageContent () {
            return this.filteredContent.filter((entry, index) => index >= this.currentPage * Configuration.ui.pagination.pageSize &&
                index < (this.currentPage + 1) * Configuration.ui.pagination.pageSize)
        },
        header () {
            if (this.content.length > 0 && this.content[0].fixed) {
                return this.content[0]
            } else {
                return null
            }
        },
        footer () {
            if (this.content.length > 0 && this.content[this.content.length - 1].fixed) {
                return this.content[this.content.length - 1]
            } else {
                return null
            }
        },
        pageCount () {
            return Math.ceil(this.filteredContent.length / Configuration.ui.pagination.pageSize)
        },
        pageSize () {
            return Configuration.ui.pagination.pages
        },
        pageLinks () {
            const PAGES = Configuration.ui.pagination.pages
            const links = []

            let start = this.currentPage - PAGES
            if (start < 0) {
                start = 0
            }
            for (let i = start; i < this.currentPage; i++) {
                links.push({
                    text: `${i + 1}`,
                    page: i
                })
            }

            let end = this.currentPage + PAGES
            if (links.length < PAGES) {
                end += PAGES - links.length
            }
            if (end >= this.pageCount) {
                end = this.pageCount - 1
            }

            links.push({
                text: `${this.currentPage + 1}`,
                page: null
            })
            for (let i = this.currentPage + 1; i <= end; i++) {
                links.push({
                    text: `${i + 1}`,
                    page: i
                })
            }

            return links
        }
    },
    watch: {
        language () {
            this.loadArticle()
        },
        path () {
            this.currentPage = 0
            this.loadArticle()
        },
        '$route.query' () {
            if (typeof this.$route.query.page !== 'undefined') {
                this.currentPage = parseInt(this.$route.query.page)
            }
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
        changePage (page) {
            this.$router.push({
                path: this.$route.path,
                query: {
                    page
                }
            })
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
        if (typeof this.$route.query.page !== 'undefined') {
            this.currentPage = parseInt(this.$route.query.page)
        }
        this.loadArticle()
    }
}
</script>

<style scoped>
.read-full-link {
    text-align: end;
}

.pagination {
    display: flex;
    flex-direction: row;
    justify-content: center;
}

.pagination span {
    padding: 0px 1px 0px 1px;
}
</style>
