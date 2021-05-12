package tests

import "testwork/incrementor"

type test struct {
    testingFoo  interface{}
    fooParam    interface{}
    estResult   interface{}
    estError    bool
}

func getTests() []test {
    var tests []test

    tests = append(tests, test{  // 0 test of GetNumber request && started value
        testingFoo: incrementor.Incrementor.GetNumber,
        estResult: 0,
    })

    tests = append(tests, test{ // 1 test of IncrementNumber request
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 2
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 3 test of IncrementNumber result
        testingFoo: incrementor.Incrementor.GetNumber,
        estResult: 2,
    })

    tests = append(tests, test{ // 4 test of SetMaximumValue request error
        testingFoo: incrementor.Incrementor.SetMaximumValue,
        fooParam: -2,
        estError: true,
    })

    tests = append(tests, test{ // 5 test of SetMaximumValue request
        testingFoo: incrementor.Incrementor.SetMaximumValue,
        fooParam: 2,
    })

    tests = append(tests, test{ // 6
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 7 test of SetMaximumValue result
        testingFoo: incrementor.Incrementor.GetNumber,
        estResult: 0,
    })

    tests = append(tests, test{ // 8
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 9 test of IncrementNumber result
        testingFoo: incrementor.Incrementor.GetNumber,
        estResult: 1,
    })

    tests = append(tests, test{ // 10
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 11 test of IncrementNumber maxValue result
        testingFoo: incrementor.Incrementor.IncrementNumber,
    })

    tests = append(tests, test{ // 12  test of IncrementNumber maxValue result
        testingFoo: incrementor.Incrementor.GetNumber,
        estResult: 0,
    })

    return tests
}
