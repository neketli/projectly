import { defineStore } from 'pinia'
import { UserRole, type UserWithRoles } from '~/types/user'
import type { Team } from '~/types/team'
import type { Project } from '~/types/project'

export const useTeamStore = defineStore('projectly-team', {
    persist: {
        storage: persistedState.localStorage,
    },
    state: () => ({
        team: {} as Team,
        users: [] as UserWithRoles[],
        project: {} as Project,
    }),
    getters: {
        getUserRole(state): UserRole {
            const authStore = useAuthStore()
            const user = state.users.find(u => u.id === authStore.user.id)
            return user?.role.role_name || UserRole.UNKNOWN
        },
        isDeleteAvailable(): boolean {
            const access = [UserRole.OWNER]
            return access.includes(this.getUserRole)
        },
        isUsersActionsAvailable(): boolean {
            const access = [UserRole.OWNER]
            return access.includes(this.getUserRole)
        },
        isEditAvailable(): boolean {
            const access = [UserRole.OWNER, UserRole.EDITOR]
            return access.includes(this.getUserRole)
        },
        isInviteAvailable(): boolean {
            const access = [UserRole.OWNER, UserRole.EDITOR]
            return access.includes(this.getUserRole)
        },
        canLeave(state): boolean {
            if (this.getUserRole !== UserRole.OWNER) return true
            return state.users.filter(user => user.role.role_name === UserRole.OWNER).length > 1
        },
    },
})
