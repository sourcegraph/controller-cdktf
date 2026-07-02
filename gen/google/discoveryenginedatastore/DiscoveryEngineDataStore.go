package discoveryenginedatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginedatastore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_store google_discovery_engine_data_store}.
type DiscoveryEngineDataStore interface {
	cdktf.TerraformResource
	AdvancedSiteSearchConfig() DiscoveryEngineDataStoreAdvancedSiteSearchConfigOutputReference
	AdvancedSiteSearchConfigInput() *DiscoveryEngineDataStoreAdvancedSiteSearchConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContentConfig() *string
	SetContentConfig(val *string)
	ContentConfigInput() *string
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateAdvancedSiteSearch() any
	SetCreateAdvancedSiteSearch(val any)
	CreateAdvancedSiteSearchInput() any
	CreateTime() *string
	DataStoreId() *string
	SetDataStoreId(val *string)
	DataStoreIdInput() *string
	DefaultSchemaId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DocumentProcessingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference
	DocumentProcessingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfig
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
	IndustryVertical() *string
	SetIndustryVertical(val *string)
	IndustryVerticalInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
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
	SkipDefaultSchemaCreation() any
	SetSkipDefaultSchemaCreation(val any)
	SkipDefaultSchemaCreationInput() any
	SolutionTypes() *[]*string
	SetSolutionTypes(val *[]*string)
	SolutionTypesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DiscoveryEngineDataStoreTimeoutsOutputReference
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
	PutAdvancedSiteSearchConfig(value *DiscoveryEngineDataStoreAdvancedSiteSearchConfig)
	PutDocumentProcessingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfig)
	PutTimeouts(value *DiscoveryEngineDataStoreTimeouts)
	ResetAdvancedSiteSearchConfig()
	ResetContentConfig()
	ResetCreateAdvancedSiteSearch()
	ResetDocumentProcessingConfig()
	ResetId()
	ResetKmsKeyName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSkipDefaultSchemaCreation()
	ResetSolutionTypes()
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

// The jsii proxy struct for DiscoveryEngineDataStore
type jsiiProxy_DiscoveryEngineDataStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineDataStore) AdvancedSiteSearchConfig() DiscoveryEngineDataStoreAdvancedSiteSearchConfigOutputReference {
	var returns DiscoveryEngineDataStoreAdvancedSiteSearchConfigOutputReference
	_jsii_.Get(
		j,
		"advancedSiteSearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) AdvancedSiteSearchConfigInput() *DiscoveryEngineDataStoreAdvancedSiteSearchConfig {
	var returns *DiscoveryEngineDataStoreAdvancedSiteSearchConfig
	_jsii_.Get(
		j,
		"advancedSiteSearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) ContentConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) ContentConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) CreateAdvancedSiteSearch() any {
	var returns any
	_jsii_.Get(
		j,
		"createAdvancedSiteSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) CreateAdvancedSiteSearchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"createAdvancedSiteSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DataStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DataStoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DefaultSchemaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSchemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DocumentProcessingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference
	_jsii_.Get(
		j,
		"documentProcessingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) DocumentProcessingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfig
	_jsii_.Get(
		j,
		"documentProcessingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) IndustryVertical() *string {
	var returns *string
	_jsii_.Get(
		j,
		"industryVertical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) IndustryVerticalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"industryVerticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SkipDefaultSchemaCreation() any {
	var returns any
	_jsii_.Get(
		j,
		"skipDefaultSchemaCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SkipDefaultSchemaCreationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"skipDefaultSchemaCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SolutionTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"solutionTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SolutionTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"solutionTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) Timeouts() DiscoveryEngineDataStoreTimeoutsOutputReference {
	var returns DiscoveryEngineDataStoreTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStore) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_store google_discovery_engine_data_store} Resource.
func NewDiscoveryEngineDataStore(scope constructs.Construct, id *string, config *DiscoveryEngineDataStoreConfig) DiscoveryEngineDataStore {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineDataStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineDataStore{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_store google_discovery_engine_data_store} Resource.
func NewDiscoveryEngineDataStore_Override(d DiscoveryEngineDataStore, scope constructs.Construct, id *string, config *DiscoveryEngineDataStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetContentConfig(val *string) {
	if err := j.validateSetContentConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentConfig",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetCreateAdvancedSiteSearch(val any) {
	if err := j.validateSetCreateAdvancedSiteSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAdvancedSiteSearch",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetDataStoreId(val *string) {
	if err := j.validateSetDataStoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStoreId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetIndustryVertical(val *string) {
	if err := j.validateSetIndustryVerticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"industryVertical",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetSkipDefaultSchemaCreation(val any) {
	if err := j.validateSetSkipDefaultSchemaCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDefaultSchemaCreation",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStore) SetSolutionTypes(val *[]*string) {
	if err := j.validateSetSolutionTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"solutionTypes",
		val,
	)
}

// Generates CDKTF code for importing a DiscoveryEngineDataStore resource upon running "cdktf plan <stack-name>".
func DiscoveryEngineDataStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
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
func DiscoveryEngineDataStore_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineDataStore_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineDataStore_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineDataStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) PutAdvancedSiteSearchConfig(value *DiscoveryEngineDataStoreAdvancedSiteSearchConfig) {
	if err := d.validatePutAdvancedSiteSearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdvancedSiteSearchConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) PutDocumentProcessingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfig) {
	if err := d.validatePutDocumentProcessingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDocumentProcessingConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) PutTimeouts(value *DiscoveryEngineDataStoreTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetAdvancedSiteSearchConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAdvancedSiteSearchConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetContentConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetContentConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetCreateAdvancedSiteSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateAdvancedSiteSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetDocumentProcessingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDocumentProcessingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetSkipDefaultSchemaCreation() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipDefaultSchemaCreation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetSolutionTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetSolutionTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStore) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStore) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
