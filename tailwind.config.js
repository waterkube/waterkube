const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
    mode: 'jit',
    purge: [
        './resources/**/*.gohtml',
        './resources/**/*.js',
        './resources/**/*.vue'
    ],
    darkMode: 'media', // or 'media' or 'class'
    theme: {
        fontFamily: {
            sans: ['Sarabun', ...defaultTheme.fontFamily.sans]
        },
        extend: {
            colors: {}
        }
    },
    variants: {
        extend: {}
    },
    plugins: [
        require('@tailwindcss/forms')
    ]
};
