<template>
    <div class="flex items-start gap-2 p-2 border border-gray-200 dark:border-gray-700 bg-gray-100 dark:bg-gray-800 bg-opacity-80 rounded-lg">
        <UserAvatar
            :size="32"
            :file-name="activity.user?.avatar"
            class="mr-2"
        />
        <div class="flex flex-col gap-1">
            <div class="flex items-center gap-2 text-sm text-gray-700 dark:text-gray-300">
                <span class="font-bold">
                    {{ activity.user?.name }} {{ activity.user?.surname }}
                </span>
                <span class="text-gray-500 font-light">
                    {{ dayjs.unix(activity.created_at).format('DD MMM YYYY, HH:mm') }}
                </span>
            </div>
            <div class="text-sm text-gray-600 dark:text-gray-400">
                {{ actionText }}
            </div>
            <div
                v-if="activity.old_value || activity.new_value"
                class="flex items-center gap-1 text-xs"
            >
                <span
                    v-if="activity.old_value"
                    class="px-1.5 py-0.5 rounded bg-red-100 dark:bg-red-900/40 text-red-700 dark:text-red-300 line-through"
                >
                    {{ displayOldValue }}
                </span>
                <span
                    v-if="activity.old_value && activity.new_value"
                    class="text-gray-400"
                >→</span>
                <span
                    v-if="activity.new_value"
                    class="px-1.5 py-0.5 rounded bg-green-100 dark:bg-green-900/40 text-green-700 dark:text-green-300"
                >
                    {{ displayNewValue }}
                </span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { computed } from 'vue'
import type { TaskActivity } from '~/types/task'

const props = defineProps<{
    activity: TaskActivity
}>()

const { t } = useI18n()

const actionKeys: Record<string, string> = {
    task_created: 'task.activity.action.task_created',
    status_changed: 'task.activity.action.status_changed',
    comment_added: 'task.activity.action.comment_added',
    comment_deleted: 'task.activity.action.comment_deleted',
    attachment_added: 'task.activity.action.attachment_added',
    attachment_deleted: 'task.activity.action.attachment_deleted',
}

const actionText = computed(() => {
    const a = props.activity
    if (a.action_type === 'task_updated') {
        const fieldKey = `task.activity.field.${a.field_name}`
        const field = t(fieldKey)
        return t('task.activity.action.task_updated', { field: field !== fieldKey ? field : a.field_name })
    }
    const key = actionKeys[a.action_type]
    return key ? t(key) : a.action_type
})

const boardStore = useBoardStore()
const teamStore = useTeamStore()

const displayOldValue = computed(() => resolveValue(props.activity.old_value))
const displayNewValue = computed(() => resolveValue(props.activity.new_value))

function resolveValue(value: string): string {
    if (!value) return ''
    const a = props.activity

    if (a.action_type === 'status_changed' || (a.action_type === 'task_updated' && a.field_name === 'status_id')) {
        const status = boardStore.statusList.find(s => s.id === Number(value))
        if (status) return status.title
    }

    if (a.action_type === 'task_updated' && a.field_name === 'assigned_user_id') {
        if (Number(value) === 0) return t('task.activity.user_none')
        const user = teamStore.users.find(u => u.id === Number(value))
        if (user) return `${user.name} ${user.surname}`
    }

    return value
}
</script>
