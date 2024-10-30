package authenticator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/authenticator/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator okta_authenticator}.
type Authenticator interface {
	cdktf.TerraformResource
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	LegacyIgnoreName() interface{}
	SetLegacyIgnoreName(val interface{})
	LegacyIgnoreNameInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderAuthPort() *float64
	SetProviderAuthPort(val *float64)
	ProviderAuthPortInput() *float64
	ProviderHost() *string
	SetProviderHost(val *string)
	ProviderHostInput() *string
	ProviderHostname() *string
	SetProviderHostname(val *string)
	ProviderHostnameInput() *string
	ProviderInstanceId() *string
	ProviderIntegrationKey() *string
	SetProviderIntegrationKey(val *string)
	ProviderIntegrationKeyInput() *string
	ProviderJson() *string
	SetProviderJson(val *string)
	ProviderJsonInput() *string
	ProviderSecretKey() *string
	SetProviderSecretKey(val *string)
	ProviderSecretKeyInput() *string
	ProviderSharedSecret() *string
	SetProviderSharedSecret(val *string)
	ProviderSharedSecretInput() *string
	ProviderType() *string
	ProviderUserNameTemplate() *string
	SetProviderUserNameTemplate(val *string)
	ProviderUserNameTemplateInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Settings() *string
	SetSettings(val *string)
	SettingsInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	ResetLegacyIgnoreName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProviderAuthPort()
	ResetProviderHost()
	ResetProviderHostname()
	ResetProviderIntegrationKey()
	ResetProviderJson()
	ResetProviderSecretKey()
	ResetProviderSharedSecret()
	ResetProviderUserNameTemplate()
	ResetSettings()
	ResetStatus()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Authenticator
type jsiiProxy_Authenticator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Authenticator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) LegacyIgnoreName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legacyIgnoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) LegacyIgnoreNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legacyIgnoreNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderAuthPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerAuthPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderAuthPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerAuthPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderIntegrationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIntegrationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderIntegrationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIntegrationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderSharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderSharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderUserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerUserNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) ProviderUserNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerUserNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Settings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) SettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Authenticator) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator okta_authenticator} Resource.
func NewAuthenticator(scope constructs.Construct, id *string, config *AuthenticatorConfig) Authenticator {
	_init_.Initialize()

	if err := validateNewAuthenticatorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Authenticator{}

	_jsii_.Create(
		"@cdktf/provider-okta.authenticator.Authenticator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/authenticator okta_authenticator} Resource.
func NewAuthenticator_Override(a Authenticator, scope constructs.Construct, id *string, config *AuthenticatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.authenticator.Authenticator",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Authenticator)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetLegacyIgnoreName(val interface{}) {
	if err := j.validateSetLegacyIgnoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legacyIgnoreName",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderAuthPort(val *float64) {
	if err := j.validateSetProviderAuthPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerAuthPort",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderHost(val *string) {
	if err := j.validateSetProviderHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerHost",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderHostname(val *string) {
	if err := j.validateSetProviderHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerHostname",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderIntegrationKey(val *string) {
	if err := j.validateSetProviderIntegrationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerIntegrationKey",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderJson(val *string) {
	if err := j.validateSetProviderJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerJson",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderSecretKey(val *string) {
	if err := j.validateSetProviderSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerSecretKey",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderSharedSecret(val *string) {
	if err := j.validateSetProviderSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerSharedSecret",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProviderUserNameTemplate(val *string) {
	if err := j.validateSetProviderUserNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerUserNameTemplate",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetSettings(val *string) {
	if err := j.validateSetSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_Authenticator)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a Authenticator resource upon running "cdktf plan <stack-name>".
func Authenticator_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuthenticator_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authenticator.Authenticator",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
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
func Authenticator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthenticator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authenticator.Authenticator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Authenticator_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthenticator_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authenticator.Authenticator",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Authenticator_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthenticator_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authenticator.Authenticator",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Authenticator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.authenticator.Authenticator",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Authenticator) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_Authenticator) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Authenticator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_Authenticator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_Authenticator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_Authenticator) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_Authenticator) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_Authenticator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_Authenticator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_Authenticator) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_Authenticator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_Authenticator) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_Authenticator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_Authenticator) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Authenticator) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_Authenticator) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Authenticator) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Authenticator) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetLegacyIgnoreName() {
	_jsii_.InvokeVoid(
		a,
		"resetLegacyIgnoreName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderAuthPort() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderAuthPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderHost() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderIntegrationKey() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderIntegrationKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderJson() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderSharedSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderSharedSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetProviderUserNameTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetProviderUserNameTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Authenticator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Authenticator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

