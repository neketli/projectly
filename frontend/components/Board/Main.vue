<template>
    <section class="main-board w-full h-full py-4">
        <ElScrollbar
            v-if="sortedStatusList.length"
            height="72vh"
            view-class="p-2"
        >
            <div
                ref="parent"
                class="!w-full max-h-[70vh] flex gap-4"
                :class="{
                    'justify-center': sortedStatusList.length <= 4,
                }"
            >
                <StatusItem
                    v-for="status in sortedStatusList"
                    :key="status.id+status.order"
                    :status="status"
                    @create="handleCreateStatus"
                    @update="handleUpdateStatus"
                    @delete="handleDeleteStatus"
                />
                <!-- TODO: fix without tricks -->
                <div class="min-w-1" />
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
        const statuses = await getStatusList(boardStore.board.id)
        boardStore.setStatusList(statuses)
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
