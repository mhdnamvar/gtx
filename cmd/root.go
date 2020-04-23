package cmd

import (
	"github.com/fatih/color"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gtx",
	Short: "\nGTX: Transaction management system",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("%v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// // Find home directory.
	// 	// home, err := homedir.Dir()
	// 	// if err != nil {
	// 	// 	fmt.Println(err)
	// 	// 	os.Exit(1)
	// 	// }

	// 	viper.SetConfigName("gtx")       // name of config file (without extension)
	// 	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	// 	viper.AddConfigPath("/etc/gtx/") // path to look for the config file in
	// 	viper.AddConfigPath("$HOME")     // call multiple times to add many search paths
	// 	viper.AddConfigPath(".")         // optionally look for config in the working directory

	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		//log.Println(err)
	// 	} else {
	// 		log.Println("Using config file:", viper.ConfigFileUsed())
	// 	}
	// }

	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	log.Println("Config file changed:", e.Name)
	// })
}
