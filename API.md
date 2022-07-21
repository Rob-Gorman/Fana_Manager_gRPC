Conditions.vals []string

# API

## Attributes
### Create Attribute

POST /api/attributes

Expected Payload:
```js
{
    "key": "country",
    "attrType": "STR"
}
```

Expected Response:
```js


```

## Flags

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