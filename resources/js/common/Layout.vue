<template>
    <!-- eslint-disable max-len -->
    <div v-show="isLeftSidebarOpen || isRightSidebarOpen"
         class="bg-black bg-opacity-50 fixed inset-0 z-10 xl:hidden"
         @click="isLeftSidebarOpen = false; isRightSidebarOpen = false"></div>
    <nav class="nav-left fixed w-80 left-0 top-0 z-30 bg-right-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}">
        <div class="h-16 w-64 mb-12 text-2xl flex items-center justify-center">
            <shopping-bag-icon class="h-8 w-8 text-amber-300 mr-2" />
            Inventory
        </div>
        <div class="sidebar-content w-60 overflow-y-auto" scroll-region>
            <div class="grid grid-cols-2 gap-4">
                <div v-for="artifact in artifacts"
                     :key="artifact.name"
                     class="flex flex-col items-center justify-center">
                    <img class="w-16" :src="`/images/${artifact.name.toLowerCase()}.png`">
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
    <nav class="nav-top fixed h-24 left-1/2 -translate-x-1/2 bg-bottom bg-no-repeat">
        <div class="flex justify-center pt-2">
            <button class="btn xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = true; isRightSidebarOpen = false">
                <shopping-bag-icon class="h-6 w-6" />
            </button>
            <div class="w-60 flex items-center justify-center">
                <div class="grid grid-rows-2 mr-6">
                    <div class="flex items-center justify-center">
                        <div class="mr-2">
                            9
                        </div>
                        <div class="text-sm text-stone-400">
                            /50
                        </div>
                    </div>
                    <div class="text-center text-stone-500 text-xs -mt-1">
                        Boat
                    </div>
                </div>
                <div class="grid grid-rows-2 mr-6">
                    <div class="flex items-center justify-center">
                        <div class="text-cyan-400 mr-2">
                            1
                        </div>
                        <div class="text-sm text-stone-400">
                            /34
                        </div>
                    </div>
                    <div class="text-center text-stone-500 text-xs -mt-1">
                        Diver
                    </div>
                </div>
                <div class="grid grid-rows-2">
                    <div class="flex items-center justify-center">
                        <div class="text-purple-400 mr-2">
                            3
                        </div>
                        <div class="text-sm text-stone-400">
                            /11
                        </div>
                    </div>
                    <div class="text-center text-stone-500 text-xs -mt-1">
                        Submarine
                    </div>
                </div>
            </div>
            <button class="btn xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = false; isRightSidebarOpen = true">
                <library-icon class="h-6 w-6" />
            </button>
        </div>
        <div class="flex items-center justify-center text-lg -mt-2.5">
            <div class="text-amber-300 mr-1">
                $
            </div>
            <div class="text-lime-400">
                100000
            </div>
        </div>
    </nav>
    <nav class="nav-right fixed w-80 right-0 top-0 z-30 bg-left-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isRightSidebarOpen, 'translate-x-full': !isRightSidebarOpen}">
        <div class="h-16 w-64 ml-16 mb-12 text-2xl flex items-center justify-center">
            <library-icon class="h-8 w-8 text-amber-300 mr-2" />
            Museum
        </div>
        <div class="sidebar-content w-60 ml-16 overflow-y-auto" scroll-region>
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
    <!-- eslint-enable max-len -->
    <slot></slot>
</template>

<script>
import {
    LibraryIcon,
    ShoppingBagIcon
} from '@heroicons/vue/outline';

import { ref } from 'vue';

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
        const isRightSidebarOpen = ref(false);

        return {
            isLeftSidebarOpen,
            isRightSidebarOpen
        };
    }
};
</script>
