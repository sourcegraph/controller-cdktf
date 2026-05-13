package googlestorageinsightsdatasetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestorageinsightsdatasetconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config google_storage_insights_dataset_config}.
type GoogleStorageInsightsDatasetConfig interface {
	cdktf.TerraformResource
	ActivityDataRetentionPeriodDays() *float64
	SetActivityDataRetentionPeriodDays(val *float64)
	ActivityDataRetentionPeriodDaysInput() *float64
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
	DatasetConfigId() *string
	SetDatasetConfigId(val *string)
	DatasetConfigIdInput() *string
	DatasetConfigState() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExcludeCloudStorageBuckets() GoogleStorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference
	ExcludeCloudStorageBucketsInput() *GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets
	ExcludeCloudStorageLocations() GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference
	ExcludeCloudStorageLocationsInput() *GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations
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
	Identity() GoogleStorageInsightsDatasetConfigIdentityOutputReference
	IdentityInput() *GoogleStorageInsightsDatasetConfigIdentity
	IdInput() *string
	IncludeCloudStorageBuckets() GoogleStorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference
	IncludeCloudStorageBucketsInput() *GoogleStorageInsightsDatasetConfigIncludeCloudStorageBuckets
	IncludeCloudStorageLocations() GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference
	IncludeCloudStorageLocationsInput() *GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations
	IncludeNewlyCreatedBuckets() interface{}
	SetIncludeNewlyCreatedBuckets(val interface{})
	IncludeNewlyCreatedBucketsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Link() GoogleStorageInsightsDatasetConfigLinkList
	LinkDataset() interface{}
	SetLinkDataset(val interface{})
	LinkDatasetInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OrganizationNumber() *string
	SetOrganizationNumber(val *string)
	OrganizationNumberInput() *string
	OrganizationScope() interface{}
	SetOrganizationScope(val interface{})
	OrganizationScopeInput() interface{}
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
	RetentionPeriodDays() *float64
	SetRetentionPeriodDays(val *float64)
	RetentionPeriodDaysInput() *float64
	SourceFolders() GoogleStorageInsightsDatasetConfigSourceFoldersOutputReference
	SourceFoldersInput() *GoogleStorageInsightsDatasetConfigSourceFolders
	SourceProjects() GoogleStorageInsightsDatasetConfigSourceProjectsOutputReference
	SourceProjectsInput() *GoogleStorageInsightsDatasetConfigSourceProjects
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleStorageInsightsDatasetConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
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
	PutExcludeCloudStorageBuckets(value *GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets)
	PutExcludeCloudStorageLocations(value *GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations)
	PutIdentity(value *GoogleStorageInsightsDatasetConfigIdentity)
	PutIncludeCloudStorageBuckets(value *GoogleStorageInsightsDatasetConfigIncludeCloudStorageBuckets)
	PutIncludeCloudStorageLocations(value *GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations)
	PutSourceFolders(value *GoogleStorageInsightsDatasetConfigSourceFolders)
	PutSourceProjects(value *GoogleStorageInsightsDatasetConfigSourceProjects)
	PutTimeouts(value *GoogleStorageInsightsDatasetConfigTimeouts)
	ResetActivityDataRetentionPeriodDays()
	ResetDescription()
	ResetExcludeCloudStorageBuckets()
	ResetExcludeCloudStorageLocations()
	ResetId()
	ResetIncludeCloudStorageBuckets()
	ResetIncludeCloudStorageLocations()
	ResetIncludeNewlyCreatedBuckets()
	ResetLinkDataset()
	ResetOrganizationNumber()
	ResetOrganizationScope()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSourceFolders()
	ResetSourceProjects()
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

// The jsii proxy struct for GoogleStorageInsightsDatasetConfig
type jsiiProxy_GoogleStorageInsightsDatasetConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ActivityDataRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activityDataRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ActivityDataRetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activityDataRetentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) DatasetConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) DatasetConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) DatasetConfigState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ExcludeCloudStorageBuckets() GoogleStorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"excludeCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ExcludeCloudStorageBucketsInput() *GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets {
	var returns *GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets
	_jsii_.Get(
		j,
		"excludeCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ExcludeCloudStorageLocations() GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"excludeCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ExcludeCloudStorageLocationsInput() *GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations {
	var returns *GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations
	_jsii_.Get(
		j,
		"excludeCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Identity() GoogleStorageInsightsDatasetConfigIdentityOutputReference {
	var returns GoogleStorageInsightsDatasetConfigIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IdentityInput() *GoogleStorageInsightsDatasetConfigIdentity {
	var returns *GoogleStorageInsightsDatasetConfigIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeCloudStorageBuckets() GoogleStorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"includeCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeCloudStorageBucketsInput() *GoogleStorageInsightsDatasetConfigIncludeCloudStorageBuckets {
	var returns *GoogleStorageInsightsDatasetConfigIncludeCloudStorageBuckets
	_jsii_.Get(
		j,
		"includeCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeCloudStorageLocations() GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"includeCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeCloudStorageLocationsInput() *GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations {
	var returns *GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations
	_jsii_.Get(
		j,
		"includeCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeNewlyCreatedBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNewlyCreatedBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) IncludeNewlyCreatedBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNewlyCreatedBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Link() GoogleStorageInsightsDatasetConfigLinkList {
	var returns GoogleStorageInsightsDatasetConfigLinkList
	_jsii_.Get(
		j,
		"link",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) LinkDataset() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) LinkDatasetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) OrganizationNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) OrganizationNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) OrganizationScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) OrganizationScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) RetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) RetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) SourceFolders() GoogleStorageInsightsDatasetConfigSourceFoldersOutputReference {
	var returns GoogleStorageInsightsDatasetConfigSourceFoldersOutputReference
	_jsii_.Get(
		j,
		"sourceFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) SourceFoldersInput() *GoogleStorageInsightsDatasetConfigSourceFolders {
	var returns *GoogleStorageInsightsDatasetConfigSourceFolders
	_jsii_.Get(
		j,
		"sourceFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) SourceProjects() GoogleStorageInsightsDatasetConfigSourceProjectsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigSourceProjectsOutputReference
	_jsii_.Get(
		j,
		"sourceProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) SourceProjectsInput() *GoogleStorageInsightsDatasetConfigSourceProjects {
	var returns *GoogleStorageInsightsDatasetConfigSourceProjects
	_jsii_.Get(
		j,
		"sourceProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Timeouts() GoogleStorageInsightsDatasetConfigTimeoutsOutputReference {
	var returns GoogleStorageInsightsDatasetConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config google_storage_insights_dataset_config} Resource.
func NewGoogleStorageInsightsDatasetConfig(scope constructs.Construct, id *string, config *GoogleStorageInsightsDatasetConfigConfig) GoogleStorageInsightsDatasetConfig {
	_init_.Initialize()

	if err := validateNewGoogleStorageInsightsDatasetConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageInsightsDatasetConfig{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_insights_dataset_config google_storage_insights_dataset_config} Resource.
func NewGoogleStorageInsightsDatasetConfig_Override(g GoogleStorageInsightsDatasetConfig, scope constructs.Construct, id *string, config *GoogleStorageInsightsDatasetConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetActivityDataRetentionPeriodDays(val *float64) {
	if err := j.validateSetActivityDataRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activityDataRetentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetDatasetConfigId(val *string) {
	if err := j.validateSetDatasetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetConfigId",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetIncludeNewlyCreatedBuckets(val interface{}) {
	if err := j.validateSetIncludeNewlyCreatedBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNewlyCreatedBuckets",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetLinkDataset(val interface{}) {
	if err := j.validateSetLinkDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDataset",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetOrganizationNumber(val *string) {
	if err := j.validateSetOrganizationNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationNumber",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetOrganizationScope(val interface{}) {
	if err := j.validateSetOrganizationScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationScope",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsDatasetConfig)SetRetentionPeriodDays(val *float64) {
	if err := j.validateSetRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodDays",
		val,
	)
}

// Generates CDKTF code for importing a GoogleStorageInsightsDatasetConfig resource upon running "cdktf plan <stack-name>".
func GoogleStorageInsightsDatasetConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleStorageInsightsDatasetConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
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
func GoogleStorageInsightsDatasetConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageInsightsDatasetConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleStorageInsightsDatasetConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageInsightsDatasetConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleStorageInsightsDatasetConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleStorageInsightsDatasetConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleStorageInsightsDatasetConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleStorageInsightsDatasetConfig.GoogleStorageInsightsDatasetConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutExcludeCloudStorageBuckets(value *GoogleStorageInsightsDatasetConfigExcludeCloudStorageBuckets) {
	if err := g.validatePutExcludeCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutExcludeCloudStorageLocations(value *GoogleStorageInsightsDatasetConfigExcludeCloudStorageLocations) {
	if err := g.validatePutExcludeCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeCloudStorageLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutIdentity(value *GoogleStorageInsightsDatasetConfigIdentity) {
	if err := g.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdentity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutIncludeCloudStorageBuckets(value *GoogleStorageInsightsDatasetConfigIncludeCloudStorageBuckets) {
	if err := g.validatePutIncludeCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutIncludeCloudStorageLocations(value *GoogleStorageInsightsDatasetConfigIncludeCloudStorageLocations) {
	if err := g.validatePutIncludeCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeCloudStorageLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutSourceFolders(value *GoogleStorageInsightsDatasetConfigSourceFolders) {
	if err := g.validatePutSourceFoldersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceFolders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutSourceProjects(value *GoogleStorageInsightsDatasetConfigSourceProjects) {
	if err := g.validatePutSourceProjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceProjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) PutTimeouts(value *GoogleStorageInsightsDatasetConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetActivityDataRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		g,
		"resetActivityDataRetentionPeriodDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetExcludeCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeCloudStorageBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetExcludeCloudStorageLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeCloudStorageLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetIncludeCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeCloudStorageBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetIncludeCloudStorageLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeCloudStorageLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetIncludeNewlyCreatedBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeNewlyCreatedBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetLinkDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetOrganizationNumber() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationNumber",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetOrganizationScope() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetSourceFolders() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceFolders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetSourceProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsDatasetConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

