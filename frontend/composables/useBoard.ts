import type { AxiosError } from 'axios'
import { defaultStatusColors, type Board } from '~/types/board'

export const useBoard = () => {
    const { $api } = useNuxtApp()
    const { t } = useI18n()

    const getBoard = async (id: number): Promise<Board> => {
        try {
            const { data } = await $api.get(`/board/${id}`)
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get board')
        }
    }

    const getUserBoards = async (user_id: number): Promise<Board[]> => {
        try {
            const { data } = await $api.get(`/board/list-user`, {
                params: {
                    user_id,
                },
            })
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get board')
        }
    }

    const createBoard = async (data: Omit<Board, 'id'>): Promise<Board> => {
        try {
            const { data: board } = await $api.post('/board/create', data)
            return board
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create board')
        }
    }

    const setupDefaultBoard = async (board_id: number): Promise<void> => {
        const { createStatus } = useStatus()

        const statuses = [
            {
                title: t('status.default.todo'),
                order: 0,
                hex_color: defaultStatusColors[0],
                board_id,
            },
            {
                title: t('status.default.in_progress'),
                order: 1,
                hex_color: defaultStatusColors[1],
                board_id,
            },
            {
                title: t('status.default.done'),
                order: 2,
                hex_color: defaultStatusColors[2],
                board_id,
            },
        ]

        for (const status of statuses) {
            await createStatus(status)
        }
    }

    const updateBoard = async (updatedBoard: Board): Promise<Board> => {
        try {
            await $api.patch(`/board/${updatedBoard.id}`, {
                title: updatedBoard.title,
            })
            return updatedBoard
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update board')
        }
    }

    const deleteBoard = async (id: number): Promise<void> => {
        try {
            await $api.delete(`/board/${id}`)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete board')
        }
    }

    const getBoardsList = async (project_id: number): Promise<Board[]> => {
        try {
            const { data: boards } = await $api.get(`/board/list`, {
                params: {
                    project_id,
                },
            })
            return boards
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get board list')
        }
    }

    return { getBoard, getUserBoards, getBoardsList, createBoard, setupDefaultBoard, updateBoard, deleteBoard }
}
