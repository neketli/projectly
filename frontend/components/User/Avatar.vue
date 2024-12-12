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
const config = useRuntimeConfig()

const props = defineProps({
    fileName: {
        type: String,
        default: '',
    },
    size: {
        type: Number,
        default: 64,
    },
})

const avatarUrl = computed(() => props.fileName ? `${config.public.S3_URL}/${props.fileName}` : '')
</script>
