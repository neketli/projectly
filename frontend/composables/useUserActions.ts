import type { AxiosError } from 'axios'
import { SHA256 } from 'crypto-js'

export const useUserActions = () => {
    const { $api } = useNuxtApp()

    const updateUserInfo = async (user: {
        name: string
        surname: string
        email: string
    }) => {
        try {
            await $api.patch('/user/update', user)
        }
        catch (e) {
            const error = e as AxiosError<{ message: string }>
            throw new Error(error.response?.data?.message || error.message)
        }
    }

    const changePassword = async (password: string) => {
        try {
            await $api.patch('/user/change-password', {
                password: SHA256(password).toString(),
            })
        }
        catch (e) {
            const error = e as AxiosError<{ message: string }>
            throw new Error(error.response?.data?.message || error.message)
        }
    }

    const uploadAvatar = async (avatar: File) => {
        const formData = new FormData()

        formData.append('image', avatar)

        try {
            await $api.post('/user/upload-avatar', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            })
        }
        catch (e) {
            const error = e as AxiosError<{ message: string }>
            throw new Error(error.response?.data?.message || error.message)
        }
    }

    const removeAvatar = async () => {
        try {
            await $api.delete('/user/remove-avatar')
        }
        catch (e) {
            const error = e as AxiosError<{ message: string }>
            throw new Error(error.response?.data?.message || error.message)
        }
    }

    return {
        updateUserInfo,
        changePassword,
        uploadAvatar,
        removeAvatar,
    }
}
