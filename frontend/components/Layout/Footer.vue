<template>
    <ElFooter class="flex items-center gap-6 justify-center mt-auto border-t dark:border-gray-700 text-gray-400">
        <div class="flex justify-center gap-4">
            <a
                v-for="item in social"
                :key="item.name"
                :href="item.href"
                class="h-5 w-5 hover:text-gray-500 dark:hover:text-gray-300"
            >
                <span class="sr-only">{{ item.name }}</span>
                <Icon
                    :name="item.icon"
                    class="h-5 w-5"
                    aria-hidden="true"
                />
            </a>
        </div>

        <p class="text-center text-base">
            {{ $t('footer.copyright', { year: new Date().getFullYear() }) }}
        </p>
    </ElFooter>
</template>

<script setup>
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
