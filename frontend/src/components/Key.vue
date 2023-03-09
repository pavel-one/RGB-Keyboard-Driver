<script lang="ts" setup>
import {keyboard} from "../../wailsjs/go/models";

const props = defineProps<{
  k: keyboard.Key
}>()
</script>

<template>
  <div @click="$emit('click')" :style="{background: 'rgb('+k.red+', '+k.green+', '+k.blue+')'}" class="key" :class="k.name.toLowerCase()">
    {{ k.name }}
  </div>
</template>

<style lang="scss">
.key {
  cursor: pointer;
}
</style>

