package main

//import (
//	"net"
//	"os"
//	"time"
//)
//
//// $GOROOT/src/net/http/server.go
//func (srv *Server) Serve(l net.Listener) error {
//	... ...
//	for {
//		rw, e := l.Accept()
//		if e != nil {
//			select {
//			case <-srv.getDoneChan():
//				return ErrServerClosed
//			default:
//			}
//			if ne, ok := e.(net.Error); ok && ne.Temporary() {
//				// 注：这里对临时性(temporary)错误进行处理
//				... ...
//				time.Sleep(tempDelay)
//				continue
//			}
//			return e
//		}
//		...
//	}
//	... ...
//}
//
//// $GOROOT/src/net/net.go
//type OpError struct {
//	... ...
//	// Err is the error that occurred during the operation.
//	Err error
//}
//
//type temporary interface {
//	Temporary() bool
//}
//
//func (e *OpError) Temporary() bool {
//	if ne, ok := e.Err.(*os.SyscallError); ok {
//		t, ok := ne.Err.(temporary)
//		return ok && t.Temporary()
//	}
//	t, ok := e.Err.(temporary)
//	return ok && t.Temporary()
//}
