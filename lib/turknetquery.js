/**
 * API Wrapper for Turk.net speed query.
 * 
 * @author linuxgemini
 * @license MIT
 */

"use strict";

const fetch = require("node-fetch");
const turknetAPIError = require("./turknetAPIError");

class TurknetQuery {
    constructor() {
        this.__tokenEndpoint = "https://turk.net/service/AddressServ.svc/GetToken";
        this.__ilceEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKCountyList";
        this.__bucakEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKBucakList";
        this.__koyEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKKoyList";
        this.__mahalleEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKMahalleList";
        this.__caddesokakEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKCaddeList";
        this.__binaEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKBinaList";
        this.__daireEndpoint = "https://turk.net/service/AddressServ.svc/GetBBKList";
        this.__queryEndpoint = "https://turk.net/service/AddressServ.svc/CheckServiceAvailability";
        this.__requestType = "PUT";
        this._token = "";
    }

    /**
     * @param {Object|Array} body 
     */
    _buildRequestOptions(body) {
        return {
            method: this.__requestType,
            headers: {
                "Content-Type": "application/json",
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
                "Referer": "https://turk.net/internet-hiz-altyapi-sorgulama/",
                "Origin": "https://turk.net",
                ...(this._token !== "" && {"Token": this._token})
            },
            body: JSON.stringify(body)
        };
    }

    async handleToken() {
        try {
            if (this._token !== "") return Promise.resolve(this._token);

            let reqraw = await fetch(this.__tokenEndpoint, this._buildRequestOptions({}));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            return Promise.resolve(this._token = obj.Token);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} plateCode 
     */
    async handleIl(plateCode) {
        try {
            if (isNaN(plateCode) || (plateCode < 0 || plateCode > 81)) throw new Error("Plate Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__ilceEndpoint, this._buildRequestOptions({
                "IlKod": `${plateCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let ilceler = {};

            for (const ilceObj of obj.lstIlce) {
                ilceler[ilceObj.Name.replace("\u000a", "")] = parseInt(ilceObj.Id);
            }

            return Promise.resolve(ilceler);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} ilceCode 
     */
    async handleIlce(ilceCode) {
        try {
            if (isNaN(ilceCode)) throw new Error("Ilce Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__bucakEndpoint, this._buildRequestOptions({
                "IlceKod": `${ilceCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let bucaklar = {};

            for (const bucakObj of obj.lstBucak) {
                bucaklar[bucakObj.Name.replace("\u000a", "")] = parseInt(bucakObj.Id);
            }

            return Promise.resolve(bucaklar);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} bucakCode 
     */
    async handleBucak(bucakCode) {
        try {
            if (isNaN(bucakCode)) throw new Error("Bucak Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__koyEndpoint, this._buildRequestOptions({
                "BucakKod": `${bucakCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let koyler = {};

            for (const koyObj of obj.lstKoy) {
                koyler[koyObj.Name.replace("\u000a", "")] = parseInt(koyObj.Id);
            }

            return Promise.resolve(koyler);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} koyCode 
     */
    async handleKoy(koyCode) {
        try {
            if (isNaN(koyCode)) throw new Error("Koy Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__mahalleEndpoint, this._buildRequestOptions({
                "KoyKod": `${koyCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let mahalleler = {};

            for (const mahalleObj of obj.lstMahalle) {
                mahalleler[mahalleObj.Name.replace("\u000a", "")] = parseInt(mahalleObj.Id);
            }

            return Promise.resolve(mahalleler);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} mahalleCode 
     */
    async handleMahalle(mahalleCode) {
        try {
            if (isNaN(mahalleCode)) throw new Error("Mahalle Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__caddesokakEndpoint, this._buildRequestOptions({
                "MahalleKod": `${mahalleCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let caddeler = {};

            for (const caddeObj of obj.lstCadde) {
                caddeler[caddeObj.Name.replace("\u000a", "")] = parseInt(caddeObj.Id);
            }

            return Promise.resolve(caddeler);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} caddeCode 
     */
    async handleCadde(caddeCode) {
        try {
            if (isNaN(caddeCode)) throw new Error("Cadde Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__binaEndpoint, this._buildRequestOptions({
                "CaddeKod": `${caddeCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let binalar = {};

            for (const binaObj of obj.lstBina) {
                binalar[binaObj.Name.replace("\u000a", "")] = parseInt(binaObj.Id);
            }

            return Promise.resolve(binalar);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {number} binaCode 
     */
    async handleBina(binaCode) {
        try {
            if (isNaN(binaCode)) throw new Error("Bina Code is not valid!");

            await this.handleToken();

            let reqraw = await fetch(this.__daireEndpoint, this._buildRequestOptions({
                "BinaKod": `${binaCode}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            /**
             * @type {{string: number}}
             */
            let daireler = {};

            for (const daireObj of obj.lstDaire) {
                daireler[daireObj.Name.replace("\u000a", "")] = parseInt(daireObj.Id);
            }

            return Promise.resolve(daireler);
        } catch (error) {
            return Promise.reject(error);
        }
    }

    /**
     * 
     * @param {"PSTN" | "BBK"} queryType 
     * @param {number} queryValue 
     */
    async makeQuery(queryType, queryValue) {
        try {
            if (!(["PSTN", "BBK"].includes(queryType.toUpperCase()))) throw new Error("Not a valid query type!");
            if (isNaN(queryValue)) throw new Error("Not a valid query value!");

            await this.handleToken();

            let reqraw = await fetch(this.__queryEndpoint, this._buildRequestOptions({
                "Key": queryType.toUpperCase(),
                "Value": `${queryValue}`
            }));
            let obj = await reqraw.json();
            if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

            let res = obj.Result;

            let fin = {
                "turknetFiberAvailability": {
                    /** @type {boolean} */
                    "isAvailable": res.FiberServiceAvailablity.IsAvailable,
                    /** @type {boolean} */
                    "isGigaFiber": res.FiberServiceAvailablity.IsGigaFiber,
                    /** @type {number} */
                    "maxCapacity": res.FiberServiceAvailablity.MaxCapacity
                },
                "vaeFiberAvailability": {
                    /** @type {boolean} */
                    "isAvailable": res.VAEFiberServiceAvailability.IsAvailable,
                    /** @type {number} */
                    "maxCapacity": res.VAEFiberServiceAvailability.MaxCapacity,
                    /** @type {number} */
                    "maxCapacityServiceType": res.VAEFiberServiceAvailability.MaxCapacityServiceType,
                    /** @type {number} */
                    "nmsMax": res.VAEFiberServiceAvailability.NmsMax,
                    /** @type {number} */
                    "type": res.VAEFiberServiceAvailability.Type,
                    /** @type {string?} */
                    "description": res.VAEFiberServiceAvailability.Description
                },
                "vdslAvailability": {
                    /** @type {boolean} */
                    "isAvailable": res.VDSLServiceAvailability.IsAvailable,
                    /** @type {number} */
                    "maxCapacity": res.VDSLServiceAvailability.MaxCapacity,
                    /** @type {number} */
                    "maxCapacityServiceType": res.VDSLServiceAvailability.MaxCapacityServiceType,
                    /** @type {number} */
                    "nmsMax": res.VDSLServiceAvailability.NmsMax,
                    /** @type {string?} */
                    "description": res.VDSLServiceAvailability.Description
                },
                "xdslAvailability": {
                    /** @type {boolean} */
                    "isAvailable": res.XDSLServiceAvailability.IsAvailable,
                    /** @type {number} */
                    "maxCapacity": res.XDSLServiceAvailability.MaxCapacity,
                    /** @type {number} */
                    "nmsMax": res.XDSLServiceAvailability.NmsMax,
                    /** @type {string?} */
                    "description": res.XDSLServiceAvailability.Description
                },
                "yapaAvailability": {
                    /** @type {boolean} */
                    "isAvailable": res.YapaServiceAvailability.IsAvailable,
                    /** @type {boolean} */
                    "isIndoor": res.YapaServiceAvailability.IsIndoor,
                    /** @type {boolean} */
                    "isTurknetActiveOnExchange": res.YapaServiceAvailability.IsTurknetStatusActiveForSantral,
                    /** @type {string?} */
                    "description": res.YapaServiceAvailability.Description
                }

            };

            return Promise.resolve(fin);
        } catch (error) {
            return Promise.reject(error);
        }
    }
}

module.exports = TurknetQuery;
