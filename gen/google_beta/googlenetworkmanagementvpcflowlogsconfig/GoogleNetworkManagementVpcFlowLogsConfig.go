package googlenetworkmanagementvpcflowlogsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkmanagementvpcflowlogsconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config}.
type GoogleNetworkManagementVpcFlowLogsConfig interface {
	cdktf.TerraformResource
	AggregationInterval() *string
	SetAggregationInterval(val *string)
	AggregationIntervalInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	FilterExpr() *string
	SetFilterExpr(val *string)
	FilterExprInput() *string
	FlowSampling() *float64
	SetFlowSampling(val *float64)
	FlowSamplingInput() *float64
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
	InterconnectAttachment() *string
	SetInterconnectAttachment(val *string)
	InterconnectAttachmentInput() *string
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
	Metadata() *string
	SetMetadata(val *string)
	MetadataFields() *[]*string
	SetMetadataFields(val *[]*string)
	MetadataFieldsInput() *[]*string
	MetadataInput() *string
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
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNetworkManagementVpcFlowLogsConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VpcFlowLogsConfigId() *string
	SetVpcFlowLogsConfigId(val *string)
	VpcFlowLogsConfigIdInput() *string
	VpnTunnel() *string
	SetVpnTunnel(val *string)
	VpnTunnelInput() *string
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
	PutTimeouts(value *GoogleNetworkManagementVpcFlowLogsConfigTimeouts)
	ResetAggregationInterval()
	ResetDescription()
	ResetFilterExpr()
	ResetFlowSampling()
	ResetId()
	ResetInterconnectAttachment()
	ResetLabels()
	ResetMetadata()
	ResetMetadataFields()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetState()
	ResetTimeouts()
	ResetVpnTunnel()
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

// The jsii proxy struct for GoogleNetworkManagementVpcFlowLogsConfig
type jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) AggregationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) AggregationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) FilterExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) FilterExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) FlowSampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) FlowSamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) InterconnectAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) InterconnectAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) Timeouts() GoogleNetworkManagementVpcFlowLogsConfigTimeoutsOutputReference {
	var returns GoogleNetworkManagementVpcFlowLogsConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) VpcFlowLogsConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) VpcFlowLogsConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) VpnTunnel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) VpnTunnelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnTunnelInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config} Resource.
func NewGoogleNetworkManagementVpcFlowLogsConfig(scope constructs.Construct, id *string, config *GoogleNetworkManagementVpcFlowLogsConfigConfig) GoogleNetworkManagementVpcFlowLogsConfig {
	_init_.Initialize()

	if err := validateNewGoogleNetworkManagementVpcFlowLogsConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_management_vpc_flow_logs_config google_network_management_vpc_flow_logs_config} Resource.
func NewGoogleNetworkManagementVpcFlowLogsConfig_Override(g GoogleNetworkManagementVpcFlowLogsConfig, scope constructs.Construct, id *string, config *GoogleNetworkManagementVpcFlowLogsConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetAggregationInterval(val *string) {
	if err := j.validateSetAggregationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetFilterExpr(val *string) {
	if err := j.validateSetFilterExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterExpr",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetFlowSampling(val *float64) {
	if err := j.validateSetFlowSamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowSampling",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetInterconnectAttachment(val *string) {
	if err := j.validateSetInterconnectAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnectAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetVpcFlowLogsConfigId(val *string) {
	if err := j.validateSetVpcFlowLogsConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcFlowLogsConfigId",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig)SetVpnTunnel(val *string) {
	if err := j.validateSetVpnTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnTunnel",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNetworkManagementVpcFlowLogsConfig resource upon running "cdktf plan <stack-name>".
func GoogleNetworkManagementVpcFlowLogsConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementVpcFlowLogsConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
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
func GoogleNetworkManagementVpcFlowLogsConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementVpcFlowLogsConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkManagementVpcFlowLogsConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementVpcFlowLogsConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkManagementVpcFlowLogsConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementVpcFlowLogsConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetworkManagementVpcFlowLogsConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNetworkManagementVpcFlowLogsConfig.GoogleNetworkManagementVpcFlowLogsConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) PutTimeouts(value *GoogleNetworkManagementVpcFlowLogsConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetAggregationInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregationInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetFilterExpr() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterExpr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetFlowSampling() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowSampling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetInterconnectAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetInterconnectAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetMetadataFields() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetState() {
	_jsii_.InvokeVoid(
		g,
		"resetState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ResetVpnTunnel() {
	_jsii_.InvokeVoid(
		g,
		"resetVpnTunnel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementVpcFlowLogsConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

