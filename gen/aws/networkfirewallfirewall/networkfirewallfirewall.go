package networkfirewallfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/networkfirewallfirewall/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall}.
type NetworkfirewallFirewall interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DeleteProtection() any
	SetDeleteProtection(val any)
	DeleteProtectionInput() any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EncryptionConfiguration() NetworkfirewallFirewallEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() *NetworkfirewallFirewallEncryptionConfiguration
	FirewallPolicyArn() *string
	SetFirewallPolicyArn(val *string)
	FirewallPolicyArnInput() *string
	FirewallPolicyChangeProtection() any
	SetFirewallPolicyChangeProtection(val any)
	FirewallPolicyChangeProtectionInput() any
	FirewallStatus() NetworkfirewallFirewallFirewallStatusList
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
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	SubnetChangeProtection() any
	SetSubnetChangeProtection(val any)
	SubnetChangeProtectionInput() any
	SubnetMapping() NetworkfirewallFirewallSubnetMappingList
	SubnetMappingInput() any
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	UpdateToken() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutEncryptionConfiguration(value *NetworkfirewallFirewallEncryptionConfiguration)
	PutSubnetMapping(value any)
	ResetDeleteProtection()
	ResetDescription()
	ResetEncryptionConfiguration()
	ResetFirewallPolicyChangeProtection()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubnetChangeProtection()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for NetworkfirewallFirewall
type jsiiProxy_NetworkfirewallFirewall struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkfirewallFirewall) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) DeleteProtection() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) DeleteProtectionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deleteProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) EncryptionConfiguration() NetworkfirewallFirewallEncryptionConfigurationOutputReference {
	var returns NetworkfirewallFirewallEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) EncryptionConfigurationInput() *NetworkfirewallFirewallEncryptionConfiguration {
	var returns *NetworkfirewallFirewallEncryptionConfiguration
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FirewallPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FirewallPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FirewallPolicyChangeProtection() any {
	var returns any
	_jsii_.Get(
		j,
		"firewallPolicyChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FirewallPolicyChangeProtectionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"firewallPolicyChangeProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FirewallStatus() NetworkfirewallFirewallFirewallStatusList {
	var returns NetworkfirewallFirewallFirewallStatusList
	_jsii_.Get(
		j,
		"firewallStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) SubnetChangeProtection() any {
	var returns any
	_jsii_.Get(
		j,
		"subnetChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) SubnetChangeProtectionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"subnetChangeProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) SubnetMapping() NetworkfirewallFirewallSubnetMappingList {
	var returns NetworkfirewallFirewallSubnetMappingList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) SubnetMappingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) UpdateToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkfirewallFirewall) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall} Resource.
func NewNetworkfirewallFirewall(scope constructs.Construct, id *string, config *NetworkfirewallFirewallConfig) NetworkfirewallFirewall {
	_init_.Initialize()

	if err := validateNewNetworkfirewallFirewallParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkfirewallFirewall{}

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall} Resource.
func NewNetworkfirewallFirewall_Override(n NetworkfirewallFirewall, scope constructs.Construct, id *string, config *NetworkfirewallFirewallConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		[]any{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetDeleteProtection(val any) {
	if err := j.validateSetDeleteProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteProtection",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetFirewallPolicyArn(val *string) {
	if err := j.validateSetFirewallPolicyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyArn",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetFirewallPolicyChangeProtection(val any) {
	if err := j.validateSetFirewallPolicyChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyChangeProtection",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetSubnetChangeProtection(val any) {
	if err := j.validateSetSubnetChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetChangeProtection",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_NetworkfirewallFirewall) SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a NetworkfirewallFirewall resource upon running "cdktf plan <stack-name>".
func NetworkfirewallFirewall_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkfirewallFirewall_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
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
func NetworkfirewallFirewall_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkfirewallFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkfirewallFirewall_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkfirewallFirewall_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkfirewallFirewall_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkfirewallFirewall_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func NetworkfirewallFirewall_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.networkfirewallFirewall.NetworkfirewallFirewall",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) AddOverride(path *string, value any) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]any{path, value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]any{id, provider},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) MoveTo(moveTarget *string, index any) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) PutEncryptionConfiguration(value *NetworkfirewallFirewallEncryptionConfiguration) {
	if err := n.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putEncryptionConfiguration",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) PutSubnetMapping(value any) {
	if err := n.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSubnetMapping",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetDeleteProtection() {
	_jsii_.InvokeVoid(
		n,
		"resetDeleteProtection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetFirewallPolicyChangeProtection() {
	_jsii_.InvokeVoid(
		n,
		"resetFirewallPolicyChangeProtection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetSubnetChangeProtection() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnetChangeProtection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) ResetTagsAll() {
	_jsii_.InvokeVoid(
		n,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkfirewallFirewall) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkfirewallFirewall) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
