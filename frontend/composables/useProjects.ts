import type { AxiosError } from 'axios'
import type { Project } from '~/types/project'

export const useProjects = () => {
    const { $api } = useNuxtApp()

    const getProject = async (id: number): Promise<Project> => {
        try {
            const { data } = await $api.get(`/project/${id}`)
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get project')
        }
    }

    const getTeamProject = async (team_id: number, code: string): Promise<Project> => {
        try {
            const { data } = await $api.get(`/project/`, {
                params: {
                    team_id,
                    code,
                },
            })
            return data
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get project')
        }
    }

    const createProject = async (data: Omit<Project, 'id'>): Promise<Project> => {
        try {
            const { data: project } = await $api.post('/project/create', data)
            return project
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create project')
        }
    }

    const updateProject = async (updatedProject: Project): Promise<Project> => {
        try {
            await $api.patch(`/project/${updatedProject.id}`, {
                title: updatedProject.title,
                description: updatedProject.description,
            })
            return updatedProject
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update project')
        }
    }

    const deleteProject = async (id: number): Promise<void> => {
        try {
            await $api.delete(`/project/${id}`)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete project')
        }
    }

    const getProjectsList = async (params: {
        team_id?: number
        user_id?: number
    }): Promise<Project[]> => {
        try {
            const { data: projects } = await $api.get(`/project/list`, { params })
            return projects
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get project list')
        }
    }

    return {
        getProject,
        getTeamProject,
        getProjectsList,
        createProject,
        updateProject,
        deleteProject,
    }
}
