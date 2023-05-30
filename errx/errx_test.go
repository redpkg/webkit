package errx_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/redpkg/webkit/errx"
	"github.com/stretchr/testify/assert"
)

func TestSetMessage(t *testing.T) {
	assert := assert.New(t)

	code := "e33e6116"
	message := "test set message"
	err := errx.New(code, "foo").SetMessage(message)

	assert.Equal(code, err.Code)
	assert.Equal(message, err.Message)
	assert.Equal(http.StatusInternalServerError, err.StatusCode)
	assert.Equal(err.Error(), fmt.Sprintf("(%s) %s", err.Code, err.Message))
}

func TestSetStatusCode(t *testing.T) {
	assert := assert.New(t)

	code := "e33e6116"
	message := "test set status code"
	err := errx.New(code, message).SetStatusCode(http.StatusUnauthorized)

	assert.Equal(code, err.Code)
	assert.Equal(message, err.Message)
	assert.Equal(http.StatusUnauthorized, err.StatusCode)
	assert.Equal(err.Error(), fmt.Sprintf("(%s) %s", err.Code, err.Message))
}

func TestSetInternal(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New("foobar")

	code := "59a7fc0a"
	message := "test set internal 1"
	err2 := errx.New(code, message).SetInternal(err1)

	assert.Equal(code, err2.Code)
	assert.Equal(message, err2.Message)
	assert.Equal(http.StatusInternalServerError, err2.StatusCode)
	assert.Equal(err2.Error(), fmt.Sprintf("(%s) %s: %s", err2.Code, err2.Message, err1.Error()))

	code = "2fd2162b"
	message = "test set internal 2"
	err3 := errx.New(code, message).SetInternal(err2)

	assert.Equal(code, err3.Code)
	assert.Equal(message, err3.Message)
	assert.Equal(http.StatusInternalServerError, err3.StatusCode)
	assert.Equal(err3.Error(), fmt.Sprintf("(%s) %s: %s", err3.Code, err3.Message, err2.Error()))
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	code := "6cd11ffd"
	message := "test new"
	err := errx.New(code, message)

	assert.Equal(code, err.Code)
	assert.Equal(message, err.Message)
	assert.Equal(http.StatusInternalServerError, err.StatusCode)
	assert.Equal(err.Error(), fmt.Sprintf("(%s) %s", err.Code, err.Message))
}

func TestCodex(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New("foobar")

	code := "932c1322"
	err2 := errx.Wrap(code, err1)

	assert.Equal(code, err2.Code)
	assert.Equal("Internal server error", err2.Message)
	assert.Equal(http.StatusInternalServerError, err2.StatusCode)
	assert.Equal(err2.Error(), fmt.Sprintf("(%s) %s: %s", err2.Code, err2.Message, err1.Error()))

	code = "2c321293"
	err3 := errx.Wrap(code, nil)

	assert.Equal(code, err3.Code)
	assert.Equal("Internal server error", err3.Message)
	assert.Equal(http.StatusInternalServerError, err3.StatusCode)
	assert.Equal(err3.Error(), fmt.Sprintf("(%s) %s", err3.Code, err3.Message))
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	err1 := errors.New("test flatten 1")
	err2 := errx.New("6912d754", "test flatten 2").SetInternal(err1)
	err3 := errx.New("d21622fb", "test flatten 3").SetInternal(err2)

	assert.EqualValues([]error{err3, err2, err1}, errx.Flatten(err3))
	assert.Equal(err3.Error(), fmt.Sprintf("(%s) %s: (%s) %s: %s", err3.Code, err3.Message, err2.Code, err2.Message, err1.Error()))
}
