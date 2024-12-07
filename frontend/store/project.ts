import { defineStore } from 'pinia'
import type { Project } from '~/types/project'
import { UserRole } from '~/types/user'

export const useProjectStore = defineStore('projectly-project', {
    state: () => ({
        project: {} as Project,
    }),
    getters: {
        isDeleteAvailable(): boolean {
            const teamStore = useTeamStore()
            const access = [UserRole.OWNER]
            return access.includes(teamStore.getUserRole)
        },
        isEditAvailable(): boolean {
            const teamStore = useTeamStore()
            const access = [UserRole.OWNER, UserRole.EDITOR]
            return access.includes(teamStore.getUserRole)
        },
    },
})
