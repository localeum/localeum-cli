package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"localeum-cli/core"
	"os"
)

var (
	cfgFile   string
	Api       *core.Api
	Directory string
	Format    string
	ApiKey    string
)

const Version = "1.0.0"
const ConfigFileName = "localeum.yml"

const ApiKeyFlag = "api-key"
const DirectoryFlag = "directory"
const FormatFlag = "format"

var rootCmd = &cobra.Command{
	Use:     "localeum-cli",
	Short:   `Localeum helps organize the localization process in your company`,
	Long:    GetHeader(),
	Version: Version,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	Api = core.NewApiClient()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+ConfigFileName+")")

	rootCmd.PersistentFlags().StringVarP(&ApiKey, ApiKeyFlag, "k", "", "Project API key. You can find it on the project settings page.")
	if err := viper.BindPFlag(ApiKeyFlag, rootCmd.PersistentFlags().Lookup(ApiKeyFlag)); err != nil {
		fmt.Println(err)
	}
	viper.SetDefault(ApiKeyFlag, "")

	rootCmd.PersistentFlags().StringVarP(&Directory, DirectoryFlag, "d", "", "Directory for download localization files.")
	if err := viper.BindPFlag(DirectoryFlag, rootCmd.PersistentFlags().Lookup(DirectoryFlag)); err != nil {
		fmt.Println(err)
	}
	viper.SetDefault(DirectoryFlag, "")

	rootCmd.PersistentFlags().StringVarP(&Format, FormatFlag, "f", "", "File format for localization files.")
	if err := viper.BindPFlag(FormatFlag, rootCmd.PersistentFlags().Lookup(FormatFlag)); err != nil {
		fmt.Println(err)
	}
	viper.SetDefault(FormatFlag, "")
}

func initConfig() {
	viper.SetConfigFile(ConfigFileName)
	viper.SetConfigType("yaml")
}

func CheckConfigFile() {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found!")
		os.Exit(1)
	}
}

func MakeDirectory() {
	path := viper.GetString(DirectoryFlag)

	if path != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func GetHeader() string {
	v := Version

	return `
+--------------------------------------------+
|                                            |
|             LOCALEUM CLI TOOL              |
|                  v` + v + `                    |
|                                            |
+--------------------------------------------+

Read more docs on https://docs.localeum.com/cli
`
}
