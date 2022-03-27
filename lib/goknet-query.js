/**
 * Partial API Wrapper for Goknet infrastructure query.
 *
 * @author İlteriş Yağıztegin Eroğlu "linuxgemini" <ilteris@asenkron.com.tr>
 * @license MIT
 */

"use strict";

//const goknetAPIError = require("./goknetAPIError");
const fetch = require("node-fetch");

class GoknetQuery {
    constructor() {}

    /**
     * Builds request options.
     */
    _buildRequestOptions() {
        return {
            headers: {
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
                "Referer": "https://user.goknet.com.tr/inactive/service_availability.php"
            }
        };
    }

    _convertToReadableKV(name, value) {
        if (name === "BSPRT") {
            name = "bosPort";
            value = (value === "1" ? true : false);
        }

        if (name === "IPDSLM") {
            name = "ipDSLAM";
            value = (value === "1" ? true : false);
        }

        if (name === "SNTRLHZMT") {
            name = "santralHizmeti";
            value = (value === "1" ? true : false);
        }

        if (name === "IPTVHZMT") {
            name = "iptvHizmeti";
            value = (value === "1" ? true : false);
        }

        if (name === "ACKISEMRI") {
            name = "acikIsEmri";
            value = value.replace(/\|/g, "").trim() || "";
        }

        if (name === "GREENBROWN") {
            name = "tabloRengi";
            value = (value === "G" ? "yesil" : (value === "B" ? "kahverengi" : "renksiz"));
        }

        if (name === "ISFTTC") {
            name = "isFTTC";
            value = (value === "1" ? true : false);
        }
        if (name === "ISINDOOR") {
            name = "isIndoor";
            value = (value === "1" ? true : false);
        }
        if (name === "ISIPVOK") {
            name = "isIPVOK";
            value = (value === "1" ? true : false);
        }

        if (name === "NDSLX") value = (value === "1" ? true : false);
        if (name === "FIBERX") value = (value === "1" ? true : false);

        if (name === "SNTRLT11") name = "t11Santral";
        if (name === "SNTRLPRTT11") name = "t11SantralPortu";
        if (name === "DSLMXSPD") name = "dslMaxSpeed";
        if (name === "IPVMXSPD") name = "ipvMaxSpeed";
        if (name === "NMSMAX") name = "nmsMax";
        if (name === "FTTXTYPE") name = "fttxTipi";
        if (name === "SNTRLMSF") name = "santralMesafe";

        if (value === "N/A") value = null;

        return [name, value];
    }

    async makeQuery(queryType, queryValue) {
        if (!(["BBK"].includes(queryType.toUpperCase()))) throw new Error("Not a valid query type!");
        if (isNaN(queryValue)) throw new Error("Not a valid query value!");

        let reqraw = await fetch(`https://user.goknet.com.tr/sistem/getTTAddressWebservice.php?kod=${queryValue}&datatype=checkAddress`, this._buildRequestOptions());
        if (!reqraw.ok) throw new Error("API Request Failure");

        let obj = {
            adsl: {},
            vdsl: {},
            ftth: {}
        };
        let objraw = await reqraw.json();

        obj.adsl.errorCode = objraw["1"].hataKod;
        obj.adsl.errorMessage = objraw["1"].hataMesaj;
        for (const resultItem of objraw["1"].flexList.flexList) {
            let kv = this._convertToReadableKV(resultItem.name, resultItem.value);
            if (kv[0] === "tabloRengi" && obj.adsl.errorCode !== "100") kv[1] = "renksiz";
            obj.adsl[kv[0]] = kv[1];
        }

        obj.vdsl.errorCode = objraw["6"].hataKod;
        obj.vdsl.errorMessage = objraw["6"].hataMesaj;
        for (const resultItem of objraw["6"].flexList.flexList) {
            let kv = this._convertToReadableKV(resultItem.name, resultItem.value);
            if (kv[0] === "tabloRengi" && obj.vdsl.errorCode !== "100") kv[1] = "renksiz";
            obj.vdsl[kv[0]] = kv[1];
        }

        obj.ftth.errorCode = objraw["7"].hataKod;
        obj.ftth.errorMessage = objraw["7"].hataMesaj;
        for (const resultItem of objraw["7"].flexList.flexList) {
            let kv = this._convertToReadableKV(resultItem.name, resultItem.value);
            if (kv[0] === "tabloRengi" && obj.ftth.errorCode !== "100") kv[1] = "renksiz";
            obj.ftth[kv[0]] = kv[1];
        }

        return obj;
    }
}

module.exports = GoknetQuery;
