// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type TlsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs#alias TlsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs#proxy TlsProvider#proxy}
	Proxy interface{} `field:"optional" json:"proxy" yaml:"proxy"`
}

