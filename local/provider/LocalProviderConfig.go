// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type LocalProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.1/docs#alias LocalProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

