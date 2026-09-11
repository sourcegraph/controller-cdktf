package googlevertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpointwithmodelgardendeployment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment}.
type GoogleVertexAiEndpointWithModelGardenDeployment interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DeployConfig() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference
	DeployConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig
	DeployedModelDisplayName() *string
	DeployedModelId() *string
	Endpoint() *string
	EndpointConfig() GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference
	EndpointConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfig
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ModelConfig() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
	ModelConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublisherModelName() *string
	SetPublisherModelName(val *string)
	PublisherModelNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleVertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutDeployConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig)
	PutEndpointConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfig)
	PutModelConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig)
	PutTimeouts(value *GoogleVertexAiEndpointWithModelGardenDeploymentTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GoogleVertexAiEndpointWithModelGardenDeployment
type jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) DeployConfig() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference
	_jsii_.Get(
		j,
		"deployConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) DeployConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig
	_jsii_.Get(
		j,
		"deployConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) DeployedModelDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedModelDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) DeployedModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployedModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) EndpointConfig() GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) EndpointConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfig {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) HuggingFaceModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) HuggingFaceModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceModelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ModelConfig() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
	_jsii_.Get(
		j,
		"modelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ModelConfigInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig
	_jsii_.Get(
		j,
		"modelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PublisherModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PublisherModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) Timeouts() GoogleVertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment} Resource.
func NewGoogleVertexAiEndpointWithModelGardenDeployment(scope constructs.Construct, id *string, config *GoogleVertexAiEndpointWithModelGardenDeploymentConfig) GoogleVertexAiEndpointWithModelGardenDeployment {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointWithModelGardenDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment google_vertex_ai_endpoint_with_model_garden_deployment} Resource.
func NewGoogleVertexAiEndpointWithModelGardenDeployment_Override(g GoogleVertexAiEndpointWithModelGardenDeployment, scope constructs.Construct, id *string, config *GoogleVertexAiEndpointWithModelGardenDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetHuggingFaceModelId(val *string) {
	if err := j.validateSetHuggingFaceModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceModelId",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment)SetPublisherModelName(val *string) {
	if err := j.validateSetPublisherModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisherModelName",
		val,
	)
}

// Generates CDKTN code for importing a GoogleVertexAiEndpointWithModelGardenDeployment resource upon running "cdktn plan <stack-name>".
func GoogleVertexAiEndpointWithModelGardenDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleVertexAiEndpointWithModelGardenDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
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
func GoogleVertexAiEndpointWithModelGardenDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiEndpointWithModelGardenDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiEndpointWithModelGardenDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiEndpointWithModelGardenDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiEndpointWithModelGardenDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiEndpointWithModelGardenDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVertexAiEndpointWithModelGardenDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PutDeployConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig) {
	if err := g.validatePutDeployConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeployConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PutEndpointConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfig) {
	if err := g.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PutModelConfig(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig) {
	if err := g.validatePutModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) PutTimeouts(value *GoogleVertexAiEndpointWithModelGardenDeploymentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetDeployConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDeployConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetHuggingFaceModelId() {
	_jsii_.InvokeVoid(
		g,
		"resetHuggingFaceModelId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetModelConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetModelConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetPublisherModelName() {
	_jsii_.InvokeVoid(
		g,
		"resetPublisherModelName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeployment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

