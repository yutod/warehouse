<template>
  <div id="app">
    <v-app>
      <Navigation :version="version"/>
      <v-content>
          <v-container grid-list-xl fluid fill-height>
            <router-view :data="data"></router-view>
          </v-container>
      </v-content>
    </v-app>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import Navigation from '@/components/Navigation.vue'
import { apiEndpoint } from './constants'

interface ApiData {
 version?: string,
 installed?: {
   name?: string,
   version?: {
     current?: string,
     latest?: string,
   },
 },
}

export default Vue.extend({
  name: 'app',
  components: {
    Navigation,
  },
  data() {
    return {
      version: '',
      data: {} as ApiData,
      // version: '',
      formulas: [],
    }
  },
  created() {
    Vue.axios.get(`${apiEndpoint}?query={version,installed{name,version{current,latest}}}`).then((response: any) => {
      this.data = response.data.data
    })
  },
})
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
