<template>
  <v-layout>
    <v-flex>
      <v-tabs fixed-tabs v-model="model">
        <v-tab :key="'installed'">Installed</v-tab>
        <v-tab :key="'available'">Available</v-tab>
      </v-tabs>
      <v-tabs-items class="mt-3" v-model="model">
        <v-tab-item :key="'installed'">
          <v-spacer></v-spacer>
          <v-flex v-if="isLoading" offset-xs1 xs10>
            <div class="text-xs-center">
              <v-progress-circular :size="50" color="primary" indeterminate></v-progress-circular>
            </div>
          </v-flex>
          <v-card v-else>
            <v-list two-line>
              <template v-for="(item, index) in data.installed">
                <v-list-tile :key="item.name">
                  <v-list-tile-content>
                    <v-list-tile-title
                      class="title font-weight-thin my-1"
                      color="blue-grey darken-1"
                    >{{ item.name }}</v-list-tile-title>
                    <v-list-tile-sub-title>
                      <v-layout row xs12>
                        <v-flex xs3 align-self-center>
                          <span class="subheading">
                            current:
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
                        <v-flex align-self-center>
                          <v-btn flat icon @click.stop="dialog = true; uninstall = item.name">
                            <v-icon>far fa-trash-alt</v-icon>
                          </v-btn>
                        </v-flex>
                      </v-layout>
                    </v-list-tile-sub-title>
                  </v-list-tile-content>
                </v-list-tile>
                <v-divider class="my-3" :key="index"></v-divider>
              </template>
            </v-list>
          </v-card>
          <v-dialog v-model="dialog" max-width="520">
            <v-card>
              <v-card-title class="headline font-weight-thin" primary-title>Confirm Uninstall</v-card-title>
              <v-card-text class="title font-weight-light pa-5">Are you sure to uninstall
                <span class="headline font-weight-medium font-italic">{{ uninstall }}</span> ?
              </v-card-text>
              <v-card-actions class="pa-4">
                <v-spacer></v-spacer>
                <v-btn outline flat="flat" @click="dialog = false">Cancel</v-btn>
                <v-btn color="error darker-1" @click="dialog = false">Uninstall</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-tab-item>
        <v-tab-item :key="'available'">
          <v-flex offset-xs1 xs10>
            <v-card-text>
              <v-autocomplete
                v-model="searchKeyword"
                :hint="!isEditing ? 'Filter by keyword' : 'Click the icon to save'"
                :items="states"
                :label="`Search...`"
                :multiple="true"
                chips
                persistent-hint
                clearable
              ></v-autocomplete>
            </v-card-text>
          </v-flex>
          <v-card>
            <v-list two-line>
              <v-list-tile v-for="formula in formulaPerPage(page)" :key="formula.name" class="my-3">
                <v-list-tile-content>
                  <v-list-tile-title
                    class="title font-weight-thin my-1"
                    color="blue-grey darken-1"
                  >{{ formula.name }}</v-list-tile-title>
                  <v-list-tile-sub-title>
                    <v-layout row xs12>
                      <v-flex xs3 align-self-center>
                        <span class="subheading">
                          stable version:
                          <span class="font-weight-bold">{{ formula.stable }}</span>
                        </span>
                      </v-flex>
                      <v-flex align-self-center>
                        <v-btn round outline color="indigo">Install</v-btn>
                      </v-flex>
                    </v-layout>
                  </v-list-tile-sub-title>
                </v-list-tile-content>
              </v-list-tile>
            </v-list>
          </v-card>
          <div class="text-xs-center pa-3">
            <v-pagination v-model="page" :length="maxPage()" :total-visible="10"></v-pagination>
          </div>
        </v-tab-item>
      </v-tabs-items>
    </v-flex>
  </v-layout>
</template>

<script lang='ts'>
import Vue, { PropOptions } from 'vue'
import { formulaEndpoint } from '../constants'

const countPerPage = 20

interface Formula {
  name: string,
  stable: string,
}
type Formulas = Formula[]
interface Response {
  name: string,
  versions: Versions,
}
interface Versions {
  stable: string,
}

export default Vue.extend({
  name: 'Dependency',
  props: {
    data: Object,
    // formulas: Array as PropOptions<Formula[]>,
  },
  data() {
    return {
      isEditing: false,
      dialog: false,
      uninstall: '',
      searchKeyword: '',
      model: null,
      states: [] as any[],
      page: 1,
      formulas: [] as Formulas,
    }
  },
  methods: {
    maxPage(): number {
      return Math.floor(this.filteredFormula.length / countPerPage)
    },
    formulaPerPage(page: number): Formula[] {
      const startIndex: number = countPerPage * (page - 1)

      return this.filteredFormula.slice(startIndex, startIndex + countPerPage)
    },
  },
  computed: {
    filteredFormula(): Formulas {
      if (this.searchKeyword && this.searchKeyword.length > 0) {
        return this.formulas.filter((formula) =>
          this.searchKeyword.includes(formula.name),
        )
      }
      return this.formulas
    },
    isLoading(): boolean {
      return this.data.installed === undefined
    },
  },
  created() {
    Vue.axios.get(`${formulaEndpoint}`).then((response) => {
      response.data.forEach((element: Response) => {
        this.states.push(element.name)
        const formula: Formula = {
          name: element.name,
          stable: element.versions.stable,
        }
        this.formulas.push(formula)
      })
    })
  },
})
</script>
