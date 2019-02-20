package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yutod/warehouse/api"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "warehouse",
	Short: "Warehouse is a great homebrew management tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize application",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := exec.Command("pwd").Output()
		command := exec.Command("yarn", "install")
		dir := string(out)
		dir = strings.TrimSuffix(dir, "\n") + "/gui"
		command.Dir = dir
		out1, err := command.Output()
		if err != nil {
			fmt.Println("Please make sure 'yarn' command is available")
		} else {
			fmt.Println(string(out1))
		}
		os.Exit(0)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start Homebrew Client - Warehouse!!! -",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := exec.Command("pwd").Output()
		baseDir := strings.TrimSuffix(string(pwd), "\n")

		// build
		command := exec.Command("./node_modules/.bin/vue-cli-service", "build", "./src/main.ts")
		command.Dir = baseDir + "/gui"
		exec.Command("rm", "-rf", baseDir+"/gui/node_modules/.cache").Run()
		out, err := command.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Please make sure './gui/node_modules/.bin/vue-cli-service' command is available")
			os.Exit(1)
		} else {
			fmt.Println(string(out))
		}

		// serve
		// command1 := exec.Command("./node_modules/.bin/vue-cli-service", "serve")
		// command1.Dir = baseDir + "/gui"
		// command1.Run()
		// out1, err := command1.CombinedOutput()
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	fmt.Println("Please make sure './gui/node_modules/.bin/vue-cli-service' command is available")
		// 	os.Exit(1)
		// } else {
		// 	fmt.Println(string(out1))
		// }

		api.Start()
	},
}

func init() {
	rootCmd.AddCommand(initCmd, runCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// cobra.OnInitialize(initConfig)
// rootCmd.PersistentFlags().StringVar(&cfgFile, "init", "", "")
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".cobra")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
