<template>
    <div class="min-h-screen">
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
                    class="!text-xl"
                >
                    Task tracker
                    <Icon
                        name="mdi:list-box"
                        class="ml-2"
                    />
                </ElMenuItem>

                <LayoutLangSwitcher class="ml-auto mr-2" />
                <LayoutThemeSwitcher class="mr-2" />

                <ElMenuItem
                    index="profile"
                >
                    {{ getUserInfo.name }}
                    {{ getUserInfo.surname }}

                    <UserAvatar
                        :size="32"
                        class="ml-2"
                    />
                </ElMenuItem>

                <ElMenuItem @click="exit">
                    {{ $t('common.logout') }}
                    <Icon
                        class="ml-2"
                        name="mdi:logout"
                    />
                </ElMenuItem>
            </ElMenu>
        </header>

        <div class="h-full text-grey-400 ">
            <div class="h-full w-full flex gap-4">
                <ElScrollbar
                    max-height="92vh"
                    class="w-full"
                >
                    <ElMain>
                        <slot />
                    </ElMain>
                </ElScrollbar>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const route = useRoute()
const { getUserInfo, logout } = useAuthStore()

const exit = () => {
    logout()
    navigateTo('/auth/login')
}
</script>

<style>
.menu .el-menu-item .el-menu-tooltip__trigger,
.menu .el-sub-menu .el-tooltip__trigger {
  justify-content: center;
}
</style>
