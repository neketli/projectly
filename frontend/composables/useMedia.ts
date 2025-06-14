import type { AxiosError } from 'axios'

export const useMedia = () => {
    const { $api } = useNuxtApp()

    const getMedia = async (path: string): Promise<Blob> => {
        try {
            const { data } = await $api.get(`/media/${path}`, {
                responseType: 'blob',
            })

            if (!(data instanceof Blob)) {
                throw new Error('Media is not a blob')
            }

            return data
        }
        catch (err: unknown) {
            const error = err as AxiosError<{ error: string }>

            if (error.response) {
                const errorMessage = error.response.data?.error
                  || `HTTP ${error.response.status}: ${error.response.statusText}`
                throw new Error(errorMessage)
            }
            else {
                throw new Error(error.message)
            }
        }
    }

    return { getMedia }
}
