<template>
    <span>
        <NuxtLink
            v-if="user?.email"
            class="flex gap-2 items-center cursor-pointer hover:text-blue-500 transition-all"
            :to="{ name: 'user-email', params: { email: user.email } }"
            @click.stop
        >
            <UserAvatar
                v-if="user?.avatar"
                :size="32"
                :file-name="user.avatar"
                class="ml-2 min-w-8 h-8"
            />

            <span class="inline">
                {{ user.name }}
                {{ user.surname }}
                ({{ user.email }})
            </span>
        </NuxtLink>

        <ElTag
            v-else
            type="info"
        >
            -
        </ElTag>
    </span>
</template>

<script setup lang="ts">
import type { PlainUser } from '~/types/user'

defineProps<{ user: PlainUser }>()
</script>
