<template>
    <div class="container mx-auto">
        <h1 class="text-4xl pb-4 mx-auto max-w-xl">
            {{ $t('profile.title') }}
        </h1>

        <ElCard
            v-loading="isLoading"
            class="relative mx-auto max-w-xl"
        >
            <div class="absolute top-4 right-4">
                <ElButton
                    v-if="!isEdit"
                    circle
                    @click="handleEdit"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>
            </div>

            <div
                v-if="!isEdit"
                class="flex flex-col gap-4"
            >
                <div class="flex gap-4 items-center">
                    <UserAvatar
                        :file-name="getUserInfo.meta?.avatar"
                        :size="64"
                    />

                    <div>
                        <p class="text-2xl font-medium">
                            {{ getUserInfo.name }}
                            {{ getUserInfo.surname }}
                        </p>

                        <p>
                            {{ getUserInfo.email }}
                        </p>
                    </div>
                </div>

                <div
                    v-if="!isChangePassword && !isEdit"
                    class="flex flex-col sm:flex-row gap-2"
                >
                    <ElButton
                        class="w-full sm:w-fit"
                        plain
                        @click="handleChangePassword"
                    >
                        {{ $t('profile.form.change_password') }}
                    </ElButton>

                    <ElButton
                        v-if="getUserInfo?.meta?.avatar"
                        class="w-full !ml-0 sm:w-fit"
                        type="warning"
                        plain
                        @click="handleRemoveAvatar"
                    >
                        {{ $t('profile.form.remove_avatar') }}
                    </ElButton>
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
                                getUserInfo.meta?.birthday
                                    ? dayjs.unix(getUserInfo.meta?.birthday).format('DD MMM, YYYY')
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
                            {{ getUserInfo.meta?.location || t('profile.form.placeholder') }}
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
                            {{ getUserInfo.meta?.about || t('profile.form.placeholder') }}
                        </div>
                    </ElDescriptionsItem>
                </ElDescriptions>
            </div>

            <UserEditForm
                v-else
                @success="handleSave"
                @cancel="handleCancel"
            />

            <UserChangePasswordForm
                v-if="isChangePassword"
                class="pt-4"
                @success="handleSave"
                @cancel="handleCancel"
            />
        </ElCard>

        <UserTerms class="mt-4 mx-auto max-w-xl" />
    </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'

const { t } = useI18n()

useHead({
    title: t('profile.title'),
})

const authStore = useAuthStore()
const { getUserInfo } = toRefs(authStore)

const { getUserByEmail, removeAvatar } = useUserActions()

const isEdit = ref(false)
const isChangePassword = ref(false)
const isLoading = ref(false)

const handleEdit = () => {
    handleCancel()
    isEdit.value = true
}

const handleCancel = () => {
    isEdit.value = false
    isChangePassword.value = false
}

const handleChangePassword = () => {
    handleCancel()
    isChangePassword.value = true
}

const handleRemoveAvatar = async () => {
    try {
        await removeAvatar()
        await authStore.refresh()
        ElMessage.success(t('profile.form.success.image'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleSave = async () => {
    isLoading.value = true
    try {
        await authStore.refresh()
        authStore.user = await getUserByEmail(getUserInfo.value.email)

        handleCancel()
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

onMounted(async () => {
    isLoading.value = true
    authStore.user = await getUserByEmail(getUserInfo.value.email)
    isLoading.value = false
})
</script>
