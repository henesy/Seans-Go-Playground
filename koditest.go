package main

import (
kodi	"github.com/pdf/kodirpc"
	"fmt"
)


/* Send notification to kodi @ iseage */
func main() {
	client, err := kodi.NewClient(`10.0.5.201:9090`, kodi.NewConfig())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			panic(err)
		}
	}()

	params := map[string]interface{}{
		`title`:       `Hello`,
		`message`:     `From kodirpc`,
		`displaytime`: 5000,
	}
	res, err := client.Call(`GUI.ShowNotification`, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
