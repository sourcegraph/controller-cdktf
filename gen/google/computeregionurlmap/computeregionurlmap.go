package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeregionurlmap/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_url_map google_compute_region_url_map}.
type ComputeRegionUrlMap interface {
	cdktf.TerraformResource
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
	DefaultRouteAction() ComputeRegionUrlMapDefaultRouteActionOutputReference
	DefaultRouteActionInput() *ComputeRegionUrlMapDefaultRouteAction
	DefaultService() *string
	SetDefaultService(val *string)
	DefaultServiceInput() *string
	DefaultUrlRedirect() ComputeRegionUrlMapDefaultUrlRedirectOutputReference
	DefaultUrlRedirectInput() *ComputeRegionUrlMapDefaultUrlRedirect
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeaderAction() ComputeRegionUrlMapHeaderActionOutputReference
	HeaderActionInput() *ComputeRegionUrlMapHeaderAction
	HostRule() ComputeRegionUrlMapHostRuleList
	HostRuleInput() any
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MapId() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PathMatcher() ComputeRegionUrlMapPathMatcherList
	PathMatcherInput() any
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
	SelfLink() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Test() ComputeRegionUrlMapTestList
	TestInput() any
	Timeouts() ComputeRegionUrlMapTimeoutsOutputReference
	TimeoutsInput() any
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
	PutDefaultRouteAction(value *ComputeRegionUrlMapDefaultRouteAction)
	PutDefaultUrlRedirect(value *ComputeRegionUrlMapDefaultUrlRedirect)
	PutHeaderAction(value *ComputeRegionUrlMapHeaderAction)
	PutHostRule(value any)
	PutPathMatcher(value any)
	PutTest(value any)
	PutTimeouts(value *ComputeRegionUrlMapTimeouts)
	ResetDefaultRouteAction()
	ResetDefaultService()
	ResetDefaultUrlRedirect()
	ResetDescription()
	ResetHeaderAction()
	ResetHostRule()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathMatcher()
	ResetProject()
	ResetRegion()
	ResetTest()
	ResetTimeouts()
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

// The jsii proxy struct for ComputeRegionUrlMap
type jsiiProxy_ComputeRegionUrlMap struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionUrlMap) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultRouteAction() ComputeRegionUrlMapDefaultRouteActionOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionOutputReference
	_jsii_.Get(
		j,
		"defaultRouteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultRouteActionInput() *ComputeRegionUrlMapDefaultRouteAction {
	var returns *ComputeRegionUrlMapDefaultRouteAction
	_jsii_.Get(
		j,
		"defaultRouteActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultUrlRedirect() ComputeRegionUrlMapDefaultUrlRedirectOutputReference {
	var returns ComputeRegionUrlMapDefaultUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"defaultUrlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DefaultUrlRedirectInput() *ComputeRegionUrlMapDefaultUrlRedirect {
	var returns *ComputeRegionUrlMapDefaultUrlRedirect
	_jsii_.Get(
		j,
		"defaultUrlRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) HeaderAction() ComputeRegionUrlMapHeaderActionOutputReference {
	var returns ComputeRegionUrlMapHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) HeaderActionInput() *ComputeRegionUrlMapHeaderAction {
	var returns *ComputeRegionUrlMapHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) HostRule() ComputeRegionUrlMapHostRuleList {
	var returns ComputeRegionUrlMapHostRuleList
	_jsii_.Get(
		j,
		"hostRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) HostRuleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) MapId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) PathMatcher() ComputeRegionUrlMapPathMatcherList {
	var returns ComputeRegionUrlMapPathMatcherList
	_jsii_.Get(
		j,
		"pathMatcher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) PathMatcherInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pathMatcherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Test() ComputeRegionUrlMapTestList {
	var returns ComputeRegionUrlMapTestList
	_jsii_.Get(
		j,
		"test",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) TestInput() any {
	var returns any
	_jsii_.Get(
		j,
		"testInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) Timeouts() ComputeRegionUrlMapTimeoutsOutputReference {
	var returns ComputeRegionUrlMapTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMap) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_url_map google_compute_region_url_map} Resource.
func NewComputeRegionUrlMap(scope constructs.Construct, id *string, config *ComputeRegionUrlMapConfig) ComputeRegionUrlMap {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMap{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_region_url_map google_compute_region_url_map} Resource.
func NewComputeRegionUrlMap_Override(c ComputeRegionUrlMap, scope constructs.Construct, id *string, config *ComputeRegionUrlMapConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetDefaultService(val *string) {
	if err := j.validateSetDefaultServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultService",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMap) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionUrlMap resource upon running "cdktf plan <stack-name>".
func ComputeRegionUrlMap_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionUrlMap_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
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
func ComputeRegionUrlMap_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionUrlMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionUrlMap_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionUrlMap_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionUrlMap_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRegionUrlMap_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ComputeRegionUrlMap_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMap",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMap) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMap) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMap) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutDefaultRouteAction(value *ComputeRegionUrlMapDefaultRouteAction) {
	if err := c.validatePutDefaultRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultRouteAction",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutDefaultUrlRedirect(value *ComputeRegionUrlMapDefaultUrlRedirect) {
	if err := c.validatePutDefaultUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultUrlRedirect",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutHeaderAction(value *ComputeRegionUrlMapHeaderAction) {
	if err := c.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeaderAction",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutHostRule(value any) {
	if err := c.validatePutHostRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostRule",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutPathMatcher(value any) {
	if err := c.validatePutPathMatcherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPathMatcher",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutTest(value any) {
	if err := c.validatePutTestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTest",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) PutTimeouts(value *ComputeRegionUrlMapTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetDefaultRouteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRouteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetDefaultService() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetDefaultUrlRedirect() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultUrlRedirect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetHostRule() {
	_jsii_.InvokeVoid(
		c,
		"resetHostRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetPathMatcher() {
	_jsii_.InvokeVoid(
		c,
		"resetPathMatcher",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetTest() {
	_jsii_.InvokeVoid(
		c,
		"resetTest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMap) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMap) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
