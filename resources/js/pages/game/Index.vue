<template>
    <Head :title="`Deep Sea - ${title}`" />
    <div v-show="isLeftSidebarOpen || isRightSidebarOpen"
         class="bg-black bg-opacity-50 fixed inset-0 z-10 xl:hidden"
         @click="isLeftSidebarOpen = false; isRightSidebarOpen = false"></div>
    <!-- eslint-disable max-len -->
    <nav class="nav-left fixed w-80 left-0 top-0 z-30 bg-right-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}">
        <div class="h-16 w-64 mb-12 text-2xl flex items-center justify-center">
            Inventory
            <shopping-bag-icon class="h-8 w-8 text-amber-300 ml-2" />
        </div>
        <div ref="leftSidebarContent" class="sidebar-content-left w-60 relative overflow-hidden">
            <div class="grid grid-cols-2 gap-4">
                <div v-for="artifact in discoveredArtifacts"
                     :key="artifact.name"
                     class="flex flex-col items-center justify-center relative">
                    <div v-if="artifact.quantity > 1" class="extra left-3">
                        {{ artifact.quantity }}
                    </div>
                    <div v-if="artifact.type === 'combinable'" class="extra right-3">
                        <plus-icon class="h-4 w-4" />
                    </div>
                    <div v-else-if="artifact.type === 'legendary'" class="extra right-3">
                        <star-icon class="h-4 w-4" />
                    </div>
                    <div :class="`artifact-${artifact.name.toLowerCase()}`"
                         class="w-16 h-16 bg-no-repeat bg-center bg-contain"></div>
                    {{ artifact.name }}
                    <div class="flex items-center text-sm">
                        <div class="text-amber-300 mr-1">
                            $
                        </div>
                        <div class="text-lime-400">
                            {{ artifact.price }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </nav>
    <section class="sidebar fixed z-20 w-64 inset-y-0 bg-black bg-opacity-50 bg-right-top bg-repeat-y transform transition-transform xl:translate-x-0"
             :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}"></section>
    <nav class="nav-top z-10 fixed h-24 left-1/2 -translate-x-1/2 bg-bottom bg-no-repeat">
        <div class="flex justify-center pt-1">
            <button class="btn mt-1 xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = true; isRightSidebarOpen = false">
                <shopping-bag-icon class="h-6 w-6" />
            </button>
            <div class="w-64 px-1 pb-1 grid grid-cols-4">
                <div class="flex flex-col items-center">
                    <div class="dollar w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-lime-400">
                        {{ player.money }}
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="boat w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-slate-300">
                        {{ freeBoat }}
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="diver w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-cyan-400">
                        {{ freeDiver }}
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="submarine w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-center text-purple-400">
                        {{ freeSubmarine }}
                    </div>
                </div>
            </div>
            <button class="btn mt-1 xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = false; isRightSidebarOpen = true">
                <library-icon class="h-6 w-6" />
            </button>
        </div>
        <div class="w-56 h-1.5 mt-1 bg-stone-700 bg-opacity-50 rounded mx-auto">
            <div class="bg-cyan-400 rounded h-1.5" :style="{width: `${progress}%`}"></div>
        </div>
    </nav>
    <nav class="nav-right fixed w-80 right-0 top-0 z-30 bg-left-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isRightSidebarOpen, 'translate-x-full': !isRightSidebarOpen}">
        <div class="h-16 w-64 ml-16 mb-12 text-2xl flex items-center justify-center">
            <library-icon class="h-8 w-8 text-amber-300 mr-2" />
            Museum
        </div>
        <div ref="rightSidebarContent" class="sidebar-content-right w-60 ml-20 relative overflow-hidden">
            <div class="grid grid-cols-2 gap-4">
                <div v-for="artifact in donatedArtifacts"
                     :key="artifact.name"
                     class="flex flex-col items-center justify-center relative">
                    <div v-if="artifact.quantity > 1" class="extra left-3">
                        {{ artifact.quantity }}
                    </div>
                    <div v-if="artifact.type === 'combinable'" class="extra right-3">
                        <plus-icon class="h-4 w-4" />
                    </div>
                    <div v-else-if="artifact.type === 'legendary'" class="extra right-3">
                        <star-icon class="h-4 w-4" />
                    </div>
                    <div :class="`artifact-${artifact.name.toLowerCase()}`"
                         class="w-16 h-16 bg-no-repeat bg-center bg-contain"></div>
                    {{ artifact.name }}
                    <div class="text-sm text-lime-400">
                        +1 Boat
                    </div>
                </div>
            </div>
        </div>
    </nav>
    <section class="sidebar fixed z-20 w-64 right-0 inset-y-0 bg-black bg-opacity-50 bg-left-top bg-repeat-y transform transition-transform xl:translate-x-0"
             :class="{'translate-x-0': isRightSidebarOpen, 'translate-x-full': !isRightSidebarOpen}"></section>
    <div v-if="progress < 100" class="flex pt-32">
        <div class="grid gap-1 grid-cols-12 mx-auto">
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
                    <template v-if="hasExploration(grid)">
                        <span v-if="grid.type === 'shallow'"
                              class="flex absolute h-3 w-3 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
                            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75"></span>
                            <span class="relative inline-flex rounded-full h-3 w-3 bg-cyan-500"></span>
                        </span>
                        <span v-else
                              class="flex absolute h-3 w-3 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
                            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-purple-400 opacity-75"></span>
                            <span class="relative inline-flex rounded-full h-3 w-3 bg-purple-500"></span>
                        </span>
                    </template>
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
        </div>
    </div>
    <div v-else class="container mx-auto pt-32 relative">
        <div class="animate-pulse bg-no-repeat bg-contain bg-center trophy-beam"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-ray"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-glow"></div>
        <div class="absolute inset-x-0 top-32 bg-no-repeat bg-contain bg-center trophy"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-light"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-01"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-02"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-03"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-04"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-05"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-06"></div>
        <div class="absolute inset-x-0 top-32 animate-pulse bg-no-repeat bg-contain bg-center trophy-gleam-07"></div>
    </div>
    <!-- eslint-enable max-len -->
    <footer class="flex items-center justify-center my-4 text-sm text-slate-600">
        Fork me on
        <a class="text-cyan-400 hover:text-cyan-500 ml-2"
           href="https://github.com/waterkube/waterkube"
           target="_blank">
            GitHub
        </a>
    </footer>
</template>

<script>
import {
    ref,
    toRefs,
    computed,
    onMounted,
    onUnmounted
} from 'vue';

import {
    LibraryIcon,
    PlusIcon,
    ShoppingBagIcon,
    StarIcon
} from '@heroicons/vue/outline';

import { Inertia } from '@inertiajs/inertia';
import { Head } from '@inertiajs/inertia-vue3';
import PerfectScrollbar from 'perfect-scrollbar';

export default {
    components: {
        LibraryIcon,
        PlusIcon,
        ShoppingBagIcon,
        StarIcon,
        Head
    },

    props: {
        title: {
            type: String,
            required: true
        },

        player: {
            type: Object,
            required: true
        },

        freeBoat: {
            type: Number,
            required: true
        },

        freeDiver: {
            type: Number,
            required: true
        },

        freeSubmarine: {
            type: Number,
            required: true
        },

        progress: {
            type: Number,
            required: true
        },

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
            default: () => []
        },

        explorations: {
            type: Array,
            default: () => []
        },

        discoveredArtifacts: {
            type: Array,
            default: () => []
        },

        donatedArtifacts: {
            type: Array,
            default: () => []
        }
    },

    setup(props) {
        const { cols, rows, explorations } = toRefs(props);
        const isLeftSidebarOpen = ref(false);
        const leftSidebarContent = ref();
        const leftSidebarPs = ref(undefined);
        const isRightSidebarOpen = ref(false);
        const rightSidebarContent = ref();
        const rightSidebarPs = ref(undefined);
        const letters = ref(['', ...cols.value, '']);
        const reloadInterval = ref(undefined);
        const reloadTimer = 2500;

        const cellCount = computed(() => cols.value.length * rows.value.length);

        const hasExploration = grid => explorations.value
            && explorations.value.indexOf(grid.name) !== -1;

        const gridClass = grid => {
            if (grid.status === 'undiscovered' || hasExploration(grid)) {
                return `grid-${grid.type}`;
            }

            return 'grid-empty';
        };

        onMounted(() => {
            leftSidebarPs.value = new PerfectScrollbar(leftSidebarContent.value);
            rightSidebarPs.value = new PerfectScrollbar(rightSidebarContent.value);
            reloadInterval.value = setInterval(() => Inertia.reload(), reloadTimer);
        });

        onUnmounted(() => {
            clearInterval(reloadInterval.value);
        });

        return {
            isLeftSidebarOpen,
            isRightSidebarOpen,
            leftSidebarContent,
            rightSidebarContent,
            letters,
            cellCount,
            hasExploration,
            gridClass
        };
    }
};
</script>
