<template>
    <inertia-head title="Deep Sea" />
    <div class="flex pt-72">
        <div class="inline-grid gap-1 lg:gap-2 grid-cols-10 mx-auto">
            <div v-for="item in 100" class="bg-lightBlue-500 w-4 h-4 sm:w-8 sm:h-8 lg:w-16 lg:h-16"></div>
        </div>
    </div>
</template>

<script>
import Layout from '../../common/Layout.vue';

export default {
    layout: Layout
};
</script>
