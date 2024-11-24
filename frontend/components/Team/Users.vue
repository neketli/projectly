<template>
    <ElCard class="mt-4">
        <template #header>
            <div class="flex justify-between">
                <h3 class="text-xl">
                    {{ $t('team.users.title') }}
                </h3>

                <ElButton
                    v-if="teamStore.isInviteAvailable"
                    @click="dialog.user = true"
                >
                    <Icon name="mdi:account-plus" />
                </ElButton>
            </div>
        </template>

        <ElTable :data="users">
            <ElTableColumn
                prop="name"
                width="150"
                :label="$t('team.users.table.name')"
            />
            <ElTableColumn
                prop="surname"
                width="150"
                :label="$t('team.users.table.surname')"
            />
            <ElTableColumn
                prop="email"
                :label="$t('team.users.table.email')"
            />

            <ElTableColumn
                prop="role"
                :label="$t('team.users.table.role')"
            >
                <template #default="{ row }">
                    <ElSelect
                        v-if="teamStore.isUsersActionsAvailable"
                        v-model="row.role.id"
                        @change="(newRoleId) => handleUpdateRole(row.id, newRoleId)"
                    >
                        <ElOption
                            v-for="option in roleNames"
                            :key="option.value"
                            :label="option.label"
                            :value="option.value"
                        />
                    </ElSelect>
                    <div v-else>
                        {{ $t(RoleLabels[row.role.role_name as UserRole]) }}
                    </div>
                </template>
            </ElTableColumn>

            <ElTableColumn
                v-if="teamStore.isUsersActionsAvailable"
                align="right"
                :label="$t('team.users.table.actions')"
            >
                <template #default="{ row }">
                    <ElPopconfirm
                        width="250"
                        :title="$t('team.users.table.confirm_delete')"
                        :confirm-button-text="$t('common.button.confirm')"
                        confirm-button-type="danger"
                        :cancel-button-text="$t('common.button.cancel')"
                        @confirm="handleConfirmRemoveUser(row.id)"
                    >
                        <template #reference>
                            <ElButton
                                circle
                                type="danger"
                            >
                                <Icon name="mdi:close" />
                            </ElButton>
                        </template>
                    </ElPopconfirm>
                </template>
            </ElTableColumn>
        </ElTable>

        <ElDialog
            v-model="dialog.user"
            :title="$t('team.users.dialog.title')"
            align-center
            destroy-on-close
        >
            <span>{{ $t('team.users.dialog.email') }}</span>
            <ElInput
                v-model="newUserEmail"
                class="mb-4"
            >
                <template #prefix>
                    <Icon name="mdi:email" />
                </template>
            </ElInput>

            <ElButton
                plain
                @click="handleAddUser"
            >
                {{ $t('team.users.dialog.invite') }}
            </ElButton>

            <ElButton
                type="danger"
                plain
                @click="handleCancelAddUser"
            >
                {{ $t('common.button.cancel') }}
            </ElButton>
        </ElDialog>
    </ElCard>
</template>

<script lang="ts" setup>
import { RoleLabels, type Role, type UserRole } from '~/types/user'

const props = defineProps<{ teamId: number }>()
const { t } = useI18n()

const teamStore = useTeamStore()
const { users } = toRefs(teamStore)

const { getTeamUsers, addTeamUser, removeTeamUser, getRoles, setRole } = useTeam()

const roles = ref([] as Role[])
const newUserEmail = ref('')
const isLoading = ref(false)

const dialog = reactive({
    user: false,
})

const roleNames = computed(() => {
    return roles.value.map(r => ({
        value: r.id,
        label: t(RoleLabels[r.role_name as UserRole]),
    }))
})

const handleConfirmRemoveUser = async (id: number) => {
    isLoading.value = true
    try {
        await removeTeamUser(props.teamId, id)
        users.value = users.value.filter(u => u.id !== id)
        ElMessage.success(t('team.success.users.remove'))
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
    finally {
        isLoading.value = false
    }
}

const handleAddUser = async () => {
    try {
        isLoading.value = true
        await addTeamUser(props.teamId, newUserEmail.value)
        users.value = await getTeamUsers(props.teamId)
        dialog.user = false
        ElMessage.success(t('team.success.users.invite'))
    }
    catch (err) {
        const _ = err as Error
        ElMessage.error(t('team.error.users.invite'))
    }
    finally {
        isLoading.value = false
    }
}

const handleCancelAddUser = () => {
    newUserEmail.value = ''
    dialog.user = false
}

const handleUpdateRole = async (userId: number, newRoleId: number) => {
    try {
        isLoading.value = true
        await setRole(props.teamId, userId, newRoleId)
        users.value = await getTeamUsers(props.teamId)
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
        users.value = await getTeamUsers(props.teamId)
        roles.value = await getRoles()
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
