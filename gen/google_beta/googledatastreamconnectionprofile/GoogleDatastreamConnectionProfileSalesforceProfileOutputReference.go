package googledatastreamconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamconnectionprofile/internal"
)

type GoogleDatastreamConnectionProfileSalesforceProfileOutputReference interface {
	cdktf.ComplexObject
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
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDatastreamConnectionProfileSalesforceProfile
	SetInternalValue(val *GoogleDatastreamConnectionProfileSalesforceProfile)
	Oauth2ClientCredentials() GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentialsOutputReference
	Oauth2ClientCredentialsInput() *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserCredentials() GoogleDatastreamConnectionProfileSalesforceProfileUserCredentialsOutputReference
	UserCredentialsInput() *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials
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
	PutOauth2ClientCredentials(value *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials)
	PutUserCredentials(value *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials)
	ResetOauth2ClientCredentials()
	ResetUserCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamConnectionProfileSalesforceProfileOutputReference
type jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) InternalValue() *GoogleDatastreamConnectionProfileSalesforceProfile {
	var returns *GoogleDatastreamConnectionProfileSalesforceProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) Oauth2ClientCredentials() GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentialsOutputReference {
	var returns GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) Oauth2ClientCredentialsInput() *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials {
	var returns *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) UserCredentials() GoogleDatastreamConnectionProfileSalesforceProfileUserCredentialsOutputReference {
	var returns GoogleDatastreamConnectionProfileSalesforceProfileUserCredentialsOutputReference
	_jsii_.Get(
		j,
		"userCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) UserCredentialsInput() *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials {
	var returns *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials
	_jsii_.Get(
		j,
		"userCredentialsInput",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamConnectionProfileSalesforceProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamConnectionProfileSalesforceProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamConnectionProfileSalesforceProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileSalesforceProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamConnectionProfileSalesforceProfileOutputReference_Override(g GoogleDatastreamConnectionProfileSalesforceProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamConnectionProfile.GoogleDatastreamConnectionProfileSalesforceProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetInternalValue(val *GoogleDatastreamConnectionProfileSalesforceProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) PutOauth2ClientCredentials(value *GoogleDatastreamConnectionProfileSalesforceProfileOauth2ClientCredentials) {
	if err := g.validatePutOauth2ClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2ClientCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) PutUserCredentials(value *GoogleDatastreamConnectionProfileSalesforceProfileUserCredentials) {
	if err := g.validatePutUserCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ResetOauth2ClientCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ResetUserCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetUserCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamConnectionProfileSalesforceProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

