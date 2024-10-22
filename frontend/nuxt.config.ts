// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
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

  eslint: {
    config: {
      stylistic: {
        indent: 2,
        quotes: 'single',
        semi: false,
      },
    },
  },

  site: { indexable: false },

  devtools: {
    enabled: process.env.NODE_ENV !== 'production',
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

  components: {
    global: true,
    dirs: ['~/components'],
  },

  imports: {
    dirs: ['store'],
  },

  vite: {
    optimizeDeps: {
      esbuildOptions: {
        define: {
          global: 'globalThis',
        },
      },
    },
  },
});
