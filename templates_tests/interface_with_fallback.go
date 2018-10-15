package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../templates/fallback template

//go:generate gowrap gen -d . -i TestInterface -t ../templates/fallback -o interface_with_fallback.go

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// TestInterfaceWithFallback implements TestInterface interface wrapped with Prometheus metrics
type TestInterfaceWithFallback struct {
	implementations []TestInterface
	interval        time.Duration
}

// NewTestInterfaceWithFallback takes several implementations of the TestInterface and returns an instance of TestInterface
// which calls all implementations concurrently with given interval and returns first non-error response.
func NewTestInterfaceWithFallback(interval time.Duration, impls ...TestInterface) TestInterfaceWithFallback {
	return TestInterfaceWithFallback{implementations: impls, interval: interval}
}

// F implements TestInterface
func (_d TestInterfaceWithFallback) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	type _resultStruct struct {
		result1 string
		result2 string
		err     error
	}
	var _res _resultStruct
	var _ch = make(chan _resultStruct, 0)
	var _errorsList []string
	var _ticker = time.NewTicker(_d.interval)
	defer _ticker.Stop()
	ctx, _cancelFunc := context.WithCancel(ctx)
	defer _cancelFunc()

	for _i := 0; _i < len(_d.implementations); _i++ {
		go func(_impl TestInterface) {
			result1, result2, err := _impl.F(ctx, a1, a2...)
			if err != nil {
				err = fmt.Errorf("%T: %v", _impl, err)
			}
			_ch <- _resultStruct{result1, result2, err}
		}(_d.implementations[_i])
		select {
		case _res = <-_ch:
			if _res.err == nil {
				return _res.result1, _res.result2, _res.err
			}
			_errorsList = append(_errorsList, _res.err.Error())
		case <-_ticker.C:
		}
	}
	err = fmt.Errorf(strings.Join(_errorsList, ";"))
	return
}
