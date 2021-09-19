package jsonparse

import (
	"encoding/json"
	"errors"
	"fmt"
)

type JSON struct {
	// v initial json data
	v interface{}

	// array json array
	array []interface{}

	// object json object
	object map[string]interface{}

	// fail to get information
	fail bool

	// failMessage message about fail to get or interface typing
	failMessage string
}

func New(bb []byte) (JSON, error) {
	j := JSON{}
	if err := json.Unmarshal(bb, &j.v); err != nil {
		return JSON{}, err
	}

	return j, nil
}

func (j JSON) Array() JSON {
	if j.fail {
		return j
	}

	arr, ok := j.v.([]interface{})
	if !ok {
		j.setFail(fmt.Sprintf("ошибка преобразования %v", j.v))
	}else{
		j.array = arr
		j.v = arr
	}

	return j
}

func (j JSON) Object() JSON {
	if j.fail {
		return j
	}

	obj, ok := j.v.(map[string]interface{})
	if !ok {
		j.setFail(fmt.Sprintf("ошибка преобразования %v", j.v))
	}else{
		j.object = obj
		j.v = obj
	}

	return j
}

func (j JSON) Select(v interface{}) JSON {
	switch vv := v.(type) {
	case string:
		if _, ok := j.object[vv]; !ok {
			j.setFail(fmt.Sprintf("ошибка при получении информации по индексу %s", vv))
		}else {
			j.v = j.object[vv]
		}
	case int:
		if i := j.array[vv]; i == nil {
			j.setFail(fmt.Sprintf("ошибка при получении информации по индексу %d", vv))
		}else{
			j.v = j.array[vv]
		}
	}
	return j
}

func (j JSON) Fail() error {
	if j.fail {
		return errors.New(j.failMessage)
	}
	return nil
}

func (j JSON) Get(v interface{}) (interface{}, error) {
	var (
		rt interface{}
		ok bool
	)

	switch vv := v.(type) {
	case string:
		rt, ok = j.object[vv]
		if !ok {
			return nil, errors.New("по данному ключу не найдено информации")
		}
	case int:
		rt = j.array[vv]
		if rt == nil {
			return nil, errors.New("по данному ключу не найдено информации")
		}
	}

	return rt, j.Fail()
}

func (j JSON) setFail(m string)  {
	j.fail = true
	j.failMessage = m
}