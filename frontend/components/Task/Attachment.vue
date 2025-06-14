<template>
    <div v-loading="isLoading">
        <ElPopover v-if="variant === 'compact'">
            <template #reference>
                <ElImage
                    v-if="isImageFile(fileName)"
                    :src="blobUrl"
                    :title="fileName"
                    :preview-src-list="[blobUrl]"
                    :initial-index="0"
                    hide-on-click-modal
                    class="w-24 h-24 !flex items-center justify-center  border border-gray-100 dark:border-gray-800 p-1 rounded"
                    fit="cover"
                >
                    <template #error>
                        <Icon
                            size="32"
                            name="mdi:file-image"
                        />
                    </template>
                </ElImage>
                <a
                    v-else
                    class="w-24 h-24 flex items-center justify-center rounded bg-gray-100 dark:bg-gray-800"
                    :title="fileName"
                    @click="copyFilePath"
                >
                    <Icon
                        size="32"
                        :name="extMdiMap[getFileExtension(fileName)] || 'mdi:file-outline'"
                    />
                </a>
            </template>

            <div class="flex justify-center">
                <ElButton
                    :href="blobUrl"
                    :title="fileName"
                    :download="fileName"
                    circle
                    tag="a"
                >
                    <Icon
                        name="mdi:download"
                    />
                </ElButton>

                <ElButton
                    circle
                    @click.stop="copyFilePath"
                >
                    <Icon
                        name="mdi:content-copy"
                    />
                </ElButton>

                <ElButton
                    circle
                    type="danger"
                    @click.stop="emit('delete', fileName)"
                >
                    <Icon
                        name="mdi:delete"
                    />
                </ElButton>
            </div>
        </ElPopover>
        <div
            v-else
            class="flex flex-col justify-center items-center"
        >
            <h4 class="text-xl mb-4">
                {{ fileName.split('/').at(-1) }}
            </h4>
            <ElImage
                v-if="isImageFile(fileName)"
                :src="blobUrl"
                :title="fileName"
                :preview-src-list="[blobUrl]"
                :initial-index="0"
                hide-on-click-modal
                class="w-2/3 !flex items-center justify-center  border border-gray-100 dark:border-gray-800 p-1 rounded"
                fit="cover"
            >
                <template #error>
                    <Icon
                        size="32"
                        name="mdi:file-image"
                    />
                </template>
            </ElImage>
            <a
                v-else
                class="w-48 h-48 flex items-center justify-center rounded bg-gray-100 dark:bg-gray-800"
                :title="fileName"
                @click="copyFilePath"
            >
                <Icon
                    size="32"
                    :name="extMdiMap[getFileExtension(fileName)] || 'mdi:file-outline'"
                />
            </a>

            <div class="mt-4 flex justify-center">
                <ElButton
                    :href="blobUrl"
                    :title="fileName"
                    :download="fileName"
                    tag="a"
                >
                    {{ t('attachment.download') }}
                    <Icon
                        name="mdi:download"
                        class="ml-2"
                    />
                </ElButton>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
const props = withDefaults(
    defineProps<{
        fileName: string
        variant?: 'compact' | 'full'
    }>(),
    {
        variant: 'compact',
    },
)

const emit = defineEmits<{
    (event: 'delete', fileName: string): void
}>()

const { t } = useI18n()
const { copy } = useClipboard()
const { getMedia } = useMedia()

const blobUrl = ref<string>('')
const isLoading = ref(false)

const extMdiMap: Record<string, string> = {
    'doc': 'mdi:file-document-outline',
    'docx': 'mdi:file-document-outline',
    'txt': 'mdi:file-document-outline',
    'json': 'mdi:file-document-outline',
    'pdf': 'mdi:file-pdf-outline',
    'xls': 'mdi:file-excel-outline',
    'xlsx': 'mdi:file-excel-outline',
    'ppt': 'mdi:file-powerpoint-outline',
    'pptx': 'mdi:file-powerpoint-outline',
    'zip': 'mdi:folder-zip-outline',
    'rar': 'mdi:folder-zip-outline',
    'mp3': 'mdi:file-music-outline',
    'mp4': 'mdi:file-video-outline',
    'avi': 'mdi:file-video-outline',
    'mov': 'mdi:file-video-outline',
    'wmv': 'mdi:file-video-outline',
    '': 'mdi:file-outline',
}

const getFileExtension = (fileName: string) => {
    const ext = fileName.split('.').pop()
    return ext ? ext.toLowerCase() : ''
}

const isImageFile = (fileName: string) => {
    const imageExtensions = ['png', 'jpg', 'jpeg', 'gif', 'svg']
    const ext = getFileExtension(fileName)
    return imageExtensions.includes(ext)
}

const updateBlobUrl = async () => {
    isLoading.value = true
    try {
        if (!props.fileName) return

        if (blobUrl.value) {
            URL.revokeObjectURL(blobUrl.value)
            blobUrl.value = ''
        }

        const blob = await getMedia(props.fileName)
        blobUrl.value = URL.createObjectURL(blob)
    }
    finally {
        isLoading.value = false
    }
}

const copyFilePath = () => {
    copy(`${window.location.origin}/${props.fileName}`)
    ElMessage.success(t('attachment.copied'))
}

watch(() => props.fileName, updateBlobUrl, { immediate: true })

onUnmounted(() => {
    if (blobUrl.value) {
        URL.revokeObjectURL(blobUrl.value)
    }
})
</script>
