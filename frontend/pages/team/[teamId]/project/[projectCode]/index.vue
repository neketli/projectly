<template>
    <div class="project">
        <ElPageHeader @back="navigateTo(`/team/${teamId}`)">
            <template #content>
                <h1 class="text-2xl">
                    {{ $t('project.title') }}
                </h1>
            </template>

            <template #title>
                <h3>
                    {{ $t('team.title') }}
                </h3>
            </template>

            <template #extra>
                <ElPopconfirm
                    v-if="projectStore.isDeleteAvailable"
                    width="250"
                    :title="$t('project.delete.confirm')"
                    :confirm-button-text="$t('common.button.confirm')"
                    :cancel-button-text="$t('common.button.cancel')"
                    confirm-button-type="danger"
                    @confirm="handleDeleteProject"
                >
                    <template #reference>
                        <ElButton
                            :title="$t('project.delete.title')"
                            circle
                            plain
                            type="danger"
                        >
                            <Icon name="mdi:delete" />
                        </ElButton>
                    </template>
                </ElPopconfirm>

                <ElButton
                    v-if="projectStore.isEditAvailable"
                    :title="$t('project.update.title')"
                    circle
                    plain
                    @click="dialog.project = true"
                >
                    <Icon name="mdi:pencil" />
                </ElButton>
            </template>
        </ElPageHeader>

        <h3 class="text-4xl mt-2">
            {{ project.title }}
        </h3>

        <p class="mt-2">
            {{ project.description }}
        </p>

        <BoardList
            v-if="project.id"
            :project-id="project.id"
            class="mt-8 max-w-lg mx-auto"
        />

        <ElDialog
            v-model="dialog.project"
            :title="$t('project.update.title')"
            align-center
            destroy-on-close
        >
            <ProjectForm
                :project="project"
                :team-id="Number(teamId)"
                @success="handleProjectUpdated"
                @cancel="dialog.project = false"
            />
        </ElDialog>
    </div>
</template>

<script lang="ts" setup>
import type { Project } from '~/types/project'

const { teamId, projectCode } = useRoute().params
const { t } = useI18n()

useHead({
    title: t('project.title'),
})

definePageMeta({
    layout: 'team',
})

const projectStore = useProjectStore()
const { project } = toRefs(projectStore)
const { getTeamProject, deleteProject } = useProjects()

const dialog = reactive({
    project: false,
})

const isLoading = ref(false)

const handleProjectUpdated = (updated: Project) => {
    project.value = updated
    dialog.project = false
}

const handleDeleteProject = async () => {
    try {
        await deleteProject(project.value.id)
        ElMessage.success(t('project.success.delete'))
        navigateTo(`/team/${teamId}`)
    }
    catch (err) {
        const error = err as Error
        ElMessage.error(error.message)
    }
}

onMounted(async () => {
    try {
        isLoading.value = true
        project.value = await getTeamProject(Number(teamId), `${projectCode}`)
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
