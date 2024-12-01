import type { AxiosError } from 'axios'
import type { Status } from '~/types/board'

export const useStatus = () => {
    const { $api } = useNuxtApp()

    const createStatus = async (data: Omit<Status, 'id'>): Promise<Status> => {
        try {
            const { data: status } = await $api.post('/status/create', data)
            return status
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create status')
        }
    }

    const updateStatus = async (updatedStatus: Status): Promise<Status> => {
        try {
            await $api.patch(`/status/${updatedStatus.id}`, {
                title: updatedStatus.title,
                order: updatedStatus.order,
                hex_color: updatedStatus.hex_color,
            })
            return updatedStatus
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update status')
        }
    }

    const deleteStatus = async (status: Status): Promise<void> => {
        try {
            await $api.delete(`/status/delete`, { params: { id: status.id, order: status.order } })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete status')
        }
    }

    const getStatusList = async (board_id: number): Promise<Status[]> => {
        try {
            const { data } = await $api.get(`/status/list`, {
                params: {
                    board_id,
                },
            })
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get status list')
        }
    }

    return { getStatusList, createStatus, updateStatus, deleteStatus }
}
