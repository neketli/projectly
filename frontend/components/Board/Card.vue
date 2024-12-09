<template>
    <div class="relative bg-neutral-100 dark:bg-slate-800 dark:shadow-none rounded-lg shadow-md py-2 px-4">
        <span class="text-lg">
            {{ board.title }}
        </span>
    </div>
</template>

<script lang="ts" setup>
import type { Board } from '~/types/board'

defineProps<{ board: Board }>()
</script>
