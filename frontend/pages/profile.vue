<template>
    <div class="container">
        <h1 class="text-4xl pb-4 mx-auto max-w-xl">
            {{ $t('profile.title') }}
        </h1>

        <ElCard class="relative mx-auto max-w-xl">
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
                    <UserAvatar :size="64" />

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

                <div v-if="!isChangePassword && !isEdit">
                    <ElButton
                        plain
                        @click="handleChangePassword"
                    >
                        {{ $t('profile.form.change_password') }}
                    </ElButton>

                    <ElButton
                        v-if="getUserInfo?.meta?.avatar"
                        type="warning"
                        plain
                        @click="handleRemoveAvatar"
                    >
                        {{ $t('profile.form.remove_avatar') }}
                    </ElButton>
                </div>
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
    </div>
</template>

<script setup lang="ts">
const { t } = useI18n()

useHead({
    title: t('profile.title'),
})

const authStore = useAuthStore()

const { getUserInfo } = toRefs(authStore)

const { removeAvatar } = useUserActions()

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
</script>
