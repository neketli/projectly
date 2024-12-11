<template>
    <section
        class="w-full min-h-[92vh] flex flex-col gap-8 justify-center items-center bg-sky-50 dark:bg-slate-900 px-4"
    >
        <ElForm
            ref="formElement"
            :rules="rules"
            :model="authForm"
            class="max-w-md w-full container bg-white dark:bg-slate-800 dark:shadow-none dark:ring-2 dark:ring-slate-600 p-8 shadow-lg rounded-lg"
            label-position="top"
            @submit.prevent
        >
            <h1 class="text-xl text-center font-medium mb-8">
                {{ $t('auth.register.title') }}
            </h1>
            <ElFormItem
                :label="$t('auth.register.form.name')"
                prop="name"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.name"
                >
                    <template #prefix>
                        <Icon name="mdi:account" />
                    </template>
                </ElInput>
            </ElFormItem>

            <ElFormItem
                :label="$t('auth.register.form.surname')"
                prop="surname"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.surname"
                >
                    <template #prefix>
                        <Icon name="mdi:account" />
                    </template>
                </ElInput>
            </ElFormItem>

            <ElFormItem
                :label="$t('auth.register.form.email')"
                prop="email"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.email"
                    type="email"
                >
                    <template #prefix>
                        <Icon name="mdi:email" />
                    </template>
                </ElInput>
            </ElFormItem>

            <ElFormItem
                :label="$t('auth.register.form.password')"
                prop="password"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.password"
                    type="password"
                    show-password
                >
                    <template #prefix>
                        <Icon name="mdi:lock" />
                    </template>
                </ElInput>
            </ElFormItem>

            <ElFormItem
                :label="$t('auth.register.form.confirm_password')"
                prop="confirmPassword"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.confirmPassword"
                    type="password"
                    show-password
                    @keyup.enter="check"
                >
                    <template #prefix>
                        <Icon name="mdi:lock" />
                    </template>
                </ElInput>
            </ElFormItem>

            <div class="flex flex-col sm:flex-row gap-4 mt-2">
                <ElButton
                    :disabled="state.isLoading"
                    :loading="state.isLoading"
                    class="w-full"
                    type="primary"
                    @click="check"
                >
                    {{ $t('auth.register.form.register') }}
                    <Icon
                        class="ml-2"
                        name="mdi:account-plus"
                    />
                </ElButton>

                <ElButton
                    :disabled="state.isLoading"
                    :loading="state.isLoading"
                    class="w-full !ml-0"
                    @click="navigateTo('/auth/login')"
                >
                    {{ $t('auth.register.form.login') }}

                    <Icon
                        class="ml-2"
                        name="mdi:login"
                    />
                </ElButton>
            </div>
        </ElForm>
    </section>
</template>

<script setup lang="ts">
import type { FormRules, FormInstance } from 'element-plus'

const { t } = useI18n()

useHead({
    title: t('auth.register.title'),
})
definePageMeta({
    isPublic: true,
    layout: 'empty',
    pageTransition: { name: 'slide', mode: 'out-in' },
})

const authStore = useAuthStore()
const validators = useValidator()

const state = reactive({
    isLoading: false,
})

const formElement = ref<FormInstance>()
const authForm = ref({
    name: '',
    surname: '',
    email: '',
    password: '',
    confirmPassword: '',
})
const rules = reactive<FormRules<typeof authForm.value>>({
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
                if (value !== authForm.value.password) {
                    callback(new Error(t('auth.register.form.error.confirm_password')))
                }
                else {
                    callback()
                }
            },
            trigger: 'blur',
        },
    ],
})

const register = async () => {
    state.isLoading = true
    try {
        await authStore.register({
            name: authForm.value.name,
            surname: authForm.value.surname,
            email: authForm.value.email,
            password: authForm.value.password,
        })
        navigateTo('/', { external: true })
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        state.isLoading = false
    }
}

const check = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            await register()
        }
    })
}
</script>
