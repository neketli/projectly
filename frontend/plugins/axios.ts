/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from 'axios'
import { jwtDecode } from 'jwt-decode'
import { useAuthStore } from '~/store/auth'

export default defineNuxtPlugin(({ $config }) => {
    const api = axios.create({
        baseURL: `${$config.public.API_URL}/api/v1`,
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Credentials': true,
            'Access-Control-Allow-Origin': '*',
        },
    })

    const userStore = useAuthStore()

    api.interceptors.request.use(
        async (config: any) => {
            if (!userStore.accessToken) {
                return config
            }

            try {
                const payload = jwtDecode(userStore.accessToken)

                const expTime = Number(payload.exp || 0) * 1000

                if (expTime - new Date().getTime() <= 0) {
                    userStore.accessToken = ''
                    await userStore.refresh()
                }

                config.headers.Authorization = `Bearer ${userStore.accessToken}`

                return config
            }
            catch (err: unknown) {
                const error = err as Error

                ElNotification({
                    title: 'Ошибка',
                    message: error.message,
                    type: 'error',
                })
                return config
            }
        },
        error => Promise.reject(error),
    )

    api.interceptors.response.use(
        (response: any) => response,
        (error: any) => {
            if (userStore.accessToken && error.response.status === 401) {
                error.config.headers.Authorization = ''
                userStore.logout()
                navigateTo('/auth/login', { external: true })
            }

            return Promise.reject(error)
        },
    )

    return {
        provide: {
            api,
        },
    }
})
