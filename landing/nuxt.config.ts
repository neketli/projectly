// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
        '@nuxtjs/robots',
        '@nuxtjs/sitemap',
        '@nuxt/icon',
        '@nuxt/eslint',
    ],
    devtools: { enabled: false },

    app: {
        head: {
            title: 'Projectly - идеальный инструмент для управления проектами',
            meta: [
                { charset: 'utf-8' },
                { name: 'viewport', content: 'width=device-width, initial-scale=1' },
                { hid: 'description', name: 'description', content: 'Projectly помогает командам работать вместе, отслеживать задачи и завершать проекты эффективно' },
            ],
            link: [
                { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
                { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap' },
            ],
        },
    },

    site: {
        url: 'https://projectly.ru',
        indexable: true,
    },

    compatibilityDate: '2024-04-03',

    eslint: {
        config: {
            stylistic: {
                indent: 2,
                quotes: 'single',
                semi: false,
            },
        },
    },

    i18n: {
        locales: [
            {
                code: 'ru',
                label: 'Русский',
                file: 'ru.json',
            },
            {
                code: 'en',
                label: 'English',
                file: 'en.json',
            },
        ],
        strategy: 'prefix_except_default',
        defaultLocale: 'ru',
        compilation: {
            strictMessage: false,
        },

        detectBrowserLanguage: {
            useCookie: true,
            cookieKey: 'i18n_redirected',
            redirectOn: 'root',
        },
    },

    robots: {
        UserAgent: '*',
        Allow: '/',
        Sitemap: 'https://projectly.ru/sitemap.xml',
    },
    sitemap: {
        hostname: 'https://projectly.ru',
        gzip: true,
    },

    yandexMetrika: {
        id: process.env.NUXT_YM_ID,
        debug: process.env.NODE_ENV !== 'production',
        clickmap: true,
        trackLinks: true,
        accurateTrackBounce: true,
        webvisor: true,
        consoleLog: false,
    },

})
