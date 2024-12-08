import type { AxiosError } from 'axios'
import type { Team, TeamProject } from '~/types/team'
import type { Role, UserWithRoles } from '~/types/user'

export const useTeam = () => {
    const { $api } = useNuxtApp()

    const getTeam = async (id: number) => {
        try {
            const { data: team } = await $api.get<Team>(`/team/${id}`)
            return team
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get team details')
        }
    }

    const createTeam = async (data: { name: string, description: string }) => {
        try {
            const { data: team } = await $api.post<Team>(`/team/create`, data)
            return team
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to create team')
        }
    }

    const updateTeam = async (data: { id: number, name?: string, description?: string }) => {
        try {
            const { data: team } = await $api.put<Team>(`/team/update`, data)
            return team
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to update team')
        }
    }

    const deleteTeam = async (id: number) => {
        try {
            await $api.delete(`/team/${id}`)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to delete team')
        }
    }

    const getUserTeams = async () => {
        try {
            const { data: teams } = await $api.get<Team[]>(`/team/user`)
            return teams
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get user teams')
        }
    }

    const getTeamUsers = async (teamId: number) => {
        try {
            const { data: users } = await $api.get<UserWithRoles[]>(`/team/${teamId}/users`)
            return users
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get team users')
        }
    }

    const addTeamUser = async (teamId: number, email: string) => {
        try {
            await $api.post(`/team/${teamId}/add-user`, { email })
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to add team user')
        }
    }

    const removeTeamUser = async (teamId: number, userId: number) => {
        try {
            await $api.delete(`/team/${teamId}/remove-user/${userId}`)
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to remove team user')
        }
    }

    const getRoles = async () => {
        try {
            const { data: roles } = await $api.get<Role[]>('/team/roles')
            return roles
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get roles')
        }
    }

    const setRole = async (teamId: number, userId: number, roleId: number) => {
        try {
            const { data: roles } = await $api.post(`/team/${teamId}/role`, {
                user_id: userId,
                role_id: roleId,
            })
            return roles
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get roles')
        }
    }

    const getProjectsStatistic = async (team_id: number): Promise<TeamProject[]> => {
        try {
            const { data: statistic } = await $api.get(`/team/${team_id}/statistic`)
            return statistic
        }
        catch (error) {
            const axiosError = error as AxiosError<{ message: string }>
            throw new Error(axiosError.response?.data?.message || 'Failed to get team statistic')
        }
    }

    return {
        getTeam,
        createTeam,
        updateTeam,
        deleteTeam,
        getUserTeams,
        getTeamUsers,
        addTeamUser,
        removeTeamUser,
        getRoles,
        setRole,
        getProjectsStatistic,
    }
}
