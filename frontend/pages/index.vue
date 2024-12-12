<template>
    <div class="dashboard container mx-auto ">
        <h1 class="text-3xl font-semibold mb-8">
            {{ $t('dashboard.title') }}
        </h1>

        <div class="flex flex-wrap gap-4">
            <UserTeams class="h-fit w-full max-w-md" />

            <ElCalendar
                v-model="date"
                class="max-w-2xl"
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
