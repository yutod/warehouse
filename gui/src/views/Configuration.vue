<template>
  <v-layout>
    <!-- <v-toolbar app></v-toolbar> -->
    <v-flex>
      <!-- <v-parallax dark src="/brew.jpg"> -->
      <v-list two-line>
        <v-list-tile v-for="item in data.installed" :key="item.name">
          <v-list-tile-content>
            <v-list-tile-title>{{ item.name }}</v-list-tile-title>
            <v-list-tile-sub-title>version: {{ item.version }}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
      <!-- </v-parallax> -->
    </v-flex>
  </v-layout>
</template>
  <!-- <div class="dependency">
    <v-parallax src="../assets/brew.jpg">
      <h1>This is an about page</h1>
    </v-parallax>
  </div> -->
  <!-- <div class="dependency">aaaa</div> -->
  <!-- <v-parallax dark src="/brew.jpg"> -->
    <!-- <v-layout align-center column justify-center>
      <h1 class="display-2 font-weight-thin mb-3">Vuetify.js</h1>
      <h4 class="subheading">Build your application today!</h4>
    </v-layout> -->
  <!-- </v-parallax> -->

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'Configuration',
  props: { data: Object },
})
</script>
