package googlenetworksecuritysecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworksecuritysecurityprofile/internal"
)

type GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference interface {
	cdktf.ComplexObject
	AntivirusOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileAntivirusOverridesList
	AntivirusOverridesInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkSecuritySecurityProfileThreatPreventionProfile
	SetInternalValue(val *GoogleNetworkSecuritySecurityProfileThreatPreventionProfile)
	SeverityOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesList
	SeverityOverridesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreatOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesList
	ThreatOverridesInput() interface{}
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
	PutAntivirusOverrides(value interface{})
	PutSeverityOverrides(value interface{})
	PutThreatOverrides(value interface{})
	ResetAntivirusOverrides()
	ResetSeverityOverrides()
	ResetThreatOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference
type jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) AntivirusOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileAntivirusOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileThreatPreventionProfileAntivirusOverridesList
	_jsii_.Get(
		j,
		"antivirusOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) AntivirusOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antivirusOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) InternalValue() *GoogleNetworkSecuritySecurityProfileThreatPreventionProfile {
	var returns *GoogleNetworkSecuritySecurityProfileThreatPreventionProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) SeverityOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesList
	_jsii_.Get(
		j,
		"severityOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) SeverityOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ThreatOverrides() GoogleNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesList {
	var returns GoogleNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesList
	_jsii_.Get(
		j,
		"threatOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ThreatOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatOverridesInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkSecuritySecurityProfile.GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference_Override(g GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkSecuritySecurityProfile.GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference)SetInternalValue(val *GoogleNetworkSecuritySecurityProfileThreatPreventionProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) PutAntivirusOverrides(value interface{}) {
	if err := g.validatePutAntivirusOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntivirusOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) PutSeverityOverrides(value interface{}) {
	if err := g.validatePutSeverityOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSeverityOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) PutThreatOverrides(value interface{}) {
	if err := g.validatePutThreatOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThreatOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ResetAntivirusOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetAntivirusOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ResetSeverityOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetSeverityOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ResetThreatOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetThreatOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecuritySecurityProfileThreatPreventionProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

