import dayjs from 'dayjs'
import { defineStore } from 'pinia'

export const useTimerStore = defineStore('projectly-timer', {
    persist: {
        storage: persistedState.localStorage,
    },
    state: () => ({
        timeLeft: 25 * 60, // seconds
        isRunning: false,
        currentSessionType: 'work' as 'work' | 'shortBreak' | 'longBreak',
        completedSessions: 0,
        focusedSeconds: 0,
        lastStartTimestamp: 0,
        startTimestamp: 0,
        endTimestamp: 0,
        pausedAt: 0,
        settings: {
            workDuration: 25, // minutes
            shortBreakDuration: 5,
            longBreakDuration: 15,
            sessionsUntilLongBreak: 4,
            autoStartBreaks: false,
            soundNotifications: true,
            headerWidget: true,
        },

        timer: undefined as ReturnType<typeof setInterval> | undefined,
    }),

    getters: {
        formattedTime: (state) => {
            const minutes = Math.floor(state.timeLeft / 60)
            const seconds = state.timeLeft % 60
            return `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`
        },

        progressPercentage: (state): number => {
            const totalTime = state.currentSessionType === 'work'
                ? state.settings.workDuration * 60
                : (state.completedSessions && state.completedSessions % state.settings.sessionsUntilLongBreak === 0)
                    ? state.settings.longBreakDuration * 60
                    : state.settings.shortBreakDuration * 60

            return ((totalTime - state.timeLeft) / totalTime) * 100
        },

        endTime: (state) => {
            return state.endTimestamp ? dayjs(state.endTimestamp) : null
        },
    },

    actions: {
        start() {
            const now = dayjs()
            this.startTimestamp = now.valueOf()
            this.endTimestamp = now.add(this.timeLeft, 'seconds').valueOf()
            this.lastStartTimestamp = now.valueOf()

            this.isRunning = true
            this.pausedAt = 0
        },

        pause() {
            this.pausedAt = dayjs().valueOf()
            this.isRunning = false
            this.startTimestamp = 0
            this.endTimestamp = 0
        },

        reset() {
            this.isRunning = false
            this.startTimestamp = 0
            this.endTimestamp = 0
            this.pausedAt = 0

            this.timeLeft = this.currentSessionType === 'work'
                ? this.settings.workDuration * 60
                : (this.completedSessions > 0 && this.completedSessions % this.settings.sessionsUntilLongBreak === 0)
                    ? this.settings.longBreakDuration * 60
                    : this.settings.shortBreakDuration * 60
        },

        tick() {
            if (!this.isRunning || !this.endTimestamp) return

            if (this.currentSessionType === 'work') {
                this.focusedSeconds++
            }

            const now = dayjs()
            const remainingTime = Math.max(0, dayjs(this.endTimestamp).diff(now, 'seconds'))

            this.timeLeft = remainingTime

            if (remainingTime === 0) {
                this.completeSession()
            }
        },

        completeSession() {
            const { $i18n } = useNuxtApp()
            this.isRunning = false
            this.startTimestamp = 0
            this.endTimestamp = 0
            this.pausedAt = 0

            if (this.settings.soundNotifications) {
                this.playNotificationSound()
            }

            if (this.currentSessionType === 'work') {
                this.completedSessions++

                if (this.completedSessions % this.settings.sessionsUntilLongBreak === 0) {
                    this.currentSessionType = 'longBreak'
                    this.timeLeft = this.settings.longBreakDuration * 60
                }
                else {
                    this.currentSessionType = 'shortBreak'
                    this.timeLeft = this.settings.shortBreakDuration * 60
                }

                ElNotification({
                    title: $i18n.t('user.timer.notifications.session_completed.title'),
                    message: $i18n.t('user.timer.notifications.session_completed.message'),
                    type: 'success',
                    duration: 3000,
                })
            }
            else {
                this.currentSessionType = 'work'
                this.timeLeft = this.settings.workDuration * 60

                ElNotification({
                    title: $i18n.t('user.timer.notifications.break_completed.title'),
                    message: $i18n.t('user.timer.notifications.break_completed.message'),
                    type: 'info',
                    duration: 3000,
                })
            }

            if (this.settings.autoStartBreaks) {
                setTimeout(() => {
                    this.start()
                }, 1000)
            }
        },

        skipSession() {
            this.completeSession()
        },

        playNotificationSound() {
            const audioContext = new window.AudioContext()
            const oscillator = audioContext.createOscillator()
            const gainNode = audioContext.createGain()

            oscillator.connect(gainNode)
            gainNode.connect(audioContext.destination)

            oscillator.frequency.value = 900
            oscillator.type = 'sine'

            gainNode.gain.setValueAtTime(0.3, audioContext.currentTime)
            gainNode.gain.exponentialRampToValueAtTime(0.01, audioContext.currentTime + 0.5)

            oscillator.start(audioContext.currentTime)
            oscillator.stop(audioContext.currentTime + 1)
        },

        synchronizeTime() {
            if (this.timer) {
                clearInterval(this.timer)
            }
            this.timer = setInterval(() => {
                this.tick()
            }, 1000)

            if (this.isRunning && this.endTimestamp) {
                const now = dayjs()
                const remainingTime = Math.max(0, dayjs(this.endTimestamp).diff(now, 'seconds'))
                this.timeLeft = remainingTime

                if (remainingTime === 0) {
                    this.completeSession()
                }
            }
        },

        clearStatistics() {
            this.completedSessions = 0
            this.focusedSeconds = 0
        },
    },
})
