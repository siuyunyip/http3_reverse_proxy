package main

import (
	"fmt"
	"github.com/txthinking/runnergroup"
	"net/http"
	"net/url"
)

var backends = []string{
	"http://172.17.0.2:8001",
	"http://172.17.0.3:8002",
	"http://172.17.0.4:8003",
	"http://172.17.0.5:8004",
}

type reqRsp struct {
	req *http.Request
	rsp http.ResponseWriter
}

type Worker struct {
	channel chan *reqRsp
	handler http.Handler
	state   int
	id      int
}

type WorkerPool struct {
	Workers map[string]Worker
}

func InitWorker() (runnergroup.RunnerGroup, WorkerPool) {
	pool := &WorkerPool{Workers: make(map[string]Worker)}
	g := runnergroup.New()

	var id = 1
	for _, u := range backends {
		w, err := newWorker(u, id)
		if err != nil {
			fmt.Println("Fail to create new worker with url: %s", u)
			continue
		}
		pool.Workers[u] = *w
		g.Add(&runnergroup.Runner{
			Start: func() error {
				return w.Start()
			},
			Stop: func() error {
				return w.Stop()
			},
		})
		id += 1
	}

	return *g, *pool
}

func newWorker(to string, id int) (*Worker, error) {
	u, err := url.Parse(to)
	if err != nil {
		fmt.Println("Incorrect upstream url")
		return nil, err
	}

	return &Worker{
		channel: make(chan *reqRsp, 100),
		handler: NewSingleHostReverseProxy(u),
		state:   0,
		id:      id,
	}, nil
}

func (w Worker) BuffReq(req *http.Request, rsp http.ResponseWriter) {
	w.channel <- &reqRsp{
		req: req,
		rsp: rsp,
	}
}

func (w Worker) Start() error {
	go func() {
		for {
			if w.state == 1 {
				break
			}
			reqrsp := <-w.channel
			fmt.Println("Get request: ", reqrsp.req.URL)
			fmt.Println("Remote Addr: ", reqrsp.req.RemoteAddr)
			w.handler.ServeHTTP(reqrsp.rsp, reqrsp.req)
		}
	}()

	return nil
}

func (w *Worker) Stop() error {
	w.state = 1
	return nil
}
