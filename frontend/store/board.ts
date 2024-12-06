import { defineStore } from 'pinia'
import type { Board, Status } from '~/types/board'
import type { Task } from '~/types/task'

export const useBoardStore = defineStore('task-tracker-board', {
    state: () => ({
        board: {} as Board,
        statusList: [] as Status[],
        tasks: [] as { [status_id: number]: Task[] },
    }),

    getters: {
        sortedStatusList: (state) => {
            return state.statusList.sort((a, b) => a.order - b.order)
        },
        statusCount: (state) => {
            return state.statusList.length
        },
    },

    actions: {
        changeTaskStatus(oldStatusId: number, newStatusId: number, taskId: number) {
            this.tasks[newStatusId].push(this.tasks[oldStatusId].find(t => t.id === taskId) as Task)
            this.tasks[oldStatusId] = this.tasks[oldStatusId].filter(t => t.id !== taskId)
        },

        setStatusList(statusList: Status[]) {
            this.statusList = statusList
            this.tasks = statusList.reduce((acc, status) => ({
                ...acc,
                [status.id]: this.tasks[status.id] || [],
            }), {})
        },
        setTaskList(taskMap: { [status_id: number]: Task[] }) {
            Object.entries(taskMap).forEach(([status_id, tasks]) => {
                this.tasks[Number(status_id)] = tasks
            })
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
