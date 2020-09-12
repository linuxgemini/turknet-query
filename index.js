#!/usr/bin/env node

/* eslint-disable no-case-declarations */

/**
 * Turknet Query CLI
 * 
 * @author linuxgemini
 * @license MIT
 */

"use strict";

const fs = require("fs");
const https = require("https");
const path = require("path");
https.globalAgent.options.ca = fs.readFileSync(path.resolve(__dirname, "./ca_bundle/ca_intermediate_root_bundle.pem"));

const program = require("commander");
const inquirer = require("inquirer"); 
const chalk = require("chalk");
const iller = require("./iller.json");
const queryapi = require("./lib/turknetquery");

const exit = () => {
    return setTimeout(() => {
        process.exit(0);
    }, 1000);
};

const exitWithError = (err) => {
    console.error("\n\nBir hata meydana geldi!\n");
    if (err.stack) console.error(`\nStacktrace:\n${err.stack}\n`);
    return setTimeout(() => {
        process.exit(1);
    }, 1000);
};

const main = () => {
    program
        .version("0.1.5", "-v, --version")
        .description("Adresiniz veya telefon numaranızı girerek Türknet altyapı durumunu sorgulayın!");

    program
        .action(async () => {
            try {
                console.log(chalk.yellow("Gerekli paketler yükleniyor, lütfen bekleyin...\n"));
                let api = new queryapi();

                let tipPrompt = await inquirer.prompt([
                    {
                        "type": "list",
                        "name": "method",
                        "message": "Lütfen sorgulama metodunu seçin:",
                        "choices": [
                            "Adres",
                            "Telefon Numarası"
                        ],
                        "default": 0
                    }
                ]);
    
                let res;
    
                switch (tipPrompt.method) {
                    case "Adres":
                        let ilPrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "ilAdi",
                                "message": "Lütfen ilinizi seçin:",
                                "choices": Object.keys(iller)
                            }
                        ]);
    
                        let ilceler = await api.handleIl(iller[ilPrompt.ilAdi]);
                        let ilcePrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "ilceAdi",
                                "message": "Lütfen ilçenizi seçin:",
                                "choices": Object.keys(ilceler)
                            }
                        ]);
    
                        let bucaklar = await api.handleIlce(ilceler[ilcePrompt.ilceAdi]);
                        let useBucak;
                        if (Object.keys(bucaklar).length === 1) {
                            useBucak = bucaklar[Object.keys(bucaklar)[0]];
                        } else {
                            let bucakPrompt = await inquirer.prompt([
                                {
                                    "type": "list",
                                    "name": "bucakAdi",
                                    "message": "Lütfen bucağınızı seçin:",
                                    "choices": Object.keys(bucaklar)
                                }
                            ]);
                            useBucak = bucaklar[bucakPrompt.bucakAdi];
                        }
    
                        let koyler = await api.handleBucak(useBucak);
                        let useKoy;
                        if (Object.keys(koyler).length === 1) {
                            useKoy = koyler[Object.keys(koyler)[0]];
                        } else {
                            let koyPrompt = await inquirer.prompt([
                                {
                                    "type": "list",
                                    "name": "koyAdi",
                                    "message": "Lütfen köyünüzü seçin:",
                                    "choices": Object.keys(koyler)
                                }
                            ]);
                            useKoy = koyler[koyPrompt.koyAdi];
                        }
    
                        let mahalleler = await api.handleKoy(useKoy);
                        let mahallePrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "mahalleAdi",
                                "message": "Lütfen mahallenizi seçin:",
                                "choices": Object.keys(mahalleler)
                            }
                        ]);
    
                        let caddeSokaklar = await api.handleMahalle(mahalleler[mahallePrompt.mahalleAdi]);
                        let caddeSokakPrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "caddeSokakAdi",
                                "message": "Lütfen caddenizi/sokağınızı seçin:",
                                "choices": Object.keys(caddeSokaklar)
                            }
                        ]);
    
                        let binalar = await api.handleCadde(caddeSokaklar[caddeSokakPrompt.caddeSokakAdi]);
                        let binaPrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "binaAdi",
                                "message": "Lütfen bina numaranızı seçin:",
                                "choices": Object.keys(binalar)
                            }
                        ]);
    
                        let daireler = await api.handleBina(binalar[binaPrompt.binaAdi]);
                        let dairePrompt = await inquirer.prompt([
                            {
                                "type": "list",
                                "name": "daireAdi",
                                "message": "Lütfen daire numaranızı seçin:",
                                "choices": Object.keys(daireler)
                            }
                        ]);
    
                        res = await api.makeQuery("BBK", daireler[dairePrompt.daireAdi]);
                        break;
                    case "Telefon Numarası":
                        let telefonPrompt = await inquirer.prompt([
                            {
                                "type": "input",
                                "name": "telno",
                                "message": "Lütfen sabit telefon numaranızı **başında 0 olmadan** girin:",
                                "validate": (i) => {
                                    if (!(i.match(/^([2-4][1-9][1-9])(\d{7})$/g))) return false;
                                    return true;
                                }
                            }
                        ]);
    
                        res = await api.makeQuery("PSTN", telefonPrompt.telno);
                        break;
                    default:
                        break;
                }
                if (res) {
                    console.log(`
${chalk.cyan("Türknet Fiber Durumu:")}
    Var mı?: ${(res.turknetFiberAvailability.isAvailable ? chalk.green("Evet") : chalk.red("Hayır"))}
    GigaFiber mi?: ${(res.turknetFiberAvailability.isGigaFiber ? chalk.green("Evet") : chalk.red("Hayır"))}
    Maksimum kapasite: ${res.turknetFiberAvailability.maxCapacity}
${chalk.cyan("VAE Fiber Durumu:")}
    Var mı?: ${(res.vaeFiberAvailability.isAvailable ? chalk.green("Evet") : chalk.red("Hayır"))}
    Maksimum kapasite: ${res.vaeFiberAvailability.maxCapacity}
    Maksimum kapasite servis tipi: ${res.vaeFiberAvailability.maxCapacityServiceType}
    NmsMax: ${res.vaeFiberAvailability.nmsMax}
    Tip: ${res.vaeFiberAvailability.type}
    Açıklama: ${(res.vaeFiberAvailability.description === null ? chalk.red("Yok") : res.vaeFiberAvailability.description)}
${chalk.cyan("VDSL Durumu:")}
    Var mı?: ${(res.vdslAvailability.isAvailable ? chalk.green("Evet") : chalk.red("Hayır"))}
    Maksimum kapasite: ${res.vdslAvailability.maxCapacity}
    Maksimum kapasite servis tipi: ${res.vdslAvailability.maxCapacityServiceType}
    NmsMax: ${res.vdslAvailability.nmsMax}
    Açıklama: ${(res.vdslAvailability.description === null ? chalk.red("Yok") : res.vdslAvailability.description)}
${chalk.cyan("xDSL Durumu:")}
    Var mı?: ${(res.xdslAvailability.isAvailable ? chalk.green("Evet") : chalk.red("Hayır"))}
    Maksimum kapasite: ${res.xdslAvailability.maxCapacity}
    NmsMax: ${res.vdslAvailability.nmsMax}
    Açıklama: ${(res.xdslAvailability.description === null ? chalk.red("Yok") : res.xdslAvailability.description)}
${chalk.cyan("YAPA Durumu:")}
    Var mı?: ${(res.yapaAvailability.isAvailable ? chalk.green("Evet") : chalk.red("Hayır"))}
    "Indoor" mu?: ${(res.yapaAvailability.isIndoor ? chalk.green("Evet") : chalk.red("Hayır"))}
    Türknet santralde aktif mi?: ${(res.yapaAvailability.isTurknetActiveOnExchange ? chalk.green("Evet") : chalk.red("Hayır"))}
    Açıklama: ${(res.yapaAvailability.description === null ? chalk.red("Yok") : res.yapaAvailability.description)}
`);
                }
                exit();
            } catch (error) {
                exitWithError(error);
            }
        });

    program.parse(process.argv);
};

main();
