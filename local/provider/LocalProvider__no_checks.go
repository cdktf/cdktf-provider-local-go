// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LocalProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_LocalProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateLocalProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLocalProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLocalProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLocalProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewLocalProviderParameters(scope constructs.Construct, id *string, config *LocalProviderConfig) error {
	return nil
}

