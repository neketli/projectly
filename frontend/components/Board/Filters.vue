<template>
    <div class="board-filters mb-4">
        <div class="flex flex-wrap items-center gap-1 mb-2">
            <ElButton
                v-for="preset in presets"
                :key="preset.key"
                size="small"
                :type="activePreset === preset.key ? 'primary' : 'default'"
                :plain="activePreset !== preset.key"
                @click="applyPreset(preset.key)"
            >
                <Icon
                    :name="preset.icon"
                    class="mr-1"
                />
                {{ preset.label }}
            </ElButton>

            <ElButton
                v-if="hasActiveFilters"
                size="small"
                circle
                class="!ml-1"
                @click="handleReset"
            >
                <Icon name="mdi:filter-remove" />
            </ElButton>
        </div>

        <ElCollapse v-model="activeCollapse">
            <ElCollapseItem
                :title="$t('board.filters.title') + (hasActiveFilters ? ` (${activeFilterCount})` : '')"
                name="filters"
            >
                <div class="flex flex-wrap items-start gap-3 py-2">
                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.search') }}</span>
                        <ElInput
                            :model-value="filters.search"
                            clearable
                            class="!w-56"
                            size="small"
                            @input="filters.search = $event"
                            @clear="filters.search = ''"
                        />
                    </div>

                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.assignee') }}</span>
                        <ElSelect
                            :model-value="filters.assignedUserId"
                            clearable
                            filterable
                            class="!w-44"
                            size="small"
                            @update:model-value="filters.assignedUserId = $event"
                        >
                            <ElOption
                                v-for="user in availableAssignees"
                                :key="user.id"
                                :label="userLabel(user)"
                                :value="user.id"
                            />
                        </ElSelect>
                    </div>

                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.priority') }}</span>
                        <ElSelect
                            :model-value="filters.priority"
                            clearable
                            class="!w-36"
                            size="small"
                            @update:model-value="filters.priority = $event"
                        >
                            <ElOption
                                v-for="p in priorityOptions"
                                :key="p.value"
                                :label="p.label"
                                :value="p.value"
                            />
                        </ElSelect>
                    </div>

                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.deadline') }}</span>
                        <ElDatePicker
                            :model-value="deadlineRange"
                            type="daterange"
                            range-separator="—"
                            :start-placeholder="$t('board.filters.deadline_from')"
                            :end-placeholder="$t('board.filters.deadline_to')"
                            value-format="x"
                            class="!w-60"
                            size="small"
                            clearable
                            @update:model-value="handleDeadlineChange"
                        />
                    </div>

                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.created_at') }}</span>
                        <ElDatePicker
                            :model-value="createdAtRange"
                            type="daterange"
                            range-separator="—"
                            :start-placeholder="$t('board.filters.created_at_from')"
                            :end-placeholder="$t('board.filters.created_at_to')"
                            value-format="x"
                            class="!w-60"
                            size="small"
                            clearable
                            @update:model-value="handleCreatedAtChange"
                        />
                    </div>

                    <div class="flex flex-col gap-1">
                        <span class="text-xs text-gray-500">{{ $t('board.filters.story_points') }}</span>
                        <div class="flex items-center gap-1">
                            <ElInputNumber
                                :model-value="filters.storyPointsFrom"
                                :placeholder="$t('board.filters.min')"
                                :min="0"
                                :max="999"
                                class="!w-20"
                                size="small"
                                controls-position="right"
                                clearable
                                @update:model-value="filters.storyPointsFrom = $event"
                            />
                            <span class="text-gray-400">—</span>
                            <ElInputNumber
                                :model-value="filters.storyPointsTo"
                                :placeholder="$t('board.filters.max')"
                                :min="0"
                                :max="999"
                                class="!w-20"
                                size="small"
                                controls-position="right"
                                clearable
                                @update:model-value="filters.storyPointsTo = $event"
                            />
                        </div>
                    </div>
                </div>
            </ElCollapseItem>
        </ElCollapse>
    </div>
</template>

<script lang="ts" setup>
import type { PlainUser } from '~/types/user'

const { t } = useI18n()

const authStore = useAuthStore()
const boardStore = useBoardStore()
const { filters, tasks } = toRefs(boardStore)

const activeCollapse = ref<string[]>([])
const deadlineRange = ref<[number, number] | undefined>(undefined)
const createdAtRange = ref<[number, number] | undefined>(undefined)

const priorityOptions = [
    { value: 1, label: '⚡' },
    { value: 2, label: '⚡⚡' },
    { value: 3, label: '⚡⚡⚡' },
]

const userLabel = (user: PlainUser) => `${user.name} ${user.surname}`

const availableAssignees = computed<PlainUser[]>(() => {
    const seen = new Set<number>()
    const users: PlainUser[] = []
    for (const taskList of Object.values(tasks.value)) {
        for (const task of taskList) {
            if (task.assigned_user?.id && !seen.has(task.assigned_user.id)) {
                seen.add(task.assigned_user.id)
                users.push(task.assigned_user)
            }
        }
    }
    return users.sort((a, b) => {
        const aLabel = `${a.surname} ${a.name}`.toLowerCase()
        const bLabel = `${b.surname} ${b.name}`.toLowerCase()
        return aLabel.localeCompare(bLabel)
    })
})

const presets = computed(() => [
    {
        key: 'my',
        icon: 'mdi:account',
        label: t('board.filters.preset.my'),
    },
    {
        key: 'active',
        icon: 'mdi:play-circle',
        label: t('board.filters.preset.active'),
    },
    {
        key: 'noDeadline',
        icon: 'mdi:calendar-remove',
        label: t('board.filters.preset.no_deadline'),
    },
])

const activePreset = computed(() => {
    const f = filters.value
    if (f.assignedUserId === authStore.user.id
      && !f.priority
      && !f.deadlineFrom && !f.deadlineTo
      && !f.createdAtFrom && !f.createdAtTo
      && f.storyPointsFrom === undefined && f.storyPointsTo === undefined
      && !f.search
      && !f.hideFinishStatus) return 'my'
    if (f.hideFinishStatus && !f.assignedUserId
      && !f.priority
      && !f.deadlineFrom && !f.deadlineTo
      && !f.createdAtFrom && !f.createdAtTo
      && f.storyPointsFrom === undefined && f.storyPointsTo === undefined
      && !f.search) return 'active'
    if (f.deadlineTo === 0 && !f.assignedUserId && !f.priority
      && !f.deadlineFrom
      && !f.createdAtFrom && !f.createdAtTo
      && f.storyPointsFrom === undefined && f.storyPointsTo === undefined
      && !f.search && !f.hideFinishStatus) return 'noDeadline'
    return null
})

const hasActiveFilters = computed(() => activeFilterCount.value > 0)

const activeFilterCount = computed(() => {
    const f = filters.value
    let count = 0
    if (f.search) count++
    if (f.assignedUserId) count++
    if (f.priority) count++
    if (f.deadlineFrom || f.deadlineTo) count++
    if (f.createdAtFrom || f.createdAtTo) count++
    if (f.storyPointsFrom !== undefined || f.storyPointsTo !== undefined) count++
    if (f.hideFinishStatus) count++
    return count
})

const applyPreset = (key: string) => {
    boardStore.resetFilters()
    activeCollapse.value = ['filters']

    switch (key) {
        case 'my':
            filters.value.assignedUserId = authStore.user.id
            break
        case 'active':
            filters.value.hideFinishStatus = true
            break
        case 'noDeadline':
            filters.value.deadlineTo = 0
            break
    }
}

const handleDeadlineChange = (val: [number, number] | undefined) => {
    deadlineRange.value = val
    if (val) {
        filters.value.deadlineFrom = Math.floor(val[0] / 1000)
        filters.value.deadlineTo = Math.floor(val[1] / 1000)
    }
    else {
        filters.value.deadlineFrom = undefined
        filters.value.deadlineTo = undefined
    }
}

const handleCreatedAtChange = (val: [number, number] | undefined) => {
    createdAtRange.value = val
    if (val) {
        filters.value.createdAtFrom = Math.floor(val[0] / 1000)
        filters.value.createdAtTo = Math.floor(val[1] / 1000)
    }
    else {
        filters.value.createdAtFrom = undefined
        filters.value.createdAtTo = undefined
    }
}

const handleReset = () => {
    boardStore.resetFilters()
    deadlineRange.value = undefined
    createdAtRange.value = undefined
}
</script>
