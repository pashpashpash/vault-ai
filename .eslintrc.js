module.exports = {
    env: {
        browser: true,
        es6: true,
        amd: true,
        node: true,
    },
    globals: {
        PIXI: 'readonly',
    },
    extends: ['eslint:recommended', 'plugin:react/recommended'],
    parser: 'babel-eslint',
    parserOptions: {
        ecmaVersion: 8,
        ecmaFeatures: {
            experimentalObjectRestSpread: true,
            jsx: true,
        },
        sourceType: 'module',
    },
    plugins: [
        'react',
        'eslint-plugin-google-camelcase',
        'eslint-plugin-flowtype',
        'react-hooks',
    ],
    rules: {
        'react-hooks/rules-of-hooks': 'error', // Checks rules of Hooks
        'react-hooks/exhaustive-deps': 'warn', // Checks effect dependencies
        'accessor-pairs': 'error',
        'array-bracket-newline': 'off',
        'array-bracket-spacing': 'off',
        'array-callback-return': 'error',
        'array-element-newline': 'off',
        'arrow-body-style': 'error',
        'arrow-spacing': 'error',
        'block-scoped-var': 'error',
        'block-spacing': 'error',
        'callback-return': 'off',
        camelcase: 'off',
        'google-camelcase/google-camelcase': 'error',
        'capitalized-comments': 'off',
        'class-methods-use-this': 'off',
        'comma-dangle': 'only-multiline',
        'comma-spacing': [
            'error',
            {
                after: true,
                before: false,
            },
        ],
        'comma-style': ['error', 'last'],
        complexity: ['error', { max: 26 }],
        'computed-property-spacing': 'error',
        'consistent-return': 'error',
        'consistent-this': 'off',
        curly: 'off',
        'default-case': 'error',
        'eol-last': 'error',
        eqeqeq: ['error', 'always', { null: 'ignore' }],
        'for-direction': 'error',
        'func-call-spacing': 'error',
        'func-name-matching': 'error',
        'func-names': ['error', 'never'],
        'func-style': 'error',
        'generator-star-spacing': 'error',
        'getter-return': 'error',
        'global-require': 'off',
        'guard-for-in': 'error',
        'handle-callback-err': 'error',
        'id-blacklist': 'error',
        'id-length': 'off',
        'id-match': 'error',
        'implicit-arrow-linebreak': 'off',
        indent: 'off',
        'indent-legacy': 'off',
        'init-declarations': 'off',
        'jsx-quotes': ['error', 'prefer-double'],
        'key-spacing': 'off',
        'keyword-spacing': [
            'error',
            {
                after: true,
                before: true,
            },
        ],
        'line-comment-position': 'off',
        'linebreak-style': ['error', 'unix'],
        'lines-around-comment': 'error',
        'lines-around-directive': 'error',
        'lines-between-class-members': [
            'error',
            'always',
            { exceptAfterSingleLine: true },
        ],
        'max-depth': 'error',
        'max-len': 'off',
        'max-nested-callbacks': 'error',
        'max-params': 'off',
        'max-statements': 'off',
        'max-statements-per-line': 'error',
        'multiline-comment-style': 'off',
        'new-cap': ['error', { newIsCapExceptionPattern: 'PIXI.*' }],
        'new-parens': 'error',
        'newline-after-var': 'off',
        'newline-before-return': 'off',
        'newline-per-chained-call': 'off',
        'no-alert': 'error',
        'no-array-constructor': 'error',
        'no-await-in-loop': 'off',
        'no-bitwise': 'error',
        'no-buffer-constructor': 'error',
        'no-caller': 'error',
        'no-catch-shadow': 'error',
        'no-confusing-arrow': 'off',
        'no-continue': 'error',
        'no-console': 'off',
        'no-constant-condition': 'off',
        'no-div-regex': 'error',
        'no-duplicate-imports': 'error',
        'no-else-return': 'off',
        'no-empty': 'off',
        'no-empty-function': 'off',
        'no-eval': 'error',
        'no-extend-native': 'error',
        'no-extra-bind': 'off',
        'no-extra-label': 'error',
        'no-extra-parens': 'off',
        'no-floating-decimal': 'error',
        'no-implicit-coercion': 'off',
        'no-implicit-globals': 'error',
        'no-implied-eval': 'error',
        'no-inline-comments': 'off',
        'no-invalid-this': 'off',
        'no-iterator': 'error',
        'no-label-var': 'error',
        'no-labels': 'error',
        'no-lone-blocks': 'error',
        'no-lonely-if': 'error',
        'no-loop-func': 'error',
        'no-magic-numbers': 'off',
        'no-mixed-operators': 'off',
        'no-mixed-requires': 'error',
        'no-multi-assign': 'error',
        'no-multi-spaces': 'off',
        'no-multi-str': 'error',
        'no-multiple-empty-lines': 'error',
        'no-native-reassign': 'error',
        'no-negated-condition': 'off',
        'no-negated-in-lhs': 'error',
        'no-nested-ternary': 'error',
        'no-new': 'error',
        'no-new-func': 'error',
        'no-new-object': 'error',
        'no-new-require': 'error',
        'no-new-wrappers': 'error',
        'no-octal-escape': 'error',
        'no-param-reassign': 'off',
        'no-path-concat': 'error',
        'no-plusplus': 'off',
        'no-process-env': 'error',
        'no-process-exit': 'error',
        'no-proto': 'error',
        'no-prototype-builtins': 'error',
        'no-restricted-globals': 'error',
        'no-restricted-imports': 'error',
        'no-restricted-modules': 'error',
        'no-restricted-properties': 'error',
        'no-restricted-syntax': 'error',
        'no-return-assign': 'error',
        'no-return-await': 'error',
        'no-script-url': 'error',
        'no-self-compare': 'error',
        'no-sequences': 'error',
        'no-shadow': 'error',
        'no-shadow-restricted-names': 'error',
        'no-spaced-func': 'error',
        'no-sync': 'error',
        'no-tabs': 'error',
        'no-template-curly-in-string': 'error',
        'no-ternary': 'off',
        'no-throw-literal': 'error',
        'no-trailing-spaces': 'error',
        'no-undef-init': 'off',
        'no-undefined': 'off',
        'no-underscore-dangle': 'off',
        'no-unmodified-loop-condition': 'error',
        'no-unneeded-ternary': 'error',
        'no-unused-expressions': ['error', { allowShortCircuit: true }],
        'no-unused-vars': ['error', { args: 'none' }],
        'no-use-before-define': 'error',
        'no-useless-call': 'error',
        'no-useless-computed-key': 'error',
        'no-useless-concat': 'error',
        'no-useless-constructor': 'off',
        'no-useless-rename': 'error',
        'no-useless-return': 'off',
        'no-var': 'off',
        'no-void': 'error',
        'no-warning-comments': 'off',
        'no-whitespace-before-property': 'error',
        'no-with': 'error',
        'nonblock-statement-body-position': 'error',
        'object-curly-newline': 'off',
        'object-curly-spacing': 'off',
        'object-property-newline': 'off',
        'object-shorthand': 'error',
        'one-var': 'off',
        'one-var-declaration-per-line': 'error',
        'operator-assignment': 'error',
        'operator-linebreak': 'off',
        'padded-blocks': 'off',
        'padding-line-between-statements': 'error',
        'prefer-arrow-callback': 'off',
        'prefer-const': 'error',
        'prefer-numeric-literals': 'off',
        'prefer-promise-reject-errors': 'error',
        'prefer-reflect': 'off',
        'prefer-rest-params': 'error',
        'prefer-spread': 'error',
        'quote-props': 'off',
        quotes: 'off',
        radix: 'error',
        'react/no-unknown-property': ['error', { ignore: ['class'] }],
        'react/prop-types': 'off',
        'react/no-string-refs': 'off',
        'react/no-unescaped-entities': 'off',
        'react/display-name': 'off',
        'require-await': 'off',
        'require-jsdoc': 'error',
        'rest-spread-spacing': 'error',
        semi: 'off',
        'semi-spacing': 'error',
        'semi-style': ['error', 'last'],
        'sort-imports': 'off',
        'sort-keys': 'off',
        'sort-vars': 'off',
        'space-before-blocks': 'error',
        'space-before-function-paren': 'off',
        'space-in-parens': ['error', 'never'],
        'space-infix-ops': 'error',
        'space-unary-ops': 'error',
        'spaced-comment': ['error', 'always'],
        strict: 'error',
        'switch-colon-spacing': 'error',
        'symbol-description': 'error',
        'template-curly-spacing': 'error',
        'template-tag-spacing': 'error',
        'unicode-bom': ['error', 'never'],
        'valid-jsdoc': 'error',
        'vars-on-top': 'off',
        'wrap-iife': 'error',
        'wrap-regex': 'error',
        'yield-star-spacing': 'error',
        yoda: ['error', 'never'],
        'flowtype/boolean-style': [2, 'boolean'],
        'flowtype/define-flow-type': 1,
        'flowtype/delimiter-dangle': [2, 'never'],
        'flowtype/object-type-delimiter': [2, 'comma'],
        'flowtype/require-parameter-type': 2,
        'flowtype/require-readonly-react-props': 0,
        'flowtype/require-return-type': [
            2,
            'always',
            {
                any: true,
                Object: false,
                Function: false,
            },
        ],
        'flowtype/require-valid-file-annotation': 2,
        'flowtype/semi': [2, 'always'],
        'flowtype/space-after-type-colon': [2, 'always'],
        'flowtype/space-before-generic-bracket': [2, 'never'],
        'flowtype/space-before-type-colon': [2, 'never'],
        'flowtype/type-id-match': [2, '^([A-Z][a-z0-9]+)+$'],
        'flowtype/union-intersection-spacing': [2, 'always'],
        'flowtype/use-flow-type': 1,
        'flowtype/valid-syntax': 1,
    },
};
