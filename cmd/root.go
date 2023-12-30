/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pr0ph0z/fgo-certificate-extractor/extract"
	"github.com/spf13/cobra"
	"os"
)

var (
	file               string
	encodedCertificate string
	rootCmd            = &cobra.Command{
		Use:   "fgo-certificate-extractor",
		Short: "Extract the certificate from the FGO data file",
		Long:  `Extract the certificate from the FGO data file`,
		Run: func(cmd *cobra.Command, args []string) {
			if file == "" && len(args) == 0 {
				fmt.Println("Please specify the certificate file or the certificate string")
				os.Exit(1)
			}
			if file != "" {
				certificateByte, err := os.ReadFile(file)
				if err != nil {
					panic(err)
				}

				encodedCertificate = string(certificateByte[2:])
			}
			if len(args) != 0 {
				encodedCertificate = args[0]
			}

			certificate, err := extract.Extract(encodedCertificate)
			if err != nil {
				panic(err)
			}

			var obj map[string]interface{}
			json.Unmarshal(certificate, &obj)

			b, err := json.MarshalIndent(obj, "", "    ")
			fmt.Println(string(b))
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "Certificate file")
}
