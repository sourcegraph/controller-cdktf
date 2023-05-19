package origincacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/origincacertificate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate cloudflare_origin_ca_certificate}.
type OriginCaCertificate interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
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
	Csr() *string
	SetCsr(val *string)
	CsrInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExpiresOn() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostnames() *[]*string
	SetHostnames(val *[]*string)
	HostnamesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinDaysForRenewal() *float64
	SetMinDaysForRenewal(val *float64)
	MinDaysForRenewalInput() *float64
	// The tree node.
	Node() constructs.Node
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
	RequestedValidity() *float64
	SetRequestedValidity(val *float64)
	RequestedValidityInput() *float64
	RequestType() *string
	SetRequestType(val *string)
	RequestTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCsr()
	ResetId()
	ResetMinDaysForRenewal()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestedValidity()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OriginCaCertificate
type jsiiProxy_OriginCaCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OriginCaCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Csr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) CsrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) ExpiresOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Hostnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) HostnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) MinDaysForRenewal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDaysForRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) MinDaysForRenewalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDaysForRenewalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) RequestedValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) RequestedValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) RequestType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) RequestTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginCaCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate cloudflare_origin_ca_certificate} Resource.
func NewOriginCaCertificate(scope constructs.Construct, id *string, config *OriginCaCertificateConfig) OriginCaCertificate {
	_init_.Initialize()

	if err := validateNewOriginCaCertificateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginCaCertificate{}

	_jsii_.Create(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate cloudflare_origin_ca_certificate} Resource.
func NewOriginCaCertificate_Override(o OriginCaCertificate, scope constructs.Construct, id *string, config *OriginCaCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetCsr(val *string) {
	if err := j.validateSetCsrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csr",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetHostnames(val *[]*string) {
	if err := j.validateSetHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnames",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetMinDaysForRenewal(val *float64) {
	if err := j.validateSetMinDaysForRenewalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDaysForRenewal",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetRequestedValidity(val *float64) {
	if err := j.validateSetRequestedValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedValidity",
		val,
	)
}

func (j *jsiiProxy_OriginCaCertificate)SetRequestType(val *string) {
	if err := j.validateSetRequestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestType",
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
func OriginCaCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginCaCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OriginCaCertificate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginCaCertificate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OriginCaCertificate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginCaCertificate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OriginCaCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.originCaCertificate.OriginCaCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OriginCaCertificate) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OriginCaCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OriginCaCertificate) ResetCsr() {
	_jsii_.InvokeVoid(
		o,
		"resetCsr",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginCaCertificate) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginCaCertificate) ResetMinDaysForRenewal() {
	_jsii_.InvokeVoid(
		o,
		"resetMinDaysForRenewal",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginCaCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginCaCertificate) ResetRequestedValidity() {
	_jsii_.InvokeVoid(
		o,
		"resetRequestedValidity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginCaCertificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginCaCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

