<template>
    <header class="bg-white shadow-sm sticky top-0 z-50">
        <nav
            class="max-w-7xl mx-auto px-4 sm:px-6 xl:px-8"
            aria-label="Top"
        >
            <div class="w-full py-6 flex items-center justify-between  border-gray-200">
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
                    <div class="hidden ml-10 space-x-8 lg:block">
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

                <LangSwitch class="hidden ml-auto lg:flex" />

                <div class="hidden lg:block ml-4 space-x-4">
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

                <button
                    class="lg:hidden inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500"
                    @click="menu = !menu"
                >
                    <Icon :name="menu ? 'mdi:close' : 'mdi:menu' " />
                </button>
            </div>

            <Transition
                enter-active-class="transition ease-out duration-100"
                enter-from-class="transform opacity-0 scale-95 -translate-y-[80px]"
                enter-to-class="transform opacity-100 scale-100"
                leave-active-class="transition ease-in duration-75"
                leave-from-class="transform opacity-100 scale-100"
                leave-to-class="transform opacity-0 scale-95 -translate-y-[80px]"
            >
                <div
                    v-show="menu"
                    class="lg:hidden absolute top-[80px] right-0 w-full bg-white border border-gray-200 rounded-md shadow-md"
                >
                    <div class="py-4 flex flex-wrap justify-center gap-4">
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

                    <LangSwitch />

                    <div
                        class="py-4 flex flex-wrap justify-center space-x-6"
                        @click="menu = false"
                    >
                        <a
                            v-for="link in navigation"
                            :key="link.name"
                            :href="link.href"
                            class="text-base font-medium text-gray-600 hover:text-gray-900"
                        >
                            {{ link.name }}
                        </a>
                    </div>
                </div>
            </Transition>
        </nav>
    </header>
</template>

<script setup>
import LangSwitch from './ui/LangSwitch.vue'

const { t } = useI18n()

const menu = ref(false)

const navigation = computed(() => [
    { name: t('header.nav.features'), href: '#features' },
    { name: t('header.nav.how_it_works'), href: '#how-it-works' },
    { name: t('header.nav.demo'), href: '#demo' },
    { name: t('header.nav.pricing'), href: '#pricing' },
    { name: t('header.nav.faq'), href: '#faq' },
])
</script>
