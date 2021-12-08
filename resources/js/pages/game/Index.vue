<template>
    <inertia-head title="Deep Sea" />
    <div class="flex py-72">
        <div class="inline-grid gap-1 sm:gap-1.5 lg:gap-2 grid-cols-10 mx-auto">
            <div v-for="item in 100"
                 :key="item"
                 class="bg-sky-500 w-6 h-6 sm:w-12 sm:h-12 lg:w-16 lg:h-16"></div>
        </div>
    </div>
</template>

<script>
import Layout from '../../common/Layout.vue';

export default {
    layout: Layout
};
</script>
