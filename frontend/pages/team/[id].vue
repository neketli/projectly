<template>
    <div class="team container mx-auto">
        <ElPageHeader @back="navigateTo('/')">
            <template #content>
                <h1 class="text-2xl">
                    {{ $t('team.title') }}
                </h1>
            </template>

            <template #title>
                <h3>
                    {{ $t('dashboard.title') }}
                </h3>
            </template>

            <template #extra>
                <ElPopconfirm
                    v-if="teamStore.isDeleteAvailable"
                    width="250"
                    :title="$t('team.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteTeam"
                >
                    <template #reference>
                        <ElButton
                            :title="$t('team.delete.title')"
                            circle
                            plain
                            type="danger"
                        >
                            <Icon name="mdi:delete" />
                        </ElButton>
                    </template>
                </ElPopconfirm>

                <ElButton
                    v-if="teamStore.isEditAvailable"
                    :title="$t('team.update.title')"
                    circle
                    plain
                    @click="dialog.team = true"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>

                <ElPopconfirm
                    v-if="teamStore.canLeave"
                    width="250"
                    :title="$t('team.leave.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleLeaveTeam"
                >
                    <template #reference>
                        <ElButton
                            :title="$t('team.leave.title')"
                            circle
                            plain
                            type="warning"
                        >
                            <Icon name="mdi:logout-variant" />
                        </ElButton>
                    </template>
                </ElPopconfirm>
            </template>
        </ElPageHeader>

        <h3 class="text-4xl mt-2">
            {{ $t('team.title') }} - {{ team.name }}
        </h3>

        <p class="mt-2">
            {{ team.description }}
        </p>

        <TeamUsers
            v-if="team.id"
            :team-id="team.id"
        />

        <ElDialog
            v-model="dialog.team"
            :title="$t('team.update.title')"
            align-center
            destroy-on-close
        >
            <TeamForm
                :team="team"
                @success="handleTeamUpdated"
                @cancel="dialog.team = false"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { Team } from '~/types/team'

const { params } = useRoute()
const { t } = useI18n()

useHead({
    title: t('team.title'),
})

const { getTeam, deleteTeam, removeTeamUser } = useTeam()
const { getUserInfo } = useAuthStore()
const teamStore = useTeamStore()

const { team } = toRefs(teamStore)

const dialog = reactive({
    team: false,
})

const handleDeleteTeam = async () => {
    try {
        await deleteTeam(team.value.id)
        ElMessage.success(t('team.success.delete'))
        navigateTo('/')
    }
    catch (err) {
        const _ = err as Error
        ElMessage.error(t('team.error.delete'))
    }
}

const handleLeaveTeam = async () => {
    try {
        await removeTeamUser(team.value.id, getUserInfo.id)
        navigateTo('/')
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

const handleTeamUpdated = (updated: Team) => {
    team.value = updated
    dialog.team = false
}

onMounted(async () => {
    try {
        team.value = await getTeam(+params.id)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
})

onUnmounted(() => {
    teamStore.$reset()
})
</script>
