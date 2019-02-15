package logging

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// ContextPrintf - Formats and logs a string, while also prepending it with context.
// "contexts" can be either a string or a string slice.
func ContextPrintf(contexts interface{}, s string, a ...interface{}) {
	contextsType := reflect.TypeOf(contexts).Kind()

	var contextString string
	if contextsType == reflect.String {
		contextString = contexts.(string)
	} else {
		contextString = strings.Join(contexts.([]string), "][")
	}

	s = fmt.Sprintf(s, a...)
	log.Printf("[%s] %s", contextString, s)
}

// InfoPrintf - Calls contextPrintf with a prepended "INFO" context.
// "contexts" can be either a string or a string slice.
func InfoPrintf(contexts interface{}, s string, a ...interface{}) {
	contextsType := reflect.TypeOf(contexts).Kind()

	if contextsType == reflect.String {
		contexts = append([]string{"INFO"}, contexts.(string))
	} else {
		contexts = append([]string{"INFO"}, contexts.([]string)...)
	}

	ContextPrintf(contexts, s, a...)
}

// ErrorPrintf - Calls contextPrintf with a prepended "ERROR" context.
// "contexts" can be either a string or a string slice.
func ErrorPrintf(contexts interface{}, s string, a ...interface{}) {
	contextsType := reflect.TypeOf(contexts).Kind()

	if contextsType == reflect.String {
		contexts = append([]string{"ERROR"}, contexts.(string))
	} else {
		contexts = append([]string{"ERROR"}, contexts.([]string)...)
	}

	ContextPrintf(contexts, s, a...)
}

// ServicePrintf - Calls infoPrintf with a prepended "SERVICE" context.
func ServicePrintf(context, s string, a ...interface{}) {
	contexts := []string{"SERVICE", context}
	InfoPrintf(contexts, s, a...)
}

// ServiceErrorPrintf - Calls errorPrintf with a prepended "SERVICE" context.
func ServiceErrorPrintf(context, s string, a ...interface{}) {
	contexts := []string{"SERVICE", context}
	ErrorPrintf(contexts, s, a...)
}

// EventPrintf - Calls infoPrintf with a prepended "EVENT" context.
func EventPrintf(context, s string, a ...interface{}) {
	contexts := []string{"EVENT", context}
	InfoPrintf(contexts, s, a...)
}

// EventPrintf - Calls errorPrintf with a prepended "EVENT" context.
func EventErrorPrintf(context, s string, a ...interface{}) {
	contexts := []string{"EVENT", context}
	ErrorPrintf(contexts, s, a...)
}

// ControllerPrintf - Calls infoPrintf with a prepended "CONTROLLER" context.
func ControllerPrintf(context, s string, a ...interface{}) {
	contexts := []string{"CONTROLLER", context}
	InfoPrintf(contexts, s, a...)
}

// ControllerErrorPrintf - Calls errorPrintf with a prepended "CONTROLLER" context.
func ControllerErrorPrintf(context, s string, a ...interface{}) {
	contexts := []string{"CONTROLLER", context}
	ErrorPrintf(contexts, s, a...)
}
