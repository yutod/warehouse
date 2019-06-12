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
                <v-hover>
                  <template slot-scope="{ hover }">
                    <v-list-tile :key="item.name">
                      <div v-if="hover && isExistNewerVersion(item)" 
                           class="d-flex transition-fast-in-fast-out darken-2 hover-label subheading white--text">
                        <v-chip label outline color="orange">
                          <v-icon left>label</v-icon><strong>Newer version is now available</strong>
                        </v-chip>
                      </div>
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
                              <span class="subheading mx-2">latest:
                                <span class="font-weight-bold">{{ item.version.latest }}</span>
                              </span>
                              <template v-if="isExistNewerVersion(item)">
                                <v-icon :size="15" color="orange darken-2" class="mb-1">fas fa-exclamation-circle</v-icon>
                              </template>
                            </v-flex>
                            <v-flex xs3 align-self-center>
                              <v-icon :size="15" color="green" class="mb-1">fas fa-check</v-icon>
                              <span class="subheading mx-2">Installed</span>
                            </v-flex>
                            <v-flex align-self-center>
                              <v-btn
                                round
                                outline
                                color="indigo"
                                :disabled="isLatestVersion(item)"
                                :loading="isUpgrading(item.name)"
                                @click="upgradeFormula(item.name, item.version)">Upgrade</v-btn>
                            </v-flex>
                            <v-flex align-self-center>
                              <v-btn flat icon @click.stop="dialog = true; deleted = false; uninstall = item.name">
                                <v-icon>far fa-trash-alt</v-icon>
                              </v-btn>
                            </v-flex>
                          </v-layout>
                        </v-list-tile-sub-title>
                      </v-list-tile-content>
                    </v-list-tile>
                  </template>
                </v-hover>
                <v-divider class="my-3" :key="index"></v-divider>
              </template>
            </v-list>
          </v-card>
          <v-dialog v-model="dialog" max-width="520">
            <v-card>
              <template v-if="!isDeleted">
                <v-card-title class="headline font-weight-regular" primary-title>Confirmation</v-card-title>
                <v-card-text class="title font-weight-light pa-5">Are you sure to uninstall
                  <span class="headline font-weight-medium font-italic">{{ uninstall }}</span> ?
                </v-card-text>
              </template>
              <template v-else>
                <v-card-title class="headline font-weight-regular" primary-title>Finished!!</v-card-title>
                <v-card-text class="title font-weight-light pa-5">
                  <v-icon :size="25" color="teal accent-2">fas fa-check</v-icon>
                  Successfully removed<span class="headline font-weight-medium font-italic"> {{ uninstall }} </span>!!
                </v-card-text>
              </template>
              <v-card-actions class="pa-4">
                <v-spacer></v-spacer>
                <v-btn outline flat="flat" @click="dialog = false">Cancel</v-btn>
                <v-btn
                  depressed
                  color="error darker-1"
                  :loading="isDeleting(uninstall)"
                  :disabled="isDeleted"
                  @click="deleteFormula(uninstall)">Uninstall</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-tab-item>
        <v-tab-item :key="'available'">
          <v-alert dismissible :value="isLatest" color="teal accent-4" icon="check_circle" outline>
            Ready to install stable version !!
          </v-alert>
          <v-alert dismissible :value="!isLatest" color="deep-orange darken-2" icon="priority_high" outline>
            Recommend to update homebrew ... ( Cannot install stable version now )
          </v-alert>
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
                        <v-btn
                          round
                          outline
                          color="indigo"
                          :disabled="isInstalled(formula.name)"
                          :loading="isInstalling(formula.name)"
                          @click="installFormula(formula.name, formula.stable)">{{ isInstalled(formula.name) ? 'Installed' : 'Install' }}</v-btn>
                        <v-chip :value="versionOfInstalledNow(formula.name) !== ''" color="teal" text-color="white">
                          <v-avatar>
                            <v-icon>check_circle</v-icon>
                          </v-avatar>
                          {{ versionOfInstalledNow(formula.name) }}
                        </v-chip>
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
import { formulaEndpoint, apiEndpoint } from '../constants'

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
interface Installed {
    name: string,
    version: {
      current: string,
      latest: string,
    },
}

export default Vue.extend({
  name: 'Dependency',
  props: {
    data: Object,
    isLatest: Boolean,
    // formulas: Array as PropOptions<Formula[]>,
  },
  data() {
    return {
      isEditing: false,
      installingLibs: [] as string[],
      upgradingLibs: [] as string[],
      deletingLibs: [] as string[],
      installedNow: {} as {[key: string]: string},
      dialog: false,
      deleted: false,
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
      const remainder = this.filteredFormula.length % countPerPage
      return Math.floor(this.filteredFormula.length / countPerPage) + (remainder !== 0 ? 1 : 0)
    },
    formulaPerPage(page: number): Formula[] {
      const startIndex: number = countPerPage * (page - 1)

      return this.filteredFormula.slice(startIndex, startIndex + countPerPage)
    },
    isExistNewerVersion(item: Installed): boolean {
      const newest: Formula | undefined = this.formulas.find((formula: Formula) => item.name === formula.name)

      return newest !== undefined && item.version.latest !== newest.stable
    },
    isLatestVersion(item: Installed): boolean {
      return item.version.current !== undefined && item.version.current === item.version.latest
    },
    isInstalled(name: string) {
      return this.installedLibs.includes(name)
    },
    isInstalling(name: string) {
      return this.installingLibs.includes(name)
    },
    isUpgrading(name: string) {
      return this.upgradingLibs.includes(name)
    },
    isDeleting(name: string) {
      return this.deletingLibs.includes(name)
    },
    versionOfInstalledNow(name: string) {
      if (this.installedNow[name] === undefined) {
        return ''
      }

      return this.installedNow[name]
    },
    installFormula(name: string, version: string): void {
      this.installingLibs.push(name)
      Vue.axios
        .get(`${apiEndpoint}?query=mutation+_{install(name:\"${name}\",version:\"${version}\"){name,version,status}}`)
        .then((response) => {
          this.installingLibs = this.installingLibs.filter((lib) => lib !== name)
          this.installedLibs.push(name)
          this.installedNow[response.data.data.install.name] = response.data.data.install.version
          const newLib: Installed = {
            name,
            version: {
              current: version,
              latest: version,
            },
          }
          this.data.installed.push(newLib)
          this.data.installed.sort((a: Installed, b: Installed) => {
            const charA = a.name.toLowerCase()
            const charB = b.name.toLowerCase()
            if (charA > charB) {
              return 1
            } else if (charA < charB) {
              return -1
            }
            return 0
          })
        })
    },
    upgradeFormula(name: string, version: string): void {
      this.upgradingLibs.push(name)
      Vue.axios
        .get(`${apiEndpoint}?query=mutation+_{upgrade(name:\"${name}\",version:\"${version}\"){name,version,status}}`)
        .then((response) => {
          this.upgradingLibs = this.upgradingLibs.filter((lib) => lib !== name)
          const upgrade = response.data.data.upgrade
          if (upgrade.status) {
            this.data.installed.forEach((lib: Installed): void => {
              if (lib.name === name) {
                lib.version.current = upgrade.version
              }
            })
          }
        })
    },
    deleteFormula(name: string): void {
      this.deletingLibs.push(name)
      Vue.axios
        .get(`${apiEndpoint}?query=mutation+_{delete(name:\"${name}\"){name,status}}`)
        .then((response) => {
          this.deletingLibs = this.deletingLibs.filter((lib) => lib !== name)
          if (response.data.data.delete.status) {
            this.deleted = true
            this.data.installed = this.data.installed.filter((lib: Installed) => lib.name !== name)
            if (this.installedNow[name] !== undefined) {
              delete this.installedNow[name]
            }
          }
        })
    },
  },
  computed: {
    filteredFormula(): Formulas {
      if (this.searchKeyword && this.searchKeyword.length > 0) {
        this.page = 1
        const filtered = this.formulas.filter((formula) => this.searchKeyword.includes(formula.name))
        const remainder = filtered.length % countPerPage
        this.page += remainder !== 0 ? Math.floor(filtered.length / countPerPage) : 0
        return filtered
      }
      return this.formulas
    },
    installedLibs(): string[] {
      if (this.data.installed === undefined) {
        return []
      }
      return this.data.installed.map((lib: Installed) => lib.name)
    },
    isLoading(): boolean {
      return this.data.installed === undefined
    },
    isDeleted(): boolean {
      return this.deleted
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

<style>
.hover-label {
  position: absolute;
  top: -10%;
  left: 25%;
}
</style>
