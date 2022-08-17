/*
Package api
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
	"io"
	"net/http"
	"strconv"
	"strings"
)

type TurknetAPIClient struct {
	ClientToken     string
	ClientUserAgent string
}

type TurknetAPIServiceResult struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
	Message     string `json:"Message"`
	ResultType  int    `json:"ResultType"`
}

type TurknetAPIFiberServiceAvailablityResult struct {
	IsAvailable bool `json:"IsAvailable,omitempty"`
	IsGigaFiber bool `json:"IsGigaFiber,omitempty"`
	MaxCapacity int  `json:"MaxCapacity,omitempty"`
}

type TurknetAPIVaeFiberServiceAvailabilityResult struct {
	Description            string `json:"Description,omitempty"`
	IsAvailable            bool   `json:"IsAvailable,omitempty"`
	MaxCapacity            int    `json:"MaxCapacity,omitempty"`
	MaxCapacityServiceType int    `json:"MaxCapacityServiceType,omitempty"`
	NmsMax                 int    `json:"NmsMax,omitempty"`
	Type                   int    `json:"Type,omitempty"`
}

type TurknetAPIVdslServiceAvailabilityResult struct {
	Description            string `json:"Description,omitempty"`
	IsAvailable            bool   `json:"IsAvailable,omitempty"`
	MaxCapacity            int    `json:"MaxCapacity,omitempty"`
	MaxCapacityServiceType int    `json:"MaxCapacityServiceType,omitempty"`
	NmsMax                 int    `json:"NmsMax,omitempty"`
	TnRealSpeed            int    `json:"TnRealSpeed,omitempty"`
}

type TurknetAPIXdslServiceAvailabilityResult struct {
	Description       string `json:"Description,omitempty"`
	FiberStatusID     int    `json:"FiberStatusId,omitempty"`
	IsAvailable       bool   `json:"IsAvailable,omitempty"`
	IsExistedCustomer bool   `json:"IsExistedCustomer,omitempty"`
	IsIndividualFiber bool   `json:"IsIndividualFiber,omitempty"`
	MaxCapacity       int    `json:"MaxCapacity,omitempty"`
	NmsMax            int    `json:"NmsMax,omitempty"`
	TnRealSpeed       int    `json:"TnRealSpeed,omitempty"`
}

type TurknetAPIYapaServiceAvailabilityResult struct {
	Description                     string `json:"Description,omitempty"`
	IsAvailable                     bool   `json:"IsAvailable,omitempty"`
	IsIndoor                        bool   `json:"IsIndoor,omitempty"`
	IsTurknetStatusActiveForSantral bool   `json:"IsTurknetStatusActiveForSantral,omitempty"`
}

type TurknetAPIServiceAvailabilityResult struct {
	FiberServiceAvailablity     TurknetAPIFiberServiceAvailablityResult     `json:"FiberServiceAvailablity,omitempty"`
	IsGigaFiberPlanned          bool                                        `json:"IsGigaFiberPlanned,omitempty"`
	VAEFiberServiceAvailability TurknetAPIVaeFiberServiceAvailabilityResult `json:"VAEFiberServiceAvailability,omitempty"`
	VDSLServiceAvailability     TurknetAPIVdslServiceAvailabilityResult     `json:"VDSLServiceAvailability,omitempty"`
	XDSLServiceAvailability     TurknetAPIXdslServiceAvailabilityResult     `json:"XDSLServiceAvailability,omitempty"`
	YapaServiceAvailability     TurknetAPIYapaServiceAvailabilityResult     `json:"YapaServiceAvailability,omitempty"`
}

type TurknetAPIIDname struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type TurknetAPIKv map[string]int

type TurknetAPIGetTokenResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	Token         string                  `json:"Token,omitempty"`
}

type TurknetAPIGetBBKCountyListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstIlce       []TurknetAPIIDname      `json:"lstIlce,omitempty"`
}

type TurknetAPIGetBBKBucakListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstBucak      []TurknetAPIIDname      `json:"lstBucak,omitempty"`
}

type TurknetAPIGetBBKKoyListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstKoy        []TurknetAPIIDname      `json:"lstKoy,omitempty"`
}

type TurknetAPIGetBBKMahalleListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstMahalle    []TurknetAPIIDname      `json:"lstMahalle,omitempty"`
}

type TurknetAPIGetBBKCaddeListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstCadde      []TurknetAPIIDname      `json:"lstCadde,omitempty"`
}

type TurknetAPIGetBBKBinaListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstBina       []TurknetAPIIDname      `json:"lstBina,omitempty"`
}

type TurknetAPIGetBBKListResponse struct {
	ServiceResult TurknetAPIServiceResult `json:"ServiceResult"`
	LstDaire      []TurknetAPIIDname      `json:"lstDaire,omitempty"`
}

type TurknetAPICheckServiceAvailabilityResponse struct {
	ServiceResult TurknetAPIServiceResult             `json:"ServiceResult"`
	Result        TurknetAPIServiceAvailabilityResult `json:"Result,omitempty"`
}

func CreateTurknetAPIClient() TurknetAPIClient {
	return TurknetAPIClient{"", "github.com/linuxgemini/turknet-query/cmd TurknetAPIClient/1.1"}
}

func (tac *TurknetAPIClient) GetToken() (err error) {
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result TurknetAPIGetTokenResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return errors.New("failed to unmarshal TurknetAPIGetTokenResponse")
	}

	tac.ClientToken = result.Token

	return
}

func (tac *TurknetAPIClient) _CreateWebRequest(url string, jsonStr []byte) ([]byte, error) {
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (tac *TurknetAPIClient) _CreateGetWebRequest(url string) ([]byte, error) {
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (tac *TurknetAPIClient) _ConvertIDnameToKV(idn []TurknetAPIIDname) TurknetAPIKv {
	var result = make(TurknetAPIKv)
	for _, ta := range idn {
		i, _ := strconv.Atoi(ta.ID)
		result[ta.Name] = i
	}

	return result
}

func (tac *TurknetAPIClient) GetBBKCountyList(ilPlakaKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKCountyList/%d", ilPlakaKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKCountyListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKCountyListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstIlce), nil
}

func (tac *TurknetAPIClient) GetBBKBucakList(ilceKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKBucakList/%d", ilceKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKBucakListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKBucakListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstBucak), nil
}

func (tac *TurknetAPIClient) GetBBKKoyList(bucakKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKKoyList/%d", bucakKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKKoyListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKKoyListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstKoy), nil
}

func (tac *TurknetAPIClient) GetBBKMahalleList(koyKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKMahalleList/%d", koyKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKMahalleListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKMahalleListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstMahalle), nil
}

func (tac *TurknetAPIClient) GetBBKCaddeList(mahalleKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKCaddeList/%d", mahalleKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKCaddeListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKCaddeListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstCadde), nil
}

func (tac *TurknetAPIClient) GetBBKBinaList(caddeKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKBinaList/%d", caddeKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKBinaListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKBinaListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstBina), nil
}

func (tac *TurknetAPIClient) GetBBKList(binaKodu int) (TurknetAPIKv, error) {
	url := fmt.Sprintf("https://turk.net/service/AddressGetServ.svc/GetBBKList/%d", binaKodu)

	body, err := tac._CreateGetWebRequest(url)
	if err != nil {
		return nil, err
	}

	var result TurknetAPIGetBBKListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("failed to unmarshal TurknetAPIGetBBKListResponse")
	}

	return tac._ConvertIDnameToKV(result.LstDaire), nil
}

func (tac *TurknetAPIClient) CheckServiceAvailability(valueType string, value string) (TurknetAPIServiceAvailabilityResult, error) {
	url := "https://turk.net/service/AddressServ.svc/CheckServiceAvailability"

	var jsonStr = []byte(fmt.Sprintf(`{"InquirySource":2,"IsInfrastructureInquiry":true,"Key":"%s","Value":"%s"}`, valueType, value))

	body, err := tac._CreateWebRequest(url, jsonStr)
	if err != nil {
		return TurknetAPIServiceAvailabilityResult{}, err
	}

	var result TurknetAPICheckServiceAvailabilityResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return TurknetAPIServiceAvailabilityResult{}, errors.New("failed to unmarshal TurknetAPICheckServiceAvailabilityResponse")
	}

	return result.Result, nil
}
