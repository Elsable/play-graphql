package common

import (
	"github.com/graphql-go/graphql"
	"log"
	"reflect"
	"time"
)

func ObjToFields(i interface{}) (graphql.Fields) {
	gFields := make(graphql.Fields)
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		v := &graphql.Field{
			Name: typ.Field(i).Name,
			Description: typ.Field(i).Name,
		}
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v.Type = graphql.Int
		case uint, uint8, uint16, uint32, uint64:
			v.Type = graphql.Int
		case float32, float64:
			v.Type = graphql.Float
		case []byte:
			v.Type = graphql.NewList(graphql.Int)
		case string:
			v.Type = graphql.String
		case []string:
			v.Type = graphql.NewList(graphql.String)
		case time.Time, *time.Time:
			v.Type = graphql.DateTime
		}
		//if _, ok := gFields[typ.Field(i).Name]; !ok {
		//	gFields[typ.Field(i).Name] = &graphql.Field{}
		//}
		log.Printf("%v %v",typ.Field(i).Name, f.Kind())
		gFields[ typ.Field(i).Tag.Get("json")] = v
	}
	return gFields
}