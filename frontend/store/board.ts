import { defineStore } from 'pinia'
import type { Board, Status } from '~/types/board'
import type { DetailedTask } from '~/types/task'

export const useBoardStore = defineStore('projectly-board', {
    state: () => ({
        board: {} as Board,
        statusList: [] as Status[],
        tasks: [] as { [status_id: number]: DetailedTask[] },
    }),

    getters: {
        sortedStatusList: (state) => {
            return state.statusList.sort((a, b) => a.order - b.order)
        },
        statusCount: (state) => {
            return state.statusList.length
        },
        finishStatus: (state) => {
            return state.statusList.slice(-1)[0]
        },
    },

    actions: {
        changeTaskStatus(oldStatusId: number, newStatusId: number, taskId: number, finishedAt: number | null) {
            const task = {
                ...this.tasks[oldStatusId].find(t => t.id === taskId) as DetailedTask,
                finished_at: finishedAt || 0,
            }
            this.tasks[newStatusId].push(task)
            this.tasks[oldStatusId] = this.tasks[oldStatusId].filter(t => t.id !== taskId)
        },
        updateTasksStatusMap() {
            this.tasks = this.statusList.reduce((acc, status) => ({
                ...acc,
                [status.id]: this.tasks[status.id] || [],
            }), {})
        },
        setStatusList(statusList: Status[]) {
            this.statusList = statusList
            this.updateTasksStatusMap()
        },
        setTaskList(taskMap: { [status_id: number]: DetailedTask[] }) {
            Object.entries(taskMap).forEach(([status_id, tasks]) => {
                this.tasks[Number(status_id)] = tasks
            })
        },
        replaceStatus(status: Status) {
            this.statusList.splice(status.order, 1, status)
            this.updateTasksStatusMap()
        },
        deleteStatus(status: Status) {
            this.statusList.splice(this.statusList.indexOf(status), 1)
            this.statusList = this.statusList.map(s => (s.order > status.order ? { ...s, order: s.order - 1 } : s))
            this.updateTasksStatusMap()
        },
    },
})
