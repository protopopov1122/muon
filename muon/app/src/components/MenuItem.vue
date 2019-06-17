<template>
    <div class="menu-item" @click="openRoute">
        {{ label }}
    </div>
</template>

<script>
export default {
    name: 'MenuItem',
    props: {
        label: {
            required: true,
            type: String
        },
        to: {
            required: true,
            type: String
        }
    },
    methods: {
        openRoute () {
            this.$router.push(this.to)
        }
    }
}
</script>

<style scoped>
.menu-item {
    display: inline-block;
    padding: 5px;
    cursor: pointer;
    width: fit-content;
    text-decoration: underline;
}
</style>
