login:
	gcloud auth login
install:
	go get github.com/rakyll/statik@latest

resource:
	statik -src=static

test:
	go test ./...

qrme:
	gcloud config set project qrme

deploy: qrme vendoring
	GO111MODULE=off gcloud --quiet app deploy app.yaml --project qrme-286611

vendoring:
	go mod vendor
