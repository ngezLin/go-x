package clue

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	httpCode := http.StatusOK
	code := "200"
	data := map[string]string{"key": "value"}
	message := "Success"

	b := Build(httpCode, code, data, message)

	// Assert that the builder is not nil
	assert.NotNil(t, b)

	// Assert that the underlying clue is set correctly
	clue := b.(*builder).clue
	assert.Equal(t, httpCode, clue.HttpCode)
	assert.Equal(t, code, clue.Meta.GetCode())
	assert.Equal(t, message, clue.Meta.GetMessage())
	assert.Equal(t, data, clue.Data)
}

// Test SnapBI method
func TestSnapBI(t *testing.T) {
	b := Build(http.StatusOK, "200", nil, "Success").SnapBI()

	clue := b.(*builder).clue
	assert.Equal(t, "200", clue.Meta.GetCode())
	assert.Equal(t, "Success", clue.Meta.GetMessage())
}

// Test Std method
func TestStd(t *testing.T) {
	b := Build(http.StatusOK, "200", nil, "Success").Std()

	clue := b.(*builder).clue
	assert.Equal(t, "200", clue.Meta.GetCode())
	assert.Equal(t, "Success", clue.Meta.GetMessage())
}

// Test MarshalJSON method
func TestMarshalJSON(t *testing.T) {
	data := map[string]interface{}{"key": "value"}
	b := Build(http.StatusOK, "200", data, "Success")

	jsonData, err := b.(*builder).clue.MarshalJSON()
	assert.NoError(t, err, "Expected no error during JSON marshaling")
	fmt.Println(string(jsonData))
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	assert.NoError(t, err, "Expected valid JSON output")
	assert.Equal(t, "200", result["status"], "Expected responseCode to be '200'")
	assert.Equal(t, "Success", result["message"], "Expected responseMessage to be 'Success'")
	assert.Equal(t, data, result["data"], "Expected data to match the input")
}

// Test CoverBuilder function
func TestCoverBuilder(t *testing.T) {
	b := Build(http.StatusOK, "200", nil, "Success")
	coveredBuilder := CoverBuilder(b, map[string]string{"key": "new value"})

	clue := coveredBuilder.(*builder).clue
	assert.Equal(t, http.StatusOK, clue.HttpCode)
	assert.Equal(t, "200", clue.Meta.GetCode())
	assert.Equal(t, "Success", clue.Meta.GetMessage())
	assert.Equal(t, map[string]string{"key": "new value"}, clue.Data)

	// Test with an error
	err := errors.New("an error occurred")
	coveredBuilder = CoverBuilder(err, nil)

	clue = coveredBuilder.(*builder).clue
	assert.Equal(t, http.StatusInternalServerError, clue.HttpCode)
	assert.Equal(t, "00", clue.Meta.GetCode())
	assert.Equal(t, "an error occurred", clue.Meta.GetMessage())
	assert.Nil(t, clue.Data)
}
