import { SHA256 } from 'crypto-js'
import { defineStore } from 'pinia'
import { jwtDecode } from 'jwt-decode'
import type { AxiosError } from 'axios'
import type { User } from '~/types/user'

type AuthData = { email: string, password: string }
type Tokens = { access: string, refresh: string }

export const useAuthStore = defineStore('task-tracker-auth', {
    persist: {
        storage: persistedState.localStorage,
    },
    state: () => ({
        user: {} as User,
        accessToken: '',
        refreshToken: '',
        status: '',
        refreshingRequest: false as false | Promise<{ data: Tokens }>,
    }),
    getters: {
        getUserInfo(state) {
            return state.user
        },
        isLogged(state) {
            return !!state.accessToken
        },
    },
    actions: {
        authSuccess({ access, refresh }: Tokens) {
            this.status = 'success'
            this.accessToken = access
            this.refreshToken = refresh
            this.user = <User>jwtDecode(access)
        },

        async login({ email, password }: AuthData) {
            this.status = 'loading'

            const { $api } = useNuxtApp()
            try {
                const { data } = await $api.post<Tokens>('/auth/login', {
                    email,
                    password: SHA256(password).toString(),
                })

                this.authSuccess(data)
            }
            catch (e) {
                const error = e as AxiosError<{ message: string }>
                throw new Error(error.response?.data?.message || error.message)
            }
        },

        async register(user: {
            name: string
            surname: string
            email: string
            password: string
        }) {
            this.status = 'loading'

            const { $api } = useNuxtApp()
            try {
                await $api.post<{ data: Tokens }>(
                    '/auth/register',
                    {
                        ...user,
                        password: SHA256(user.password).toString(),
                    },
                )
            }
            catch (e) {
                const error = e as AxiosError<{ message: string }>
                throw new Error(error.response?.data?.message || error.message)
            }
        },

        logout() {
            this.$reset()
        },

        async refresh() {
            const { $api } = useNuxtApp()

            try {
                if (!this.refreshingRequest) {
                    this.refreshingRequest = $api.post<Tokens>(
                        '/user/refresh',
                        {
                            refreshToken: this.refreshToken,
                        },
                    )
                }

                const { data } = await this.refreshingRequest

                this.authSuccess(data)
                this.refreshingRequest = false
            }
            catch (error) {
                this.$reset()
                throw error
            }
        },
    },
})
