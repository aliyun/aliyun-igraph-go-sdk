package aliyun_igraph_go_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadRequest_BuildUri(t *testing.T) {
	m := make(map[string]string)
	queryString := "g(\"idmapping\").V(\"901f8903d55333fcbf29c530885aeef2\").hasLabel(\"oaid\").out()"
	readRequest := &ReadRequest{QueryString: queryString, QueryParams: m}
	url := readRequest.BuildUri()
	expectedUrl := "ip=127.0.0.1?config%7Boutfmt%3Djson%26no_cache%3Dfalse%26cache_only%3Dfalse%7D%26%26g%28%22idmapping%22%29.V%28%22901f8903d55333fcbf29c530885aeef2%22%29.hasLabel%28%22oaid%22%29.out%28%29"
	assert.Equal(t, url, expectedUrl)
}

func TestWriteRequest_BuildUri(t *testing.T) {
	graphName := "testWrite"
	instanceName := "igraph-cn-testWrite"
	request := NewWriteRequest(WriteTypeAdd, instanceName, graphName, "sku_feature", "id", "", map[string]string{})
	request.AddContent("goods_size_list", "去")
	request.AddContent("go_gender_ratio", "|")
	request.AddContent("go_gender_rank", "1a")
	request.AddContent("id", "1")
	url := request.BuildUri()
	expectedUrl := "update?go_gender_rank=1a&go_gender_ratio=%7C&goods_size_list=%E5%8E%BB&pkey=1&table=testWrite_sku_feature&type=1"
	assert.Equal(t, url.RequestURI(), expectedUrl)
}
