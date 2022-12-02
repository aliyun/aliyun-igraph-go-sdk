package aliyun_igraph_go_sdk

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

type ReadRequest struct {
	QueryString string            `json:"query_string"`
	QueryParams map[string]string `json:"query_params"`
}

func (r *ReadRequest) Validate() error {
	if len(r.QueryString) == 0 {
		return InvalidParamsError{"Empty query string"}
	}
	return nil
}

func (r *ReadRequest) AddQueryParam(key string, value string) *ReadRequest {
	r.QueryParams[key] = value
	return r
}

func (r *ReadRequest) SetQueryParams(params map[string]string) *ReadRequest {
	r.QueryParams = params
	return r
}

func (r *ReadRequest) BuildUri() url.URL {
	uri := url.URL{Path: "app"}
	query := map[string]string{}
	query["app"] = "gremlin"
	if len(r.QueryParams) != 0 {
		for k, v := range r.QueryParams {
			query[k] = v
		}
	}

	IP, err := LocalIP()
	if err != nil {
		fmt.Println(err)
		query["ip"] = "127.0.0.1"
	} else {
		query["ip"] = IP.String()
	}
	var configStr = "config{outfmt=json&cache_only=false&no_cache=false}"
	query["ip"] = query["ip"] + "?" + configStr + "&&" + r.QueryString

	var params []string
	for k, v := range query {
		params = append(params, k+"="+v)
	}
	uri.RawQuery = strings.Join(params[:], "&")
	return uri
}

func LocalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			return ip, nil
		}
	}

	return nil, errors.New("no IP")
}
