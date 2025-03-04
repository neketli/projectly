<template>
    <div>
        <ElCard>
            <template #header>
                <div class="flex justify-between">
                    <h3 class="text-xl">
                        {{ $t('board.table.title') }}
                    </h3>

                    <div class="flex items-center gap-4">
                        <ElRadioGroup v-model="view">
                            <ElRadioButton value="grid">
                                <Icon name="mdi:view-grid-outline" />
                            </ElRadioButton>
                            <ElRadioButton value="list">
                                <Icon name="mdi:view-list-outline" />
                            </ElRadioButton>
                        </ElRadioGroup>

                        <ElButton
                            v-if="projectStore.isEditAvailable"
                            @click="dialog.board = true"
                        >
                            <Icon name="mdi:plus" />
                        </ElButton>
                    </div>
                </div>
            </template>

            <div
                v-if="view === 'grid'"
                class="w-full flex flex-col gap-4"
            >
                <ElInput
                    v-model.trim="search"
                    :placeholder="$t('common.search')"
                />

                <div
                    v-if="boardItems.length"
                    v-loading="isLoading"
                    class="flex flex-wrap gap-4"
                >
                    <BoardCard
                        v-for="board in boardItems"
                        :key="board.id"
                        :board="board"
                        class="cursor-pointer min-w-32 max-w-64"
                        @click="handleBoardClick(board)"
                    />
                </div>

                <ElEmpty
                    v-else
                    :description="$t('common.no_data')"
                />
            </div>

            <ElTable
                v-else
                :data="boardItems"
                :loading="isLoading"
                height="250"
                row-class-name="cursor-pointer"
                class="w-full"
                :empty-text="$t('common.no_data')"
                @row-click="handleBoardClick"
            >
                <ElTableColumn
                    prop="title"
                    :label="$t('board.form.title')"
                />

                <ElTableColumn
                    width="150"
                    align="right"
                >
                    <template #header>
                        <ElInput
                            v-model.trim="search"
                            size="small"
                            :placeholder="$t('common.search')"
                        />
                    </template>
                </ElTableColumn>
            </ElTable>
        </ElCard>

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

const route = useRoute()

const projectStore = useProjectStore()
const { getBoardsList } = useBoard()

const boards: Ref<Board[]> = ref([])
const search = ref('')
const view = ref('list')

const isLoading = ref(false)

const dialog = reactive({
    board: false,
})

const boardItems = computed(() =>
    !search.value
        ? boards.value
        : boards.value.filter(item =>
            item.title.toLowerCase().includes(search.value.toLowerCase()),
        ),
)

const handleBoardClick = async (board: Board) => {
    navigateTo(`${route.path}/${board.id}`)
}

const handleBoardCreated = (board: Board) => {
    boards.value.push(board)
    dialog.board = false
}

onMounted(async () => {
    try {
        isLoading.value = true
        boards.value = await getBoardsList(props.projectId)
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
