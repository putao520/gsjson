# gsjson

gsjson is an easy to use JSON library for go. It is designed just like [Fastjson in Java](https://github.com/alibaba/fastjson).

## Install

```shell
go get github.com/putao520/gsjson
```

## Import

```go
import (
  "github.com/putao520/gsjson"
)
```

## Usage

### Create a JsonObject from string

```go
jsonObject, err := ParseObject(jsonStr)
```

### Create a JsonArray from string

```go
jsonArray, err := ParseArray(jsonStr)
```

### Get a Child JsonObject

```go
jsonObject := parentObject.GetJsonObject("user")
jsonObject := parentArray.GetJsonObject(0)
```

### Get values

```go
name, err := parentObject.GetString("name")
age, err := parentObject.GetInt("age")
id, err := parentObject.GetInt64("id")
isMarried, err := parentObject.GetBoolean("is_married")
```

### Get values which could be `null`

```go
age, err := parentObject.GetNullInt("age")
id, err := parentObject.GetNullLong("id")
isMarried, err := parentObject.GetNullBoolean("is_married")
```

### Chaining call

```go
balance, err := parentObject.GetJsonObject("user").GetJsonObject("account").GetFloat("balance")
email, err := parentObject.GetJsonObject("user").GetJsonArray("emails").GetString(0)
```

### JsonObject map to Struct
```go
testStruct, _ := MapperObject[StructName](jsonObject)
testStructArray, _ := MapperArray[StructName](jsonArray)
```
### Struct map to JsonObject
```go
jsonObject, _ := StructObject(testStruct)
jsonArray _ := StructArray(testStructArray)
```

## Example


```go
package main

import (
  "fmt"
  "github.com/putao520/gsjson"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `{
		"user": {
			"name": "aa",
			"age": 10,
			"phone": "12222222222",
			"emails": [
				"aa@164.com",
				"aa@165.com"
			],
			"address": [
				{
					"number": "101",
					"now_live": true
				},
				{
					"number": "102",
					"now_live": null
				}
			],
			"account": {
				"balance": 999.9
			}
		}
	}
	`
	jsonObject, err := gsjson.ParseObject(jsonStr)
	fmt.Println(jsonObject, err)

	user := jsonObject.GetJsonObject("user")
	fmt.Println(user)

	name, err := user.GetString("name")
	fmt.Println(name, err)

	phone, err := user.GetInt64("phone")
	fmt.Println(phone, err)

	age, err := user.GetInt64("age")
	fmt.Println(age, err)

	account := user.GetJsonObject("account")
	fmt.Println(account)

	balance, err := account.GetFloat64("balance")
	fmt.Println(balance, err)

	email1, err := user.GetJsonArray("emails").GetString(0)
	fmt.Println(email1, err)

	address := user.GetJsonArray("address")
	fmt.Println(address)

	address1nowLive, err := user.GetJsonArray("address").GetJsonObject(0).GetBoolean("now_live")
	fmt.Println(address1nowLive, err)

	address2, err := address.Get(1)
	fmt.Println(address2, err)

	address2NowLive, err := address2.JsonObject().GetNullBoolean("now_live")
	fmt.Println(address2NowLive, err)
	
	k, _ := address2.iter() {
		fmt.Println(k)
	}
	
	fmt.Println(address2.ToString())

	fmt.Println(user.ToString())

	j := gsjson.New()
	
	j2 := gsjson.FromObject(j)
	
	ja := gsjson.NewArray()
	
	ja2 := gsjson.FromArray(ja)
	
	for k, v := range jsonObject.Values() {
		fmt.Println("key:" + k)
		fmt.Println("val:" + v)
    }
    
    j.replace(jsonObject)
	
	j.Clear()
	
	j.remove("address")
	
	size := len(jsonObject.Values())

	fmt.Println(jsonObject)
	
	str := jsonObject.String()
	fmt.Println(str)
	
}

```
