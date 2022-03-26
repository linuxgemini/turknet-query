/**
 * Simple error wrapper
 * 
 * @author linuxgemini
 * @license MIT
 */

"use strict";

class goknetAPIError extends Error {
    constructor(errcode = "0000", errmessage = "Critical code error on wrapper, contact @linuxgemini", ...params) {
        // Pass remaining arguments (including vendor specific ones) to parent constructor
        super(...params);

        // Maintains proper stack trace for where our error was thrown (only available on V8)
        if (Error.captureStackTrace) {
            Error.captureStackTrace(this, goknetAPIError);
        }

        // Custom debugging information
        this.code = errcode;
        this.message = errmessage;
        this.date = new Date();
        this.type = "goknetAPIError";
    }
}

module.exports = goknetAPIError;
 