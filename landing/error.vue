<script setup>
useHead({
    title: 'Projectly — Удобное управление проектами',
    meta: [
        { name: 'description', content: 'Projectly помогает командам эффективно планировать, отслеживать и завершать проекты с помощью Kanban-досок.' },
        { name: 'keywords', content: 'управление проектами, Kanban, таск-трекер, трекер задач, проектный менеджмент' },
        { property: 'og:title', content: 'Projectly — Удобное управление проектами' },
        { property: 'og:description', content: 'Projectly помогает командам эффективно планировать, отслеживать и завершать проекты с помощью Kanban-досок.' },
        { property: 'og:image', content: 'https://projectly.ru/media/og-image.jpg' },
        { property: 'og:url', content: 'https://projectly.ru' },
        { name: 'yandex-verification', content: '27f59918e45e4cc0' },
    ],
    link: [
        { rel: 'canonical', href: 'https://projectly.ru' },
    ],
})

navigateTo('/')
</script>

<template>
    <div class="font-sans text-gray-800 antialiased !scroll-pt-24">
        <AppHeader />
        <NuxtLayout>
            <NuxtPage />
        </NuxtLayout>
        <AppFooter />
    </div>
</template>

<style>
html {
    scroll-behavior: smooth;
}
</style>
