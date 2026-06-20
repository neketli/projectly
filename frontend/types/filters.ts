export type BoardFilters = {
    search: string
    assignedUserId?: number
    priority?: number
    deadlineFrom?: number
    deadlineTo?: number
    createdAtFrom?: number
    createdAtTo?: number
    storyPointsFrom?: number
    storyPointsTo?: number
    hideFinishStatus?: boolean
}

export const defaultBoardFilters: BoardFilters = {
    search: '',
}
