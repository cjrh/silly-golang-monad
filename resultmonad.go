package main

import (
	"strings"
	//"errors"
)

type Result interface {
	Bind(f func(v interface{}) (r interface{}, err error)) Result
}

type Ok struct {
	value interface{}
}

type Err struct {
	err error
}

func (m Ok) Bind(f func(v interface{}) (r interface{}, err error)) Result {
	value, err := f(m.value)
	if err != nil {
		return Err{err}
	} else {
		return Ok{value}
	}
}

func (m Err) Bind(f func(v interface{}) (r interface{}, err error)) Result { return m }

func Work(data Result) Result {
	return data.Bind(func(v interface{}) (r interface{}, err error) {
		return strings.ToUpper(v.(string)), nil
	}).Bind(func(v interface{}) (r interface{}, err error) {
		return len(v.(string)) - 5, nil
	}).Bind(func(v interface{}) (r interface{}, err error) {
		r = nil
		err = nil
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}
		}()

		return 10 / v.(int), err
	}).Bind(func(v interface{}) (r interface{}, err error) {
		return v.(int) + 7, nil
	})
}
