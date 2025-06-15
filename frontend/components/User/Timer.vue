<template>
    <ElCard class="pomodoro-widget bg-white">
        <div class="text-center mb-6">
            <h2 class="flex items-center justify-center gap-2 text-2xl font-bold text-gray-800 dark:text-gray-200 mb-2">
                {{ t('user.timer.title') }}

                <Icon name="mdi-timer-check" />
            </h2>

            <ElTag
                :type="currentSession.type === 'work' ? 'danger' : 'success'"
                size="large"
            >
                {{ currentSession.label }}
            </ElTag>
        </div>

        <div class="text-center mb-6">
            <div
                class="text-6xl font-mono font-bold text-gray-800 dark:text-gray-200 mb-2"
            >
                {{ formattedTime }}
            </div>
            <ElProgress
                :percentage="progressPercentage"
                :stroke-width="8"
                :show-text="false"
                :color="currentSession.type === 'work' ? '#007BFF' : '#28A745'"
            />
        </div>

        <div class="flex flex-col gap-2 justify-center mb-6">
            <ElButton
                :type="isRunning ? 'warning' : 'primary'"
                size="large"
                @click="toggleTimer"
            >
                {{ isRunning ? t('user.timer.pause') : t('user.timer.start') }}
                <Icon
                    :name="isRunning ? 'mdi-pause' : 'mdi-play'"
                    class="ml-2"
                />
            </ElButton>
            <div class="flex w-full">
                <ElButton
                    type="info"
                    size="large"
                    class="w-1/2"
                    @click="resetTimer"
                >
                    {{ t('user.timer.reset') }}
                    <Icon
                        name="mdi-restore"
                        class="ml-2"
                    />
                </ElButton>

                <ElButton
                    size="large"
                    class="w-1/2"
                    @click="skipSession"
                >
                    {{ t('user.timer.skip') }}
                    <Icon
                        name="mdi-debug-step-over"
                        class="ml-2"
                    />
                </ElButton>
            </div>
        </div>

        <ElCollapse v-model="settingsOpen">
            <ElCollapseItem
                :title="t('user.timer.settings.title')"
                name="settings"
                class="dark:text-gray-300"
            >
                <div class="space-y-4">
                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.work_duration') }}
                        </label>
                        <ElInputNumber
                            v-model="settings.workDuration"
                            :min="1"
                            :max="60"
                        />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.short_break') }}
                        </label>
                        <ElInputNumber
                            v-model="settings.shortBreakDuration"
                            :min="1"
                            :max="60"
                        />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.long_break') }}
                        </label>
                        <ElInputNumber
                            v-model="settings.longBreakDuration"
                            :min="1"
                            :max="60"
                        />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.sessions') }}
                        </label>
                        <ElInputNumber
                            v-model="settings.sessionsUntilLongBreak"
                            :min="2"
                        />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.auto_break') }}
                        </label>
                        <ElSwitch v-model="settings.autoStartBreaks" />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.sound') }}
                        </label>
                        <ElSwitch v-model="settings.soundNotifications" />
                    </div>

                    <div class="flex items-center justify-between gap-2">
                        <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
                            {{ t('user.timer.settings.widget') }}
                        </label>
                        <ElSwitch v-model="settings.headerWidget" />
                    </div>
                </div>
            </ElCollapseItem>
        </ElCollapse>

        <div class="mt-6 p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
            <div class="text-center">
                <div class="text-lg font-semibold text-gray-800 dark:text-gray-200 mb-2">
                    {{ t('user.timer.statistics.title') }}
                </div>
                <div class="grid grid-cols-2 gap-4 text-sm">
                    <div>
                        <div class="font-medium text-gray-600 dark:text-gray-300">
                            {{ t('user.timer.statistics.today') }}
                        </div>
                        <div class="text-xl font-bold text-cyan-500">
                            {{ completedSessions }}
                        </div>
                    </div>
                    <div>
                        <div class="font-medium text-gray-600 dark:text-gray-300">
                            {{ t('user.timer.statistics.total') }}
                        </div>
                        <div class="text-xl font-bold text-blue-500">
                            {{ totalActiveTime }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </ElCard>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import duration from 'dayjs/plugin/duration'

dayjs.extend(duration)

const { t } = useI18n()

const store = useTimerStore()
const settingsOpen = ref([])

const formattedTime = computed(() => store.formattedTime)
const progressPercentage = computed(() => store.progressPercentage)
const isRunning = computed(() => store.isRunning)
const completedSessions = computed(() => store.completedSessions)
const settings = computed(() => store.settings)
const currentSession = computed(() => {
    const sessions = {
        work: { label: t('user.timer.type.work'), type: 'work' },
        shortBreak: { label: t('user.timer.type.short_break'), type: 'break' },
        longBreak: { label: t('user.timer.type.long_break'), type: 'break' },
    }
    return sessions[store.currentSessionType as keyof typeof sessions]
})
const totalActiveTime = computed(() => {
    return dayjs.duration(store.focusedSeconds, 'second').format('HH:mm:ss')
})

const toggleTimer = () => {
    if (store.isRunning) {
        store.pause()
    }
    else {
        store.start()
    }
}

const resetTimer = () => {
    store.reset()
}

const skipSession = () => {
    store.skipSession()
}

watch(() => store.settings, () => {
    store.reset()
}, { deep: true })

onMounted(() => {
    store.synchronizeTime()

    const lastStart = dayjs(store.lastStartTimestamp)
    if (lastStart.isBefore(dayjs().startOf('day'))) {
        store.clearStatistics()
    }
})
</script>
