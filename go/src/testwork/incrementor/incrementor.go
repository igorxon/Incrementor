package incrementor

import (
    "fmt"
    "sync"
)

type Incrementor interface {
    GetNumber() int                     // get current counter value
    IncrementNumber()                   // increment counter value
    SetMaximumValue(int) error          // set maximal counter value, return error if value <= 0
}

type incrementor struct {
    maxValue    int
    currValue   int
    mutex       sync.Mutex
}

const defMaxValue = 5                   // default Maximal value for counter

func Init() (Incrementor, int) {        // class instance initialise
    inc := incrementor{
        currValue: 0,                   // set initial counter value
        maxValue: defMaxValue,          // set initial Maximal counter value
    }
    return &inc, defMaxValue            // return address of class instance
}

func (c *incrementor) GetNumber() int {
    c.mutex.Lock()
    retValue := c.currValue             // get current value to return
    c.mutex.Unlock()
    return retValue
}

func (c *incrementor) IncrementNumber() {
    c.mutex.Lock()
    c.currValue += 1                    // increment current value
    if c.currValue > c.maxValue {
        c.currValue = 0
    }
    c.mutex.Unlock()
}

func (c *incrementor) SetMaximumValue(maxValue int) error {
    if maxValue <= 0 {
        return fmt.Errorf("maxValue must be > 0, but: %d", maxValue)
    }
    c.mutex.Lock()
    c.maxValue = maxValue
    if c.currValue > c.maxValue {
        c.currValue = 0                 // set current value to 0 if it was > maximal value
    }
    c.mutex.Unlock()
    return nil
}
