export type User = {
    id: number
    name: string
    surname: string
    email: string
    meta?: {
        avatar?: string
    }
}

export type UserWithRoles = User & {
    role: Role
}

export type Role = {
    id: number
    role_name: UserRole
}

export enum UserRole {
    OWNER = 'owner',
    EDITOR = 'editor',
    DEVELOPER = 'developer',
    USER = 'user',
    UNKNOWN = '',
}

export const RoleLabels: Record<UserRole, string> = {
    [UserRole.OWNER]: 'team.roles.owner',
    [UserRole.EDITOR]: 'team.roles.editor',
    [UserRole.DEVELOPER]: 'team.roles.developer',
    [UserRole.USER]: 'team.roles.user',
    [UserRole.UNKNOWN]: '-',
}
