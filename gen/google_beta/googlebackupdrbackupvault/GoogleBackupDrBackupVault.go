package googlebackupdrbackupvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrbackupvault/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_vault google_backup_dr_backup_vault}.
type GoogleBackupDrBackupVault interface {
	cdktf.TerraformResource
	AccessRestriction() *string
	SetAccessRestriction(val *string)
	AccessRestrictionInput() *string
	AllowMissing() interface{}
	SetAllowMissing(val interface{})
	AllowMissingInput() interface{}
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BackupCount() *string
	BackupMinimumEnforcedRetentionDuration() *string
	SetBackupMinimumEnforcedRetentionDuration(val *string)
	BackupMinimumEnforcedRetentionDurationInput() *string
	BackupVaultId() *string
	SetBackupVaultId(val *string)
	BackupVaultIdInput() *string
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
	Deletable() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	EffectiveTime() *string
	SetEffectiveTime(val *string)
	EffectiveTimeInput() *string
	Etag() *string
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
	ForceUpdate() interface{}
	SetForceUpdate(val interface{})
	ForceUpdateInput() interface{}
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
	IgnoreBackupPlanReferences() interface{}
	SetIgnoreBackupPlanReferences(val interface{})
	IgnoreBackupPlanReferencesInput() interface{}
	IgnoreInactiveDatasources() interface{}
	SetIgnoreInactiveDatasources(val interface{})
	IgnoreInactiveDatasourcesInput() interface{}
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
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	ServiceAccount() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBackupDrBackupVaultTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalStoredBytes() *string
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
	PutTimeouts(value *GoogleBackupDrBackupVaultTimeouts)
	ResetAccessRestriction()
	ResetAllowMissing()
	ResetAnnotations()
	ResetDescription()
	ResetEffectiveTime()
	ResetForceDelete()
	ResetForceUpdate()
	ResetId()
	ResetIgnoreBackupPlanReferences()
	ResetIgnoreInactiveDatasources()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for GoogleBackupDrBackupVault
type jsiiProxy_GoogleBackupDrBackupVault struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) AccessRestriction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) AccessRestrictionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) AllowMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) AllowMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) BackupCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) BackupMinimumEnforcedRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) BackupMinimumEnforcedRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) BackupVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) BackupVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Deletable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) EffectiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) EffectiveTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ForceUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ForceUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) IgnoreBackupPlanReferences() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreBackupPlanReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) IgnoreBackupPlanReferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreBackupPlanReferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) IgnoreInactiveDatasources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreInactiveDatasources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) IgnoreInactiveDatasourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreInactiveDatasourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Timeouts() GoogleBackupDrBackupVaultTimeoutsOutputReference {
	var returns GoogleBackupDrBackupVaultTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) TotalStoredBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalStoredBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupVault) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_vault google_backup_dr_backup_vault} Resource.
func NewGoogleBackupDrBackupVault(scope constructs.Construct, id *string, config *GoogleBackupDrBackupVaultConfig) GoogleBackupDrBackupVault {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrBackupVaultParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrBackupVault{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_backup_dr_backup_vault google_backup_dr_backup_vault} Resource.
func NewGoogleBackupDrBackupVault_Override(g GoogleBackupDrBackupVault, scope constructs.Construct, id *string, config *GoogleBackupDrBackupVaultConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetAccessRestriction(val *string) {
	if err := j.validateSetAccessRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessRestriction",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetAllowMissing(val interface{}) {
	if err := j.validateSetAllowMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMissing",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetBackupMinimumEnforcedRetentionDuration(val *string) {
	if err := j.validateSetBackupMinimumEnforcedRetentionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinimumEnforcedRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetBackupVaultId(val *string) {
	if err := j.validateSetBackupVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVaultId",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetEffectiveTime(val *string) {
	if err := j.validateSetEffectiveTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveTime",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetForceUpdate(val interface{}) {
	if err := j.validateSetForceUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdate",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetIgnoreBackupPlanReferences(val interface{}) {
	if err := j.validateSetIgnoreBackupPlanReferencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreBackupPlanReferences",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetIgnoreInactiveDatasources(val interface{}) {
	if err := j.validateSetIgnoreInactiveDatasourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreInactiveDatasources",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupVault)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBackupDrBackupVault resource upon running "cdktf plan <stack-name>".
func GoogleBackupDrBackupVault_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBackupDrBackupVault_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
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
func GoogleBackupDrBackupVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrBackupVault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBackupDrBackupVault_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrBackupVault_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBackupDrBackupVault_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBackupDrBackupVault_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBackupDrBackupVault_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBackupDrBackupVault.GoogleBackupDrBackupVault",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrBackupVault) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) PutTimeouts(value *GoogleBackupDrBackupVaultTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetAccessRestriction() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessRestriction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetAllowMissing() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowMissing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetEffectiveTime() {
	_jsii_.InvokeVoid(
		g,
		"resetEffectiveTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetForceDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetForceUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetForceUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetIgnoreBackupPlanReferences() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreBackupPlanReferences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetIgnoreInactiveDatasources() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreInactiveDatasources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupVault) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

