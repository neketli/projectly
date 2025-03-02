<template>
    <section
        id="demo"
        class="bg-gray-50 py-16 overflow-hidden"
    >
        <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="relative">
                <div class="text-center">
                    <h2 class="text-base font-semibold text-primary-600 tracking-wide uppercase">
                        {{ $t('demo.title') }}
                    </h2>
                    <p class="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight">
                        {{ $t('demo.subtitle') }}
                    </p>
                </div>
                <div class="mt-6">
                    <Splide :options="options">
                        <SplideSlide
                            v-for="item in slides"
                            :key="item.id"
                        >
                            <p class="text-2xl text-center mb-2">
                                {{ item.title }}
                            </p>

                            <p class="max-w-3xl text-gray-600 mx-auto text-center mb-4">
                                {{ item.description }}
                            </p>
                            <img
                                class="h-72 md:h-96 lg:h-[600px] w-auto mx-auto mb-2 shadow-md rounded-md"
                                :src="item.image"
                                :alt="item.alt"
                                loading="lazy"
                            >
                        </SplideSlide>
                    </Splide>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import { Splide, SplideSlide } from '@splidejs/vue-splide'
import '@splidejs/vue-splide/css'

import boardImage from '@/assets/board.webp'
import detailImage from '@/assets/detail.webp'
import analyticsImage from '@/assets/analytics.webp'

const { t } = useI18n()

const slides = computed(() => [
    {
        id: 'board',
        image: boardImage,
        title: t('demo.features.board.title'),
        description: t('demo.features.board.description'),
    },
    {
        id: 'detail',
        image: detailImage,
        title: t('demo.features.detail.title'),
        description: t('demo.features.detail.description'),
    },
    {
        id: 'analytics',
        image: analyticsImage,
        title: t('demo.features.analytics.title'),
        description: t('demo.features.analytics.description'),
    },
])

const options = {
    rewind: true,
    type: 'loop',
    gap: '1rem',
    padding: '1rem',
    perPage: 1,
    lazyLoad: 'nearby',
    autoplay: true,
    pagination: false,
}
</script>
