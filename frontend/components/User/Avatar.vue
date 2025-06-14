<template>
    <ElAvatar
        :size="size"
        :src="blobUrl"
    >
        <Icon
            name="mdi:account"
            :size="size * 0.75"
        />
    </ElAvatar>
</template>

<script setup>
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

const { getMedia } = useMedia()

const blobUrl = ref(null)

const updateBlobUrl = async () => {
    if (!props.fileName) return

    if (blobUrl.value) {
        URL.revokeObjectURL(blobUrl.value)
        blobUrl.value = null
    }

    const blob = await getMedia(props.fileName)
    blobUrl.value = URL.createObjectURL(blob)
}

watch(() => props.fileName, updateBlobUrl, { immediate: true })

onUnmounted(() => {
    if (blobUrl.value) {
        URL.revokeObjectURL(blobUrl.value)
    }
})
</script>
