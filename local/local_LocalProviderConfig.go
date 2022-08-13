// Prebuilt local Provider for Terraform CDK (cdktf)
package local


type LocalProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/local#alias LocalProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

