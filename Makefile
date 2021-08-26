.PHONY: update-pkg-cache
update-pkg-cache:
	GOPROXY=https://proxy.golang.org go get github.com/ishihaya/cloudlog@v${VERSION}