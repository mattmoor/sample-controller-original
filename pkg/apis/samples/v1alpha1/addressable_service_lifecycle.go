/*
Copyright 2019 The Knative Authors.

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
	"github.com/knative/pkg/apis"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var condSet = apis.NewLivingConditionSet()

// GetGroupVersionKind implements kmeta.OwnerRefable
func (as *AddressableService) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AddressableService")
}

func (ass *AddressableServiceStatus) InitializeConditions() {
	condSet.Manage(ass).InitializeConditions()
}

func (ass *AddressableServiceStatus) MarkServiceUnavailable(name string) {
	condSet.Manage(ass).MarkFalse(
		AddressableServiceConditionReady,
		"ServiceUnavailable",
		"Service %q wasn't found.", name)
}

func (ass *AddressableServiceStatus) MarkServiceAvailable() {
	condSet.Manage(ass).MarkTrue(AddressableServiceConditionReady)
}
