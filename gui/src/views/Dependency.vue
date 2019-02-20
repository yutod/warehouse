<template>
  <v-layout>
    <!-- <v-toolbar app></v-toolbar> -->
    <v-flex>
      <v-tabs fixed-tabs v-model="model">
        <v-tab :key="'installed'">Installed</v-tab>
        <v-tab :key="'available'">Available</v-tab>
      </v-tabs>
      <v-tabs-items v-model="model">
        <v-tab-item :key="'installed'">
          <v-card>
            <v-list two-line>
              <v-list-tile v-for="item in data.installed" :key="item.name" class="my-3">
                <v-list-tile-content>
                  <v-list-tile-title
                    class="title font-weight-thin my-1"
                    color="blue-grey darken-1"
                  >{{ item.name }}</v-list-tile-title>
                  <v-list-tile-sub-title>
                    <v-layout row xs12>
                      <v-flex xs3 align-self-center>
                        <span class="subheading">current:
                          <span class="font-weight-bold">{{ item.version.current }}</span>
                        </span>
                      </v-flex>
                      <v-flex xs3 align-self-center>
                        <span class="subheading">latest:
                          <span class="font-weight-bold">{{ item.version.latest }}</span>
                        </span>
                      </v-flex>
                      <v-flex xs3 align-self-center>
                        <v-icon :size="15" color="green" class="text-md-center">fas fa-check</v-icon>
                        <span class="subheading mx-2">Installed</span>
                      </v-flex>
                      <v-flex align-self-center>
                        <v-btn round outline color="indigo">Upgrade</v-btn>
                      </v-flex>
                    </v-layout>
                  </v-list-tile-sub-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list>
            <!-- <v-card-text></v-card-text> -->
          </v-card>
        </v-tab-item>
        <v-tab-item :key="'available'">
          <v-card>
            <v-card-text>bb</v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
      <!-- <v-parallax dark src="/brew.jpg"> -->
      <!-- <v-list two-line>
        <v-list-tile v-for="item in data.installed" :key="item.name">
          <v-list-tile-content>
            <v-list-tile-title>{{ item.name }}</v-list-tile-title>
            <v-list-tile-sub-title>version: {{ item.version }}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>-->
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

<script lang='ts'>
import Vue from 'vue'

export default Vue.extend({
  name: 'Dependency',
  props: { data: Object },
  data() {
    return {
      model: null,
    }
  },
})
</script>
