package main

import (
	"context"
	"testing"
	"connectrpc.com/connect"
	calculator "calculatorRPC/backend/gen/calculator"
)

func makeRequest(a, b float64, op string) *connect.Request[calculator.CalculateRequest] {
	return connect.NewRequest(&calculator.CalculateRequest{A: a, B: b, Operation: op})
}

func TestAdd(t *testing.T) {
	res, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 5, "add"))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.Result != 15 {
		t.Errorf("expected 15, got %f", res.Msg.Result)
	}
}

func TestSubtract(t *testing.T) {
	res, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 5, "subtract"))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.Result != 5 {
		t.Errorf("expected 5, got %f", res.Msg.Result)
	}
}

func TestMultiply(t *testing.T) {
	res, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 5, "multiply"))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.Result != 50 {
		t.Errorf("expected 50, got %f", res.Msg.Result)
	}
}

func TestDivide(t *testing.T) {
	res, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 5, "divide"))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.Result != 2 {
		t.Errorf("expected 2, got %f", res.Msg.Result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 0, "divide"))
	if err == nil {
		t.Error("expected an error for divide by zero, got error")
	}
}

func TestUnknownOperation(t *testing.T) {
	_, err := (&CalculatorServer{}).Calculate(context.Background(), makeRequest(10, 5, "mod"))
	if err == nil {
		t.Error("expected an error for unknown operation, got error")
	}
}
