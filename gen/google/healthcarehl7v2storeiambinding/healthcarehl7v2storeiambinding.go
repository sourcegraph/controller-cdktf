package healthcarehl7v2storeiambinding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/healthcarehl7v2storeiambinding/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_hl7_v2_store_iam_binding google_healthcare_hl7_v2_store_iam_binding}.
type HealthcareHl7V2StoreIamBinding interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() HealthcareHl7V2StoreIamBindingConditionOutputReference
	ConditionInput() *HealthcareHl7V2StoreIamBindingCondition
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
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hl7V2StoreId() *string
	SetHl7V2StoreId(val *string)
	Hl7V2StoreIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
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
	PutCondition(value *HealthcareHl7V2StoreIamBindingCondition)
	ResetCondition()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for HealthcareHl7V2StoreIamBinding
type jsiiProxy_HealthcareHl7V2StoreIamBinding struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Condition() HealthcareHl7V2StoreIamBindingConditionOutputReference {
	var returns HealthcareHl7V2StoreIamBindingConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) ConditionInput() *HealthcareHl7V2StoreIamBindingCondition {
	var returns *HealthcareHl7V2StoreIamBindingCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Hl7V2StoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hl7V2StoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Hl7V2StoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hl7V2StoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_hl7_v2_store_iam_binding google_healthcare_hl7_v2_store_iam_binding} Resource.
func NewHealthcareHl7V2StoreIamBinding(scope constructs.Construct, id *string, config *HealthcareHl7V2StoreIamBindingConfig) HealthcareHl7V2StoreIamBinding {
	_init_.Initialize()

	if err := validateNewHealthcareHl7V2StoreIamBindingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareHl7V2StoreIamBinding{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_hl7_v2_store_iam_binding google_healthcare_hl7_v2_store_iam_binding} Resource.
func NewHealthcareHl7V2StoreIamBinding_Override(h HealthcareHl7V2StoreIamBinding, scope constructs.Construct, id *string, config *HealthcareHl7V2StoreIamBindingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		[]any{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetHl7V2StoreId(val *string) {
	if err := j.validateSetHl7V2StoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hl7V2StoreId",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HealthcareHl7V2StoreIamBinding) SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

// Generates CDKTF code for importing a HealthcareHl7V2StoreIamBinding resource upon running "cdktf plan <stack-name>".
func HealthcareHl7V2StoreIamBinding_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHealthcareHl7V2StoreIamBinding_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
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
func HealthcareHl7V2StoreIamBinding_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareHl7V2StoreIamBinding_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareHl7V2StoreIamBinding_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareHl7V2StoreIamBinding_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareHl7V2StoreIamBinding_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareHl7V2StoreIamBinding_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func HealthcareHl7V2StoreIamBinding_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.healthcareHl7V2StoreIamBinding.HealthcareHl7V2StoreIamBinding",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) AddOverride(path *string, value any) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]any{path, value},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]any{id, provider},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]any{id},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) MoveTo(moveTarget *string, index any) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]any{id},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) PutCondition(value *HealthcareHl7V2StoreIamBindingCondition) {
	if err := h.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putCondition",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ResetCondition() {
	_jsii_.InvokeVoid(
		h,
		"resetCondition",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareHl7V2StoreIamBinding) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
