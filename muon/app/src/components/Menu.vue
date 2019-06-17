<template>
    <div class="menu-box theme-background theme-text-color">
        <menu-item :key="item.path" v-for="item of menu"
            :label="$t(item.label)" :to="itemPath(item)"></menu-item>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
import MenuItem from '@/components/MenuItem'

export default {
    name: 'AppMenu',
    components: {
        MenuItem
    },
    computed: {
        ...mapGetters({
            menu: 'getAppMenu'
        })
    },
    methods: {
        itemPath ({ path }) {
            return `/article/${encodeURIComponent(path)}`
        }
    }
}
</script>

<style scoped>
.menu-box {
    width: 100%;
    height: fit-content;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
}
</style>
