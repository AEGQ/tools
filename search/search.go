package search

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/astaxie/beego/httplib"
)

var (
	//REGISTRYURL (docker info |grep "Registry:")
	REGISTRYURL = "https://index.docker.io/v1/"
	//PROXY read proxy from environment
	PROXY = getProxyEnv("http_proxy")
)

func getProxyEnv(key string) string {
	proxyValue := os.Getenv(strings.ToUpper(key))
	if proxyValue == "" {
		return os.Getenv(strings.ToLower(key))
	}
	return proxyValue
}

//Images docker images on hub.docker.com
//eg: https://index.docker.io/v1/search?q=ubuntu&n=5
func Images(image string, limit int) (*ImageResults, error) {
	u := REGISTRYURL + "search?q=" + image + "&n=" + strconv.Itoa(limit)
	req := httplib.NewBeegoRequest(u, "GET")
	proxy := func(req *http.Request) (*url.URL, error) {
		u, _ := url.ParseRequestURI(PROXY)
		return u, nil
	}
	req.SetProxy(proxy)
	results := new(ImageResults)
	if err := req.ToJSON(results); err != nil {
		return nil, fmt.Errorf("ERROR Search->ToJSON: %s", err.Error())
	}
	return results, nil
}

//Tags list all tags
//eg: https://index.docker.io/v1/repositories/ubuntu/tags
func Tags(image string) (*[]TagResult, error) {
	u := REGISTRYURL + "repositories/" + image + "/tags"
	req := httplib.NewBeegoRequest(u, "GET")
	proxy := func(req *http.Request) (*url.URL, error) {
		u, _ := url.ParseRequestURI(PROXY)
		return u, nil
	}
	req.SetProxy(proxy)
	tags := new([]TagResult)
	if err := req.ToJSON(tags); err != nil {
		return nil, fmt.Errorf("ERROR Tags->ToJSON: %s", err.Error())
	}
	return tags, nil
}

//PrintImages  STAR  OFFICIAL  NAME  URL
func PrintImages(results *ImageResults) {
	images := ImageResultsByStars(results.Results)
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)
	fmt.Fprintf(w, "STAR\tOFFICIAL\tNAME\tURL\n")
	for _, v := range images {
		url := "https://hub.docker.com/r/" + v.Name
		if v.IsOfficial {
			url = "https://hub.docker.com/r/library/" + v.Name
		}
		fmt.Fprintf(w, "%d\t", v.StarCount)
		if v.IsOfficial {
			fmt.Fprint(w, "[OK]")
		}
		fmt.Fprint(w, "\t")
		fmt.Fprintf(w, "%s\t%s\n", v.Name, url)
	}
	w.Flush()
}
