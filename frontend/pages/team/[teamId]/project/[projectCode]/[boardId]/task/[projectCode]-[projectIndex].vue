<template>
    <div class="task">
        <ElPageHeader @back="navigateTo(`/team/${teamId}/project/${projectCode}/${boardId}`)">
            <template #content>
                <h1 class="text-2xl">
                    {{ task.project_code }}-{{ task.project_index }}
                </h1>
            </template>

            <template #title>
                <h3>
                    {{ $t('board.title') }}
                </h3>
            </template>

            <template #extra>
                <ElPopconfirm
                    v-if="projectStore.isDeleteAvailable"
                    width="250"
                    :title="$t('task.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteTask"
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
                    v-if="projectStore.isEditAvailable"
                    :title="$t('task.update.title')"
                    circle
                    plain
                    @click="dialog.task = true"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>
            </template>
        </ElPageHeader>

        <h3 class="text-4xl mt-2">
            {{ task.title }}
        </h3>

        <TaskForm
            :task="task"
            :status-id="task.status_id"
            @success="handleTaskUpdated"
            @cancel="dialog.task = false"
        />

        <ElDialog
            v-model="dialog.task"
            :title="$t('task.update.title')"
            align-center
            destroy-on-close
        >
            <TaskForm
                :task="task"
                :status-id="task.status_id"
                @success="handleTaskUpdated"
                @cancel="dialog.task = false"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { DetailedTask } from '~/types/task'

const { teamId, projectCode, boardId, projectIndex } = useRoute().params
const { t } = useI18n()

useHead({
    title: t('task.title'),
})

definePageMeta({
    layout: 'team',
})

const projectStore = useProjectStore()
const boardStore = useBoardStore()

const { getTaskDetail, deleteTask } = useTask()
const { getStatusList } = useStatus()

const dialog = reactive({
    task: false,
})

const task = ref({} as DetailedTask)

const isLoading = ref(false)

const handleTaskUpdated = (updated: DetailedTask) => {
    task.value = updated
    dialog.task = false
}

const handleDeleteTask = async () => {
    try {
        await deleteTask(task.value.id)
        ElMessage.success(t('task.success.delete'))
        navigateTo(`/team/${teamId}/project/${projectCode}/${boardId}`)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

onMounted(async () => {
    try {
        isLoading.value = true
        task.value = await getTaskDetail(`${projectCode}`, Number(projectIndex))
        const statuses = await getStatusList(Number(boardId))
        boardStore.setStatusList(statuses)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
})
</script>
