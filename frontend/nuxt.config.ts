// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        '@nuxt/devtools',
        '@nuxt/image',
        '@nuxt/icon',
        '@nuxt/eslint',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
        '@nuxtjs/color-mode',
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
        '@element-plus/nuxt',
        '@formkit/auto-animate',
        '@vueuse/nuxt',
        'nuxt-svgo',
        'yandex-metrika-module-nuxt3',
    ],
    ssr: false,

    components: {
        global: true,
        dirs: ['~/components'],
    },

    imports: {
        dirs: ['store'],
    },
    devtools: { enabled: false }, // process.env.NODE_ENV !== 'production'

    site: { indexable: false },

    colorMode: {
        preference: 'light',
        fallback: 'light',
        classSuffix: '',
    },

    runtimeConfig: {
        public: {
            API_HOST: process.env.NUXT_PUBLIC_API_HOST,
            S3_HOST: process.env.NUXT_PUBLIC_S3_HOST,
        },
    },
    compatibilityDate: '2024-04-03',

    vite: {
        optimizeDeps: {
            esbuildOptions: {
                define: {
                    global: 'globalThis',
                },
            },
        },
    },
    baseUrl: process.env.NUXT_APP_BASE_URL,

    elementPlus: {
        importStyle: 'css',
        themes: ['dark'],
    },

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
                code: 'ru-RU',
                label: 'Русский',
                flag: 'ru',
                file: 'ru-RU.yaml',
            },
            {
                code: 'en-US',
                label: 'English',
                flag: 'us',
                file: 'en-US.yaml',
            },
        ],
        strategy: 'no_prefix',
        lazy: true,
        defaultLocale: 'en-US',
        detectBrowserLanguage: {
            useCookie: true,
            cookieKey: 'i18n_redirected',
            redirectOn: 'root',
        },
        compilation: {
            strictMessage: false,
        },
    },

    icon: {
        serverBundle: false,
        clientBundle: {
            icons: [
                'flag:ru-4x3',
                'flag:us-4x3',
            ],
            scan: true,
            sizeLimitKb: 256,
        },
    },

    pinia: {
        autoImports: ['defineStore', 'acceptHMRUpdate'],
    },

    piniaPersistedstate: {
        cookieOptions: {
            sameSite: 'strict',
        },
        storage: 'localStorage',
    },

    yandexMetrika: {
        id: 100137750,
        clickmap: true,
        trackLinks: true,
        accurateTrackBounce: true,
        webvisor: true,
        consoleLog: false,
    },
})
