<script lang="ts" setup>
import {reactive} from 'vue'
import {GetKeyboardKeys} from '../../wailsjs/go/main/App'
import {keyboard} from "../../wailsjs/go/models";
import Key from "./Key.vue";

interface KeyboardData {
  keys: Array<keyboard.Key> | null;
}

const data: KeyboardData = reactive({
  keys: null
})

setInterval(() => {
  GetKeyboardKeys().then(result => {
    data.keys = result
  })
}, 10)
</script>

<template>
  <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; align-items: center">
    <div class="keyboard-base">
      <key v-for="key in data.keys" :k="key"></key>
    </div>
  </div>
</template>

<style lang="scss">
$columnWidth: 22px;
$keyWidth: ($columnWidth+3) * 2;

.keyboard-base {
  min-width: 700px;
  max-width: 100%;
  padding: 20px;
  border-radius: 10px;
  display: grid;
  grid-template-columns: repeat(38, $columnWidth);
  grid-template-rows: repeat(6, 60px);
  grid-gap: 5px;
  justify-items: start;
}


.key {
  border: 2px solid transparent;
  border-radius: 5px;
  grid-column: span 2;
  font-size: 1em;
  text-align: center;
  padding-top: 17px;
  box-sizing: border-box;
  width: $keyWidth;
  transition: .1s;
  box-shadow: 2px 2px 5px rgba(0,0,0,.4);
}

.key:hover {
  border: 1px solid #eeeeee;
}

.backspace {
  grid-column: span 4;
  width: $keyWidth*2.05;
}

.f1 {
  grid-column: 5 / 7;
}
.f5 {
  grid-column: 14 / 16;
}
.f9 {
  grid-column: 23 / 25;
}
.prt, .ins, .del, .left {
  grid-column: 32 / 34;
}

.tab {
  grid-column: span 3;
  width: $keyWidth*1.5;
}

.\| {
  grid-column: span 3;
  width: $keyWidth*1.50;
}

.caps {
  grid-column: span 4;
  width: $keyWidth*2;
}

.enter {
  grid-column-start: 27;
  grid-column-end: 38;
  justify-self: start;
  width: $keyWidth*2.05;
}

.lshift {
  grid-column: span 5;
  width: $keyWidth*2.5;
}

.rshift {
  grid-column: span 5;
  width: $keyWidth*2.60;
}
.up {
  grid-column: 34 / 38;
  justify-self: start;
  width: $keyWidth;
}

.space {
  grid-column: span 16;
  width: 100%;
}
</style>