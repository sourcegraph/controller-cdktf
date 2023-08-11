package googlecertificatemanagercertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecertificatemanagercertificate/internal"
)

type GoogleCertificateManagerCertificateManagedOutputReference interface {
	cdktf.ComplexObject
	AuthorizationAttemptInfo() GoogleCertificateManagerCertificateManagedAuthorizationAttemptInfoList
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
	DnsAuthorizations() *[]*string
	SetDnsAuthorizations(val *[]*string)
	DnsAuthorizationsInput() *[]*string
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCertificateManagerCertificateManaged
	SetInternalValue(val *GoogleCertificateManagerCertificateManaged)
	ProvisioningIssue() GoogleCertificateManagerCertificateManagedProvisioningIssueList
	State() *string
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
	ResetDnsAuthorizations()
	ResetDomains()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCertificateManagerCertificateManagedOutputReference
type jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) AuthorizationAttemptInfo() GoogleCertificateManagerCertificateManagedAuthorizationAttemptInfoList {
	var returns GoogleCertificateManagerCertificateManagedAuthorizationAttemptInfoList
	_jsii_.Get(
		j,
		"authorizationAttemptInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) DnsAuthorizations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAuthorizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) DnsAuthorizationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsAuthorizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) InternalValue() *GoogleCertificateManagerCertificateManaged {
	var returns *GoogleCertificateManagerCertificateManaged
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ProvisioningIssue() GoogleCertificateManagerCertificateManagedProvisioningIssueList {
	var returns GoogleCertificateManagerCertificateManagedProvisioningIssueList
	_jsii_.Get(
		j,
		"provisioningIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCertificateManagerCertificateManagedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCertificateManagerCertificateManagedOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCertificateManagerCertificateManagedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCertificateManagerCertificate.GoogleCertificateManagerCertificateManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCertificateManagerCertificateManagedOutputReference_Override(g GoogleCertificateManagerCertificateManagedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCertificateManagerCertificate.GoogleCertificateManagerCertificateManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetDnsAuthorizations(val *[]*string) {
	if err := j.validateSetDnsAuthorizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsAuthorizations",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetDomains(val *[]*string) {
	if err := j.validateSetDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetInternalValue(val *GoogleCertificateManagerCertificateManaged) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ResetDnsAuthorizations() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsAuthorizations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ResetDomains() {
	_jsii_.InvokeVoid(
		g,
		"resetDomains",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCertificateManagerCertificateManagedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

