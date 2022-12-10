package pusher

import "github.com/pusher/pusher-http-go/v5"

func Authenticate() pusher.Client {
	return pusher.Client{
		AppID:   "1497589",
		Key:     "b54d3ce1a7f22d5d694c",
		Secret:  "849193a3f7ac73b61969",
		Cluster: "ap1",
	}
}
