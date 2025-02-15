import type { AxiosError } from 'axios'
import type { DetailedTask, Task } from '~/types/task'

export const useTask = () => {
    const { $api } = useNuxtApp()

    const getTask = async (id: number): Promise<DetailedTask> => {
        try {
            const { data } = await $api.get(`/task/${id}`)
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task')
        }
    }

    const getUserTasks = async (user_id: number): Promise<Task[]> => {
        try {
            const { data } = await $api.get(`/task/list-user`, {
                params: {
                    user_id,
                },
            })
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task')
        }
    }

    const createTask = async (data: {
        title: string
        description?: string
        priority: number
        story_points?: number
        deadline?: number
        tracked_time?: number
        status_id: number
        assigned_user_id?: number | string
        finished_at?: number
    }): Promise<DetailedTask> => {
        try {
            const { data: task } = await $api.post('/task/create', data)
            return task
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create task')
        }
    }

    const updateTask = async (updatedTask: Task): Promise<DetailedTask> => {
        try {
            const { data } = await $api.put(`/task/${updatedTask.id}`, updatedTask)
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update task')
        }
    }

    const updateTaskStatus = async (taskId: number, newStatusId: number, finishedAt: number | null): Promise<void> => {
        try {
            await $api.patch(`/task/${taskId}/change-status`, {
                status_id: newStatusId,
                finished_at: finishedAt,
            })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update task status')
        }
    }

    const deleteTask = async (id: number): Promise<void> => {
        try {
            await $api.delete(`/task/${id}`)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete task')
        }
    }

    const getTasksList = async (board_id: number): Promise<Record<number, DetailedTask[]>> => {
        try {
            const { data: tasks } = await $api.get<DetailedTask[]>(`/task/`, {
                params: {
                    board_id,
                },
            })

            return tasks.reduce((acc, curr) => {
                if (!acc[curr.status.id]) {
                    acc[curr.status.id] = []
                }

                acc[curr.status.id].push(curr)
                return acc
            }, {} as Record<number, DetailedTask[]>)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task list')
        }
    }

    const getTaskDetail = async (project_code: string, project_index: number): Promise<DetailedTask> => {
        try {
            const { data: task } = await $api.get<DetailedTask[]>(`/task/`, {
                params: {
                    project_code,
                    project_index,
                },
            })

            return task[0]
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task')
        }
    }

    return { getTask, getUserTasks, getTasksList, createTask, updateTask, updateTaskStatus, deleteTask, getTaskDetail }
}
