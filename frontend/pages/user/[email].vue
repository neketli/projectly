<template>
    <div class="container mx-auto">
        <ElCard
            v-loading="isLoading"
            class="relative mx-auto max-w-xl"
        >
            <div class="flex flex-col gap-4">
                <div class="flex gap-4 items-center">
                    <UserAvatar
                        :file-name="user.meta?.avatar"
                        :size="64"
                    />

                    <div>
                        <p class="text-2xl font-medium">
                            {{ user.name }}
                            {{ user.surname }}
                        </p>

                        <p>
                            {{ user.email }}
                        </p>
                    </div>
                </div>

                <ElDivider />

                <ElDescriptions
                    :column="1"
                    border
                >
                    <ElDescriptionsItem>
                        <template #label>
                            <div class="flex gap-2 items-center">
                                <Icon name="mdi:calendar" />
                                {{ t('profile.form.birthday') }}
                            </div>
                        </template>

                        <div class="whitespace-pre-wrap">
                            {{
                                user.meta?.birthday
                                    ? dayjs.unix(user.meta?.birthday).format('DD MMM, YYYY')
                                    : t('profile.form.placeholder')
                            }}
                        </div>
                    </ElDescriptionsItem>

                    <ElDescriptionsItem>
                        <template #label>
                            <div class="flex gap-2 items-center">
                                <Icon name="mdi:map-marker" />
                                {{ t('profile.form.location') }}
                            </div>
                        </template>

                        <div class="whitespace-pre-wrap">
                            {{ user.meta?.location || t('profile.form.placeholder') }}
                        </div>
                    </ElDescriptionsItem>

                    <ElDescriptionsItem>
                        <template #label>
                            <div class="flex gap-2 items-center">
                                <Icon name="mdi:information" />
                                {{ t('profile.form.about') }}
                            </div>
                        </template>

                        <div class="whitespace-pre-wrap">
                            {{ user.meta?.about || t('profile.form.placeholder') }}
                        </div>
                    </ElDescriptionsItem>
                </ElDescriptions>
            </div>
        </ElCard>
    </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import type { User } from '~/types/user'

const route = useRoute()
const { t } = useI18n()

useHead({
    title: t('profile.title'),
})

const { getUserByEmail } = useUserActions()

const isLoading = ref(false)
const user = ref({} as User)

onMounted(async () => {
    isLoading.value = true
    user.value = await getUserByEmail(route.params.email)
    isLoading.value = false
})
</script>
