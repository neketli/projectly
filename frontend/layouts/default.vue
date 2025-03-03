<template>
    <div class="min-h-screen flex flex-col">
        <LayoutHeader />

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
        <LayoutFooter />
    </div>
</template>
