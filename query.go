package appwrite

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct{}

func (q Query) Equal(attribute string, value interface{}) string {
	return q.addQuery(attribute, "equal", value)
}

func (q Query) NotEqual(attribute string, value interface{}) string {
	return q.addQuery(attribute, "notEqual", value)
}

func (q Query) LessThan(attribute string, value interface{}) string {
	return q.addQuery(attribute, "lessThan", value)
}

func (q Query) LessThanEqual(attribute string, value interface{}) string {
	return q.addQuery(attribute, "lessThanEqual", value)
}

func (q Query) GreaterThan(attribute string, value interface{}) string {
	return q.addQuery(attribute, "greaterThan", value)
}

func (q Query) GreaterThanEqual(attribute string, value interface{}) string {
	return q.addQuery(attribute, "greaterThanEqual", value)
}

func (q Query) Search(attribute string, value interface{}) string {
	return q.addQuery(attribute, "search", value)
}

func (q Query) OrderDesc(attribute string) string {
	return fmt.Sprintf(`orderDesc("%s")`, attribute)
}

func (q Query) OrderAsc(attribute string) string {
	return fmt.Sprintf(`orderAsc("%s")`, attribute)
}

func (q Query) CursorAfter(documentId string) string {
	return fmt.Sprintf(`cursorAfter("%s")`, documentId)
}

func (q Query) CursorBefore(documentId string) string {
	return fmt.Sprintf(`cursorBefore("%s")`, documentId)
}

func (q Query) Limit(limit int) string {
	return fmt.Sprintf(`limit(%d)`, limit)
}

func (q Query) Offset(offset int) string {
	return fmt.Sprintf(`offset(%d)`, offset)
}

func (q Query) addQuery(attribute string, method string, value interface{}) string {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		values := reflect.ValueOf(value)
		strValues := make([]string, values.Len())
		for i := 0; i < values.Len(); i++ {
			strValues[i] = q.parseValues(values.Index(i).Interface())
		}
		return fmt.Sprintf(`%s("%s", [%s])`, method, attribute, strings.Join(strValues, ","))
	}
	return fmt.Sprintf(`%s("%s", [%s])`, method, attribute, q.parseValues(value))
}

func (Query) parseValues(value interface{}) string {
	switch value.(type) {
	case string:
		return fmt.Sprintf(`"%s"`, value)
	default:
		return fmt.Sprint(value)
	}
}
