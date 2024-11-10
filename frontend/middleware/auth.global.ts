import { useAuthStore } from '~/store/auth'

export default defineNuxtRouteMiddleware((to) => {
    const { isLogged } = useAuthStore()

    console.log(isLogged, to.path)

    if (!isLogged && !to.meta.isPublic && to.path !== '/auth/login') {
        console.log(1)

        return navigateTo('/auth/login', { external: true })
    }
})
