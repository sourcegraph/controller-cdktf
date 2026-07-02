package healthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/healthcarefhirstore/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store}.
type HealthcareFhirStore interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComplexDataTypeReferenceParsing() *string
	SetComplexDataTypeReferenceParsing(val *string)
	ComplexDataTypeReferenceParsingInput() *string
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
	Dataset() *string
	SetDataset(val *string)
	DatasetInput() *string
	DefaultSearchHandlingStrict() any
	SetDefaultSearchHandlingStrict(val any)
	DefaultSearchHandlingStrictInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableReferentialIntegrity() any
	SetDisableReferentialIntegrity(val any)
	DisableReferentialIntegrityInput() any
	DisableResourceVersioning() any
	SetDisableResourceVersioning(val any)
	DisableResourceVersioningInput() any
	EffectiveLabels() cdktf.StringMap
	EnableHistoryImport() any
	SetEnableHistoryImport(val any)
	EnableHistoryImportInput() any
	EnableUpdateCreate() any
	SetEnableUpdateCreate(val any)
	EnableUpdateCreateInput() any
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NotificationConfig() HealthcareFhirStoreNotificationConfigOutputReference
	NotificationConfigInput() *HealthcareFhirStoreNotificationConfig
	NotificationConfigs() HealthcareFhirStoreNotificationConfigsList
	NotificationConfigsInput() any
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
	SelfLink() *string
	StreamConfigs() HealthcareFhirStoreStreamConfigsList
	StreamConfigsInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HealthcareFhirStoreTimeoutsOutputReference
	TimeoutsInput() any
	ValidationConfig() HealthcareFhirStoreValidationConfigOutputReference
	ValidationConfigInput() *HealthcareFhirStoreValidationConfig
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutNotificationConfig(value *HealthcareFhirStoreNotificationConfig)
	PutNotificationConfigs(value any)
	PutStreamConfigs(value any)
	PutTimeouts(value *HealthcareFhirStoreTimeouts)
	PutValidationConfig(value *HealthcareFhirStoreValidationConfig)
	ResetComplexDataTypeReferenceParsing()
	ResetDefaultSearchHandlingStrict()
	ResetDisableReferentialIntegrity()
	ResetDisableResourceVersioning()
	ResetEnableHistoryImport()
	ResetEnableUpdateCreate()
	ResetId()
	ResetLabels()
	ResetNotificationConfig()
	ResetNotificationConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamConfigs()
	ResetTimeouts()
	ResetValidationConfig()
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

// The jsii proxy struct for HealthcareFhirStore
type jsiiProxy_HealthcareFhirStore struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HealthcareFhirStore) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ComplexDataTypeReferenceParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ComplexDataTypeReferenceParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexDataTypeReferenceParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DefaultSearchHandlingStrict() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultSearchHandlingStrict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DefaultSearchHandlingStrictInput() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultSearchHandlingStrictInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableReferentialIntegrity() any {
	var returns any
	_jsii_.Get(
		j,
		"disableReferentialIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableReferentialIntegrityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableReferentialIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableResourceVersioning() any {
	var returns any
	_jsii_.Get(
		j,
		"disableResourceVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) DisableResourceVersioningInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableResourceVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableHistoryImport() any {
	var returns any
	_jsii_.Get(
		j,
		"enableHistoryImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableHistoryImportInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableHistoryImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableUpdateCreate() any {
	var returns any
	_jsii_.Get(
		j,
		"enableUpdateCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) EnableUpdateCreateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableUpdateCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfig() HealthcareFhirStoreNotificationConfigOutputReference {
	var returns HealthcareFhirStoreNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigInput() *HealthcareFhirStoreNotificationConfig {
	var returns *HealthcareFhirStoreNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigs() HealthcareFhirStoreNotificationConfigsList {
	var returns HealthcareFhirStoreNotificationConfigsList
	_jsii_.Get(
		j,
		"notificationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) NotificationConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"notificationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) StreamConfigs() HealthcareFhirStoreStreamConfigsList {
	var returns HealthcareFhirStoreStreamConfigsList
	_jsii_.Get(
		j,
		"streamConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) StreamConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"streamConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Timeouts() HealthcareFhirStoreTimeoutsOutputReference {
	var returns HealthcareFhirStoreTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ValidationConfig() HealthcareFhirStoreValidationConfigOutputReference {
	var returns HealthcareFhirStoreValidationConfigOutputReference
	_jsii_.Get(
		j,
		"validationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) ValidationConfigInput() *HealthcareFhirStoreValidationConfig {
	var returns *HealthcareFhirStoreValidationConfig
	_jsii_.Get(
		j,
		"validationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStore) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewHealthcareFhirStore(scope constructs.Construct, id *string, config *HealthcareFhirStoreConfig) HealthcareFhirStore {
	_init_.Initialize()

	if err := validateNewHealthcareFhirStoreParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareFhirStore{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/healthcare_fhir_store google_healthcare_fhir_store} Resource.
func NewHealthcareFhirStore_Override(h HealthcareFhirStore, scope constructs.Construct, id *string, config *HealthcareFhirStoreConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		[]any{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetComplexDataTypeReferenceParsing(val *string) {
	if err := j.validateSetComplexDataTypeReferenceParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexDataTypeReferenceParsing",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetDefaultSearchHandlingStrict(val any) {
	if err := j.validateSetDefaultSearchHandlingStrictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSearchHandlingStrict",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetDisableReferentialIntegrity(val any) {
	if err := j.validateSetDisableReferentialIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableReferentialIntegrity",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetDisableResourceVersioning(val any) {
	if err := j.validateSetDisableResourceVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableResourceVersioning",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetEnableHistoryImport(val any) {
	if err := j.validateSetEnableHistoryImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHistoryImport",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetEnableUpdateCreate(val any) {
	if err := j.validateSetEnableUpdateCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUpdateCreate",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStore) SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a HealthcareFhirStore resource upon running "cdktf plan <stack-name>".
func HealthcareFhirStore_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
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
func HealthcareFhirStore_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareFhirStore_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareFhirStore_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateHealthcareFhirStore_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func HealthcareFhirStore_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStore",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) AddOverride(path *string, value any) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]any{path, value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (h *jsiiProxy_HealthcareFhirStore) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareFhirStore) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcareFhirStore) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcareFhirStore) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcareFhirStore) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcareFhirStore) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcareFhirStore) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcareFhirStore) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcareFhirStore) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]any{id, provider},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareFhirStore) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]any{id},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) MoveTo(moveTarget *string, index any) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]any{id},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutNotificationConfig(value *HealthcareFhirStoreNotificationConfig) {
	if err := h.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNotificationConfig",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutNotificationConfigs(value any) {
	if err := h.validatePutNotificationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNotificationConfigs",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutStreamConfigs(value any) {
	if err := h.validatePutStreamConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStreamConfigs",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutTimeouts(value *HealthcareFhirStoreTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) PutValidationConfig(value *HealthcareFhirStoreValidationConfig) {
	if err := h.validatePutValidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putValidationConfig",
		[]any{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetComplexDataTypeReferenceParsing() {
	_jsii_.InvokeVoid(
		h,
		"resetComplexDataTypeReferenceParsing",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDefaultSearchHandlingStrict() {
	_jsii_.InvokeVoid(
		h,
		"resetDefaultSearchHandlingStrict",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDisableReferentialIntegrity() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableReferentialIntegrity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetDisableResourceVersioning() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableResourceVersioning",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetEnableHistoryImport() {
	_jsii_.InvokeVoid(
		h,
		"resetEnableHistoryImport",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetEnableUpdateCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetEnableUpdateCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetLabels() {
	_jsii_.InvokeVoid(
		h,
		"resetLabels",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		h,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetNotificationConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetNotificationConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetStreamConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetStreamConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) ResetValidationConfig() {
	_jsii_.InvokeVoid(
		h,
		"resetValidationConfig",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStore) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStore) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
