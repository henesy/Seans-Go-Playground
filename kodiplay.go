package main

import (
	"github.com/pdf/kodirpc"
	"fmt"
)

func main() {
	client, err := kodirpc.NewClient(`10.0.5.201:9090`, kodirpc.NewConfig())
	if err != nil {
	    panic(err)
	}
	defer func() {
	    if err = client.Close(); err != nil {
	        panic(err)
	    }
	}()
	
	
	var params = make(map[string]map[string]interface{})
	params["item"] = make(map[string]interface{})
	
	params["item"]["file"] = `plugin://plugin.video.youtube/play/?video_id=K3Qzzggn--s`

	res, err := client.Call(`Player.Open`, params)
	if err != nil {
	    panic(err)
	}
	
	fmt.Println(res)
}
