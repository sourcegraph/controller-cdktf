package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference interface {
	cdktf.ComplexObject
	CloudComposerConnectionSubnetwork() *string
	SetCloudComposerConnectionSubnetwork(val *string)
	CloudComposerConnectionSubnetworkInput() *string
	CloudComposerNetworkIpv4CidrBlock() *string
	SetCloudComposerNetworkIpv4CidrBlock(val *string)
	CloudComposerNetworkIpv4CidrBlockInput() *string
	CloudSqlIpv4CidrBlock() *string
	SetCloudSqlIpv4CidrBlock(val *string)
	CloudSqlIpv4CidrBlockInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnablePrivateEndpoint() interface{}
	SetEnablePrivateEndpoint(val interface{})
	EnablePrivateEndpointInput() interface{}
	EnablePrivatelyUsedPublicIps() interface{}
	SetEnablePrivatelyUsedPublicIps(val interface{})
	EnablePrivatelyUsedPublicIpsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig
	SetInternalValue(val *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig)
	MasterIpv4CidrBlock() *string
	SetMasterIpv4CidrBlock(val *string)
	MasterIpv4CidrBlockInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebServerIpv4CidrBlock() *string
	SetWebServerIpv4CidrBlock(val *string)
	WebServerIpv4CidrBlockInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCloudComposerConnectionSubnetwork()
	ResetCloudComposerNetworkIpv4CidrBlock()
	ResetCloudSqlIpv4CidrBlock()
	ResetConnectionType()
	ResetEnablePrivateEndpoint()
	ResetEnablePrivatelyUsedPublicIps()
	ResetMasterIpv4CidrBlock()
	ResetWebServerIpv4CidrBlock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
type jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerConnectionSubnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerConnectionSubnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerConnectionSubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerConnectionSubnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerNetworkIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerNetworkIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudComposerNetworkIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudComposerNetworkIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudSqlIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CloudSqlIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSqlIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivateEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivateEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivatelyUsedPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatelyUsedPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) EnablePrivatelyUsedPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatelyUsedPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InternalValue() *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig {
	var returns *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) MasterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) MasterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) WebServerIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webServerIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) WebServerIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webServerIpv4CidrBlockInput",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference_Override(g GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudComposerConnectionSubnetwork(val *string) {
	if err := j.validateSetCloudComposerConnectionSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudComposerConnectionSubnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudComposerNetworkIpv4CidrBlock(val *string) {
	if err := j.validateSetCloudComposerNetworkIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudComposerNetworkIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetCloudSqlIpv4CidrBlock(val *string) {
	if err := j.validateSetCloudSqlIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudSqlIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetEnablePrivateEndpoint(val interface{}) {
	if err := j.validateSetEnablePrivateEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetEnablePrivatelyUsedPublicIps(val interface{}) {
	if err := j.validateSetEnablePrivatelyUsedPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivatelyUsedPublicIps",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetInternalValue(val *GoogleComposerEnvironmentConfigPrivateEnvironmentConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetMasterIpv4CidrBlock(val *string) {
	if err := j.validateSetMasterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference)SetWebServerIpv4CidrBlock(val *string) {
	if err := j.validateSetWebServerIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webServerIpv4CidrBlock",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudComposerConnectionSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudComposerConnectionSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudComposerNetworkIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudComposerNetworkIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetCloudSqlIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSqlIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetConnectionType() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetEnablePrivateEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePrivateEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetEnablePrivatelyUsedPublicIps() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePrivatelyUsedPublicIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetMasterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ResetWebServerIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetWebServerIpv4CidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

