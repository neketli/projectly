<template>
    <div class="min-h-screen">
        <LayoutHeader />

        <div class="h-full text-grey-400 ">
            <div class="h-full w-full flex gap-4">
                <aside class="hidden lg:block">
                    <BoardMenu
                        v-if="projectStore.project.id"
                        :project-id="projectStore.project.id"
                        class="h-full min-h-[92vh]"
                    />
                </aside>

                <ElScrollbar
                    max-height="92vh"
                    class="w-full"
                >
                    <ElMain>
                        <div class="mb-2 flex items-center justify-between">
                            <ElBreadcrumb separator="/">
                                <ElBreadcrumbItem :to="{ path: '/' }">
                                    {{ $t('dashboard.title') }}
                                </ElBreadcrumbItem>

                                <ElBreadcrumbItem
                                    v-if="teamStore.team.id"
                                    :to="{ path: `/team/${teamStore.team.id}` }"
                                >
                                    <span class="text-blue-500">{{ teamStore.team.name }}</span>
                                </ElBreadcrumbItem>

                                <ElBreadcrumbItem
                                    v-if="projectStore.project?.code"
                                    :to="{ path: `/team/${teamStore.team.id}/project/${projectStore.project.code}` }"
                                >
                                    <span class="text-blue-500">{{ projectStore.project.title }}</span>
                                </ElBreadcrumbItem>

                                <ElBreadcrumbItem
                                    v-if="route.params.boardId"
                                    :to="{
                                        // eslint-disable-next-line @stylistic/max-len
                                        path: `/team/${teamStore.team.id}/project/${projectStore.project.code}/${route.params.boardId}`,
                                    }"
                                >
                                    <span class="text-blue-500">{{ $t('board.title') }} {{ route.params.boardId }}</span>
                                </ElBreadcrumbItem>

                                <ElBreadcrumbItem
                                    v-if="route.params.projectIndex"
                                    :to="{
                                        // eslint-disable-next-line @stylistic/max-len
                                        path: `/team/${teamStore.team.id}/project/${projectStore.project.code}/${route.params.boardId}/task/${route.params.projectCode}-${route.params.projectIndex}`,
                                    }"
                                >
                                    <span class="text-blue-500">{{ route.params.projectCode }}-{{ route.params.projectIndex }}</span>
                                </ElBreadcrumbItem>
                            </ElBreadcrumb>

                            <ElButton
                                class="lg:!hidden"
                                @click="handleToggleMenu"
                            >
                                {{ $t('board.list.title') }}
                            </ElButton>
                        </div>

                        <slot />
                    </ElMain>
                </ElScrollbar>
            </div>
        </div>

        <ElDrawer
            v-model="projectStore.boardsDrawer"
            size="80%"
            direction="btt"
        >
            <BoardMenu
                v-if="projectStore.project.id"
                :project-id="projectStore.project.id"
                class="h-full"
                @board-clicked="projectStore.boardsDrawer = false"
            />
        </ElDrawer>
    </div>
</template>

<script setup lang="ts">
const route = useRoute()

const teamStore = useTeamStore()
const projectStore = useProjectStore()

const { getTeamUsers, getTeam } = useTeam()
const { getTeamProject } = useProjects()

const updateStore = async () => {
    if (route.params.teamId) {
        if (!teamStore.team?.id || teamStore.team.id !== Number(route.params.teamId)) {
            try {
                teamStore.team = await getTeam(Number(route.params.teamId))
            }
            catch (err) {
                const error = err as Error
                ElMessage.error(error.message)
                navigateTo('/404')
            }
        }
        if (!teamStore.users.length) {
            try {
                teamStore.users = await getTeamUsers(Number(route.params.teamId))
            }
            catch (err) {
                const error = err as Error
                ElMessage.error(error.message)
                navigateTo('/404')
            }
        }
    }
    else {
        teamStore.$reset()
        return
    }

    if (route.params.projectCode) {
        if (!projectStore.project?.id) {
            try {
                projectStore.project = await getTeamProject(Number(route.params.teamId), `${route.params.projectCode}`)
            }
            catch (err) {
                const error = err as Error
                ElMessage.error(error.message)
            }
        }
    }
    else {
        projectStore.$reset()
    }
}

const handleToggleMenu = () => {
    projectStore.boardsDrawer = !projectStore.boardsDrawer
}

watch(() => route.params, updateStore)

onMounted(() => updateStore())
</script>

<style>
.menu .el-menu-item .el-menu-tooltip__trigger,
.menu .el-sub-menu .el-tooltip__trigger {
  justify-content: center;
}
</style>
