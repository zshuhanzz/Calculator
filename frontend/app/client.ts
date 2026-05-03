import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { CalculatorService } from "../gen/calculator/calculator_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

export const calculatorClient = createClient(CalculatorService, transport);
