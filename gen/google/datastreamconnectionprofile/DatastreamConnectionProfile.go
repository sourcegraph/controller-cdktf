package datastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamconnectionprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_connection_profile google_datastream_connection_profile}.
type DatastreamConnectionProfile interface {
	cdktn.TerraformResource
	BigqueryProfile() DatastreamConnectionProfileBigqueryProfileOutputReference
	BigqueryProfileInput() *DatastreamConnectionProfileBigqueryProfile
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionProfileId() *string
	SetConnectionProfileId(val *string)
	ConnectionProfileIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateWithoutValidation() interface{}
	SetCreateWithoutValidation(val interface{})
	CreateWithoutValidationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktn.StringMap
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	ForwardSshConnectivity() DatastreamConnectionProfileForwardSshConnectivityOutputReference
	ForwardSshConnectivityInput() *DatastreamConnectionProfileForwardSshConnectivity
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsProfile() DatastreamConnectionProfileGcsProfileOutputReference
	GcsProfileInput() *DatastreamConnectionProfileGcsProfile
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MongodbProfile() DatastreamConnectionProfileMongodbProfileOutputReference
	MongodbProfileInput() *DatastreamConnectionProfileMongodbProfile
	MysqlProfile() DatastreamConnectionProfileMysqlProfileOutputReference
	MysqlProfileInput() *DatastreamConnectionProfileMysqlProfile
	Name() *string
	// The tree node.
	Node() constructs.Node
	OracleProfile() DatastreamConnectionProfileOracleProfileOutputReference
	OracleProfileInput() *DatastreamConnectionProfileOracleProfile
	PostgresqlProfile() DatastreamConnectionProfilePostgresqlProfileOutputReference
	PostgresqlProfileInput() *DatastreamConnectionProfilePostgresqlProfile
	PrivateConnectivity() DatastreamConnectionProfilePrivateConnectivityOutputReference
	PrivateConnectivityInput() *DatastreamConnectionProfilePrivateConnectivity
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
	SqlServerProfile() DatastreamConnectionProfileSqlServerProfileOutputReference
	SqlServerProfileInput() *DatastreamConnectionProfileSqlServerProfile
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatastreamConnectionProfileTimeoutsOutputReference
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
	PutBigqueryProfile(value *DatastreamConnectionProfileBigqueryProfile)
	PutForwardSshConnectivity(value *DatastreamConnectionProfileForwardSshConnectivity)
	PutGcsProfile(value *DatastreamConnectionProfileGcsProfile)
	PutMongodbProfile(value *DatastreamConnectionProfileMongodbProfile)
	PutMysqlProfile(value *DatastreamConnectionProfileMysqlProfile)
	PutOracleProfile(value *DatastreamConnectionProfileOracleProfile)
	PutPostgresqlProfile(value *DatastreamConnectionProfilePostgresqlProfile)
	PutPrivateConnectivity(value *DatastreamConnectionProfilePrivateConnectivity)
	PutSqlServerProfile(value *DatastreamConnectionProfileSqlServerProfile)
	PutTimeouts(value *DatastreamConnectionProfileTimeouts)
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
	ResetBigqueryProfile()
	ResetCreateWithoutValidation()
	ResetForwardSshConnectivity()
	ResetGcsProfile()
	ResetId()
	ResetLabels()
	ResetMongodbProfile()
	ResetMysqlProfile()
	ResetOracleProfile()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresqlProfile()
	ResetPrivateConnectivity()
	ResetProject()
	ResetSqlServerProfile()
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

// The jsii proxy struct for DatastreamConnectionProfile
type jsiiProxy_DatastreamConnectionProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DatastreamConnectionProfile) BigqueryProfile() DatastreamConnectionProfileBigqueryProfileOutputReference {
	var returns DatastreamConnectionProfileBigqueryProfileOutputReference
	_jsii_.Get(
		j,
		"bigqueryProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) BigqueryProfileInput() *DatastreamConnectionProfileBigqueryProfile {
	var returns *DatastreamConnectionProfileBigqueryProfile
	_jsii_.Get(
		j,
		"bigqueryProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ConnectionProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ConnectionProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) CreateWithoutValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createWithoutValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) CreateWithoutValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createWithoutValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ForwardSshConnectivity() DatastreamConnectionProfileForwardSshConnectivityOutputReference {
	var returns DatastreamConnectionProfileForwardSshConnectivityOutputReference
	_jsii_.Get(
		j,
		"forwardSshConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ForwardSshConnectivityInput() *DatastreamConnectionProfileForwardSshConnectivity {
	var returns *DatastreamConnectionProfileForwardSshConnectivity
	_jsii_.Get(
		j,
		"forwardSshConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) GcsProfile() DatastreamConnectionProfileGcsProfileOutputReference {
	var returns DatastreamConnectionProfileGcsProfileOutputReference
	_jsii_.Get(
		j,
		"gcsProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) GcsProfileInput() *DatastreamConnectionProfileGcsProfile {
	var returns *DatastreamConnectionProfileGcsProfile
	_jsii_.Get(
		j,
		"gcsProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) MongodbProfile() DatastreamConnectionProfileMongodbProfileOutputReference {
	var returns DatastreamConnectionProfileMongodbProfileOutputReference
	_jsii_.Get(
		j,
		"mongodbProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) MongodbProfileInput() *DatastreamConnectionProfileMongodbProfile {
	var returns *DatastreamConnectionProfileMongodbProfile
	_jsii_.Get(
		j,
		"mongodbProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) MysqlProfile() DatastreamConnectionProfileMysqlProfileOutputReference {
	var returns DatastreamConnectionProfileMysqlProfileOutputReference
	_jsii_.Get(
		j,
		"mysqlProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) MysqlProfileInput() *DatastreamConnectionProfileMysqlProfile {
	var returns *DatastreamConnectionProfileMysqlProfile
	_jsii_.Get(
		j,
		"mysqlProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) OracleProfile() DatastreamConnectionProfileOracleProfileOutputReference {
	var returns DatastreamConnectionProfileOracleProfileOutputReference
	_jsii_.Get(
		j,
		"oracleProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) OracleProfileInput() *DatastreamConnectionProfileOracleProfile {
	var returns *DatastreamConnectionProfileOracleProfile
	_jsii_.Get(
		j,
		"oracleProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) PostgresqlProfile() DatastreamConnectionProfilePostgresqlProfileOutputReference {
	var returns DatastreamConnectionProfilePostgresqlProfileOutputReference
	_jsii_.Get(
		j,
		"postgresqlProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) PostgresqlProfileInput() *DatastreamConnectionProfilePostgresqlProfile {
	var returns *DatastreamConnectionProfilePostgresqlProfile
	_jsii_.Get(
		j,
		"postgresqlProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) PrivateConnectivity() DatastreamConnectionProfilePrivateConnectivityOutputReference {
	var returns DatastreamConnectionProfilePrivateConnectivityOutputReference
	_jsii_.Get(
		j,
		"privateConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) PrivateConnectivityInput() *DatastreamConnectionProfilePrivateConnectivity {
	var returns *DatastreamConnectionProfilePrivateConnectivity
	_jsii_.Get(
		j,
		"privateConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) SqlServerProfile() DatastreamConnectionProfileSqlServerProfileOutputReference {
	var returns DatastreamConnectionProfileSqlServerProfileOutputReference
	_jsii_.Get(
		j,
		"sqlServerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) SqlServerProfileInput() *DatastreamConnectionProfileSqlServerProfile {
	var returns *DatastreamConnectionProfileSqlServerProfile
	_jsii_.Get(
		j,
		"sqlServerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) Timeouts() DatastreamConnectionProfileTimeoutsOutputReference {
	var returns DatastreamConnectionProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamConnectionProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_connection_profile google_datastream_connection_profile} Resource.
func NewDatastreamConnectionProfile(scope constructs.Construct, id *string, config *DatastreamConnectionProfileConfig) DatastreamConnectionProfile {
	_init_.Initialize()

	if err := validateNewDatastreamConnectionProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamConnectionProfile{}

	_jsii_.Create(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_connection_profile google_datastream_connection_profile} Resource.
func NewDatastreamConnectionProfile_Override(d DatastreamConnectionProfile, scope constructs.Construct, id *string, config *DatastreamConnectionProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetConnectionProfileId(val *string) {
	if err := j.validateSetConnectionProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProfileId",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetCreateWithoutValidation(val interface{}) {
	if err := j.validateSetCreateWithoutValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createWithoutValidation",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatastreamConnectionProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a DatastreamConnectionProfile resource upon running "cdktn plan <stack-name>".
func DatastreamConnectionProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDatastreamConnectionProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
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
func DatastreamConnectionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastreamConnectionProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatastreamConnectionProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastreamConnectionProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatastreamConnectionProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatastreamConnectionProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatastreamConnectionProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.datastreamConnectionProfile.DatastreamConnectionProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutBigqueryProfile(value *DatastreamConnectionProfileBigqueryProfile) {
	if err := d.validatePutBigqueryProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBigqueryProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutForwardSshConnectivity(value *DatastreamConnectionProfileForwardSshConnectivity) {
	if err := d.validatePutForwardSshConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForwardSshConnectivity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutGcsProfile(value *DatastreamConnectionProfileGcsProfile) {
	if err := d.validatePutGcsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcsProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutMongodbProfile(value *DatastreamConnectionProfileMongodbProfile) {
	if err := d.validatePutMongodbProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongodbProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutMysqlProfile(value *DatastreamConnectionProfileMysqlProfile) {
	if err := d.validatePutMysqlProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutOracleProfile(value *DatastreamConnectionProfileOracleProfile) {
	if err := d.validatePutOracleProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutPostgresqlProfile(value *DatastreamConnectionProfilePostgresqlProfile) {
	if err := d.validatePutPostgresqlProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresqlProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutPrivateConnectivity(value *DatastreamConnectionProfilePrivateConnectivity) {
	if err := d.validatePutPrivateConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrivateConnectivity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutSqlServerProfile(value *DatastreamConnectionProfileSqlServerProfile) {
	if err := d.validatePutSqlServerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlServerProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) PutTimeouts(value *DatastreamConnectionProfileTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetBigqueryProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetBigqueryProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetCreateWithoutValidation() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateWithoutValidation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetForwardSshConnectivity() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardSshConnectivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetGcsProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetGcsProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetMongodbProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetMysqlProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetOracleProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetPostgresqlProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresqlProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetPrivateConnectivity() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateConnectivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetSqlServerProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlServerProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamConnectionProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamConnectionProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

