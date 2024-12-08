export type Team = {
    id: number
    name: string
    description: string
}

export type TeamProject = {
    id: number
    code: string
    total_tasks_count: number
    completed_tasks_count: number
    avg_task_duration: number
    cluster: number
}
