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

export const defaultStatusColors = [
    '#007BFF', //  Blue
    '#FD7E14', //  Orange
    '#28A745', //  Green
    '#DC3545', //  Red
    '#6F42C1', //  Purple
    '#FFC107', //  Yellow
]
