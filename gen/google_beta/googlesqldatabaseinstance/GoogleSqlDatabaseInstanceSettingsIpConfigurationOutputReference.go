package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllocatedIpRange() *string
	SetAllocatedIpRange(val *string)
	AllocatedIpRangeInput() *string
	AuthorizedNetworks() GoogleSqlDatabaseInstanceSettingsIpConfigurationAuthorizedNetworksList
	AuthorizedNetworksInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnablePrivatePathForGoogleCloudServices() interface{}
	SetEnablePrivatePathForGoogleCloudServices(val interface{})
	EnablePrivatePathForGoogleCloudServicesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceSettingsIpConfiguration
	SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsIpConfiguration)
	Ipv4Enabled() interface{}
	SetIpv4Enabled(val interface{})
	Ipv4EnabledInput() interface{}
	PrivateNetwork() *string
	SetPrivateNetwork(val *string)
	PrivateNetworkInput() *string
	RequireSsl() interface{}
	SetRequireSsl(val interface{})
	RequireSslInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAuthorizedNetworks(value interface{})
	ResetAllocatedIpRange()
	ResetAuthorizedNetworks()
	ResetEnablePrivatePathForGoogleCloudServices()
	ResetIpv4Enabled()
	ResetPrivateNetwork()
	ResetRequireSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) AllocatedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) AllocatedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) AuthorizedNetworks() GoogleSqlDatabaseInstanceSettingsIpConfigurationAuthorizedNetworksList {
	var returns GoogleSqlDatabaseInstanceSettingsIpConfigurationAuthorizedNetworksList
	_jsii_.Get(
		j,
		"authorizedNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) AuthorizedNetworksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizedNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) EnablePrivatePathForGoogleCloudServices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatePathForGoogleCloudServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) EnablePrivatePathForGoogleCloudServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivatePathForGoogleCloudServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) InternalValue() *GoogleSqlDatabaseInstanceSettingsIpConfiguration {
	var returns *GoogleSqlDatabaseInstanceSettingsIpConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) Ipv4Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) Ipv4EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) PrivateNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) PrivateNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) RequireSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) RequireSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetAllocatedIpRange(val *string) {
	if err := j.validateSetAllocatedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedIpRange",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetEnablePrivatePathForGoogleCloudServices(val interface{}) {
	if err := j.validateSetEnablePrivatePathForGoogleCloudServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivatePathForGoogleCloudServices",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsIpConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetIpv4Enabled(val interface{}) {
	if err := j.validateSetIpv4EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetPrivateNetwork(val *string) {
	if err := j.validateSetPrivateNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetRequireSsl(val interface{}) {
	if err := j.validateSetRequireSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSsl",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) PutAuthorizedNetworks(value interface{}) {
	if err := g.validatePutAuthorizedNetworksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizedNetworks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetAllocatedIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetAllocatedIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetAuthorizedNetworks() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizedNetworks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetEnablePrivatePathForGoogleCloudServices() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePrivatePathForGoogleCloudServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetIpv4Enabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv4Enabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetPrivateNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ResetRequireSsl() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireSsl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

