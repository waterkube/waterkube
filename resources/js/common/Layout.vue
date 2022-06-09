<template>
    <!-- eslint-disable max-len -->
    <div v-show="isLeftSidebarOpen || isRightSidebarOpen"
         class="bg-black bg-opacity-50 fixed inset-0 z-10 xl:hidden"
         @click="isLeftSidebarOpen = false; isRightSidebarOpen = false"></div>
    <nav class="nav-left fixed w-80 left-0 top-0 z-30 bg-right-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}">
        <div class="h-16 w-64 mb-12 text-2xl flex items-center justify-center">
            Inventory
            <shopping-bag-icon class="h-8 w-8 text-amber-300 ml-2" />
        </div>
        <div ref="leftSidebarContent" class="sidebar-content-left w-60 relative overflow-hidden">
            <div class="grid grid-cols-2 gap-4">
                <div v-for="artifact in artifacts"
                     :key="artifact.name"
                     class="flex flex-col items-center justify-center">
                    <img class="w-16 h-16" :src="`/images/${artifact.name.toLowerCase()}.png`">
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
                        1000
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="boat w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-slate-300">
                        9
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="diver w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-cyan-400">
                        1
                    </div>
                </div>
                <div class="flex flex-col items-center">
                    <div class="submarine w-12 h-12 bg-no-repeat bg-center bg-contain"></div>
                    <div class="text-center text-purple-400">
                        3
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
            <div class="bg-cyan-400 rounded h-1.5" style="width: 11%"></div>
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
                <div v-for="artifact in artifacts"
                     :key="artifact.name"
                     class="flex flex-col items-center justify-center">
                    <img class="w-16" :src="`/images/${artifact.name.toLowerCase()}.png`">
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
    <slot></slot>
    <div class="flex items-center justify-center my-4 text-sm text-slate-600">
        Fork me on
        <a class="text-cyan-400 hover:text-cyan-500 ml-2"
           href="https://github.com/waterkube/waterkube"
           target="_blank">
            GitHub
        </a>
    </div>
    <!-- eslint-enable max-len -->
</template>

<script>
import {
    LibraryIcon,
    ShoppingBagIcon
} from '@heroicons/vue/outline';

import { ref, onMounted } from 'vue';
import PerfectScrollbar from 'perfect-scrollbar';

export default {
    components: {
        LibraryIcon,
        ShoppingBagIcon
    },

    props: {
        artifacts: {
            type: Array,
            default: () => []
        }
    },

    setup() {
        const isLeftSidebarOpen = ref(false);
        const leftSidebarContent = ref();
        const leftSidebarPs = ref(undefined);

        const isRightSidebarOpen = ref(false);
        const rightSidebarContent = ref();
        const rightSidebarPs = ref(undefined);

        onMounted(() => {
            leftSidebarPs.value = new PerfectScrollbar(leftSidebarContent.value);
            rightSidebarPs.value = new PerfectScrollbar(rightSidebarContent.value);
        });

        return {
            isLeftSidebarOpen,
            isRightSidebarOpen,
            leftSidebarContent,
            rightSidebarContent
        };
    }
};
</script>
