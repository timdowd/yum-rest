// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// based on the gorilla logging handler found at:
//       https://github.com/gorilla/handlers/blob/master/handlers.go

package httpmw

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"time"

	"github.com/pizzahutdigital/phdmw/phdcid"
	"github.com/pizzahutdigital/phdmw/phdlog"
)

type key int

const (
	// HandlerNameKey is the key for context values sent from handler to the logger.  This is to prevent package conflicts.
	HandlerNameKey key = iota
)

// HandlerStart - Used at begining of handler to add handler name in the context so it can be logged in the HTTPHandlerSummary Event log
func HandlerStart(r *http.Request, handlerName string) string {
	// Prevent panics during tests.
	if r == nil {
		return ""
	}

	iPtr := r.Context().Value(HandlerNameKey)
	if iPtr != nil {
		ptr := iPtr.(*string)
		*ptr = handlerName
	}
	return phdcid.GetCIDFromContext(r.Context())
}

// loggingHandler is the http.Handler implementation for LoggingHandlerTo and its
// friends
type loggingHandler struct {
	handler http.Handler
}

// LoggingMiddleware - logging interceptor
func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return loggingHandler{h}
	}
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	logger := makeLogger(w)

	// Create a variable that we can assign a value to in the handlers using the pointer.
	// I am not proud of this, but it solved a specific problem.
	handlerName := "unknown"

	// Create a new context and add our variable pointer to the context for use in the handlers.
	nReq := req.WithContext(context.WithValue(req.Context(), HandlerNameKey, &handlerName))

	h.handler.ServeHTTP(logger, nReq)

	// calculate number of milliseconds the request took to process
	duration := ((time.Since(t).Nanoseconds()) / 1000000)
	phdlog.HTTPHandlerSummary(req,
		logger.Size(),
		logger.Status(),
		duration,
		phdcid.GetCIDFromContext(req.Context()),
		handlerName)
}

func makeLogger(w http.ResponseWriter) loggingResponseWriter {
	var logger loggingResponseWriter = &responseLogger{w: w}
	if _, ok := w.(http.Hijacker); ok {
		logger = &hijackLogger{responseLogger{w: w}}
	}
	h, ok1 := logger.(http.Hijacker)
	c, ok2 := w.(http.CloseNotifier)
	if ok1 && ok2 {
		return hijackCloseNotifier{logger, h, c}
	}
	if ok2 {
		return &closeNotifyWriter{logger, c}
	}
	return logger
}

type loggingResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	Status() int
	Size() int
}

// responseLogger is wrapper of http.ResponseWriter that keeps track of its HTTP
// status code and body size
type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *responseLogger) Write(b []byte) (int, error) {
	if l.status == 0 {
		// The status will be StatusOK if WriteHeader has not been called yet
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *responseLogger) Status() int {
	return l.status
}

func (l *responseLogger) Size() int {
	return l.size
}

func (l *responseLogger) Flush() {
	f, ok := l.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}

type hijackLogger struct {
	responseLogger
}

func (l *hijackLogger) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h := l.responseLogger.w.(http.Hijacker)
	conn, rw, err := h.Hijack()
	if err == nil && l.responseLogger.status == 0 {
		// The status will be StatusSwitchingProtocols if there was no error and
		// WriteHeader has not been called yet
		l.responseLogger.status = http.StatusSwitchingProtocols
	}
	return conn, rw, err
}

type closeNotifyWriter struct {
	loggingResponseWriter
	http.CloseNotifier
}

type hijackCloseNotifier struct {
	loggingResponseWriter
	http.Hijacker
	http.CloseNotifier
}
