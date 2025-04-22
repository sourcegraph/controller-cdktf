package googlecolabruntimetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecolabruntimetemplate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template google_colab_runtime_template}.
type GoogleColabRuntimeTemplate interface {
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
	DataPersistentDiskSpec() GoogleColabRuntimeTemplateDataPersistentDiskSpecOutputReference
	DataPersistentDiskSpecInput() *GoogleColabRuntimeTemplateDataPersistentDiskSpec
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionSpec() GoogleColabRuntimeTemplateEncryptionSpecOutputReference
	EncryptionSpecInput() *GoogleColabRuntimeTemplateEncryptionSpec
	EucConfig() GoogleColabRuntimeTemplateEucConfigOutputReference
	EucConfigInput() *GoogleColabRuntimeTemplateEucConfig
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
	IdleShutdownConfig() GoogleColabRuntimeTemplateIdleShutdownConfigOutputReference
	IdleShutdownConfigInput() *GoogleColabRuntimeTemplateIdleShutdownConfig
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
	MachineSpec() GoogleColabRuntimeTemplateMachineSpecOutputReference
	MachineSpecInput() *GoogleColabRuntimeTemplateMachineSpec
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSpec() GoogleColabRuntimeTemplateNetworkSpecOutputReference
	NetworkSpecInput() *GoogleColabRuntimeTemplateNetworkSpec
	NetworkTags() *[]*string
	SetNetworkTags(val *[]*string)
	NetworkTagsInput() *[]*string
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
	ShieldedVmConfig() GoogleColabRuntimeTemplateShieldedVmConfigOutputReference
	ShieldedVmConfigInput() *GoogleColabRuntimeTemplateShieldedVmConfig
	SoftwareConfig() GoogleColabRuntimeTemplateSoftwareConfigOutputReference
	SoftwareConfigInput() *GoogleColabRuntimeTemplateSoftwareConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleColabRuntimeTemplateTimeoutsOutputReference
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
	PutDataPersistentDiskSpec(value *GoogleColabRuntimeTemplateDataPersistentDiskSpec)
	PutEncryptionSpec(value *GoogleColabRuntimeTemplateEncryptionSpec)
	PutEucConfig(value *GoogleColabRuntimeTemplateEucConfig)
	PutIdleShutdownConfig(value *GoogleColabRuntimeTemplateIdleShutdownConfig)
	PutMachineSpec(value *GoogleColabRuntimeTemplateMachineSpec)
	PutNetworkSpec(value *GoogleColabRuntimeTemplateNetworkSpec)
	PutShieldedVmConfig(value *GoogleColabRuntimeTemplateShieldedVmConfig)
	PutSoftwareConfig(value *GoogleColabRuntimeTemplateSoftwareConfig)
	PutTimeouts(value *GoogleColabRuntimeTemplateTimeouts)
	ResetDataPersistentDiskSpec()
	ResetDescription()
	ResetEncryptionSpec()
	ResetEucConfig()
	ResetId()
	ResetIdleShutdownConfig()
	ResetLabels()
	ResetMachineSpec()
	ResetName()
	ResetNetworkSpec()
	ResetNetworkTags()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetShieldedVmConfig()
	ResetSoftwareConfig()
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

// The jsii proxy struct for GoogleColabRuntimeTemplate
type jsiiProxy_GoogleColabRuntimeTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DataPersistentDiskSpec() GoogleColabRuntimeTemplateDataPersistentDiskSpecOutputReference {
	var returns GoogleColabRuntimeTemplateDataPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"dataPersistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DataPersistentDiskSpecInput() *GoogleColabRuntimeTemplateDataPersistentDiskSpec {
	var returns *GoogleColabRuntimeTemplateDataPersistentDiskSpec
	_jsii_.Get(
		j,
		"dataPersistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) EncryptionSpec() GoogleColabRuntimeTemplateEncryptionSpecOutputReference {
	var returns GoogleColabRuntimeTemplateEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) EncryptionSpecInput() *GoogleColabRuntimeTemplateEncryptionSpec {
	var returns *GoogleColabRuntimeTemplateEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) EucConfig() GoogleColabRuntimeTemplateEucConfigOutputReference {
	var returns GoogleColabRuntimeTemplateEucConfigOutputReference
	_jsii_.Get(
		j,
		"eucConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) EucConfigInput() *GoogleColabRuntimeTemplateEucConfig {
	var returns *GoogleColabRuntimeTemplateEucConfig
	_jsii_.Get(
		j,
		"eucConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) IdleShutdownConfig() GoogleColabRuntimeTemplateIdleShutdownConfigOutputReference {
	var returns GoogleColabRuntimeTemplateIdleShutdownConfigOutputReference
	_jsii_.Get(
		j,
		"idleShutdownConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) IdleShutdownConfigInput() *GoogleColabRuntimeTemplateIdleShutdownConfig {
	var returns *GoogleColabRuntimeTemplateIdleShutdownConfig
	_jsii_.Get(
		j,
		"idleShutdownConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) MachineSpec() GoogleColabRuntimeTemplateMachineSpecOutputReference {
	var returns GoogleColabRuntimeTemplateMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) MachineSpecInput() *GoogleColabRuntimeTemplateMachineSpec {
	var returns *GoogleColabRuntimeTemplateMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) NetworkSpec() GoogleColabRuntimeTemplateNetworkSpecOutputReference {
	var returns GoogleColabRuntimeTemplateNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) NetworkSpecInput() *GoogleColabRuntimeTemplateNetworkSpec {
	var returns *GoogleColabRuntimeTemplateNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) NetworkTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) NetworkTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) ShieldedVmConfig() GoogleColabRuntimeTemplateShieldedVmConfigOutputReference {
	var returns GoogleColabRuntimeTemplateShieldedVmConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedVmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) ShieldedVmConfigInput() *GoogleColabRuntimeTemplateShieldedVmConfig {
	var returns *GoogleColabRuntimeTemplateShieldedVmConfig
	_jsii_.Get(
		j,
		"shieldedVmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) SoftwareConfig() GoogleColabRuntimeTemplateSoftwareConfigOutputReference {
	var returns GoogleColabRuntimeTemplateSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) SoftwareConfigInput() *GoogleColabRuntimeTemplateSoftwareConfig {
	var returns *GoogleColabRuntimeTemplateSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) Timeouts() GoogleColabRuntimeTemplateTimeoutsOutputReference {
	var returns GoogleColabRuntimeTemplateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template google_colab_runtime_template} Resource.
func NewGoogleColabRuntimeTemplate(scope constructs.Construct, id *string, config *GoogleColabRuntimeTemplateConfig) GoogleColabRuntimeTemplate {
	_init_.Initialize()

	if err := validateNewGoogleColabRuntimeTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleColabRuntimeTemplate{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template google_colab_runtime_template} Resource.
func NewGoogleColabRuntimeTemplate_Override(g GoogleColabRuntimeTemplate, scope constructs.Construct, id *string, config *GoogleColabRuntimeTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetNetworkTags(val *[]*string) {
	if err := j.validateSetNetworkTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTags",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleColabRuntimeTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleColabRuntimeTemplate resource upon running "cdktf plan <stack-name>".
func GoogleColabRuntimeTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleColabRuntimeTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
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
func GoogleColabRuntimeTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabRuntimeTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleColabRuntimeTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabRuntimeTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleColabRuntimeTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleColabRuntimeTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleColabRuntimeTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleColabRuntimeTemplate.GoogleColabRuntimeTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleColabRuntimeTemplate) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutDataPersistentDiskSpec(value *GoogleColabRuntimeTemplateDataPersistentDiskSpec) {
	if err := g.validatePutDataPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutEncryptionSpec(value *GoogleColabRuntimeTemplateEncryptionSpec) {
	if err := g.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutEucConfig(value *GoogleColabRuntimeTemplateEucConfig) {
	if err := g.validatePutEucConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEucConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutIdleShutdownConfig(value *GoogleColabRuntimeTemplateIdleShutdownConfig) {
	if err := g.validatePutIdleShutdownConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdleShutdownConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutMachineSpec(value *GoogleColabRuntimeTemplateMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutNetworkSpec(value *GoogleColabRuntimeTemplateNetworkSpec) {
	if err := g.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutShieldedVmConfig(value *GoogleColabRuntimeTemplateShieldedVmConfig) {
	if err := g.validatePutShieldedVmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedVmConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutSoftwareConfig(value *GoogleColabRuntimeTemplateSoftwareConfig) {
	if err := g.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) PutTimeouts(value *GoogleColabRuntimeTemplateTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetDataPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetDataPersistentDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetEucConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEucConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetIdleShutdownConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIdleShutdownConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetNetworkTags() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetShieldedVmConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedVmConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabRuntimeTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

