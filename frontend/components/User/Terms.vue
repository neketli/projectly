<template>
    <ElCard>
        <div class="grid grid-cols-2 gap-2">
            <a
                v-for="item in documents"
                :key="item.name"
                :href="item.href"
                target="_blank"
                nofollow
                noreferrer
                class="text-blue-500 hover:text-blue-600 dark:text-blue-400 dark:hover:text-blue-300"
            >
                {{ item.name }}
                <Icon name="mdi:open-in-new" />
            </a>
        </div>
    </ElCard>
</template>

<script lang="ts" setup>
const { t, locale } = useI18n()
const config = useRuntimeConfig()

const baseUrl = `${config.public.S3_HOST}/media/documents`

const shortLocale = computed(() => locale.value.split('-')[0])
const documents = computed(() => [
    {
        name: t('footer.privacy'),
        href: `${baseUrl}/privacy_policy(${shortLocale.value}).pdf`,
    },
    {
        name: t('footer.terms_of_service'),
        href: `${baseUrl}/terms_of_service(${shortLocale.value}).pdf`,
    },
    {
        name: t('footer.terms_of_use'),
        href: `${baseUrl}/terms_of_use(${shortLocale.value}).pdf`,
    },
    {
        name: t('footer.disclaimer'),
        href: `${baseUrl}/disclaimer(${shortLocale.value}).pdf`,
    },
])
</script>
