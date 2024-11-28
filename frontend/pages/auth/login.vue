<template>
    <section
        class="w-full min-h-[92vh] flex flex-col gap-8 justify-center bg-sky-50 dark:bg-slate-900 items-center "
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
                {{ $t('auth.login.title') }}
            </h1>
            <ElFormItem
                :label="$t('auth.login.form.email')"
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
                :label="$t('auth.login.form.password')"
                prop="password"
                class="w-full"
            >
                <ElInput
                    v-model="authForm.password"
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
                    {{ $t('auth.login.form.login') }}
                    <Icon
                        class="ml-2"
                        name="mdi:login"
                    />
                </ElButton>

                <ElButton
                    :disabled="state.isLoading"
                    :loading="state.isLoading"
                    class="w-full"
                    @click="navigateTo('/auth/register')"
                >
                    {{ $t('auth.login.form.register') }}
                    <Icon
                        class="ml-2"
                        name="mdi:account-plus"
                    />
                </ElButton>
            </div>
        </ElForm>
    </section>
</template>

<script setup lang="ts">
import type { FormRules, FormInstance } from 'element-plus'
import { validators } from '~/utils/validators'

const { t } = useI18n()

useHead({
    title: t('auth.login.title'),
})
definePageMeta({
    isPublic: true,
    layout: 'empty',
    pageTransition: { name: 'slide' },
})

const authStore = useAuthStore()

const state = reactive({
    isLoading: false,
})

const formElement = ref<FormInstance>()
const authForm = ref({
    email: '',
    password: '',
})
const rules = reactive<FormRules<{ email: string, password: string }>>({
    email: [
        validators.required,
    ],
    password: [
        validators.required,
    ],
})

const auth = async () => {
    state.isLoading = true
    try {
        await authStore.login({
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
            await auth()
        }
    })
}
</script>
