<template>
    <div class="task-detail">
        <ElDescriptions
            :column="isMobile ? 1 : 2"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:account" />

                        <span>
                            {{ t('task.form.created_user') }}
                        </span>
                    </div>
                </template>

                <a
                    v-if="task.created_user"
                    class="flex gap-2 items-center cursor-pointer hover:text-blue-500 transition-all"
                    @click="handleUserProfile(task.created_user.email)"
                >
                    <UserAvatar
                        v-if="task.created_user?.avatar"
                        :size="32"
                        :file-name="task.created_user.avatar"
                        class="ml-2"
                    />

                    <span class="inline">
                        {{ task.created_user.name }}
                        {{ task.created_user.surname }}
                        ({{ task.created_user.email }})
                    </span>
                </a>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:account-check" />
                        {{ t('task.form.assigned_user') }}
                    </div>
                </template>

                <a
                    v-if="task.assigned_user?.email"
                    class="flex gap-2 items-center cursor-pointer hover:text-blue-500 transition-all"
                    @click="handleUserProfile(task.assigned_user.email)"
                >
                    <UserAvatar
                        v-if="task.assigned_user.avatar"
                        :size="32"
                        :file-name="task.assigned_user.avatar"
                        class="ml-2"
                    />

                    <span class="inline">
                        {{ task.assigned_user.name }}
                        {{ task.assigned_user.surname }}
                        ({{ task.assigned_user.email }})
                    </span>
                </a>
                <ElTag
                    v-else
                    type="info"
                >
                    {{ t('task.form.placeholder.user') }}
                </ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="isMobile ? 1 : 3"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:list-status" />
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
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:timer" />
                        {{ t('task.form.story_point') }}
                    </div>
                </template>

                <ElTag>
                    {{ task.story_points || t('task.form.placeholder.time') }}
                </ElTag>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:timelapse" />
                        {{ t('task.form.tracked_time') }}
                    </div>
                </template>

                <ElTag>
                    {{ task.tracked_time || t('task.form.placeholder.time') }}
                </ElTag>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="isMobile ? 1 : 2"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:calendar-clock" />
                        {{ t('task.form.deadline') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="warning">
                        {{ task.deadline
                            ? dayjs.unix(task.deadline).format('DD.MM.YYYY HH:mm')
                            : t('task.form.placeholder.time') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:label-variant" />
                        {{ t('task.form.priority') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    {{ priorityOptions[task.priority] }}
                </div>
            </ElDescriptionsItem>
        </ElDescriptions>

        <ElDescriptions
            :column="isMobile ? 1 : 3"
            border
        >
            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:calendar" />
                        {{ t('task.form.created_at') }}
                    </div>
                </template>

                <div class="flex gap-2 items-center">
                    <ElTag type="info">
                        {{ dayjs.unix(task.created_at).format('DD.MM.YYYY HH:mm') }}
                    </ElTag>
                </div>
            </ElDescriptionsItem>

            <ElDescriptionsItem>
                <template #label>
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:calendar-edit" />
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
                    <div class="flex gap-2 items-center">
                        <Icon name="mdi:calendar-check" />
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
const { width } = useWindowSize()

const boardStore = useBoardStore()

const isMobile = computed(() => {
    return width.value < 768
})

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

const handleUserProfile = (email: string) => navigateTo(`/user/${email}`)
</script>
