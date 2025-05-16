import type { AxiosError } from 'axios'
import { SHA256 } from 'crypto-js'
import type { User } from '~/types/user'

export const useUserActions = () => {
    const { $api, $i18n } = useNuxtApp()

    const getUserByEmail = async (email: string): Promise<User> => {
        try {
            const { data: user } = await $api.get<User>(`/user/${email}`)
            return user
        }
        catch (e) {
            const error = e as AxiosError<{ message: string }>
            throw new Error(error.response?.data?.message || error.message)
        }
    }

    const updateUserInfo = async (user: {
        name: string
        surname: string
        email: string
        birthday: string
        location: string
        about: string
    }) => {
        try {
            await $api.patch('/user/update', {
                ...user,
                language: $i18n.locale.value,
            })
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
        getUserByEmail,
        updateUserInfo,
        changePassword,
        uploadAvatar,
        removeAvatar,
    }
}
