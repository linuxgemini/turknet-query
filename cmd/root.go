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
	"github.com/golang-demos/chalk"
	"github.com/linuxgemini/turknet-query/api"
	"github.com/manifoldco/promptui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "turknet-query",
	Short: "Query your estimated line speed if you choose Turk.net as your ISP.",
	Long:  `A CLI applet to query your infra speed without loading the entire turk.net website.`,
	Run: func(cmd *cobra.Command, args []string) {
		tnAPI := api.CreateTurknetAPIClient()

		var tnRes api.TurknetAPI_ServiceAvailability
		var valueType string
		var queryType string
		var answer string
		var err error

		flagBbk, _ := cmd.Flags().GetString("bbk")
		flagPstn, _ := cmd.Flags().GetString("pstn")

		// TODO: query by address
		if flagBbk != "" {
			queryType = "BBK"
			answer = flagBbk
		} else if flagPstn != "" {
			queryType = "Telefon Numarası"
			answer = flagPstn
		} else {
			qtPrompt := promptui.Select{
				Label: "Sorgu Tipi",
				Items: []string{"BBK", "Telefon Numarası"},
			}

			_, queryType, err = qtPrompt.Run()

			if err != nil {
				panic(err)
			}
		}

		if queryType == "BBK" {
			valueType = "BBK"
			if answer == "" {
				answer = askQuestion("BBK Numarası")
			}
		} else if queryType == "Telefon Numarası" {
			valueType = "PSTN"
			if answer == "" {
				answer = askQuestion("Telefon Numarası")
			}
		}

		tnRes, err = tnAPI.CheckServiceAvailability(valueType, answer)

		if err != nil {
			panic(err)
		}

		fmt.Printf(`%s
	Var mı?: %s
	GigaFiber mi?: %s
	Maksimum Kapasite: %d
%s
	Var mı?: %s
	Maksimum kapasite: %d
	Maksimum kapasite servis tipi: %d
	NmsMax: %d
	Tip: %d
	Açıklama: %s
%s
	Var mı?: %s
	Maksimum kapasite: %d
	Maksimum kapasite servis tipi: %d
	NmsMax: %d
	Açıklama: %s
%s
	Var mı?: %s
	Maksimum kapasite: %d
	NmsMax: %d
	Açıklama: %s
%s
	Var mı?: %s
	"Indoor" mu?: %s
	Türknet santralde aktif mi?: %s
	Açıklama: %s
`,
			chalk.Cyan("Türknet Fiber Durumu:"),
			getBooleanText(tnRes.FiberServiceAvailablity.IsAvailable),
			getBooleanText(tnRes.FiberServiceAvailablity.IsGigaFiber),
			tnRes.FiberServiceAvailablity.MaxCapacity,

			chalk.Cyan("VAE Fiber Durumu:"),
			getBooleanText(tnRes.VAEFiberServiceAvailability.IsAvailable),
			tnRes.VAEFiberServiceAvailability.MaxCapacity,
			tnRes.VAEFiberServiceAvailability.MaxCapacityServiceType,
			tnRes.VAEFiberServiceAvailability.NmsMax,
			tnRes.VAEFiberServiceAvailability.Type,
			getDescription(tnRes.VAEFiberServiceAvailability.Description),

			chalk.Cyan("VDSL Durumu:"),
			getBooleanText(tnRes.VDSLServiceAvailability.IsAvailable),
			tnRes.VDSLServiceAvailability.MaxCapacity,
			tnRes.VDSLServiceAvailability.MaxCapacityServiceType,
			tnRes.VDSLServiceAvailability.NmsMax,
			getDescription(tnRes.VDSLServiceAvailability.Description),

			chalk.Cyan("xDSL Durumu:"),
			getBooleanText(tnRes.XDSLServiceAvailability.IsAvailable),
			tnRes.XDSLServiceAvailability.MaxCapacity,
			tnRes.XDSLServiceAvailability.NmsMax,
			getDescription(tnRes.XDSLServiceAvailability.Description),

			chalk.Cyan("YAPA Durumu:"),
			getBooleanText(tnRes.YapaServiceAvailability.IsAvailable),
			getBooleanText(tnRes.YapaServiceAvailability.IsIndoor),
			getBooleanText(tnRes.YapaServiceAvailability.IsTurknetStatusActiveForSantral),
			getDescription(tnRes.YapaServiceAvailability.Description),
		)
	},
}

func getDescription(description string) interface{} {
	if description != "" {
		return description
	}
	return chalk.Red("Yok")
}

func getBooleanText(truthy bool) *chalk.Color {
	if truthy {
		return chalk.Green("Evet")
	} else {
		return chalk.Red("Hayır")
	}
}

func askQuestion(question string) string {
	prompt := promptui.Prompt{
		Label: question,
	}

	res, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	return res
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
