# Snippetbox

Web application written in go, from the fantastic book [Let's Go, by Alex Edwards](https://lets-go.alexedwards.net/).

## Development

I use [`air`](https://github.com/air-verse/air) to live reload the app:

```sh
go install github.com/air-verse/air@latest
```

To run `air` from the CLI, add the `$GOPATH/bin` to the `$PATH`. On zsh, edit the file`~/.zshrc` and tweak the `PATH` variable like so:
```
export PATH=$HOME/go/bin:$PATH
```

Run `air` to start the app:
```sh
air
```
<span style="background-color=rgba(245, 215, 39, 0.28)">💡 The app itself runs on port 8080, but `air` proxy it to [http://localhost:8008](http://localhost:8008) allowing to reload the browser after a new build.</span>
