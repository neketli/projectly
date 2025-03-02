<template>
    <section
        id="how-it-works"
        class="py-16 bg-white overflow-hidden"
    >
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="text-center">
                <h2 class="text-base font-semibold text-primary-600 tracking-wide uppercase">
                    {{ $t('how_it_works.section') }}
                </h2>
                <h3 class="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight">
                    {{ $t('how_it_works.title') }}
                </h3>
                <p class="max-w-xl mt-5 mx-auto text-xl text-gray-500">
                    {{ $t('how_it_works.subtitle') }}
                </p>
            </div>

            <div class="mt-16">
                <div class="space-y-12 lg:space-y-0 lg:grid lg:grid-cols-3 lg:gap-x-8">
                    <div
                        v-for="(_, index) in 3"
                        :key="index"
                        class="relative"
                    >
                        <div class="relative pb-12 lg:pb-0">
                            <div class="relative flex items-center justify-center h-12 w-12 rounded-md bg-primary-500 text-white mx-auto">
                                {{ index + 1 }}
                            </div>
                            <h4 class="mt-6 text-lg font-medium text-gray-900 text-center">
                                {{ $t(`how_it_works.steps.${index+1}.title`) }}
                            </h4>
                            <p class="mt-2 text-base text-gray-500 text-center">
                                {{ $t(`how_it_works.steps.${index+1}.description`) }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
