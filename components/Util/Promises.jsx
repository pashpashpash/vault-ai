// Promise related helpers

const Promises = {
    sleep: ms => {
        return new Promise(resolve => setTimeout(resolve, ms));
    },

    makeCancelable: promise => {
        let hasCanceled_ = false;

        const wrappedPromise = new Promise((resolve, reject) => {
            promise.then(
                val =>
                    hasCanceled_ ? reject({ isCanceled: true }) : resolve(val),
                error =>
                    hasCanceled_ ? reject({ isCanceled: true }) : reject(error)
            );
        });

        return {
            promise: wrappedPromise,
            cancel() {
                hasCanceled_ = true;
            },
        };
    },

    resolvedPromise: ret =>
        new Promise((resolve, reject) => {
            resolve(ret);
        }),

    failPromise: message =>
        new Promise((resolve, reject) => {
            reject(new Error(message));
        }),
};

export default Promises;
