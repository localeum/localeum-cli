package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Download your current translations in various file format.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		CheckConfigFile()

		Api.SetAuthToken(viper.GetString(ApiKeyFlag))

		resp, err := Api.Get("/v1/languages", map[string]string{})
		if err != nil {
			return err
		}

		if resp.StatusCode() == 401 {
			fmt.Println("API key is incorrect!")
			os.Exit(1)
		}

		if resp.StatusCode() != 200 {
			fmt.Println("Server error!", resp)
			os.Exit(1)
		}

		var langs []string

		if err = json.Unmarshal(resp.Body(), &langs); err != nil {
			return err
		}

		MakeDirectory()

		var wg sync.WaitGroup
		wg.Add(len(langs))
		for _, lang := range langs {
			go fetchFile(lang, &wg)
		}

		wg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func fetchFile(lang string, wg *sync.WaitGroup) {
	format := viper.GetString(FormatFlag)
	resp, err := Api.Get("/v1/export", map[string]string{
		"format": format,
		"lang":   lang,
	})
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	if resp.StatusCode() != 200 {
		fmt.Println("Server error!", resp)
		wg.Done()
		return
	}

	filename := viper.GetString(DirectoryFlag) + "/" + lang + "." + format
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	_, err = file.Write(resp.Body())
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	fmt.Println("File -", filename, "downloaded")
	wg.Done()
}
