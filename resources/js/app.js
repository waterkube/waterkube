import { createApp, h } from 'vue';
import { createInertiaApp, InertiaLink } from '@inertiajs/inertia-vue3';
import { InertiaProgress } from '@inertiajs/progress';
import AppTitle from './common/AppTitle.vue';

window._ = require('lodash');

InertiaProgress.init();

createInertiaApp({
    // eslint-disable-next-line import/no-dynamic-require
    resolve: name => require(`./pages/${name}`),
    setup({
        el, App, props, plugin
    }) {
        createApp({ render: () => h(App, props) })
            .use(plugin)
            .component('InertiaLink', InertiaLink)
            .component('AppTitle', AppTitle)
            .mount(el);
    }
});
