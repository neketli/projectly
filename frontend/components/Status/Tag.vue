<template>
    <div
        class="text-xl px-4 py-2 rounded flex justify-center items-center"
        :style="{
            'background-color': `${color}20`,
            'color': color,
        }"
    >
        <slot />
    </div>
</template>

<script lang="ts" setup>
withDefaults(
    defineProps<{ color: string }>(),
    { color: '#409EFF' },
)
</script>
