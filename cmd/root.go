/*
Package cmd
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
	fmt "fmt"
	"os"

	"github.com/fatih/color"
	"github.com/linuxgemini/turknet-query/api"
	"github.com/linuxgemini/turknet-query/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "turknet-query",
	Short: "Türknet Altyapı Sorgulama uygulaması.",
	Long:  "Türknet'in kendi Altyapı Sorgulama sayfasına girmeden altyapı sorgulama şeysi.",
	RunE: func(cmd *cobra.Command, args []string) error {
		tnAPI := api.CreateTurknetAPIClient()

		var (
			result                api.TurknetAPIServiceAvailabilityResult
			valueType             string
			value                 string
			promptUserForMoreData bool = true
			err                   error
			ilPlakaKodu           int
			ilceKodu              int
			bucakKodu             int
			koyKodu               int
			mahalleKodu           int
			caddeKodu             int
			binaKodu              int
			daireKodu             int
			countyList            api.TurknetAPIKv
			bucakList             api.TurknetAPIKv
			koyList               api.TurknetAPIKv
			mahalleList           api.TurknetAPIKv
			caddeList             api.TurknetAPIKv
			binaList              api.TurknetAPIKv
			daireList             api.TurknetAPIKv
		)

		flagValueBBK, _ := cmd.Flags().GetString("bbk")
		flagValuePSTN, _ := cmd.Flags().GetString("pstn")

		if flagValueBBK != "" {
			err = utils.ValidateIfNumber(flagValueBBK)
			if err != nil {
				return err
			}

			valueType = "BBK"
			value = flagValueBBK
			promptUserForMoreData = false
		} else if flagValuePSTN != "" {
			err = utils.ValidatePSTN(flagValuePSTN)
			if err != nil {
				return err
			}

			valueType = "PSTN"
			value = flagValuePSTN
			promptUserForMoreData = false
		} else {
			valueType, err = utils.AskUserForValueType()
			if err != nil {
				panic(err)
			}
		}

		if promptUserForMoreData {
			if valueType == "BBK" {
				ilPlakaKodu, err = utils.AskUserForIl()
				if err != nil {
					panic(err)
				}

				countyList, err = tnAPI.GetBBKCountyList(ilPlakaKodu)
				if err != nil {
					panic(err)
				}

				ilceKodu, err = utils.AskUserForSomething("ilçenizi", countyList)
				if err != nil {
					panic(err)
				}

				bucakList, err = tnAPI.GetBBKBucakList(ilceKodu)
				if err != nil {
					panic(err)
				}

				bucakKodu, err = utils.AskUserForSomething("bucağınızı", bucakList)
				if err != nil {
					panic(err)
				}

				koyList, err = tnAPI.GetBBKKoyList(bucakKodu)
				if err != nil {
					panic(err)
				}

				koyKodu, err = utils.AskUserForSomething("köyünüzü", koyList)
				if err != nil {
					panic(err)
				}

				mahalleList, err = tnAPI.GetBBKMahalleList(koyKodu)
				if err != nil {
					panic(err)
				}

				mahalleKodu, err = utils.AskUserForSomething("mahallenizi", mahalleList)
				if err != nil {
					panic(err)
				}

				caddeList, err = tnAPI.GetBBKCaddeList(mahalleKodu)
				if err != nil {
					panic(err)
				}

				caddeKodu, err = utils.AskUserForSomething("caddenizi/sokağınızı", caddeList)
				if err != nil {
					panic(err)
				}

				binaList, err = tnAPI.GetBBKBinaList(caddeKodu)
				if err != nil {
					panic(err)
				}

				binaKodu, err = utils.AskUserForSomething("bina numaranızı", binaList)
				if err != nil {
					panic(err)
				}

				daireList, err = tnAPI.GetBBKList(binaKodu)
				if err != nil {
					panic(err)
				}

				if len(daireList) == 0 {
					value = fmt.Sprint(binaKodu)
					promptUserForMoreData = false
				} else {
					daireKodu, err = utils.AskUserForSomething("daire numaranızı", daireList)
					if err != nil {
						panic(err)
					}
					value = fmt.Sprint(daireKodu)
					promptUserForMoreData = false
				}
			} else if valueType == "PSTN" {
				value, err = utils.AskUserForPSTN()
				if err != nil {
					panic(err)
				}
			} else {
				panic("data somehow went poof")
			}
		}

		result, err = tnAPI.CheckServiceAvailability(valueType, value)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(color.Output, `%s
	Var mı?: %s
	GigaFiber mi?: %s
	GigaFiber kurulumu planda var mı?: %s
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
	TürkNet'in tespit ettiği gerçek hız: %d
	NmsMax: %d
	Açıklama: %s
%s
	Var mı?: %s
	Halihazırda Türknet müşterisi mi?: %s
	Bireysel Fiber mi?: %s
	Fiber durum ID'si: %d
	Maksimum kapasite: %d
	TürkNet'in tespit ettiği gerçek hız: %d
	NmsMax: %d
	Açıklama: %s
%s
	Var mı?: %s
	"Indoor" mu?: %s
	Türknet santralde aktif mi?: %s
	Açıklama: %s
`,
			color.CyanString("Türknet Fiber Durumu:"),
			utils.GetColoredBoolText(result.FiberServiceAvailablity.IsAvailable),
			utils.GetColoredBoolText(result.FiberServiceAvailablity.IsGigaFiber),
			utils.GetColoredBoolText(result.IsGigaFiberPlanned),
			result.FiberServiceAvailablity.MaxCapacity,

			color.CyanString("VAE Fiber Durumu:"),
			utils.GetColoredBoolText(result.VAEFiberServiceAvailability.IsAvailable),
			result.VAEFiberServiceAvailability.MaxCapacity,
			result.VAEFiberServiceAvailability.MaxCapacityServiceType,
			result.VAEFiberServiceAvailability.NmsMax,
			result.VAEFiberServiceAvailability.Type,
			utils.GetDescriptionButFancier(result.VAEFiberServiceAvailability.Description),

			color.CyanString("VDSL Durumu:"),
			utils.GetColoredBoolText(result.VDSLServiceAvailability.IsAvailable),
			result.VDSLServiceAvailability.MaxCapacity,
			result.VDSLServiceAvailability.MaxCapacityServiceType,
			result.VDSLServiceAvailability.TnRealSpeed,
			result.VDSLServiceAvailability.NmsMax,
			utils.GetDescriptionButFancier(result.VDSLServiceAvailability.Description),

			color.CyanString("xDSL Durumu:"),
			utils.GetColoredBoolText(result.XDSLServiceAvailability.IsAvailable),
			utils.GetColoredBoolText(result.XDSLServiceAvailability.IsExistedCustomer),
			utils.GetColoredBoolText(result.XDSLServiceAvailability.IsIndividualFiber),
			result.XDSLServiceAvailability.FiberStatusID,
			result.XDSLServiceAvailability.MaxCapacity,
			result.XDSLServiceAvailability.TnRealSpeed,
			result.XDSLServiceAvailability.NmsMax,
			utils.GetDescriptionButFancier(result.XDSLServiceAvailability.Description),

			color.CyanString("YAPA Durumu:"),
			utils.GetColoredBoolText(result.YapaServiceAvailability.IsAvailable),
			utils.GetColoredBoolText(result.YapaServiceAvailability.IsIndoor),
			utils.GetColoredBoolText(result.YapaServiceAvailability.IsTurknetStatusActiveForSantral),
			utils.GetDescriptionButFancier(result.YapaServiceAvailability.Description),
		)
		if err != nil {
			panic(err)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("bbk", "", "(opsiyonel) Bağımsız Bölüm Kodu")
	rootCmd.Flags().String("pstn", "", "(opsiyonel) Başında 0 olmadan sabit telefon numarası")
}
