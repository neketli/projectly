<template>
    <div>
        <NuxtLayout>
            <NuxtPage />
        </NuxtLayout>
    </div>
</template>

<script setup lang="ts">
import 'element-plus/theme-chalk/dark/css-vars.css'

useHead({
    titleTemplate: (title) => {
        return title ? `${title} | Projectly` : 'Projectly'
    },
})
</script>

<style>
.slide-enter-active,
.slide-leave-active {
    transition: all 0.2s;
}
.slide-enter-from {
    opacity: 0;
    transform: translate(50px, 0);
}
.slide-leave-to {
    opacity: 0;
    transform: translate(-50px, 0);
}

@media only screen and (max-width: 500px) {
    .el-page-header__header {
        flex-direction: column;
        gap: 4px;
    }
}
</style>
