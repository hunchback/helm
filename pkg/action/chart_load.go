/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package action

import (
	"helm.sh/helm/v3/internal/experimental/registry"
	"helm.sh/helm/v3/pkg/chart"
)

// ChartLoad performs a chart load operation.
type ChartLoad struct {
	cfg *Configuration
}

// NewChartLoad creates a new ChartLoad object with the given configuration.
func NewChartLoad(cfg *Configuration) *ChartLoad {
	return &ChartLoad{
		cfg: cfg,
	}
}

// Run executes the chart load operation
func (a *ChartLoad) Run(ref string) (*chart.Chart, error) {
	r, err := registry.ParseReference(ref)
	if err != nil {
		return nil, err
	}

	// If no tag is present, use default chart version
	if r.Tag == "" {
		r.Tag = "0.1.0"
	}

	return a.cfg.RegistryClient.LoadChart(r)
}
