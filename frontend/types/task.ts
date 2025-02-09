export type Task = {
    id: number
    project_index: number
    title: string
    description?: string
    priority: number
    story_points?: number
    deadline?: number
    tracked_time?: number
    status_id: number
    assigned_user_id?: number | string
    created_user_id: number
    updated_at: number
    created_at: number
    finished_at?: number
}

export type DetailedTask = Task & {
    project_code: string
    status: {
        id: number
        title: string
        hex_color: string
    }
    created_user: {
        id: number
        name: string
        surname: string
        email: string
        avatar: string
    }
    assigned_user: {
        id: number
        name: string
        surname: string
        email: string
        avatar: string
    }
    meta: {
        team_id: number
        project_id: number
        board_id: number
    }
}
