/*
Copyright © 2022 linuxgemini

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
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type turknetAPIClient struct {
	ClientToken     string
	ClientUserAgent string
}

type TurknetAPI_ServiceResult struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
	Message     string `json:"Message"`
	ResultType  int    `json:"ResultType"`
}

type TurknetAPI_ServiceAvailability struct {
	FiberServiceAvailablity struct {
		IsAvailable bool `json:"IsAvailable,omitempty"`
		IsGigaFiber bool `json:"IsGigaFiber,omitempty"`
		MaxCapacity int  `json:"MaxCapacity,omitempty"`
	} `json:"FiberServiceAvailablity,omitempty"`
	IsGigaFiberPlanned          bool `json:"IsGigaFiberPlanned,omitempty"`
	VAEFiberServiceAvailability struct {
		Description            interface{} `json:"Description,omitempty"`
		IsAvailable            bool        `json:"IsAvailable,omitempty"`
		MaxCapacity            int         `json:"MaxCapacity,omitempty"`
		MaxCapacityServiceType int         `json:"MaxCapacityServiceType,omitempty"`
		NmsMax                 int         `json:"NmsMax,omitempty"`
		Type                   int         `json:"Type,omitempty"`
	} `json:"VAEFiberServiceAvailability,omitempty"`
	VDSLServiceAvailability struct {
		Description            string `json:"Description,omitempty"`
		IsAvailable            bool   `json:"IsAvailable,omitempty"`
		MaxCapacity            int    `json:"MaxCapacity,omitempty"`
		MaxCapacityServiceType int    `json:"MaxCapacityServiceType,omitempty"`
		NmsMax                 int    `json:"NmsMax,omitempty"`
		TnRealSpeed            int    `json:"TnRealSpeed,omitempty"`
	} `json:"VDSLServiceAvailability,omitempty"`
	XDSLServiceAvailability struct {
		Description       interface{} `json:"Description,omitempty"`
		FiberStatusID     int         `json:"FiberStatusId,omitempty"`
		IsAvailable       bool        `json:"IsAvailable,omitempty"`
		IsExistedCustomer bool        `json:"IsExistedCustomer,omitempty"`
		IsIndividualFiber bool        `json:"IsIndividualFiber,omitempty"`
		MaxCapacity       int         `json:"MaxCapacity,omitempty"`
		NmsMax            int         `json:"NmsMax,omitempty"`
		TnRealSpeed       int         `json:"TnRealSpeed,omitempty"`
	} `json:"XDSLServiceAvailability,omitempty"`
	YapaServiceAvailability struct {
		Description                     string `json:"Description,omitempty"`
		IsAvailable                     bool   `json:"IsAvailable,omitempty"`
		IsIndoor                        bool   `json:"IsIndoor,omitempty"`
		IsTurknetStatusActiveForSantral bool   `json:"IsTurknetStatusActiveForSantral,omitempty"`
	} `json:"YapaServiceAvailability,omitempty"`
}

type TurknetAPI_IDname struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type TurknetAPI_GetToken_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	Token         string                   `json:"Token,omitempty"`
}

type TurknetAPI_GetBBKCountyList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstIlce       []TurknetAPI_IDname      `json:"lstIlce,omitempty"`
}

type TurknetAPI_GetBBKBucakList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstBucak      []TurknetAPI_IDname      `json:"lstBucak,omitempty"`
}

type TurknetAPI_GetBBKKoyList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstKoy        []TurknetAPI_IDname      `json:"lstKoy,omitempty"`
}

type TurknetAPI_GetBBKMahalleList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstMahalle    []TurknetAPI_IDname      `json:"lstMahalle,omitempty"`
}

type TurknetAPI_GetBBKCaddeList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstCadde      []TurknetAPI_IDname      `json:"lstCadde,omitempty"`
}

type TurknetAPI_GetBBKBinaList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstBina       []TurknetAPI_IDname      `json:"lstBina,omitempty"`
}

type TurknetAPI_GetBBKList_Response struct {
	ServiceResult TurknetAPI_ServiceResult `json:"ServiceResult"`
	LstDaire      []TurknetAPI_IDname      `json:"lstDaire,omitempty"`
}

type TurknetAPI_CheckServiceAvailability_Response struct {
	ServiceResult TurknetAPI_ServiceResult       `json:"ServiceResult"`
	Result        TurknetAPI_ServiceAvailability `json:"Result,omitempty"`
}

func CreateTurknetAPIClient() turknetAPIClient {
	return turknetAPIClient{"", "github.com/linuxgemini/turknet-query/cmd TurknetAPIClient/1.1"}
}

func (tac *turknetAPIClient) GetToken() (err error) {
	url := "https://turk.net/service/AddressServ.svc/GetToken"

	var jsonStr = []byte(`{}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", tac.ClientUserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result TurknetAPI_GetToken_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return errors.New("failed to unmarshal TurknetAPI_GetToken_Response")
	}

	tac.ClientToken = result.Token

	return
}

func (tac *turknetAPIClient) _CreateWebRequest(url string, jsonStr []byte) ([]byte, error) {
	if tac.ClientToken == "" {
		if err := tac.GetToken(); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", tac.ClientUserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Token", tac.ClientToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (tac *turknetAPIClient) _CreateGetWebRequest(url string) ([]byte, error) {
	if tac.ClientToken == "" {
		if err := tac.GetToken(); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("GET", url, strings.NewReader(``))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", tac.ClientUserAgent)
	req.Header.Set("Token", tac.ClientToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (tac *turknetAPIClient) GetBBKCountyList(ilPlakaKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKCountyList/%d", ilPlakaKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKCountyList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKCountyList_Response")
	}

	return result.LstIlce, nil
}

func (tac *turknetAPIClient) GetBBKBucakList(ilceKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKBucakList/%d", ilceKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKBucakList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKBucakList_Response")
	}

	return result.LstBucak, nil
}

func (tac *turknetAPIClient) GetBBKKoyList(bucakKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKKoyList/%d", bucakKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKKoyList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKKoyList_Response")
	}

	return result.LstKoy, nil
}

func (tac *turknetAPIClient) GetBBKMahalleList(koyKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKMahalleList/%d", koyKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKMahalleList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKMahalleList_Response")
	}

	return result.LstMahalle, nil
}

func (tac *turknetAPIClient) GetBBKCaddeList(mahalleKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKCaddeList/%d", mahalleKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKCaddeList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKCaddeList_Response")
	}

	return result.LstCadde, nil
}

func (tac *turknetAPIClient) GetBBKBinaList(caddeKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKBinaList/%d", caddeKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKBinaList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKBinaList_Response")
	}

	return result.LstBina, nil
}

func (tac *turknetAPIClient) GetBBKList(binaKodu int) ([]TurknetAPI_IDname, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKList/%d", binaKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPI_GetBBKList_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPI_GetBBKList_Response")
	}

	return result.LstDaire, nil
}

func (tac *turknetAPIClient) CheckServiceAvailability(valueType string, value string) (TurknetAPI_ServiceAvailability, error) {
	url := "https://turk.net/service/AddressServ.svc/CheckServiceAvailability"

	var jsonStr = []byte(fmt.Sprintf(`{"InquirySource":2,"IsInfrastructureInquiry":true,"Key":"%s","Value":"%s"}`, valueType, value))

	body, err := tac._CreateWebRequest(url, jsonStr)
	if err != nil {
		return TurknetAPI_ServiceAvailability{}, err
	}

	var result TurknetAPI_CheckServiceAvailability_Response
	if err := json.Unmarshal(body, &result); err != nil {
		return TurknetAPI_ServiceAvailability{}, errors.New("failed to unmarshal TurknetAPI_CheckServiceAvailability_Response")
	}

	return result.Result, nil
}
