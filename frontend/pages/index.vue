<template>
    <div class="dashboard container mx-auto ">
        <h1 class="text-3xl font-semibold mb-8">
            {{ $t('dashboard.title') }}
        </h1>

        <div class="grid grid-cols-3 gap-4">
            <UserTeams class="h-fit" />

            <ElCalendar
                v-model="date"
                class="col-start-2 col-end-4"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'

definePageMeta({
    layout: 'default',
})

const date = ref(dayjs().toDate())

onMounted(() => {
    useTeamStore().$reset()
    useProjectStore().$reset()
})
</script>
