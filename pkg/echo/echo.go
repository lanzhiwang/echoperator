/*
MIT License

Copyright (c) 2021 Martín Montes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package echo

import (
	version "github.com/mmontes11/echoperator/pkg/echo/version"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GroupName = "mmontes.io"

	EchoKind     string = "Echo"
	EchoListKind string = "EchoList"
	EchoPlural   string = "echos"
	EchoSingular string = "echo"
	EchoCRDName  string = EchoPlural + "." + GroupName

	ScheduledEchoKind     string = "ScheduledEcho"
	ScheduledEchoListKind string = "ScheduledEchoList"
	ScheduledEchoPlural   string = "scheduledechos"
	ScheduledEchoSingular string = "scheduledecho"
	ScheduledEchoCRDName  string = ScheduledEchoPlural + "." + GroupName
)

var (
	EchoShortNames          []string = []string{"ec"}
	ScheduledEchoShortNames []string = []string{"sec"}
)

var EchoCRD = extv1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{
		Name: EchoCRDName,
	},
	Spec: extv1.CustomResourceDefinitionSpec{
		Group: GroupName,
		Names: extv1.CustomResourceDefinitionNames{
			Kind:       EchoKind,
			ListKind:   EchoListKind,
			Plural:     EchoPlural,
			Singular:   EchoSingular,
			ShortNames: EchoShortNames,
		},
		Scope: extv1.NamespaceScoped,
		Versions: []extv1.CustomResourceDefinitionVersion{
			{
				Name:    version.V1alpha1,
				Served:  true,
				Storage: true,
				Schema: &extv1.CustomResourceValidation{
					OpenAPIV3Schema: &extv1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]extv1.JSONSchemaProps{
							"spec": {
								Type: "object",
								Properties: map[string]extv1.JSONSchemaProps{
									"message": {Type: "string"},
								},
								Required: []string{"message"},
							},
						},
						Required: []string{"spec"},
					},
				},
			},
		},
	},
}

var ScheduledEchoCRD = extv1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{
		Name: ScheduledEchoCRDName,
	},
	Spec: extv1.CustomResourceDefinitionSpec{
		Group: GroupName,
		Names: extv1.CustomResourceDefinitionNames{
			Kind:       ScheduledEchoKind,
			ListKind:   ScheduledEchoListKind,
			Plural:     ScheduledEchoPlural,
			Singular:   ScheduledEchoSingular,
			ShortNames: ScheduledEchoShortNames,
		},
		Scope: extv1.NamespaceScoped,
		Versions: []extv1.CustomResourceDefinitionVersion{
			{
				Name:    version.V1alpha1,
				Served:  true,
				Storage: true,
				Schema: &extv1.CustomResourceValidation{
					OpenAPIV3Schema: &extv1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]extv1.JSONSchemaProps{
							"spec": {
								Type: "object",
								Properties: map[string]extv1.JSONSchemaProps{
									"message":  {Type: "string"},
									"schedule": {Type: "string"},
								},
								Required: []string{"message", "schedule"},
							},
						},
						Required: []string{"spec"},
					},
				},
			},
		},
	},
}
