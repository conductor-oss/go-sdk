package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func getWeather(_ context.Context, in cityIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
}

func getStockPrice(_ context.Context, in symbolIn) (map[string]any, error) {
	return map[string]any{"symbol": in.Symbol, "price": 182.50, "change": "+1.2%"}, nil
}

// SimpleTools is the Python SDK's examples/agents/02a_simple_tools.py: two
// worker tools, and the model picks the right one.
func SimpleTools(model string) *ai.Agent {
	return &ai.Agent{
		Name:  "weather_stock_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("get_weather", getWeather, "Get the current weather for a city."),
			tool.Func("get_stock_price", getStockPrice, "Get the current stock price for a ticker symbol."),
		),
		Instructions: "You are a helpful assistant. Use tools to answer questions.",
	}
}

func runSimpleTools(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, SimpleTools(Model()), "What's the weather like in San Francisco?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
