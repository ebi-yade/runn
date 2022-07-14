package runn

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type httpRunnerConfig struct {
	Endpoint             string `yaml:"endpoint"`
	OpenApi3DocLocation  string `yaml:"openapi3,omitempty"`
	SkipValidateRequest  bool   `yaml:"skipValidateRequest,omitempty"`
	SkipValidateResponse bool   `yaml:"skipValidateResponse,omitempty"`

	openApi3Doc *openapi3.T
}

type grpcRunnerConfig struct {
	Addr   string `yaml:"addr"`
	TLS    *bool  `yaml:"tls,omitempty"`
	CACert string `yaml:"cacert,omitempty"`
	Cert   string `yaml:"cert,omitempty"`
	Key    string `yaml:"key,omitempty"`
}

type httpRunnerOption func(*httpRunnerConfig) error

func OpenApi3(l string) httpRunnerOption {
	return func(c *httpRunnerConfig) error {
		c.OpenApi3DocLocation = l
		return nil
	}
}

func OpenApi3FromData(d []byte) httpRunnerOption {
	return func(c *httpRunnerConfig) error {
		ctx := context.Background()
		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromData(d)
		if err != nil {
			return err
		}
		if err := doc.Validate(ctx); err != nil {
			return fmt.Errorf("openapi document validation error: %w", err)
		}
		c.openApi3Doc = doc
		return nil
	}
}

func SkipValidateRequest(skip bool) httpRunnerOption {
	return func(c *httpRunnerConfig) error {
		c.SkipValidateRequest = skip
		return nil
	}
}

func SkipValidateResponse(skip bool) httpRunnerOption {
	return func(c *httpRunnerConfig) error {
		c.SkipValidateResponse = skip
		return nil
	}
}
