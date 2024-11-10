<template>
    <section
        class="w-full min-h-screen flex flex-col gap-8 justify-center items-center bg-sky-50"
    >
        <ElForm
            ref="formElement"
            :rules="rules"
            :model="authForm"
            class="max-w-md w-full container bg-white p-8 shadow-lg rounded-lg"
            label-position="top"
        >
            <h1 class="text-xl text-center font-medium mb-8">
                Регистрация
            </h1>
            <ElFormItem
                label="Имя"
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
                label="Фамилия"
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
                label="Адрес электронной почты"
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
                label="Пароль"
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
                label="Повторите пароль"
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

            <div class="flex mt-2">
                <ElButton
                    :disabled="state.isLoading"
                    :loading="state.isLoading"
                    class="w-full"
                    type="primary"
                    @click="check"
                >
                    Зарегистрироваться
                    <Icon
                        class="ml-2"
                        name="mdi:account-plus"
                    />
                </ElButton>

                <ElButton
                    :disabled="state.isLoading"
                    :loading="state.isLoading"
                    class="w-full"
                    @click="navigateTo('/auth/login')"
                >
                    Уже есть аккаунт
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
import { validators } from '~/utils/validators'

useHead({
    title: 'Регистрация',
})
definePageMeta({
    isPublic: true,
    layout: 'empty',
    pageTransition: { name: 'slide', mode: 'out-in' },
})

const authStore = useAuthStore()

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

        return ElNotification({
            title: 'Ошибка',
            message: error.message,
            type: 'error',
        })
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
