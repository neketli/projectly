<template>
    <section
        id="faq"
        class="bg-gray-50 py-24"
    >
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="max-w-3xl mx-auto">
                <h3 class="text-center text-3xl font-extrabold text-gray-900 sm:text-4xl">
                    {{ $t('faq.title') }}
                </h3>

                <dl class="my-6">
                    <Disclosure
                        v-for="(faq, index) in faqs"
                        :key="faq.question"
                        v-slot="{ open }"
                        :default-open="index === 0"
                        as="dt"
                        class="py-4 border-b border-gray-200"
                    >
                        <DisclosureButton
                            class="text-left w-full flex justify-between items-start text-gray-400"
                        >
                            <span class="font-medium text-gray-900">{{ faq.question }}</span>
                            <span class="ml-6 h-7 flex items-center">
                                <Icon
                                    class="h-6 w-6 transform duration-200"
                                    :class="open ? '-rotate-90' : 'rotate-90'"
                                    name="mdi:chevron-right"
                                />
                            </span>
                        </DisclosureButton>
                        <TransitionRoot
                            :show="open"
                            :unmount="false"
                            as="dd"
                            class="mt-2 pr-12"
                            enter="transition ease-in-out duration-200 transform opacity-0"
                            enter-from="-translate-y-[10px] opacity-0"
                            enter-to="translate-y-0 opacity-100"
                            leave="transition ease-in-out duration-200 transform opacity-100"
                            leave-from="translate-y-0 opacity-100"
                            leave-to="-translate-y-[10px] opacity-0"
                        >
                            <p class="text-base text-gray-500">
                                {{ faq.answer }}
                            </p>
                        </TransitionRoot>
                    </Disclosure>
                </dl>
            </div>
        </div>
    </section>
</template>

<script setup>
import { Disclosure, DisclosureButton, TransitionRoot } from '@headlessui/vue'

const { t } = useI18n()
const faqs = computed(() => [
    {
        question: t('faq.questions.projectly.question'),
        answer: t('faq.questions.projectly.answer'),
    },
    {
        question: t('faq.questions.free.question'),
        answer: t('faq.questions.free.answer'),
    },
    {
        question: t('faq.questions.how_to_start.question'),
        answer: t('faq.questions.how_to_start.answer'),
    },
    {
        question: t('faq.questions.customization.question'),
        answer: t('faq.questions.customization.answer'),
    },
    {
        question: t('faq.questions.integration.question'),
        answer: t('faq.questions.integration.answer'),
    },
    {
        question: t('faq.questions.security.question'),
        answer: t('faq.questions.security.answer'),
    },
    {
        question: t('faq.questions.updates.question'),
        answer: t('faq.questions.updates.answer'),
    },
])
</script>
