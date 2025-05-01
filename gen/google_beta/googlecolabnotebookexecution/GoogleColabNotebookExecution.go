package googlecolabnotebookexecution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecolabnotebookexecution/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_notebook_execution google_colab_notebook_execution}.
type GoogleColabNotebookExecution interface {
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
	DataformRepositorySource() GoogleColabNotebookExecutionDataformRepositorySourceOutputReference
	DataformRepositorySourceInput() *GoogleColabNotebookExecutionDataformRepositorySource
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectNotebookSource() GoogleColabNotebookExecutionDirectNotebookSourceOutputReference
	DirectNotebookSourceInput() *GoogleColabNotebookExecutionDirectNotebookSource
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
	ExecutionUser() *string
	SetExecutionUser(val *string)
	ExecutionUserInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsNotebookSource() GoogleColabNotebookExecutionGcsNotebookSourceOutputReference
	GcsNotebookSourceInput() *GoogleColabNotebookExecutionGcsNotebookSource
	GcsOutputUri() *string
	SetGcsOutputUri(val *string)
	GcsOutputUriInput() *string
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
	// The tree node.
	Node() constructs.Node
	NotebookExecutionJobId() *string
	SetNotebookExecutionJobId(val *string)
	NotebookExecutionJobIdInput() *string
	NotebookRuntimeTemplateResourceName() *string
	SetNotebookRuntimeTemplateResourceName(val *string)
	NotebookRuntimeTemplateResourceNameInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleColabNotebookExecutionTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDataformRepositorySource(value *GoogleColabNotebookExecutionDataformRepositorySource)
	PutDirectNotebookSource(value *GoogleColabNotebookExecutionDirectNotebookSource)
	PutGcsNotebookSource(value *GoogleColabNotebookExecutionGcsNotebookSource)
	PutTimeouts(value *GoogleColabNotebookExecutionTimeouts)
	ResetDataformRepositorySource()
	ResetDirectNotebookSource()
	ResetExecutionTimeout()
	ResetExecutionUser()
	ResetGcsNotebookSource()
	ResetId()
	ResetNotebookExecutionJobId()
	ResetNotebookRuntimeTemplateResourceName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetServiceAccount()
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
}

// The jsii proxy struct for GoogleColabNotebookExecution
type jsiiProxy_GoogleColabNotebookExecution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleColabNotebookExecution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DataformRepositorySource() GoogleColabNotebookExecutionDataformRepositorySourceOutputReference {
	var returns GoogleColabNotebookExecutionDataformRepositorySourceOutputReference
	_jsii_.Get(
		j,
		"dataformRepositorySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DataformRepositorySourceInput() *GoogleColabNotebookExecutionDataformRepositorySource {
	var returns *GoogleColabNotebookExecutionDataformRepositorySource
	_jsii_.Get(
		j,
		"dataformRepositorySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DirectNotebookSource() GoogleColabNotebookExecutionDirectNotebookSourceOutputReference {
	var returns GoogleColabNotebookExecutionDirectNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"directNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DirectNotebookSourceInput() *GoogleColabNotebookExecutionDirectNotebookSource {
	var returns *GoogleColabNotebookExecutionDirectNotebookSource
	_jsii_.Get(
		j,
		"directNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ExecutionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ExecutionUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) GcsNotebookSource() GoogleColabNotebookExecutionGcsNotebookSourceOutputReference {
	var returns GoogleColabNotebookExecutionGcsNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"gcsNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) GcsNotebookSourceInput() *GoogleColabNotebookExecutionGcsNotebookSource {
	var returns *GoogleColabNotebookExecutionGcsNotebookSource
	_jsii_.Get(
		j,
		"gcsNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) GcsOutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) GcsOutputUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) NotebookExecutionJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookExecutionJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) NotebookExecutionJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookExecutionJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) NotebookRuntimeTemplateResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) NotebookRuntimeTemplateResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) Timeouts() GoogleColabNotebookExecutionTimeoutsOutputReference {
	var returns GoogleColabNotebookExecutionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecution) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_notebook_execution google_colab_notebook_execution} Resource.
func NewGoogleColabNotebookExecution(scope constructs.Construct, id *string, config *GoogleColabNotebookExecutionConfig) GoogleColabNotebookExecution {
	_init_.Initialize()

	if err := validateNewGoogleColabNotebookExecutionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleColabNotebookExecution{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_notebook_execution google_colab_notebook_execution} Resource.
func NewGoogleColabNotebookExecution_Override(g GoogleColabNotebookExecution, scope constructs.Construct, id *string, config *GoogleColabNotebookExecutionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetExecutionTimeout(val *string) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetExecutionUser(val *string) {
	if err := j.validateSetExecutionUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUser",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetGcsOutputUri(val *string) {
	if err := j.validateSetGcsOutputUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputUri",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetNotebookExecutionJobId(val *string) {
	if err := j.validateSetNotebookExecutionJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookExecutionJobId",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetNotebookRuntimeTemplateResourceName(val *string) {
	if err := j.validateSetNotebookRuntimeTemplateResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookRuntimeTemplateResourceName",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecution)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

// Generates CDKTF code for importing a GoogleColabNotebookExecution resource upon running "cdktf plan <stack-name>".
func GoogleColabNotebookExecution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleColabNotebookExecution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
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
func GoogleColabNotebookExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabNotebookExecution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleColabNotebookExecution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabNotebookExecution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleColabNotebookExecution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabNotebookExecution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleColabNotebookExecution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleColabNotebookExecution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) PutDataformRepositorySource(value *GoogleColabNotebookExecutionDataformRepositorySource) {
	if err := g.validatePutDataformRepositorySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataformRepositorySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) PutDirectNotebookSource(value *GoogleColabNotebookExecutionDirectNotebookSource) {
	if err := g.validatePutDirectNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDirectNotebookSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) PutGcsNotebookSource(value *GoogleColabNotebookExecutionGcsNotebookSource) {
	if err := g.validatePutGcsNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsNotebookSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) PutTimeouts(value *GoogleColabNotebookExecutionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetDataformRepositorySource() {
	_jsii_.InvokeVoid(
		g,
		"resetDataformRepositorySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetDirectNotebookSource() {
	_jsii_.InvokeVoid(
		g,
		"resetDirectNotebookSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetExecutionUser() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetGcsNotebookSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsNotebookSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetNotebookExecutionJobId() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebookExecutionJobId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetNotebookRuntimeTemplateResourceName() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebookRuntimeTemplateResourceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

