# GORM
[29K Stars on Github ORM for Go](https://gorm.io/docs/index.html)
This is like Mongoose, but for SQL systems, and in Go

It more tightly couples your code with your DB schema.
Primarily because it has to: Go is statically-typed and so there's more requisite structure to get a reasonable workflow

These are the notes I jotted down on how we define the structure of our tables and establish the relationships between them such that our program ultimately reflects the nature of our database.

## Go Structs
Think JavaScript Objects; it's a field and a value.
If it's all the same type (and/or unbounded in size), that's a map:
```go
var kindaJSON = map[string]int{"a": 1, "b":2}
```
But maps can only accommodate specifically Ta:Tb for every element (`String:Int` above)

That's not flexible enough to represent database tables.
In order to have multiple different types represented in your fields, there are custom structs.
These are rigidly defined, and strictly enforced, because Go is statically-typed.
Each instance has to be composed of exactly the same fields and types as their definition
```go
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
}
```
Worth mentioning, any field can have a type of another struct, even self-referencing:
```go
type User struct {
  ID           uint
  Name         string
  Friends     []User  // pointers not all that relevant in this project, i think
}
```

## Models
[Models docs](https://gorm.io/docs/models.html)
Models are just Go structs. GORM (and EVERY DB package in Go) uses them to represent tables.
GORM's framework basically translates these structs to SQL schema and associated qualities.
*Tags are used for constrains and foreign references*
*However* a lot of _implied_ functionality via syntactic sugar for SQL conventions:
- ID is primary key (don't need this, see below)
- field and struct names are CamelCase in Go, but translated to snake_case in postgres

gorm.Model is a convenience stand-alone struct to put in a model that adds the following:
```go
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```
including `gorm.Model` as a field itself in the struct definition of your model will simply add those fields to the table schema

## Associations
[Many:many docs](https://gorm.io/docs/many_to_many.html)
The only ones we care about, I think, are many-to-many.
In GORM, you declare it in tags:
```go
type Person struct {
  ID        int
  Name      string
  Addresses []Address `gorm:"many2many:person_addresses;"`
}

type Address struct {
  ID   uint
  Name string
}

// HOWEVER, you can also declare the table outright
// maybe to add more fields to the join (like flag status)
// or add hooks (read below)
// UNCLEAR if the `SetupJoinTable` is required if you have the
// foresight to do it at the outset. I think it would.
type PersonAddress struct {
  PersonID  int `gorm:"primaryKey"`
  AddressID int `gorm:"primaryKey"`
  CreatedAt time.Time
  DeletedAt gorm.DeletedAt
}

// Change model Person's field Addresses' join table to PersonAddress
// PersonAddress must defined all required foreign keys or it will raise error
err := db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})
```
This is the really ORM-y stuff that most obscures the actual SQL happening.

## Hooks
These should be super useful for keeping clean our automatically pushed updates.

GORM hooks: model (struct) methods with pre-defined names that are invoked whenever the method name's associated CRUD action is executed

```go
func (fl *Flag) AfterCreate(db *gorm.DB) err {
  // eg., push update to flag bearer
  return
}
```

## Queries
TBD
[Queries Docs](https://gorm.io/docs/create.html)

