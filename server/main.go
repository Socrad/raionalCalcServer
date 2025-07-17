package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type RationalCalcParams struct {
	Operand1 string `json:"operand1"`
	Operator string `json:"operator"`
	Operand2 string `json:"operand2"`
}

func RationalCalc(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[RationalCalcParams]) (*mcp.CallToolResultFor[any], error) {
	calculator_path := "./rational_calculator"
	cmd := exec.Command(calculator_path, params.Arguments.Operand1, params.Arguments.Operator, params.Arguments.Operand2)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("calculator 실행 오류: %s, Output: %s", err, string(output))
		return nil, fmt.Errorf("rational_calculator 실행 실패: %s", strings.TrimSpace(string(output)))
	}
	result := strings.TrimSpace(string(output))
	log.Printf("rational_calculator 도구 호출됨: %s %s %s = %s", params.Arguments.Operand1, params.Arguments.Operator, params.Arguments.Operand2, result)
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: result}},
	}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "adder-server", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "rational_calculator", Description: "정수와 유리수를 오차없이 연산합니다. 일상적인 사칙연산과 크기비교는 이 툴을 사용합니다"}, RationalCalc)

	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
