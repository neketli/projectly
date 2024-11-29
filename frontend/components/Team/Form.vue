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
            <ElInput v-model="form[item.value]">
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
                @click="handleSaveTeam"
                @keyup.enter="handleSaveTeam"
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
import type { Team } from '~/types/team'

const props = defineProps<{
    team?: Team
}>()

const emit = defineEmits(['success', 'cancel'])

const { t } = useI18n()
const { createTeam, updateTeam } = useTeam()
const validators = useValidator()

const isLoading = ref(false)

const formElement = ref<FormInstance>()
const form = ref({
    name: props.team?.name || '',
    description: props.team?.description || '',
})
const rules = reactive<FormRules<typeof form.value>>({
    name: [
        validators.required,
        validators.len(),
    ],
    description: [
        validators.len(0, 255),
    ],
})

const items: {
    label: string
    value: keyof typeof form.value
    icon: string
}[] = [
    {
        label: t('team.form.name'),
        value: 'name',
        icon: 'mdi:account-group-outline',
    },
    {
        label: t('team.form.description'),
        value: 'description',
        icon: 'mdi:file-document-outline',
    },
]

const handleSaveTeam = async () => {
    if (!formElement.value) {
        return
    }
    await formElement.value.validate(async (valid) => {
        if (valid) {
            await saveTeam()
        }
    })
}

const saveTeam = async () => {
    isLoading.value = true
    try {
        let team
        if (props.team) {
            team = await updateTeam({
                ...props.team,
                ...form.value,
            })
        }
        else {
            team = await createTeam(form.value)
        }

        emit('success', team)
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
