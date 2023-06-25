<script lang="ts">
export default {
  inheritAttrs: false,
};
</script>

<script lang="ts" setup>
import { ref, toRef } from "vue";
const props = defineProps({ value: String });

const modelValue = toRef(props, "value");
function change(opt: string) {
  modelValue.value = opt;
}
</script>

<template>
  <div class="mood-control">
    <div
      class="verybad"
      @click="change('verybad')"
      :class="modelValue == 'verybad' ? 'selected' : ''"
    >
      <font-awesome-icon icon="fa-solid fa-face-sad-tear" />
    </div>
    <div
      class="bad"
      @click="change('bad')"
      :class="modelValue == 'bad' ? 'selected' : ''"
    >
      <font-awesome-icon icon="fa-solid fa-face-frown" />
    </div>
    <div
      class="neutral"
      @click="change('neutral')"
      :class="modelValue == 'neutral' ? 'selected' : ''"
    >
      <font-awesome-icon icon="fa-solid fa-face-meh" />
    </div>
    <div
      class="good"
      @click="change('good')"
      :class="modelValue == 'good' ? 'selected' : ''"
    >
      <font-awesome-icon icon="fa-solid fa-face-smile" />
    </div>
    <div
      class="verygood"
      @click="change('verygood')"
      :class="modelValue == 'verygood' ? 'selected' : ''"
    >
      <font-awesome-icon icon="fa-solid fa-face-laugh-beam" />
    </div>
    <input
      type="text"
      style="display: none"
      :value="modelValue"
      v-bind="$attrs"
    />
  </div>
</template>

<style lang="scss" scoped>
@use "sass:map";
@import "@/assets/base/theme";
@import "@/assets/base/functions";

.mood-control {
  display: flex;
  justify-content: center;
  width: 100%;

  & div {
    border: $border-width $border-style
      map.get(map.get($themes, "light"), "background");
    border-radius: $border-radius;
    cursor: pointer;
    font-size: $font-size-large;
    padding: $gap-smaller;
    transition: box-shadow $transition, border-color $transition;

    &.selected {
      border-color: get-color("primary", 1);
      box-shadow: 0 0 0 0.15rem #{rgba(get-color("primary", 1), 25%)};
    }

    &.verygood {
      color: $mood-verygood;
    }

    &.good {
      color: $mood-good;
    }

    &.neutral {
      color: $mood-neutral;
    }

    &.bad {
      color: $mood-bad;
    }

    &.verybad {
      color: $mood-verybad;
    }
  }

  & div:not(:first-child) {
    margin-left: $gap;
  }
}
</style>
