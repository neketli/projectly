<template>
    <ElCard>
        <template #header>
            <div class="flex justify-between">
                <h4 class="text-xl">
                    {{ $t('user.teams.title') }}
                </h4>

                <ElButton @click="handleCreateTeam">
                    {{ $t('common.button.create') }}
                </ElButton>
            </div>
        </template>

        <ElTable
            :data="items"
            :loading="isLoading"
            height="250"
            row-class-name="cursor-pointer"
            :empty-text="$t('user.teams.table.empty')"
            @row-click="handleTeamClick"
        >
            <ElTableColumn
                prop="name"
                :label="$t('user.teams.table.name')"
            />

            <ElTableColumn align="right">
                <template #header>
                    <ElInput
                        v-model="search"
                        size="small"
                        :placeholder="$t('common.search')"
                    />
                </template>
            </ElTableColumn>
        </ElTable>

        <ElDialog
            v-model="dialog.team"
            :title="$t('user.teams.dialog.title')"
            align-center
            destroy-on-close
        >
            <TeamForm
                @success="handleTeamCreated"
                @cancel="dialog.team = false"
            />
        </ElDialog>
    </ElCard>
</template>

<script lang="ts" setup>
import type { Team } from '~/types/team'

const { getUserTeams } = useTeam()

const dialog = reactive({
    team: false,
})

const isLoading = ref(false)

const search = ref('')
const teams = ref([] as Team[])

const items = computed(() =>
    !search.value
        ? teams.value
        : teams.value.filter(team => team.name.toLowerCase().includes(search.value.toLowerCase())),
)

const handleCreateTeam = () => {
    dialog.team = true
}

const handleTeamCreated = (team: Team) => {
    teams.value.push(team)
    dialog.team = false
}

const handleTeamClick = ({ id }: Team) => {
    navigateTo(`/team/${id}`)
}

onMounted(async () => {
    try {
        isLoading.value = true
        teams.value = await getUserTeams()
        isLoading.value = false
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
})
</script>
