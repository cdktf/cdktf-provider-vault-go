// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavaultkubernetesauthbackendconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVaultKubernetesAuthBackendConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Unique name of the kubernetes backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#backend DataVaultKubernetesAuthBackendConfig#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#disable_iss_validation DataVaultKubernetesAuthBackendConfig#disable_iss_validation}
	DisableIssValidation interface{} `field:"optional" json:"disableIssValidation" yaml:"disableIssValidation"`
	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#disable_local_ca_jwt DataVaultKubernetesAuthBackendConfig#disable_local_ca_jwt}
	DisableLocalCaJwt interface{} `field:"optional" json:"disableLocalCaJwt" yaml:"disableLocalCaJwt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#id DataVaultKubernetesAuthBackendConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#issuer DataVaultKubernetesAuthBackendConfig#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#kubernetes_ca_cert DataVaultKubernetesAuthBackendConfig#kubernetes_ca_cert}
	KubernetesCaCert *string `field:"optional" json:"kubernetesCaCert" yaml:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#kubernetes_host DataVaultKubernetesAuthBackendConfig#kubernetes_host}
	KubernetesHost *string `field:"optional" json:"kubernetesHost" yaml:"kubernetesHost"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#namespace DataVaultKubernetesAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs.
	//
	// If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/data-sources/kubernetes_auth_backend_config#pem_keys DataVaultKubernetesAuthBackendConfig#pem_keys}
	PemKeys *[]*string `field:"optional" json:"pemKeys" yaml:"pemKeys"`
}

