<script lang="ts" setup>
import {reactive} from "vue";
import { DarkModeFilled, LightModeRound } from '@vicons/material'

const data = reactive({
  active: true,
})
</script>

<template>
  <div class="theme-switcher">
    <n-switch @updateValue="$emit('change', data.active)" class="theme-switcher" v-model:value="data.active" size="large">
      <template #checked-icon>

        <n-icon :component="DarkModeFilled" />
      </template>
      <template #unchecked-icon>
        <n-icon :component="LightModeRound" />
      </template>
    </n-switch>
  </div>
</template>

<style scoped>
.theme-switcher {
  position: absolute;
  right: 9%;
  bottom: 5.5%;
  transition: .25s;
  opacity: .8;
}
.theme-switcher:hover {
  opacity: 1;
}
</style>