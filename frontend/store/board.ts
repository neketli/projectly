import { defineStore } from 'pinia'
import type { Board, Status } from '~/types/board'

export const useBoardStore = defineStore('task-tracker-board', {
    state: () => ({
        board: {} as Board,
        statusList: [] as Status[],
    }),

    getters: {
        sortedStatusList: (state) => {
            return state.statusList.sort((a, b) => a.order - b.order)
        },
    },

    actions: {

        updateStatus(status: Status) {
            this.statusList = this.statusList.map(s => (s.id === status.id ? status : s))
        },
        replaceStatus(status: Status) {
            this.statusList.splice(status.order, 1, status)
        },
        deleteStatus(status: Status) {
            this.statusList.splice(this.statusList.indexOf(status), 1)
            this.statusList = this.statusList.map(s => (s.order > status.order ? { ...s, order: s.order - 1 } : s))
        },
    },
})
