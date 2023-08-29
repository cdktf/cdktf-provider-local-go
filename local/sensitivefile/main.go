// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivefile

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-local.sensitiveFile.SensitiveFile",
		reflect.TypeOf((*SensitiveFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentBase64", GoGetter: "ContentBase64"},
			_jsii_.MemberProperty{JsiiProperty: "contentBase64Input", GoGetter: "ContentBase64Input"},
			_jsii_.MemberProperty{JsiiProperty: "contentBase64Sha256", GoGetter: "ContentBase64Sha256"},
			_jsii_.MemberProperty{JsiiProperty: "contentBase64Sha512", GoGetter: "ContentBase64Sha512"},
			_jsii_.MemberProperty{JsiiProperty: "contentInput", GoGetter: "ContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentMd5", GoGetter: "ContentMd5"},
			_jsii_.MemberProperty{JsiiProperty: "contentSha1", GoGetter: "ContentSha1"},
			_jsii_.MemberProperty{JsiiProperty: "contentSha256", GoGetter: "ContentSha256"},
			_jsii_.MemberProperty{JsiiProperty: "contentSha512", GoGetter: "ContentSha512"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryPermission", GoGetter: "DirectoryPermission"},
			_jsii_.MemberProperty{JsiiProperty: "directoryPermissionInput", GoGetter: "DirectoryPermissionInput"},
			_jsii_.MemberProperty{JsiiProperty: "filename", GoGetter: "Filename"},
			_jsii_.MemberProperty{JsiiProperty: "filenameInput", GoGetter: "FilenameInput"},
			_jsii_.MemberProperty{JsiiProperty: "filePermission", GoGetter: "FilePermission"},
			_jsii_.MemberProperty{JsiiProperty: "filePermissionInput", GoGetter: "FilePermissionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetContent", GoMethod: "ResetContent"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentBase64", GoMethod: "ResetContentBase64"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectoryPermission", GoMethod: "ResetDirectoryPermission"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilePermission", GoMethod: "ResetFilePermission"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SensitiveFile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-local.sensitiveFile.SensitiveFileConfig",
		reflect.TypeOf((*SensitiveFileConfig)(nil)).Elem(),
	)
}
