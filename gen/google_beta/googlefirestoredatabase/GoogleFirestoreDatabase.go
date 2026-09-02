package googlefirestoredatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefirestoredatabase/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_firestore_database google_firestore_database}.
type GoogleFirestoreDatabase interface {
	cdktn.TerraformResource
	AppEngineIntegrationMode() *string
	SetAppEngineIntegrationMode(val *string)
	AppEngineIntegrationModeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CmekConfig() GoogleFirestoreDatabaseCmekConfigOutputReference
	CmekConfigInput() *GoogleFirestoreDatabaseCmekConfig
	ConcurrencyMode() *string
	SetConcurrencyMode(val *string)
	ConcurrencyModeInput() *string
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
	CreateTime() *string
	DatabaseEdition() *string
	SetDatabaseEdition(val *string)
	DatabaseEditionInput() *string
	DeleteProtectionState() *string
	SetDeleteProtectionState(val *string)
	DeleteProtectionStateInput() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EarliestVersionTime() *string
	Etag() *string
	FirestoreDataAccessMode() *string
	SetFirestoreDataAccessMode(val *string)
	FirestoreDataAccessModeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeyPrefix() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocationId() *string
	SetLocationId(val *string)
	LocationIdInput() *string
	MongodbCompatibleDataAccessMode() *string
	SetMongodbCompatibleDataAccessMode(val *string)
	MongodbCompatibleDataAccessModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PointInTimeRecoveryEnablement() *string
	SetPointInTimeRecoveryEnablement(val *string)
	PointInTimeRecoveryEnablementInput() *string
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
	RealtimeUpdatesMode() *string
	SetRealtimeUpdatesMode(val *string)
	RealtimeUpdatesModeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleFirestoreDatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Uid() *string
	UpdateTime() *string
	VersionRetentionPeriod() *string
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
	PutCmekConfig(value *GoogleFirestoreDatabaseCmekConfig)
	PutTimeouts(value *GoogleFirestoreDatabaseTimeouts)
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
	ResetAppEngineIntegrationMode()
	ResetCmekConfig()
	ResetConcurrencyMode()
	ResetDatabaseEdition()
	ResetDeleteProtectionState()
	ResetDeletionPolicy()
	ResetFirestoreDataAccessMode()
	ResetId()
	ResetMongodbCompatibleDataAccessMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecoveryEnablement()
	ResetProject()
	ResetRealtimeUpdatesMode()
	ResetTags()
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

// The jsii proxy struct for GoogleFirestoreDatabase
type jsiiProxy_GoogleFirestoreDatabase struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleFirestoreDatabase) AppEngineIntegrationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineIntegrationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) AppEngineIntegrationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineIntegrationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) CmekConfig() GoogleFirestoreDatabaseCmekConfigOutputReference {
	var returns GoogleFirestoreDatabaseCmekConfigOutputReference
	_jsii_.Get(
		j,
		"cmekConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) CmekConfigInput() *GoogleFirestoreDatabaseCmekConfig {
	var returns *GoogleFirestoreDatabaseCmekConfig
	_jsii_.Get(
		j,
		"cmekConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) ConcurrencyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) ConcurrencyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DatabaseEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DatabaseEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DeleteProtectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteProtectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DeleteProtectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteProtectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) EarliestVersionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"earliestVersionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) FirestoreDataAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreDataAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) FirestoreDataAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreDataAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) LocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) LocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) MongodbCompatibleDataAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbCompatibleDataAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) MongodbCompatibleDataAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongodbCompatibleDataAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) PointInTimeRecoveryEnablement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnablement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) PointInTimeRecoveryEnablementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnablementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) RealtimeUpdatesMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeUpdatesMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) RealtimeUpdatesModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeUpdatesModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Timeouts() GoogleFirestoreDatabaseTimeoutsOutputReference {
	var returns GoogleFirestoreDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreDatabase) VersionRetentionPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionRetentionPeriod",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_firestore_database google_firestore_database} Resource.
func NewGoogleFirestoreDatabase(scope constructs.Construct, id *string, config *GoogleFirestoreDatabaseConfig) GoogleFirestoreDatabase {
	_init_.Initialize()

	if err := validateNewGoogleFirestoreDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFirestoreDatabase{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_firestore_database google_firestore_database} Resource.
func NewGoogleFirestoreDatabase_Override(g GoogleFirestoreDatabase, scope constructs.Construct, id *string, config *GoogleFirestoreDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetAppEngineIntegrationMode(val *string) {
	if err := j.validateSetAppEngineIntegrationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appEngineIntegrationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetConcurrencyMode(val *string) {
	if err := j.validateSetConcurrencyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetDatabaseEdition(val *string) {
	if err := j.validateSetDatabaseEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseEdition",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetDeleteProtectionState(val *string) {
	if err := j.validateSetDeleteProtectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteProtectionState",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetFirestoreDataAccessMode(val *string) {
	if err := j.validateSetFirestoreDataAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firestoreDataAccessMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetLocationId(val *string) {
	if err := j.validateSetLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationId",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetMongodbCompatibleDataAccessMode(val *string) {
	if err := j.validateSetMongodbCompatibleDataAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongodbCompatibleDataAccessMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetPointInTimeRecoveryEnablement(val *string) {
	if err := j.validateSetPointInTimeRecoveryEnablementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRecoveryEnablement",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetRealtimeUpdatesMode(val *string) {
	if err := j.validateSetRealtimeUpdatesModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realtimeUpdatesMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreDatabase)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a GoogleFirestoreDatabase resource upon running "cdktn plan <stack-name>".
func GoogleFirestoreDatabase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleFirestoreDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
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
func GoogleFirestoreDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleFirestoreDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleFirestoreDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleFirestoreDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleFirestoreDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleFirestoreDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleFirestoreDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleFirestoreDatabase.GoogleFirestoreDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleFirestoreDatabase) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) PutCmekConfig(value *GoogleFirestoreDatabaseCmekConfig) {
	if err := g.validatePutCmekConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCmekConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) PutTimeouts(value *GoogleFirestoreDatabaseTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetAppEngineIntegrationMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineIntegrationMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetCmekConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCmekConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetConcurrencyMode() {
	_jsii_.InvokeVoid(
		g,
		"resetConcurrencyMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetDatabaseEdition() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseEdition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetDeleteProtectionState() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteProtectionState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetFirestoreDataAccessMode() {
	_jsii_.InvokeVoid(
		g,
		"resetFirestoreDataAccessMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetMongodbCompatibleDataAccessMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMongodbCompatibleDataAccessMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetPointInTimeRecoveryEnablement() {
	_jsii_.InvokeVoid(
		g,
		"resetPointInTimeRecoveryEnablement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetRealtimeUpdatesMode() {
	_jsii_.InvokeVoid(
		g,
		"resetRealtimeUpdatesMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreDatabase) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

