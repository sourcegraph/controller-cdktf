package certificatepack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/certificatepack/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack cloudflare_certificate_pack}.
type CertificatePack interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateAuthority() *string
	SetCertificateAuthority(val *string)
	CertificateAuthorityInput() *string
	CloudflareBranding() interface{}
	SetCloudflareBranding(val interface{})
	CloudflareBrandingInput() interface{}
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
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidationErrors() CertificatePackValidationErrorsList
	ValidationErrorsInput() interface{}
	ValidationMethod() *string
	SetValidationMethod(val *string)
	ValidationMethodInput() *string
	ValidationRecords() CertificatePackValidationRecordsList
	ValidationRecordsInput() interface{}
	ValidityDays() *float64
	SetValidityDays(val *float64)
	ValidityDaysInput() *float64
	WaitForActiveStatus() interface{}
	SetWaitForActiveStatus(val interface{})
	WaitForActiveStatusInput() interface{}
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
	PutValidationErrors(value interface{})
	PutValidationRecords(value interface{})
	ResetCloudflareBranding()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetValidationErrors()
	ResetValidationRecords()
	ResetWaitForActiveStatus()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CertificatePack
type jsiiProxy_CertificatePack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CertificatePack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) CertificateAuthority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) CertificateAuthorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) CloudflareBranding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudflareBranding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) CloudflareBrandingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudflareBrandingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationErrors() CertificatePackValidationErrorsList {
	var returns CertificatePackValidationErrorsList
	_jsii_.Get(
		j,
		"validationErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationRecords() CertificatePackValidationRecordsList {
	var returns CertificatePackValidationRecordsList
	_jsii_.Get(
		j,
		"validationRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidationRecordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidityDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ValidityDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) WaitForActiveStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForActiveStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) WaitForActiveStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForActiveStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePack) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack cloudflare_certificate_pack} Resource.
func NewCertificatePack(scope constructs.Construct, id *string, config *CertificatePackConfig) CertificatePack {
	_init_.Initialize()

	if err := validateNewCertificatePackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificatePack{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack cloudflare_certificate_pack} Resource.
func NewCertificatePack_Override(c CertificatePack, scope constructs.Construct, id *string, config *CertificatePackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CertificatePack)SetCertificateAuthority(val *string) {
	if err := j.validateSetCertificateAuthorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthority",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetCloudflareBranding(val interface{}) {
	if err := j.validateSetCloudflareBrandingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudflareBranding",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetHosts(val *[]*string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetValidationMethod(val *string) {
	if err := j.validateSetValidationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationMethod",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetValidityDays(val *float64) {
	if err := j.validateSetValidityDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validityDays",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetWaitForActiveStatus(val interface{}) {
	if err := j.validateSetWaitForActiveStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForActiveStatus",
		val,
	)
}

func (j *jsiiProxy_CertificatePack)SetZoneId(val *string) {
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
func CertificatePack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificatePack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CertificatePack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificatePack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CertificatePack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificatePack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CertificatePack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.certificatePack.CertificatePack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CertificatePack) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CertificatePack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CertificatePack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CertificatePack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CertificatePack) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CertificatePack) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CertificatePack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CertificatePack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CertificatePack) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CertificatePack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CertificatePack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CertificatePack) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CertificatePack) PutValidationErrors(value interface{}) {
	if err := c.validatePutValidationErrorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putValidationErrors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CertificatePack) PutValidationRecords(value interface{}) {
	if err := c.validatePutValidationRecordsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putValidationRecords",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CertificatePack) ResetCloudflareBranding() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudflareBranding",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) ResetValidationErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetValidationErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) ResetValidationRecords() {
	_jsii_.InvokeVoid(
		c,
		"resetValidationRecords",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) ResetWaitForActiveStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForActiveStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

