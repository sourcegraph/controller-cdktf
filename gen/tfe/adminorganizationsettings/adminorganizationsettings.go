package adminorganizationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tfe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/adminorganizationsettings/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/admin_organization_settings tfe_admin_organization_settings}.
type AdminOrganizationSettings interface {
	cdktf.TerraformResource
	AccessBetaTools() interface{}
	SetAccessBetaTools(val interface{})
	AccessBetaToolsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalModuleSharing() interface{}
	SetGlobalModuleSharing(val interface{})
	GlobalModuleSharingInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModuleSharingConsumerOrganizations() *[]*string
	SetModuleSharingConsumerOrganizations(val *[]*string)
	ModuleSharingConsumerOrganizationsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SsoEnabled() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkspaceLimit() *float64
	SetWorkspaceLimit(val *float64)
	WorkspaceLimitInput() *float64
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessBetaTools()
	ResetGlobalModuleSharing()
	ResetId()
	ResetModuleSharingConsumerOrganizations()
	ResetOrganization()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetWorkspaceLimit()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AdminOrganizationSettings
type jsiiProxy_AdminOrganizationSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AdminOrganizationSettings) AccessBetaTools() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBetaTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) AccessBetaToolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBetaToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) GlobalModuleSharing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalModuleSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) GlobalModuleSharingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalModuleSharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) ModuleSharingConsumerOrganizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"moduleSharingConsumerOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) ModuleSharingConsumerOrganizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"moduleSharingConsumerOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) SsoEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ssoEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) WorkspaceLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workspaceLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdminOrganizationSettings) WorkspaceLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workspaceLimitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/admin_organization_settings tfe_admin_organization_settings} Resource.
func NewAdminOrganizationSettings(scope constructs.Construct, id *string, config *AdminOrganizationSettingsConfig) AdminOrganizationSettings {
	_init_.Initialize()

	if err := validateNewAdminOrganizationSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdminOrganizationSettings{}

	_jsii_.Create(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.42.0/docs/resources/admin_organization_settings tfe_admin_organization_settings} Resource.
func NewAdminOrganizationSettings_Override(a AdminOrganizationSettings, scope constructs.Construct, id *string, config *AdminOrganizationSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetAccessBetaTools(val interface{}) {
	if err := j.validateSetAccessBetaToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessBetaTools",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetGlobalModuleSharing(val interface{}) {
	if err := j.validateSetGlobalModuleSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalModuleSharing",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetModuleSharingConsumerOrganizations(val *[]*string) {
	if err := j.validateSetModuleSharingConsumerOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"moduleSharingConsumerOrganizations",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AdminOrganizationSettings)SetWorkspaceLimit(val *float64) {
	if err := j.validateSetWorkspaceLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceLimit",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AdminOrganizationSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdminOrganizationSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdminOrganizationSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdminOrganizationSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdminOrganizationSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdminOrganizationSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AdminOrganizationSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"tfe.adminOrganizationSettings.AdminOrganizationSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetAccessBetaTools() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessBetaTools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetGlobalModuleSharing() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalModuleSharing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetModuleSharingConsumerOrganizations() {
	_jsii_.InvokeVoid(
		a,
		"resetModuleSharingConsumerOrganizations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetOrganization() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) ResetWorkspaceLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdminOrganizationSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdminOrganizationSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

