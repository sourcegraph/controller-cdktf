package databasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/databasemigrationserviceconnectionprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/database_migration_service_connection_profile google_database_migration_service_connection_profile}.
type DatabaseMigrationServiceConnectionProfile interface {
	cdktn.TerraformResource
	Alloydb() DatabaseMigrationServiceConnectionProfileAlloydbOutputReference
	AlloydbInput() *DatabaseMigrationServiceConnectionProfileAlloydb
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Cloudsql() DatabaseMigrationServiceConnectionProfileCloudsqlOutputReference
	CloudsqlInput() *DatabaseMigrationServiceConnectionProfileCloudsql
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
	CreateTime() *string
	Dbprovider() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktn.StringMap
	Error() DatabaseMigrationServiceConnectionProfileErrorList
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
	Mysql() DatabaseMigrationServiceConnectionProfileMysqlOutputReference
	MysqlInput() *DatabaseMigrationServiceConnectionProfileMysql
	Name() *string
	// The tree node.
	Node() constructs.Node
	Oracle() DatabaseMigrationServiceConnectionProfileOracleOutputReference
	OracleInput() *DatabaseMigrationServiceConnectionProfileOracle
	Postgresql() DatabaseMigrationServiceConnectionProfilePostgresqlOutputReference
	PostgresqlInput() *DatabaseMigrationServiceConnectionProfilePostgresql
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
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabaseMigrationServiceConnectionProfileTimeoutsOutputReference
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
	PutAlloydb(value *DatabaseMigrationServiceConnectionProfileAlloydb)
	PutCloudsql(value *DatabaseMigrationServiceConnectionProfileCloudsql)
	PutMysql(value *DatabaseMigrationServiceConnectionProfileMysql)
	PutOracle(value *DatabaseMigrationServiceConnectionProfileOracle)
	PutPostgresql(value *DatabaseMigrationServiceConnectionProfilePostgresql)
	PutTimeouts(value *DatabaseMigrationServiceConnectionProfileTimeouts)
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
	ResetAlloydb()
	ResetCloudsql()
	ResetDisplayName()
	ResetId()
	ResetLabels()
	ResetLocation()
	ResetMysql()
	ResetOracle()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresql()
	ResetProject()
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

// The jsii proxy struct for DatabaseMigrationServiceConnectionProfile
type jsiiProxy_DatabaseMigrationServiceConnectionProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Alloydb() DatabaseMigrationServiceConnectionProfileAlloydbOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileAlloydbOutputReference
	_jsii_.Get(
		j,
		"alloydb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) AlloydbInput() *DatabaseMigrationServiceConnectionProfileAlloydb {
	var returns *DatabaseMigrationServiceConnectionProfileAlloydb
	_jsii_.Get(
		j,
		"alloydbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Cloudsql() DatabaseMigrationServiceConnectionProfileCloudsqlOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileCloudsqlOutputReference
	_jsii_.Get(
		j,
		"cloudsql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) CloudsqlInput() *DatabaseMigrationServiceConnectionProfileCloudsql {
	var returns *DatabaseMigrationServiceConnectionProfileCloudsql
	_jsii_.Get(
		j,
		"cloudsqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ConnectionProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ConnectionProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Dbprovider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbprovider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Error() DatabaseMigrationServiceConnectionProfileErrorList {
	var returns DatabaseMigrationServiceConnectionProfileErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Mysql() DatabaseMigrationServiceConnectionProfileMysqlOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) MysqlInput() *DatabaseMigrationServiceConnectionProfileMysql {
	var returns *DatabaseMigrationServiceConnectionProfileMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Oracle() DatabaseMigrationServiceConnectionProfileOracleOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileOracleOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) OracleInput() *DatabaseMigrationServiceConnectionProfileOracle {
	var returns *DatabaseMigrationServiceConnectionProfileOracle
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Postgresql() DatabaseMigrationServiceConnectionProfilePostgresqlOutputReference {
	var returns DatabaseMigrationServiceConnectionProfilePostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PostgresqlInput() *DatabaseMigrationServiceConnectionProfilePostgresql {
	var returns *DatabaseMigrationServiceConnectionProfilePostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) Timeouts() DatabaseMigrationServiceConnectionProfileTimeoutsOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/database_migration_service_connection_profile google_database_migration_service_connection_profile} Resource.
func NewDatabaseMigrationServiceConnectionProfile(scope constructs.Construct, id *string, config *DatabaseMigrationServiceConnectionProfileConfig) DatabaseMigrationServiceConnectionProfile {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceConnectionProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceConnectionProfile{}

	_jsii_.Create(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/database_migration_service_connection_profile google_database_migration_service_connection_profile} Resource.
func NewDatabaseMigrationServiceConnectionProfile_Override(d DatabaseMigrationServiceConnectionProfile, scope constructs.Construct, id *string, config *DatabaseMigrationServiceConnectionProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetConnectionProfileId(val *string) {
	if err := j.validateSetConnectionProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProfileId",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a DatabaseMigrationServiceConnectionProfile resource upon running "cdktn plan <stack-name>".
func DatabaseMigrationServiceConnectionProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceConnectionProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
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
func DatabaseMigrationServiceConnectionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceConnectionProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMigrationServiceConnectionProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceConnectionProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseMigrationServiceConnectionProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseMigrationServiceConnectionProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseMigrationServiceConnectionProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutAlloydb(value *DatabaseMigrationServiceConnectionProfileAlloydb) {
	if err := d.validatePutAlloydbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlloydb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutCloudsql(value *DatabaseMigrationServiceConnectionProfileCloudsql) {
	if err := d.validatePutCloudsqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudsql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutMysql(value *DatabaseMigrationServiceConnectionProfileMysql) {
	if err := d.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutOracle(value *DatabaseMigrationServiceConnectionProfileOracle) {
	if err := d.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutPostgresql(value *DatabaseMigrationServiceConnectionProfilePostgresql) {
	if err := d.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) PutTimeouts(value *DatabaseMigrationServiceConnectionProfileTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetAlloydb() {
	_jsii_.InvokeVoid(
		d,
		"resetAlloydb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetCloudsql() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudsql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetMysql() {
	_jsii_.InvokeVoid(
		d,
		"resetMysql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetOracle() {
	_jsii_.InvokeVoid(
		d,
		"resetOracle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetPostgresql() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

