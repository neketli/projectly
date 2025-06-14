<template>
    <div class=" flex justify-center items-center">
        <TaskAttachment
            :file-name="fileName"
            variant="full"
        />
    </div>
</template>

<script setup lang="ts">
const route = useRoute()
const { filePath } = route.params

const path = computed(() => Array.isArray(filePath) ? filePath.join('/') : filePath)

const fileName = computed(() => `attachments/${path.value}`)
</script>
