const mix = require('laravel-mix');
const glob = require('glob');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

glob.sync('resources/images/**/!(favicon.ico)').map(
    file => mix.copy(file, 'static/images')
);

mix.options({ processCssUrls: false })
    .webpackConfig({
        watchOptions: {
            ignored: /node_modules|dist|static|mix-manifest.json/
        }
    })
    .js('resources/js/app.js', 'static/js')
    .postCss('resources/css/app.css', 'css', [
        require('@tailwindcss/postcss')
    ])
    .copy('resources/images/favicon.ico', 'static/favicon.ico')
    .vue()
    .version()
    .sourceMaps()
    .setPublicPath('static');
