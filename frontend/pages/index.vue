<template>
    <div class="dashboard container mx-auto ">
        <h1 class="text-3xl font-semibold mb-8">
            {{ $t('dashboard.title') }}
        </h1>

        <div class="grid grid-cols-3">
            <UserTeams />
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    layout: 'default',
})

onMounted(() => {
    useTeamStore().$reset()
    useProjectStore().$reset()
})
</script>
