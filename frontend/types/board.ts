export type Board = {
    id: number
    title: string
    project_id: number
}

export type Status = {
    id: number
    title: string
    board_id: number
    order: number
    hex_color: string
}
