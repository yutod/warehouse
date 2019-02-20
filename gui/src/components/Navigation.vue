<template>
  <v-navigation-drawer app parmanent class=".navigation">
    <v-toolbar flat>
      <v-avatar :size="48">
        <img alt="homebrew icon" src="../assets/homebrew.png">
      </v-avatar>
      <v-list>
        <v-list-tile>
          <v-list-tile-title class="title">Warehouse</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-toolbar>
    <v-divider></v-divider>
    <v-list dense class="pt-0">
      <v-list-tile v-for="item in items" :key="item.title" @click="" :to="{path: item.route}">
        <v-list-tile-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-tile-action>

        <v-list-tile-content>
          <v-list-tile-title>{{ item.title }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  name: 'Navigation',
  props: {
    version: String,
  },
  data() {
    return {
      items: [
        { title: 'Homebrew', icon: 'dashboard', route: '/' },
        {
          title: 'Dependency',
          icon: 'collections_bookmark',
          route: '/dependency',
        },
        {
          title: 'Configuration',
          icon: 'settings_applications',
          route: '/configuration',
        },
      ],
      right: null,
    }
  },
})
</script>
