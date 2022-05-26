<template>
    <app-title title="Deep Sea" />
    <div class="flex pt-40">
        <div class="inline-grid gap-1 sm:gap-1.5 lg:gap-2 grid-cols-10 mx-auto">
            <!-- eslint-disable max-len -->
            <div v-for="item in 100"
                 :key="item"
                 class="bg-no-repeat bg-contain border border-slate-800 w-6 h-6 sm:w-12 sm:h-12 lg:w-16 lg:h-16"
                 :class="randomGrid()"></div>
            <!-- eslint-enable max-len -->
        </div>
    </div>
</template>

<script>
import Layout from '../../common/Layout.vue';

export default {
    layout: Layout,

    setup() {
        const randomGrid = () => {
            const random = Math.floor(Math.random() * 3);

            return ['grid-empty', 'grid-shallow', 'grid-deep'][random];
        };

        return {
            randomGrid
        };
    }
};
</script>
