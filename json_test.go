package gsjson

import (
	"fmt"
	"testing"
)

func TestParseObject(t *testing.T) {
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
	jsonObject, err := ParseObject(jsonStr)
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

	jsonStrCmp := jsonObject.String()
	fmt.Println(jsonStrCmp)

	// fast
	userObj := jsonObject.GetJsonObject("user")
	userObj.Put("test1", 123)
	fmt.Println(jsonObject.String())
	userObj.Put("test2", "ttt")
	fmt.Println(jsonObject.String())

	// cmp json
	jsonObjCmp, _ := ParseObject(jsonStrCmp)
	if !jsonObjCmp.compareTo(jsonObject) {
		t.Errorf("json compare false-> String faild")
	}
	// replace json
	nJson := New()
	nJson.Replace(jsonObject)
	fmt.Println(nJson.String())
}

func TestParseArray(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `[
		{
			"name": "Will"
		},
		{
			"name": "Uzi"
		},
	"asd",123,9.8,null,false
	]
	`
	jsonArray, err := ParseArray(jsonStr)
	fmt.Println(jsonArray, err)

	name1, err := jsonArray.GetJsonObject(0).GetString("name")
	fmt.Println(name1, err)

	arrayStrCmp := jsonArray.String()
	fmt.Println(arrayStrCmp)

	arrayObjCmp, _ := ParseArray(arrayStrCmp)
	if !arrayObjCmp.compareTo(jsonArray) {
		t.Errorf("jsonArray compare false-> String faild")
	}

	userObj := jsonArray.GetJsonObject(1)
	userObj.Put("test1", New().Put("caption", "value"))
	fmt.Println(jsonArray.String())
	jsonArray.Put("test2")
	fmt.Println(jsonArray.String())

	// replace jsonArray
	nArray := NewArray()
	nArray.Replace(jsonArray)
	fmt.Println(nArray.String())
}

type test1 struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Timestamp int64  `json:"timestamp"`
}

func TestMapJson(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `{
		"name": "aa",
		"age": 10,
		"timestamp": 12222222222
	}
	`
	j, _ := ParseObject(jsonStr)
	structTest1, _ := MapperObject[test1](j)
	fmt.Println(structTest1)

	jsonArrayStr := `[{
		"name": "aa",
		"age": 10,
		"timestamp": 12222222222
	}, {
		"name": "bb",
		"age": 12,
		"timestamp": 21111111111
	}]
	`
	a, _ := ParseArray(jsonArrayStr)
	structArrayTest1, _ := MapperArray[test1](a)
	fmt.Println(structArrayTest1)

	jo, _ := StructObject(structTest1)
	println(jo.String())

	ja, _ := StructArray(structArrayTest1)
	println(ja.String())
}
