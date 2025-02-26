<template>
    <div class="relative bg-white dark:bg-slate-800 dark:shadow-none rounded-lg shadow-md py-2 px-4">
        <a
            class="hover:underline hover:text-blue-500 hover:cursor-pointer transition-all text-lg"
            @click="handleTaskClick"
        >
            {{ project.code }}-{{ task.project_index }}
        </a>

        <h4 class="mt-1 text-xl font-bold max-w-48 line-clamp-3 text-ellipsis">
            {{ task.title }}
        </h4>

        <div class="absolute top-2 right-2">
            <ElPopconfirm
                width="250"
                :title="$t('task.delete.confirm')"
                :confirm-button-text="$t('common.button.confirm')"
                :cancel-button-text="$t('common.button.cancel')"
                confirm-button-type="danger"
                @confirm="emit('delete', task)"
            >
                <template #reference>
                    <ElButton
                        :title="$t('task.delete.title')"
                        circle
                        plain
                        type="danger"
                    >
                        <Icon name="mdi:delete" />
                    </ElButton>
                </template>
            </ElPopconfirm>

            <ElButton
                :title="$t('task.update.title')"
                circle
                plain
                @click="dialog.task = true"
            >
                <Icon name="mdi:pencil" />
            </ElButton>
        </div>

        <div
            v-if="task.assigned_user?.avatar"
            class="absolute bottom-2 right-2"
        >
            <UserAvatar
                :size="20"
                :file-name="task.assigned_user?.avatar"
                :title="task.assigned_user?.email"
            />
        </div>

        <div class="mt-2 flex flex-wrap gap-2">
            <ElTag v-if="task.story_points">
                {{ task.story_points }}
            </ElTag>

            <ElTag
                v-if="task.priority"
                type="warning"
            >
                {{ priority }}
            </ElTag>

            <ElTag
                v-if="task.deadline"
                type="success"
            >
                {{ dayjs.unix(task.deadline).format('ddd DD, MMM YYYY') }}
            </ElTag>
        </div>

        <ElDialog
            v-model="dialog.task"
            :title="$t('task.update.title')"
            align-center
            destroy-on-close
        >
            <TaskForm
                :task="task"
                :status-id="task.status_id"
                @cancel="dialog.task = false"
                @success="handleUpdateTask"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import dayjs from 'dayjs'
import type { DetailedTask } from '~/types/task'

const props = defineProps<{
    task: DetailedTask
}>()
const emit = defineEmits <{
    (event: 'update', task: DetailedTask): void
    (event: 'delete', task: DetailedTask): void
}>()

const route = useRoute()
const { project } = toRefs(useProjectStore())

const dialog = reactive({
    task: false,
})

const priority = computed(() => '⚡'.repeat(props.task.priority || 0))

const handleUpdateTask = (task: DetailedTask) => {
    emit('update', task)
    dialog.task = false
}

const handleTaskClick = () => {
    navigateTo(`${route.path}/task/${project.value.code}-${props.task.project_index}`)
}
</script>
