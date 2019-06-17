<template>
    <div class="footer theme-background theme-text-color">
        <div>&copy; {{ $t('footer.copyright') }}</div>
        <a :href="emailHref" class="theme-text-color">{{ $t('footer.contactLink') }}</a>
    </div>
</template>

<script>
export default {
    name: 'AppFooter',
    computed: {
        emailHref () {
            return `mailto:${atob(this.config.ui.email)}`
        }
    }
}
</script>

<style scoped>
.footer {
    font-size: smaller;
    padding: 5px 25px 5px 15px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    flex-wrap: wrap;
}
</style>
