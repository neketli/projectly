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
                class="flex gap-4 items-center"
            >
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

                    <ElFormItem
                        :label="$t('profile.form.password')"
                        prop="password"
                        class="w-full"
                    >
                        <ElInput
                            v-model="form.password"
                            type="password"
                            autocomplete="off"
                            show-password
                        >
                            <template #prefix>
                                <Icon name="mdi:lock" />
                            </template>
                        </ElInput>
                    </ElFormItem>

                    <ElAlert
                        type="info"
                        show-icon
                        :closable="false"
                    >
                        {{ $t('profile.form.info') }}
                    </ElAlert>
                </ElForm>
            </div>
        </ElCard>
    </div>
</template>

<script setup lang="ts">
import type { FormInstance, FormRules, UploadProps, UploadUserFile } from 'element-plus'

const { t } = useI18n()

const authStore = useAuthStore()

const { getUserInfo } = toRefs(authStore)

const { updateUserInfo, uploadAvatar } = useUserActions()

const isEdit = ref(false)
const isLoading = ref(false)

const files: Ref<UploadUserFile[]> = ref([])
const imageUrl = ref('')

const formElement = ref<FormInstance>()
const form = ref({
    name: getUserInfo.value.name,
    surname: getUserInfo.value.surname,
    email: getUserInfo.value.email,
    password: '',
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
    password: [],
})

const handleEdit = () => {
    files.value = []
    isEdit.value = true
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

const handleCancel = () => {
    isEdit.value = false
}

const handleChange: UploadProps['onChange'] = (uploadFile) => {
    const MAX_SIZE_MB = 2
    if (uploadFile.size && (uploadFile.size / 1024 / 1024 > MAX_SIZE_MB)) {
        ElMessage.error(t('profile.image.size_error', { x: MAX_SIZE_MB }))
        imageUrl.value = ''
        return false
    }
    files.value = [uploadFile]
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
    return true
}

const saveChanges = async () => {
    isLoading.value = true
    if (files.value.length) {
        await uploadAvatar(files.value[0].raw as File)
    }

    await updateUserInfo(form.value)
    await authStore.refresh()

    isEdit.value = false
    isLoading.value = false
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
