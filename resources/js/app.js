import { createApp, h } from 'vue';
import { createInertiaApp, InertiaLink, InertiaHead } from '@inertiajs/inertia-vue3';
import { InertiaProgress } from '@inertiajs/progress';

window._ = require('lodash');

InertiaProgress.init();

createInertiaApp({
    // eslint-disable-next-line import/no-dynamic-require
    resolve: name => require(`./pages/${name}`),
    setup({
        el, App, props, plugin
    }) {
        props.titleCallback = title => (title
            ? `${title} - ${props.initialPage.props.title}`
            : props.initialPage.props.title);

        createApp({ render: () => h(App, props) })
            .use(plugin)
            .component('InertiaLink', InertiaLink)
            .component('InertiaHead', InertiaHead)
            .mount(el);
    }
});