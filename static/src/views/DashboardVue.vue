<script lang="ts" setup>
import { ref, type Ref } from "vue";
import SiteMood from "@/components/SiteMood.vue";

interface Mood {
  Date: string;
  Global: string;
  Sleep: string;
  Anxiety: string;
  DarkThoughts: string;
  Ruminations: string;
  Notes: string;
}

const now = ref(new Date());

const first = ref(new Date());
first.value.setDate(0);

now.value.setMonth(now.value.getMonth() + 1);
now.value.setDate(0);

const displayPopup = ref(false);
const currentDay = ref(new Date());

const today = new Date();
const results: Ref<Mood[]> = ref(new Array(32).fill({ Global: "" }));

const global = ref("");
const sleep = ref("");
const anxiety = ref("");
const darkthoughts = ref("");
const ruminations = ref("");
const notes = ref("");

function prevMonth() {
  let year = now.value.getFullYear();
  let month = now.value.getMonth();
  if (month == 0) {
    year--;
    month = 11;
  } else {
    month--;
  }
  now.value = new Date(year, month + 1, 0);
  first.value = new Date(year, month, 0);
  get();
}

function nextMonth() {
  let year = now.value.getFullYear();
  let month = now.value.getMonth();
  if (month == 11) {
    year++;
    month = 0;
  } else {
    month++;
  }
  now.value = new Date(year, month + 1, 0);
  first.value = new Date(year, month, 0);
  get();
}

function open(day: number) {
  currentDay.value = new Date(
    now.value.getFullYear(),
    now.value.getMonth(),
    day
  );
  notes.value = results.value[currentDay.value.getDate()].Notes;
  global.value = results.value[currentDay.value.getDate()].Global;
  sleep.value = results.value[currentDay.value.getDate()].Sleep;
  anxiety.value = results.value[currentDay.value.getDate()].Anxiety;
  darkthoughts.value = results.value[currentDay.value.getDate()].DarkThoughts;
  ruminations.value = results.value[currentDay.value.getDate()].Ruminations;
  displayPopup.value = true;
}

function moodClass(day: number): string {
  let c = "";

  if (
    today.getDate() === day &&
    today.getMonth() == now.value.getMonth() &&
    today.getFullYear() == now.value.getFullYear()
  )
    c += "today ";
  c += results.value[day].Global;
  return c;
}

async function submit(e: Event) {
  e.preventDefault();
  const formData = new FormData(e.target as HTMLFormElement);
  formData.set("date", currentDay.value.toISOString());

  let global = formData.get("global");
  let sleep = formData.get("sleep");
  let anxiety = formData.get("anxiety");
  let darkthoughts = formData.get("darkthoughts");
  let ruminations = formData.get("ruminations");
  let notes = formData.get("notes");
  if (!global) global = "";
  if (!sleep) sleep = "";
  if (!anxiety) anxiety = "";
  if (!darkthoughts) darkthoughts = "";
  if (!ruminations) ruminations = "";
  if (!notes) notes = "";

  const mood: Mood = {
    Date: currentDay.value.toString(),
    Global: global.toString(),
    Sleep: sleep.toString(),
    Anxiety: anxiety.toString(),
    DarkThoughts: darkthoughts.toString(),
    Ruminations: ruminations.toString(),
    Notes: notes.toString(),
  };

  results.value[currentDay.value.getDate()] = mood;

  const f = await fetch("./api/update", {
    method: "POST",
    body: formData,
  });
  if (f.status !== 200) alert("error while updating the mood");
  displayPopup.value = false;
  return;
}

async function get() {
  const formData = new FormData();
  const from = new Date(now.value);
  from.setDate(0);
  formData.set("from", from.toISOString());
  formData.set("to", now.value.toISOString());
  const f = await fetch("./api/get", {
    method: "POST",
    body: formData,
  });
  if (f.status !== 200) {
    alert("error while updating the mood");
    return;
  }
  const json = await f.json();
  results.value = new Array(32).fill({ Global: "" });
  json.forEach((element: Mood) => {
    const date = new Date(element.Date);
    results.value[date.getDate()] = element;
  });

  return;
}

get();
</script>

<template>
  <div
    id="popup"
    :style="displayPopup ? 'visibility:visible' : 'visibility:hidden'"
  >
    <div class="background" @click="displayPopup = false"></div>
    <div class="content">
      <h1 class="title">{{ $d(currentDay, "day") }}</h1>
      <form @submit="submit">
        <div class="field">
          <label class="label">{{ $t("dashboard.moodavg") }}</label>
          <div class="control">
            <SiteMood name="global" :value="global"></SiteMood>
          </div>
        </div>
        <div class="field">
          <label class="label">{{ $t("dashboard.sleep") }}</label>
          <div class="control">
            <SiteMood name="sleep" :value="sleep"></SiteMood>
          </div>
        </div>
        <div class="field">
          <label class="label">{{ $t("dashboard.anxiety") }}</label>
          <div class="control">
            <SiteMood name="anxiety" :value="anxiety"></SiteMood>
          </div>
        </div>
        <div class="field">
          <label class="label">{{ $t("dashboard.darkthoughts") }}</label>
          <div class="control">
            <SiteMood name="darkthoughts" :value="darkthoughts"></SiteMood>
          </div>
        </div>
        <div class="field">
          <label class="label">{{ $t("dashboard.ruminations") }}</label>
          <div class="control">
            <SiteMood name="ruminations" :value="ruminations"></SiteMood>
          </div>
        </div>
        <div class="field">
          <label for="notes" class="label">{{ $t("dashboard.notes") }}</label>
          <div class="control">
            <textarea
              id="notes"
              name="notes"
              type="text"
              :value="notes"
              rows="5"
              cols="30"
              class="input is-light"
            />
          </div>
        </div>
        <div class="field">
          <div class="control is-not-full-width">
            <button type="submit" class="input button">
              {{ $t("dashboard.submit") }}
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
  <h1 class="title">{{ $t("dashboard.title") }}</h1>
  <main>
    <h1 class="title month">
      <button
        class="input button"
        :title="$t('dashboard.prevMonth')"
        @click="prevMonth"
      >
        &lt;
      </button>
      <span>{{ $d(now, "month") }}</span>
      <button
        class="input button"
        :title="$t('dashboard.nextMonth')"
        @click="nextMonth"
      >
        >
      </button>
    </h1>

    <div
      class="mood-container"
      v-for="i in Math.ceil((now.getDate() + first.getDay()) / 7)"
      :key="i"
    >
      <template v-if="i == 1">
        <button
          class="input button mood"
          v-for="j in first.getDay()"
          :key="j"
        ></button>
        <button
          class="input button mood"
          :class="moodClass(j + 7 * (i - 1))"
          @click="open(j + 7 * (i - 1))"
          v-for="j in 7 - first.getDay()"
          :key="j"
        >
          {{ j + 7 * (i - 1) }}
        </button>
      </template>
      <template
        v-else-if="i != Math.ceil((now.getDate() + first.getDay()) / 7)"
      >
        <button
          class="input button mood"
          :class="moodClass(j + 7 * (i - 1) - first.getDay())"
          @click="open(j + 7 * (i - 1) - first.getDay())"
          v-for="j in 7"
          :key="j"
        >
          {{ j + 7 * (i - 1) - first.getDay() }}
        </button>
      </template>
      <template v-else>
        <button
          class="input button mood"
          v-for="j in now.getDate() - (i - 1) * 7 + first.getDay()"
          :class="results[j + 7 * (i - 1) - first.getDay()].Global"
          @click="open(j + 7 * (i - 1) - first.getDay())"
          :key="j"
        >
          {{ j + 7 * (i - 1) - first.getDay() }}
        </button>
        <button
          class="input button mood"
          v-for="j in now.getDay() === 0 ? 0 : 7 - now.getDay()"
          :key="j"
        ></button>
      </template>
    </div>
  </main>
</template>

<style lang="scss" scoped>
@use "sass:map";
@import "@/assets/base/theme";
@import "@/assets/base/functions";
@import "@/assets/base/mixins";

.month {
  align-items: center;
  display: flex;
  justify-content: space-between;
}

.mood-container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.mood-container:not(:first-child) {
  margin-top: $gap;
}

.mood {
  background-color: var(--c-background);
  border: $border-width dashed $primary;
  color: var(--c-text);
  width: 100%;

  &.today {
    border: $border-width $border-style $primary;
    box-shadow: 0 0 0 0.15rem #{rgba(get-color("primary", 1), 25%)} !important;
  }

  &.verygood {
    background-color: $mood-verygood !important;
    border-color: $mood-verygood !important;
    color: $dark;
  }

  &.good {
    background-color: $mood-good !important;
    border-color: $mood-good !important;
    color: $dark;
  }

  &.neutral {
    background-color: $mood-neutral !important;
    border-color: $mood-neutral !important;
    color: $dark;
  }

  &.bad {
    background-color: $mood-bad !important;
    border-color: $mood-bad !important;
    color: $dark;
  }

  &.verybad {
    background-color: $mood-verybad !important;
    border-color: $mood-verybad !important;
    color: $dark;
  }

  @include until($mobile) {
    padding-left: 0 !important;
    padding-right: 0 !important;
  }
}

.mood:not(:first-child) {
  @include from($mobile) {
    margin-left: $gap;
  }
}

#popup {
  bottom: 0;
  color: map.get(map.get($themes, "light"), "text");
  display: flex;
  left: 0;
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1;

  & > .background {
    background-color: rgb(0 0 0 / 50%);
    inset: 0;
    position: absolute;
  }

  & > .content {
    background-color: map.get(map.get($themes, "light"), "background");
    border-radius: $border-radius;
    margin: auto;
    max-height: 100%;
    overflow: auto;
    padding: $gap;
    z-index: 2;
  }

  & h1 {
    border-bottom: 1px solid $dark;
  }
}
</style>
