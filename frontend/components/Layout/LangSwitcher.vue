<template>
    <ElSelect
        v-model="lang"
        value-key="code"
        class="!w-28"
        @change="handleSwitchLang"
    >
        <ElOption
            v-for="item in locales"
            :key="item.code"
            :value="item"
            :label="item.label"
        >
            {{ item.label }}
            <Icon :name="`flag:${item.flag}-4x3`" />
        </ElOption>
    </ElSelect>
</template>

<script setup lang="ts">
const { locale, locales, setLocale } = useI18n()

const lang = ref(locales.value.find(item => item.code === locale.value))

const handleSwitchLang = () => {
    setLocale(lang.value.code)
}
</script>
