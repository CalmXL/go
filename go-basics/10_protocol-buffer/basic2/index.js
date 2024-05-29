{ 
  Syntax: proto3, 
  Path: "myInfo.proto", 
  Messages: [
    { 
      Name: MyInfoRequest, 
      Fields: [
        {
          Name: name,
          Number: 1, 
          Cardinality: optional,
          Kind: string,
          HasJSONName: true,
          JSONName: "name" 
        },
        {
          Name: age, 
          Number: 2, 
          Cardinality: optional,
          Kind: int32,
          HasJSONName: true,
          JSONName: "age" 
         },
        { 
          Name: isMarriage,
          Number: 3, 
          Cardinality: optional, 
          Kind: bool, 
          HasJSONName: true,
          JSONName: "isMarriage" 
        }
      ]
    }, 
    { 
      Name: MyInfoResponse,
      Fields: [
        { 
          Name: name, 
          Number: 1, 
          Cardinality: optional,
          Kind: string,
          HasJSONName: true,
          JSONName: "name" 
        }, 
        { 
          Name: age, 
          Number: 2,
          Cardinality: optional,
          Kind: int32,
          HasJSONName: true,
          JSONName: "age"
        }, 
        { 
          Name: isMarriage,
           Number: 3, 
           Cardinality: optional,
            Kind: bool, 
            HasJSONName: true,
            JSONName: "isMarriage" 
          }
        ] 
      }
    ],
    Services: [
      { 
        Name: MyInfo, 
        Methods: [
          { 
            Name: getData, 
            Input: MyInfoRequest, 
            Output: MyInfoResponse 
          }
        ] 
      }
    ] 
  }
