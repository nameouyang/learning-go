package validator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
)

func init() {
	binding.Validator = new(defaultValidator)
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

func (v *defaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}

func kindOfData(d interface{}) reflect.Kind {
	valueOf := reflect.ValueOf(d)
	valueType := valueOf.Kind()
	if valueType == reflect.Ptr {
		valueType = valueOf.Elem().Kind()
	}
	return valueType
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

//实现接口
var _ binding.StructValidator = &defaultValidator{}
