<template>
    <ElAvatar
        :size="size"
        :src="avatarUrl"
    >
        <Icon
            name="mdi:account"
            :size="size * 0.75"
        />
    </ElAvatar>
</template>

<script setup>
const { getUserInfo } = toRefs(useAuthStore())

defineProps({
    size: {
        type: Number,
        default: 64,
    },
})

// TODO: вынести обращения к S3
const avatarUrl = computed(() => `http://194.67.84.127/media/${getUserInfo.value?.meta?.avatar}`)
</script>
