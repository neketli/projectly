<template>
    <div class="bg-neutral-100 dark:bg-neutral-900 dark:ring-2 dark:ring-neutral-700 min-w-[22vw] rounded-md p-2 relative">
        <div class="flex justify-between items-center gap-4 p-2 sticky top-0">
            <template v-if="!isEdit">
                <div
                    class="flex items-center gap-4"
                >
                    <StatusTag :color="status.hex_color">
                        {{ status.title }}
                    </StatusTag>

                    <p class="text-gray-600 dark:text-gray-500 text-sm">
                        <!-- TODO: task counter -->
                    </p>
                </div>

                <div class="flex gap-2 items-center text-gray-300">
                    <ElDropdown trigger="click">
                        <ElButton circle>
                            <Icon name="mdi:dots-horizontal" />
                        </ElButton>

                        <template #dropdown>
                            <ElDropdownItem @click="handleEditStatus">
                                <Icon
                                    name="mdi:pencil"
                                    class="mr-2"
                                />
                                {{ $t('common.button.edit') }}
                            </ElDropdownItem>
                            <ElDropdownItem @click="dialog.delete = true">
                                <Icon
                                    name="mdi:delete"
                                    class="mr-2"
                                />
                                {{ $t('common.button.delete') }}
                            </ElDropdownItem>

                            <ElDropdownItem
                                :disabled="status.order === statusCount - 1"
                                @click="handleMove('right')"
                            >
                                <Icon
                                    name="mdi:arrow-right"
                                    class="mr-2"
                                />
                                {{ $t('status.form.right') }}
                            </ElDropdownItem>

                            <ElDropdownItem
                                :disabled="status.order === 0"
                                @click="handleMove('left')"
                            >
                                <Icon
                                    name="mdi:arrow-left"
                                    class="mr-2"
                                />
                                {{ $t('status.form.left') }}
                            </ElDropdownItem>
                        </template>
                    </ElDropdown>

                    <ElButton
                        :title="$t('status.create.task')"
                        type="primary"
                        circle
                    >
                        <Icon name="mdi:plus" />
                    </ElButton>
                </div>
            </template>
            <template v-else>
                <div class="w-full flex flex-col gap-2">
                    <label>
                        {{ $t('status.form.title') }}
                    </label>
                    <div class="w-full flex gap-2">
                        <ElInput
                            v-model="title"
                            @keyup.enter="handleSaveStatus"
                        />
                        <ElColorPicker
                            v-model="color"
                            :predefine="defaultStatusColors"
                        />
                    </div>
                </div>

                <div class="flex self-end">
                    <ElButton
                        circle
                        @click="handleCancel"
                    >
                        <Icon name="mdi:close" />
                    </ElButton>
                    <ElButton
                        circle
                        type="success"
                        :disabled="!title || !color"
                        @click="handleSaveStatus"
                    >
                        <Icon name="mdi:content-save" />
                    </ElButton>
                </div>
            </template>
        </div>

        <!-- <ElScrollbar
            class="w-full"
        /> -->

        <ElDialog
            v-model="dialog.delete"
            :title="$t('status.delete.confirm')"
            align-center
            destroy-on-close
            class="!w-1/3"
        >
            <ElButton
                @click="dialog.delete = false"
            >
                {{ $t('common.button.cancel') }}
            </ElButton>
            <ElButton
                type="danger"
                @click="emit('delete', status)"
            >
                {{ $t('common.button.confirm') }}
            </ElButton>
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import { defaultStatusColors, type Status } from '~/types/board'

const props = defineProps<{ status: Status }>()
const emit = defineEmits(['create', 'update', 'delete'])

const { statusCount } = useBoardStore()

const isEdit = ref(false)
const dialog = reactive({
    delete: false,
})
const title = ref('')
const color = ref(props.status.hex_color || defaultStatusColors[0])

const handleEditStatus = () => {
    isEdit.value = true
    title.value = props.status.title || ''
}

const handleCancel = () => {
    isEdit.value = false

    if (props.status.id === 0) {
        emit('delete', props.status)
    }
}

const handleSaveStatus = () => {
    isEdit.value = false

    emit(props.status.id === 0 ? 'create' : 'update', {
        ...props.status,
        title: title.value,
        hex_color: color.value,
    })

    title.value = ''
}

const handleMove = (direction: 'left' | 'right') => {
    if ((direction === 'left' && props.status.order === 0)
      || (direction === 'right' && props.status.order === statusCount - 1)) return

    emit('update', {
        ...props.status,
        order: direction === 'left'
            ? props.status.order - 1
            : props.status.order + 1,
    })
}

onMounted(() => {
    if (props.status.id === 0) {
        isEdit.value = true
        title.value = props.status.title || ''
    }
})
</script>
