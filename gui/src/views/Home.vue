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
                <div class="ma-3 display-1 font-weight-mediumx" :class="{ 'pink--text darken-1': errorCount > 0 }">{{ errorCount }}</div>
              </v-flex>
              <v-flex xs6>
                <div class="ma-3 headline font-weight-thin">warning</div>
                <div class="ma-3 display-1 font-weight-medium" :class="{'deep-orange--text': warningCount > 0}">{{ warningCount }}</div>
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
    <template>
      <v-expansion-panel v-if="errorCount !== 0">
        <v-expansion-panel-content v-for="(error, index) in data.doctor.errors" :key="index">
          <template v-slot:header>
            <div>
              <span class="subtitle-1 font-weight-bold pink--text darken-1 mr-2">Error: </span>
              <span class="subtitle-1">{{ format(error.subject) }}</span>
          </div>
          </template>
          <v-card>
            <v-card-text v-for="(msg, i) in error.detail" :key="i">{{ msg }}</v-card-text>
          </v-card>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel v-if="warningCount !== 0">
        <v-expansion-panel-content v-for="(warning, index) in data.doctor.warnings" :key="index">
          <template v-slot:actions>
            <v-icon color="primary" :size="15">$vuetify.icons.expand</v-icon>
          </template>
          <template v-slot:header>
            <div>
              <span class="subtitle-1 font-weight-bold deep-orange--text mr-2">Warning: </span>
              <span class="subtitle-1">{{ format(warning.subject) }}</span>
          </div>
          </template>
          <v-card>
            <v-card-text class="grey lighten-3 py-2" v-for="(msg, i) in warning.detail" :key="i">{{ msg }}</v-card-text>
          </v-card>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </template>
  </v-layout>
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
  methods: {
    format: function(subject: string): string {
      return subject.replace(/Warning:|Error:/, "")
    },
  },
  computed: {
    errorCount(): number {
      if (this.data.doctor === undefined) {
        return 0
      }
      return this.data.doctor.errors.length
    },
    warningCount(): number {
      if (this.data.doctor === undefined) {
        return 0
      }
      console.log(this.data.doctor)
      return this.data.doctor.warnings.length
    },
  },
})
</script>
