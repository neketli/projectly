<template>
    <div class="!text-2xl flex items-center gap-2 font-['Helvetica']">
        Projectly

        <Logo />

        <ElTooltip :content="$t('common.beta')">
            <ElTag
                round
                effect="plain"
            >
                <Icon name="mdi:beta" />
            </ElTag>
        </ElTooltip>
    </div>
</template>

<script setup lang="ts">
import Logo from '@/assets/icons/logo.svg'
</script>
