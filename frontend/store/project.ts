import { defineStore } from 'pinia'
import type { Board } from '~/types/board'
import type { Project } from '~/types/project'
import { UserRole } from '~/types/user'

export const useProjectStore = defineStore('projectly-project', {
    state: () => ({
        project: {} as Project,
        boardList: [] as Board[],

        boardsDrawer: false,
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
