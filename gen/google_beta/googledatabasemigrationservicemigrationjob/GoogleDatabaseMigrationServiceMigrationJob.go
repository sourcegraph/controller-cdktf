package googledatabasemigrationservicemigrationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatabasemigrationservicemigrationjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_migration_job google_database_migration_service_migration_job}.
type GoogleDatabaseMigrationServiceMigrationJob interface {
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DumpFlags() GoogleDatabaseMigrationServiceMigrationJobDumpFlagsOutputReference
	DumpFlagsInput() *GoogleDatabaseMigrationServiceMigrationJobDumpFlags
	DumpPath() *string
	SetDumpPath(val *string)
	DumpPathInput() *string
	DumpType() *string
	SetDumpType(val *string)
	DumpTypeInput() *string
	EffectiveLabels() cdktf.StringMap
	Error() GoogleDatabaseMigrationServiceMigrationJobErrorList
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
	MigrationJobId() *string
	SetMigrationJobId(val *string)
	MigrationJobIdInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PerformanceConfig() GoogleDatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference
	PerformanceConfigInput() *GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig
	Phase() *string
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
	ReverseSshConnectivity() GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference
	ReverseSshConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivity
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	State() *string
	StaticIpConnectivity() GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference
	StaticIpConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivity
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDatabaseMigrationServiceMigrationJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VpcPeeringConnectivity() GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference
	VpcPeeringConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity
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
	PutDumpFlags(value *GoogleDatabaseMigrationServiceMigrationJobDumpFlags)
	PutPerformanceConfig(value *GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig)
	PutReverseSshConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivity)
	PutStaticIpConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivity)
	PutTimeouts(value *GoogleDatabaseMigrationServiceMigrationJobTimeouts)
	PutVpcPeeringConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity)
	ResetDisplayName()
	ResetDumpFlags()
	ResetDumpPath()
	ResetDumpType()
	ResetId()
	ResetLabels()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceConfig()
	ResetProject()
	ResetReverseSshConnectivity()
	ResetStaticIpConnectivity()
	ResetTimeouts()
	ResetVpcPeeringConnectivity()
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

// The jsii proxy struct for GoogleDatabaseMigrationServiceMigrationJob
type jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpFlags() GoogleDatabaseMigrationServiceMigrationJobDumpFlagsOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobDumpFlagsOutputReference
	_jsii_.Get(
		j,
		"dumpFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpFlagsInput() *GoogleDatabaseMigrationServiceMigrationJobDumpFlags {
	var returns *GoogleDatabaseMigrationServiceMigrationJobDumpFlags
	_jsii_.Get(
		j,
		"dumpFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) DumpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Error() GoogleDatabaseMigrationServiceMigrationJobErrorList {
	var returns GoogleDatabaseMigrationServiceMigrationJobErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) MigrationJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) MigrationJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PerformanceConfig() GoogleDatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"performanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PerformanceConfigInput() *GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig {
	var returns *GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig
	_jsii_.Get(
		j,
		"performanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Phase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ReverseSshConnectivity() GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivityOutputReference
	_jsii_.Get(
		j,
		"reverseSshConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ReverseSshConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivity {
	var returns *GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivity
	_jsii_.Get(
		j,
		"reverseSshConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) StaticIpConnectivity() GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivityOutputReference
	_jsii_.Get(
		j,
		"staticIpConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) StaticIpConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivity {
	var returns *GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivity
	_jsii_.Get(
		j,
		"staticIpConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Timeouts() GoogleDatabaseMigrationServiceMigrationJobTimeoutsOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) VpcPeeringConnectivity() GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference {
	var returns GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivityOutputReference
	_jsii_.Get(
		j,
		"vpcPeeringConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) VpcPeeringConnectivityInput() *GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity {
	var returns *GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity
	_jsii_.Get(
		j,
		"vpcPeeringConnectivityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_migration_job google_database_migration_service_migration_job} Resource.
func NewGoogleDatabaseMigrationServiceMigrationJob(scope constructs.Construct, id *string, config *GoogleDatabaseMigrationServiceMigrationJobConfig) GoogleDatabaseMigrationServiceMigrationJob {
	_init_.Initialize()

	if err := validateNewGoogleDatabaseMigrationServiceMigrationJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_database_migration_service_migration_job google_database_migration_service_migration_job} Resource.
func NewGoogleDatabaseMigrationServiceMigrationJob_Override(g GoogleDatabaseMigrationServiceMigrationJob, scope constructs.Construct, id *string, config *GoogleDatabaseMigrationServiceMigrationJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetDumpPath(val *string) {
	if err := j.validateSetDumpPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dumpPath",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetDumpType(val *string) {
	if err := j.validateSetDumpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dumpType",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetMigrationJobId(val *string) {
	if err := j.validateSetMigrationJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationJobId",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDatabaseMigrationServiceMigrationJob resource upon running "cdktf plan <stack-name>".
func GoogleDatabaseMigrationServiceMigrationJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceMigrationJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
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
func GoogleDatabaseMigrationServiceMigrationJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceMigrationJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatabaseMigrationServiceMigrationJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceMigrationJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDatabaseMigrationServiceMigrationJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDatabaseMigrationServiceMigrationJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDatabaseMigrationServiceMigrationJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDatabaseMigrationServiceMigrationJob.GoogleDatabaseMigrationServiceMigrationJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutDumpFlags(value *GoogleDatabaseMigrationServiceMigrationJobDumpFlags) {
	if err := g.validatePutDumpFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDumpFlags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutPerformanceConfig(value *GoogleDatabaseMigrationServiceMigrationJobPerformanceConfig) {
	if err := g.validatePutPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPerformanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutReverseSshConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobReverseSshConnectivity) {
	if err := g.validatePutReverseSshConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReverseSshConnectivity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutStaticIpConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobStaticIpConnectivity) {
	if err := g.validatePutStaticIpConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticIpConnectivity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutTimeouts(value *GoogleDatabaseMigrationServiceMigrationJobTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) PutVpcPeeringConnectivity(value *GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity) {
	if err := g.validatePutVpcPeeringConnectivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcPeeringConnectivity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetDumpFlags() {
	_jsii_.InvokeVoid(
		g,
		"resetDumpFlags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetDumpPath() {
	_jsii_.InvokeVoid(
		g,
		"resetDumpPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetDumpType() {
	_jsii_.InvokeVoid(
		g,
		"resetDumpType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetPerformanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPerformanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetReverseSshConnectivity() {
	_jsii_.InvokeVoid(
		g,
		"resetReverseSshConnectivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetStaticIpConnectivity() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticIpConnectivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ResetVpcPeeringConnectivity() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcPeeringConnectivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatabaseMigrationServiceMigrationJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

