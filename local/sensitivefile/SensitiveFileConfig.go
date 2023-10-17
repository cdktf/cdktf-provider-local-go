// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivefile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SensitiveFileConfig struct {
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
	// The path to the file that will be created.
	//
	// Missing parent directories will be created.
	//  If the file already exists, it will be overridden with the given content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#filename SensitiveFile#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Sensitive Content to store in the file, expected to be a UTF-8 encoded string.
	//
	// Conflicts with `content_base64` and `source`.
	//  Exactly one of these three arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#content SensitiveFile#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Sensitive Content to store in the file, expected to be binary encoded as base64 string.
	//
	// Conflicts with `content` and `source`.
	//  Exactly one of these three arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#content_base64 SensitiveFile#content_base64}
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Permissions to set for directories created (before umask), expressed as string in  [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).  Default value is `"0700"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#directory_permission SensitiveFile#directory_permission}
	DirectoryPermission *string `field:"optional" json:"directoryPermission" yaml:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in  [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).  Default value is `"0700"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#file_permission SensitiveFile#file_permission}
	FilePermission *string `field:"optional" json:"filePermission" yaml:"filePermission"`
	// Path to file to use as source for the one we are creating.
	//
	// Conflicts with `content` and `content_base64`.
	//  Exactly one of these three arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/sensitive_file#source SensitiveFile#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

