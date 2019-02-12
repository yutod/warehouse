<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import gql from 'graphql-tag';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

export default Vue.extend({
  name: 'home',
  // data() {
  //   return {
  //     version: ''
  //   }
  // },
  // apollo: {
  //   version: gql`query {
  //     version
  //   }`,
  // },
  components: {
    HelloWorld,
  },
  created() {
    Vue.axios.get('http://localhost:3001/graphql?query={version}').then((response) => {
      console.log(response.data);
    })
  }
});
</script>
