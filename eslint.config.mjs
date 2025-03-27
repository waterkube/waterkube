import { defineConfig } from 'eslint/config';
import globals from 'globals';
import parser from 'vue-eslint-parser';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import js from '@eslint/js';
import { FlatCompat } from '@eslint/eslintrc';
import pluginVue from 'eslint-plugin-vue';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const compat = new FlatCompat({
    baseDirectory: __dirname,
    recommendedConfig: js.configs.recommended,
    allConfig: js.configs.all
});

export default defineConfig([
    ...pluginVue.configs['flat/recommended'],
    {
        extends: compat.extends('airbnb-base'),

        languageOptions: {
            globals: {
                ...globals.browser,
                ...globals.commonjs,
                ...globals.jquery,
                ...globals.node,
                _: 'readonly'
            },

            parser,
            ecmaVersion: 6,
            sourceType: 'module',

            parserOptions: {
                parser: '@babel/eslint-parser',
                requireConfigFile: false
            }
        },

        settings: {
            'import/core-modules': ['webpack']
        },

        rules: {
            'arrow-parens': ['error', 'as-needed'],
            'comma-dangle': ['error', 'never'],
            'global-require': 'off',
            'function-paren-newline': ['error', 'consistent'],
            'import/extensions': ['error', 'never', {
                vue: 'always'
            }],
            'import/no-unresolved': 'off',
            'import/no-extraneous-dependencies': ['error', {
                devDependencies: true
            }],
            indent: ['error', 4],
            'keyword-spacing': ['error', {
                overrides: {
                    this: {
                        before: false
                    }
                }
            }],
            'object-shorthand': ['error', 'always', {
                avoidQuotes: false
            }],
            'no-console': 'off',
            'no-empty': ['error', {
                allowEmptyCatch: true
            }],
            'no-multi-assign': 'off',
            'no-param-reassign': ['error', {
                props: false
            }],
            'no-plusplus': ['error', {
                allowForLoopAfterthoughts: true
            }],
            'no-underscore-dangle': 'off',
            'no-unused-vars': 'off',
            'vue/first-attribute-linebreak': ['error', {
                multiline: 'beside'
            }],
            'vue/html-closing-bracket-newline': 'off',
            'vue/html-indent': ['error', 4],
            'vue/html-self-closing': ['error', {
                html: {
                    normal: 'never'
                }
            }],
            'vue/max-attributes-per-line': ['error', {
                singleline: 2
            }],
            'vue/multi-word-component-names': 'off',
            'vue/no-v-html': 'off'
        }
    }
]);
