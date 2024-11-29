<template>
    <div class="w-full h-screen flex items-center justify-center">
        <ElResult
            :title="$t('error.page.title')"
            :sub-title="$t('error.page.not_found')"
            icon="error"
        >
            <template #extra>
                <ElButton
                    class="ml-2"
                    @click="handleError"
                >
                    {{ $t('error.page.back') }}
                </ElButton>
            </template>
        </ElResult>
    </div>
</template>

<script setup lang="ts">
const { t } = useI18n()

useHead({
    title: t('common.error'),
})

const handleError = () => clearError({ redirect: '/' })
</script>
