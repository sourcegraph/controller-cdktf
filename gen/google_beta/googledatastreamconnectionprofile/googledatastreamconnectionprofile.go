package googledatastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamconnectionprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_datastream_connection_profile google_datastream_connection_profile}.
type GoogleDatastreamConnectionProfile interface {
	cdktf.TerraformResource
	BigqueryProfile() GoogleDatastreamConnectionProfileBigqueryProfileOutputReference
	BigqueryProfileInput() *GoogleDatastreamConnectionProfileBigqueryProfile
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardSshConnectivity() GoogleDatastreamConnectionProfileForwardSshConnectivityOutputReference
	ForwardSshConnectivityInput() *GoogleDatastreamConnectionProfileForwardSshConnectivity
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsProfile() GoogleDatastreamConnectionProfileGcsProfileOutputReference
	GcsProfileInput() *GoogleDatastreamConnectionProfileGcsProfile
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
	MysqlProfile() GoogleDatastreamConnectionProfileMysqlProfileOutputReference
	MysqlProfileInput() *GoogleDatastreamConnectionProfileMysqlProfile
	Name() *string
	// The tree node.
	Node() constructs.Node
	OracleProfile() GoogleDatastreamConnectionProfileOracleProfileOutputReference
	OracleProfileInput() *GoogleDatastreamConnectionProfileOracleProfile
	PostgresqlProfile() GoogleDatastreamConnectionProfilePostgresqlProfileOutputReference
	PostgresqlProfileInput() *GoogleDatastreamConnectionProfilePostgresqlProfile
	PrivateConnectivity() GoogleDatastreamConnectionProfilePrivateConnectivityOutputReference
	PrivateConnectivityInput() *GoogleDatastreamConnectionProfilePrivateConnectivity
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
	SqlServerProfile() GoogleDatastreamConnectionProfileSqlServerProfileOutputReference
	SqlServerProfileInput() *GoogleDatastreamConnectionProfileSqlServerProfile
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDatastreamConnectionProfileTimeoutsOutputReference
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
	PutBigqueryProfile(value *GoogleDatastreamConnectionProfileBigqueryProfile)
	PutForwardSshConnectivity(value *GoogleDatastreamConnectionProfileForwardSshConnectivity)
	PutGcsProfile(value *GoogleDatastreamConnectionProfileGcsProfile)
	PutMysqlProfile(value *GoogleDatastreamConnectionProfileMysqlProfile)
	PutOracleProfile(value *GoogleDatastreamConnectionProfileOracleProfile)
	PutPostgresqlProfile(value *GoogleDatastreamConnectionProfilePostgresqlProfile)
	PutPrivateConnectivity(value *GoogleDatastreamConnectionProfilePrivateConnectivity)
	PutSqlServerProfile(value *GoogleDatastreamConnectionProfileSqlServerProfile)
	PutTimeouts(value *GoogleDatastreamConnectionProfileTimeouts)
	ResetBigqueryProfile()
	ResetForwardSshConnectivity()
	ResetGcsProfile()
	ResetId()
	ResetLabels()
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
}

// The jsii proxy struct for GoogleDatastreamConnectionProfile
type jsiiProxy_GoogleDatastreamConnectionProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) BigqueryProfile() GoogleDatastreamConnectionProfileBigqueryProfileOutputReference {
	var returns GoogleDatastreamConnectionProfileBigqueryProfileOutputReference
	_jsii_.Get(
		j,
		"bigqueryProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) BigqueryProfileInput() *GoogleDatastreamConnectionProfileBigqueryProfile {
	var returns *GoogleDatastreamConnectionProfileBigqueryProfile
	_jsii_.Get(
		j,
		"bigqueryProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ConnectionProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ConnectionProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ForwardSshConnectivity() GoogleDatastreamConnectionProfileForwardSshConnectivityOutputReference {
	var returns GoogleDatastreamConnectionProfileForwardSshConnectivityOutputReference
	_jsii_.Get(
		j,
		"forwardSshConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ForwardSshConnectivityInput() *GoogleDatastreamConnectionProfileForwardSshConnectivity {
	var returns *GoogleDatastreamConnectionProfileForwardSshConnectivity
	_jsii_.Get(
		j,
		"forwardSshConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) GcsProfile() GoogleDatastreamConnectionProfileGcsProfileOutputReference {
	var returns GoogleDatastreamConnectionProfileGcsProfileOutputReference
	_jsii_.Get(
		j,
		"gcsProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) GcsProfileInput() *GoogleDatastreamConnectionProfileGcsProfile {
	var returns *GoogleDatastreamConnectionProfileGcsProfile
	_jsii_.Get(
		j,
		"gcsProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) MysqlProfile() GoogleDatastreamConnectionProfileMysqlProfileOutputReference {
	var returns GoogleDatastreamConnectionProfileMysqlProfileOutputReference
	_jsii_.Get(
		j,
		"mysqlProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) MysqlProfileInput() *GoogleDatastreamConnectionProfileMysqlProfile {
	var returns *GoogleDatastreamConnectionProfileMysqlProfile
	_jsii_.Get(
		j,
		"mysqlProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) OracleProfile() GoogleDatastreamConnectionProfileOracleProfileOutputReference {
	var returns GoogleDatastreamConnectionProfileOracleProfileOutputReference
	_jsii_.Get(
		j,
		"oracleProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) OracleProfileInput() *GoogleDatastreamConnectionProfileOracleProfile {
	var returns *GoogleDatastreamConnectionProfileOracleProfile
	_jsii_.Get(
		j,
		"oracleProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) PostgresqlProfile() GoogleDatastreamConnectionProfilePostgresqlProfileOutputReference {
	var returns GoogleDatastreamConnectionProfilePostgresqlProfileOutputReference
	_jsii_.Get(
		j,
		"postgresqlProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) PostgresqlProfileInput() *GoogleDatastreamConnectionProfilePostgresqlProfile {
	var returns *GoogleDatastreamConnectionProfilePostgresqlProfile
	_jsii_.Get(
		j,
		"postgresqlProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) PrivateConnectivity() GoogleDatastreamConnectionProfilePrivateConnectivityOutputReference {
	var returns GoogleDatastreamConnectionProfilePrivateConnectivityOutputReference
	_jsii_.Get(
		j,
		"privateConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) PrivateConnectivityInput() *GoogleDatastreamConnectionProfilePrivateConnectivity {
	var returns *GoogleDatastreamConnectionProfilePrivateConnectivity
	_jsii_.Get(
		j,
		"privateConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) SqlServerProfile() GoogleDatastreamConnectionProfileSqlServerProfileOutputReference {
	var returns GoogleDatastreamConnectionProfileSqlServerProfileOutputReference
	_jsii_.Get(
		j,
		"sqlServerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) SqlServerProfileInput() *GoogleDatastreamConnectionProfileSqlServerProfile {
	var returns *GoogleDatastreamConnectionProfileSqlServerProfile
	_jsii_.Get(
		j,
		"sqlServerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) Timeouts() GoogleDatastreamConnectionProfileTimeoutsOutputReference {
	var returns GoogleDatastreamConnectionProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_datastream_connection_profile google_datastream_connection_profile} Resource.
func NewGoogleDatastreamConnectionProfile(scope constructs.Construct, id *string, config *GoogleDatastreamConnectionProfileConfig) GoogleDatastreamConnectionProfile {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamConnectionProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamConnectionProfile{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_datastream_connection_profile google_datastream_connection_profile} Resource.
func NewGoogleDatastreamConnectionProfile_Override(g GoogleDatastreamConnectionProfile, scope constructs.Construct, id *string, config *GoogleDatastreamConnectionProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetConnectionProfileId(val *string) {
	if err := j.validateSetConnectionProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDatastreamConnectionProfile resource upon running "cdktf plan <stack-name>".
func GoogleDatastreamConnectionProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDatastreamConnectionProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
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
func GoogleDatastreamConnectionProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatastreamConnectionProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatastreamConnectionProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatastreamConnectionProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatastreamConnectionProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatastreamConnectionProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDatastreamConnectionProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutBigqueryProfile(value *GoogleDatastreamConnectionProfileBigqueryProfile) {
	if err := g.validatePutBigqueryProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutForwardSshConnectivity(value *GoogleDatastreamConnectionProfileForwardSshConnectivity) {
	if err := g.validatePutForwardSshConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putForwardSshConnectivity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutGcsProfile(value *GoogleDatastreamConnectionProfileGcsProfile) {
	if err := g.validatePutGcsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutMysqlProfile(value *GoogleDatastreamConnectionProfileMysqlProfile) {
	if err := g.validatePutMysqlProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysqlProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutOracleProfile(value *GoogleDatastreamConnectionProfileOracleProfile) {
	if err := g.validatePutOracleProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutPostgresqlProfile(value *GoogleDatastreamConnectionProfilePostgresqlProfile) {
	if err := g.validatePutPostgresqlProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutPrivateConnectivity(value *GoogleDatastreamConnectionProfilePrivateConnectivity) {
	if err := g.validatePutPrivateConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateConnectivity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutSqlServerProfile(value *GoogleDatastreamConnectionProfileSqlServerProfile) {
	if err := g.validatePutSqlServerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlServerProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) PutTimeouts(value *GoogleDatastreamConnectionProfileTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetBigqueryProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetForwardSshConnectivity() {
	_jsii_.InvokeVoid(
		g,
		"resetForwardSshConnectivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetGcsProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetMysqlProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetMysqlProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetOracleProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetPostgresqlProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetPrivateConnectivity() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateConnectivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetSqlServerProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlServerProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

