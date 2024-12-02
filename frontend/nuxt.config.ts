// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        '@nuxt/devtools',
        '@nuxt/image',
        '@nuxt/icon',
        '@nuxt/eslint',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
        '@nuxtjs/sitemap',
        '@nuxtjs/robots',
        '@nuxtjs/color-mode',
        '@element-plus/nuxt',
        '@nuxtjs/tailwindcss',
        '@formkit/auto-animate',
        '@nuxtjs/i18n',
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
        preference: 'system',
        fallback: 'light',
        classSuffix: '',
    },

    runtimeConfig: {
        public: {
            API_URL: process.env.VITE_API_URL,
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
    },

    icon: {
        serverBundle: {
            collections: ['mdi'],
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

})
