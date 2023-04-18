/* eslint-disable flowtype/no-weak-types */
// @flow
// Handy helpers + caching for Nugbase Post API

import Promises from './Promises';

import TimedLocalStorage from './TimedLocalStorage';
const API_TIMEOUT = 1000 * 60 * 10; // 10 min in ms

const questions = {
    submitQuestion: (question: string, model: string): Promise<any> => {
        const resultPromise = new Promise((resolve: any, reject: any) => {
            var data = new FormData();
            data.append('question', question);
            data.append('model', model);

            const timeoutID = setTimeout(
                function() {
                    console.log('[/api/questions] Timeout:', question);
                    reject(new Error('Timeout'));
                }.bind(this),
                API_TIMEOUT
            );
            fetch('/api/questions', {
                method: 'POST',
                body: data,
            })
                .then((res: any): any => {
                    console.log(
                        `[/api/questions RESPONSE] STAT: ${res.status} | OK: ${res.ok}`
                    );
                    clearTimeout(timeoutID);
                    if (res.ok) return res.json();

                    res.text().then((text: string) => {
                        console.log('>>>>>>status ok false', { text });
                        reject(text);
                    });
                })
                .catch((err: Error): void => {
                    console.log('[/api/questions] Error sending POST', err);
                    reject(err);
                })
                .then((responseData: any) => {
                    console.log('>>>>>>hello?', { responseData });
                    if (responseData == null) {
                        if (!res.ok) return; // Do not reject the promise with a new error if it was already rejected
                        reject(new Error('404'));
                        return;
                    }
                    try {
                        console.log('[/api/questions] Success', responseData);
                        resolve(responseData);
                    } catch (err) {
                        console.log(
                            '[/api/questions ERR] unable to handle',
                            err
                        );
                        reject(err);
                    }
                });
        });

        return Promises.makeCancelable(resultPromise);
    },
};

const upload = {
    uploadFiles: (files: Array<File>): Promise<any> => {
        const resultPromise = new Promise((resolve: any, reject: any) => {
            var data = new FormData();
            files.forEach(file => {
                data.append('files', file);
            });

            const timeoutID = setTimeout(() => {
                console.log('[/upload] Timeout:', files);
                reject(new Error('Timeout'));
            }, API_TIMEOUT);

            fetch('/upload', {
                method: 'POST',
                body: data,
            })
                .then((res: any): any => {
                    console.log(
                        `[/upload RESPONSE] STAT: ${res.status} | OK: ${res.ok}`
                    );
                    clearTimeout(timeoutID);
                    if (res.ok) return res.json();

                    return res.text().then((text: string) => {
                        throw new Error(res.status + ' | ' + text);
                    });
                })
                .then((responseData: any) => {
                    console.log('>>>>>>hello?', { responseData });
                    if (responseData == null) {
                        reject(new Error('404'));
                        return;
                    }
                    console.log('[/upload] Success', responseData);
                    resolve(responseData);
                })
                .catch((err: Error): void => {
                    console.log('[/upload] Error', err);
                    reject(err);
                });
        });

        return Promises.makeCancelable(resultPromise);
    },
};

const PostAPI = {
    questions,
    upload,
};

export default PostAPI;
