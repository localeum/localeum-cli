package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"localeum-cli/core"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure your CLI tool",
	Long:  GetHeader(),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err:= checkExistConfigFile(); err != nil {
			return nil
		}

		fmt.Print("Enter the API access token for your project: ")
		if _, err := fmt.Fscan(os.Stdin, &ApiKey); err != nil {
			return err
		}

		fmt.Print("Enter the path to the DirectoryFlag for the localization files: ")
		if _, err := fmt.Fscan(os.Stdin, &Directory); err != nil {
			return err
		}

		fmt.Print("Enter the file FormatFlag for export (json, csv, yaml): ")
		if _, err := fmt.Fscan(os.Stdin, &Format); err != nil {
			return err
		}

		viper.Set(ApiKeyFlag, ApiKey)
		viper.Set(DirectoryFlag, Directory)
		viper.Set(FormatFlag, Format)

		err := viper.WriteConfig()
		if err != nil {
			return err
		}

		fmt.Println("Success! You finished initializing your application!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func checkExistConfigFile() error {
	file, _ := os.Open(ConfigFileName)

	if file != nil {
		isConfirm, err := core.Confirm(`Config file (` + ConfigFileName + `) already exists. Do you want to override this file?`)
		if err != nil {
			return err
		}

		if !isConfirm {
			os.Exit(0)
		}
	}

	return nil
}
