/**
 * API Wrapper for Turk.net infrastructure query.
 *
 * @author İlteriş Yağıztegin Eroğlu "linuxgemini" <ilteris@asenkron.com.tr>
 * @license MIT
 */

"use strict";

const fetch = require("node-fetch");
const turknetAPIError = require("./turknetAPIError");

/**
 * @typedef {Object} TurknetFiberAvailabilityObject
 * @property {boolean} isAvailable
 * @property {boolean} isGigaFiber
 * @property {boolean} isGigaFiberPlanned
 * @property {number} maxCapacity
 */

/**
 * @typedef {Object} VAEFiberAvailabilityObject
 * @property {boolean} isAvailable
 * @property {number} maxCapacity
 * @property {number} maxCapacityServiceType
 * @property {number} nmsMax
 * @property {number} type
 * @property {string} description
 */

/**
 * @typedef {Object} VDSLAvailabilityObject
 * @property {boolean} isAvailable
 * @property {number} maxCapacity
 * @property {number} maxCapacityServiceType
 * @property {number} nmsMax
 * @property {string} description
 */

/**
 * @typedef {Object} XDSLAvailabilityObject
 * @property {boolean} isAvailable
 * @property {number} maxCapacity
 * @property {number} nmsMax
 * @property {string} description
 */

/**
 * @typedef {Object} YAPAAvailabilityObject
 * @property {boolean} isAvailable
 * @property {boolean} isIndoor
 * @property {boolean} isTurknetActiveOnExchange
 * @property {string} description
 */

/**
 * @typedef {Object} TurknetInfrastructureQueryResult
 * @property {TurknetFiberAvailabilityObject} turknetFiberAvailability
 * @property {VAEFiberAvailabilityObject} vaeFiberAvailability
 * @property {VDSLAvailabilityObject} vdslAvailability
 * @property {XDSLAvailabilityObject} xdslAvailability
 * @property {YAPAAvailabilityObject} yapaAvailability
 */

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
     * Builds request options.
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

    /**
     * Checks if authorized to use API.
     * @returns {Promise<string>} API Access Token.
     */
    async handleToken() {
        if (this._token !== "") return this._token;

        let reqraw = await fetch(this.__tokenEndpoint, this._buildRequestOptions({}));
        let obj = await reqraw.json();
        if (obj.ServiceResult.Code !== 0) throw new turknetAPIError(obj.ServiceResult.Code.toString(), obj.ServiceResult.Message);

        return this._token = obj.Token;
    }

    /**
     * Gets the BBK number list for cities inside given province
     * plate code.
     * @param {number} ilPlateCode Province Plate Code.
     * @returns {Promise<Object.<string, number>>} A dictionary of city names with BBK code values.
     */
    async getIlceler(ilPlateCode) {
        if (isNaN(ilPlateCode) || (ilPlateCode < 0 || ilPlateCode > 81)) throw new Error("Plate Code is not valid!");

        await this.handleToken();

        let reqraw = await fetch(this.__ilceEndpoint, this._buildRequestOptions({
            "IlKod": `${ilPlateCode}`
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

        return ilceler;
    }

    /**
     * Gets the BBK number list for regions inside given city
     * BBK code.
     * @param {number} ilceCode
     * @returns {Promise<Object.<string, number>>} A dictionary of region names with BBK code values.
     */
    async getBucaklar(ilceCode) {
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

        return bucaklar;
    }

    /**
     * Gets the BBK number list for villages inside given region
     * BBK code.
     * @param {number} bucakCode
     * @returns {Promise<Object.<string, number>>} A dictionary of village names with BBK code values.
     */
    async getKoyler(bucakCode) {
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

        return koyler;
    }

    /**
     * Gets the BBK number list for neighborhoods inside given village
     * BBK code.
     * @param {number} koyCode
     * @returns {Promise<Object.<string, number>>} A dictionary of neighborhood names with BBK code values.
     */
    async getMahalleler(koyCode) {
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

        return mahalleler;
    }

    /**
     * Gets the BBK number list for streets inside given neighborhood
     * BBK code.
     * @param {number} mahalleCode
     * @returns {Promise<Object.<string, number>>} A dictionary of street names with BBK code values.
     */
    async getCaddeSokaklar(mahalleCode) {
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

        return caddeler;
    }

    /**
     * Gets the BBK number list for buildings inside given street
     * BBK code.
     * @param {number} caddeCode
     * @returns {Promise<Object.<string, number>>} A dictionary of building numbers with BBK code values.
     */
    async getBinalar(caddeCode) {
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

        return binalar;
    }

    /**
     * Gets the BBK number list for apartments inside given building
     * BBK code.
     * @param {number} binaCode
     * @returns {Promise<Object.<string, number>>} A dictionary of apartment numbers with BBK code values.
     */
    async getDaireler(binaCode) {
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

        return daireler;
    }

    /**
     * Gets the Infrastructure Information for given value.
     *
     * Value can either be:
     *   - Legacy PSTN/POTS number, with type "PSTN"
     *   - BBK number, with type "BBK"
     * @param {"PSTN" | "BBK"} queryType
     * @param {number} queryValue
     * @returns {Promise<TurknetInfrastructureQueryResult>}
     */
    async makeQuery(queryType, queryValue) {
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

        /**
         * @type {TurknetInfrastructureQueryResult}
         */
        let fin = {
            "turknetFiberAvailability": {
                "isAvailable": res.FiberServiceAvailablity.IsAvailable,
                "isGigaFiber": res.FiberServiceAvailablity.IsGigaFiber,
                "isGigaFiberPlanned": res.IsGigaFiberPlanned,
                "maxCapacity": res.FiberServiceAvailablity.MaxCapacity
            },
            "vaeFiberAvailability": {
                "isAvailable": res.VAEFiberServiceAvailability.IsAvailable,
                "maxCapacity": res.VAEFiberServiceAvailability.MaxCapacity,
                "maxCapacityServiceType": res.VAEFiberServiceAvailability.MaxCapacityServiceType,
                "nmsMax": res.VAEFiberServiceAvailability.NmsMax,
                "type": res.VAEFiberServiceAvailability.Type,
                "description": res.VAEFiberServiceAvailability.Description || ""
            },
            "vdslAvailability": {
                "isAvailable": res.VDSLServiceAvailability.IsAvailable,
                "maxCapacity": res.VDSLServiceAvailability.MaxCapacity,
                "maxCapacityServiceType": res.VDSLServiceAvailability.MaxCapacityServiceType,
                "nmsMax": res.VDSLServiceAvailability.NmsMax,
                "description": res.VDSLServiceAvailability.Description || ""
            },
            "xdslAvailability": {
                "isAvailable": res.XDSLServiceAvailability.IsAvailable,
                "maxCapacity": res.XDSLServiceAvailability.MaxCapacity,
                "nmsMax": res.XDSLServiceAvailability.NmsMax,
                "description": res.XDSLServiceAvailability.Description || ""
            },
            "yapaAvailability": {
                "isAvailable": res.YapaServiceAvailability.IsAvailable,
                "isIndoor": res.YapaServiceAvailability.IsIndoor,
                "isTurknetActiveOnExchange": res.YapaServiceAvailability.IsTurknetStatusActiveForSantral,
                "description": res.YapaServiceAvailability.Description || ""
            }

        };

        return fin;
    }
}

module.exports = TurknetQuery;
