<template>
  <div id="app">
    <v-app>
      <Navigation :version="version"/>
      <v-content>
        <v-container grid-list-xl fluid fill-height>
          <router-view :data="data" :isLatest="isLatest"></router-view>
        </v-container>
      </v-content>
    </v-app>
  </div>
</template>

<script lang='ts'>
import Vue from 'vue'
import Navigation from '@/components/Navigation.vue'
import { apiEndpoint, githubApiEndpoint, homebrewEndpoint } from './constants'

interface ApiData {
  latest?: string,
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
    const data: ApiData = {}
    Vue.axios.get(`${githubApiEndpoint}${homebrewEndpoint}`).then((response: any) => {
      data.latest = response.data.name
      this.data = Object.assign({}, data)
    })
    Vue.axios.get(`${apiEndpoint}?query={version}`).then((response: any) => {
      data.version = response.data.data.version
      this.data = Object.assign({}, data)
    })
    Vue.axios.get(`${apiEndpoint}?query={installed{name,version{current,latest}}}`).then((response: any) => {
      data.installed = response.data.data.installed
      this.data = Object.assign({}, data)
    })
  },
  computed: {
    isLatest(): boolean {
      if (this.data.version === undefined) {
        return false
      }
      const current: string[] = this.data.version.split(' ')
      return this.data.latest === current[1]
    },
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
