package googlednsmanagedzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlednsmanagedzone/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dns_managed_zone google_dns_managed_zone}.
type GoogleDnsManagedZone interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudLoggingConfig() GoogleDnsManagedZoneCloudLoggingConfigOutputReference
	CloudLoggingConfigInput() *GoogleDnsManagedZoneCloudLoggingConfig
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
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	DnssecConfig() GoogleDnsManagedZoneDnssecConfigOutputReference
	DnssecConfigInput() *GoogleDnsManagedZoneDnssecConfig
	EffectiveLabels() cdktf.StringMap
	ForceDestroy() interface{}
	SetForceDestroy(val interface{})
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardingConfig() GoogleDnsManagedZoneForwardingConfigOutputReference
	ForwardingConfigInput() *GoogleDnsManagedZoneForwardingConfig
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
	ManagedZoneId() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NameServers() *[]*string
	// The tree node.
	Node() constructs.Node
	PeeringConfig() GoogleDnsManagedZonePeeringConfigOutputReference
	PeeringConfigInput() *GoogleDnsManagedZonePeeringConfig
	PrivateVisibilityConfig() GoogleDnsManagedZonePrivateVisibilityConfigOutputReference
	PrivateVisibilityConfigInput() *GoogleDnsManagedZonePrivateVisibilityConfig
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
	ReverseLookup() interface{}
	SetReverseLookup(val interface{})
	ReverseLookupInput() interface{}
	ServiceDirectoryConfig() GoogleDnsManagedZoneServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *GoogleDnsManagedZoneServiceDirectoryConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDnsManagedZoneTimeoutsOutputReference
	TimeoutsInput() interface{}
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
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
	PutCloudLoggingConfig(value *GoogleDnsManagedZoneCloudLoggingConfig)
	PutDnssecConfig(value *GoogleDnsManagedZoneDnssecConfig)
	PutForwardingConfig(value *GoogleDnsManagedZoneForwardingConfig)
	PutPeeringConfig(value *GoogleDnsManagedZonePeeringConfig)
	PutPrivateVisibilityConfig(value *GoogleDnsManagedZonePrivateVisibilityConfig)
	PutServiceDirectoryConfig(value *GoogleDnsManagedZoneServiceDirectoryConfig)
	PutTimeouts(value *GoogleDnsManagedZoneTimeouts)
	ResetCloudLoggingConfig()
	ResetDescription()
	ResetDnssecConfig()
	ResetForceDestroy()
	ResetForwardingConfig()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeeringConfig()
	ResetPrivateVisibilityConfig()
	ResetProject()
	ResetReverseLookup()
	ResetServiceDirectoryConfig()
	ResetTimeouts()
	ResetVisibility()
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

// The jsii proxy struct for GoogleDnsManagedZone
type jsiiProxy_GoogleDnsManagedZone struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDnsManagedZone) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) CloudLoggingConfig() GoogleDnsManagedZoneCloudLoggingConfigOutputReference {
	var returns GoogleDnsManagedZoneCloudLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"cloudLoggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) CloudLoggingConfigInput() *GoogleDnsManagedZoneCloudLoggingConfig {
	var returns *GoogleDnsManagedZoneCloudLoggingConfig
	_jsii_.Get(
		j,
		"cloudLoggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DnssecConfig() GoogleDnsManagedZoneDnssecConfigOutputReference {
	var returns GoogleDnsManagedZoneDnssecConfigOutputReference
	_jsii_.Get(
		j,
		"dnssecConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) DnssecConfigInput() *GoogleDnsManagedZoneDnssecConfig {
	var returns *GoogleDnsManagedZoneDnssecConfig
	_jsii_.Get(
		j,
		"dnssecConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ForwardingConfig() GoogleDnsManagedZoneForwardingConfigOutputReference {
	var returns GoogleDnsManagedZoneForwardingConfigOutputReference
	_jsii_.Get(
		j,
		"forwardingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ForwardingConfigInput() *GoogleDnsManagedZoneForwardingConfig {
	var returns *GoogleDnsManagedZoneForwardingConfig
	_jsii_.Get(
		j,
		"forwardingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ManagedZoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) NameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) PeeringConfig() GoogleDnsManagedZonePeeringConfigOutputReference {
	var returns GoogleDnsManagedZonePeeringConfigOutputReference
	_jsii_.Get(
		j,
		"peeringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) PeeringConfigInput() *GoogleDnsManagedZonePeeringConfig {
	var returns *GoogleDnsManagedZonePeeringConfig
	_jsii_.Get(
		j,
		"peeringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) PrivateVisibilityConfig() GoogleDnsManagedZonePrivateVisibilityConfigOutputReference {
	var returns GoogleDnsManagedZonePrivateVisibilityConfigOutputReference
	_jsii_.Get(
		j,
		"privateVisibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) PrivateVisibilityConfigInput() *GoogleDnsManagedZonePrivateVisibilityConfig {
	var returns *GoogleDnsManagedZonePrivateVisibilityConfig
	_jsii_.Get(
		j,
		"privateVisibilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ReverseLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ReverseLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ServiceDirectoryConfig() GoogleDnsManagedZoneServiceDirectoryConfigOutputReference {
	var returns GoogleDnsManagedZoneServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) ServiceDirectoryConfigInput() *GoogleDnsManagedZoneServiceDirectoryConfig {
	var returns *GoogleDnsManagedZoneServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Timeouts() GoogleDnsManagedZoneTimeoutsOutputReference {
	var returns GoogleDnsManagedZoneTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZone) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dns_managed_zone google_dns_managed_zone} Resource.
func NewGoogleDnsManagedZone(scope constructs.Construct, id *string, config *GoogleDnsManagedZoneConfig) GoogleDnsManagedZone {
	_init_.Initialize()

	if err := validateNewGoogleDnsManagedZoneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsManagedZone{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dns_managed_zone google_dns_managed_zone} Resource.
func NewGoogleDnsManagedZone_Override(g GoogleDnsManagedZone, scope constructs.Construct, id *string, config *GoogleDnsManagedZoneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetReverseLookup(val interface{}) {
	if err := j.validateSetReverseLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reverseLookup",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZone)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDnsManagedZone resource upon running "cdktf plan <stack-name>".
func GoogleDnsManagedZone_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDnsManagedZone_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
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
func GoogleDnsManagedZone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDnsManagedZone_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDnsManagedZone_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDnsManagedZone_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDnsManagedZone_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDnsManagedZone_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDnsManagedZone_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDnsManagedZone.GoogleDnsManagedZone",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsManagedZone) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsManagedZone) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsManagedZone) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutCloudLoggingConfig(value *GoogleDnsManagedZoneCloudLoggingConfig) {
	if err := g.validatePutCloudLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutDnssecConfig(value *GoogleDnsManagedZoneDnssecConfig) {
	if err := g.validatePutDnssecConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnssecConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutForwardingConfig(value *GoogleDnsManagedZoneForwardingConfig) {
	if err := g.validatePutForwardingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putForwardingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutPeeringConfig(value *GoogleDnsManagedZonePeeringConfig) {
	if err := g.validatePutPeeringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPeeringConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutPrivateVisibilityConfig(value *GoogleDnsManagedZonePrivateVisibilityConfig) {
	if err := g.validatePutPrivateVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateVisibilityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutServiceDirectoryConfig(value *GoogleDnsManagedZoneServiceDirectoryConfig) {
	if err := g.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) PutTimeouts(value *GoogleDnsManagedZoneTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetCloudLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetDnssecConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDnssecConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetForwardingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetForwardingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetPeeringConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPeeringConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetPrivateVisibilityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateVisibilityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetReverseLookup() {
	_jsii_.InvokeVoid(
		g,
		"resetReverseLookup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) ResetVisibility() {
	_jsii_.InvokeVoid(
		g,
		"resetVisibility",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZone) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZone) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

