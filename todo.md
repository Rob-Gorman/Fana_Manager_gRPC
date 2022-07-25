# Todo GoManager
## Populate Associations:
- Audiences
- Condition Attributes

## FLAG BEARER ENDPOINT
- [See Below](#flag-bearer-get-endpoint)

## Push Updates
- AfterUpdate hooks
- AfterCreate hooks

## Audit Trail Pages
Seed first

## API Contract
Manager - Dashboard

```js
[
  {
    date: string,
    event_description: string
  }
]
```

## Flag Bearer GET endpoint
**Data Structure:**
```js
{
  sdkKeys: [
    "sdk_key_1",
    "sdk_key_2"
  ],
  flag_key_1: {
      audience_key_1: {
        combine: "ANY",
        conditions: [
          condition_1, // this is an object
          condition_2
        ]
      },
      audience_key_2: {
        combine: "ALL",
        conditions: [
          condition_2,
          condition_3
        ]
      }
  }
  flag_key_2: {
    combine: "ALL",
    audiences
    ...
  }
}
```

```js
condition_1 = {
  "attribute": "student",
  "operator": "EQ",
  "value": "true",
  "negate": false
}
```


### Eventually
- How to manage secrets besides `.env`?


#### Hacks thusfar:
- condition vals intake as a string
  - postgres doesn't hold arrays
  - figure out how to translate to different struct
- subrouter/pathprefix just hardcoded (there's a better way)


~~## ROUTES WIP~~
~~s.HandleFunc("/api/flags", s.H.GetAllFlags).Methods("GET")~~
~~s.HandleFunc("/api/flags", s.H.CreateFlag).Methods("POST")~~
~~s.HandleFunc("/api/flags/{id}", s.H.GetFlag).Methods("GET")~~
~~s.HandleFunc("/api/flags/{id}", s.H.UpdateFlag).Methods("PATCH")~~
~~s.HandleFunc("/api/flags/{id}/toggle", s.H.ToggleFlag).Methods("PATCH")~~

~~s.HandleFunc("/api/audiences", s.H.GetAllAudiences).Methods("GET")~~
~~s.HandleFunc("/api/audiences", s.H.CreateAudience).Methods("POST")~~
~~s.HandleFunc("/api/audiences/{id}", s.H.GetAudience).Methods("GET")~~

~~s.HandleFunc("/api/attributes", s.H.GetAllAttributes).Methods("GET")~~
~~s.HandleFunc("/api/attributes", s.H.CreateAttribute).Methods("POST")~~


UPDATE STRUCTURE ==> redis publish only updates in shape of bearer object