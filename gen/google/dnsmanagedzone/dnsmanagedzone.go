package dnsmanagedzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dnsmanagedzone/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_managed_zone google_dns_managed_zone}.
type DnsManagedZone interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudLoggingConfig() DnsManagedZoneCloudLoggingConfigOutputReference
	CloudLoggingConfigInput() *DnsManagedZoneCloudLoggingConfig
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
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
	DnssecConfig() DnsManagedZoneDnssecConfigOutputReference
	DnssecConfigInput() *DnsManagedZoneDnssecConfig
	EffectiveLabels() cdktf.StringMap
	ForceDestroy() any
	SetForceDestroy(val any)
	ForceDestroyInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardingConfig() DnsManagedZoneForwardingConfigOutputReference
	ForwardingConfigInput() *DnsManagedZoneForwardingConfig
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
	PeeringConfig() DnsManagedZonePeeringConfigOutputReference
	PeeringConfigInput() *DnsManagedZonePeeringConfig
	PrivateVisibilityConfig() DnsManagedZonePrivateVisibilityConfigOutputReference
	PrivateVisibilityConfigInput() *DnsManagedZonePrivateVisibilityConfig
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DnsManagedZoneTimeoutsOutputReference
	TimeoutsInput() any
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCloudLoggingConfig(value *DnsManagedZoneCloudLoggingConfig)
	PutDnssecConfig(value *DnsManagedZoneDnssecConfig)
	PutForwardingConfig(value *DnsManagedZoneForwardingConfig)
	PutPeeringConfig(value *DnsManagedZonePeeringConfig)
	PutPrivateVisibilityConfig(value *DnsManagedZonePrivateVisibilityConfig)
	PutTimeouts(value *DnsManagedZoneTimeouts)
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
	ResetTimeouts()
	ResetVisibility()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for DnsManagedZone
type jsiiProxy_DnsManagedZone struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DnsManagedZone) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) CloudLoggingConfig() DnsManagedZoneCloudLoggingConfigOutputReference {
	var returns DnsManagedZoneCloudLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"cloudLoggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) CloudLoggingConfigInput() *DnsManagedZoneCloudLoggingConfig {
	var returns *DnsManagedZoneCloudLoggingConfig
	_jsii_.Get(
		j,
		"cloudLoggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DnssecConfig() DnsManagedZoneDnssecConfigOutputReference {
	var returns DnsManagedZoneDnssecConfigOutputReference
	_jsii_.Get(
		j,
		"dnssecConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) DnssecConfigInput() *DnsManagedZoneDnssecConfig {
	var returns *DnsManagedZoneDnssecConfig
	_jsii_.Get(
		j,
		"dnssecConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ForceDestroy() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ForceDestroyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ForwardingConfig() DnsManagedZoneForwardingConfigOutputReference {
	var returns DnsManagedZoneForwardingConfigOutputReference
	_jsii_.Get(
		j,
		"forwardingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ForwardingConfigInput() *DnsManagedZoneForwardingConfig {
	var returns *DnsManagedZoneForwardingConfig
	_jsii_.Get(
		j,
		"forwardingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ManagedZoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) NameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) PeeringConfig() DnsManagedZonePeeringConfigOutputReference {
	var returns DnsManagedZonePeeringConfigOutputReference
	_jsii_.Get(
		j,
		"peeringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) PeeringConfigInput() *DnsManagedZonePeeringConfig {
	var returns *DnsManagedZonePeeringConfig
	_jsii_.Get(
		j,
		"peeringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) PrivateVisibilityConfig() DnsManagedZonePrivateVisibilityConfigOutputReference {
	var returns DnsManagedZonePrivateVisibilityConfigOutputReference
	_jsii_.Get(
		j,
		"privateVisibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) PrivateVisibilityConfigInput() *DnsManagedZonePrivateVisibilityConfig {
	var returns *DnsManagedZonePrivateVisibilityConfig
	_jsii_.Get(
		j,
		"privateVisibilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Timeouts() DnsManagedZoneTimeoutsOutputReference {
	var returns DnsManagedZoneTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsManagedZone) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_managed_zone google_dns_managed_zone} Resource.
func NewDnsManagedZone(scope constructs.Construct, id *string, config *DnsManagedZoneConfig) DnsManagedZone {
	_init_.Initialize()

	if err := validateNewDnsManagedZoneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsManagedZone{}

	_jsii_.Create(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dns_managed_zone google_dns_managed_zone} Resource.
func NewDnsManagedZone_Override(d DnsManagedZone, scope constructs.Construct, id *string, config *DnsManagedZoneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetDnsName(val *string) {
	if err := j.validateSetDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetForceDestroy(val any) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DnsManagedZone) SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

// Generates CDKTF code for importing a DnsManagedZone resource upon running "cdktf plan <stack-name>".
func DnsManagedZone_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDnsManagedZone_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func DnsManagedZone_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDnsManagedZone_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsManagedZone_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDnsManagedZone_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DnsManagedZone_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDnsManagedZone_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DnsManagedZone_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dnsManagedZone.DnsManagedZone",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DnsManagedZone) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DnsManagedZone) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DnsManagedZone) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DnsManagedZone) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DnsManagedZone) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DnsManagedZone) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DnsManagedZone) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutCloudLoggingConfig(value *DnsManagedZoneCloudLoggingConfig) {
	if err := d.validatePutCloudLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudLoggingConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutDnssecConfig(value *DnsManagedZoneDnssecConfig) {
	if err := d.validatePutDnssecConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDnssecConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutForwardingConfig(value *DnsManagedZoneForwardingConfig) {
	if err := d.validatePutForwardingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForwardingConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutPeeringConfig(value *DnsManagedZonePeeringConfig) {
	if err := d.validatePutPeeringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPeeringConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutPrivateVisibilityConfig(value *DnsManagedZonePrivateVisibilityConfig) {
	if err := d.validatePutPrivateVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrivateVisibilityConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) PutTimeouts(value *DnsManagedZoneTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetCloudLoggingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudLoggingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetDnssecConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDnssecConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetForwardingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetForwardingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetPeeringConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPeeringConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetPrivateVisibilityConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateVisibilityConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) ResetVisibility() {
	_jsii_.InvokeVoid(
		d,
		"resetVisibility",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsManagedZone) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsManagedZone) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
