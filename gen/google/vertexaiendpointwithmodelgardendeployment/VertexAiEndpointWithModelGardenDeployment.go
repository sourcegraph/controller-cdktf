package vertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaiendpointwithmodelgardendeployment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment}.
type VertexAiEndpointWithModelGardenDeployment interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployConfig() VertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference
	DeployConfigInput() *VertexAiEndpointWithModelGardenDeploymentDeployConfig
	DeployedModelDisplayName() *string
	DeployedModelId() *string
	Endpoint() *string
	EndpointConfig() VertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference
	EndpointConfigInput() *VertexAiEndpointWithModelGardenDeploymentEndpointConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HuggingFaceModelId() *string
	SetHuggingFaceModelId(val *string)
	HuggingFaceModelIdInput() *string
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
	ModelConfig() VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
	ModelConfigInput() *VertexAiEndpointWithModelGardenDeploymentModelConfig
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
	PublisherModelName() *string
	SetPublisherModelName(val *string)
	PublisherModelNameInput() *string
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference
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
	PutDeployConfig(value *VertexAiEndpointWithModelGardenDeploymentDeployConfig)
	PutEndpointConfig(value *VertexAiEndpointWithModelGardenDeploymentEndpointConfig)
	PutModelConfig(value *VertexAiEndpointWithModelGardenDeploymentModelConfig)
	PutTimeouts(value *VertexAiEndpointWithModelGardenDeploymentTimeouts)
	ResetDeployConfig()
	ResetEndpointConfig()
	ResetHuggingFaceModelId()
	ResetId()
	ResetModelConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPublisherModelName()
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

// The jsii proxy struct for VertexAiEndpointWithModelGardenDeployment
type jsiiProxy_VertexAiEndpointWithModelGardenDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) DeployConfig() VertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference
	_jsii_.Get(
		j,
		"deployConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) DeployConfigInput() *VertexAiEndpointWithModelGardenDeploymentDeployConfig {
	var returns *VertexAiEndpointWithModelGardenDeploymentDeployConfig
	_jsii_.Get(
		j,
		"deployConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) DeployedModelDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedModelDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) DeployedModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) EndpointConfig() VertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) EndpointConfigInput() *VertexAiEndpointWithModelGardenDeploymentEndpointConfig {
	var returns *VertexAiEndpointWithModelGardenDeploymentEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) HuggingFaceModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) HuggingFaceModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceModelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ModelConfig() VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
	_jsii_.Get(
		j,
		"modelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ModelConfigInput() *VertexAiEndpointWithModelGardenDeploymentModelConfig {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfig
	_jsii_.Get(
		j,
		"modelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PublisherModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PublisherModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) Timeouts() VertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment} Resource.
func NewVertexAiEndpointWithModelGardenDeployment(scope constructs.Construct, id *string, config *VertexAiEndpointWithModelGardenDeploymentConfig) VertexAiEndpointWithModelGardenDeployment {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointWithModelGardenDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointWithModelGardenDeployment{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment} Resource.
func NewVertexAiEndpointWithModelGardenDeployment_Override(v VertexAiEndpointWithModelGardenDeployment, scope constructs.Construct, id *string, config *VertexAiEndpointWithModelGardenDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		[]any{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetHuggingFaceModelId(val *string) {
	if err := j.validateSetHuggingFaceModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceModelId",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SetPublisherModelName(val *string) {
	if err := j.validateSetPublisherModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisherModelName",
		val,
	)
}

// Generates CDKTF code for importing a VertexAiEndpointWithModelGardenDeployment resource upon running "cdktf plan <stack-name>".
func VertexAiEndpointWithModelGardenDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVertexAiEndpointWithModelGardenDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
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
func VertexAiEndpointWithModelGardenDeployment_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateVertexAiEndpointWithModelGardenDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiEndpointWithModelGardenDeployment_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateVertexAiEndpointWithModelGardenDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiEndpointWithModelGardenDeployment_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateVertexAiEndpointWithModelGardenDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func VertexAiEndpointWithModelGardenDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) AddOverride(path *string, value any) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]any{path, value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]any{id, provider},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]any{id},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) MoveTo(moveTarget *string, index any) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]any{id},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PutDeployConfig(value *VertexAiEndpointWithModelGardenDeploymentDeployConfig) {
	if err := v.validatePutDeployConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDeployConfig",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PutEndpointConfig(value *VertexAiEndpointWithModelGardenDeploymentEndpointConfig) {
	if err := v.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEndpointConfig",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PutModelConfig(value *VertexAiEndpointWithModelGardenDeploymentModelConfig) {
	if err := v.validatePutModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putModelConfig",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) PutTimeouts(value *VertexAiEndpointWithModelGardenDeploymentTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetDeployConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetDeployConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetHuggingFaceModelId() {
	_jsii_.InvokeVoid(
		v,
		"resetHuggingFaceModelId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetModelConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetModelConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetProject() {
	_jsii_.InvokeVoid(
		v,
		"resetProject",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetPublisherModelName() {
	_jsii_.InvokeVoid(
		v,
		"resetPublisherModelName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeployment) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
