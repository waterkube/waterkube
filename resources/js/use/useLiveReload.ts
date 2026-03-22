import { ref, onMounted, onUnmounted } from 'vue';

import { router } from '@inertiajs/vue3';

export default (interval = 2500) => {
    const isLive = ref(true);
    let reloadInterval: ReturnType<typeof setInterval>;

    onMounted(() => {
        reloadInterval = setInterval(() => {
            if (isLive.value) {
                router.reload();
            }
        }, interval);
    });

    onUnmounted(() => {
        clearInterval(reloadInterval);
    });

    return {
        isLive
    };
};
