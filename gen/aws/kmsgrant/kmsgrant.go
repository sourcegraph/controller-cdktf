package kmsgrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/kmsgrant/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kms_grant aws_kms_grant}.
type KmsGrant interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	Constraints() KmsGrantConstraintsList
	ConstraintsInput() any
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
	GrantCreationTokens() *[]*string
	SetGrantCreationTokens(val *[]*string)
	GrantCreationTokensInput() *[]*string
	GranteePrincipal() *string
	SetGranteePrincipal(val *string)
	GranteePrincipalInput() *string
	GrantId() *string
	GrantToken() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Operations() *[]*string
	SetOperations(val *[]*string)
	OperationsInput() *[]*string
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
	RetireOnDelete() any
	SetRetireOnDelete(val any)
	RetireOnDeleteInput() any
	RetiringPrincipal() *string
	SetRetiringPrincipal(val *string)
	RetiringPrincipalInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
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
	PutConstraints(value any)
	ResetConstraints()
	ResetGrantCreationTokens()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetireOnDelete()
	ResetRetiringPrincipal()
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

// The jsii proxy struct for KmsGrant
type jsiiProxy_KmsGrant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KmsGrant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Constraints() KmsGrantConstraintsList {
	var returns KmsGrantConstraintsList
	_jsii_.Get(
		j,
		"constraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) ConstraintsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"constraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GrantCreationTokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantCreationTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GrantCreationTokensInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantCreationTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GranteePrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granteePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GranteePrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granteePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GrantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) GrantToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Operations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) OperationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) RetireOnDelete() any {
	var returns any
	_jsii_.Get(
		j,
		"retireOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) RetireOnDeleteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"retireOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) RetiringPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retiringPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) RetiringPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retiringPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsGrant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kms_grant aws_kms_grant} Resource.
func NewKmsGrant(scope constructs.Construct, id *string, config *KmsGrantConfig) KmsGrant {
	_init_.Initialize()

	if err := validateNewKmsGrantParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmsGrant{}

	_jsii_.Create(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/kms_grant aws_kms_grant} Resource.
func NewKmsGrant_Override(k KmsGrant, scope constructs.Construct, id *string, config *KmsGrantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		[]any{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmsGrant) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetGrantCreationTokens(val *[]*string) {
	if err := j.validateSetGrantCreationTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantCreationTokens",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetGranteePrincipal(val *string) {
	if err := j.validateSetGranteePrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granteePrincipal",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetKeyId(val *string) {
	if err := j.validateSetKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetOperations(val *[]*string) {
	if err := j.validateSetOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operations",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetRetireOnDelete(val any) {
	if err := j.validateSetRetireOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retireOnDelete",
		val,
	)
}

func (j *jsiiProxy_KmsGrant) SetRetiringPrincipal(val *string) {
	if err := j.validateSetRetiringPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retiringPrincipal",
		val,
	)
}

// Generates CDKTF code for importing a KmsGrant resource upon running "cdktf plan <stack-name>".
func KmsGrant_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKmsGrant_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
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
func KmsGrant_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateKmsGrant_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmsGrant_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateKmsGrant_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmsGrant_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateKmsGrant_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func KmsGrant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.kmsGrant.KmsGrant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmsGrant) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (k *jsiiProxy_KmsGrant) AddOverride(path *string, value any) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]any{path, value},
	)
}

func (k *jsiiProxy_KmsGrant) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]any{id, provider},
	)
}

func (k *jsiiProxy_KmsGrant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]any{id},
	)
}

func (k *jsiiProxy_KmsGrant) MoveTo(moveTarget *string, index any) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (k *jsiiProxy_KmsGrant) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]any{id},
	)
}

func (k *jsiiProxy_KmsGrant) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (k *jsiiProxy_KmsGrant) PutConstraints(value any) {
	if err := k.validatePutConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putConstraints",
		[]any{value},
	)
}

func (k *jsiiProxy_KmsGrant) ResetConstraints() {
	_jsii_.InvokeVoid(
		k,
		"resetConstraints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetGrantCreationTokens() {
	_jsii_.InvokeVoid(
		k,
		"resetGrantCreationTokens",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetName() {
	_jsii_.InvokeVoid(
		k,
		"resetName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetRetireOnDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetRetireOnDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) ResetRetiringPrincipal() {
	_jsii_.InvokeVoid(
		k,
		"resetRetiringPrincipal",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsGrant) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsGrant) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
