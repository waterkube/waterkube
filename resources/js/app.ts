import { createInertiaApp, Head, Link } from '@inertiajs/vue3';

import '../css/app.css';

import.meta.glob([
    '../images/**'
]);

createInertiaApp({
    pages: './pages',
    withApp(app) {
        app
            .component('InertiaHead', Head)
            .component('InertiaLink', Link);
    }
});
