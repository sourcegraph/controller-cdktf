package computeregionnetworkfirewallpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeregionnetworkfirewallpolicyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_network_firewall_policy_rule google_compute_region_network_firewall_policy_rule}.
type ComputeRegionNetworkFirewallPolicyRule interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Disabled() any
	SetDisabled(val any)
	DisabledInput() any
	EnableLogging() any
	SetEnableLogging(val any)
	EnableLoggingInput() any
	FirewallPolicy() *string
	SetFirewallPolicy(val *string)
	FirewallPolicyInput() *string
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
	Kind() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() ComputeRegionNetworkFirewallPolicyRuleMatchOutputReference
	MatchInput() *ComputeRegionNetworkFirewallPolicyRuleMatch
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	RuleTupleCount() *float64
	SecurityProfileGroup() *string
	SetSecurityProfileGroup(val *string)
	SecurityProfileGroupInput() *string
	TargetForwardingRules() *[]*string
	SetTargetForwardingRules(val *[]*string)
	TargetForwardingRulesInput() *[]*string
	TargetSecureTags() ComputeRegionNetworkFirewallPolicyRuleTargetSecureTagsList
	TargetSecureTagsInput() any
	TargetServiceAccounts() *[]*string
	SetTargetServiceAccounts(val *[]*string)
	TargetServiceAccountsInput() *[]*string
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionNetworkFirewallPolicyRuleTimeoutsOutputReference
	TimeoutsInput() any
	TlsInspect() any
	SetTlsInspect(val any)
	TlsInspectInput() any
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
	PutMatch(value *ComputeRegionNetworkFirewallPolicyRuleMatch)
	PutTargetSecureTags(value any)
	PutTimeouts(value *ComputeRegionNetworkFirewallPolicyRuleTimeouts)
	ResetDescription()
	ResetDisabled()
	ResetEnableLogging()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetRuleName()
	ResetSecurityProfileGroup()
	ResetTargetForwardingRules()
	ResetTargetSecureTags()
	ResetTargetServiceAccounts()
	ResetTargetType()
	ResetTimeouts()
	ResetTlsInspect()
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

// The jsii proxy struct for ComputeRegionNetworkFirewallPolicyRule
type jsiiProxy_ComputeRegionNetworkFirewallPolicyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Disabled() any {
	var returns any
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) DisabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) EnableLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) EnableLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) FirewallPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) FirewallPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Match() ComputeRegionNetworkFirewallPolicyRuleMatchOutputReference {
	var returns ComputeRegionNetworkFirewallPolicyRuleMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) MatchInput() *ComputeRegionNetworkFirewallPolicyRuleMatch {
	var returns *ComputeRegionNetworkFirewallPolicyRuleMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) RuleTupleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleTupleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SecurityProfileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SecurityProfileGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetForwardingRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetForwardingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetForwardingRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetForwardingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetSecureTags() ComputeRegionNetworkFirewallPolicyRuleTargetSecureTagsList {
	var returns ComputeRegionNetworkFirewallPolicyRuleTargetSecureTagsList
	_jsii_.Get(
		j,
		"targetSecureTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetSecureTagsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"targetSecureTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) Timeouts() ComputeRegionNetworkFirewallPolicyRuleTimeoutsOutputReference {
	var returns ComputeRegionNetworkFirewallPolicyRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TlsInspect() any {
	var returns any
	_jsii_.Get(
		j,
		"tlsInspect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) TlsInspectInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tlsInspectInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_network_firewall_policy_rule google_compute_region_network_firewall_policy_rule} Resource.
func NewComputeRegionNetworkFirewallPolicyRule(scope constructs.Construct, id *string, config *ComputeRegionNetworkFirewallPolicyRuleConfig) ComputeRegionNetworkFirewallPolicyRule {
	_init_.Initialize()

	if err := validateNewComputeRegionNetworkFirewallPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionNetworkFirewallPolicyRule{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_network_firewall_policy_rule google_compute_region_network_firewall_policy_rule} Resource.
func NewComputeRegionNetworkFirewallPolicyRule_Override(c ComputeRegionNetworkFirewallPolicyRule, scope constructs.Construct, id *string, config *ComputeRegionNetworkFirewallPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetDisabled(val any) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetEnableLogging(val any) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetFirewallPolicy(val *string) {
	if err := j.validateSetFirewallPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetSecurityProfileGroup(val *string) {
	if err := j.validateSetSecurityProfileGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileGroup",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetTargetForwardingRules(val *[]*string) {
	if err := j.validateSetTargetForwardingRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetForwardingRules",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SetTlsInspect(val any) {
	if err := j.validateSetTlsInspectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspect",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionNetworkFirewallPolicyRule resource upon running "cdktf plan <stack-name>".
func ComputeRegionNetworkFirewallPolicyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionNetworkFirewallPolicyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
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
func ComputeRegionNetworkFirewallPolicyRule_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionNetworkFirewallPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionNetworkFirewallPolicyRule_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionNetworkFirewallPolicyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionNetworkFirewallPolicyRule_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionNetworkFirewallPolicyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ComputeRegionNetworkFirewallPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyRule.ComputeRegionNetworkFirewallPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) PutMatch(value *ComputeRegionNetworkFirewallPolicyRuleMatch) {
	if err := c.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMatch",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) PutTargetSecureTags(value any) {
	if err := c.validatePutTargetSecureTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetSecureTags",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) PutTimeouts(value *ComputeRegionNetworkFirewallPolicyRuleTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetRuleName() {
	_jsii_.InvokeVoid(
		c,
		"resetRuleName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetSecurityProfileGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityProfileGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTargetForwardingRules() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetForwardingRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTargetSecureTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetSecureTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTargetType() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ResetTlsInspect() {
	_jsii_.InvokeVoid(
		c,
		"resetTlsInspect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyRule) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
