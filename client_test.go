package aliyun_igraph_go_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_BuildReadUri(t *testing.T) {
	var client = NewClient("endpoint", "userName", "password", "ut")
	m := make(map[string]string)
	queryString := "g(\"idmapping\").V(\"901f8903d55333fcbf29c530885aeef2\").hasLabel(\"oaid\").out()"
	readRequest := &ReadRequest{QueryString: queryString, QueryParams: m}
	url := client.buildReadUrl(readRequest)
	expectedUrl := "app?app=gremlin&src=ut&ip=127.0.0.1?config%7Boutfmt%3Djson%26no_cache%3Dfalse%26cache_only%3Dfalse%7D%26%26g%28%22idmapping%22%29.V%28%22901f8903d55333fcbf29c530885aeef2%22%29.hasLabel%28%22oaid%22%29.out%28%29"
	assert.Equal(t, url.RequestURI(), expectedUrl)
}

func TestClient_BuildReadUriWithEmptySrc(t *testing.T) {
	var client = NewClient("endpoint", "userName", "password", "")
	m := make(map[string]string)
	queryString := "g(\"idmapping\").V(\"901f8903d55333fcbf29c530885aeef2\").hasLabel(\"oaid\").out()"
	readRequest := &ReadRequest{QueryString: queryString, QueryParams: m}
	url := client.buildReadUrl(readRequest)
	expectedUrl := "app?app=gremlin&src=userName_endpoint&ip=127.0.0.1?config%7Boutfmt%3Djson%26no_cache%3Dfalse%26cache_only%3Dfalse%7D%26%26g%28%22idmapping%22%29.V%28%22901f8903d55333fcbf29c530885aeef2%22%29.hasLabel%28%22oaid%22%29.out%28%29"
	assert.Equal(t, url.RequestURI(), expectedUrl)
}
