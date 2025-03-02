import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
    const authStore = useAuthStore()

    if (!authStore.isLogged && !to.meta.isPublic && to.path !== '/auth/login') {
        return navigateTo('/auth/login')
    }
})
