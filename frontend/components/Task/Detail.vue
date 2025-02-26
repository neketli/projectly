<template>
    <div class="task-detail">
        <ElDescriptions
            :column="2"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.created_user') }}
                    </div>
                </template>

                <div
                    v-if="task.created_user"
                    class="flex gap-2 items-center"
                >
                    <UserAvatar
                        v-if="task.created_user?.avatar"
                        :size="32"
                        :file-name="task.created_user.avatar"
                        class="sm:ml-2"
                    />

                    <span class="hidden sm:inline">
                        {{ task.created_user.name }}
                        {{ task.created_user.surname }}
                        ({{ task.created_user.email }})
                    </span>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.assigned_user') }}
                    </div>
                </template>

                <div
                    v-if="task.assigned_user?.email"
                    class="flex gap-2 items-center"
                >
                    <UserAvatar
                        v-if="task.assigned_user.avatar"
                        :size="32"
                        :file-name="task.assigned_user.avatar"
                        class="sm:ml-2"
                    />

                    <span class="hidden sm:inline">
                        {{ task.assigned_user.name }}
                        {{ task.assigned_user.surname }}
                        ({{ task.assigned_user.email }})
                    </span>
                </div>
                <ElTag
                    v-else
                    type="info"
                >
                    {{ t('task.form.placeholder.user') }}
                </ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="3"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.status') }}
                    </div>
                </template>

                <ElSelect
                    :value="task?.status?.id"
                    multiple
                    placeholder=""
                    @change="handleStatusChange"
                >
                    <ElOption
                        v-for="item in boardStore.statusList"
                        :key="item.id"
                        :value="item.id"
                    >
                        <StatusTag
                            :color="item.hex_color"
                            class="!py-0"
                        >
                            {{ item.title }}
                        </StatusTag>
                    </ElOption>
                    <template #tag>
                        <StatusTag
                            :color="task.status?.hex_color"
                            class="!p-1 mx-auto"
                        >
                            {{ task.status?.title }}
                        </StatusTag>
                    </template>
                </ElSelect>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.story_point') }}
                    </div>
                </template>

                <ElTag>
                    {{ task.story_points || t('task.form.placeholder.time') }}
                </ElTag>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.tracked_time') }}
                    </div>
                </template>

                <ElTag>
                    {{ task.tracked_time || t('task.form.placeholder.time') }}
                </ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="2"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.deadline') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="warning">
                        {{ task.deadline || t('task.form.placeholder.deadline') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.priority') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    {{ priorityOptions[task.priority] }}
                </div>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="3"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.created_at') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="info">
                        {{ dayjs.unix(task.created_at).format('DD.MM.YYYY HH:mm') || t('task.form.placeholder.time') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.updated_at') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="info">
                        {{ dayjs.unix(task.updated_at).format('DD.MM.YYYY HH:mm') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="cell-item">
                        <Icon name="mdi:account" />
                        {{ t('task.form.finished_at') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="success">
                        {{ task.finished_at
                            ? dayjs.unix(task.finished_at).format('DD.MM.YYYY HH:mm')
                            : t('task.form.placeholder.time') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>
        </ElDescriptions>
    </div>
</template>

<script lang="ts" setup>
import dayjs from 'dayjs'
import type { DetailedTask } from '~/types/task'

defineProps<{ task: DetailedTask }>()
const emit = defineEmits<{
    (event: 'update-status', taskId: number): void
}>()

const { t } = useI18n()

const boardStore = useBoardStore()

const priorityOptions = computed(() => {
    return [
        t('task.form.placeholder.priority'),
        '⚡',
        '⚡⚡',
        '⚡⚡⚡',
    ]
})

const handleStatusChange = ([taskId]: number[]) => {
    emit('update-status', taskId)
}
</script>
