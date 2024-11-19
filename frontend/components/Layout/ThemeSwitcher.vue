<template>
    <div>
        <ElButton
            :title="
                isDark ? $t('common.theme.light') : $t('common.theme.dark')"
            @click="handleSwitchTheme"
        >
            <Icon :name="icon" />
        </ElButton>
    </div>
</template>

<script setup lang="ts">
const colorMode = useColorMode()

const isDark = computed(() => colorMode.preference === 'dark')
const icon = computed(() => isDark.value ? 'mdi:lightbulb-on' : 'mdi:lightbulb-off')

const handleSwitchTheme = () => {
    colorMode.preference = isDark.value ? 'light' : 'dark'
}
</script>
