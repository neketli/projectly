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
                    {{ $t('project.title') }}
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
            </template>
        </ElPageHeader>

        <h3 class="text-4xl mt-2">
            {{ board.title }}
        </h3>

        <ElDialog
            v-model="dialog.board"
            :title="$t('board.update.title')"
            align-center
            destroy-on-close
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
    layout: 'team',
})

const projectStore = useProjectStore()
const { getBoard, deleteBoard } = useBoard()

const board = ref({} as Board)

const dialog = reactive({
    board: false,
})

const handleBoardUpdated = (updated: Board) => {
    board.value = updated
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

onMounted(async () => {
    try {
        board.value = await getBoard(Number(boardId))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
})
</script>
