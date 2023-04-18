// @flow

const Text = {
    prettyEthAccount: (account: string, chunkSize: ?number): string => {
        chunkSize = chunkSize || 4;
        const len = account.length;
        return (
            account.slice(0, chunkSize + 2) +
            '...' +
            account.slice(len - chunkSize)
        );
    },

    // converts chainID into a pretty network name
    prettyChainName: (chainID: number): string => {
        if (chainID === 137) return 'Polygon';
        return 'Ethereum';
    },
};

export default Text;
