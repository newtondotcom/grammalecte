// Grammalecte - Thésaurus

/* jshint esversion:6 */
/* jslint esversion:6 */
/* global __dirname */

"use strict";

${string}


if (typeof(process) !== 'undefined') {
    var helpers = require("../graphspell/helpers.js");
}


var thesaurus = {
    _dWord: new Map(),

    bInit: false,
    init: function (sJSONData1, sJSONData2) {
        try {
            // As addons.mozilla.org doesn’t accept file bigger than 5 Mb,
            // we had to split the thesaurus in two parts. And now we merge them.
            let _oData1 = JSON.parse(sJSONData1);
            let _oData2 = JSON.parse(sJSONData2);
            let _oData = { ..._oData1, ..._oData2 };
            // convert to Map
            this._dWord = helpers.objectToMap(_oData);
            this.bInit = true;
            //console.log(this._dWord);
        }
        catch (e) {
            console.error(e);
        }
    },

    getSyns: function (sWord) {
        // return list of synonyms of <sWord>
        if (!sWord) {
            return [];
        }
        if (this._dWord.has(sWord)) {
            return this._dWord.get(sWord);
        }
        if (sWord.slice(0,1).gl_isUpperCase()) {
            sWord = sWord.toLowerCase();
            if (this._dWord.has(sWord)) {
                return this._dWord.get(sWord);
            }
        }
        return [];
    }
};



// Initialization
if (!thesaurus.bInit && typeof(process) !== 'undefined') {
    // NodeJS
    thesaurus.init(helpers.loadFile(__dirname+"/thesaurus1_data.json"), helpers.loadFile(__dirname+"/thesaurus2_data.json"));
} else if (!thesaurus.bInit && typeof(browser) !== 'undefined') {
    // WebExtension Standard (but not in Worker)
    thesaurus.init(helpers.loadFile(browser.runtime.getURL("grammalecte/fr/thesaurus1_data.json")), helpers.loadFile(browser.runtime.getURL("grammalecte/fr/thesaurus2_data.json")));
} else if (!thesaurus.bInit && typeof(chrome) !== 'undefined') {
    // WebExtension Chrome (but not in Worker)
    thesaurus.init(helpers.loadFile(chrome.runtime.getURL("grammalecte/fr/thesaurus1_data.json")), helpers.loadFile(chrome.runtime.getURL("grammalecte/fr/thesaurus2_data.json")));
} else if (thesaurus.bInit){
    console.log("Module thesaurus déjà initialisé");
} else {
    //console.log("Module thesaurus non initialisé");
}


if (typeof(exports) !== 'undefined') {
    exports._dWord = thesaurus._dWord;
    exports.init = thesaurus.init;
    exports.getSyns = thesaurus.getSyns;
}
