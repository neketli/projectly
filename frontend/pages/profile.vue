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

                <template v-else>
                    <ElButton
                        circle
                        :disabled="isLoading"
                        @click="handleSave"
                    >
                        <Icon name="mdi:content-save" />
                    </ElButton>
                    <ElButton
                        type="danger"
                        circle
                        @click="handleCancel"
                    >
                        <Icon name="mdi:cancel" />
                    </ElButton>
                </template>
            </div>

            <div
                v-if="!isEdit"
                class="flex flex-col gap-4"
            >
                <div class="flex gap-4 items-center">
                    <LayoutUserAvatar :size="64" />

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
                        v-if="getUserInfo.meta.avatar"
                        type="warning"
                        plain
                        @click="handleRemoveAvatar"
                    >
                        {{ $t('profile.form.remove_avatar') }}
                    </ElButton>
                </div>
            </div>

            <div v-else>
                <div class="flex items-center gap-2">
                    <ElUpload
                        v-loading="isLoading"
                        accept=".png, .jpg, .jpeg"
                        :auto-upload="false"
                        drag
                        :show-file-list="false"
                        class="w-16 avatar-uploader"
                        :on-change="handleChange"
                    >
                        <img
                            v-if="imageUrl"
                            :src="imageUrl"
                            class="w-16 h-16 object-cover"
                        >

                        <template v-else>
                            <Icon
                                class="w-8 h-8 absolute top-4 left-4"
                                name="mdi:cloud-upload-outline"
                            />
                        </template>
                    </ElUpload>

                    <span>
                        Загрузите аватар
                    </span>
                </div>

                <ElForm
                    ref="formElement"
                    :rules="rules"
                    :model="form"
                    :loading="isLoading"
                    class="pt-4"
                    autocomplete="off"
                    label-position="top"
                >
                    <ElFormItem
                        :label="$t('profile.form.name')"
                        prop="name"
                        class="w-full"
                    >
                        <ElInput
                            v-model="form.name"
                        >
                            <template #prefix>
                                <Icon name="mdi:account" />
                            </template>
                        </ElInput>
                    </ElFormItem>

                    <ElFormItem
                        :label="$t('profile.form.surname')"
                        prop="surname"
                        class="w-full"
                    >
                        <ElInput
                            v-model="form.surname"
                        >
                            <template #prefix>
                                <Icon name="mdi:account" />
                            </template>
                        </ElInput>
                    </ElFormItem>

                    <ElFormItem
                        :label="$t('profile.form.email')"
                        prop="email"
                        class="w-full"
                    >
                        <ElInput
                            v-model="form.email"
                            type="email"
                        >
                            <template #prefix>
                                <Icon name="mdi:email" />
                            </template>
                        </ElInput>
                    </ElFormItem>
                </ElForm>
            </div>

            <ElForm
                v-if="isChangePassword"
                ref="changePasswordFormElement"
                :rules="changePasswordRules"
                :model="changePasswordForm"
                :loading="isLoading"
                class="pt-4"
                autocomplete="off"
                label-position="top"
            >
                <ElFormItem
                    :label="$t('profile.form.password')"
                    prop="password"
                    class="w-full"
                >
                    <ElInput
                        v-model="changePasswordForm.password"
                        type="password"
                        autocomplete="off"
                        show-password
                    >
                        <template #prefix>
                            <Icon name="mdi:lock" />
                        </template>
                    </ElInput>
                </ElFormItem>

                <ElFormItem
                    :label="$t('profile.form.confirm_password')"
                    prop="confirmPassword"
                    class="w-full"
                >
                    <ElInput
                        v-model="changePasswordForm.confirmPassword"
                        type="password"
                        autocomplete="off"
                        show-password
                    >
                        <template #prefix>
                            <Icon name="mdi:lock" />
                        </template>
                    </ElInput>
                </ElFormItem>

                <div>
                    <ElButton
                        :disabled="isLoading
                            || !changePasswordForm.password
                            || changePasswordForm.password !== changePasswordForm.confirmPassword"
                        type="primary"
                        plain
                        @click="handleSavePassword"
                    >
                        {{ $t('profile.form.confirm') }}
                    </ElButton>
                    <ElButton
                        type="danger"
                        plain
                        class="max-w-56"
                        @click="handleCancel"
                    >
                        {{ $t('profile.form.cancel') }}
                    </ElButton>
                </div>
            </ElForm>
        </ElCard>
    </div>
</template>

<script setup lang="ts">
import type { FormInstance, FormRules, UploadProps, UploadUserFile } from 'element-plus'

const { t } = useI18n()

const authStore = useAuthStore()

const { getUserInfo } = toRefs(authStore)

const { updateUserInfo, changePassword, uploadAvatar, removeAvatar } = useUserActions()

const isEdit = ref(false)
const isChangePassword = ref(false)
const isLoading = ref(false)

const files: Ref<UploadUserFile[]> = ref([])
const imageUrl = ref('')

const formElement = ref<FormInstance>()
const form = ref({
    name: getUserInfo.value.name,
    surname: getUserInfo.value.surname,
    email: getUserInfo.value.email,
})
const rules = reactive<FormRules<typeof form.value>>({
    name: [
        validators.required,
        validators.len(),
    ],
    surname: [
        validators.required,
        validators.len(),
    ],
    email: [
        validators.required,
        validators.email,
        validators.len(),
    ],
})

const changePasswordFormElement = ref<FormInstance>()
const changePasswordForm = ref({
    password: '',
    confirmPassword: '',
})
const changePasswordRules = reactive<FormRules<typeof changePasswordForm.value>>({
    password: [
        validators.required,
    ],
    confirmPassword: [
        validators.required,
        {
            validator: (
                _: unknown,
                value: unknown,
                callback: (error?: Error) => unknown,
            ) => {
                if (value !== changePasswordForm.value.password) {
                    callback(new Error('Пароли не совпадают'))
                }
                else {
                    callback()
                }
            },
            trigger: 'blur',
        },
    ],
})

const handleEdit = () => {
    handleCancel()
    isEdit.value = true
}

const handleCancel = () => {
    imageUrl.value = ''
    files.value = []

    changePasswordForm.value.password = ''
    changePasswordForm.value.confirmPassword = ''

    form.value.name = getUserInfo.value.name
    form.value.surname = getUserInfo.value.surname
    form.value.email = getUserInfo.value.email

    isEdit.value = false
    isChangePassword.value = false
}

const handleChangePassword = () => {
    handleCancel()
    isChangePassword.value = true
}

const handleSavePassword = async () => {
    if (!changePasswordFormElement.value) {
        return
    }
    await changePasswordFormElement.value.validate(async (valid) => {
        if (valid) {
            savePassword()
        }
    })
    isChangePassword.value = false
}

const handleSave = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            saveChanges()
        }
    })
}

const handleChange: UploadProps['onChange'] = (uploadFile) => {
    const MAX_SIZE_MB = 2
    if (uploadFile.size && (uploadFile.size / 1024 / 1024 > MAX_SIZE_MB)) {
        ElMessage.error(t('profile.form.error.image_size', { x: MAX_SIZE_MB }))
        imageUrl.value = ''
        return false
    }
    files.value = [uploadFile]
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
    return true
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

const saveChanges = async () => {
    isLoading.value = true
    try {
        if (files.value.length) {
            await uploadAvatar(files.value[0].raw as File)
        }
        await updateUserInfo(form.value)
        await authStore.refresh()
        handleCancel()
        ElMessage.success(t('profile.form.success.profile'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

const savePassword = async () => {
    isLoading.value = true
    try {
        await changePassword(changePasswordForm.value.confirmPassword)
        await authStore.refresh()
        handleCancel()
        ElMessage.success(t('profile.form.success.password'))
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

<style>
.avatar-uploader .el-upload-dragger {
    border-radius: 50%;
    width: 64px !important;
    height: 64px!important;
    padding: 0!important;
    position: relative;
}
</style>
