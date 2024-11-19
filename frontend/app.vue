<script setup lang="ts">
import 'element-plus/theme-chalk/dark/css-vars.css'
</script>

<template>
    <div>
        <NuxtLayout>
            <NuxtPage />
        </NuxtLayout>
    </div>
</template>

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
</style>
