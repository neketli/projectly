<template>
    <div class="task">
        <ElPageHeader @back="navigateTo(`/team/${teamId}/project/${projectCode}/${boardId}`)">
            <template #content>
                <h2 class="text-2xl">
                    {{ task.project_code }}-{{ task.project_index }}
                </h2>
            </template>

            <template #title>
                <h3>
                    {{ $t('board.title') }}
                </h3>
            </template>

            <template #extra>
                <ElPopconfirm
                    v-if="projectStore.isDeleteAvailable"
                    width="250"
                    :title="$t('task.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteTask"
                >
                    <template #reference>
                        <ElButton
                            :title="$t('task.delete.title')"
                            circle
                            plain
                            type="danger"
                        >
                            <Icon name="mdi:delete" />
                        </ElButton>
                    </template>
                </ElPopconfirm>

                <ElButton
                    v-if="projectStore.isEditAvailable"
                    :title="$t('task.update.title')"
                    circle
                    plain
                    @click="dialog.task = true"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>
            </template>
        </ElPageHeader>

        <div
            v-loading="isLoading"
            class="container mx-auto"
        >
            <h1 class="text-4xl my-4">
                {{ task.title }}
            </h1>

            <TaskDetail
                :task="task"
                @update-status="handleStatusChange"
            />

            <div
                class="prose dark:prose-invert my-4"
                v-html="taskDescriptionMd"
            />

            <h4 class="text-2xl border-t pt-2">
                {{ t('task.form.attachment.title') }}
            </h4>
            <ElUpload
                v-loading="isLoading"
                :auto-upload="false"
                :show-file-list="false"
                drag
                multiple
                class="my-2"
                :on-change="handleAddAttachment"
            >
                <Icon name="mdi:cloud-upload-outline" />
                {{ $t('task.form.attachment.add') }}
            </ElUpload>

            <div class="flex flex-wrap gap-4">
                <ElPopover
                    v-for="(attachment, index) in attachments"
                    :key="attachment"
                >
                    <template #reference>
                        <ElImage
                            v-if="isImageFile(attachment)"
                            :src="attachment"
                            :title="attachment"
                            :preview-src-list="[attachment]"
                            :initial-index="0"
                            hide-on-click-modal
                            class="w-24 h-24 border border-gray-100 dark:border-gray-800 p-1 rounded"
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
                            :href="attachment"
                            class="w-24 h-24 flex items-center justify-center rounded bg-gray-100 dark:bg-gray-800"
                            :title="attachment"
                        >
                            <Icon
                                size="32"
                                :name="extMdiMap[getFileExtension(attachment)] || 'mdi:file-outline'"
                            />
                        </a>
                    </template>

                    <div class="flex justify-center">
                        <ElButton
                            :href="attachment"
                            :title="attachment"
                            circle
                            tag="a"
                        >
                            <Icon
                                name="mdi:download"
                            />
                        </ElButton>

                        <ElButton
                            class="cursor-pointer"
                            circle
                            type="danger"
                            @click.stop="handleDeleteAttachment(files[index])"
                        >
                            <Icon
                                name="mdi:delete"
                            />
                        </ElButton>
                    </div>
                </ElPopover>
            </div>
        </div>

        <ElDialog
            v-model="dialog.task"
            :title="$t('task.update.title')"
            align-center
            destroy-on-close
        >
            <TaskForm
                :task="task"
                :status-id="task.status_id"
                @success="handleTaskUpdated"
                @cancel="dialog.task = false"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import markdownit from 'markdown-it'
import type { UploadProps } from 'element-plus'
import dayjs from 'dayjs'
import type { DetailedTask } from '~/types/task'

const config = useRuntimeConfig()
const { teamId, projectCode, boardId, projectIndex } = useRoute().params
const { t } = useI18n()
const md = markdownit()

useHead({
    title: t('task.title'),
})

definePageMeta({
    layout: 'team',
})

const projectStore = useProjectStore()
const boardStore = useBoardStore()

const { getTaskDetail, deleteTask, updateTaskStatus, createAttachments, deleteAttachment, getAttachments } = useTask()
const { getStatusList } = useStatus()

const dialog = reactive({
    task: false,
})

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

const task = ref({} as DetailedTask)
const files = ref([] as string[])
const isLoading = ref(false)

const taskDescriptionMd = computed(() => task.value.description ? md.render(task.value.description) : t('task.form.placeholder.description'))

const attachments = computed(() => files.value.map(item => `${config.public.S3_HOST}/media/${item}`))

const handleTaskUpdated = (updated: DetailedTask) => {
    task.value = updated
    dialog.task = false
}

const handleDeleteTask = async () => {
    try {
        await deleteTask(task.value.id)
        ElMessage.success(t('task.success.delete'))
        navigateTo(`/team/${teamId}/project/${projectCode}/${boardId}`)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleAddAttachment: UploadProps['onChange'] = async (uploadFile) => {
    const MAX_SIZE_MB = 30
    if (uploadFile.size && (uploadFile.size / 1024 / 1024 > MAX_SIZE_MB)) {
        ElMessage.error(t('task.error.file_size', { x: MAX_SIZE_MB }))
        return false
    }

    const [attachment] = await createAttachments(task.value.id, uploadFile.raw as File)
    files.value.push(attachment)

    return true
}

const handleDeleteAttachment = async (filename: string) => {
    try {
        await deleteAttachment(filename)
        files.value = files.value.filter(item => item !== filename)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleStatusChange = async (statusId: number) => {
    try {
        if (task.value.status_id === statusId) return
        if (!statusId) {
            ElMessage.error(t('task.error.status_undefined'))
            return
        }

        isLoading.value = true
        const finishedAt = boardStore.finishStatus.id === statusId ? dayjs().unix() : null

        isLoading.value = true
        await updateTaskStatus(task.value.id, statusId, finishedAt)
        task.value.status_id = statusId
        task.value.status = boardStore.statusList.find(item => item.id === statusId) || task.value.status
        if (finishedAt) {
            task.value.finished_at = finishedAt
        }
        ElMessage.success(t('task.success.update'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

onMounted(async () => {
    try {
        isLoading.value = true
        task.value = await getTaskDetail(`${projectCode}`, Number(projectIndex))
        const statuses = await getStatusList(Number(boardId))
        boardStore.setStatusList(statuses)

        const taskAttachments = await getAttachments(task.value.id)
        files.value = taskAttachments
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
})
</script>
