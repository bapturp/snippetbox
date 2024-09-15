# Snippetbox

Web application written in go, from the fantastic book [Let's Go, by Alex Edwards](https://lets-go.alexedwards.net/).

## Development

I setup [air](https://github.com/air-verse/air) to live reload the app:

Install air:
```sh
go install github.com/air-verse/air@latest
```

To run `air` from the CLI, add the `$GOPATH/bin` to the `$PATH`:

On zsh, edit `~/.zshrc`
```
export PATH=$HOME/go/bin:$PATH
```

Run `air` to start the app:
```sh
air
```

💡 The app itself runs on port 8080, but `air` proxy it to [http://localhost:8008](http://localhost:8008) allowing to reload the browser after a new build.
