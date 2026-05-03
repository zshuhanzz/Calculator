package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	calculator "calculatorRPC/backend/gen/calculator"
	"calculatorRPC/backend/gen/calculator/calculatorconnect"
)

// CalculatorServer implements the generated CalculatorServiceHandler interface
type CalculatorServer struct{}

func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[calculator.CalculateRequest],
) (*connect.Response[calculator.CalculateResponse], error) {
	a := req.Msg.A
	b := req.Msg.B
	op := req.Msg.Operation

	var result float64
	var symbol string

	switch op {
	case "add":
		result = a + b
		symbol = "+"
	case "subtract":
		result = a - b
		symbol = "-"
	case "multiply":
		result = a * b
		symbol = "×"
	case "divide":
		if b == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("cannot divide by zero"))
		}
		result = a / b
		symbol = "÷"
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unknown operation: %s", op))
	}

	expression := fmt.Sprintf("%g %s %g = %g", a, symbol, b, result)

	return connect.NewResponse(&calculator.CalculateResponse{
		Result:     result,
		Expression: expression,
	}), nil
}

func main() {
	mux := http.NewServeMux()
	path, handler := calculatorconnect.NewCalculatorServiceHandler(&CalculatorServer{})
	mux.Handle(path, handler)

	// Wrap with CORS so the browser on localhost:3000 can call us
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	fmt.Println("Backend running on http://localhost:8080")
	http.ListenAndServe(":8080", corsHandler.Handler(mux))
}
