<script lang="ts" setup>
import {reactive} from "vue";
import { DarkModeFilled, LightModeRound } from '@vicons/material'
import { ReloadCircleSharp } from '@vicons/ionicons5'
import {Reload} from "../../wailsjs/go/main/App";

const data = reactive({
  active: false,
})
</script>

<template>
  <div class="theme-switcher">
    <n-space>
      <n-switch @updateValue="$emit('change', data.active)" class="theme-switcher" v-model:value="data.active" size="large">
        <template #checked-icon>

          <n-icon :component="DarkModeFilled" />
        </template>
        <template #unchecked-icon>
          <n-icon :component="LightModeRound" />
        </template>
      </n-switch>
      <n-button @click="Reload()">
        <n-icon size="30">
          <ReloadCircleSharp />
        </n-icon>
      </n-button>
    </n-space>
  </div>
</template>

<style scoped>
.theme-switcher {
  position: absolute;
  right: 10vh;
  bottom: 10vh;
}
</style>