package phpipam

import (
	"encoding/json"
	"os"
	"testing"
)

const testBoolIntStringJSONTrue = `{"foo":"1"}`
const testBoolIntStringJSONFalse = `{"foo":"0"}`
const testBoolIntStringJSONError = `{"foo":"2"}`

type testBoolIntStringType struct {
	Foo BoolIntString `json:"foo"`
}

func setPHPIPAMenv() {
	os.Setenv("PHPIPAM_APP_ID", "foobar")
	os.Setenv("PHPIPAM_ENDPOINT_ADDR", "https://example.com/phpipam/api")
	os.Setenv("PHPIPAM_PASSWORD", "abcdefgh0123456789")
	os.Setenv("PHPIPAM_USER_NAME", "nobody")
}

func unsetPHPIPAMenv() {
	os.Unsetenv("PHPIPAM_APP_ID")
	os.Unsetenv("PHPIPAM_ENDPOINT_ADDR")
	os.Unsetenv("PHPIPAM_PASSWORD")
	os.Unsetenv("PHPIPAM_USER_NAME")
}

func TestPHPIPAMDefaultConfigProviderWithEnv(t *testing.T) {
	setPHPIPAMenv()
	c := DefaultConfigProvider()
	if c.Endpoint != "https://example.com/phpipam/api" {
		t.Fatalf("Expected Endpoint to be https://example.com/phpipam/api, got %s", c.Endpoint)
	}
	if c.Username != "nobody" {
		t.Fatalf("Expected Username to be nobody, got %s", c.Username)
	}
	if c.Password != "abcdefgh0123456789" {
		t.Fatalf("Expected Password to be abcdefgh0123456789, got %s", c.Password)
	}
	if c.AppID != "foobar" {
		t.Fatalf("Expected AppID to be foobar, got %s", c.AppID)
	}
}

func TestPHPIPAMDefaultConfigProviderNoEnv(t *testing.T) {
	unsetPHPIPAMenv()
	c := DefaultConfigProvider()
	if c.Endpoint != "http://localhost/api" {
		t.Fatalf("Expected Endpoint to be http://localhost/api, got %s", c.Endpoint)
	}
	if c.Username != "" {
		t.Fatalf("Expected Username to be empty, got %s", c.Username)
	}
	if c.Password != "" {
		t.Fatalf("Expected Password to be empty, got %s", c.Password)
	}
	if c.AppID != "" {
		t.Fatalf("Expected AppID to be empty, got %s", c.AppID)
	}
}

func TestBoolIntStringUnmarshalJSONTrue(t *testing.T) {
	var actual testBoolIntStringType
	if err := json.Unmarshal([]byte(testBoolIntStringJSONTrue), &actual); err != nil {
		t.Fatalf("Bad: %s", err)
	}
	if actual.Foo != true {
		t.Fatalf("Expected value to be true, got %t", actual)
	}
}

func TestBoolIntStringUnmarshalJSONFalse(t *testing.T) {
	var actual testBoolIntStringType
	if err := json.Unmarshal([]byte(testBoolIntStringJSONFalse), &actual); err != nil {
		t.Fatalf("Bad: %s", err)
	}
	if actual.Foo != false {
		t.Fatalf("Expected value to be false, got %t", actual)
	}
}

func TestBoolIntStringUnmarshalJSONError(t *testing.T) {
	var v testBoolIntStringType
	err := json.Unmarshal([]byte(testBoolIntStringJSONError), &v)
	if err == nil {
		t.Fatalf("Expected error, got none")
	}

	expected := "json: cannot unmarshal bool into Go struct field testBoolIntStringType.foo of type string"
	actual := err.Error()
	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestBoolIntStringMarshalJSONTrue(t *testing.T) {
	v := testBoolIntStringType{
		Foo: true,
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}
	expected := testBoolIntStringJSONTrue
	actual := string(b)
	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestBoolIntStringMarshalJSONFalse(t *testing.T) {
	v := testBoolIntStringType{
		Foo: false,
	}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("Bad: %s", err)
	}
	expected := testBoolIntStringJSONFalse
	actual := string(b)
	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}
