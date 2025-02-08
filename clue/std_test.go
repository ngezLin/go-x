package clue

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMewStd(t *testing.T) {
	code := "200"
	message := "Success"

	stdInstance := MewStd(code, message)

	assert.NotNil(t, stdInstance)
	assert.Equal(t, code, stdInstance.GetCode())
	assert.Equal(t, message, stdInstance.GetMessage())
}

// Test GetCode method
func TestStdGetCode(t *testing.T) {
	stdInstance := &std{Code: "200"}
	assert.Equal(t, "200", stdInstance.GetCode())
}

// Test GetMessage method
func TestStdGetMessage(t *testing.T) {
	stdInstance := &std{Message: "Success"}
	assert.Equal(t, "Success", stdInstance.GetMessage())
}

// Test SetCode method
func TestStdSetCode(t *testing.T) {
	stdInstance := &std{}
	stdInstance.SetCode("404")
	assert.Equal(t, "404", stdInstance.Code)
}

// Test SetMessage method
func TestStdSetMessage(t *testing.T) {
	stdInstance := &std{}
	stdInstance.SetMessage("Not Found")
	assert.Equal(t, "Not Found", stdInstance.Message)
}

// Test Templating method
func TestStdTemplating(t *testing.T) {
	clue := &Clue{
		HttpCode: http.StatusOK,
		Meta:     &std{Code: "00", Message: ""},
		Data:     nil,
	}
	ctx := context.Background()

	stdInstance := &std{}
	result := stdInstance.Templating(ctx, clue)

	assert.Equal(t, "00", result.Meta.GetCode())
	assert.Equal(t, "Successful", result.Meta.GetMessage())
}

// Test Marshall method
func TestStdMarshall(t *testing.T) {
	clue := &Clue{
		HttpCode: http.StatusOK,
		Meta:     &std{Code: "200", Message: "Success"},
		Data:     map[string]string{"key": "value"},
	}

	stdInstance := &std{Code: "200", Message: "Success"}
	jsonData, err := stdInstance.Marshall(clue)

	assert.NoError(t, err, "Expected no error during JSON marshaling")

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	assert.NoError(t, err, "Expected valid JSON output")
	assert.Equal(t, "200", result["status"], "Expected responseCode to be '200'")
	assert.Equal(t, "Success", result["message"], "Expected responseMessage to be 'Success'")
	assert.Equal(t, map[string]interface{}{
		"data":    map[string]interface{}{"key": "value"},
		"status":  "200",
		"message": "Success",
	}, result, "Expected data to match the input")
}
