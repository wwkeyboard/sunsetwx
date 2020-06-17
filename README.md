# sunsetwx
A go client to interact with [https://sunsetwx.com/](sunsetwx.com)'s API

# Prereqs

An account, go to [https://subscriptions.sunsetwx.com/](here) and click `sign up`.

`go get github.com/wwkeyboard/sunsetwx`

# Usage

There is an [https://github.com/wwkeyboard/sunsetwx/blob/master/cmd/query/main.go](example program) in the cmd directory. The steps are:

1. instantiate a new client
2. Login, which exchanges the username/password for an auth token that's stored in the client struct
3. call `GetQuality` which returns a struct that matches [https://sunburst.sunsetwx.com/v1/docs/#get-quality](the JSON) returned from the API.

```golang
	client := sunsetwx.NewClient()
	err := client.Login(config.username, config.password)
	if err != nil {
		return err
	}

	quality, err := client.GetQuality(40.1108411, -88.2075309)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", quality)
  ```
