<template>
    <!-- eslint-disable max-len -->
    <div v-show="isLeftSidebarOpen || isRightSidebarOpen"
         class="bg-black bg-opacity-50 fixed inset-0 z-10 xl:hidden"
         @click="isLeftSidebarOpen = false; isRightSidebarOpen = false"></div>
    <nav class="nav-left fixed w-80 left-0 top-0 z-30 bg-right-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}"></nav>
    <section class="sidebar fixed z-20 w-64 inset-y-0 bg-black bg-opacity-50 bg-right-top bg-repeat-y transform transition-transform xl:translate-x-0"
             :class="{'translate-x-0': isLeftSidebarOpen, '-translate-x-full': !isLeftSidebarOpen}"></section>
    <nav class="nav-top fixed h-24 left-1/2 -translate-x-1/2 bg-bottom bg-no-repeat">
        <div class="flex justify-center pt-2">
            <button class="btn xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = true; isRightSidebarOpen = false">
                <shopping-bag-icon class="h-6 w-6" />
            </button>
            <div class="h-9 w-60 flex items-center justify-center text-xl">
                <span class="mr-2">
                    9
                </span>
                <span class="text-sm text-slate-400 mr-6">
                    /50
                </span>
                <span class="text-cyan-400 mr-2">
                    1
                </span>
                <span class="text-sm text-slate-400 mr-6">
                    /34
                </span>
                <span class="text-fuchsia-400 mr-2">
                    3
                </span>
                <span class="text-sm text-slate-400">
                    /11
                </span>
            </div>
            <button class="btn xl:hidden"
                    type="button"
                    @click="isLeftSidebarOpen = false; isRightSidebarOpen = true">
                <menu-icon class="h-6 w-6" />
            </button>
        </div>
        <div class="flex items-center justify-center text-xl text-lime-500">
            <div class="text-amber-300 mr-1">
                $
            </div>
            100000
        </div>
    </nav>
    <nav class="nav-right fixed w-80 right-0 top-0 z-30 bg-left-top bg-no-repeat transform transition-transform xl:translate-x-0"
         :class="{'translate-x-0': isRightSidebarOpen, 'translate-x-full': !isRightSidebarOpen}"></nav>
    <section class="sidebar fixed z-20 w-64 right-0 inset-y-0 bg-black bg-opacity-50 bg-left-top bg-repeat-y transform transition-transform xl:translate-x-0"
             :class="{'translate-x-0': isRightSidebarOpen, 'translate-x-full': !isRightSidebarOpen}"></section>
    <!-- eslint-enable max-len -->
    <slot></slot>
</template>

<script>
import {
    MenuIcon,
    ShoppingBagIcon
} from '@heroicons/vue/outline';

import { ref } from 'vue';

export default {
    components: {
        MenuIcon,
        ShoppingBagIcon
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
