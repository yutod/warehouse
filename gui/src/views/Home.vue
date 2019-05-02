<template>
  <!-- <div class="home"> -->
  <v-layout flex-child wrap>
    <v-flex xs12 md4 d-flex>
      <!-- <v-sheet class="d-flex" color="grey lighten-3" height="400"> -->
      <v-sheet class="d-flex" color="rgb(255,255,255,0.9)" height="400">
        <v-layout xs12 justify-center column wrap>
          <v-flex xs3>
            <div class="ma-4 display-1 font-weight-bold">Version</div>
          </v-flex>
          <v-flex>
            <v-icon :size="100">fas fa-code-branch</v-icon>
          </v-flex>
          <v-flex xs1>
            <div class="ma-1 headline font-weight-thin">{{ data.version }}</div>
            <div v-show="isLatest" class="teal accent-2 font-italic">ready to brew!!</div>
            <div v-show="!isLatest" class="yellow accent-2 font-italic">recommend to update brew ...</div>
          </v-flex>
          <v-flex>
            <v-btn outline large color="indigo" :disabled="isLatest">Update ({{ data.latest }})</v-btn>
            <!-- <div class="ma-3 headline font-weight-thin">{{ data.version }}</div> -->
          </v-flex>
        </v-layout>
      </v-sheet>
    </v-flex>
    <v-flex xs12 md4 d-flex>
      <v-sheet class="d-flex" color="rgb(255,255,255,0.9)" height="400">
        <v-layout xs12 justify-center column wrap>
          <v-flex xs2>
            <div class="ma-4 display-1 font-weight-bold">Doctor</div>
          </v-flex>
          <v-flex xs3>
            <v-icon :size="100">fas fa-briefcase-medical</v-icon>
          </v-flex>
          <v-flex d-flex>
            <v-layout xs12 justify-center row wrap>
              <v-flex xs6>
                <div class="ma-3 headline font-weight-thin">error</div>
                <div class="ma-3 display-1 font-weight-medium">0</div>
              </v-flex>
              <v-flex xs6>
                <div class="ma-3 headline font-weight-thin">warning</div>
                <div class="ma-3 display-1 font-weight-medium">0</div>
                <!-- <v-badge color="red" right>
                  <span slot="badge">6</span>
                  <v-icon :size="50">fas fa-exclamation</v-icon>
                </v-badge>-->
              </v-flex>
            </v-layout>
          </v-flex>
        </v-layout>
      </v-sheet>
    </v-flex>
    <v-flex xs12 md4 d-flex>
      <v-sheet class="d-flex" color="rgb(255,255,255,0.9" height="400">
        <v-layout xs12 justify-center column wrap>
          <v-flex xs3>
            <div class="ma-4 display-1 font-weight-bold">Clean-up</div>
          </v-flex>
          <v-flex>
            <v-icon :size="100">fas fa-trash-alt</v-icon>
          </v-flex>
          <v-spacer></v-spacer>
          <v-spacer></v-spacer>
          <v-spacer></v-spacer>
          <v-flex>
            <v-btn outline large color="pink">Clean Up</v-btn>
          </v-flex>
        </v-layout>
      </v-sheet>
    </v-flex>
  </v-layout>
  <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
  <!-- <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/> -->
  <!-- </div> -->
</template>

<script lang='ts'>
import Vue from 'vue'
import gql from 'graphql-tag'
import HelloWorld from '@/components/HelloWorld.vue' // @ is an alias to /src
import { apiEndpoint } from '../constants'

export default Vue.extend({
  name: 'home',
  props: {
    data: Object,
    isLatest: Boolean,
  },
  components: {
    HelloWorld,
  },
})
</script>
