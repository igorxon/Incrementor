package tests

import (
    "fmt"
    "os"
    "reflect"
    "testing"
    "testwork/incrementor"
)

const errMsg = "%d: Test error: %v\n"
const errResult = "%d: Estimated result: %v, but: %v\n"

func TestMain(m *testing.M) {
    fmt.Println("Start testing of 'incrementor' class")

    rez := m.Run()

    fmt.Println("All tests passed")
    fmt.Println("Complete testing of 'incrementor' class")

    os.Exit(rez)
}

func TestGrpc(t *testing.T) {
    inc, _ := incrementor.Init()

    for	i, tst := range getTests() {                                    // iterate throw all tests

        params := []reflect.Value{ reflect.ValueOf(inc) }               // address of class instance

        if tst.fooParam != nil {
            params = append(params, reflect.ValueOf(tst.fooParam))      // optional parameter for class methods
        }

        hr := reflect.ValueOf(tst.testingFoo).Call(params)              // call of class method

        if len(hr) == 0 {                                               // there is return value ?
            if !tst.estError { continue }

            t.Error(fmt.Sprintf(errResult, i, tst.estResult, hr))       // must return string when tst.estError is true!
            break
        }

        rez := hr[0].Interface()                                        // fetch result value

        if reflect.TypeOf(rez) == reflect.TypeOf(fmt.Errorf("")) {
            if tst.estError { continue }                                // if result type is string && waiting for error
                                                                        // - ok
            t.Error(fmt.Sprintf(errMsg, i, rez))
            break
        }

        if rez != tst.estResult {                                       // else result type must be integer
            t.Errorf( fmt.Sprintf(errResult, i, tst.estResult, rez))    // and equal to tst.estimate result
            break
        }
    }
}
