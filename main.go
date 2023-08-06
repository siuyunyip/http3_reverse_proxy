package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/txthinking/runnergroup"
)

func main() {
	port := flag.String("port", "443", "bind to port")
	domain := flag.String("domain", "dev.cafewithbook.org", "domain name")
	www := flag.String("www", "./html", "web root")
	flag.Parse()

	//h, h2, h3, wg := NewServer(*domain, *port, *www)
	h, wg := NewServer(*domain, *port, *www)

	g := runnergroup.New()
	g.Add(&runnergroup.Runner{
		Start: func() error {
			return h.ListenAndServe()
		},
		Stop: func() error {
			return h.Shutdown(context.Background())
		},
	})

	//g.Add(&runnergroup.Runner{
	//	Start: func() error {
	//		return h2.ListenAndServeTLS("", "")
	//	},
	//	Stop: func() error {
	//		return h2.Shutdown(context.Background())
	//	},
	//})
	//g.Add(&runnergroup.Runner{
	//	Start: func() error {
	//		return h3.ListenAndServe()
	//	},
	//	Stop: func() error {
	//		return h3.Close()
	//	},
	//})

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		g.Done()
		wg.Done()
	}()

	// start workers
	wg.Wait()
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

}
