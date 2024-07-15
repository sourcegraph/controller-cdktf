package googledatabasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatabasemigrationserviceconnectionprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile google_database_migration_service_connection_profile}.
type GoogleDatabaseMigrationServiceConnectionProfile interface {
	cdktf.TerraformResource
	Alloydb() GoogleDatabaseMigrationServiceConnectionProfileAlloydbOutputReference
	AlloydbInput() *GoogleDatabaseMigrationServiceConnectionProfileAlloydb
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloudsql() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference
	CloudsqlInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsql
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
	EffectiveLabels() cdktf.StringMap
	Error() GoogleDatabaseMigrationServiceConnectionProfileErrorList
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Mysql() GoogleDatabaseMigrationServiceConnectionProfileMysqlOutputReference
	MysqlInput() *GoogleDatabaseMigrationServiceConnectionProfileMysql
	Name() *string
	// The tree node.
	Node() constructs.Node
	Oracle() GoogleDatabaseMigrationServiceConnectionProfileOracleOutputReference
	OracleInput() *GoogleDatabaseMigrationServiceConnectionProfileOracle
	Postgresql() GoogleDatabaseMigrationServiceConnectionProfilePostgresqlOutputReference
	PostgresqlInput() *GoogleDatabaseMigrationServiceConnectionProfilePostgresql
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
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDatabaseMigrationServiceConnectionProfileTimeoutsOutputReference
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
	PutAlloydb(value *GoogleDatabaseMigrationServiceConnectionProfileAlloydb)
	PutCloudsql(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsql)
	PutMysql(value *GoogleDatabaseMigrationServiceConnectionProfileMysql)
	PutOracle(value *GoogleDatabaseMigrationServiceConnectionProfileOracle)
	PutPostgresql(value *GoogleDatabaseMigrationServiceConnectionProfilePostgresql)
	PutTimeouts(value *GoogleDatabaseMigrationServiceConnectionProfileTimeouts)
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
}

// The jsii proxy struct for GoogleDatabaseMigrationServiceConnectionProfile
type jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Alloydb() GoogleDatabaseMigrationServiceConnectionProfileAlloydbOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileAlloydbOutputReference
	_jsii_.Get(
		j,
		"alloydb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) AlloydbInput() *GoogleDatabaseMigrationServiceConnectionProfileAlloydb {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileAlloydb
	_jsii_.Get(
		j,
		"alloydbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Cloudsql() GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileCloudsqlOutputReference
	_jsii_.Get(
		j,
		"cloudsql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) CloudsqlInput() *GoogleDatabaseMigrationServiceConnectionProfileCloudsql {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileCloudsql
	_jsii_.Get(
		j,
		"cloudsqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ConnectionProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ConnectionProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Dbprovider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbprovider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Error() GoogleDatabaseMigrationServiceConnectionProfileErrorList {
	var returns GoogleDatabaseMigrationServiceConnectionProfileErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Mysql() GoogleDatabaseMigrationServiceConnectionProfileMysqlOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) MysqlInput() *GoogleDatabaseMigrationServiceConnectionProfileMysql {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Oracle() GoogleDatabaseMigrationServiceConnectionProfileOracleOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileOracleOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) OracleInput() *GoogleDatabaseMigrationServiceConnectionProfileOracle {
	var returns *GoogleDatabaseMigrationServiceConnectionProfileOracle
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Postgresql() GoogleDatabaseMigrationServiceConnectionProfilePostgresqlOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfilePostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PostgresqlInput() *GoogleDatabaseMigrationServiceConnectionProfilePostgresql {
	var returns *GoogleDatabaseMigrationServiceConnectionProfilePostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) Timeouts() GoogleDatabaseMigrationServiceConnectionProfileTimeoutsOutputReference {
	var returns GoogleDatabaseMigrationServiceConnectionProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile google_database_migration_service_connection_profile} Resource.
func NewGoogleDatabaseMigrationServiceConnectionProfile(scope constructs.Construct, id *string, config *GoogleDatabaseMigrationServiceConnectionProfileConfig) GoogleDatabaseMigrationServiceConnectionProfile {
	_init_.Initialize()

	if err := validateNewGoogleDatabaseMigrationServiceConnectionProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_database_migration_service_connection_profile google_database_migration_service_connection_profile} Resource.
func NewGoogleDatabaseMigrationServiceConnectionProfile_Override(g GoogleDatabaseMigrationServiceConnectionProfile, scope constructs.Construct, id *string, config *GoogleDatabaseMigrationServiceConnectionProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetConnectionProfileId(val *string) {
	if err := j.validateSetConnectionProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDatabaseMigrationServiceConnectionProfile resource upon running "cdktf plan <stack-name>".
func GoogleDatabaseMigrationServiceConnectionProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceConnectionProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
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
func GoogleDatabaseMigrationServiceConnectionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceConnectionProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatabaseMigrationServiceConnectionProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceConnectionProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatabaseMigrationServiceConnectionProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceConnectionProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDatabaseMigrationServiceConnectionProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceConnectionProfile.GoogleDatabaseMigrationServiceConnectionProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutAlloydb(value *GoogleDatabaseMigrationServiceConnectionProfileAlloydb) {
	if err := g.validatePutAlloydbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAlloydb",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutCloudsql(value *GoogleDatabaseMigrationServiceConnectionProfileCloudsql) {
	if err := g.validatePutCloudsqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudsql",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutMysql(value *GoogleDatabaseMigrationServiceConnectionProfileMysql) {
	if err := g.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysql",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutOracle(value *GoogleDatabaseMigrationServiceConnectionProfileOracle) {
	if err := g.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracle",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutPostgresql(value *GoogleDatabaseMigrationServiceConnectionProfilePostgresql) {
	if err := g.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) PutTimeouts(value *GoogleDatabaseMigrationServiceConnectionProfileTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetAlloydb() {
	_jsii_.InvokeVoid(
		g,
		"resetAlloydb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetCloudsql() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudsql",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetMysql() {
	_jsii_.InvokeVoid(
		g,
		"resetMysql",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetOracle() {
	_jsii_.InvokeVoid(
		g,
		"resetOracle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetPostgresql() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceConnectionProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

