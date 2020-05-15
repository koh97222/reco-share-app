<template>
  <div id="app">
    <TopBar class="sticky-top"></TopBar>
    <router-link to="/"></router-link>
    <router-view />
  </div>
</template>

<script>
import TopBar from "./components/TopBar.vue";

export default {
  name: "app",
  components: { TopBar },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: black;
}
.container {
  background-color: #f4f5f7;
}
</style>
