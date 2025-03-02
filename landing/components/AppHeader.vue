<template>
    <header class="bg-white shadow-sm sticky top-0 z-50">
        <nav
            class="max-w-7xl mx-auto px-4 sm:px-6 xl:px-8"
            aria-label="Top"
        >
            <div class="w-full py-6 flex items-center justify-between border-b border-gray-200 xl:border-none">
                <div class="flex items-center">
                    <a href="#">
                        <span class="flex gap-2 text-2xl text-gray-900 font-['Helvetica']">
                            Projectly

                            <img
                                class="h-8 w-auto"
                                src="@/assets/icons/logo.svg"
                                alt="Projectly logo"
                            >
                        </span>

                    </a>
                    <div class="hidden ml-10 space-x-8 xl:block">
                        <a
                            v-for="link in navigation"
                            :key="link.name"
                            :href="link.href"
                            class="text-base font-medium text-gray-600 transition hover:text-gray-900"
                        >
                            {{ link.name }}
                        </a>
                    </div>
                </div>

                <div class="hidden ml-auto xl:flex gap-4 ">
                    <button
                        :class="{
                            'text-gray-600 cursor-default': $i18n.locale === 'ru',
                            'text-primary-600 transition hover:text-primary-700 cursor-pointer': $i18n.locale !== 'ru',
                        }"
                        @click="setLocale('ru')"
                    >
                        {{ $t('header.ru') }}
                    </button>
                    <span>
                        |
                    </span>
                    <button
                        :class="{
                            'text-gray-600 cursor-default': $i18n.locale === 'en',
                            'text-primary-600 transition hover:text-primary-700 cursor-pointer': $i18n.locale !== 'en',
                        }"
                        @click="setLocale('en')"
                    >
                        {{ $t('header.en') }}
                    </button>
                </div>
                <div class="ml-4 space-x-4">
                    <a
                        href="https://projectly.ru/my/auth/login"
                        class="inline-block bg-white py-2 px-4 border border-transparent rounded-md text-base font-medium text-primary-600 hover:bg-gray-50"
                    >
                        {{ $t('header.login') }}
                    </a>
                    <a
                        href="https://projectly.ru/my/auth/register"

                        class="inline-block bg-primary-600 py-2 px-4 border border-transparent rounded-md text-base font-medium text-white hover:bg-primary-700"
                    >
                        {{ $t('header.register') }}
                    </a>
                </div>
            </div>

            <div class="py-4 flex flex-wrap justify-center gap-4 xl:hidden">
                <button
                    :class="{
                        'text-gray-600 cursor-default': $i18n.locale === 'ru',
                        'text-primary-600 transition hover:text-primary-700 cursor-pointer': $i18n.locale !== 'ru',
                    }"
                    @click="setLocale('ru')"
                >
                    {{ $t('header.ru') }}
                </button>
                <span>
                    |
                </span>
                <button
                    :class="{
                        'text-gray-600 cursor-default': $i18n.locale === 'en',
                        'text-primary-600 transition hover:text-primary-700 cursor-pointer': $i18n.locale !== 'en',
                    }"
                    @click="setLocale('en')"
                >
                    {{ $t('header.en') }}
                </button>
            </div>

            <div class="py-4 flex flex-wrap justify-center space-x-6 xl:hidden">
                <a
                    v-for="link in navigation"
                    :key="link.name"
                    :href="link.href"
                    class="text-base font-medium text-gray-600 hover:text-gray-900"
                >
                    {{ link.name }}
                </a>
            </div>
        </nav>
    </header>
</template>

<script setup>
const { t, setLocale } = useI18n()

const navigation = computed(() => [
    { name: t('header.nav.features'), href: '#features' },
    { name: t('header.nav.how_it_works'), href: '#how-it-works' },
    { name: t('header.nav.demo'), href: '#demo' },
    { name: t('header.nav.pricing'), href: '#pricing' },
    { name: t('header.nav.faq'), href: '#faq' },
])
</script>
