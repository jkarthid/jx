// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	client "github.com/heptio/sonobuoy/pkg/client"
)

func AnyPtrToClientSonobuoyClient() *client.SonobuoyClient {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*client.SonobuoyClient))(nil)).Elem()))
	var nullValue *client.SonobuoyClient
	return nullValue
}

func EqPtrToClientSonobuoyClient(value *client.SonobuoyClient) *client.SonobuoyClient {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *client.SonobuoyClient
	return nullValue
}
