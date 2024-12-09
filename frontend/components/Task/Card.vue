<template>
    <div class="relative bg-white dark:bg-slate-800 dark:shadow-none rounded-lg shadow-md py-2 px-4">
        <span class="text-lg">
            {{ project.code }}-{{ task.project_index }}
        </span>

        <h3 class="mt-1 text-xl font-bold">
            {{ task.title }}
        </h3>

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
import type { Task } from '~/types/task'

const props = defineProps<{ task: Task }>()
const emit = defineEmits(['delete', 'update'])

const { project } = toRefs(useProjectStore())

const dialog = reactive({
    task: false,
})

const priority = computed(() => '⚡'.repeat(props.task.priority || 0))

const handleUpdateTask = (task: Task) => {
    emit('update', task)
    dialog.task = false
}
</script>
