/*
Copyright (c) 2022 İlteriş Yağıztegin Eroğlu (linuxgemini)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxgemini/turknet-query/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "turknet-query",
	Short: "Query your estimated line speed if you choose Turk.net as your ISP.",
	Long:  `A CLI applet to query your infra speed without loading the entire turk.net website.`,
	Run: func(cmd *cobra.Command, args []string) {
		tnAPI := api.CreateTurknetAPIClient()

		flag_bbk, _ := cmd.Flags().GetString("bbk")
		flag_pstn, _ := cmd.Flags().GetString("pstn")

		var result api.TurknetAPI_ServiceAvailability
		var err error

		if flag_bbk != "" {
			result, err = tnAPI.CheckServiceAvailability("BBK", flag_bbk)
		} else if flag_pstn != "" {
			result, err = tnAPI.CheckServiceAvailability("PSTN", flag_pstn)
		}

		// todo: complete this function
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("bbk", "", "Bağımsız Bölüm Kodu")
	rootCmd.Flags().String("pstn", "", "Başında 0 olmadan sabit telefon numarası")
}
