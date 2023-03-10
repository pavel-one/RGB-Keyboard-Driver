<script lang="ts" setup>
import {keyboard} from "../../wailsjs/go/models";
import { computed } from 'vue'

const props = defineProps<{
  k: keyboard.Key
}>()

const background = computed<string>(() => {
  if ((props.k.red < 20) && (props.k.green < 20) && (props.k.blue < 20)) {
    return 'rgba('+props.k.red+', '+props.k.green+', '+props.k.blue+', 0)'
  }

  return 'rgba('+props.k.red+', '+props.k.green+', '+props.k.blue+', 1)'
})
</script>

<template>
  <div @click="$emit('click')" :style="{background: background}" class="key" :class="k.name.toLowerCase()">
    {{ k.name }}
  </div>
</template>

<style lang="scss">
.key {
  cursor: pointer;
}
</style>

