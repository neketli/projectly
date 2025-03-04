<template>
    <ElPopover
        width="305"
        trigger="click"
    >
        <template #reference>
            <ElButton
                text
                circle
            >
                <Icon
                    name="mdi:emoticon"
                    class="text-xl"
                />
            </ElButton>
        </template>

        <Picker
            :data="emojiIndex"
            :show-skin-tones="false"
            :native="true"
            :show-search="false"
            :emoji-tooltip="false"
            :show-preview="false"
            :show-categories="true"
            theme="dark"
            emoji="dart"
            title="Projectly"
            per-line="7"
            color="#409EFF"
            auto-focus
            @select="handleSelectEmoji"
        />
    </ElPopover>
</template>

<script setup>
import data from 'emoji-mart-vue-fast/data/all.json'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import { Picker, EmojiIndex } from 'emoji-mart-vue-fast/src'

const emojiIndex = new EmojiIndex(data)

const emit = defineEmits(['select'])

const handleSelectEmoji = (emoji) => {
    emit('select', emoji.native)
}
</script>
