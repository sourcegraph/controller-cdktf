package bigquerydatapolicyv2datapolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigquerydatapolicyv2datapolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_datapolicyv2_data_policy google_bigquery_datapolicyv2_data_policy}.
type BigqueryDatapolicyv2DataPolicy interface {
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
	DataMaskingPolicy() BigqueryDatapolicyv2DataPolicyDataMaskingPolicyOutputReference
	DataMaskingPolicyInput() *BigqueryDatapolicyv2DataPolicyDataMaskingPolicy
	DataPolicyId() *string
	SetDataPolicyId(val *string)
	DataPolicyIdInput() *string
	DataPolicyType() *string
	SetDataPolicyType(val *string)
	DataPolicyTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Grantees() *[]*string
	SetGrantees(val *[]*string)
	GranteesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PolicyTag() *string
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
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryDatapolicyv2DataPolicyTimeoutsOutputReference
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
	PutDataMaskingPolicy(value *BigqueryDatapolicyv2DataPolicyDataMaskingPolicy)
	PutTimeouts(value *BigqueryDatapolicyv2DataPolicyTimeouts)
	ResetDataMaskingPolicy()
	ResetEtag()
	ResetGrantees()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for BigqueryDatapolicyv2DataPolicy
type jsiiProxy_BigqueryDatapolicyv2DataPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataMaskingPolicy() BigqueryDatapolicyv2DataPolicyDataMaskingPolicyOutputReference {
	var returns BigqueryDatapolicyv2DataPolicyDataMaskingPolicyOutputReference
	_jsii_.Get(
		j,
		"dataMaskingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataMaskingPolicyInput() *BigqueryDatapolicyv2DataPolicyDataMaskingPolicy {
	var returns *BigqueryDatapolicyv2DataPolicyDataMaskingPolicy
	_jsii_.Get(
		j,
		"dataMaskingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataPolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPolicyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DataPolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPolicyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Grantees() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantees",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GranteesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"granteesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) PolicyTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Timeouts() BigqueryDatapolicyv2DataPolicyTimeoutsOutputReference {
	var returns BigqueryDatapolicyv2DataPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_datapolicyv2_data_policy google_bigquery_datapolicyv2_data_policy} Resource.
func NewBigqueryDatapolicyv2DataPolicy(scope constructs.Construct, id *string, config *BigqueryDatapolicyv2DataPolicyConfig) BigqueryDatapolicyv2DataPolicy {
	_init_.Initialize()

	if err := validateNewBigqueryDatapolicyv2DataPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDatapolicyv2DataPolicy{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_datapolicyv2_data_policy google_bigquery_datapolicyv2_data_policy} Resource.
func NewBigqueryDatapolicyv2DataPolicy_Override(b BigqueryDatapolicyv2DataPolicy, scope constructs.Construct, id *string, config *BigqueryDatapolicyv2DataPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		[]any{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetDataPolicyId(val *string) {
	if err := j.validateSetDataPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPolicyId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetDataPolicyType(val *string) {
	if err := j.validateSetDataPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPolicyType",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetGrantees(val *[]*string) {
	if err := j.validateSetGranteesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantees",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a BigqueryDatapolicyv2DataPolicy resource upon running "cdktf plan <stack-name>".
func BigqueryDatapolicyv2DataPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBigqueryDatapolicyv2DataPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
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
func BigqueryDatapolicyv2DataPolicy_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatapolicyv2DataPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDatapolicyv2DataPolicy_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatapolicyv2DataPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDatapolicyv2DataPolicy_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateBigqueryDatapolicyv2DataPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func BigqueryDatapolicyv2DataPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.bigqueryDatapolicyv2DataPolicy.BigqueryDatapolicyv2DataPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) AddOverride(path *string, value any) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]any{path, value},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]any{id, provider},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]any{id},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) MoveTo(moveTarget *string, index any) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]any{id},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) PutDataMaskingPolicy(value *BigqueryDatapolicyv2DataPolicyDataMaskingPolicy) {
	if err := b.validatePutDataMaskingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDataMaskingPolicy",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) PutTimeouts(value *BigqueryDatapolicyv2DataPolicyTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]any{value},
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetDataMaskingPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetDataMaskingPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetEtag() {
	_jsii_.InvokeVoid(
		b,
		"resetEtag",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetGrantees() {
	_jsii_.InvokeVoid(
		b,
		"resetGrantees",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDatapolicyv2DataPolicy) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
