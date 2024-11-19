import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
    const { isLogged } = useAuthStore()

    if (!isLogged && !to.meta.isPublic && to.path !== '/auth/login') {
        return navigateTo('/auth/login', { external: true })
    }
})
