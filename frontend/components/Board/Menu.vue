<template>
    <div>
        <ElMenu
            :collapse="isMenuCollapsed"
            :default-active="`${route.params.boardId}`"
            class="!h-full max-lg:!border-r-0 "
        >
            <div
                v-if="!isMenuCollapsed"
                class="w-full min-w-52 p-2"
            >
                <div class="flex justify-between mb-2">
                    <h3 class="text-xl">
                        {{ $t('board.list.title') }}
                    </h3>

                    <ElButton
                        v-if="projectStore.isEditAvailable"
                        @click="dialog.board = true"
                    >
                        <Icon name="mdi:plus" />
                    </ElButton>
                </div>

                <ElInput
                    v-model.trim="search"
                    :placeholder="$t('common.search')"
                />
            </div>

            <div
                v-else-if="projectStore.isEditAvailable"
                class="w-full p-2"
            >
                <ElButton
                    class="w-full"
                    @click="dialog.board = true"
                >
                    <Icon name="mdi:plus" />
                </ElButton>
            </div>

            <ElMenuItem
                v-for="board in boardItems"
                :key="board.id"
                :index="`${board.id}`"
                @click="handleBoardClick(board)"
            >
                <ElTag
                    type="info"
                    class="mr-2 h-6 w-6"
                >
                    {{ [...board.title][0].toUpperCase() }}
                </ElTag>
                <template #title>
                    {{ board.title }}
                </template>
            </ElMenuItem>

            <ElEmpty
                v-if="!isMenuCollapsed && !boardItems.length"
                v-loading="isLoading"
                :description="$t('common.no_data')"
            />

            <ElButton
                class="w-full !hidden lg:!block"
                text
                @click="handleCollapse"
            >
                <Icon
                    name="mdi:chevron-down"
                    class="transition-all"
                    :class="isMenuCollapsed ? 'transform -rotate-90' : 'transform rotate-90'"
                />
            </ElButton>
        </ElMenu>

        <ElDialog
            v-model="dialog.board"
            :title="$t('board.create.title')"
            align-center
            destroy-on-close
            class="max-md:!w-4/5"
        >
            <BoardForm
                :project-id="projectId"
                @cancel="dialog.board = false"
                @success="handleBoardCreated"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { Board } from '~/types/board'

const props = defineProps<{ projectId: number }>()
const emit = defineEmits<{
    (event: 'board-clicked', board: Board): void
}>()

const route = useRoute()

const projectStore = useProjectStore()
const { getBoardsList } = useBoard()

const search = ref('')
const isMenuCollapsed = ref(false)

const isLoading = ref(false)

const dialog = reactive({
    board: false,
})

const boardItems = computed(() =>
    !search.value
        ? projectStore.boardList
        : projectStore.boardList.filter(item =>
            item.title.toLowerCase().includes(search.value.toLowerCase()),
        ),
)

const handleBoardClick = async (board: Board) => {
    navigateTo(`/team/${route.params.teamId}/project/${projectStore.project.code}/${board.id}`)
    emit('board-clicked', board)
}

const handleBoardCreated = (board: Board) => {
    projectStore.boardList.push(board)
    dialog.board = false
}

const handleCollapse = () => {
    isMenuCollapsed.value = !isMenuCollapsed.value

    if (isMenuCollapsed.value) {
        search.value = ''
    }
}

onMounted(async () => {
    try {
        isLoading.value = true
        if (projectStore.boardList.length === 0) {
            projectStore.boardList = await getBoardsList(props.projectId)
        }
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
