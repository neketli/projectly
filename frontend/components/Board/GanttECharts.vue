<template>
    <div class="board-gantt-echarts h-[60vh] flex flex-col">
        <div
            v-if="!chartData.length"
            class="flex items-center justify-center flex-1 text-gray-400"
        >
            {{ $t('common.no_data') }}
        </div>
        <template v-else>
            <div class="flex items-center justify-end px-2 py-1 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-900">
                <ElRadioGroup
                    v-model="groupMode"
                    size="small"
                >
                    <ElRadioButton value="task">
                        {{ t('board.gantt.group_task') }}
                    </ElRadioButton>
                    <ElRadioButton value="person">
                        {{ t('board.gantt.group_person') }}
                    </ElRadioButton>
                </ElRadioGroup>
            </div>
            <VChart
                class="flex-1 w-full"
                autoresize
                :option="chartOption"
                @click="handleChartClick"
            />
        </template>
    </div>
</template>

<script lang="ts" setup>
import 'echarts'
import VChart from 'vue-echarts'
import dayjs from 'dayjs'
import type { DetailedTask } from '~/types/task'
// eslint-disable-next-line @typescript-eslint/no-explicit-any
type EChartsApi = any
// eslint-disable-next-line @typescript-eslint/no-explicit-any
type EChartsParams = any

const { t } = useI18n()
const route = useRoute()
const groupMode = ref<'task' | 'person'>('task')

const boardStore = useBoardStore()
const { matchesFilters } = useBoardFilters()

const { tasks } = toRefs(boardStore)

const allTasks = computed<DetailedTask[]>(() => {
    return Object.values(tasks.value).flat().filter(matchesFilters)
})

function getPlanStart(task: DetailedTask): number | null {
    if (task.deadline && task.story_points) return task.deadline - (task.story_points * 3600)
    if (task.created_at) return task.created_at
    return null
}

function getPlanEnd(task: DetailedTask): number | null {
    const start = getPlanStart(task)
    if (!start) return null

    if (task.story_points) return start + (task.story_points * 3600)

    if (task.deadline) return task.deadline + 86400

    if (task.finished_at) return task.finished_at + 86400

    if (task.tracked_time) return start + (task.tracked_time * 3600)

    return null
}

function getActualEnd(task: DetailedTask): number | null {
    if (task.tracked_time) {
        const start = getPlanStart(task)
        if (start) return start + (task.tracked_time * 3600)
    }
    if (task.finished_at) return task.finished_at + 86400
    return null
}

function getUserLabel(task: DetailedTask): string {
    const user = task.assigned_user
    if (user?.name || user?.surname) return `${user.name} ${user.surname}`.trim()
    return t('board.gantt.unassigned')
}

const taskModeData = computed(() => {
    const rows: {
        xStart: number
        xEnd: number
        yCategory: string
        planColor: string
        actualEnd: number | null
        task: DetailedTask
    }[] = []

    for (const task of allTasks.value) {
        const planStart = getPlanStart(task)
        const planEnd = getPlanEnd(task)
        if (!planStart || !planEnd) continue

        rows.push({
            xStart: planStart * 1000,
            xEnd: planEnd * 1000,
            yCategory: `${task.project_code}-${task.project_index} ${task.title}`,
            planColor: task.status?.hex_color || '#409eff',
            actualEnd: getActualEnd(task) ? getActualEnd(task)! * 1000 : null,
            task,
        })
    }

    rows.sort((a, b) => a.xStart - b.xStart)
    return rows
})

const personModeData = computed(() => {
    const groups = new Map<string, {
        task: DetailedTask
        planStart: number
        planEnd: number
        actualEnd: number | null
    }[]>()

    for (const task of allTasks.value) {
        const planStart = getPlanStart(task)
        const planEnd = getPlanEnd(task)
        if (!planStart || !planEnd) continue

        const label = getUserLabel(task)
        if (!groups.has(label)) groups.set(label, [])
        groups.get(label)!.push({
            task,
            planStart: planStart * 1000,
            planEnd: planEnd * 1000,
            actualEnd: getActualEnd(task) ? getActualEnd(task)! * 1000 : null,
        })
    }

    const rows: {
        xStart: number
        xEnd: number
        yCategory: string
        planColor: string
        actualEnd: number | null
        task: DetailedTask
        barIndex: number
        barCount: number
    }[] = []

    const sortedLabels = [...groups.keys()].sort((a, b) => groups.get(b)!.length - groups.get(a)!.length)

    for (const label of sortedLabels) {
        const items = groups.get(label)!
        items.sort((a, b) => a.planStart - b.planStart)
        items.forEach((item, idx) => {
            rows.push({
                xStart: item.planStart,
                xEnd: item.planEnd,
                yCategory: label,
                planColor: item.task.status?.hex_color || '#409eff',
                actualEnd: item.actualEnd,
                task: item.task,
                barIndex: idx,
                barCount: items.length,
            })
        })
    }

    return rows
})

const chartData = computed(() => groupMode.value === 'person' ? personModeData.value : taskModeData.value)

const yCategories = computed(() => {
    const seen = new Set<string>()
    const result: string[] = []
    for (const row of chartData.value) {
        if (!seen.has(row.yCategory)) {
            seen.add(row.yCategory)
            result.push(row.yCategory)
        }
    }
    return result
})

const minTime = computed(() => {
    if (!chartData.value.length) return dayjs().subtract(1, 'month').valueOf()
    let min = Infinity
    for (const row of chartData.value) {
        if (row.xStart < min) min = row.xStart
    }
    return dayjs(min).startOf('month').valueOf()
})

const maxTime = computed(() => {
    if (!chartData.value.length) return dayjs().add(1, 'month').valueOf()
    let max = -Infinity
    for (const row of chartData.value) {
        if (row.xEnd > max) max = row.xEnd
        if (row.actualEnd && row.actualEnd > max) max = row.actualEnd
    }
    return dayjs(max).endOf('month').valueOf()
})

function renderGanttItem(params: EChartsParams, api: EChartsApi) {
    const startVal = api.value(0)
    const endVal = api.value(1)
    const yIdx = api.value(2)
    const planColor = api.value(3)
    const actualEndVal = api.value(4)
    const barIndex = api.value(5) || 0
    const barCount = (api.value(6) || 1)

    const start = api.coord([startVal, yIdx])
    const end = api.coord([endVal, yIdx])

    if (!start || !end) return

    const bandHeight = Math.max(Math.abs(api.size([0, 1])[1]), 4)
    const barBand = bandHeight / barCount
    const barHeight = Math.max(barBand * 0.85, 2)
    const barOffset = (barBand * barIndex) + ((barBand - barHeight) / 2)
    const yTop = start[1] - (bandHeight / 2) + barOffset
    const barLeft = start[0]
    const barWidthTotal = Math.max(end[0] - start[0], 1)
    const barRadius = 2

    const children: Record<string, unknown>[] = []

    if (actualEndVal) {
        const actual = api.coord([actualEndVal, yIdx])
        if (actual) {
            const actualWidth = Math.max(actual[0] - start[0], 1)
            children.push({
                type: 'rect',
                shape: {
                    x: barLeft,
                    y: yTop,
                    width: actualWidth,
                    height: barHeight,
                    r: barRadius,
                },
                style: {
                    fill: planColor + '80',
                    stroke: planColor,
                    lineWidth: 1,
                },
                emphasis: {
                    style: {
                        shadowBlur: 6,
                        shadowColor: 'rgba(0,0,0,0.2)',
                    },
                },
            })
        }
    }

    children.push({
        type: 'rect',
        shape: {
            x: barLeft,
            y: yTop,
            width: barWidthTotal,
            height: barHeight,
            r: barRadius,
        },
        style: {
            fill: planColor + '1A',
            stroke: planColor,
            lineWidth: 1,
            lineDash: [4, 3],
        },
        emphasis: {
            style: {
                shadowBlur: 6,
                shadowColor: 'rgba(0,0,0,0.2)',
            },
        },
    })

    return {
        type: 'group',
        children,
    }
}

const chartOption = computed(() => ({
    tooltip: {
        trigger: 'item',
        formatter: (params: EChartsParams) => {
            const task = chartData.value[params.dataIndex]?.task
            if (!task) return ''
            const fmt = (ts: number) => dayjs(ts).format('DD.MM.YYYY')
            const fmtH = (ts: number) => dayjs(ts).format('DD.MM.YYYY HH:mm')
            const user = task.assigned_user
            const userName = user ? `${user.name} ${user.surname}`.trim() : '-'
            return `
                <b>${task.project_code}-${task.project_index}</b><br/>
                ${task.title}<br/>
                <span style="display:inline-block;width:10px;height:10px;border-radius:50%;background:${task.status?.hex_color || '#409eff'};margin-right:4px"></span>
                ${task.status?.title || ''}<br/>
                ${t('task.form.assigned_user')}: ${userName}<br/>
                ${t('task.form.deadline')}: ${task.deadline ? fmt(task.deadline * 1000) : '-'}<br/>
                ${t('task.form.created_at')}: ${fmtH(task.created_at * 1000)}<br/>
                ${task.finished_at ? t('task.form.finished_at') + ': ' + fmt(task.finished_at! * 1000) : ''}
            `
        },
    },
    grid: {
        left: 200,
        right: 40,
        top: 10,
        bottom: 60,
    },
    xAxis: {
        type: 'time',
        min: minTime.value,
        max: maxTime.value,
        axisLabel: {
            formatter: (val: number) => dayjs(val).format('MMM DD'),
        },
        splitLine: {
            show: true,
            lineStyle: { type: 'dashed', color: '#e5e7eb' },
        },
    },
    yAxis: {
        type: 'category',
        data: yCategories.value,
        axisLabel: {
            width: 180,
            overflow: 'truncate',
            fontSize: 11,
        },
        splitLine: {
            show: true,
            lineStyle: { color: '#f3f4f6' },
        },
    },
    dataZoom: [
        {
            type: 'inside',
            xAxisIndex: 0,
            minSpan: 5,
        },
        {
            type: 'slider',
            xAxisIndex: 0,
            bottom: 10,
            height: 20,
            borderColor: '#d1d5db',
            fillerColor: 'rgba(64,158,255,0.15)',
            handleStyle: { color: '#409eff' },
            labelFormatter: (val: number) => dayjs(val).format('DD.MM'),
        },
        {
            type: 'inside',
            yAxisIndex: 0,
        },
    ],
    series: [
        {
            type: 'custom',
            renderItem: renderGanttItem,
            data: chartData.value.map(row => [
                row.xStart,
                row.xEnd,
                row.yCategory,
                row.planColor,
                row.actualEnd,
                'barIndex' in row ? row.barIndex : 0,
                'barCount' in row ? row.barCount : 1,
            ]),
            encode: {
                x: [0, 1],
                y: 2,
            },
            tooltip: {
                formatter: (params: EChartsParams) => {
                    const task = chartData.value[params.dataIndex]?.task
                    if (!task) return ''
                    const fmt = (ts: number) => dayjs(ts).format('DD.MM.YYYY')
                    const fmtH = (ts: number) => dayjs(ts).format('DD.MM.YYYY HH:mm')
                    const user = task.assigned_user
                    const userName = user ? `${user.name} ${user.surname}`.trim() : '-'
                    return `
                        <b>${task.project_code}-${task.project_index}</b><br/>
                        ${task.title}<br/>
                        <span style="display:inline-block;width:10px;height:10px;border-radius:50%;background:${task.status?.hex_color || '#409eff'};margin-right:4px"></span>
                        ${task.status?.title || ''}<br/>
                        ${t('task.form.assigned_user')}: ${userName}<br/>
                        ${t('task.form.deadline')}: ${task.deadline ? fmt(task.deadline * 1000) : '-'}<br/>
                        ${t('task.form.created_at')}: ${fmtH(task.created_at * 1000)}<br/>
                        ${task.finished_at ? t('task.form.finished_at') + ': ' + fmt(task.finished_at! * 1000) : ''}
                    `
                },
            },
        },
    ],
}))

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const handleChartClick = (params: any) => {
    const row = chartData.value[params.dataIndex]
    if (row) {
        navigateTo(`${route.path}/task/${row.task.project_code}-${row.task.project_index}`)
    }
}
</script>
