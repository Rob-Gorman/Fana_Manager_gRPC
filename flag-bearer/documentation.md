# POST `/flagset` | Receve Updated Flag Data (Full Ruleset)
- expects to receive entire set of flags and audiences for all SDK keys from Flag Manager

## Expected Payload
- each SDK ruleset requires `sdkKey` and `flags`
- `flags` is empty array if no flags created yet
- `status` represents if flag is on (`true`) or off (`false`)
- evaluation of flags will default to whether ANY audience logic is met for a user
```js
[
  {
    sdkKey: str,
    flags: [ 
      {
        flagKey: String,
        status: bool, 
        audiences: [audienceKey]
      }
    ],
    audiences: [
      {
        audienceKey: String,
        combination: String, // "ANY" or "ALL"
        conditions: [
          {
          attribute: String, // name of attribute
          type: String, // type of attribute, defined in data model
          operator: String,
          value: String
          },
        ],
      }
    ]
  }
]
```
## Example payload from flag manager to flag bearer when one flag is updated
- `sdkKey` and `flag` object with `name` is required
- `status` and `audiences` are optional (if updated) 
```js
{
  flag: {
    sdkKey: str,
    name: str,
    status: bool,
    audiences: []
  }
}
```

# GET `/subscribe/client` | Event Subscription to Flag Updates  
## Example Event Message
- Note that SDKs should listen for events with a `sdkKey` that match the client SDK key, and can ignore all other messages
- SSE events are sent when Flag Bearer receives updated flag data from Flag Manager
- Push event only sends flags that are toggled off
- `flags` property contains an array of flag objects with an updated `status` value (will always be `false`) and `value`
```js
[
  { sdkKey: 'beta_sdk_0',
    flags: [
      {
        flagKey: 'new-button',
        status: false,
        value: false,
      },
    ]
  },
  { sdkKey: 'beta_sdk_1',
    flags: [
      {
        flagKey: 'new-nav-bar',
          status: false,
          value: false,
      },
    ]
  }
]
```
# POST `/connect/clientInit` | Initialize Client SDK
Client SDK initialization request route. Returns an object of flag names and the flag evaluation for a specific user context.
- When receiving a client SDK initialization request, Flag bearer will parse through full flag dataset to identify flags pertaining to specific instance (`sdkKey`).
- Memoizes evaluated flags in-memory 

## Expected Payload
- expects `sdkKey` and a `userContext` object
  - `userContext` must provide a `userId` 

```js
{
  sdkKey: String,
  userContext: {
    userId: String,
    // ... optional attributes
  },
}
```
## Example Successful Response
- returns object of key/value pairs, where `flagKey` is the key, and evaluation `value` is a boolean. 
  - 
```js
{
  flagName1: Boolean,
  flagName2: Boolean,
  // ... more flags
} 
```
### Error Response
- `400` status code will be returned if SDK key or userId is not provided or invalid. 


## Flag Bearer Cache: Cache data structure
- memoized flag evaluation values organized by SDK key, then userId, then unique flag name
```js
// Shape of cache object
const cache = {
  sdkKey1: {
    userId1: {
      flageName1: VALUE_OF_FLAG (bool),
      flageName2: VALUE_OF_FLAG (bool),
    } ,
    userId2 : { ... }
  },
  sdkKey2 : { ... }
}

// Example cache lookup for a user -- allows O(1) lookup
const userFlagValues = cache[sdkKey][userId]

{
  flagName1: true,
  flageName2: false,
  ...
}
```
# POST `/connect/serverInit'` | Initialize Server-side SDK
- returns unevaluated flags and audience conditions
## Expected Payload
- expects `sdkKey`
```js
{
  sdkKey: String,
}
```
## Example Successful Response
- sends the flag data for server-side sdk instance
```js
 {
    sdkKey: 'beta_sdk_0',
    flags: [
      {
        flagKey: 'flag-evals-true',
        status: true,
        audiences: ['beta-testers', 'california_students'],
      },
      {
        flagKey: 'flag-evals-false',
        status: true,
        audiences: ['california_students'],
      },
      {
        flagKey: 'no-audiences-flag',
        status: true,
        audiences: [],
      },
      {
        flagKey: 'toggled-off-flag',
        status: false,
        audiences: ['beta-testers'],
      },
    ],
    audiences: [
      {
        audienceKey: 'beta-testers',
        combination: 'ANY',
        conditions: [
          {
            attribute: 'userId',
            type: 'STR',
            operator: 'EQ',
            value: 'jjuy',
          },
          {
            attribute: 'beta',
            type: 'BOOL',
            operator: 'EQ',
            value: true,
          },
        ],
      },
      {
        audienceKey: 'california_students',
        combination: 'ALL',
        conditions: [
          {
            attribute: 'state',
            type: 'STR',
            operator: 'EQ',
            value: 'california',
          },
          {
            attribute: 'userId',
            type: 'STR',
            operator: 'EQ',
            value: 'jjuy',
          },
        ],
      },
    ],
  }
  ```
  ### Error Response
- `400` status code will be returned if SDK key is not provided or invalid. 
```js
{
  error: 'SDK key is required.'
}
```