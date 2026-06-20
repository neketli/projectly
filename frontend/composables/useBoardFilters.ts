import type { DetailedTask } from '~/types/task'

export const useBoardFilters = () => {
    const boardStore = useBoardStore()

    const matchesFilters = (task: DetailedTask): boolean => {
        const f = boardStore.filters

        if (f.search) {
            const q = f.search.toLowerCase()
            if (!task.title.toLowerCase().includes(q)
              && !(task.description?.toLowerCase().includes(q) ?? false)) {
                return false
            }
        }

        if (f.assignedUserId && task.assigned_user?.id !== f.assignedUserId) {
            return false
        }

        if (f.priority && task.priority !== f.priority) {
            return false
        }

        if (f.deadlineFrom && task.deadline && task.deadline < f.deadlineFrom) {
            return false
        }

        if (f.deadlineTo && task.deadline && task.deadline > f.deadlineTo) {
            return false
        }

        if (f.deadlineTo === 0 && task.deadline && task.deadline > 0) {
            return false
        }

        if (f.createdAtFrom && task.created_at && task.created_at < f.createdAtFrom) {
            return false
        }

        if (f.createdAtTo && task.created_at && task.created_at > f.createdAtTo) {
            return false
        }

        if (f.hideFinishStatus && boardStore.finishStatus && task.status_id === boardStore.finishStatus.id) {
            return false
        }

        if (f.storyPointsFrom !== undefined && (task.story_points ?? 0) < f.storyPointsFrom) {
            return false
        }

        if (f.storyPointsTo !== undefined && (task.story_points ?? 0) > f.storyPointsTo) {
            return false
        }

        return true
    }

    return { matchesFilters }
}
