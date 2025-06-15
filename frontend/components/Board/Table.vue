<template>
    <div>
        <ElInput
            v-model.trim="search"
            :placeholder="$t('common.search')"
            class="md:!w-1/3"
        />

        <ElTable
            :data="filteredTasks"
            stripe
            class="w-full"
            height="100%"
            :empty-text="$t('common.no_data')"
            @row-click="handleTaskDetail"
        >
            <ElTableColumn
                prop="title"
                :label="t('task.form.title')"
                min-width="150"
            />

            <ElTableColumn
                prop="status"
                :label="t('task.form.status')"
                sortable
                :filters="statusFilters"
                :filter-method="(value, task) => value == task.status?.id"
                width="200"
            >
                <template #default="{ row: task }">
                    <StatusTag
                        :color="task.status?.hex_color"
                        class="!p-1"
                    >
                        {{ task.status?.title }}
                    </StatusTag>
                </template>
            </ElTableColumn>

            <ElTableColumn
                v-for="col in boardStore.columns"
                :key="col.prop"
                :prop="col.prop"
                :label="col.label"
                :formatter="col.formatter"
                :width="col.width"
                sortable
            >
                <template
                    v-if="col.prop === 'created_user'"
                    #default="{ row: task }"
                >
                    <UserLink :user="task.created_user" />
                </template>
                <template
                    v-else-if="col.prop === 'assigned_user'"
                    #default="{ row: task }"
                >
                    <UserLink :user="task.assigned_user" />
                </template>
            </ElTableColumn>

            <ElTableColumn
                :width="isEditColumns? '300' : '50'"
                align="right"
            >
                <template #header>
                    <div class="flex gap-2 items-center">
                        <ElSelect
                            v-if="isEditColumns"
                            :model-value="boardStore.columns"
                            multiple
                            filterable
                            value-key="prop"
                            :no-data-text="$t('common.no_data')"
                            :no-match-text="$t('common.no_data')"
                            :placeholder="$t('board.table.columns')"
                            :clearable="false"
                            @change="handleChangeColumns"
                        >
                            <ElOption
                                v-for="col in availableColumns"
                                :key="col.prop"
                                :label="col.label"
                                :value="col"
                            />
                        </ElSelect>

                        <ElButton
                            :type="isEditColumns ? 'success' : 'warning'"
                            plain
                            circle
                            @click="isEditColumns = !isEditColumns"
                        >
                            <Icon :name="isEditColumns ? 'mdi:check' : 'mdi:edit'" />
                        </ElButton>
                    </div>
                </template>
            </ElTableColumn>
        </ElTable>
    </div>
</template>

<script lang="ts" setup>
import dayjs from 'dayjs'
import type { Columns } from 'element-plus'
import type { DetailedTask } from '~/types/task'

const { t } = useI18n()
const route = useRoute()

const boardStore = useBoardStore()

const isEditColumns = ref(false)
const search = ref('')

const formatters = {
    story_points: ({ story_points }: DetailedTask) =>
        story_points ? String(story_points) : '-',
    deadline: ({ deadline }: DetailedTask) => deadline ? dayjs.unix(deadline).format('DD.MM.YYYY') : t('task.form.placeholder.time'),
    tracked_time: ({ tracked_time }: DetailedTask) =>
        tracked_time ? String(tracked_time) : t('task.form.placeholder.time'),
    priority: ({ priority }: DetailedTask) =>
        [
            '-',
            '⚡',
            '⚡⚡',
            '⚡⚡⚡',
        ][priority],
    created_at: ({ created_at }: DetailedTask) => dayjs.unix(created_at).format('DD.MM.YYYY HH:mm'),
}

const availableColumns = computed(() => [
    {
        prop: 'story_points',
        label: t('task.table.story_point'),
        formatter: formatters.story_points,
        width: 85,
    },
    {
        prop: 'deadline',
        label: t('task.table.deadline'),
        formatter: formatters.deadline,
        width: 125,
    },
    {
        prop: 'tracked_time',
        label: t('task.table.tracked_time'),
        formatter: formatters.tracked_time,
        width: 105,
    },
    {
        prop: 'priority',
        label: t('task.table.priority'),
        formatter: formatters.priority,
        width: 80,
    },
    {
        prop: 'assigned_user',
        label: t('task.table.assigned_user'),
        width: 0,
    },
    {
        prop: 'created_user',
        label: t('task.table.created_user'),
        width: 0,
    },
    {
        prop: 'created_at',
        label: t('task.table.created_at'),
        formatter: formatters.created_at,
        width: 150,
    },
])

const { tasks } = toRefs(boardStore)

const statusFilters = computed(() => {
    return boardStore.statusList.map(status => ({
        text: status.title,
        value: `${status.id}`,
    }))
})

const tableTasks = computed(() => {
    return Object.values(tasks.value).flat()
})

const filteredTasks = computed(() => {
    return tableTasks.value.filter((task) => {
        return task.title.toLowerCase().includes(search.value.toLowerCase())
    })
})

const handleTaskDetail = (task: DetailedTask) => {
    navigateTo(`${route.path}/task/${task.project_code}-${task.project_index}`)
}

const handleChangeColumns = (columns: Columns<string>) => {
    boardStore.setBoardTableColumns(columns)
}

onMounted(() => {
    if (!boardStore.columns.length) {
        const defaultColumns = ['deadline', 'priority', 'assigned_user']
        boardStore.setBoardTableColumns(availableColumns.value.filter(col => defaultColumns.includes(col.prop)))
    }

    boardStore.setBoardTableColumns(boardStore.columns.map(item => ({
        ...item,
        formatter: formatters[item.prop as keyof typeof formatters],
    })))
})
</script>
