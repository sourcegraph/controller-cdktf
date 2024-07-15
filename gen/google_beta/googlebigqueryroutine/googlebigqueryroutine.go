package googlebigqueryroutine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryroutine/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine google_bigquery_routine}.
type GoogleBigqueryRoutine interface {
	cdktf.TerraformResource
	Arguments() GoogleBigqueryRoutineArgumentsList
	ArgumentsInput() interface{}
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
	CreationTime() *float64
	DataGovernanceType() *string
	SetDataGovernanceType(val *string)
	DataGovernanceTypeInput() *string
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DefinitionBody() *string
	SetDefinitionBody(val *string)
	DefinitionBodyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeterminismLevel() *string
	SetDeterminismLevel(val *string)
	DeterminismLevelInput() *string
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
	ImportedLibraries() *[]*string
	SetImportedLibraries(val *[]*string)
	ImportedLibrariesInput() *[]*string
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteFunctionOptions() GoogleBigqueryRoutineRemoteFunctionOptionsOutputReference
	RemoteFunctionOptionsInput() *GoogleBigqueryRoutineRemoteFunctionOptions
	ReturnTableType() *string
	SetReturnTableType(val *string)
	ReturnTableTypeInput() *string
	ReturnType() *string
	SetReturnType(val *string)
	ReturnTypeInput() *string
	RoutineId() *string
	SetRoutineId(val *string)
	RoutineIdInput() *string
	RoutineType() *string
	SetRoutineType(val *string)
	RoutineTypeInput() *string
	SparkOptions() GoogleBigqueryRoutineSparkOptionsOutputReference
	SparkOptionsInput() *GoogleBigqueryRoutineSparkOptions
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryRoutineTimeoutsOutputReference
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
	PutArguments(value interface{})
	PutRemoteFunctionOptions(value *GoogleBigqueryRoutineRemoteFunctionOptions)
	PutSparkOptions(value *GoogleBigqueryRoutineSparkOptions)
	PutTimeouts(value *GoogleBigqueryRoutineTimeouts)
	ResetArguments()
	ResetDataGovernanceType()
	ResetDescription()
	ResetDeterminismLevel()
	ResetId()
	ResetImportedLibraries()
	ResetLanguage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteFunctionOptions()
	ResetReturnTableType()
	ResetReturnType()
	ResetSparkOptions()
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

// The jsii proxy struct for GoogleBigqueryRoutine
type jsiiProxy_GoogleBigqueryRoutine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Arguments() GoogleBigqueryRoutineArgumentsList {
	var returns GoogleBigqueryRoutineArgumentsList
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DataGovernanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataGovernanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DataGovernanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataGovernanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DefinitionBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DefinitionBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DeterminismLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"determinismLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) DeterminismLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"determinismLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ImportedLibraries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importedLibraries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ImportedLibrariesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importedLibrariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RemoteFunctionOptions() GoogleBigqueryRoutineRemoteFunctionOptionsOutputReference {
	var returns GoogleBigqueryRoutineRemoteFunctionOptionsOutputReference
	_jsii_.Get(
		j,
		"remoteFunctionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RemoteFunctionOptionsInput() *GoogleBigqueryRoutineRemoteFunctionOptions {
	var returns *GoogleBigqueryRoutineRemoteFunctionOptions
	_jsii_.Get(
		j,
		"remoteFunctionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ReturnTableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ReturnTableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ReturnType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) ReturnTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"returnTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RoutineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RoutineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RoutineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) RoutineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) SparkOptions() GoogleBigqueryRoutineSparkOptionsOutputReference {
	var returns GoogleBigqueryRoutineSparkOptionsOutputReference
	_jsii_.Get(
		j,
		"sparkOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) SparkOptionsInput() *GoogleBigqueryRoutineSparkOptions {
	var returns *GoogleBigqueryRoutineSparkOptions
	_jsii_.Get(
		j,
		"sparkOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) Timeouts() GoogleBigqueryRoutineTimeoutsOutputReference {
	var returns GoogleBigqueryRoutineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine google_bigquery_routine} Resource.
func NewGoogleBigqueryRoutine(scope constructs.Construct, id *string, config *GoogleBigqueryRoutineConfig) GoogleBigqueryRoutine {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryRoutineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryRoutine{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine google_bigquery_routine} Resource.
func NewGoogleBigqueryRoutine_Override(g GoogleBigqueryRoutine, scope constructs.Construct, id *string, config *GoogleBigqueryRoutineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDataGovernanceType(val *string) {
	if err := j.validateSetDataGovernanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataGovernanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDefinitionBody(val *string) {
	if err := j.validateSetDefinitionBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetDeterminismLevel(val *string) {
	if err := j.validateSetDeterminismLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"determinismLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetImportedLibraries(val *[]*string) {
	if err := j.validateSetImportedLibrariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedLibraries",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetReturnTableType(val *string) {
	if err := j.validateSetReturnTableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnTableType",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetReturnType(val *string) {
	if err := j.validateSetReturnTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnType",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetRoutineId(val *string) {
	if err := j.validateSetRoutineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutine)SetRoutineType(val *string) {
	if err := j.validateSetRoutineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineType",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigqueryRoutine resource upon running "cdktf plan <stack-name>".
func GoogleBigqueryRoutine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryRoutine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
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
func GoogleBigqueryRoutine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryRoutine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryRoutine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryRoutine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryRoutine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryRoutine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryRoutine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryRoutine) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) PutArguments(value interface{}) {
	if err := g.validatePutArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArguments",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) PutRemoteFunctionOptions(value *GoogleBigqueryRoutineRemoteFunctionOptions) {
	if err := g.validatePutRemoteFunctionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRemoteFunctionOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) PutSparkOptions(value *GoogleBigqueryRoutineSparkOptions) {
	if err := g.validatePutSparkOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSparkOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) PutTimeouts(value *GoogleBigqueryRoutineTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetArguments() {
	_jsii_.InvokeVoid(
		g,
		"resetArguments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetDataGovernanceType() {
	_jsii_.InvokeVoid(
		g,
		"resetDataGovernanceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetDeterminismLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetDeterminismLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetImportedLibraries() {
	_jsii_.InvokeVoid(
		g,
		"resetImportedLibraries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetLanguage() {
	_jsii_.InvokeVoid(
		g,
		"resetLanguage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetRemoteFunctionOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoteFunctionOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetReturnTableType() {
	_jsii_.InvokeVoid(
		g,
		"resetReturnTableType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetReturnType() {
	_jsii_.InvokeVoid(
		g,
		"resetReturnType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetSparkOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

