<script setup lang="ts">
import SiteHeader from "../components/SiteHeader.vue";
</script>

<template>
  <SiteHeader />
  <div id="home" class="container">
    <router-view></router-view>
  </div>
</template>
