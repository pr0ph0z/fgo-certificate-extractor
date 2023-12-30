/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pr0ph0z/fgo-certificate-extractor/extract"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	file               string
	encodedCertificate string
	rootCmd            = &cobra.Command{
		Use:   "fgo-certificate-extractor",
		Short: "Extract the certificate from the FGO data file",
		Long:  `Extract the certificate from the FGO data file`,
		Run: func(cmd *cobra.Command, args []string) {
			if isInputFromPipe() {
				var outputBuffer strings.Builder
				scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
				for scanner.Scan() {
					_, e := fmt.Fprintln(
						&outputBuffer, scanner.Text())
					if e != nil {
						panic(e)
					}
				}

				encodedCertificate = outputBuffer.String()
			} else if file != "" {
				certificateByte, err := os.ReadFile(file)
				if err != nil {
					panic(err)
				}

				encodedCertificate = string(certificateByte)
			} else if len(args) != 0 {
				encodedCertificate = args[0]
			} else {
				fmt.Println("Please specify the certificate file or the certificate string")
				os.Exit(1)
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

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
