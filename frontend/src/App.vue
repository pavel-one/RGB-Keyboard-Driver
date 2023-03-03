<script lang="ts" setup>
import Keyboard from "./components/Keyboard.vue";
import {darkTheme} from 'naive-ui'
import {reactive} from "vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";

let data = reactive({
  theme: darkTheme as any,
  background: '#151515'
})

function changeTheme(dark: boolean) {
  if (dark) {
    data.theme = darkTheme
    data.background = '#151515'
    return
  }

  data.theme = null
  data.background = '#f5f5f5'
}

</script>

<template>
  <n-config-provider :style="{background: data.background, position: 'relative'}" :theme="data.theme">
    <theme-switcher @change="changeTheme"></theme-switcher>
    <keyboard></keyboard>

  </n-config-provider>
</template>

