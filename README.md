# apitool
## Building
to build you need to have a recent version of Go installed, go into apitool`s directory and type:

`go build -o apitool github.com/qredo/apitool/cmd/cli`

this will produce a binary called `apitool`

## CLI
the following command will sign and print the required headers to make a request:

```
./apitool -api-key _lJCqc-T3cnNpw \
          -secret JDJhJDA0JGV0aGt1LnE1a0NyZUFhaHFYVnlzZnVpU0Ryd2xOVThIZTY0TkhpRzRaLjF6YS9ZSi95SExL \
          -method GET \
          -url http://127.0.0.1:8002/qapi/v1/balance sign 
```

output:

```
qredo-api-sign: ZPjvuaQHtHVKg26nd5FMNJKRyA45d5VzwonbZkRxt9w
qredo-api-key: _lJCqc-T3cnNpw 
qredo-api-ts: 1644848012
```

and this will automatically sign and send the request and print the response from the server:

```
./apitool -api-key _lJCqc-T3cnNpw \
          -secret JDJhJDA0JGV0aGt1LnE1a0NyZUFhaHFYVnlzZnVpU0Ryd2xOVThIZTY0TkhpRzRaLjF6YS9ZSi95SExL \
          -method GET \
          -url http://127.0.0.1:8002/qapi/v1/balance send 
```

output:

```
{"assets":{"BTC-TESTNET":{"total":0,"available":0,"pendingIn":0,"pendingOut":0,"scale":100000000},"QCOIN-1":{"total":200000,"available":200000,"pendingIn":0,"pendingOut":0,"scale":1},"QCOIN-2":{"total":100000,"available":100000,"pendingIn":0,"pendingOut":0,"scale":1}}}
```

## Web UI
`apitool` has a simplistic web ui for generating signatures and also sending signed requests, to start it type:

```
./apitool ui
WebUI listening on http://127.0.0.1:4569
```

and click on the link or copy&paste it in your browser

# Final notes

This project is released under the terms of the Apache 2.0 License - see LICENSE for details.
The copyright owner are listed in the `.reuse/dep5` file. Feel free to send Pull Requests (see CONTRIBUTING.md for instructions)
