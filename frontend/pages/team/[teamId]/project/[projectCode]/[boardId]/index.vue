<template>
    <div class="board">
        <ElPageHeader @back="navigateTo(`/team/${teamId}/project/${projectCode}`)">
            <template #content>
                <h1 class="text-2xl">
                    {{ $t('board.title') }}
                </h1>
            </template>

            <template #title>
                <h3>
                    {{ $t('project.overview') }}
                </h3>
            </template>

            <template #extra>
                <ElPopconfirm
                    v-if="projectStore.isDeleteAvailable"
                    width="250"
                    :title="$t('board.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteBoard"
                >
                    <template #reference>
                        <ElButton
                            :title="$t('board.delete.title')"
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
                    :title="$t('board.update.title')"
                    circle
                    plain
                    @click="dialog.board = true"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>

                <ElButton
                    v-if="projectStore.isEditAvailable"
                    :title="$t('status.create.title')"
                    type="primary"
                    circle
                    plain
                    @click="handleAddStatus"
                >
                    <Icon name="mdi:file-plus" />
                </ElButton>
            </template>
        </ElPageHeader>

        <h3 class="text-4xl mt-2">
            {{ board.title }}
        </h3>

        <ElTabs
            v-model="activeTab"
            v-loading="isLoading"
        >
            <ElTabPane
                :label="$t('board.tab.board')"
                name="board"
            >
                <BoardMain />
            </ElTabPane>

            <ElTabPane
                :label="$t('board.tab.table')"
                name="table"
            >
                <BoardTable />
            </ElTabPane>
        </ElTabs>

        <ElDialog
            v-model="dialog.board"
            :title="$t('board.update.title')"
            align-center
            destroy-on-close
            class="max-md:!w-4/5"
        >
            <BoardForm
                :board="board"
                :project-id="Number(projectStore.project.id)"
                @success="handleBoardUpdated"
                @cancel="dialog.board = false"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { Board } from '~/types/board'

const { teamId, projectCode, boardId } = useRoute().params
const { t } = useI18n()

useHead({
    title: t('board.title'),
})

definePageMeta({
    layout: 'project',
})

const projectStore = useProjectStore()
const boardStore = useBoardStore()
const { board, statusList } = toRefs(boardStore)

const { getBoard, deleteBoard } = useBoard()
const { getStatusList } = useStatus()
const { getTasksList } = useTask()

const dialog = reactive({
    board: false,
})

const activeTab = ref('board')
const isLoading = ref(false)

const handleBoardUpdated = (updated: Board) => {
    board.value = updated
    projectStore.boardList = projectStore.boardList.map(item => item.id === Number(boardId) ? updated : item)
    dialog.board = false
}

const handleDeleteBoard = async () => {
    try {
        await deleteBoard(board.value.id)
        ElMessage.success(t('board.success.delete'))
        navigateTo(`/team/${teamId}/project/${projectCode}`)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleAddStatus = () => {
    activeTab.value = 'board'
    statusList.value.push({
        id: 0,
        title: t('status.create.default'),
        order: statusList.value.length,
        board_id: Number(boardId),
        hex_color: '',
    })
}

onMounted(async () => {
    try {
        isLoading.value = true
        board.value = await getBoard(Number(boardId))
        const statuses = await getStatusList(Number(boardId))
        boardStore.setStatusList(statuses)
        const tasks = await getTasksList(Number(boardId))
        boardStore.setTaskList(tasks)
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
