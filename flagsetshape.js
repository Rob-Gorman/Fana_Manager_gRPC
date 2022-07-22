{
  "sdkKeys": {
    "c3e-db3100c-8":true,
    "de9-6bf1a0c-3":true,
    "fa4-d731f0e-4":true
  },
  
  "flags": {
    "0": {
      "status":false,
      "audiences":[
        {
          "combine":"ANY",
          "conditions":[{
            "attribute":"state",
            "operator":"EQ",
            "vals":["california"],
            "negate":false
          },
          {
            "attribute":"student",
            "operator":"EQ",
            "vals":["true"],
            "negate":false
          }]
        }]
      },
    "1":{
      "status":false,
      "audiences":[]
    },
    "2":{
      "status":false,
      "audiences":[
        {
          "combine":"ANY",
          "conditions":[
            {
              "attribute":"beta",
              "operator":"EQ",
              "vals":["true"],
              "negate":false
            }]
          }]
        }
      }
    }