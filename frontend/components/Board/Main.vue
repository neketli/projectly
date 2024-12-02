<template>
    <section class="main-board w-full h-full p-4">
        <ElScrollbar v-if="sortedStatusList.length">
            <div
                ref="parent"
                class="flex h-[72vh] py-1 mx-4 gap-4"
            >
                <StatusItem
                    v-for="status in sortedStatusList"
                    :key="status.id+status.order"
                    :status="status"
                    @create="handleCreateStatus"
                    @update="handleUpdateStatus"
                    @delete="handleDeleteStatus"
                />
            </div>
        </ElScrollbar>

        <ElEmpty
            v-else
            :description="t('board.empty')"
        />
    </section>
</template>

<script lang="ts" setup>
import type { Status } from '~/types/board'

const { t } = useI18n()
const [parent] = useAutoAnimate()

const boardStore = useBoardStore()

const { statusList, sortedStatusList } = toRefs(boardStore)
const { getStatusList, createStatus, updateStatus, deleteStatus } = useStatus()

const handleCreateStatus = async (status: Status) => {
    try {
        const newStatus = await createStatus(status)
        boardStore.replaceStatus(newStatus)
        ElMessage.success(t('status.success.create'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleUpdateStatus = async (status: Status) => {
    try {
        await updateStatus(status, statusList.value.find(s => s.id === status.id)?.order)
        statusList.value = await getStatusList(boardStore.board.id)
        ElMessage.success(t('status.success.update'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleDeleteStatus = async (status: Status) => {
    try {
        if (status.id !== 0) {
            await deleteStatus(status)
            ElMessage.success(t('status.success.delete'))
        }
        boardStore.deleteStatus(status)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}
</script>
