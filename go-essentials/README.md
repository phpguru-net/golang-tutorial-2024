# Install Go

> For Apple Silicon

```sh
sudo rm -rf /usr/local/go
brew uninstall go
brew install wget
wget https://golang.org/dl/go1.22.3.darwin-arm64.tar.gz
sudo tar -C /usr/local -xzf go1.22.3.darwin-arm64.tar.gz
source ~/.zshrc  # or source ~/.bash_profile
```

> go version go1.22.3 darwin/arm64
