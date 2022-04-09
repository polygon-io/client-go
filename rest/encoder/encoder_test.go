package encoder_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/encoder"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestEncodeParams(t *testing.T) {
	testPath := "/v1/{num}/{str}"

	type Params struct {
		Num float64 `validate:"required" path:"num"`
		Str string  `validate:"required" path:"str"`

		NumQ *float64 `query:"num"`
		StrQ *string  `query:"str"`
	}

	num := 1.6302
	str := "testing"
	params := Params{
		Num:  1.6302,
		Str:  str,
		NumQ: &num,
		StrQ: &str,
	}

	expected := "/v1/1.6302/testing?num=1.6302&str=testing"
	actual, err := encoder.New().EncodeParams(testPath, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeTime(t *testing.T) {
	testPath := "/v1/{time}"

	type Params struct {
		Time  models.Time  `validate:"required" path:"time"`
		TimeQ *models.Time `query:"time"`
	}

	ptime := models.Time(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Time:  ptime,
		TimeQ: &ptime,
	}

	expected := "/v1/2021-07-22T00:00:00.000Z?time=2021-07-22T00%3A00%3A00.000Z"
	actual, err := encoder.New().EncodeParams(testPath, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeDate(t *testing.T) {
	testPath := "/v1/{date}"

	type Params struct {
		Date  models.Date  `validate:"required" path:"date"`
		DateQ *models.Date `query:"date"`
	}

	pdate := models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Date:  pdate,
		DateQ: &pdate,
	}

	expected := "/v1/2021-07-22?date=2021-07-22"
	actual, err := encoder.New().EncodeParams(testPath, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeMillis(t *testing.T) {
	testPath := "/v1/{millis}"

	type Params struct {
		Millis  models.Millis  `validate:"required" path:"millis"`
		MillisQ *models.Millis `query:"millis"`
	}

	pmillis := models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Millis:  pmillis,
		MillisQ: &pmillis,
	}

	expected := "/v1/1626912000000?millis=1626912000000"
	actual, err := encoder.New().EncodeParams(testPath, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestEncodeNanos(t *testing.T) {
	testPath := "/v1/{nanos}"

	type Params struct {
		Nanos  models.Nanos  `validate:"required" path:"nanos"`
		NanosQ *models.Nanos `query:"nanos"`
	}

	pnanos := models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := Params{
		Nanos:  pnanos,
		NanosQ: &pnanos,
	}

	expected := "/v1/1626912000000000000?nanos=1626912000000000000"
	actual, err := encoder.New().EncodeParams(testPath, params)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestValidateError(t *testing.T) {
	_, err := encoder.New().EncodeParams("/v1/test", nil)
	assert.NotNil(t, err)
}
