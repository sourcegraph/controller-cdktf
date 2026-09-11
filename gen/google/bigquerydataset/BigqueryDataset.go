package bigquerydataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigquerydataset/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset}.
type BigqueryDataset interface {
	cdktn.TerraformResource
	Access() BigqueryDatasetAccessList
	AccessInput() interface{}
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
	CreationTime() *float64
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DefaultCollation() *string
	SetDefaultCollation(val *string)
	DefaultCollationInput() *string
	DefaultEncryptionConfiguration() BigqueryDatasetDefaultEncryptionConfigurationOutputReference
	DefaultEncryptionConfigurationInput() *BigqueryDatasetDefaultEncryptionConfiguration
	DefaultPartitionExpirationMs() *float64
	SetDefaultPartitionExpirationMs(val *float64)
	DefaultPartitionExpirationMsInput() *float64
	DefaultTableExpirationMs() *float64
	SetDefaultTableExpirationMs(val *float64)
	DefaultTableExpirationMsInput() *float64
	DeleteContentsOnDestroy() interface{}
	SetDeleteContentsOnDestroy(val interface{})
	DeleteContentsOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktn.StringMap
	Etag() *string
	ExternalCatalogDatasetOptions() BigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	ExternalCatalogDatasetOptionsInput() *BigqueryDatasetExternalCatalogDatasetOptions
	ExternalDatasetReference() BigqueryDatasetExternalDatasetReferenceOutputReference
	ExternalDatasetReferenceInput() *BigqueryDatasetExternalDatasetReference
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	SetFriendlyName(val *string)
	FriendlyNameInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsCaseInsensitive() interface{}
	SetIsCaseInsensitive(val interface{})
	IsCaseInsensitiveInput() interface{}
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxTimeTravelHours() *string
	SetMaxTimeTravelHours(val *string)
	MaxTimeTravelHoursInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	ResourceTags() *map[string]*string
	SetResourceTags(val *map[string]*string)
	ResourceTagsInput() *map[string]*string
	SelfLink() *string
	StorageBillingModel() *string
	SetStorageBillingModel(val *string)
	StorageBillingModelInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryDatasetTimeoutsOutputReference
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
	PutAccess(value interface{})
	PutDefaultEncryptionConfiguration(value *BigqueryDatasetDefaultEncryptionConfiguration)
	PutExternalCatalogDatasetOptions(value *BigqueryDatasetExternalCatalogDatasetOptions)
	PutExternalDatasetReference(value *BigqueryDatasetExternalDatasetReference)
	PutTimeouts(value *BigqueryDatasetTimeouts)
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
	ResetAccess()
	ResetDefaultCollation()
	ResetDefaultEncryptionConfiguration()
	ResetDefaultPartitionExpirationMs()
	ResetDefaultTableExpirationMs()
	ResetDeleteContentsOnDestroy()
	ResetDescription()
	ResetExternalCatalogDatasetOptions()
	ResetExternalDatasetReference()
	ResetFriendlyName()
	ResetId()
	ResetIsCaseInsensitive()
	ResetLabels()
	ResetLocation()
	ResetMaxTimeTravelHours()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResourceTags()
	ResetStorageBillingModel()
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

// The jsii proxy struct for BigqueryDataset
type jsiiProxy_BigqueryDataset struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BigqueryDataset) Access() BigqueryDatasetAccessList {
	var returns BigqueryDatasetAccessList
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) AccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultEncryptionConfiguration() BigqueryDatasetDefaultEncryptionConfigurationOutputReference {
	var returns BigqueryDatasetDefaultEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"defaultEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultEncryptionConfigurationInput() *BigqueryDatasetDefaultEncryptionConfiguration {
	var returns *BigqueryDatasetDefaultEncryptionConfiguration
	_jsii_.Get(
		j,
		"defaultEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultPartitionExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultPartitionExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPartitionExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultTableExpirationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DefaultTableExpirationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTableExpirationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DeleteContentsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteContentsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DeleteContentsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteContentsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalCatalogDatasetOptions() BigqueryDatasetExternalCatalogDatasetOptionsOutputReference {
	var returns BigqueryDatasetExternalCatalogDatasetOptionsOutputReference
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalCatalogDatasetOptionsInput() *BigqueryDatasetExternalCatalogDatasetOptions {
	var returns *BigqueryDatasetExternalCatalogDatasetOptions
	_jsii_.Get(
		j,
		"externalCatalogDatasetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalDatasetReference() BigqueryDatasetExternalDatasetReferenceOutputReference {
	var returns BigqueryDatasetExternalDatasetReferenceOutputReference
	_jsii_.Get(
		j,
		"externalDatasetReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ExternalDatasetReferenceInput() *BigqueryDatasetExternalDatasetReference {
	var returns *BigqueryDatasetExternalDatasetReference
	_jsii_.Get(
		j,
		"externalDatasetReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IsCaseInsensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) IsCaseInsensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaseInsensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) MaxTimeTravelHours() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) MaxTimeTravelHoursInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeTravelHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) StorageBillingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) StorageBillingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBillingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) Timeouts() BigqueryDatasetTimeoutsOutputReference {
	var returns BigqueryDatasetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataset) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset} Resource.
func NewBigqueryDataset(scope constructs.Construct, id *string, config *BigqueryDatasetConfig) BigqueryDataset {
	_init_.Initialize()

	if err := validateNewBigqueryDatasetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDataset{}

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/bigquery_dataset google_bigquery_dataset} Resource.
func NewBigqueryDataset_Override(b BigqueryDataset, scope constructs.Construct, id *string, config *BigqueryDatasetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDefaultCollation(val *string) {
	if err := j.validateSetDefaultCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCollation",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDefaultPartitionExpirationMs(val *float64) {
	if err := j.validateSetDefaultPartitionExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPartitionExpirationMs",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDefaultTableExpirationMs(val *float64) {
	if err := j.validateSetDefaultTableExpirationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTableExpirationMs",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDeleteContentsOnDestroy(val interface{}) {
	if err := j.validateSetDeleteContentsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteContentsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetIsCaseInsensitive(val interface{}) {
	if err := j.validateSetIsCaseInsensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCaseInsensitive",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetMaxTimeTravelHours(val *string) {
	if err := j.validateSetMaxTimeTravelHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeTravelHours",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataset)SetStorageBillingModel(val *string) {
	if err := j.validateSetStorageBillingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBillingModel",
		val,
	)
}

// Generates CDKTN code for importing a BigqueryDataset resource upon running "cdktn plan <stack-name>".
func BigqueryDataset_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBigqueryDataset_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
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
func BigqueryDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataset_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataset_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataset_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BigqueryDataset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.bigqueryDataset.BigqueryDataset",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryDataset) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BigqueryDataset) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BigqueryDataset) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BigqueryDataset) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := b.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryDataset) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BigqueryDataset) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigqueryDataset) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutAccess(value interface{}) {
	if err := b.validatePutAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAccess",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutDefaultEncryptionConfiguration(value *BigqueryDatasetDefaultEncryptionConfiguration) {
	if err := b.validatePutDefaultEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDefaultEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutExternalCatalogDatasetOptions(value *BigqueryDatasetExternalCatalogDatasetOptions) {
	if err := b.validatePutExternalCatalogDatasetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExternalCatalogDatasetOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutExternalDatasetReference(value *BigqueryDatasetExternalDatasetReference) {
	if err := b.validatePutExternalDatasetReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExternalDatasetReference",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) PutTimeouts(value *BigqueryDatasetTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataset) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := b.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetAccess() {
	_jsii_.InvokeVoid(
		b,
		"resetAccess",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultCollation() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultCollation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultEncryptionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultPartitionExpirationMs() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultPartitionExpirationMs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDefaultTableExpirationMs() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultTableExpirationMs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDeleteContentsOnDestroy() {
	_jsii_.InvokeVoid(
		b,
		"resetDeleteContentsOnDestroy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetExternalCatalogDatasetOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalCatalogDatasetOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetExternalDatasetReference() {
	_jsii_.InvokeVoid(
		b,
		"resetExternalDatasetReference",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		b,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetIsCaseInsensitive() {
	_jsii_.InvokeVoid(
		b,
		"resetIsCaseInsensitive",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetLocation() {
	_jsii_.InvokeVoid(
		b,
		"resetLocation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetMaxTimeTravelHours() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxTimeTravelHours",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetResourceTags() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetStorageBillingModel() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageBillingModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataset) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}

