package dereference

import (
	"reflect"
	"testing"
)

// Struct definitions for testing
type Spec struct {
	Enabled  bool
	Path     string
	User     string
	Password string
}

type Profiler struct {
	Enabled  bool
	User     string
	Password string
}

type Gzip struct {
	Enabled bool
	Level   int
}

type TLS struct {
	Enabled  bool
	CertFile string
	KeyFile  string
}

type Config struct {
	Server                *int64
	Name                  *string
	Type                  *string
	Version               *string
	Port                  int
	Spec                  Spec
	Profiler              Profiler
	InternalRespDetail    bool
	CORS                  *string
	Endpoint              *string
	Gzip                  Gzip
	TLS                   TLS
	OmitSuccessLogs       bool
	EnableConfigEndpoints bool
}

// Know that this test won't include CORS which will not pass the test if included (set the pointer to nil)
func TestDereferenceStruct(t *testing.T) {
	// Variables used for pointers
	name := "dataset-api"
	server := int64(123456789)
	endpoint := "http://localhost:9720"
	cors := "CORS config here"

	// Struct with pointers
	configWithPointers := &Config{
		Server:                &server,
		Name:                  &name,
		Type:                  &name,
		Version:               &name,
		Port:                  9720,
		Spec:                  Spec{Enabled: false, Path: "", User: "", Password: ""},
		Profiler:              Profiler{Enabled: true, User: "ookla", Password: "thundarr"},
		InternalRespDetail:    true,
		CORS:                  &cors,
		Endpoint:              &endpoint,
		Gzip:                  Gzip{Enabled: false, Level: 0},
		TLS:                   TLS{Enabled: false, CertFile: "", KeyFile: ""},
		OmitSuccessLogs:       false,
		EnableConfigEndpoints: true,
	}

	// We need to dereference the values manually.
	configWithoutPointers := Config{
		Server:                &server,
		Name:                  &name,
		Type:                  &name,
		Version:               &name,
		Port:                  9720,
		Spec:                  Spec{Enabled: false, Path: "", User: "", Password: ""},
		Profiler:              Profiler{Enabled: true, User: "ookla", Password: "thundarr"},
		InternalRespDetail:    true,
		CORS:                  &cors,
		Endpoint:              &endpoint,
		Gzip:                  Gzip{Enabled: false, Level: 0},
		TLS:                   TLS{Enabled: false, CertFile: "", KeyFile: ""},
		OmitSuccessLogs:       false,
		EnableConfigEndpoints: true,
	}

	result := Dereference(configWithPointers)

	// To check if the type is of config
	dereferencedStruct, ok := result.(Config)
	if !ok {
		t.Fatalf("Expected result to be of type Config, got %T", result)
	}

	// Compare the two structs
	if !reflect.DeepEqual(dereferencedStruct, configWithoutPointers) {
		t.Errorf("Dereferenced struct does not match expected struct.\nExpected: %+v\nGot: %+v", configWithoutPointers, dereferencedStruct)
	}
}
