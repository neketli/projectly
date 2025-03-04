<template>
    <section
        class="w-full min-h-[88vh] flex flex-col gap-8 justify-center bg-sky-50 dark:bg-slate-900 items-center px-4"
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

            <div class="flex  flex-col sm:flex-row gap-4 mt-2">
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
                    class="w-full !ml-0"
                    @click="navigateTo('/auth/register')"
                >
                    {{ $t('auth.login.form.register') }}
                    <Icon
                        class="ml-2"
                        name="mdi:account-plus"
                    />
                </ElButton>
            </div>

            <ElDivider
                class="form-divider my-4"
                content-position="center"
            >
                <span class="p-2 bg-white dark:bg-slate-800">
                    {{ $t('auth.login.form.or') }}
                </span>
            </ElDivider>

            <ElButton
                class="w-full"
                tag="a"
                :href="authLinks.google"
            >
                <Icon
                    class="mr-2"
                    name="fa-brands:google"
                />
                {{ $t('auth.login.form.with_google') }}
            </ElButton>

            <ElButton
                class="w-full !ml-0 mt-4"
                tag="a"
                :href="authLinks.yandex"
            >
                <Icon
                    class="mr-2"
                    name="fa-brands:yandex"
                />
                {{ $t('auth.login.form.with_yandex') }}
            </ElButton>

            <ElAlert
                type="info"
                show-icon
                :closable="false"
                class="!mt-4 text-blue"
            >
                <div
                    v-html="$t('auth.terms', {
                        type: $t('auth.login.form.login'),
                        ...termsLinks,
                    })"
                />
            </ElAlert>
        </ElForm>
    </section>
</template>

<script setup lang="ts">
import type { FormRules, FormInstance } from 'element-plus'

const { t, locale } = useI18n()
const config = useRuntimeConfig()

useHead({
    title: t('auth.login.title'),
})
definePageMeta({
    isPublic: true,
    layout: 'empty',
    pageTransition: { name: 'slide' },
})

const authStore = useAuthStore()
const validators = useValidator()

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

const authLinks = computed(() => {
    const BASE_URL = `${config.public.API_HOST}/api/v1/auth`

    return {
        google: `${BASE_URL}/google`,
        yandex: `${BASE_URL}/yandex`,
    }
})
const termsLinks = computed(() => {
    const baseUrl = `${config.public.S3_HOST}/media/documents`
    const shortLocale = locale.value.split('-')[0]

    return {
        terms: `${baseUrl}/terms_of_service(${shortLocale}).pdf`,
        privacy: `${baseUrl}/privacy_policy(${shortLocale}).pdf`,
    }
})

const auth = async () => {
    state.isLoading = true
    try {
        await authStore.login({
            email: authForm.value.email,
            password: authForm.value.password,
        })
        navigateTo('/')
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

<style>
.form-divider .el-divider__text {
    background-color: transparent !important;
}
.auth-link {
    color: #2563eb;
    text-decoration: none;
}
.auth-link::after {
    content: '➚';
}
</style>
