const protobuf = require("protobufjs");
const {resolve} = require("path");
const axios = require("axios");

async function protoLoader({
                               file,
                               api,
                               method,
                               data,
                               type
                           }) {
    try {
        const proto = await protobuf.load(
            resolve(__dirname, file)
        );
        const protoType = proto.lookupType(type);
        const result = await axios({
            url: api,
            method,
            data,
            responseType: 'arraybuffer'
        })

        let res = protoType.decode(result.data)
        return res
    }catch (e) {
        return Promise.reject(e)
    }
}

module.exports = {
    protoLoader
};