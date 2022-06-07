package main

import (
    "testing"
)

func TestContext(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    
    NewHandlerContext(nil)
}
