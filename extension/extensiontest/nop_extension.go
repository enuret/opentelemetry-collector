// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extensiontest // import "go.opentelemetry.io/collector/extension/extensiontest"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/extension"
)

var nopType = component.MustNewType("nop")

// NewNopCreateSettings returns a new nop settings for extension.Factory Create* functions.
func NewNopCreateSettings() extension.CreateSettings {
	return extension.CreateSettings{
		ID:                component.NewID(nopType),
		TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

// NewNopFactory returns an extension.Factory that constructs nop extensions.
func NewNopFactory() extension.Factory {
	return extension.NewFactory(
		nopType,
		func() component.Config {
			return &nopConfig{}
		},
		func(context.Context, extension.CreateSettings, component.Config) (extension.Extension, error) {
			return nopInstance, nil
		},
		component.StabilityLevelStable)
}

type nopConfig struct{}

var nopInstance = &nopExtension{}

// nopExtension stores consumed traces and metrics for testing purposes.
type nopExtension struct {
	component.StartFunc
	component.ShutdownFunc
}

// NewNopBuilder returns a extension.Builder that constructs nop receivers.
func NewNopBuilder() *extension.Builder {
	nopFactory := NewNopFactory()
	return extension.NewBuilder(
		map[component.ID]component.Config{component.NewID(nopType): nopFactory.CreateDefaultConfig()},
		map[component.Type]extension.Factory{nopType: nopFactory})
}
