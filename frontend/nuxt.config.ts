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
    ],
    ssr: false,

    components: {
        global: true,
        dirs: ['~/components'],
    },

    imports: {
        dirs: ['store'],
    },
    devtools: { enabled: process.env.NODE_ENV !== 'production' },

    site: { indexable: false },

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

    eslint: {
        config: {
            stylistic: {
                indent: 2,
                quotes: 'single',
                semi: false,
            },
        },
    },

    icon: {
        localApiEndpoint: '/iconify_api',
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
