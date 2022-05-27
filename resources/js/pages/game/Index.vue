<template>
    <app-title title="Deep Sea" />
    <div class="flex pt-40">
        <div class="grid grid-cols-11 mx-auto">
            <div></div>
            <div v-for="letter in letters"
                 :key="letter"
                 class="pb-2 flex items-end justify-center text-xs text-cyan-400">
                {{ letter }}
            </div>
            <!-- eslint-disable max-len -->
            <template v-for="(item, index) in 100"
                      :key="item">
                <div v-if="index % 10 === 0"
                     class="pr-2 flex items-center justify-end text-xs text-cyan-400">
                    {{ index / 10 }}
                </div>
                <div class="bg-no-repeat bg-contain w-6 h-6 sm:w-12 sm:h-12 lg:w-16 lg:h-16"
                     :class="randomGrid()"></div>
            </template>
            <!-- eslint-enable max-len -->
        </div>
    </div>
</template>

<script>
import { ref } from 'vue';
import Layout from '../../common/Layout.vue';

export default {
    layout: Layout,

    setup() {
        const letters = ref([
            'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'
        ]);

        const randomGrid = () => {
            const random = Math.floor(Math.random() * 3);

            return ['grid-empty', 'grid-shallow', 'grid-deep'][random];
        };

        return {
            letters,
            randomGrid
        };
    }
};
</script>
