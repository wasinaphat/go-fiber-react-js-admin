docker run -it  \
    -w "/go/src/github.com/wasinaphatlilawatthananan/go-admin" \
    -v $(pwd):/go/src/github.com/wasinaphatlilawatthananan/go-admin \
    -p 9090:9090 \
    cosmtrek/air