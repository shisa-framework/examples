// generated by "charlatan -output=./httphandler_charlatan.go http.Handler".  DO NOT EDIT.

package httpx

import "net/http"

import "reflect"

// HandlerServeHTTPInvocation represents a single call of FakeHandler.ServeHTTP
type HandlerServeHTTPInvocation struct {
	Parameters struct {
		Ident25 http.ResponseWriter
		Ident26 *http.Request
	}
}

// HandlerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type HandlerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeHandler is a mock implementation of Handler for testing.
Use it in your tests as in this example:

	package example

	func TestWithHandler(t *testing.T) {
		f := &httpx.FakeHandler{
			ServeHTTPHook: func(ident25 http.ResponseWriter, ident26 *http.Request) () {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeServeHTTP ...
		f.AssertServeHTTPCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeServeHTTP.
*/
type FakeHandler struct {
	ServeHTTPHook func(http.ResponseWriter, *http.Request)

	ServeHTTPCalls []*HandlerServeHTTPInvocation
}

// NewFakeHandlerDefaultPanic returns an instance of FakeHandler with all hooks configured to panic
func NewFakeHandlerDefaultPanic() *FakeHandler {
	return &FakeHandler{
		ServeHTTPHook: func(http.ResponseWriter, *http.Request) {
			panic("Unexpected call to Handler.ServeHTTP")
		},
	}
}

// NewFakeHandlerDefaultFatal returns an instance of FakeHandler with all hooks configured to call t.Fatal
func NewFakeHandlerDefaultFatal(t HandlerTestingT) *FakeHandler {
	return &FakeHandler{
		ServeHTTPHook: func(http.ResponseWriter, *http.Request) {
			t.Fatal("Unexpected call to Handler.ServeHTTP")
			return
		},
	}
}

// NewFakeHandlerDefaultError returns an instance of FakeHandler with all hooks configured to call t.Error
func NewFakeHandlerDefaultError(t HandlerTestingT) *FakeHandler {
	return &FakeHandler{
		ServeHTTPHook: func(http.ResponseWriter, *http.Request) {
			t.Error("Unexpected call to Handler.ServeHTTP")
			return
		},
	}
}

func (f *FakeHandler) Reset() {
	f.ServeHTTPCalls = []*HandlerServeHTTPInvocation{}
}

func (_f1 *FakeHandler) ServeHTTP(ident25 http.ResponseWriter, ident26 *http.Request) {
	if _f1.ServeHTTPHook == nil {
		panic("Handler.ServeHTTP() called but FakeHandler.ServeHTTPHook is nil")
	}

	invocation := new(HandlerServeHTTPInvocation)
	_f1.ServeHTTPCalls = append(_f1.ServeHTTPCalls, invocation)

	invocation.Parameters.Ident25 = ident25
	invocation.Parameters.Ident26 = ident26

	_f1.ServeHTTPHook(ident25, ident26)

	return
}

// ServeHTTPCalled returns true if FakeHandler.ServeHTTP was called
func (f *FakeHandler) ServeHTTPCalled() bool {
	return len(f.ServeHTTPCalls) != 0
}

// AssertServeHTTPCalled calls t.Error if FakeHandler.ServeHTTP was not called
func (f *FakeHandler) AssertServeHTTPCalled(t HandlerTestingT) {
	t.Helper()
	if len(f.ServeHTTPCalls) == 0 {
		t.Error("FakeHandler.ServeHTTP not called, expected at least one")
	}
}

// ServeHTTPNotCalled returns true if FakeHandler.ServeHTTP was not called
func (f *FakeHandler) ServeHTTPNotCalled() bool {
	return len(f.ServeHTTPCalls) == 0
}

// AssertServeHTTPNotCalled calls t.Error if FakeHandler.ServeHTTP was called
func (f *FakeHandler) AssertServeHTTPNotCalled(t HandlerTestingT) {
	t.Helper()
	if len(f.ServeHTTPCalls) != 0 {
		t.Error("FakeHandler.ServeHTTP called, expected none")
	}
}

// ServeHTTPCalledOnce returns true if FakeHandler.ServeHTTP was called exactly once
func (f *FakeHandler) ServeHTTPCalledOnce() bool {
	return len(f.ServeHTTPCalls) == 1
}

// AssertServeHTTPCalledOnce calls t.Error if FakeHandler.ServeHTTP was not called exactly once
func (f *FakeHandler) AssertServeHTTPCalledOnce(t HandlerTestingT) {
	t.Helper()
	if len(f.ServeHTTPCalls) != 1 {
		t.Errorf("FakeHandler.ServeHTTP called %d times, expected 1", len(f.ServeHTTPCalls))
	}
}

// ServeHTTPCalledN returns true if FakeHandler.ServeHTTP was called at least n times
func (f *FakeHandler) ServeHTTPCalledN(n int) bool {
	return len(f.ServeHTTPCalls) >= n
}

// AssertServeHTTPCalledN calls t.Error if FakeHandler.ServeHTTP was called less than n times
func (f *FakeHandler) AssertServeHTTPCalledN(t HandlerTestingT, n int) {
	t.Helper()
	if len(f.ServeHTTPCalls) < n {
		t.Errorf("FakeHandler.ServeHTTP called %d times, expected >= %d", len(f.ServeHTTPCalls), n)
	}
}

// ServeHTTPCalledWith returns true if FakeHandler.ServeHTTP was called with the given values
func (_f2 *FakeHandler) ServeHTTPCalledWith(ident25 http.ResponseWriter, ident26 *http.Request) (found bool) {
	for _, call := range _f2.ServeHTTPCalls {
		if reflect.DeepEqual(call.Parameters.Ident25, ident25) && reflect.DeepEqual(call.Parameters.Ident26, ident26) {
			found = true
			break
		}
	}

	return
}

// AssertServeHTTPCalledWith calls t.Error if FakeHandler.ServeHTTP was not called with the given values
func (_f3 *FakeHandler) AssertServeHTTPCalledWith(t HandlerTestingT, ident25 http.ResponseWriter, ident26 *http.Request) {
	t.Helper()
	var found bool
	for _, call := range _f3.ServeHTTPCalls {
		if reflect.DeepEqual(call.Parameters.Ident25, ident25) && reflect.DeepEqual(call.Parameters.Ident26, ident26) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeHandler.ServeHTTP not called with expected parameters")
	}
}

// ServeHTTPCalledOnceWith returns true if FakeHandler.ServeHTTP was called exactly once with the given values
func (_f4 *FakeHandler) ServeHTTPCalledOnceWith(ident25 http.ResponseWriter, ident26 *http.Request) bool {
	var count int
	for _, call := range _f4.ServeHTTPCalls {
		if reflect.DeepEqual(call.Parameters.Ident25, ident25) && reflect.DeepEqual(call.Parameters.Ident26, ident26) {
			count++
		}
	}

	return count == 1
}

// AssertServeHTTPCalledOnceWith calls t.Error if FakeHandler.ServeHTTP was not called exactly once with the given values
func (_f5 *FakeHandler) AssertServeHTTPCalledOnceWith(t HandlerTestingT, ident25 http.ResponseWriter, ident26 *http.Request) {
	t.Helper()
	var count int
	for _, call := range _f5.ServeHTTPCalls {
		if reflect.DeepEqual(call.Parameters.Ident25, ident25) && reflect.DeepEqual(call.Parameters.Ident26, ident26) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeHandler.ServeHTTP called %d times with expected parameters, expected one", count)
	}
}
