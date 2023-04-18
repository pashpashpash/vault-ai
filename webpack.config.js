var path = require('path');
const webpack = require('webpack');

const BUILD_DIR = path.resolve(__dirname, 'static/js');
const APP_DIR = path.resolve(__dirname, './');

const COMPILED_DEPS = [
    APP_DIR,
    path.resolve(__dirname, 'node_modules/keccak'),
    path.resolve(__dirname, 'node_modules/secp256k1'),
    path.resolve(__dirname, 'node_modules/eth-block-tracker'),
    path.resolve(__dirname, 'node_modules/ansi-styles'),
    path.resolve(__dirname, 'node_modules/asn1.js'),
    path.resolve(__dirname, 'node_modules/asn1'),
    path.resolve(__dirname, 'node_modules/parse-asn1'),
    path.resolve(__dirname, 'node_modules/chalk'),
    path.resolve(__dirname, 'node_modules/merkle-patricia-tree'),
    path.resolve(__dirname, 'node_modules/ethereum-bloom-filters'),
];

const IS_DEV = process.env.NODE_ENV !== 'production';

var config = {
    mode: IS_DEV ? 'development' : 'production',
    entry: {
        app: APP_DIR + '/components/index.jsx',
    },
    output: {
        path: BUILD_DIR,
        filename: 'bundle.js',
    },
    module: {
        rules: [
            {
                test: /\.(png|jpg|gif|svg|eot|ttf|woff|woff2)$/,
                type: 'asset/resource',
            },
            {
                test: /\.(jsx|js)$/,
                include: [APP_DIR],
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'babel-loader',
                        options: {
                            presets: [
                                '@babel/preset-env',
                                '@babel/preset-react',
                                '@babel/preset-flow',
                            ],
                            plugins: [
                                '@babel/plugin-proposal-class-properties',
                                '@babel/transform-runtime',
                            ],
                        },
                    },
                ],
            },
            {
                test: /\.(css|less)$/i,
                use: [
                    {
                        loader: 'style-loader',
                    },
                    {
                        loader: 'css-loader',
                        options: {
                            url: false,
                            modules: {
                                exportLocalsConvention: 'camelCase',
                                localIdentName: `${
                                    IS_DEV ? '[local]_' : ''
                                }[hash:base64:5]`,
                            },
                            sourceMap: IS_DEV,
                        },
                    },
                    {
                        loader: 'less-loader',
                        options: {
                            lessOptions: {
                                strictMath: true,
                                sourceMap: IS_DEV,
                            },
                        },
                    },
                ],
            },
        ],
    },
    resolve: {
        modules: ['node_modules', APP_DIR],
        extensions: ['.jsx', '.js'],
        alias: {
            process: 'process/browser',
        },
        fallback: {
            assert: require.resolve('assert'),
            buffer: require.resolve('buffer'),
            crypto: require.resolve('crypto-browserify'),
            os: require.resolve('os-browserify/browser'),
            http: require.resolve('stream-http'),
            https: require.resolve('https-browserify'),
            util: require.resolve('util'),
            stream: require.resolve('stream-browserify'),
            'react-select': false,
        },
    },
    plugins: [
        new webpack.ProvidePlugin({
            Buffer: ['buffer', 'Buffer'],
        }),
        new webpack.ProvidePlugin({
            process: 'process/browser',
        }),
    ],
};

module.exports = config;
