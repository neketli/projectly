<template>
    <section
        id="features"
        class="py-16 bg-gray-50"
    >
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="lg:text-center">
                <h2 class="text-base text-primary-600 font-semibold tracking-wide uppercase">
                    {{ $t('features.section') }}
                </h2>
                <h3 class="mt-2 text-3xl leading-8 font-extrabold tracking-tight text-gray-900 sm:text-4xl">
                    {{ $t('features.title') }}
                </h3>
                <p class="mt-4 max-w-2xl text-xl text-gray-500 lg:mx-auto">
                    {{ $t('features.subtitle') }}
                </p>
            </div>

            <div class="mt-10">
                <dl class="space-y-10 md:space-y-0 md:grid md:grid-cols-2 md:gap-x-8 md:gap-y-10">
                    <template
                        v-for="feature in features"
                        :key="feature.name"
                    >
                        <div
                            v-if="!feature.isHidden"
                            class="relative"
                        >
                            <dt>
                                <div
                                    class="absolute flex items-center justify-center h-12 w-12 rounded-md bg-primary-500 text-white shadow-md"
                                    @mouseover="currentIcon = feature.key"
                                    @mouseleave="currentIcon = ''"
                                >
                                    <Icon
                                        :key="feature.icon + currentIcon"
                                        :name="currentIcon === feature.key ? feature.iconActive : feature.icon"
                                        class="h-6 w-6 transition-all"
                                    />
                                </div>
                                <p class="ml-16 text-lg leading-6 font-medium text-gray-900">
                                    {{ feature.title }}
                                </p>
                            </dt>
                            <dd class="mt-2 ml-16 text-base text-gray-500">
                                {{ feature.description }}
                            </dd>
                        </div>
                    </template>
                </dl>
            </div>
        </div>
    </section>
</template>

<script setup>
const { t, locale } = useI18n()

const currentIcon = ref('')
const features = computed(() => [
    {
        key: 'kanban',
        title: t('features.list.kanban.title'),
        description: t('features.list.kanban.description'),
        icon: 'line-md:clipboard',
        iconActive: 'line-md:clipboard-check',
    },
    {
        key: 'tasks',
        title: t('features.list.tasks.title'),
        description: t('features.list.tasks.description'),
        icon: 'line-md:check-list-3',
        iconActive: 'line-md:check-list-3-filled',
    },
    {
        key: 'local',
        title: t('features.list.local.title'),
        description: t('features.list.local.description'),
        icon: 'line-md:map-marker',
        iconActive: 'line-md:map-marker-loop',
        isHidden: locale.value !== 'ru',
    },
    {
        key: 'deadline',
        title: t('features.list.deadline.title'),
        description: t('features.list.deadline.description'),
        icon: 'line-md:speed',
        iconActive: 'line-md:speed-loop',
        isHidden: locale.value !== 'en',
    },
    {
        key: 'team',
        title: t('features.list.team.title'),
        description: t('features.list.team.description'),
        icon: 'line-md:account-small',
        iconActive: 'line-md:account',
    },
    {
        key: 'workflows',
        title: t('features.list.workflows.title'),
        description: t('features.list.workflows.description'),
        icon: 'line-md:cog',
        iconActive: 'line-md:cog-loop',
    },
    {
        key: 'analytics',
        title: t('features.list.analytics.title'),
        description: t('features.list.analytics.description'),
        icon: 'line-md:document-report',
        iconActive: 'line-md:document-report-twotone',
    },
])
</script>
