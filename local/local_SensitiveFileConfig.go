// Prebuilt local Provider for Terraform CDK (cdktf)
package local

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SensitiveFileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// If the file already exists, it will be overridden with the given content.
	//
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#filename SensitiveFile#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Sensitive content to store in the file, expected to be an UTF-8 encoded string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#content SensitiveFile#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Sensitive content to store in the file, expected to be binary encoded as base64 string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#content_base64 SensitiveFile#content_base64}
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Permissions to set for directories created (in numeric notation).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#directory_permission SensitiveFile#directory_permission}
	DirectoryPermission *string `field:"optional" json:"directoryPermission" yaml:"directoryPermission"`
	// Permissions to set for the output file (in numeric notation).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#file_permission SensitiveFile#file_permission}
	FilePermission *string `field:"optional" json:"filePermission" yaml:"filePermission"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#id SensitiveFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Path to file to use as source for the one we are creating.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local/r/sensitive_file#source SensitiveFile#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

