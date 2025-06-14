import type { AxiosError } from 'axios'
import type { DetailedTask, Task, TaskComment } from '~/types/task'

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

    const searchTask = async (query:
    { search: string, board_id?: number, project_code?: string, team_id?: number },
    ): Promise<DetailedTask[]> => {
        try {
            const { data: tasks } = await $api.get<DetailedTask[]>(`/task/`, {
                params: {
                    ...query,
                    limit: 5,
                },
            })

            return tasks
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task')
        }
    }

    const getAttachments = async (task_id: number): Promise<string[]> => {
        try {
            const { data: attachments } = await $api.get<{ id: number, name: string }[]>(`/task/${task_id}/attachments`)

            return attachments.map(item => item.name)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get task attachments')
        }
    }

    const createAttachments = async (task_id: number, file: File): Promise<string[]> => {
        const formData = new FormData()

        formData.append('files', file)

        try {
            const { data: attachments } = await $api.post(`/task/${task_id}/create-attachments`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            })

            return attachments
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create task attachments')
        }
    }

    const deleteAttachment = async (filename: string): Promise<void> => {
        try {
            await $api.delete(`/task/delete-attachment`, {
                params: {
                    filename,
                },
            })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete task attachment')
        }
    }

    const getComments = async (task_id: number, last_comment_id?: number): Promise<TaskComment[]> => {
        try {
            const { data: comments } = await $api.get<TaskComment[]>(`/task/${task_id}/comments`, {
                params: {
                    last_comment_id,
                },
            })
            return comments
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get comments')
        }
    }

    const createComment = async (task_id: number, text: string): Promise<void> => {
        try {
            await $api.post<TaskComment>(`/task/${task_id}/create-comment`, { text })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create comment')
        }
    }

    const deleteComment = async (task_id: number, comment_id: number): Promise<void> => {
        try {
            await $api.delete(`/task/${task_id}/delete-comment`, {
                params: {
                    comment_id,
                },
            })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete comment')
        }
    }

    return {
        getTask,
        searchTask,
        getUserTasks,
        getTasksList,
        createTask,
        updateTask,
        updateTaskStatus,
        deleteTask,
        getTaskDetail,

        getAttachments,
        createAttachments,
        deleteAttachment,

        getComments,
        createComment,
        deleteComment,
    }
}
