import pluginVue from 'eslint-plugin-vue'
import stylistic from '@stylistic/eslint-plugin'
import withNuxt from './.nuxt/eslint.config.mjs'

export default withNuxt(
    ...pluginVue.configs['flat/recommended'],
    {
        plugins: {
            '@stylistic': stylistic,
        },
    },
    {
        rules: {
            'no-console': ['error'],
            '@stylistic/max-len': [
                2,
                120,
                {
                    ignoreUrls: true,
                    ignoreComments: false,
                    ignoreRegExpLiterals: true,
                    ignoreStrings: true,
                    ignoreTemplateLiterals: true,
                    ignorePattern: 'class="*"',
                },
            ],
            '@stylistic/indent': ['error', 4],
            '@stylistic/array-bracket-spacing': ['error', 'never'],
            '@stylistic/no-trailing-spaces': ['error'],
            '@stylistic/no-multiple-empty-lines': [
                'error',
                {
                    max: 1,
                    maxBOF: 0,
                    maxEOF: 0,
                },
            ],

            'object-shorthand': ['error', 'always', { avoidQuotes: false }],
            'no-proto': 'error',
            'no-mixed-operators': 'error',
            'no-shadow': 'off',
            'vue/v-on-style': ['error', 'shorthand'],

            'vue/html-indent': [
                'error',
                4,
                {
                    attribute: 1,
                    baseIndent: 1,
                    closeBracket: 0,
                    alignAttributesVertically: true,
                    ignores: [],
                },
            ],
            'vue/multi-word-component-names': 'off',
            'vue/max-attributes-per-line': [
                'error',
                {
                    singleline: { max: 1 },
                    multiline: { max: 1 },
                },
            ],
            'vue/no-dupe-keys': 'warn',
            'vue/require-toggle-inside-transition': 'warn',
            'vue/no-v-html': 'off',
            'vue/no-deprecated-v-on-native-modifier': 'warn',

            'vue/attributes-order': [
                'error',
                {
                    order: [
                        'DEFINITION',
                        'LIST_RENDERING',
                        'CONDITIONALS',
                        'RENDER_MODIFIERS',
                        'GLOBAL',
                        ['UNIQUE', 'SLOT'],
                        'TWO_WAY_BINDING',
                        'OTHER_DIRECTIVES',
                        'OTHER_ATTR',
                        'EVENTS',
                        'CONTENT',
                    ],
                    alphabetical: false,
                },
            ],
        },
    },
)
