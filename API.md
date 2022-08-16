# Manager - Dashboard API

[Attributes](#attributes)
- [Get Attributes](#get-all-attributes)
- [Get One Attribute](#get-attribute)
- [Create Attribute](#create-attribute)
- [Delete Attribute](#delete-attribute)

[Audiences](#audiences)
- [Get Audiences](#get-audiences)
- [Get One Audience](#get-single-audience)
- [Create Audience](#create-audience)
- [Update Audience](#update-audience)
- [Delete Audience](#delete-audience)

[Flags](#flags)
- [Get Flags](#get-flags)
- [Get One Flag](#get-single-flag)
- [Create Flag](#create-flag)
- [Toggle Flag](#toggle-flag)
- [Update Flag](#update-flag)
- [Delete Flag](#delete-flag)

[SDK Keys](#sdk-keys)
- [Get SDK Keys](#get-sdk-keys)
- [Regenerate SDK Key](#regenerate-sdk-key)

[Audit Logs](#audit-logs)
- [Get Audit Logs](#get-audit-logs)

## Attributes

### Get All Attributes

GET /api/attributes

Expected Response:
```js
[
  {
    "id": 1,
    "key": "state",
    "attrType": "STR",
    "displayName": "State",
    "Conditions": null,
    "created_at": "2022-08-16T19:34:24.037795Z",
    "deleted_at": null
  },
  {
    "id": 2,
    "key": "student",
    "attrType": "BOOL",
    "displayName": "Student",
    "Conditions": null,
    "created_at": "2022-08-16T19:34:24.037795Z",
    "deleted_at": null
  }
]
```

### Get Attribute

GET /api/attributes/{id}

Expected Response:
```js
{
  "id": 1,
  "key": "state",
  "attrType": "STR",
  "displayName": "State",
  "created_at": "2022-08-16T19:34:24.037795Z",
  "deleted_at": null,
  "audiences": [
    {
      "id": 1,
      "displayName": "California Students",
      "key": "california_students",
      "created_at": "2022-08-16T19:34:24.041163Z",
      "updated_at": "2022-08-16T19:34:24.04882Z"
    }
  ]
}
```

### Create Attribute

POST /api/attributes

Expected Payload:
```js
{
    "key": "country",
    "displayName": "Country",
    "attrType": "STR"
}
```

Expected Response:
```js
{
  "id": 4,
  "key": "country",
  "type": "STR",
  "displayName": "Country",
  "created_at": "2022-07-21T10:07:54.1756404-04:00"
}

```

### Delete Attribute

DELETE /api/attributes/{id}

Expected Response: 204 No Content

## Audiences

### Get Audiences

GET /api/audiences

Expected Response:
```js
[
  {
    "id": 1,
    "displayName": "California Students",
    "key": "california_students",
    "created_at": "2022-07-21T14:38:06.198091Z",
    "updated_at": "2022-07-21T14:38:06.202699Z"
  },
  {
    "id": 2,
    "displayName": "Beta Testers",
    "key": "beta_testers",
    "created_at": "2022-07-21T14:38:06.198091Z",
    "updated_at": "2022-07-21T14:38:06.202699Z"
  }
]
```

### Get Single Audience

GET /api/audiences/{id}

Expected Response:
```js
{
  "id": 1,
  "displayName": "California Students",
  "key": "california_students",
  "combine": "ANY",
  "created_at": "2022-07-21T16:32:43.229267Z",
  "updated_at": "2022-07-21T16:32:43.234051Z",
  "conditions": [
    {
      "id": 1,
      "negate": false,
      "operator": "EQ",
      "vals": "california",
      "attribute": {
        "id": 1,
        "key": "state",
        "type": "STR",
        "displayName": "State"
      }
    },
    {
      "id": 2,
      "negate": false,
      "operator": "EQ",
      "vals": "true",
      "attribute": {
        "id": 2,
        "key": "student",
        "type": "BOOL",
        "displayName": "Student"
      }
    }
  ]
}
```

### Create Audience

POST /api/audiences

Expected Payload:
```js
{
  "key": "texas_students",
  "displayName": "TX Students",
  "combine": "ALL",
  "conditions": [
      {
      "attributeID": 1,
      "operator": "EQ",
      "vals": "texas"
      },
      {
      "attributeID": 2,
      "operator": "EQ",
      "vals": "true",
      "negate": true
      }        
  ]
}
```

Expected Response:
```js
{
  "id": 5,
  "displayName": "TX Students",
  "key": "texas_students",
  "combine": "ALL",
  "created_at": "2022-07-22T14:59:08.041869Z",
  "updated_at": "2022-07-22T14:59:08.041869Z",
  "conditions": [
    {
      "negate": false,
      "operator": "EQ",
      "attribute": "state",
      "vals": [
        "texas"
      ]
    },
    {
      "negate": true,
      "operator": "EQ",
      "attribute": "student",
      "vals": [
        "true"
      ]
    }
  ]
}
```

### Update Audience

PATCH /api/audiences/{id}

Expected Payload:
```js
{
  "displayName": "NewName",
  "combine": "ALL",
  "conditions": [
      {
      "attributeID": 1,
      "operator": "EQ",
      "vals": "texas"
      },
      {
      "attributeID": 2,
      "operator": "EQ",
      "vals": "true",
      "negate": true
      }        
  ]
}
```

Expected Response:
```js
{
  "id": 1,
  "displayName": "NewName",
  "key": "california_students",
  "combine": "ALL",
  "created_at": "2022-08-16T19:56:01.358514Z",
  "updated_at": "2022-08-16T19:56:04.663571Z",
  "conditions": [
    {
      "negate": false,
      "attributeID": 1,
      "operator": "EQ",
      "vals": "texas",
      "attribute": "state"
    },
    {
      "negate": true,
      "attributeID": 2,
      "operator": "EQ",
      "vals": "true",
      "attribute": "student"
    }
  ],
  "flags": [
    {
      "id": 1,
      "key": "fake-flag-1",
      "displayName": "FAKE FLAG ONE",
      "status": false,
      "created_at": "2022-08-16T19:56:01.350562Z",
      "updated_at": "2022-08-16T19:56:01.362871Z"
    }
  ]
}
```

### Delete Audience

DELETE /api/audiences/{id}

Expected Response: 204 No Content

## Flags

### Get Flags

GET /api/flags

Expected Response:
```js
[
  {
    "id": 1,
    "key": "fake-flag-1",
    "displayName": "FAKE FLAG ONE",
    "status": false,
    "created_at": "2022-07-21T14:38:06.191632Z",
    "updated_at": "2022-07-21T14:38:06.202338Z"
  },
  {
    "id": 2,
    "key": "experimental-flag-1",
    "displayName": "Exp Fl 1",
    "status": false,
    "created_at": "2022-07-21T14:38:06.191632Z",
    "updated_at": "2022-07-21T14:38:06.191632Z"
  }
]
```

### Get Single Flag

GET /api/flags/{id}

Expected Response:
```js
{
  "id": 1,
  "key": "fake-flag-1",
  "displayName": "FAKE FLAG ONE",
  "status": false,
  "created_at": "2022-07-21T14:38:06.191632Z",
  "updated_at": "2022-07-21T14:38:06.202338Z",
  "audiences": [
    {
      "id": 1,
      "displayName": "California Students",
      "key": "california_students",
      "created_at": "2022-07-21T14:38:06.198091Z",
      "updated_at": "2022-07-21T14:38:06.202699Z"
    }
  ]
}
```

### Create Flag

POST /api/audiences

Expected Payload:
```js
{
    "key": "inauthentic_flag_2",
    "displayName": "Inauthentic Flag 2",
    "sdkKey": "beta_sdk_0",
    "audiences": []
}
```

Expected Response:
```js
{
  "id": 4,
  "key": "inauthentic_flag_2",
  "displayName": "Inauthentic Flag 2",
  "status": false,
  "created_at": "2022-08-16T20:02:22.94274Z",
  "updated_at": "2022-08-16T20:02:22.94274Z",
  "audiences": []
}
```

### Toggle Flag

PATCH /api/flags/{id}/toggle

Expected Payload
```js
{
    "status": true
}
```

Expected Response:
```js
{
  "id": 2,
  "key": "experimental-flag-1",
  "displayName": "",
  "sdkKey": "not_used_sdk_key",
  "status": true,
  "created_at": "2022-07-21T00:38:31.661643Z",
  "updated_at": "2022-07-21T00:43:14.121991Z"
}
```

### Update Flag

PATCH /api/flags/{id}

Expected Payload:
```js
{
    "displayName": "Is a real flag",
    "audiences": ["california_students", "beta_testers"]
}
```

Expected Response:
```js
{
  "id": 1,
  "key": "fake-flag-1",
  "displayName": "Is a real flag",
  "status": false,
  "created_at": "2022-07-21T13:54:06.463222Z",
  "updated_at": "2022-07-21T13:54:15.706588Z",
  "audiences": [
    {
      "id": 1,
      "displayName": "",
      "key": "california_students",
      "combine": "ANY",
      "created_at": "2022-07-21T13:54:06.467348Z",
      "updated_at": "2022-07-21T13:54:15.706932Z"
    },
    {
      "id": 2,
      "displayName": "",
      "key": "beta_testers",
      "combine": "ANY",
      "created_at": "2022-07-21T13:54:06.467348Z",
      "updated_at": "2022-07-21T13:54:15.706932Z"
    }
  ]
}
```

### Delete Flag

DELETE /api/flags/{id}

Expected Response: 204 No Content

## SDK Keys

### Get SDK Keys

GET /api/sdkkeys

Expected Response:
```js
[
  {
    "id": 1,
    "key": "1f7-b169c-84",
    "status": true,
    "type": "client",
    "created_at": "2022-08-16T19:56:01.36771Z",
    "updated_at": "2022-08-16T19:56:01.36771Z",
    "deleted_at": null
  },
  {
    "id": 2,
    "key": "6f2-18ab5-52",
    "status": true,
    "type": "server",
    "created_at": "2022-08-16T19:56:01.36771Z",
    "updated_at": "2022-08-16T19:56:01.36771Z",
    "deleted_at": null
  }
]
```

### Regenerate SDK Key

DELETE /api/sdkkeys/{2}

Expected Response:
```js
{
  "id": 3,
  "key": "fa8-2fbf8-67",
  "status": true,
  "type": "server",
  "created_at": "2022-08-16T20:19:17.420347Z",
  "updated_at": "2022-08-16T20:19:17.420347Z",
  "deleted_at": null
}
```

## Audit Logs

### Get Audit Logs

GET /api/auditlogs

Expected Response:
```js
{
  "flagLogs": [
    {
      "logID": 1,
      "id": 1,
      "key": "fake-flag-1",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.962036Z"
    },
    {
      "logID": 2,
      "id": 2,
      "key": "experimental-flag-1",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.962586Z"
    },
    {
      "logID": 3,
      "id": 3,
      "key": "development-flag-1",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.962892Z"
    },
    {
      "logID": 4,
      "id": 1,
      "key": "fake-flag-1",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.976404Z"
    },
    {
      "logID": 5,
      "id": 3,
      "key": "development-flag-1",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.976699Z"
    }
  ],
  "audienceLogs": [
    {
      "logID": 1,
      "id": 1,
      "key": "california_students",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.970743Z"
    },
    {
      "logID": 2,
      "id": 2,
      "key": "beta_testers",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.971246Z"
    },
    {
      "logID": 3,
      "id": 1,
      "key": "california_students",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.975065Z"
    },
    {
      "logID": 4,
      "id": 2,
      "key": "beta_testers",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.975347Z"
    }
  ],
  "attributeLogs": [
    {
      "logID": 1,
      "id": 1,
      "key": "state",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.964977Z"
    },
    {
      "logID": 2,
      "id": 2,
      "key": "student",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.965507Z"
    },
    {
      "logID": 3,
      "id": 3,
      "key": "beta",
      "action": "created",
      "created_at": "2022-08-16T20:24:00.965938Z"
    }
  ]
}
```
[Back To Top](#manager---dashboard-api)