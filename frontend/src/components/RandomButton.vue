<template>
  <n-button type="primary" quaternary circle class="random-button" @click="Reload()">
    <n-icon size="40">
      <ReloadCircleSharp />
    </n-icon>
  </n-button>
</template>

<script lang="ts" setup>
import { ReloadCircleSharp } from '@vicons/ionicons5'
import {Reload} from "../../wailsjs/go/main/App";
</script>

<style scoped>
.random-button {
  position: absolute;
  right: 5%;
  bottom: 5%;
  transition: .25s;
  opacity: .8;
}
.random-button:hover {
  opacity: 1;
}
</style>