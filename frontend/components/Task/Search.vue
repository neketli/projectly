<template>
    <ElSelect
        v-model="value"
        filterable
        remote
        reserve-keyword
        value-key="id"
        :placeholder="$t('task.search.placeholder')"
        :no-data-text="$t('task.search.empty')"
        :no-match-text="$t('task.search.empty')"
        placement="bottom-end"
        :remote-method="handleSearch"
        :loading="isLoading"
        class="max-w-64 min-w-32"
        @change="handleSelect"
    >
        <ElOption
            v-for="item in options"
            :key="item.id"
            :label="item.title"
            :value="item"
            class="!h-auto"
        >
            <div class="flex flex-col">
                <h5
                    class="text-md font-semibold"
                    v-html="highlightSearchText(item.title, currentSearch)"
                />
                <span
                    v-if="item.description && currentSearch"
                    class="pb-2 text-sm font-light text-gray-700 dark:text-gray-300"
                    v-html="highlightSearchText(
                        truncateString(item.description, currentSearch),
                        currentSearch,
                    )"
                />
            </div>
        </ElOption>

        <template #loading>
            <Icon
                name="mdi:loading"
                class="w-8 h-8 animate-spin text-blue-600"
            />
        </template>
    </ElSelect>
</template>

<script setup lang="ts">
import type { DetailedTask } from '~/types/task'

const props = defineProps<{
    teamId: number
    projectCode?: string
    boardId?: number
}>()

const { searchTask } = useTask()

const value = ref('')
const options = ref([] as DetailedTask[])
const isLoading = ref(false)
const currentSearch = ref('')

const handleSearch = async (search: string) => {
    isLoading.value = true
    currentSearch.value = search

    try {
        options.value = await searchTask({
            search,
            team_id: props.teamId,
            project_code: props.projectCode,
            board_id: props.boardId,
        })
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

const handleSelect = (task: DetailedTask) => {
    navigateTo(`/team/${props.teamId}/project/${task.project_code}/${task.meta.board_id}/task/${task.project_code}-${task.project_index}`)
}

const highlightSearchText = (text: string, searchTerm: string): string => {
    if (!text || !searchTerm) return text

    const escaped = searchTerm.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')

    const regex = new RegExp(`(${escaped})`, 'gi')
    return text.replace(regex, '<mark class="bg-yellow-300 bg-opacity-80 rounded">$1</mark>')
}

const truncateString = (text?: string, searchTerm?: string, contextLength: number = 15): string => {
    if (!text || !searchTerm) return ''

    const lowerText = text.toLowerCase()
    const lowerSearchTerm = searchTerm.toLowerCase()
    const matchIndex = lowerText.indexOf(lowerSearchTerm)

    if (matchIndex === -1) {
        return text.length > contextLength * 2
            ? text.substring(0, contextLength * 2) + '...'
            : text
    }

    const startIndex = Math.max(0, matchIndex - contextLength)
    const endIndex = Math.min(text.length, matchIndex + searchTerm.length + contextLength)

    let result = text.substring(startIndex, endIndex)

    if (startIndex > 0) {
        result = '...' + result
    }

    if (endIndex < text.length) {
        result = result + '...'
    }

    return result
}
</script>
