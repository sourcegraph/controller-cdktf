package dataoktaappoauth

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaAppOauth.DataOktaAppOauth",
		reflect.TypeOf((*DataOktaAppOauth)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeOnly", GoGetter: "ActiveOnly"},
			_jsii_.MemberProperty{JsiiProperty: "activeOnlyInput", GoGetter: "ActiveOnlyInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbar", GoGetter: "AutoSubmitToolbar"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientUri", GoGetter: "ClientUri"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "grantTypes", GoGetter: "GrantTypes"},
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "hideIos", GoGetter: "HideIos"},
			_jsii_.MemberProperty{JsiiProperty: "hideWeb", GoGetter: "HideWeb"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "labelPrefix", GoGetter: "LabelPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "labelPrefixInput", GoGetter: "LabelPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "links", GoGetter: "Links"},
			_jsii_.MemberProperty{JsiiProperty: "loginMode", GoGetter: "LoginMode"},
			_jsii_.MemberProperty{JsiiProperty: "loginScopes", GoGetter: "LoginScopes"},
			_jsii_.MemberProperty{JsiiProperty: "loginUri", GoGetter: "LoginUri"},
			_jsii_.MemberProperty{JsiiProperty: "logoUri", GoGetter: "LogoUri"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyUri", GoGetter: "PolicyUri"},
			_jsii_.MemberProperty{JsiiProperty: "postLogoutRedirectUris", GoGetter: "PostLogoutRedirectUris"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUris", GoGetter: "RedirectUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveOnly", GoMethod: "ResetActiveOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabelPrefix", GoMethod: "ResetLabelPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGroups", GoMethod: "ResetSkipGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipUsers", GoMethod: "ResetSkipUsers"},
			_jsii_.MemberProperty{JsiiProperty: "responseTypes", GoGetter: "ResponseTypes"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroups", GoGetter: "SkipGroups"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroupsInput", GoGetter: "SkipGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsers", GoGetter: "SkipUsers"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsersInput", GoGetter: "SkipUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "wildcardRedirect", GoGetter: "WildcardRedirect"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaAppOauth{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaAppOauth.DataOktaAppOauthConfig",
		reflect.TypeOf((*DataOktaAppOauthConfig)(nil)).Elem(),
	)
}
