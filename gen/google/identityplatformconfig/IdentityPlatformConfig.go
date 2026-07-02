package identityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/identityplatformconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/identity_platform_config google_identity_platform_config}.
type IdentityPlatformConfig interface {
	cdktf.TerraformResource
	AuthorizedDomains() *[]*string
	SetAuthorizedDomains(val *[]*string)
	AuthorizedDomainsInput() *[]*string
	AutodeleteAnonymousUsers() any
	SetAutodeleteAnonymousUsers(val any)
	AutodeleteAnonymousUsersInput() any
	BlockingFunctions() IdentityPlatformConfigBlockingFunctionsOutputReference
	BlockingFunctionsInput() *IdentityPlatformConfigBlockingFunctions
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Client() IdentityPlatformConfigClientOutputReference
	ClientInput() *IdentityPlatformConfigClient
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mfa() IdentityPlatformConfigMfaOutputReference
	MfaInput() *IdentityPlatformConfigMfa
	Monitoring() IdentityPlatformConfigMonitoringOutputReference
	MonitoringInput() *IdentityPlatformConfigMonitoring
	MultiTenant() IdentityPlatformConfigMultiTenantOutputReference
	MultiTenantInput() *IdentityPlatformConfigMultiTenant
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	Quota() IdentityPlatformConfigQuotaOutputReference
	QuotaInput() *IdentityPlatformConfigQuota
	// Experimental.
	RawOverrides() any
	SignIn() IdentityPlatformConfigSignInOutputReference
	SignInInput() *IdentityPlatformConfigSignIn
	SmsRegionConfig() IdentityPlatformConfigSmsRegionConfigOutputReference
	SmsRegionConfigInput() *IdentityPlatformConfigSmsRegionConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IdentityPlatformConfigTimeoutsOutputReference
	TimeoutsInput() any
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBlockingFunctions(value *IdentityPlatformConfigBlockingFunctions)
	PutClient(value *IdentityPlatformConfigClient)
	PutMfa(value *IdentityPlatformConfigMfa)
	PutMonitoring(value *IdentityPlatformConfigMonitoring)
	PutMultiTenant(value *IdentityPlatformConfigMultiTenant)
	PutQuota(value *IdentityPlatformConfigQuota)
	PutSignIn(value *IdentityPlatformConfigSignIn)
	PutSmsRegionConfig(value *IdentityPlatformConfigSmsRegionConfig)
	PutTimeouts(value *IdentityPlatformConfigTimeouts)
	ResetAuthorizedDomains()
	ResetAutodeleteAnonymousUsers()
	ResetBlockingFunctions()
	ResetClient()
	ResetId()
	ResetMfa()
	ResetMonitoring()
	ResetMultiTenant()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetQuota()
	ResetSignIn()
	ResetSmsRegionConfig()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for IdentityPlatformConfig
type jsiiProxy_IdentityPlatformConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityPlatformConfig) AuthorizedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) AuthorizedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) AutodeleteAnonymousUsers() any {
	var returns any
	_jsii_.Get(
		j,
		"autodeleteAnonymousUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) AutodeleteAnonymousUsersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"autodeleteAnonymousUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) BlockingFunctions() IdentityPlatformConfigBlockingFunctionsOutputReference {
	var returns IdentityPlatformConfigBlockingFunctionsOutputReference
	_jsii_.Get(
		j,
		"blockingFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) BlockingFunctionsInput() *IdentityPlatformConfigBlockingFunctions {
	var returns *IdentityPlatformConfigBlockingFunctions
	_jsii_.Get(
		j,
		"blockingFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Client() IdentityPlatformConfigClientOutputReference {
	var returns IdentityPlatformConfigClientOutputReference
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) ClientInput() *IdentityPlatformConfigClient {
	var returns *IdentityPlatformConfigClient
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Mfa() IdentityPlatformConfigMfaOutputReference {
	var returns IdentityPlatformConfigMfaOutputReference
	_jsii_.Get(
		j,
		"mfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) MfaInput() *IdentityPlatformConfigMfa {
	var returns *IdentityPlatformConfigMfa
	_jsii_.Get(
		j,
		"mfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Monitoring() IdentityPlatformConfigMonitoringOutputReference {
	var returns IdentityPlatformConfigMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) MonitoringInput() *IdentityPlatformConfigMonitoring {
	var returns *IdentityPlatformConfigMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) MultiTenant() IdentityPlatformConfigMultiTenantOutputReference {
	var returns IdentityPlatformConfigMultiTenantOutputReference
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) MultiTenantInput() *IdentityPlatformConfigMultiTenant {
	var returns *IdentityPlatformConfigMultiTenant
	_jsii_.Get(
		j,
		"multiTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Quota() IdentityPlatformConfigQuotaOutputReference {
	var returns IdentityPlatformConfigQuotaOutputReference
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) QuotaInput() *IdentityPlatformConfigQuota {
	var returns *IdentityPlatformConfigQuota
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) SignIn() IdentityPlatformConfigSignInOutputReference {
	var returns IdentityPlatformConfigSignInOutputReference
	_jsii_.Get(
		j,
		"signIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) SignInInput() *IdentityPlatformConfigSignIn {
	var returns *IdentityPlatformConfigSignIn
	_jsii_.Get(
		j,
		"signInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) SmsRegionConfig() IdentityPlatformConfigSmsRegionConfigOutputReference {
	var returns IdentityPlatformConfigSmsRegionConfigOutputReference
	_jsii_.Get(
		j,
		"smsRegionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) SmsRegionConfigInput() *IdentityPlatformConfigSmsRegionConfig {
	var returns *IdentityPlatformConfigSmsRegionConfig
	_jsii_.Get(
		j,
		"smsRegionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) Timeouts() IdentityPlatformConfigTimeoutsOutputReference {
	var returns IdentityPlatformConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfig) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/identity_platform_config google_identity_platform_config} Resource.
func NewIdentityPlatformConfig(scope constructs.Construct, id *string, config *IdentityPlatformConfigConfig) IdentityPlatformConfig {
	_init_.Initialize()

	if err := validateNewIdentityPlatformConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPlatformConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/identity_platform_config google_identity_platform_config} Resource.
func NewIdentityPlatformConfig_Override(i IdentityPlatformConfig, scope constructs.Construct, id *string, config *IdentityPlatformConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		[]any{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetAuthorizedDomains(val *[]*string) {
	if err := j.validateSetAuthorizedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedDomains",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetAutodeleteAnonymousUsers(val any) {
	if err := j.validateSetAutodeleteAnonymousUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodeleteAnonymousUsers",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfig) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IdentityPlatformConfig resource upon running "cdktf plan <stack-name>".
func IdentityPlatformConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIdentityPlatformConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func IdentityPlatformConfig_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateIdentityPlatformConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityPlatformConfig_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateIdentityPlatformConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityPlatformConfig_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateIdentityPlatformConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func IdentityPlatformConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) AddOverride(path *string, value any) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]any{path, value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]any{id, provider},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]any{id},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) MoveTo(moveTarget *string, index any) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]any{id},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutBlockingFunctions(value *IdentityPlatformConfigBlockingFunctions) {
	if err := i.validatePutBlockingFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBlockingFunctions",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutClient(value *IdentityPlatformConfigClient) {
	if err := i.validatePutClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClient",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutMfa(value *IdentityPlatformConfigMfa) {
	if err := i.validatePutMfaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMfa",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutMonitoring(value *IdentityPlatformConfigMonitoring) {
	if err := i.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMonitoring",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutMultiTenant(value *IdentityPlatformConfigMultiTenant) {
	if err := i.validatePutMultiTenantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMultiTenant",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutQuota(value *IdentityPlatformConfigQuota) {
	if err := i.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putQuota",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutSignIn(value *IdentityPlatformConfigSignIn) {
	if err := i.validatePutSignInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSignIn",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutSmsRegionConfig(value *IdentityPlatformConfigSmsRegionConfig) {
	if err := i.validatePutSmsRegionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSmsRegionConfig",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) PutTimeouts(value *IdentityPlatformConfigTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]any{value},
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetAuthorizedDomains() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthorizedDomains",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetAutodeleteAnonymousUsers() {
	_jsii_.InvokeVoid(
		i,
		"resetAutodeleteAnonymousUsers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetBlockingFunctions() {
	_jsii_.InvokeVoid(
		i,
		"resetBlockingFunctions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetClient() {
	_jsii_.InvokeVoid(
		i,
		"resetClient",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetMfa() {
	_jsii_.InvokeVoid(
		i,
		"resetMfa",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetMonitoring() {
	_jsii_.InvokeVoid(
		i,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetMultiTenant() {
	_jsii_.InvokeVoid(
		i,
		"resetMultiTenant",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetProject() {
	_jsii_.InvokeVoid(
		i,
		"resetProject",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetQuota() {
	_jsii_.InvokeVoid(
		i,
		"resetQuota",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetSignIn() {
	_jsii_.InvokeVoid(
		i,
		"resetSignIn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetSmsRegionConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetSmsRegionConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfig) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfig) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
