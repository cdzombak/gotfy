# gotfy

GoTFY is an API client for sending notifications using [ntfy](https://ntfy.sh) servers.

## Installation

```shell
go get github.com/cdzombak/gotfy
```

## Example Usage

```go
const ntfyTimeout = 10 * time.Second
serverURL, _ := url.Parse("https://ntfy.example.com")

publisher, err := gotfy.NewTopicPublisher(nil, serverURL, nil)
if err != nil {
    panic("bad config: " + err.Error())
}

ntfyPublisher.Headers.Set("user-agent", "my-app / 1.0")
ntfyPublisher.Headers.Set("authorization", "Bearer tk_MY_NTFY_TOKEN")

ctx, cancel := context.WithTimeout(context.Background(), ntfyTimeout)
defer cancel()
publisher.SendMessage(ctx, &gotfy.Message{
    Topic:    "topic",
    Message:  "message",
    Title:    "title",
    Tags:     []string{"emoji1","emoji2","some text"},
    Priority: gotfy.High,
    Actions:  []gotfy.ActionButton{
	    Label: "label",
	    Link: "http://link.example.com",
	    Clear: true,
    },
    ClickURL: "http://click.example.com",
    IconURL:  "http://icon.example.com",
    Delay:    time.Minute * 5,
    Email:    "me@example.com",
})
```

## License

gotfy is licensed under the Apache 2.0 license. See LICENSE in this repository.

- Original library copyright 2023 [AnthonyHewins](https://github.com/AnthonyHewins)
- Fork changes copyright 2023 [Chris Dzombak](https://www.dzombak.com)
