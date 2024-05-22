const exprss = require('express')

const app = exprss();
// const protobuf = require('protobufjs')
// const axios = require('axios')
// const {resolve}  = require('path')
const { protoLoader} = require('./proto');

app.get('/user', async (req, res) => {
    // const proto = await protobuf.load(resolve(__dirname, "../proto/user.proto"))
    // const User = proto.lookupType('User')
    //
    // const result = await axios({
    //     url: "http://localhost:8080/protobuf",
    //     method: "get",
    //     responseType: 'arraybuffer'
    // })
    //
    // let user = User.decode(result.data)
    // res.json(user)

    const result = await protoLoader({
        file: "../proto/user.proto",
        api: "http://localhost:8080/protobuf",
        method: 'get',
        type: 'User'
    })
    res.json(result)

})

// async function protoLoader({
//     file,
//     api,
//     method,
//     data,
//     type
//                      }) {
//     try {
//         const proto = await protobuf.load(
//             resolve(__dirname, file)
//         );
//         const protoType = proto.lookupType(type);
//         const result = await axios({
//             url: api,
//             method,
//             data,
//             responseType: 'arraybuffer'
//         })
//
//         let res = protoType.decode(result.data)
//         return res
//     }catch (e) {
//         return Promise.reject(e)
//     }
//
//
// }

app.listen(3000, () => {
    console.log("server is running on port 3000")
});