package googleloggingorganizationsink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleloggingorganizationsink/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_organization_sink google_logging_organization_sink}.
type GoogleLoggingOrganizationSink interface {
	cdktf.TerraformResource
	BigqueryOptions() GoogleLoggingOrganizationSinkBigqueryOptionsOutputReference
	BigqueryOptionsInput() *GoogleLoggingOrganizationSinkBigqueryOptions
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	Exclusions() GoogleLoggingOrganizationSinkExclusionsList
	ExclusionsInput() interface{}
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
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
	IncludeChildren() interface{}
	SetIncludeChildren(val interface{})
	IncludeChildrenInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WriterIdentity() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBigqueryOptions(value *GoogleLoggingOrganizationSinkBigqueryOptions)
	PutExclusions(value interface{})
	ResetBigqueryOptions()
	ResetDescription()
	ResetDisabled()
	ResetExclusions()
	ResetFilter()
	ResetId()
	ResetIncludeChildren()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleLoggingOrganizationSink
type jsiiProxy_GoogleLoggingOrganizationSink struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) BigqueryOptions() GoogleLoggingOrganizationSinkBigqueryOptionsOutputReference {
	var returns GoogleLoggingOrganizationSinkBigqueryOptionsOutputReference
	_jsii_.Get(
		j,
		"bigqueryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) BigqueryOptionsInput() *GoogleLoggingOrganizationSinkBigqueryOptions {
	var returns *GoogleLoggingOrganizationSinkBigqueryOptions
	_jsii_.Get(
		j,
		"bigqueryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Exclusions() GoogleLoggingOrganizationSinkExclusionsList {
	var returns GoogleLoggingOrganizationSinkExclusionsList
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) ExclusionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) IncludeChildren() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) IncludeChildrenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeChildrenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink) WriterIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writerIdentity",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_organization_sink google_logging_organization_sink} Resource.
func NewGoogleLoggingOrganizationSink(scope constructs.Construct, id *string, config *GoogleLoggingOrganizationSinkConfig) GoogleLoggingOrganizationSink {
	_init_.Initialize()

	if err := validateNewGoogleLoggingOrganizationSinkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLoggingOrganizationSink{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_organization_sink google_logging_organization_sink} Resource.
func NewGoogleLoggingOrganizationSink_Override(g GoogleLoggingOrganizationSink, scope constructs.Construct, id *string, config *GoogleLoggingOrganizationSinkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetIncludeChildren(val interface{}) {
	if err := j.validateSetIncludeChildrenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeChildren",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingOrganizationSink)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
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
func GoogleLoggingOrganizationSink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLoggingOrganizationSink_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleLoggingOrganizationSink_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLoggingOrganizationSink_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleLoggingOrganizationSink_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLoggingOrganizationSink_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleLoggingOrganizationSink_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googlebeta.googleLoggingOrganizationSink.GoogleLoggingOrganizationSink",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLoggingOrganizationSink) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) PutBigqueryOptions(value *GoogleLoggingOrganizationSinkBigqueryOptions) {
	if err := g.validatePutBigqueryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) PutExclusions(value interface{}) {
	if err := g.validatePutExclusionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExclusions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetBigqueryOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetExclusions() {
	_jsii_.InvokeVoid(
		g,
		"resetExclusions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetIncludeChildren() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeChildren",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingOrganizationSink) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

