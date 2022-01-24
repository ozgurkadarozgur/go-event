package client

type Client interface {
	Publish(channel string, message []byte)
	Subscribe(channel string, callback func(payload string))
}
