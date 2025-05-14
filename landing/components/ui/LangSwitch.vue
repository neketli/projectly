<template>
    <div :key="$i18n.locale" class="flex justify-center gap-4">
        <button
            :class="{
                'text-gray-600 cursor-default': $i18n.locale === 'ru',
                'text-primary-600 transition hover:text-primary-700 cursor-pointer': $i18n.locale !== 'ru',
            }"
            @click="$i18n.setLocale('ru')"
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
            @click="$i18n.setLocale('en')"
        >
            {{ $t('header.en') }}
        </button>
    </div>
</template>
