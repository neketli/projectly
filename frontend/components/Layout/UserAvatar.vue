<template>
    <ElAvatar
        :size="size"
        :src="avatarUrl"
    >
        <Icon
            name="mdi:account"
            class="text-4xl"
        />
    </ElAvatar>
</template>

<script setup lang="ts">
const { getUserInfo } = toRefs(useAuthStore())

defineProps({
    size: {
        type: Number,
        default: 64,
    },
})

const avatarUrl = computed(() => `http://194.67.84.127/media/${getUserInfo.value.meta.avatar}`)
</script>
