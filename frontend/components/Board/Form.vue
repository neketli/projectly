<template>
    <ElForm
        ref="formElement"
        :rules="rules"
        :model="form"
        :loading="isLoading"
        class="pt-4"
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
                @click="handleSaveBoard"
                @keyup.enter="handleSaveBoard"
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
import type { Board } from '~/types/board'

const props = defineProps<{
    board?: Board
    projectId: number
}>()

const emit = defineEmits(['success', 'cancel'])

const { t } = useI18n()
const { createBoard, updateBoard } = useBoard()
const validators = useValidator()

const isLoading = ref(false)

const formElement = ref<FormInstance>()
const form = ref({
    title: props.board?.title || '',
})
const rules = reactive<FormRules<typeof form.value>>({
    title: [
        validators.required,
        validators.len(),
    ],
})

const items: ComputedRef<{
    label: string
    value: keyof typeof form.value
    icon: string
    isDisabled?: boolean
}[]> = computed(() => [
    {
        label: t('board.form.title'),
        value: 'title',
        icon: 'mdi:label',
    },
])

const handleSaveBoard = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            await saveBoard()
        }
    })
}

const saveBoard = async () => {
    isLoading.value = true
    try {
        let board
        if (props.board) {
            board = await updateBoard({
                ...props.board,
                ...form.value,
            })
        }
        else {
            board = await createBoard({ ...form.value, project_id: props.projectId })
        }

        emit('success', board)
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
