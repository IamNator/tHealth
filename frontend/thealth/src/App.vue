<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <!-- <Welcome msg="Welcome"/> -->
   <Table />
  </div>
</template>

<script>
// import Welcome from './components/Welcome';
import Table from './components/table.vue';

export default {
  name: 'App',
  components: {
    // Welcome,
    Table
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
