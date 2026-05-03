import { render, screen, fireEvent } from "@testing-library/react";
import Home from "./page";
import { calculatorClient } from "./client";
import { ConnectError } from "@connectrpc/connect";

jest.mock("./client", () => ({
  calculatorClient: { calculate: jest.fn() },
}));

const mockCalculate = calculatorClient.calculate as jest.Mock;

beforeEach(() => {
  mockCalculate.mockReset();
});

test("inputs and button show up", () => {
  render(<Home />);
  expect(screen.getAllByRole("spinbutton")).toHaveLength(2);
  expect(screen.getByRole("button", { name: "=" })).toBeInTheDocument();
});

test("shows the result", async () => {
  mockCalculate.mockResolvedValue({ result: 15, expression: "10 + 5 = 15" });
  render(<Home />);
  const inputs = screen.getAllByRole("spinbutton");
  fireEvent.change(inputs[0], { target: { value: "10" } });
  fireEvent.change(inputs[1], { target: { value: "5" } });
  fireEvent.click(screen.getByRole("button", { name: "=" }));

  expect(await screen.findByText("10 + 5 = 15")).toBeInTheDocument();
});

test("shows error when dividing by zero", async () => {
  mockCalculate.mockRejectedValue(new ConnectError("cant divide by zero"));
  render(<Home />);
  const inputs = screen.getAllByRole("spinbutton");
  fireEvent.change(inputs[0], { target: { value: "5" } });
  fireEvent.change(inputs[1], { target: { value: "0" } });
  fireEvent.change(screen.getByRole("combobox"), { target: { value: "divide" } });
  fireEvent.click(screen.getByRole("button", { name: "=" }));

  expect(await screen.findByText(/cant divide by zero/)).toBeInTheDocument();
});
