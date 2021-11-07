package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func startServer(server *http.Server) error {
	http.HandleFunc("/hello", helloServer)
	fmt.Println("start server")
	err := server.ListenAndServe()
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	src := &http.Server{Addr: ":9090"}
	group.Go(func() error {
		return startServer(src)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return src.Shutdown(errCtx)
	})
	channel := make(chan os.Signal, 1)
	signal.Notify(channel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-channel:
				cancel()
			}
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Println("group error:", err)
	}
	fmt.Println("all group done")
}
