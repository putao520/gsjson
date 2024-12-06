//

package gsjson

import (
	"encoding/json"
)

func New() *JsonObject {
	j, _ := ParseObject("{}")
	return j
}

func ParseObject(jsonStr string) (*JsonObject, error) {
	value := new(Value)

	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}

	return value.JsonObject(), nil
}

func FromObject(json *JsonObject) *JsonObject {
	j := New()
	return j.Replace(json)
}

func NewArray() *JsonArray {
	ja, _ := ParseArray("[]")
	return ja
}

func FromArray(json *JsonArray) *JsonArray {
	j := NewArray()
	return j.Replace(json)
}

func ParseArray(jsonStr string) (*JsonArray, error) {
	value := new(Value)
	err := json.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}
	return value.JsonArray(), nil
}

//func Mapper[T any](jsonStr string) *T {
//	var result T
//	err := json.Unmarshal([]byte(jsonStr), &result)
//	if err != nil {
//		panic(err)
//	}
//	return &result
//}

func MapperObject[T any](j *JsonObject) (*T, error) {
	return MapperObjectString[T](j.String())
}

func MapperArray[T any](j *JsonArray) (*[]T, error) {
	return MapperArrayString[T](j.String())
}

func MapperObjectString[T any](s string) (*T, error) {
	var result T
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func MapperArrayString[T any](s string) (*[]T, error) {
	var result []T
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func StructObject[T any](j *T) (*JsonObject, error) {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	r, err := ParseObject(string(jsonBytes))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func StructToString[T any](j *T) (string, error) {
	r, err := StructObject(j)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}

func StructArray[T any](j *T) (*JsonArray, error) {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	r, err := ParseArray(string(jsonBytes))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func StructArrayString[T any](j *T) (string, error) {
	r, err := StructArray(j)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}
