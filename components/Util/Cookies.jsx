/* eslint-disable flowtype/no-weak-types */
// @flow
const getCookie = (name: string): any => {
    const match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
    if (match) return match[2];
    return null;
};

const Cookies = {
    getCookie,
    isLoggedIn: (account: string): any => getCookie('NUGPUB-' + account),
};

export default Cookies;
