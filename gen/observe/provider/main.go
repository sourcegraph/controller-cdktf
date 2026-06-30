package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-observe.provider.ObserveProvider",
		reflect.TypeFor[ObserveProvider](),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiToken", GoGetter: "ApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenInput", GoGetter: "ApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "customer", GoGetter: "Customer"},
			_jsii_.MemberProperty{JsiiProperty: "customerInput", GoGetter: "CustomerInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRematerializationMode", GoGetter: "DefaultRematerializationMode"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRematerializationModeInput", GoGetter: "DefaultRematerializationModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "exportObjectBindings", GoGetter: "ExportObjectBindings"},
			_jsii_.MemberProperty{JsiiProperty: "exportObjectBindingsInput", GoGetter: "ExportObjectBindingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "flags", GoGetter: "Flags"},
			_jsii_.MemberProperty{JsiiProperty: "flagsInput", GoGetter: "FlagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientTimeout", GoGetter: "HttpClientTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientTimeoutInput", GoGetter: "HttpClientTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "managingObjectId", GoGetter: "ManagingObjectId"},
			_jsii_.MemberProperty{JsiiProperty: "managingObjectIdInput", GoGetter: "ManagingObjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiToken", GoMethod: "ResetApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRematerializationMode", GoMethod: "ResetDefaultRematerializationMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomain", GoMethod: "ResetDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportObjectBindings", GoMethod: "ResetExportObjectBindings"},
			_jsii_.MemberMethod{JsiiMethod: "resetFlags", GoMethod: "ResetFlags"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientTimeout", GoMethod: "ResetHttpClientTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagingObjectId", GoMethod: "ResetManagingObjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryCount", GoMethod: "ResetRetryCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryWait", GoMethod: "ResetRetryWait"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipDatasetDryRuns", GoMethod: "ResetSkipDatasetDryRuns"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceComment", GoMethod: "ResetSourceComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceFormat", GoMethod: "ResetSourceFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserEmail", GoMethod: "ResetUserEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserPassword", GoMethod: "ResetUserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "retryCount", GoGetter: "RetryCount"},
			_jsii_.MemberProperty{JsiiProperty: "retryCountInput", GoGetter: "RetryCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "retryWait", GoGetter: "RetryWait"},
			_jsii_.MemberProperty{JsiiProperty: "retryWaitInput", GoGetter: "RetryWaitInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipDatasetDryRuns", GoGetter: "SkipDatasetDryRuns"},
			_jsii_.MemberProperty{JsiiProperty: "skipDatasetDryRunsInput", GoGetter: "SkipDatasetDryRunsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceComment", GoGetter: "SourceComment"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCommentInput", GoGetter: "SourceCommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFormat", GoGetter: "SourceFormat"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFormatInput", GoGetter: "SourceFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userEmail", GoGetter: "UserEmail"},
			_jsii_.MemberProperty{JsiiProperty: "userEmailInput", GoGetter: "UserEmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "userPassword", GoGetter: "UserPassword"},
			_jsii_.MemberProperty{JsiiProperty: "userPasswordInput", GoGetter: "UserPasswordInput"},
		},
		func() any {
			j := jsiiProxy_ObserveProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-observe.provider.ObserveProviderConfig",
		reflect.TypeFor[ObserveProviderConfig](),
	)
}
