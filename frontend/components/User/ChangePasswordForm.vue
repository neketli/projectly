<template>
    <ElForm
        ref="formElement"
        :rules="changePasswordRules"
        :model="form"
        :loading="isLoading"
        autocomplete="off"
        label-position="top"
    >
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

        <ElFormItem
            :label="$t('profile.form.confirm_password')"
            prop="confirmPassword"
            class="w-full"
        >
            <ElInput
                v-model="form.confirmPassword"
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
                    || !form.password
                    || form.password !== form.confirmPassword"
                type="primary"
                plain
                @click="handleSavePassword"
            >
                {{ $t('common.button.confirm') }}
            </ElButton>
            <ElButton
                type="danger"
                plain
                class="max-w-56"
                @click="emit('cancel')"
            >
                {{ $t('common.button.cancel') }}
            </ElButton>
        </div>
    </ElForm>
</template>

<script lang="ts" setup>
import type { FormInstance, FormRules } from 'element-plus'

const { changePassword } = useUserActions()
const { t } = useI18n()

const emit = defineEmits(['success', 'cancel'])

const isLoading = ref(false)

const formElement = ref<FormInstance>()
const form = ref({
    password: '',
    confirmPassword: '',
})
const changePasswordRules = reactive<FormRules<typeof form.value>>({
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
                if (value !== form.value.password) {
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

const handleSavePassword = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            savePassword()
        }
    })
}

const savePassword = async () => {
    isLoading.value = true
    try {
        await changePassword(form.value.confirmPassword)
        emit('success')
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

</style>
