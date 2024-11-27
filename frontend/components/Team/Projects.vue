<template>
    <div>
        <ElCard>
            <template #header>
                <div class="flex justify-between">
                    <h3 class="text-xl">
                        {{ $t('team.projects.title') }}
                    </h3>

                    <ElButton
                        v-if="teamStore.isEditAvailable"
                        @click="dialog.project = true"
                    >
                        <Icon name="mdi:plus" />
                    </ElButton>
                </div>
            </template>

            <ElTable
                :data="projectItems"
                :loading="isLoading"
                height="250"
                row-class-name="cursor-pointer"
                :empty-text="$t('common.no_data')"
                @row-click="handleProjectClick"
            >
                <ElTableColumn
                    prop="code"
                    width="100"
                    :label="$t('team.projects.table.code')"
                />
                <ElTableColumn
                    prop="title"
                    width="150"
                    :label="$t('team.projects.table.title')"
                />

                <ElTableColumn
                    width="150"
                    align="right"
                >
                    <template #header>
                        <ElInput
                            v-model.trim="search"
                            size="small"
                            :placeholder="$t('common.search')"
                        />
                    </template>
                </ElTableColumn>
            </ElTable>
        </ElCard>

        <ElDialog
            v-model="dialog.project"
            :title="$t('project.create.title')"
            align-center
            destroy-on-close
        >
            <ProjectForm
                :team-id="teamId"
                @cancel="dialog.project = false"
                @success="handleProjectCreated"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { Project } from '~/types/project'

const props = defineProps<{ teamId: number }>()

const teamStore = useTeamStore()
const { getProjectsList } = useProjects()

const projects: Ref<Project[]> = ref([])
const search = ref('')

const isLoading = ref(false)

const dialog = reactive({
    project: false,
})

const projectItems = computed(() =>
    !search.value
        ? projects.value
        : projects.value.filter(item =>
            item.title.toLowerCase().includes(search.value.toLowerCase())
            || item.code.toLowerCase().includes(search.value.toLowerCase()),
        ),
)

const handleProjectClick = async (project: Project) => {
    navigateTo(`/team/${props.teamId}/project/${project.code}`)
}

const handleProjectCreated = (project: Project) => {
    projects.value.push(project)
    dialog.project = false
}

onMounted(async () => {
    isLoading.value = true
    try {
        projects.value = await getProjectsList({ team_id: props.teamId })
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
