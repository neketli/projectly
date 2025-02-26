<template>
    <div class="flex items-start gap-2 p-2 border border-gray-200 dark:border-gray-700 bg-gray-100 dark:bg-gray-800 bg-opacity-80 rounded-lg">
        <UserAvatar
            :size="32"
            :file-name="comment.user?.avatar"
            class="mr-2"
        />
        <div class="flex flex-col gap-2">
            <div class="flex items-center gap-2 text-sm text-gray-700 dark:text-gray-300">
                <span class="font-bold">
                    {{ comment.user?.name }} {{ comment.user?.surname }} ({{ comment.user?.email }})
                </span>
                <span class="text-gray-500 font-light">
                    {{ dayjs.unix(comment.updated_at).format('DD MMM YYYY, HH:mm') }}
                </span>

                <ElPopconfirm
                    v-if="comment.user.id === authStore.user.id"
                    width="250"
                    :title="$t('task.comment.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteComment"
                >
                    <template #reference>
                        <ElButton
                            size="small"
                            class="ml-auto"
                            circle
                        >
                            <Icon name="mdi:delete" />
                        </ElButton>
                    </template>
                </ElPopconfirm>
            </div>
            <div
                class="prose dark:prose-invert"
                v-html="commentText"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import markdownit from 'markdown-it'
import dayjs from 'dayjs'
import { computed } from 'vue'
import type { TaskComment } from '~/types/task'

const props = defineProps<{
    comment: TaskComment
}>()

const authStore = useAuthStore()

const md = markdownit()

const commentText = computed(() => md.render(props.comment.text))

const emit = defineEmits<{
    (event: 'delete', commentId: number): void
}>()

const handleDeleteComment = () => {
    emit('delete', props.comment.id)
}
</script>
