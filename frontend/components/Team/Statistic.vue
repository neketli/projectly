<template>
    <ElCard v-loading="isLoading">
        <div class="flex flex-col md:flex-row justify-between gap-4">
            <VChart
                class="w-[90vw] md:w-[50vw] h-[40vh]"
                autoresize
                :option="clustersChartOption"
            />
            <div class="relative flex w-[90vw] md:w-[50vw] h-[40vh]">
                <VChart
                    class="w-[90vw] md:w-[50vw] h-[40vh]"
                    autoresize
                    :option="volumeChartOption"
                    @selectchanged="handleSelectChanged"
                />
                <ElProgress
                    class="!absolute top-1/2 right-1/2 translate-x-1/2 -translate-y-1/2"
                    type="circle"
                    :percentage="selectedProject.percent"
                    :status="selectedProject.percent === 100 ? 'success' : ''"
                >
                    <template #default="{ percentage }">
                        <span class="text-center text-2xl">{{ percentage }}%</span>
                        <div class="text-center text-md font-light">
                            {{ selectedProject.code }}
                        </div>
                    </template>
                </ElProgress>
            </div>
        </div>
    </ElCard>
</template>

<script lang="ts" setup>
import 'echarts'
import VChart from 'vue-echarts'
import type { TeamProject } from '~/types/team'

const { t } = useI18n()

const { team } = toRefs(useTeamStore())
const { getProjectsStatistic } = useTeam()
const isLoading = ref(false)

const data = ref([] as TeamProject[])
const selectedProject = ref({} as { code: string, percent: number })

const volumeData = computed(() => data.value.map(item => ({
    name: item.code,
    value: item.total_tasks_count,
    percent: (item.completed_tasks_count / item.total_tasks_count) * 100,
    total_tasks_count: item.total_tasks_count,
    completed_tasks_count: item.completed_tasks_count,
})))

const clustersChartOption = computed(() => ({
    title: { text: t('team.statistic.clusters'), left: 'center' },
    visualMap: {
        type: 'piecewise',
        showLabel: false,
        top: 'middle',
        min: 0,
        max: 2,
        left: 10,
        splitNumber: 3,
        color: ['#FFEA00', '#FF8800', '#E53E3E'],
    },
    grid: {
        left: 90,
    },
    tooltip: {
        trigger: 'item',
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        formatter: (item: any) => {
            const statisticData = item.data[2]
            return `
                <b>${statisticData.code}</b>
                <br/>
                ${t('team.statistic.cluster')}: ${item.marker} ${statisticData.cluster} 
                <br/>
                ${t('team.statistic.total')}: ${statisticData.total_tasks_count}
                <br/>
                ${t('team.statistic.completed')}: ${statisticData.completed_tasks_count}
                <br/>
                ${t('team.statistic.duration')}: ${statisticData.avg_task_duration.toFixed(2)}
            `
        },
    },
    xAxis: {
        type: 'value',
        name: t('team.statistic.completed'),
        nameLocation: 'middle',
        nameGap: 30,
    },
    yAxis: {
        type: 'value',
        name: t('team.statistic.duration'),
        nameLocation: 'middle',
        nameRotate: 90,
        nameGap: 30,
    },
    series: [
        {
            name: t('team.statistic.clusters'),
            type: 'scatter',
            symbolSize: 15,
            data: data.value.map(item => [
                item.completed_tasks_count,
                item.avg_task_duration,
                item,
                item.cluster,
            ]),
        },
    ],
}))

const volumeChartOption = computed(() => ({
    title: { text: t('team.statistic.volume'), left: 'center' },
    tooltip: {
        trigger: 'item',
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        formatter: (item: any) => {
            const { name, percent, total_tasks_count, completed_tasks_count } = item.data
            return `
                ${item.marker}<b>${name}</b>
                <br/>
                ${t('team.statistic.total')}: ${total_tasks_count}
                <br/>
                ${t('team.statistic.completed')}: ${completed_tasks_count}
                <br/>
                ${t('team.statistic.percent')}: ${percent.toFixed(2)}%
            `
        },
    },

    series: [
        {
            type: 'pie',
            radius: ['35%', '60%'],
            selectedMode: data.value.length > 1 ? 'single' : false,
            emphasis: { label: { show: true, fontSize: '30', fontWeight: 'bold' } },
            data: volumeData.value,
        },
    ],
}))

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const handleSelectChanged = ({ selected }: any) => {
    if (!data.value.length || data.value.length === 1) return
    if (!selected.length) {
        selectedProject.value = {
            percent: 0,
            code: '',
        }
        return
    }
    const [idx] = selected[0].dataIndex
    selectedProject.value = {
        percent: Math.round(volumeData.value[idx].percent),
        code: volumeData.value[idx].name,
    }
}

onMounted(async () => {
    try {
        isLoading.value = true
        data.value = await getProjectsStatistic(team.value.id)

        if (data.value.length === 1) {
            selectedProject.value = {
                percent: Math.round(volumeData.value[0].percent),
                code: volumeData.value[0].name,
            }
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
