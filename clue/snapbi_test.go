package clue

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMewSnapBI(t *testing.T) {
	code := "200"
	message := "Success"

	snapInstance := MewSnapBI(code, message)

	assert.NotNil(t, snapInstance)
	assert.Equal(t, code, snapInstance.GetCode())
	assert.Equal(t, message, snapInstance.GetMessage())
}

// Test GetCode method
func TestSnapBIGetCode(t *testing.T) {
	snapInstance := &snap{Code: "200"}
	assert.Equal(t, "200", snapInstance.GetCode())
}

// Test GetMessage method
func TestSnapBIGetMessage(t *testing.T) {
	snapInstance := &snap{Message: "Success"}
	assert.Equal(t, "Success", snapInstance.GetMessage())
}

// Test SetCode method
func TestSnapBISetCode(t *testing.T) {
	snapInstance := &snap{}
	snapInstance.SetCode("404")
	assert.Equal(t, "404", snapInstance.Code)
}

// Test SetMessage method
func TestSnapBISetMessage(t *testing.T) {
	snapInstance := &snap{}
	snapInstance.SetMessage("Not Found")
	assert.Equal(t, "Not Found", snapInstance.Message)
}

// Test Templating method
func TestSnapBITemplating(t *testing.T) {
	clue := &Clue{
		HttpCode: http.StatusOK,
		Meta:     &snap{Code: "00", Message: ""},
		Data:     nil,
	}
	ctx := context.Background()
	ctx = DefineCtxServiceCode(ctx, "XX")

	snapInstance := &snap{}
	result := snapInstance.Templating(ctx, clue)

	assert.Equal(t, "200XX00", result.Meta.GetCode())
	assert.Equal(t, "Successful", result.Meta.GetMessage())
}

// Test Marshall method
func TestSnapBIMarshall(t *testing.T) {
	clue := &Clue{
		HttpCode: http.StatusOK,
		Meta:     &snap{Code: "200", Message: "Success"},
		Data:     map[string]string{"key": "value"},
	}

	snapInstance := &snap{Code: "200", Message: "Success"}
	jsonData, err := snapInstance.Marshall(clue)

	fmt.Println(string(jsonData))

	assert.NoError(t, err, "Expected no error during JSON marshaling")

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	assert.NoError(t, err, "Expected valid JSON output")
	assert.Equal(t, "200", result["responseCode"], "Expected responseCode to be '200'")
	assert.Equal(t, "Success", result["responseMessage"], "Expected responseMessage to be 'Success'")
	assert.Equal(t, map[string]interface{}{
		"key":             "value",
		"responseCode":    "200",
		"responseMessage": "Success",
	}, result, "Expected data to match the input")
}
