<script setup lang="ts">
import { Chart, type ChartItem } from "chart.js/auto";
import "chartjs-adapter-date-fns";
import { onMounted } from "vue";

interface Mood {
  Date: string;
  Global: string;
  Sleep: string;
  Anxiety: string;
  DarkThoughts: string;
  Ruminations: string;
  Notes: string;
}

function toNum(s: string) {
  switch (s) {
    case "verybad":
      return -2;
    case "bad":
      return -1;
    case "neutral":
      return 0;
    case "good":
      return 1;
    case "verygood":
      return 2;
  }
  return 0;
}

async function get(from: Date) {
  const formData = new FormData();
  formData.set("from", from.toISOString());
  formData.set("to", new Date().toISOString());
  const f = await fetch("./api/get", {
    method: "POST",
    body: formData,
  });
  if (f.status !== 200) {
    alert("error while updating the mood");
    return;
  }
  const json = await f.json();
  const labels: string[] = [];
  const data: { x: Date; y: number }[] = [];
  json.forEach((element: Mood) => {
    labels.push(element.Date);
    data.push({ x: new Date(element.Date), y: toNum(element.Global) });
  });
  return { labels: labels, data: data };
}

async function generate(from: Date) {
  const res = await get(from);
  const ctx = document.getElementById("general");
  new Chart(ctx as ChartItem, {
    type: "line",
    data: {
      labels: res?.labels,
      datasets: [
        {
          data: res?.data,
          label: "Global",
        },
      ],
    },
    options: {
      scales: {
        y: {
          suggestedMax: 4,
          suggestedMin: -4,
        },
        x: {
          type: "time",
        },
      },
    },
  });
}

onMounted(async () => {
  const from = new Date();
  from.setDate(0);
  generate(from);
});
</script>

<template>
  <h1 class="title">{{ $t("graphes.title") }}</h1>
  <main>
    <canvas id="general"></canvas>
  </main>
</template>
