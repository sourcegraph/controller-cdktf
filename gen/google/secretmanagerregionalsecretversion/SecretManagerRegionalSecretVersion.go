package secretmanagerregionalsecretversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/secretmanagerregionalsecretversion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/secret_manager_regional_secret_version google_secret_manager_regional_secret_version}.
type SecretManagerRegionalSecretVersion interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreateTime() *string
	CustomerManagedEncryption() SecretManagerRegionalSecretVersionCustomerManagedEncryptionList
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestroyTime() *string
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
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
	IsSecretDataBase64() any
	SetIsSecretDataBase64(val any)
	IsSecretDataBase64Input() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Secret() *string
	SetSecret(val *string)
	SecretData() *string
	SetSecretData(val *string)
	SecretDataInput() *string
	SecretInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SecretManagerRegionalSecretVersionTimeoutsOutputReference
	TimeoutsInput() any
	Version() *string
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
	PutTimeouts(value *SecretManagerRegionalSecretVersionTimeouts)
	ResetDeletionPolicy()
	ResetEnabled()
	ResetId()
	ResetIsSecretDataBase64()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for SecretManagerRegionalSecretVersion
type jsiiProxy_SecretManagerRegionalSecretVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) CustomerManagedEncryption() SecretManagerRegionalSecretVersionCustomerManagedEncryptionList {
	var returns SecretManagerRegionalSecretVersionCustomerManagedEncryptionList
	_jsii_.Get(
		j,
		"customerManagedEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) DestroyTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destroyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) IsSecretDataBase64() any {
	var returns any
	_jsii_.Get(
		j,
		"isSecretDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) IsSecretDataBase64Input() any {
	var returns any
	_jsii_.Get(
		j,
		"isSecretDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SecretData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SecretDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Timeouts() SecretManagerRegionalSecretVersionTimeoutsOutputReference {
	var returns SecretManagerRegionalSecretVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/secret_manager_regional_secret_version google_secret_manager_regional_secret_version} Resource.
func NewSecretManagerRegionalSecretVersion(scope constructs.Construct, id *string, config *SecretManagerRegionalSecretVersionConfig) SecretManagerRegionalSecretVersion {
	_init_.Initialize()

	if err := validateNewSecretManagerRegionalSecretVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretManagerRegionalSecretVersion{}

	_jsii_.Create(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/secret_manager_regional_secret_version google_secret_manager_regional_secret_version} Resource.
func NewSecretManagerRegionalSecretVersion_Override(s SecretManagerRegionalSecretVersion, scope constructs.Construct, id *string, config *SecretManagerRegionalSecretVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		[]any{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetIsSecretDataBase64(val any) {
	if err := j.validateSetIsSecretDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecretDataBase64",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretVersion) SetSecretData(val *string) {
	if err := j.validateSetSecretDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretData",
		val,
	)
}

// Generates CDKTF code for importing a SecretManagerRegionalSecretVersion resource upon running "cdktf plan <stack-name>".
func SecretManagerRegionalSecretVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecretManagerRegionalSecretVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
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
func SecretManagerRegionalSecretVersion_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateSecretManagerRegionalSecretVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretManagerRegionalSecretVersion_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateSecretManagerRegionalSecretVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretManagerRegionalSecretVersion_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateSecretManagerRegionalSecretVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func SecretManagerRegionalSecretVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.secretManagerRegionalSecretVersion.SecretManagerRegionalSecretVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) AddOverride(path *string, value any) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]any{path, value},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]any{id, provider},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]any{id},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) MoveTo(moveTarget *string, index any) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]any{id},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) PutTimeouts(value *SecretManagerRegionalSecretVersionTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]any{value},
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetIsSecretDataBase64() {
	_jsii_.InvokeVoid(
		s,
		"resetIsSecretDataBase64",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretVersion) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
