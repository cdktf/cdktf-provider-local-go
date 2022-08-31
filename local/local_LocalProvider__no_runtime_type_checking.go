//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt local Provider for Terraform CDK (cdktf)
package local

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocalProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LocalProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateLocalProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLocalProviderParameters(scope constructs.Construct, id *string, config *LocalProviderConfig) error {
	return nil
}

