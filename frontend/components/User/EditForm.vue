<template>
    <div>
        <div class="absolute right-4 top-4">
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
                @click="emit('cancel')"
            >
                <Icon name="mdi:cancel" />
            </ElButton>
        </div>

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
            @submit.prevent
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
</template>

<script  setup lang="ts">
import type { FormInstance, FormRules, UploadProps, UploadUserFile } from 'element-plus'

const emit = defineEmits(['success', 'cancel'])

const { t } = useI18n()
const { getUserInfo } = toRefs(useAuthStore())
const { updateUserInfo, uploadAvatar } = useUserActions()

const validators = useValidator()

const files: Ref<UploadUserFile[]> = ref([])
const imageUrl = ref('')
const isLoading = ref(false)

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

const saveChanges = async () => {
    isLoading.value = true
    try {
        if (files.value.length) {
            await uploadAvatar(files.value[0].raw as File)
        }
        await updateUserInfo(form.value)
        emit('success')
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
