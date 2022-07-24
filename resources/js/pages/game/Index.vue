<template>
    <app-title title="Deep Sea" />
    <div class="flex pt-32">
        <div class="grid gap-1 grid-cols-12 mx-auto">
            <!-- eslint-disable max-len -->
            <div v-for="letter in letters"
                 :key="letter"
                 class="pb-2 flex items-end justify-center text-xs text-cyan-400 sm:text-sm lg:text-base">
                {{ letter }}
            </div>
            <template v-for="(grid, index) in grids"
                      :key="index">
                <div v-if="index !== 0 && index % cols.length === 0"
                     class="pl-2 flex items-center justify-start text-xs text-cyan-400 sm:text-sm lg:text-base">
                    {{ index / cols.length - 1 }}
                </div>
                <div v-if="index % cols.length === 0"
                     class="pr-2 flex items-center justify-end text-xs text-cyan-400 sm:text-sm lg:text-base">
                    {{ index / cols.length }}
                </div>
                <div class="relative bg-no-repeat bg-contain w-6 h-6 sm:w-12 sm:h-12 lg:w-16 lg:h-16"
                     :class="gridClass(grid)">
                    <span v-if="grid.type === 'shallow'"
                          class="flex absolute h-3 w-3 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
                        <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75"></span>
                        <span class="relative inline-flex rounded-full h-3 w-3 bg-cyan-500"></span>
                    </span>
                    <span v-if="grid.type === 'deep'"
                          class="flex absolute h-3 w-3 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
                        <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-purple-400 opacity-75"></span>
                        <span class="relative inline-flex rounded-full h-3 w-3 bg-purple-500"></span>
                    </span>
                </div>
                <div v-if="index === cellCount - 1"
                     class="pl-2 flex items-center justify-start text-xs text-cyan-400 sm:text-sm lg:text-base">
                    {{ index % cols.length }}
                </div>
            </template>
            <div v-for="letter in letters"
                 :key="letter"
                 class="pt-2 flex items-start justify-center text-xs text-cyan-400 sm:text-sm lg:text-base">
                {{ letter }}
            </div>
            <!-- eslint-enable max-len -->
        </div>
    </div>
</template>

<script>
import { ref, toRefs, computed } from 'vue';
import Layout from '../../common/Layout.vue';

export default {
    layout: Layout,

    props: {
        cols: {
            type: Array,
            required: true
        },

        rows: {
            type: Array,
            required: true
        },

        grids: {
            type: Array,
            required: true
        }
    },

    setup(props) {
        const { cols, rows } = toRefs(props);
        const letters = ref(['', ...cols.value, '']);

        const cellCount = computed(() => cols.value.length * rows.value.length);

        const gridClass = grid => {
            if (grid.status === 'undiscovered') {
                return `grid-${grid.type}`;
            }

            return 'grid-empty';
        };

        return {
            letters,
            cellCount,
            gridClass
        };
    }
};
</script>
