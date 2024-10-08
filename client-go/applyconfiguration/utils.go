/*
Copyright 2023.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	v1 "sigs.k8s.io/lws/api/leaderworkerset/v1"
	internal "sigs.k8s.io/lws/client-go/applyconfiguration/internal"
	leaderworkersetv1 "sigs.k8s.io/lws/client-go/applyconfiguration/leaderworkerset/v1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=leaderworkerset.x-k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("LeaderWorkerSet"):
		return &leaderworkersetv1.LeaderWorkerSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LeaderWorkerSetSpec"):
		return &leaderworkersetv1.LeaderWorkerSetSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LeaderWorkerSetStatus"):
		return &leaderworkersetv1.LeaderWorkerSetStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LeaderWorkerTemplate"):
		return &leaderworkersetv1.LeaderWorkerTemplateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RollingUpdateConfiguration"):
		return &leaderworkersetv1.RollingUpdateConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RolloutStrategy"):
		return &leaderworkersetv1.RolloutStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubGroupPolicy"):
		return &leaderworkersetv1.SubGroupPolicyApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
