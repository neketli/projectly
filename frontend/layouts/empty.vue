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
                <ElMenuItem index="/">
                    <LayoutIcon />
                </ElMenuItem>

                <LayoutLangSwitcher class="ml-auto mr-2" />
                <LayoutThemeSwitcher class="mr-2" />
            </ElMenu>
        </header>

        <div class="h-full text-grey-400 ">
            <div class="h-full w-full flex gap-4">
                <slot />
            </div>
        </div>

        <LayoutFooter />
    </div>
</template>

<script setup lang="ts">
const route = useRoute()
</script>

<style>
.menu .el-menu-item .el-menu-tooltip__trigger,
.menu .el-sub-menu .el-tooltip__trigger {
  justify-content: center;
}
</style>
