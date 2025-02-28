<script setup lang="ts">
const { t } = useI18n()

useHead({
    title: t('auth.title'),
})
definePageMeta({
    isPublic: true,
    layout: 'empty',
    pageTransition: { name: 'slide' },
})

const route = useRoute()
const authStore = useAuthStore()

onMounted(() => {
    const { access, refresh } = route.query

    if (typeof access === 'string' && refresh === 'string') {
        window.history.replaceState({}, document.title, '/')
        authStore.authSuccess({ access, refresh })
        navigateTo('/', { external: true })
    }
    else {
        navigateTo('/auth/login', { external: true })
    }
})
</script>
