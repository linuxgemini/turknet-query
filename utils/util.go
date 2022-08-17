/*
Package utils
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
package utils

import (
	"errors"
	"fmt"
	"regexp"
	"sort"

	"github.com/fatih/color"
	"github.com/linuxgemini/turknet-query/api"
	"github.com/manifoldco/promptui"
)

func GetColoredBoolText(v bool) string {
	if v {
		return color.GreenString("Evet")
	} else {
		return color.RedString("Hayır")
	}
}

func GetDescriptionButFancier(desc string) string {
	if desc != "" {
		return desc
	} else {
		return color.RedString("Yok")
	}
}

func GetKeysFromKVMap(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func ValidatePSTN(telno string) error {
	r, _ := regexp.Compile(`^([2-4][1-9][1-9])(\d{7})$`)

	if !r.MatchString(telno) {
		return errors.New("girilen telefon numarası geçerli değil")
	}
	return nil
}

func ValidateIfNumber(s string) error {
	r, _ := regexp.Compile(`^\d+$`)

	if !r.MatchString(s) {
		return errors.New("girilen değer geçerli değil")
	}
	return nil
}

func _AskUser(label string, items []string) (string, error) {
	prompt := &promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func AskUserForValueType() (string, error) {
	result, err := _AskUser("Lütfen sorgulama metodunu seçin", []string{"Adres", "Telefon Numarası"})
	if err != nil {
		return "", err
	}

	if result == "Adres" {
		return "BBK", nil
	} else if result == "Telefon Numarası" {
		return "PSTN", nil
	} else {
		return "", errors.New("AskUserForValueType failed somehow")
	}
}

func AskUserForPSTN() (string, error) {
	prompt := &promptui.Prompt{
		Label:    "Lütfen sabit telefon numaranızı **başında 0 olmadan** girin",
		Validate: ValidatePSTN,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

func AskUserForIl() (int, error) {
	result, err := _AskUser("Lütfen ilinizi seçin", GetKeysFromKVMap(illerAlfabetik))
	if err != nil {
		return 0, err
	}

	return illerAlfabetik[result], nil
}

func AskUserForSomething(midsentence string, list api.TurknetAPIKv) (int, error) {
	label := fmt.Sprintf("Lütfen %s seçin", midsentence)

	mapKeys := GetKeysFromKVMap(list)
	firstKey := mapKeys[0]

	var result int
	if len(mapKeys) == 1 {
		result = list[firstKey]
	} else {
		resstr, err := _AskUser(label, mapKeys)
		if err != nil {
			return 0, err
		}
		result = list[resstr]
	}

	return result, nil
}
