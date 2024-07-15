package googlebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerytable/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table google_bigquery_table}.
type GoogleBigqueryTable interface {
	cdktf.TerraformResource
	AllowResourceTagsOnDeletion() interface{}
	SetAllowResourceTagsOnDeletion(val interface{})
	AllowResourceTagsOnDeletionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clustering() *[]*string
	SetClustering(val *[]*string)
	ClusteringInput() *[]*string
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
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionConfiguration() GoogleBigqueryTableEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() *GoogleBigqueryTableEncryptionConfiguration
	Etag() *string
	ExpirationTime() *float64
	SetExpirationTime(val *float64)
	ExpirationTimeInput() *float64
	ExternalDataConfiguration() GoogleBigqueryTableExternalDataConfigurationOutputReference
	ExternalDataConfigurationInput() *GoogleBigqueryTableExternalDataConfiguration
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	MaterializedView() GoogleBigqueryTableMaterializedViewOutputReference
	MaterializedViewInput() *GoogleBigqueryTableMaterializedView
	MaxStaleness() *string
	SetMaxStaleness(val *string)
	MaxStalenessInput() *string
	// The tree node.
	Node() constructs.Node
	NumBytes() *float64
	NumLongTermBytes() *float64
	NumRows() *float64
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
	RangePartitioning() GoogleBigqueryTableRangePartitioningOutputReference
	RangePartitioningInput() *GoogleBigqueryTableRangePartitioning
	// Experimental.
	RawOverrides() interface{}
	RequirePartitionFilter() interface{}
	SetRequirePartitionFilter(val interface{})
	RequirePartitionFilterInput() interface{}
	ResourceTags() *map[string]*string
	SetResourceTags(val *map[string]*string)
	ResourceTagsInput() *map[string]*string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SelfLink() *string
	TableConstraints() GoogleBigqueryTableTableConstraintsOutputReference
	TableConstraintsInput() *GoogleBigqueryTableTableConstraints
	TableId() *string
	SetTableId(val *string)
	TableIdInput() *string
	TableReplicationInfo() GoogleBigqueryTableTableReplicationInfoOutputReference
	TableReplicationInfoInput() *GoogleBigqueryTableTableReplicationInfo
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimePartitioning() GoogleBigqueryTableTimePartitioningOutputReference
	TimePartitioningInput() *GoogleBigqueryTableTimePartitioning
	Type() *string
	View() GoogleBigqueryTableViewOutputReference
	ViewInput() *GoogleBigqueryTableView
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
	PutEncryptionConfiguration(value *GoogleBigqueryTableEncryptionConfiguration)
	PutExternalDataConfiguration(value *GoogleBigqueryTableExternalDataConfiguration)
	PutMaterializedView(value *GoogleBigqueryTableMaterializedView)
	PutRangePartitioning(value *GoogleBigqueryTableRangePartitioning)
	PutTableConstraints(value *GoogleBigqueryTableTableConstraints)
	PutTableReplicationInfo(value *GoogleBigqueryTableTableReplicationInfo)
	PutTimePartitioning(value *GoogleBigqueryTableTimePartitioning)
	PutView(value *GoogleBigqueryTableView)
	ResetAllowResourceTagsOnDeletion()
	ResetClustering()
	ResetDeletionProtection()
	ResetDescription()
	ResetEncryptionConfiguration()
	ResetExpirationTime()
	ResetExternalDataConfiguration()
	ResetFriendlyName()
	ResetId()
	ResetLabels()
	ResetMaterializedView()
	ResetMaxStaleness()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRangePartitioning()
	ResetRequirePartitionFilter()
	ResetResourceTags()
	ResetSchema()
	ResetTableConstraints()
	ResetTableReplicationInfo()
	ResetTimePartitioning()
	ResetView()
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

// The jsii proxy struct for GoogleBigqueryTable
type jsiiProxy_GoogleBigqueryTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryTable) AllowResourceTagsOnDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowResourceTagsOnDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) AllowResourceTagsOnDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowResourceTagsOnDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Clustering() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clustering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ClusteringInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) EncryptionConfiguration() GoogleBigqueryTableEncryptionConfigurationOutputReference {
	var returns GoogleBigqueryTableEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) EncryptionConfigurationInput() *GoogleBigqueryTableEncryptionConfiguration {
	var returns *GoogleBigqueryTableEncryptionConfiguration
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ExpirationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ExpirationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ExternalDataConfiguration() GoogleBigqueryTableExternalDataConfigurationOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationOutputReference
	_jsii_.Get(
		j,
		"externalDataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ExternalDataConfigurationInput() *GoogleBigqueryTableExternalDataConfiguration {
	var returns *GoogleBigqueryTableExternalDataConfiguration
	_jsii_.Get(
		j,
		"externalDataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) MaterializedView() GoogleBigqueryTableMaterializedViewOutputReference {
	var returns GoogleBigqueryTableMaterializedViewOutputReference
	_jsii_.Get(
		j,
		"materializedView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) MaterializedViewInput() *GoogleBigqueryTableMaterializedView {
	var returns *GoogleBigqueryTableMaterializedView
	_jsii_.Get(
		j,
		"materializedViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) MaxStaleness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxStaleness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) MaxStalenessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxStalenessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) NumBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) NumLongTermBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLongTermBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) NumRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) RangePartitioning() GoogleBigqueryTableRangePartitioningOutputReference {
	var returns GoogleBigqueryTableRangePartitioningOutputReference
	_jsii_.Get(
		j,
		"rangePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) RangePartitioningInput() *GoogleBigqueryTableRangePartitioning {
	var returns *GoogleBigqueryTableRangePartitioning
	_jsii_.Get(
		j,
		"rangePartitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) RequirePartitionFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) RequirePartitionFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableConstraints() GoogleBigqueryTableTableConstraintsOutputReference {
	var returns GoogleBigqueryTableTableConstraintsOutputReference
	_jsii_.Get(
		j,
		"tableConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableConstraintsInput() *GoogleBigqueryTableTableConstraints {
	var returns *GoogleBigqueryTableTableConstraints
	_jsii_.Get(
		j,
		"tableConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableReplicationInfo() GoogleBigqueryTableTableReplicationInfoOutputReference {
	var returns GoogleBigqueryTableTableReplicationInfoOutputReference
	_jsii_.Get(
		j,
		"tableReplicationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TableReplicationInfoInput() *GoogleBigqueryTableTableReplicationInfo {
	var returns *GoogleBigqueryTableTableReplicationInfo
	_jsii_.Get(
		j,
		"tableReplicationInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TimePartitioning() GoogleBigqueryTableTimePartitioningOutputReference {
	var returns GoogleBigqueryTableTimePartitioningOutputReference
	_jsii_.Get(
		j,
		"timePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) TimePartitioningInput() *GoogleBigqueryTableTimePartitioning {
	var returns *GoogleBigqueryTableTimePartitioning
	_jsii_.Get(
		j,
		"timePartitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) View() GoogleBigqueryTableViewOutputReference {
	var returns GoogleBigqueryTableViewOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTable) ViewInput() *GoogleBigqueryTableView {
	var returns *GoogleBigqueryTableView
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table google_bigquery_table} Resource.
func NewGoogleBigqueryTable(scope constructs.Construct, id *string, config *GoogleBigqueryTableConfig) GoogleBigqueryTable {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryTable{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table google_bigquery_table} Resource.
func NewGoogleBigqueryTable_Override(g GoogleBigqueryTable, scope constructs.Construct, id *string, config *GoogleBigqueryTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetAllowResourceTagsOnDeletion(val interface{}) {
	if err := j.validateSetAllowResourceTagsOnDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowResourceTagsOnDeletion",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetClustering(val *[]*string) {
	if err := j.validateSetClusteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clustering",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetExpirationTime(val *float64) {
	if err := j.validateSetExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTime",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetFriendlyName(val *string) {
	if err := j.validateSetFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetMaxStaleness(val *string) {
	if err := j.validateSetMaxStalenessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStaleness",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetRequirePartitionFilter(val interface{}) {
	if err := j.validateSetRequirePartitionFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePartitionFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTable)SetTableId(val *string) {
	if err := j.validateSetTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigqueryTable resource upon running "cdktf plan <stack-name>".
func GoogleBigqueryTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
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
func GoogleBigqueryTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTable) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutEncryptionConfiguration(value *GoogleBigqueryTableEncryptionConfiguration) {
	if err := g.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutExternalDataConfiguration(value *GoogleBigqueryTableExternalDataConfiguration) {
	if err := g.validatePutExternalDataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalDataConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutMaterializedView(value *GoogleBigqueryTableMaterializedView) {
	if err := g.validatePutMaterializedViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaterializedView",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutRangePartitioning(value *GoogleBigqueryTableRangePartitioning) {
	if err := g.validatePutRangePartitioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRangePartitioning",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutTableConstraints(value *GoogleBigqueryTableTableConstraints) {
	if err := g.validatePutTableConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTableConstraints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutTableReplicationInfo(value *GoogleBigqueryTableTableReplicationInfo) {
	if err := g.validatePutTableReplicationInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTableReplicationInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutTimePartitioning(value *GoogleBigqueryTableTimePartitioning) {
	if err := g.validatePutTimePartitioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimePartitioning",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) PutView(value *GoogleBigqueryTableView) {
	if err := g.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putView",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetAllowResourceTagsOnDeletion() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowResourceTagsOnDeletion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetClustering() {
	_jsii_.InvokeVoid(
		g,
		"resetClustering",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetExpirationTime() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetExternalDataConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalDataConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		g,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetMaterializedView() {
	_jsii_.InvokeVoid(
		g,
		"resetMaterializedView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetMaxStaleness() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxStaleness",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetRangePartitioning() {
	_jsii_.InvokeVoid(
		g,
		"resetRangePartitioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetRequirePartitionFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetRequirePartitionFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetResourceTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetTableConstraints() {
	_jsii_.InvokeVoid(
		g,
		"resetTableConstraints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetTableReplicationInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetTableReplicationInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetTimePartitioning() {
	_jsii_.InvokeVoid(
		g,
		"resetTimePartitioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) ResetView() {
	_jsii_.InvokeVoid(
		g,
		"resetView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

