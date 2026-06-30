package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer}.
type LoadBalancer interface {
	cdktf.TerraformResource
	AdaptiveRouting() LoadBalancerAdaptiveRoutingList
	AdaptiveRoutingInput() any
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
	CountryPools() LoadBalancerCountryPoolsList
	CountryPoolsInput() any
	CreatedOn() *string
	DefaultPoolIds() *[]*string
	SetDefaultPoolIds(val *[]*string)
	DefaultPoolIdsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	FallbackPoolId() *string
	SetFallbackPoolId(val *string)
	FallbackPoolIdInput() *string
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
	LocationStrategy() LoadBalancerLocationStrategyList
	LocationStrategyInput() any
	ModifiedOn() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PopPools() LoadBalancerPopPoolsList
	PopPoolsInput() any
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	Proxied() any
	SetProxied(val any)
	ProxiedInput() any
	RandomSteering() LoadBalancerRandomSteeringList
	RandomSteeringInput() any
	// Experimental.
	RawOverrides() any
	RegionPools() LoadBalancerRegionPoolsList
	RegionPoolsInput() any
	Rules() LoadBalancerRulesList
	RulesInput() any
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityAttributes() LoadBalancerSessionAffinityAttributesList
	SessionAffinityAttributesInput() any
	SessionAffinityInput() *string
	SessionAffinityTtl() *float64
	SetSessionAffinityTtl(val *float64)
	SessionAffinityTtlInput() *float64
	SteeringPolicy() *string
	SetSteeringPolicy(val *string)
	SteeringPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutAdaptiveRouting(value any)
	PutCountryPools(value any)
	PutLocationStrategy(value any)
	PutPopPools(value any)
	PutRandomSteering(value any)
	PutRegionPools(value any)
	PutRules(value any)
	PutSessionAffinityAttributes(value any)
	ResetAdaptiveRouting()
	ResetCountryPools()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetLocationStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPopPools()
	ResetProxied()
	ResetRandomSteering()
	ResetRegionPools()
	ResetRules()
	ResetSessionAffinity()
	ResetSessionAffinityAttributes()
	ResetSessionAffinityTtl()
	ResetSteeringPolicy()
	ResetTtl()
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

// The jsii proxy struct for LoadBalancer
type jsiiProxy_LoadBalancer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LoadBalancer) AdaptiveRouting() LoadBalancerAdaptiveRoutingList {
	var returns LoadBalancerAdaptiveRoutingList
	_jsii_.Get(
		j,
		"adaptiveRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) AdaptiveRoutingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"adaptiveRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CountryPools() LoadBalancerCountryPoolsList {
	var returns LoadBalancerCountryPoolsList
	_jsii_.Get(
		j,
		"countryPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CountryPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"countryPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DefaultPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DefaultPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FallbackPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FallbackPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LocationStrategy() LoadBalancerLocationStrategyList {
	var returns LoadBalancerLocationStrategyList
	_jsii_.Get(
		j,
		"locationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LocationStrategyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"locationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) PopPools() LoadBalancerPopPoolsList {
	var returns LoadBalancerPopPoolsList
	_jsii_.Get(
		j,
		"popPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) PopPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"popPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Proxied() any {
	var returns any
	_jsii_.Get(
		j,
		"proxied",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ProxiedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"proxiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RandomSteering() LoadBalancerRandomSteeringList {
	var returns LoadBalancerRandomSteeringList
	_jsii_.Get(
		j,
		"randomSteering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RandomSteeringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"randomSteeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RegionPools() LoadBalancerRegionPoolsList {
	var returns LoadBalancerRegionPoolsList
	_jsii_.Get(
		j,
		"regionPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RegionPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"regionPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Rules() LoadBalancerRulesList {
	var returns LoadBalancerRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) RulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityAttributes() LoadBalancerSessionAffinityAttributesList {
	var returns LoadBalancerSessionAffinityAttributesList
	_jsii_.Get(
		j,
		"sessionAffinityAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityAttributesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sessionAffinityAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SessionAffinityTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SteeringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) SteeringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer} Resource.
func NewLoadBalancer(scope constructs.Construct, id *string, config *LoadBalancerConfig) LoadBalancer {
	_init_.Initialize()

	if err := validateNewLoadBalancerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancer{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer cloudflare_load_balancer} Resource.
func NewLoadBalancer_Override(l LoadBalancer, scope constructs.Construct, id *string, config *LoadBalancerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		[]any{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LoadBalancer) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetDefaultPoolIds(val *[]*string) {
	if err := j.validateSetDefaultPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPoolIds",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetFallbackPoolId(val *string) {
	if err := j.validateSetFallbackPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackPoolId",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetProxied(val any) {
	if err := j.validateSetProxiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxied",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetSessionAffinityTtl(val *float64) {
	if err := j.validateSetSessionAffinityTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinityTtl",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetSteeringPolicy(val *string) {
	if err := j.validateSetSteeringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"steeringPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_LoadBalancer) SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a LoadBalancer resource upon running "cdktf plan <stack-name>".
func LoadBalancer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLoadBalancer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
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
func LoadBalancer_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancer_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func LoadBalancer_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateLoadBalancer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func LoadBalancer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LoadBalancer) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (l *jsiiProxy_LoadBalancer) AddOverride(path *string, value any) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]any{path, value},
	)
}

func (l *jsiiProxy_LoadBalancer) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]any{id, provider},
	)
}

func (l *jsiiProxy_LoadBalancer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]any{id},
	)
}

func (l *jsiiProxy_LoadBalancer) MoveTo(moveTarget *string, index any) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (l *jsiiProxy_LoadBalancer) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]any{id},
	)
}

func (l *jsiiProxy_LoadBalancer) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (l *jsiiProxy_LoadBalancer) PutAdaptiveRouting(value any) {
	if err := l.validatePutAdaptiveRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdaptiveRouting",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutCountryPools(value any) {
	if err := l.validatePutCountryPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCountryPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutLocationStrategy(value any) {
	if err := l.validatePutLocationStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLocationStrategy",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutPopPools(value any) {
	if err := l.validatePutPopPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPopPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRandomSteering(value any) {
	if err := l.validatePutRandomSteeringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRandomSteering",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRegionPools(value any) {
	if err := l.validatePutRegionPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRegionPools",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutRules(value any) {
	if err := l.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRules",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) PutSessionAffinityAttributes(value any) {
	if err := l.validatePutSessionAffinityAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSessionAffinityAttributes",
		[]any{value},
	)
}

func (l *jsiiProxy_LoadBalancer) ResetAdaptiveRouting() {
	_jsii_.InvokeVoid(
		l,
		"resetAdaptiveRouting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetCountryPools() {
	_jsii_.InvokeVoid(
		l,
		"resetCountryPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetLocationStrategy() {
	_jsii_.InvokeVoid(
		l,
		"resetLocationStrategy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetPopPools() {
	_jsii_.InvokeVoid(
		l,
		"resetPopPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetProxied() {
	_jsii_.InvokeVoid(
		l,
		"resetProxied",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRandomSteering() {
	_jsii_.InvokeVoid(
		l,
		"resetRandomSteering",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRegionPools() {
	_jsii_.InvokeVoid(
		l,
		"resetRegionPools",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetRules() {
	_jsii_.InvokeVoid(
		l,
		"resetRules",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinityAttributes() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityAttributes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSessionAffinityTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionAffinityTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetSteeringPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSteeringPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) ResetTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancer) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
