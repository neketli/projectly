<template>
    <ElForm
        ref="formElement"
        :rules="rules"
        :model="form"
        :loading="isLoading"
        autocomplete="off"
        label-position="top"
        @submit.prevent
    >
        <ElFormItem
            :label="$t('task.form.title')"
            prop="title"
            class="w-full"
        >
            <ElInput v-model="form.title" />
        </ElFormItem>

        <ElFormItem
            :label="$t('task.form.assigned_user')"
            prop="assigned_user_id"
            class="w-full"
        >
            <ElSelect
                v-model="form.assigned_user_id"
                :placeholder="$t('task.form.assigned_user')"
                clearable
            >
                <ElOption
                    v-for="item in usersOptions"
                    :key="item.id"
                    :value="item.id"
                    :label="item.label"
                    class="w-full flex gap-2 items-center"
                >
                    <UserAvatar
                        :size="24"
                        :file-name="item.avatar"
                    />
                    {{ item.label }}
                </ElOption>
            </ElSelect>
        </ElFormItem>

        <ElFormItem
            :label="$t('task.form.description')"
            prop="description"
            class="task-description w-full"
        >
            <ElMention
                v-model="form.description"
                :options="usersOptions"
                type="textarea"
                maxlength="4096"
                resize="none"
                show-word-limit
                :autosize="{
                    minRows: 3,
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
        </ElFormItem>

        <div class="w-full flex flex-wrap gap-8">
            <ElFormItem
                :label="$t('task.form.story_point')"
                prop="story_points"
            >
                <ElInputNumber
                    v-model="form.story_points"
                    :min="0"
                    :step="1"
                    step-strictly
                />
            </ElFormItem>

            <ElFormItem
                :label="$t('task.form.tracked_time')"
                prop="tracked_time"
            >
                <ElInputNumber
                    v-model="form.tracked_time"
                    :min="0"
                    :step="1"
                    step-strictly
                />
            </ElFormItem>
        </div>

        <ElFormItem
            :label="$t('task.form.priority.label')"
            prop="priority"
            class="w-full"
        >
            <ElSelect v-model="form.priority">
                <ElOption
                    v-for="(item, idx) in priorityOptions"
                    :key="idx"
                    :value="idx"
                    :label="item"
                >
                    {{ item }}
                </ElOption>
            </ElSelect>
        </ElFormItem>

        <ElFormItem
            :label="$t('task.form.deadline')"
            prop="deadline"
            class="w-full"
        >
            <ElDatePicker
                v-model="form.deadline"
                type="date"
                class="!w-full"
                value-format="X"
            />
        </ElFormItem>

        <div>
            <ElButton
                :disabled="isLoading"
                type="primary"
                plain
                @click="handleSaveProject"
            >
                {{ $t('common.button.save') }}
            </ElButton>
            <ElButton
                type="danger"
                plain
                class="max-w-56"
                @click="emit('cancel')"
            >
                {{ $t('common.button.cancel') }}
            </ElButton>
        </div>
    </ElForm>
</template>

<script lang="ts" setup>
import dayjs from 'dayjs'
import type { FormInstance, FormRules } from 'element-plus'
import type { DetailedTask } from '~/types/task'

const props = defineProps<{
    task?: DetailedTask
    statusId: number
    isFinished?: boolean
}>()

const emit = defineEmits(['success', 'cancel'])

const { t } = useI18n()
const validators = useValidator()

const { users } = useTeamStore()
const { createTask, updateTask } = useTask()

const isLoading = ref(false)

const formElement = ref<FormInstance>()
const form = ref({
    title: props.task?.title || '',
    description: props.task?.description || '',
    priority: props.task?.priority || 0,
    story_points: props.task?.story_points || 0,
    tracked_time: props.task?.tracked_time || 0,
    deadline: props.task?.deadline || '',
    assigned_user_id: props.task?.assigned_user_id || '',
})
const rules = reactive<FormRules<typeof form.value>>({
    title: [
        validators.required,
        validators.len(),
    ],
    description: [
        validators.len(0, 4096),
    ],
    priority: [
        validators.range(0, 3),
    ],
    story_points: [
        validators.range(0, 100),
    ],
})

const usersOptions = computed(() => {
    return users.map(user => ({
        id: user.id,
        value: `${user.email}`,
        label: `${user.name} ${user.surname} (${user.email})`,
        avatar: user?.meta?.avatar,
    }))
})

const priorityOptions = computed(() => {
    return [
        t('task.form.priority.none'),
        '⚡',
        '⚡⚡',
        '⚡⚡⚡',
    ]
})

const handleSaveProject = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            saveTask()
        }
    })
}

const saveTask = async () => {
    isLoading.value = true

    try {
        const details = {
            ...form.value,
            assigned_user_id: Number(form.value.assigned_user_id) || 0,
            deadline: Number(form.value.deadline) || 0,
            status_id: props.statusId,
        }

        const task = props.task?.id
            ? await updateTask({
                ...props.task,
                ...details,
            })
            : await createTask({
                ...details,
                finished_at: props.isFinished ? dayjs().unix() : undefined,
            })

        emit('success', task)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}
</script>

<style>
.task-description .el-mention {
    width: 100%;
}
</style>
