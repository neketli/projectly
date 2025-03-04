<template>
    <section
        id="how-it-works"
        class="py-16 bg-white overflow-hidden"
    >
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="text-center">
                <h2
                    class="text-base font-semibold text-primary-600 tracking-wide uppercase"
                >
                    {{ $t("how_it_works.section") }}
                </h2>
                <h3
                    class="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight"
                >
                    {{ $t("how_it_works.title") }}
                </h3>
                <p class="max-w-xl mt-5 mx-auto text-xl text-gray-500">
                    {{ $t("how_it_works.subtitle") }}
                </p>
            </div>

            <div class="mt-16">
                <div
                    ref="steps"
                    class="grid grid-cols-1 gap-12 lg:grid-cols-3 lg:gap-8 min-h-[180px]"
                >
                    <TransitionGroup
                        enter-active-class="transition ease-out duration-500"
                        enter-from-class="transform opacity-0 scale-95 translate-y-20"
                        enter-to-class="transform opacity-100 scale-100 translate-y-0"
                    >
                        <template
                            v-for="(_, index) in 3"
                            :key="index + stepBlocks[index]"
                        >
                            <div
                                v-show="stepBlocks[index]"
                                class="relative "
                            >
                                <div
                                    class="relative flex items-center justify-center h-12 w-12 rounded-md bg-primary-500 text-white mx-auto shadow-md"
                                >
                                    {{ index + 1 }}
                                </div>
                                <h4 class="mt-6 text-lg font-medium text-gray-900 text-center">
                                    {{ $t(`how_it_works.steps.${index + 1}.title`) }}
                                </h4>
                                <p class="mt-2 text-base text-gray-500 text-center">
                                    {{ $t(`how_it_works.steps.${index + 1}.description`) }}
                                </p>
                            </div>
                        </template>
                    </TransitionGroup>
                </div>
            </div>
        </div>
    </section>
</template>

<script setup>
import { onMounted, ref } from 'vue'

const steps = ref(null)
const stepBlocks = ref([false, false, false])

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

const animationCallback = async () => {
    stepBlocks.value[0] = true
    await sleep(300)
    stepBlocks.value[1] = true
    await sleep(300)
    stepBlocks.value[2] = true
}

onMounted(() => {
    const observer = new IntersectionObserver(
        (entries, observer) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    animationCallback()
                    observer.unobserve(entry.target)
                }
            })
        },
        {
            threshold: 0.2,
        },
    )

    observer.observe(steps.value)
})
</script>
