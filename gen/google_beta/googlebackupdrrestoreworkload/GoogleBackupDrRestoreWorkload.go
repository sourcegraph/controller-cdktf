package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload google_backup_dr_restore_workload}.
type GoogleBackupDrRestoreWorkload interface {
	cdktf.TerraformResource
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	BackupVaultId() *string
	SetBackupVaultId(val *string)
	BackupVaultIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClearOverridesFieldMask() *string
	SetClearOverridesFieldMask(val *string)
	ClearOverridesFieldMaskInput() *string
	ComputeInstanceRestoreProperties() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
	ComputeInstanceRestorePropertiesInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties
	ComputeInstanceTargetEnvironment() GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference
	ComputeInstanceTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment
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
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
	DeleteRestoredInstance() interface{}
	SetDeleteRestoredInstance(val interface{})
	DeleteRestoredInstanceInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskRestoreProperties() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
	DiskRestorePropertiesInput() *GoogleBackupDrRestoreWorkloadDiskRestoreProperties
	DiskTargetEnvironment() GoogleBackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference
	DiskTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadDiskTargetEnvironment
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RegionDiskTargetEnvironment() GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference
	RegionDiskTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment
	RequestId() *string
	SetRequestId(val *string)
	RequestIdInput() *string
	TargetResource() GoogleBackupDrRestoreWorkloadTargetResourceList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBackupDrRestoreWorkloadTimeoutsOutputReference
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
	PutComputeInstanceRestoreProperties(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties)
	PutComputeInstanceTargetEnvironment(value *GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment)
	PutDiskRestoreProperties(value *GoogleBackupDrRestoreWorkloadDiskRestoreProperties)
	PutDiskTargetEnvironment(value *GoogleBackupDrRestoreWorkloadDiskTargetEnvironment)
	PutRegionDiskTargetEnvironment(value *GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment)
	PutTimeouts(value *GoogleBackupDrRestoreWorkloadTimeouts)
	ResetClearOverridesFieldMask()
	ResetComputeInstanceRestoreProperties()
	ResetComputeInstanceTargetEnvironment()
	ResetDeleteRestoredInstance()
	ResetDiskRestoreProperties()
	ResetDiskTargetEnvironment()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegionDiskTargetEnvironment()
	ResetRequestId()
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

// The jsii proxy struct for GoogleBackupDrRestoreWorkload
type jsiiProxy_GoogleBackupDrRestoreWorkload struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) BackupVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) BackupVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ClearOverridesFieldMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearOverridesFieldMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ClearOverridesFieldMaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearOverridesFieldMaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ComputeInstanceRestoreProperties() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
	_jsii_.Get(
		j,
		"computeInstanceRestoreProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ComputeInstanceRestorePropertiesInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties
	_jsii_.Get(
		j,
		"computeInstanceRestorePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ComputeInstanceTargetEnvironment() GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"computeInstanceTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ComputeInstanceTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment
	_jsii_.Get(
		j,
		"computeInstanceTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DeleteRestoredInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRestoredInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DeleteRestoredInstanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRestoredInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DiskRestoreProperties() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference {
	var returns GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
	_jsii_.Get(
		j,
		"diskRestoreProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DiskRestorePropertiesInput() *GoogleBackupDrRestoreWorkloadDiskRestoreProperties {
	var returns *GoogleBackupDrRestoreWorkloadDiskRestoreProperties
	_jsii_.Get(
		j,
		"diskRestorePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DiskTargetEnvironment() GoogleBackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference {
	var returns GoogleBackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"diskTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) DiskTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadDiskTargetEnvironment {
	var returns *GoogleBackupDrRestoreWorkloadDiskTargetEnvironment
	_jsii_.Get(
		j,
		"diskTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) RegionDiskTargetEnvironment() GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference {
	var returns GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"regionDiskTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) RegionDiskTargetEnvironmentInput() *GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment {
	var returns *GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment
	_jsii_.Get(
		j,
		"regionDiskTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) RequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) RequestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) TargetResource() GoogleBackupDrRestoreWorkloadTargetResourceList {
	var returns GoogleBackupDrRestoreWorkloadTargetResourceList
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) Timeouts() GoogleBackupDrRestoreWorkloadTimeoutsOutputReference {
	var returns GoogleBackupDrRestoreWorkloadTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload google_backup_dr_restore_workload} Resource.
func NewGoogleBackupDrRestoreWorkload(scope constructs.Construct, id *string, config *GoogleBackupDrRestoreWorkloadConfig) GoogleBackupDrRestoreWorkload {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkload{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload google_backup_dr_restore_workload} Resource.
func NewGoogleBackupDrRestoreWorkload_Override(g GoogleBackupDrRestoreWorkload, scope constructs.Construct, id *string, config *GoogleBackupDrRestoreWorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetBackupVaultId(val *string) {
	if err := j.validateSetBackupVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVaultId",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetClearOverridesFieldMask(val *string) {
	if err := j.validateSetClearOverridesFieldMaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clearOverridesFieldMask",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetDataSourceId(val *string) {
	if err := j.validateSetDataSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetDeleteRestoredInstance(val interface{}) {
	if err := j.validateSetDeleteRestoredInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRestoredInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkload)SetRequestId(val *string) {
	if err := j.validateSetRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBackupDrRestoreWorkload resource upon running "cdktf plan <stack-name>".
func GoogleBackupDrRestoreWorkload_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBackupDrRestoreWorkload_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
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
func GoogleBackupDrRestoreWorkload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrRestoreWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBackupDrRestoreWorkload_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrRestoreWorkload_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBackupDrRestoreWorkload_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrRestoreWorkload_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBackupDrRestoreWorkload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutComputeInstanceRestoreProperties(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties) {
	if err := g.validatePutComputeInstanceRestorePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putComputeInstanceRestoreProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutComputeInstanceTargetEnvironment(value *GoogleBackupDrRestoreWorkloadComputeInstanceTargetEnvironment) {
	if err := g.validatePutComputeInstanceTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putComputeInstanceTargetEnvironment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutDiskRestoreProperties(value *GoogleBackupDrRestoreWorkloadDiskRestoreProperties) {
	if err := g.validatePutDiskRestorePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskRestoreProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutDiskTargetEnvironment(value *GoogleBackupDrRestoreWorkloadDiskTargetEnvironment) {
	if err := g.validatePutDiskTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskTargetEnvironment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutRegionDiskTargetEnvironment(value *GoogleBackupDrRestoreWorkloadRegionDiskTargetEnvironment) {
	if err := g.validatePutRegionDiskTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRegionDiskTargetEnvironment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) PutTimeouts(value *GoogleBackupDrRestoreWorkloadTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetClearOverridesFieldMask() {
	_jsii_.InvokeVoid(
		g,
		"resetClearOverridesFieldMask",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetComputeInstanceRestoreProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeInstanceRestoreProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetComputeInstanceTargetEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeInstanceTargetEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetDeleteRestoredInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteRestoredInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetDiskRestoreProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskRestoreProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetDiskTargetEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskTargetEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetRegionDiskTargetEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetRegionDiskTargetEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetRequestId() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

