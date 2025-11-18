// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalocalcommand

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLocalCommandConfig struct {
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
	// Executable name to be discovered on the PATH or absolute path to executable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.6.1/docs/data-sources/command#command DataLocalCommand#command}
	Command *string `field:"required" json:"command" yaml:"command"`
	// Indicates that the command returning a non-zero exit code should be treated as a successful execution.
	//
	// Further assertions can be made of the `exit_code` value with the [`check` block](https://developer.hashicorp.com/terraform/language/block/check). Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.6.1/docs/data-sources/command#allow_non_zero_exit_code DataLocalCommand#allow_non_zero_exit_code}
	AllowNonZeroExitCode interface{} `field:"optional" json:"allowNonZeroExitCode" yaml:"allowNonZeroExitCode"`
	// Arguments to be passed to the given command. Any `null` arguments will be removed from the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.6.1/docs/data-sources/command#arguments DataLocalCommand#arguments}
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Data to be passed to the given command's standard input as a UTF-8 string.
	//
	// [Terraform values](https://developer.hashicorp.com/terraform/language/expressions/types) can be encoded by any Terraform encode function, for example, [`jsonencode`](https://developer.hashicorp.com/terraform/language/functions/jsonencode).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.6.1/docs/data-sources/command#stdin DataLocalCommand#stdin}
	Stdin *string `field:"optional" json:"stdin" yaml:"stdin"`
	// The directory path where the command should be executed, either an absolute path or relative to the Terraform working directory.
	//
	// If not provided, defaults to the Terraform working directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.6.1/docs/data-sources/command#working_directory DataLocalCommand#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

