// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package file

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#filename File#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Content to store in the file, expected to be a UTF-8 encoded string.
	//
	// Conflicts with `sensitive_content`, `content_base64` and `source`.
	//  Exactly one of these four arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#content File#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content to store in the file, expected to be binary encoded as base64 string.
	//
	// Conflicts with `content`, `sensitive_content` and `source`.
	//  Exactly one of these four arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#content_base64 File#content_base64}
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Permissions to set for directories created (before umask), expressed as string in  [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).  Default value is `"0777"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#directory_permission File#directory_permission}
	DirectoryPermission *string `field:"optional" json:"directoryPermission" yaml:"directoryPermission"`
	// Permissions to set for the output file (before umask), expressed as string in  [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).  Default value is `"0777"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#file_permission File#file_permission}
	FilePermission *string `field:"optional" json:"filePermission" yaml:"filePermission"`
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	//
	// Will not be displayed in diffs.
	//  Conflicts with `content`, `content_base64` and `source`.
	//  Exactly one of these four arguments must be specified.
	//  If in need to use _sensitive_ content, please use the [`local_sensitive_file`](./sensitive_file.html)
	//  resource instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#sensitive_content File#sensitive_content}
	SensitiveContent *string `field:"optional" json:"sensitiveContent" yaml:"sensitiveContent"`
	// Path to file to use as source for the one we are creating.
	//
	// Conflicts with `content`, `sensitive_content` and `content_base64`.
	//  Exactly one of these four arguments must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/local/2.4.0/docs/resources/file#source File#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

