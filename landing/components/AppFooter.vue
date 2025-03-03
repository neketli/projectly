<template>
    <footer class="bg-gray-800">
        <div class="max-w-7xl mx-auto py-12 px-4 overflow-hidden sm:px-6 lg:px-8">
            <div class="mt-4 flex justify-center space-x-6">
                <a
                    v-for="item in social"
                    :key="item.name"
                    :href="item.href"
                    class="text-gray-400 hover:text-gray-300"
                >
                    <span class="sr-only">{{ item.name }}</span>
                    <Icon
                        :name="item.icon"
                        class="h-6 w-6"
                        aria-hidden="true"
                    />
                </a>
            </div>
            <p class="mt-4 text-center text-base text-gray-400">
                {{ $t('footer.copyright', { year: new Date().getFullYear() }) }}
            </p>

            <div class="mt-6 flex gap-4 flex-wrap justify-center items-center">
                <a
                    v-for="item in documents"
                    :key="item.name"
                    :href="item.href"
                    target="_blank"
                    nofollow
                    noreferrer
                    class="text-sm text-gray-400 hover:text-gray-300"
                >
                    {{ item.name }}
                </a>
            </div>
        </div>
    </footer>
</template>

<script setup>
const { t, locale } = useI18n()

const BASE_URL = `https://projectly.ru/media/documents`
const documents = computed(() => [
    {
        name: t('footer.privacy'),
        href: `${BASE_URL}/privacy_policy(${locale.value}).pdf`,
    },
    {
        name: t('footer.terms_of_service'),
        href: `${BASE_URL}/terms_of_service(${locale.value}).pdf`,
    },
    {
        name: t('footer.terms_of_use'),
        href: `${BASE_URL}/terms_of_use(${locale.value}).pdf`,
    },
    {
        name: t('footer.disclaimer'),
        href: `${BASE_URL}/disclaimer(${locale.value}).pdf`,
    },
])

const social = [
    {
        name: 'GitHub',
        href: 'https://github.com/neketli/projectly',
        icon: 'mdi:github',
    },
    {
        name: 'Telegram',
        href: 'https://t.me/projectlyru',
        icon: 'mdi:telegram',
    },
    {
        name: 'Telegram support',
        href: 'https://t.me/projectly_support_bot',
        icon: 'mdi:face-agent',
    },
]
</script>
