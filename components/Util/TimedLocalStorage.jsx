// LocalStorage wrapper that auto-expires old entries

import Text from './Text';

const TimedLocalStorage = {
    save: (key, jsonData, expirationMS) => {
        if (typeof Storage === 'undefined') return false;

        var record = {
            value: JSON.stringify(jsonData),
            timestamp: new Date().getTime() + expirationMS,
        };

        try {
            localStorage.setItem(key, JSON.stringify(record));
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }

        return true;
    },

    load: key => {
        if (typeof Storage === 'undefined') return false;

        let record = null;
        try {
            record = JSON.parse(localStorage.getItem(key));
            if (!record) return false;
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }

        if (new Date().getTime() <= record.timestamp) {
            return JSON.parse(record.value);
        }

        // Record has expired, drop it
        try {
            localStorage.removeItem(key);
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }
        return false;
    },

    saveArrayBuffer: (key, arrayBuffer, expirationMS) => {
        if (typeof Storage === 'undefined') return false;

        const stringified = Text.ab2str(arrayBuffer);
        var record = {
            value: stringified,
            timestamp: new Date().getTime() + expirationMS,
        };

        try {
            localStorage.setItem(key, JSON.stringify(record));
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }

        return true;
    },

    loadArrayBuffer: key => {
        if (typeof Storage === 'undefined') return false;

        let record = null;
        try {
            record = JSON.parse(localStorage.getItem(key));
            if (!record) return false;
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }

        if (new Date().getTime() <= record.timestamp) {
            return Text.str2ab(record.value);
        }

        // Record has expired, drop it
        try {
            localStorage.removeItem(key);
        } catch (e) {
            console.log('[TimedLocalStorage ERR]', e);
            return false;
        }

        return false;
    },
};

export default TimedLocalStorage;
