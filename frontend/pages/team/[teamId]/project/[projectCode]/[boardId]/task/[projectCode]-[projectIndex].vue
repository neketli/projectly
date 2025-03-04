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

            <ElDivider
                content-position="left"
            >
                <h4 class="text-xl">
                    {{ t('task.form.attachment.title') }}
                </h4>
            </ElDivider>

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

            <div
                ref="attachmentsParent"
                class="flex flex-wrap gap-4"
            >
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

            <ElDivider content-position="left">
                <h4 class="text-xl">
                    {{ t('task.comments.title') }}
                </h4>
            </ElDivider>
            <div
                ref="commentsParent"
                v-loading="isCommentsLoading"
                class="flex flex-col gap-2"
            >
                <ElAlert
                    v-if="comments.length === 0"
                    :closable="false"
                    :title="$t('task.comments.empty')"
                />

                <TaskComment
                    v-for="item in comments"
                    :key="item.id"
                    :comment="item"
                    @delete="handleDeleteComment"
                />

                <div class="flex gap-2">
                    <ElMention
                        v-model="comment"
                        :options="usersOptions"
                        type="textarea"
                        maxlength="1024"
                        resize="none"
                        :placeholder="t('task.comments.placeholder')"
                        prefix=":"
                        show-word-limit
                        :autosize="{
                            minRows: 2,
                            maxRows: 6,
                        }"
                        class="!w-full"
                    >
                        <template #label="{ item }">
                            <div style="display: flex; align-items: center">
                                <UserAvatar
                                    :size="24"
                                    :file-name="item.avatar"
                                />
                                <span style="margin-left: 6px">{{ item.label }}</span>
                            </div>
                        </template>
                    </ElMention>

                    <ElButton
                        circle
                        @click="handleCreateComment"
                    >
                        <Icon name="mdi:send" />
                    </ElButton>
                </div>
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
import type { DetailedTask, TaskComment } from '~/types/task'

const config = useRuntimeConfig()
const { teamId, projectCode, boardId, projectIndex } = useRoute().params
const { t } = useI18n()
const [attachmentsParent, commentsParent] = useAutoAnimate()
const md = markdownit()

useHead({
    title: t('task.title'),
})

definePageMeta({
    layout: 'project',
})

const teamStore = useTeamStore()
const projectStore = useProjectStore()
const boardStore = useBoardStore()

const {
    getTaskDetail,
    deleteTask,
    updateTaskStatus,
    createAttachments,
    deleteAttachment,
    getAttachments,
    getComments,
    createComment,
    deleteComment,
} = useTask()
const { getStatusList } = useStatus()

const dialog = reactive({
    task: false,
})

const comment = ref('')

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
const comments = ref([] as TaskComment[])
const isLoading = ref(false)
const isCommentsLoading = ref(false)
let commentsIntervalId = 0

const taskDescriptionMd = computed(() => task.value.description ? md.render(task.value.description) : t('task.form.placeholder.description'))

const attachments = computed(() => files.value.map(item => `${config.public.S3_HOST}/media/${item}`))

const usersOptions = computed(() => {
    return teamStore.users.map(user => ({
        id: user.id,
        value: `${user.email}`,
        label: `${user.name} ${user.surname} (${user.email})`,
        avatar: user?.meta?.avatar,
    }))
})

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

const handleCreateComment = async () => {
    try {
        await createComment(task.value.id, comment.value)
        comment.value = ''
        updateComments()
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleDeleteComment = async (commentId: number) => {
    try {
        await deleteComment(task.value.id, commentId)
        comments.value = comments.value.filter(item => item.id !== commentId)
        await updateComments()
        ElMessage.success(t('task.success.comment_delete'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const updateComments = async () => {
    try {
        isCommentsLoading.value = true
        const lastCommentId = comments.value.at(-1)?.id || 0
        const newComments = await getComments(task.value.id, lastCommentId)
        comments.value = [...comments.value, ...newComments]
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
        clearInterval(commentsIntervalId)
    }
    finally {
        isCommentsLoading.value = false
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

        updateComments()
        commentsIntervalId = window.setInterval(updateComments, 30000)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
})

onBeforeUnmount(() => {
    clearInterval(commentsIntervalId)
})
</script>
