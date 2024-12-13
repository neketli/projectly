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
            v-for="item in items"
            :key="item.value"
            :label="item.label"
            :prop="item.value"
            class="w-full"
        >
            <ElInput
                v-model="form[item.value]"
                :disabled="item.isDisabled"
            >
                <template #prefix>
                    <Icon :name="item.icon" />
                </template>
            </ElInput>
        </ElFormItem>

        <div>
            <ElButton
                :disabled="isLoading"
                type="primary"
                plain
                @click="handleSaveProject"
                @keyup.enter="handleSaveProject"
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
import type { FormInstance, FormRules } from 'element-plus'
import type { Project } from '~/types/project'

const props = defineProps<{
    project?: Project
    teamId: number
}>()

const emit = defineEmits(['success', 'cancel'])

const { t } = useI18n()
const { createProject, updateProject } = useProjects()
const validators = useValidator()

const isLoading = ref(false)

const formElement = ref<FormInstance>()
const form = ref({
    title: props.project?.title || '',
    description: props.project?.description || '',
    code: props.project?.code || '',
})
const rules = reactive<FormRules<typeof form.value>>({
    title: [
        validators.required,
        validators.len(),
    ],
    description: [
        validators.len(0, 128),
    ],
    code: [
        validators.required,
        validators.len(1, 5),
        validators.alpha,
    ],
})

const items: ComputedRef<{
    label: string
    value: keyof typeof form.value
    icon: string
    isDisabled?: boolean
}[]> = computed(() => [
    {
        label: t('project.form.title'),
        value: 'title',
        icon: 'mdi:bookmark',
    },
    {
        label: t('project.form.description'),
        value: 'description',
        icon: 'mdi:subtitles',
    },
    {
        label: t('project.form.code'),
        value: 'code',
        icon: 'mdi:pound-box',
        isDisabled: Boolean(props.project?.code),
    },
])

const handleSaveProject = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            await saveProject()
        }
    })
}

const saveProject = async () => {
    isLoading.value = true
    try {
        let project
        if (props.project) {
            project = await updateProject({
                ...props.project,
                ...form.value,
            })
        }
        else {
            project = await createProject({ ...form.value, team_id: props.teamId })
        }

        emit('success', project)
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
