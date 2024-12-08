<template>
    <div class="team">
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
            {{ team.name }}
        </h3>

        <p class="mt-2">
            {{ team.description }}
        </p>

        <div
            v-if="team.id"
            class="flex mt-8 gap-8"
        >
            <TeamUsers
                :team-id="team.id"
                class="w-2/3"
            />

            <ProjectTable
                :team-id="team.id"
                class="w-1/3"
            />
        </div>

        <ElCollapse
            v-if="team.id"
            v-model="collapse"
            accordion
            class="mt-8"
            @change="hasStatistic = true"
        >
            <ElCollapseItem
                :title="$t('team.statistic.title')"
                name="statistic"
            >
                <KeepAlive>
                    <TeamStatistic v-if="hasStatistic" />
                </KeepAlive>
            </ElCollapseItem>
        </ElCollapse>

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
import type { Project } from '~/types/project'
import type { Team } from '~/types/team'

const { params } = useRoute()
const { t } = useI18n()

useHead({
    title: t('team.title'),
})

definePageMeta({
    layout: 'team',
})

const { getUserInfo } = useAuthStore()
const teamStore = useTeamStore()

const { getTeam, deleteTeam, removeTeamUser } = useTeam()
const { getProjectsList } = useProjects()

const { team } = toRefs(teamStore)

const projects = ref<Project[]>([])

const collapse = ref('')
const hasStatistic = ref(false)
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
        team.value = await getTeam(Number(params.teamId))
        projects.value = await getProjectsList({ team_id: team.value.id })
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
})
</script>
