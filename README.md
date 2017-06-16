# chat-go

## Docker Way

``` docker build -t . ```

``` docker run -ti --name chat -v $(pwd):/go/src/chat-go -p 9090:9090 chat-go  ```