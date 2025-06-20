<template>
    <header>
        <ElMenu
            mode="horizontal"
            :default-active="route.path"
            router
            class="flew items-center w-full"
            :ellipsis="false"
        >
            <ElMenuItem
                index="/"
                :class="{ 'md:!mr-auto': !timerStore.settings.headerWidget }"
                class="max-md:!mr-auto"
            >
                <LayoutIcon />
            </ElMenuItem>

            <ElMenuItem
                v-if="timerStore.settings.headerWidget"
                index="/#timer"
                class="!mr-auto !hidden md:!flex"
            >
                <LayoutWidgetsTimer />
            </ElMenuItem>

            <div class="hidden md:flex gap-2 mr-2">
                <LayoutLangSwitcher />
                <LayoutThemeSwitcher />
            </div>

            <ElMenuItem index="/profile">
                <span class="hidden md:inline">
                    {{ getUserInfo.name }}
                    {{ getUserInfo.surname }}
                </span>

                <UserAvatar
                    :size="32"
                    :file-name="getUserInfo.meta?.avatar"
                    class="md:ml-2"
                />
            </ElMenuItem>

            <ElMenuItem
                v-if="isLogged"
                class="!hidden md:!inline-flex"
                @click="exit"
            >
                {{ $t('common.header.logout') }}

                <Icon
                    class="ml-2"
                    name="mdi:logout"
                />
            </ElMenuItem>

            <ElButton
                class="inline-flex md:!hidden mx-2"
                circle
                @click="isMenuOpen = !isMenuOpen"
            >
                <Icon :name="isMenuOpen ? 'mdi:close' : 'mdi:menu'" />
            </ElButton>
        </ElMenu>

        <ElDrawer
            v-model="isMenuOpen"
            direction="btt"
            size="50%"
        >
            <div class="flex flex-col gap-8">
                <div class="flex items-center gap-2">
                    <span>{{ $t('common.header.lang') }}</span>
                    <LayoutLangSwitcher />
                </div>

                <div class="flex items-center gap-2">
                    <span>{{ $t('common.header.theme') }}</span>
                    <LayoutThemeSwitcher />
                </div>

                <LayoutWidgetsTimer />

                <ElButton
                    v-if="isLogged"
                    @click="exit"
                >
                    {{ $t('common.header.logout') }}

                    <Icon
                        class="ml-2"
                        name="mdi:logout"
                    />
                </ElButton>
            </div>
        </ElDrawer>
    </header>
</template>

<script setup lang="ts">
const route = useRoute()

const timerStore = useTimerStore()
const authStore = useAuthStore()
const { getUserInfo, isLogged } = toRefs(authStore)

const isMenuOpen = ref(false)

const exit = () => {
    authStore.logout()
    navigateTo('/auth/login')
}
</script>

<style>
.menu .el-menu-item .el-menu-tooltip__trigger,
.menu .el-sub-menu .el-tooltip__trigger {
  justify-content: center;
}
</style>
