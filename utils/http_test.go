package utils

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)


type Agency struct {
	Key string `json:"key"`
	Id string `json:"agencyId"`
	Name string `json:"name"`
	Url string `json:"url"`
	Timezone string `json:"timezone"`
	Lang string `json:"lang"`
}

func TestStub(t *testing.T) {

	rr := httptest.NewRecorder()

	agency := Agency{"Key", "Id", "Name", "URL", "Timezone", "Lang"}
	SendJSON(rr, agency)

	expected :=  "{\n  \"key\": \"Key\",\n  \"agencyId\": \"Id\",\n  \"name\": \"Name\",\n  \"url\": \"URL\",\n  \"timezone\": \"Timezone\",\n  \"lang\": \"Lang\"\n}"

	assert.Equal(t, expected, string(rr.Body.String()))
}
