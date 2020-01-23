/*

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectExporterSpec defines the desired state of ObjectExporter
type ObjectExporterSpec struct {
	Identifier ResourceIdentifier `json:"identifier,omitempty"`

	// Intruments is a list of time series specs that will be exported for the Resource Identifier
	Intruments []InstrumentSpec `json:"identifier,omitempty"`

	// Description regroups information and metadata about an application.
	Description string `json:"description,omitempty"`
}

// ResourceIdentifier is the identifier of the underlying kube resource
type ResourceIdentifier struct {

	// GCK indicated the GroupVersionKind of the resource we are tracking
	metav1.GroupVersionKind `json:",inline"`

	// Selector is a label query over the indicated GVK to filter out resource not desired for this exporter
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

// InstrumentType represents the time series metric type. https://prometheus.io/docs/concepts/metric_types/
type InstrumentType string

const (
	// GaugeInstrumentType represents a prometheus gauge time series type. See: https://prometheus.io/docs/concepts/metric_types/#gauge
	GaugeInstrumentType InstrumentType = "Gauge"

	// CounterInstrumentType represents a prometheus gauge time series type. See: https://prometheus.io/docs/concepts/metric_types/#counter
	CounterInstrumentType InstrumentType = "Counter"
)

// InstrumentSpec configures a time
type InstrumentSpec struct {

	// Name is the subresource being intrumented or tracked
	Name string `json:"name,omitempty"`

	// UnitAggregation refers to how the Unit is being tracked. _count or _total for sum,
	UnitAggregation string `json:"unit,omitempty"`

	// Type represents the time series metric type. https://prometheus.io/docs/concepts/metric_types/
	Type InstrumentType `json:"type,omitempty"`

	// ValueJSONPath sets the field or strategy for determining the time series value
	ValueJSONPath string `json:"valueJSONPath,omitempty"`

	// AdditionalLabelsFromFields is formatted as LabelName:
	AdditionalLabelsFromFields map[string]string `json:"additionalLabelsFromFields,omitempty"`
}

// ObjectExporterStatus defines the observed state of ObjectExporter
type ObjectExporterStatus struct {

	// A count of objects being counted
	InstrumentedResourceCount int64 `json:"instrumentedResourceCount,omitempty"`

	// ExportedTimeSeries is a preview of the time series exported
	ExportedTimeSeries []string `json:"exportedTimeSeries,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ObjectExporter is the Schema for the objectexporters API
type ObjectExporter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectExporterSpec   `json:"spec,omitempty"`
	Status ObjectExporterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectExporterList contains a list of ObjectExporter
type ObjectExporterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectExporter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObjectExporter{}, &ObjectExporterList{})
}
