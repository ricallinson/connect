package stackr

import (
	"strconv"
	"time"
)

/*
   ResponseTime:

   Adds the X-Response-Time header displaying the response duration in milliseconds.

   Example:

       stackr.CreateServer().Use(stackr.ResponseTime())

*/
func ResponseTime() func(req *Request, res *Response, next func()) {
	return func(req *Request, res *Response, next func()) {
		start := time.Now().UnixNano()
		res.On("header", func() {
			time.Sleep(100 * time.Millisecond)
			duration := time.Now().UnixNano() - start
			res.SetHeader("X-Response-Time", strconv.FormatInt(int64(time.Duration(duration)/time.Millisecond), 10)+"ms")
		})
	}
}
