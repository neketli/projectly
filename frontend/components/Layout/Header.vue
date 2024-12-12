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
                class="!mr-auto"
            >
                <LayoutIcon />
            </ElMenuItem>

            <div class="hidden sm:flex gap-2 mr-2">
                <LayoutLangSwitcher />
                <LayoutThemeSwitcher />
            </div>

            <ElMenuItem index="/profile">
                <span class="hidden sm:inline">
                    {{ getUserInfo.name }}
                    {{ getUserInfo.surname }}
                </span>

                <UserAvatar
                    :size="32"
                    :file-name="getUserInfo.meta?.avatar"
                    class="sm:ml-2"
                />
            </ElMenuItem>

            <ElMenuItem
                v-if="getUserInfo.id"
                class="!hidden sm:!inline-flex"
                @click="exit"
            >
                {{ $t('common.header.logout') }}

                <Icon
                    class="ml-2"
                    name="mdi:logout"
                />
            </ElMenuItem>

            <ElButton
                class="inline-flex sm:!hidden mx-2"
                circle
                @click="isMenuOpen = !isMenuOpen"
            >
                <Icon :name="isMenuOpen ? 'mdi:close' : 'mdi:menu'" />
            </ElButton>
        </ElMenu>

        <ElDrawer
            v-model="isMenuOpen"
            direction="btt"
        >
            <div class="flex flex-col gap-8">
                <div class="flex justify-center gap-8">
                    <div class="flex gap-2">
                        <span>{{ $t('common.header.lang') }}</span>
                        <LayoutLangSwitcher />
                    </div>

                    <div class="flex gap-2">
                        <span>{{ $t('common.header.theme') }}</span>
                        <LayoutThemeSwitcher />
                    </div>
                </div>

                <ElButton
                    v-if="getUserInfo.id"
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
const { getUserInfo, logout } = useAuthStore()

const isMenuOpen = ref(false)

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
