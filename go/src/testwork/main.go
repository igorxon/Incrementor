package main

import (
    "fmt"
    "testwork/incrementor"
)

func main() {
    inc, defMaxValue := incrementor.Init()      // class instance initialise

    fmt.Printf("Incrementor initialized, default max value: %d\n", defMaxValue)

    for n := 0; n < defMaxValue + 1;  n++ {     // example of class usage
        curValue := inc.GetNumber()
        fmt.Printf("Current value: %d\n", curValue)
        inc.IncrementNumber()
    }
}

