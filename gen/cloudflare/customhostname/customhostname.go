package customhostname

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/customhostname/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname cloudflare_custom_hostname}.
type CustomHostname interface {
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
	CustomMetadata() *map[string]*string
	SetCustomMetadata(val *map[string]*string)
	CustomMetadataInput() *map[string]*string
	CustomOriginServer() *string
	SetCustomOriginServer(val *string)
	CustomOriginServerInput() *string
	CustomOriginSni() *string
	SetCustomOriginSni(val *string)
	CustomOriginSniInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OwnershipVerification() cdktf.StringMap
	OwnershipVerificationHttp() cdktf.StringMap
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
	Ssl() CustomHostnameSslList
	SslInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WaitForSslPendingValidation() interface{}
	SetWaitForSslPendingValidation(val interface{})
	WaitForSslPendingValidationInput() interface{}
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutSsl(value interface{})
	ResetCustomMetadata()
	ResetCustomOriginServer()
	ResetCustomOriginSni()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSsl()
	ResetWaitForSslPendingValidation()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CustomHostname
type jsiiProxy_CustomHostname struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CustomHostname) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomOriginServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomOriginServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomOriginSni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) CustomOriginSniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) OwnershipVerification() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"ownershipVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) OwnershipVerificationHttp() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"ownershipVerificationHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Ssl() CustomHostnameSslList {
	var returns CustomHostnameSslList
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) SslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) WaitForSslPendingValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSslPendingValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) WaitForSslPendingValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSslPendingValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostname) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname cloudflare_custom_hostname} Resource.
func NewCustomHostname(scope constructs.Construct, id *string, config *CustomHostnameConfig) CustomHostname {
	_init_.Initialize()

	if err := validateNewCustomHostnameParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomHostname{}

	_jsii_.Create(
		"cloudflare.customHostname.CustomHostname",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname cloudflare_custom_hostname} Resource.
func NewCustomHostname_Override(c CustomHostname, scope constructs.Construct, id *string, config *CustomHostnameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.customHostname.CustomHostname",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CustomHostname)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetCustomMetadata(val *map[string]*string) {
	if err := j.validateSetCustomMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMetadata",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetCustomOriginServer(val *string) {
	if err := j.validateSetCustomOriginServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customOriginServer",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetCustomOriginSni(val *string) {
	if err := j.validateSetCustomOriginSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customOriginSni",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetWaitForSslPendingValidation(val interface{}) {
	if err := j.validateSetWaitForSslPendingValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSslPendingValidation",
		val,
	)
}

func (j *jsiiProxy_CustomHostname)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
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
func CustomHostname_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomHostname_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.customHostname.CustomHostname",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CustomHostname_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomHostname_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.customHostname.CustomHostname",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CustomHostname_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomHostname_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.customHostname.CustomHostname",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CustomHostname_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.customHostname.CustomHostname",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CustomHostname) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CustomHostname) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CustomHostname) PutSsl(value interface{}) {
	if err := c.validatePutSslParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSsl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomHostname) ResetCustomMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetCustomOriginServer() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomOriginServer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetCustomOriginSni() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomOriginSni",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetSsl() {
	_jsii_.InvokeVoid(
		c,
		"resetSsl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) ResetWaitForSslPendingValidation() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForSslPendingValidation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostname) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostname) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

