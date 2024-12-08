<template>
    <div
        class="bg-neutral-100 dark:bg-neutral-900 dark:ring-2 dark:ring-neutral-700 min-w-[20vw] rounded-md p-2 relative overflow-hidden min-h-full"
    >
        <div class="flex justify-between items-center gap-4 p-2 sticky z-10 bg-neutral-100 dark:bg-neutral-900 top-0">
            <template v-if="!isEdit">
                <div class="flex items-center gap-4">
                    <StatusTag :color="status.hex_color">
                        {{ status.title }}
                    </StatusTag>

                    <p class="text-gray-600 dark:text-gray-500 text-sm">
                        {{ taskList?.length }}
                    </p>
                </div>

                <div class="flex gap-2 items-center text-gray-300">
                    <ElDropdown trigger="click">
                        <ElButton circle>
                            <Icon name="mdi:dots-horizontal" />
                        </ElButton>

                        <template #dropdown>
                            <ElDropdownItem @click="handleEditStatus">
                                <Icon
                                    name="mdi:pencil"
                                    class="mr-2"
                                />
                                {{ $t('common.button.edit') }}
                            </ElDropdownItem>

                            <ElDropdownItem @click="dialog.delete = true">
                                <Icon
                                    name="mdi:delete"
                                    class="mr-2"
                                />
                                {{ $t('common.button.delete') }}
                            </ElDropdownItem>

                            <ElDropdownItem
                                :disabled="status.order === statusCount - 1"
                                @click="handleMove('right')"
                            >
                                <Icon
                                    name="mdi:arrow-right"
                                    class="mr-2"
                                />
                                {{ $t('status.form.right') }}
                            </ElDropdownItem>

                            <ElDropdownItem
                                :disabled="status.order === 0"
                                @click="handleMove('left')"
                            >
                                <Icon
                                    name="mdi:arrow-left"
                                    class="mr-2"
                                />
                                {{ $t('status.form.left') }}
                            </ElDropdownItem>

                            <ElDropdownItem @click="handleChangeSort">
                                <Icon
                                    :name="sortDirection === 'asc' ? 'mdi:sort-ascending' : 'mdi:sort-descending'"
                                    class="mr-2"
                                />
                                {{ $t('status.form.sort') }} -
                                {{ sortDirection === 'asc' ? $t('status.sort.asc') : $t('status.sort.desc') }}
                            </ElDropdownItem>

                            <ElSelectV2
                                v-model="statusSort"
                                class="min-w-40"
                                :label="$t('status.form.sort')"
                                :options="statusSortOptions"
                            />
                        </template>
                    </ElDropdown>

                    <ElButton
                        :title="$t('status.create.task')"
                        type="primary"
                        circle
                        @click="dialog.task = true"
                    >
                        <Icon name="mdi:plus" />
                    </ElButton>
                </div>
            </template>
            <template v-else>
                <div class="w-full flex flex-col gap-2">
                    <label>
                        {{ $t('status.form.title') }}
                    </label>
                    <div class="w-full flex gap-2">
                        <ElInput
                            v-model="title"
                            @keyup.enter="handleSaveStatus"
                        />
                        <ElColorPicker
                            v-model="color"
                            :predefine="defaultStatusColors"
                        />
                    </div>
                </div>

                <div class="flex self-end">
                    <ElButton
                        circle
                        @click="handleCancel"
                    >
                        <Icon name="mdi:close" />
                    </ElButton>
                    <ElButton
                        circle
                        type="success"
                        :disabled="!title || !color"
                        @click="handleSaveStatus"
                    >
                        <Icon name="mdi:content-save" />
                    </ElButton>
                </div>
            </template>
        </div>

        <ElScrollbar
            height="58vh"
            class="w-full px-2 py-2"
        >
            <Draggable
                :id="`status-${status.id}`"
                v-model="taskList"
                class="flex flex-col gap-4 w-full min-h-[58vh] overflow-hidden"
                group="tasks"
                :sort="false"
                :animation="200"
                item-key="id"
                @end="handleChangeTaskStatus"
            >
                <template #item="{ element }">
                    <TaskCard
                        :id="`task-${element.id}`"
                        :key="element.id"
                        :task="element"
                        class="cursor-grab"
                        @update="handleUpdateTask"
                        @delete="handleDeleteTask"
                    />
                </template>
            </Draggable>
        </ElScrollbar>

        <ElDialog
            v-model="dialog.delete"
            :title="$t('status.delete.confirm')"
            align-center
            destroy-on-close
            class="!w-1/3"
        >
            <ElButton
                @click="dialog.delete = false"
            >
                {{ $t('common.button.cancel') }}
            </ElButton>
            <ElButton
                type="danger"
                @click="emit('delete', status)"
            >
                {{ $t('common.button.confirm') }}
            </ElButton>
        </ElDialog>

        <ElDialog
            v-model="dialog.task"
            :title="$t('task.create.title')"
            align-center
            destroy-on-close
        >
            <TaskForm
                :status-id="status.id"
                @cancel="dialog.task = false"
                @success="handleCreateTask"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import dayjs from 'dayjs'
import Draggable from 'vuedraggable'
import { defaultStatusColors, type Status } from '~/types/board'
import type { Task } from '~/types/task'

const props = defineProps<{ status: Status }>()
const emit = defineEmits(['create', 'update', 'delete'])

const { t } = useI18n()

const boardStore = useBoardStore()
const { statusCount, finishStatus, tasks } = toRefs(boardStore)
const { updateTaskStatus, deleteTask } = useTask()

const isEdit = ref(false)
const isLoading = ref(false)
const dialog = reactive({
    delete: false,
    task: false,
})
const title = ref('')
const color = ref(props.status.hex_color || defaultStatusColors[0])

const sortDirection = ref<'asc' | 'desc'>('asc')
const statusSort: Ref<keyof Task> = ref('created_at')
const statusSortOptions = [
    {
        label: t('status.sort.updated_at'),
        value: 'updated_at',
    },
    {
        label: t('status.sort.created_at'),
        value: 'created_at',
    },
    {
        label: t('status.sort.deadline'),
        value: 'deadline',
    },
    {
        label: t('status.sort.priority'),
        value: 'priority',
    },
    {
        label: t('status.sort.story_point'),
        value: 'story_points',
    },
]

const taskList = computed(() => {
    const arr = tasks.value[props.status.id]
    if (!arr) return []
    arr.sort((a: Task, b: Task) => {
        if (a[statusSort.value] === undefined && b[statusSort.value] !== undefined) return 1
        if (a[statusSort.value] !== undefined && b[statusSort.value] === undefined) return -1

        if (sortDirection.value === 'asc') {
            // @ts-expect-error
            return a[statusSort.value] > b[statusSort.value] ? 1 : -1
        }
        // @ts-expect-error
        return a[statusSort.value] < b[statusSort.value] ? 1 : -1
    })
    return arr
})

const handleEditStatus = () => {
    isEdit.value = true
    title.value = props.status.title || ''
}

const handleCancel = () => {
    isEdit.value = false

    if (props.status.id === 0) {
        emit('delete', props.status)
    }
}

const handleChangeSort = () => {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc'
}

const handleSaveStatus = () => {
    isEdit.value = false

    emit(props.status.id === 0 ? 'create' : 'update', {
        ...props.status,
        title: title.value,
        hex_color: color.value,
    })

    title.value = ''
}

const handleMove = (direction: 'left' | 'right') => {
    if ((direction === 'left' && props.status.order === 0)
      || (direction === 'right' && props.status.order === statusCount.value - 1)) return

    emit('update', {
        ...props.status,
        order: direction === 'left'
            ? props.status.order - 1
            : props.status.order + 1,
    })
}

const handleCreateTask = (task: Task) => {
    tasks.value[task.status_id].push(task)
    dialog.task = false
    ElMessage.success(t('task.success.create'))
}

const handleDeleteTask = async (task: Task) => {
    try {
        if (task.id) {
            isLoading.value = true
            await deleteTask(task.id)
            tasks.value[task.status_id] = tasks.value[task.status_id].filter(item => item.id !== task.id)
            ElMessage.success(t('task.success.delete'))
        }
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

const handleUpdateTask = (task: Task) => {
    tasks.value[task.status_id] = tasks.value[task.status_id].map(item => item.id === task.id ? task : item)
    ElMessage.success(t('task.success.update'))
}
const handleChangeTaskStatus = async (params: {
    item: { id: string }
    from: { id: string }
    to: { id: string }
}) => {
    try {
        const taskId = Number(params.item.id.split('-')[1])
        const fromStatusId = Number(params.from.id.split('-')[1])
        const toStatusId = Number(params.to.id.split('-')[1])
        if (fromStatusId === toStatusId) return

        isLoading.value = true
        const finishedAt = finishStatus.value.id === toStatusId ? dayjs().unix() : null
        await updateTaskStatus(taskId, toStatusId, finishedAt)
        boardStore.changeTaskStatus(fromStatusId, toStatusId, taskId, finishedAt)

        ElMessage.success(t('task.success.update'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

onMounted(() => {
    if (props.status.id === 0) {
        isEdit.value = true
        title.value = props.status.title || ''
    }
})
</script>
