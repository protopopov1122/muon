<template>
    <div class="content-text-color">
        <div class="loading-backdrop loading-backdrop-background" v-if="showLoadingDialog">
            <div class="loading-dialog loading-dialog-background">
                <pulse-loader :loading="showLoadingDialog"></pulse-loader>
            </div>
        </div>
        <router-view></router-view>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PulseLoader from 'vue-spinner/src/PulseLoader'

export default {
    name: 'AppContent',
    components: {
        PulseLoader
    },
    data: () => ({
        showLoadingDialog: false,
        loadingDialogTimeout: null
    }),
    computed: {
        ...mapGetters({
            resourceLoading: 'isResourceLoading'
        })
    },
    watch: {
        resourceLoading () {
            const LOADING_DIALOG_DELAY = this.config.ui.loadingDialogTimeout
            if (this.resourceLoading && this.loadingDialogTimeout === null) {
                this.loadingDialogTimeout = setTimeout(() => {
                    this.showLoadingDialog = true
                }, LOADING_DIALOG_DELAY)
            } else if (!this.resourceLoading) {
                this.showLoadingDialog = false
                if (this.loadingDialogTimeout) {
                    clearTimeout(this.loadingDialogTimeout)
                    this.loadingDialogTimeout = null
                }
            }
        }
    }
}
</script>

<style scoped>
.loading-backdrop {
    position: fixed;
    top: 0px;
    left: 0px;
    width: 100vw;
    height: 100vh;
    z-index: 9999;
}

.loading-dialog {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: fit-content;
    height: fit-content;
    padding: 15px 35px 15px 35px;
    border-radius: 5%;
}
</style>
