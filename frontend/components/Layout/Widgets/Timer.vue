<template>
    <ElPopover class="pomodoro-layout-widget">
        <template #reference>
            <div class="text-center">
                <div
                    class="text-2xl font-mono font-bold text-gray-800 dark:text-gray-200"
                >
                    {{ formattedTime }}
                </div>
            </div>
        </template>

        <div class="flex flex-col gap-2 justify-center">
            <ElButton
                :type="isRunning ? 'warning' : 'primary'"
                size="small"
                @click="toggleTimer"
            >
                <Icon
                    :name="isRunning ? 'mdi-pause' : 'mdi-play'"
                    class="ml-2"
                />
            </ElButton>
        </div>
    </ElPopover>
</template>

<script setup lang="ts">
const store = useTimerStore()

const formattedTime = computed(() => store.formattedTime)
const isRunning = computed(() => store.isRunning)

const toggleTimer = () => {
    if (store.isRunning) {
        store.pause()
    }
    else {
        store.start()
    }
}

onMounted(() => {
    if (!store.timer) {
        store.timer = setInterval(() => {
            store.tick()
        }, 1000)
    }
    store.synchronizeTime()
})
</script>

<style scoped>
:deep(.el-progress-bar__outer) {
    transition: all 0.3s ease;
}
</style>
