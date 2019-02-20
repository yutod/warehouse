<template>
  <div id="app">
    <v-app>
      <Navigation :version="version"/>
      <!-- <Navigation :version="data.version" /> -->
      <!-- <v-toolbar app></v-toolbar> -->
      <v-content>
        <!-- <v-parallax src="/bg_home1.jpg" hight="300"> -->
          <v-container grid-list-xl fluid fill-height>
            <router-view :data="data"></router-view>
          </v-container>
        <!-- </v-parallax> -->
      </v-content>
    </v-app>
    <!-- <div id="nav"> -->
    <!-- <router-link to="/">Home</router-link> | -->
    <!-- <router-link to="/about">About</router-link> -->
    <!-- </div> -->
    <!-- <router-view /> -->
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import Navigation from '@/components/Navigation.vue'
import { apiEndpoint } from './constants'

export default Vue.extend({
  name: 'app',
  components: {
    Navigation,
  },
  data() {
    return {
      version: '',
      data: {},
      // version: '',
      // formulas: [],
    }
  },
  created() {
    Vue.axios.get(`${apiEndpoint}?query={version,installed{name,version{current,latest}}}`).then((response) => {
      // console.log(response.data.data)
      this.data = response.data.data
      // this.version = response.data.data.version
      // this.formulas = response.data.data.formulas
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
