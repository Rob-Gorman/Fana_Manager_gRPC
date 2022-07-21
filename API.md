# API

## Attributes
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

## Audiences (PATCH and POST 1 sec...)

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

### Toggle Flag

PATCH /api/flags/{id}/toggle

Expected Payload
```js
{
    "status": true
}
```

Expected Response:
**OPINION: response should just be 200, no body**
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