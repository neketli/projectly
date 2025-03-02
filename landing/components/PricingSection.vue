<template>
    <section
        id="pricing"
        class="bg-white py-16"
    >
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="text-center">
                <h2 class="text-base font-semibold text-primary-600 tracking-wide uppercase">
                    {{ $t('pricing.section') }}
                </h2>
                <p class="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight">
                    {{ $t('pricing.title') }}
                </p>
                <p class="max-w-xl mt-5 mx-auto text-xl text-gray-500">
                    {{ $t('pricing.subtitle') }}
                </p>
            </div>

            <div class="mt-12 space-y-4 sm:mt-16 sm:space-y-0 sm:grid sm:gap-6 lg:max-w-4xl lg:mx-auto lg:grid-cols-3 xl:max-w-none ">
                <div
                    v-for="tier in tiers"
                    :key="tier.name"
                    class="border border-gray-200 rounded-lg shadow-sm divide-y divide-gray-200"
                >
                    <div class="p-6">
                        <h2 class="text-lg leading-6 font-medium text-gray-900">
                            {{ tier.name }}
                        </h2>
                        <p class="mt-4 text-sm text-gray-500">
                            {{ tier.description }}
                        </p>
                        <p class="mt-2 text-sm text-primary-500">
                            {{ tier.price }}
                        </p>

                        <a
                            :href="tier.href"
                            class="mt-6 block w-full bg-primary-600 border border-transparent rounded-md py-2 text-sm font-semibold text-white text-center hover:bg-primary-700"
                        >
                            {{ tier.cta }}
                        </a>
                    </div>
                    <div class="p-4">
                        <ul
                            role="list"
                            class="mt-2 space-y-4"
                        >
                            <li
                                v-for="feature in tier.features"
                                :key="feature"
                                class="flex space-x-3"
                            >
                                <Icon
                                    name="mdi:check"
                                    class="h-5 w-5 flex-shrink-0 text-primary-500"
                                    aria-hidden="true"
                                />
                                <span class="text-sm text-gray-500">{{ feature }}</span>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
const { t } = useI18n()
const tiers = computed(() => [
    {
        name: t('pricing.plans.cloud.title'),
        href: 'https://projectly.ru/my/auth/register',
        description: t('pricing.plans.cloud.description'),
        price: t('pricing.free'),
        features: [
            t('pricing.plans.cloud.features.members'),
            t('pricing.plans.cloud.features.projects'),
            t('pricing.plans.cloud.features.storage'),
            t('pricing.plans.cloud.features.support'),
        ],
        cta: t('pricing.plans.cloud.cta'),
    },
    {
        name: t('pricing.plans.self-hosted.title'),
        href: 'https://github.com/neketli/projectly?tab=readme-ov-file#for-contributors-%EF%B8%8F',
        description: t('pricing.plans.self-hosted.description'),
        price: t('pricing.free'),
        features: [
            t('pricing.plans.self-hosted.features.members'),
            t('pricing.plans.self-hosted.features.projects'),
            t('pricing.plans.self-hosted.features.storage'),
            t('pricing.plans.self-hosted.features.support'),
            t('pricing.plans.self-hosted.features.infrastructure'),
        ],
        cta: t('pricing.plans.self-hosted.cta'),
    },
    {
        name: t('pricing.plans.enterprise.title'),
        href: '#contact',
        description: t('pricing.plans.enterprise.description'),
        price: t('pricing.plans.enterprise.price'),
        features: [
            t('pricing.plans.enterprise.features.members'),
            t('pricing.plans.enterprise.features.projects'),
            t('pricing.plans.enterprise.features.storage'),
            t('pricing.plans.enterprise.features.support'),
            t('pricing.plans.enterprise.features.infrastructure'),
            t('pricing.plans.enterprise.features.develop'),
        ],
        cta: t('pricing.plans.enterprise.cta'),
    },
])
</script>
