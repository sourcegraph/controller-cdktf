package googleaccesscontextmanagerserviceperimeters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleaccesscontextmanagerserviceperimeters/internal"
)

type GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference interface {
	cdktf.ComplexObject
	AccessLevels() *[]*string
	SetAccessLevels(val *[]*string)
	AccessLevelsInput() *[]*string
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
	EgressPolicies() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesList
	EgressPoliciesInput() interface{}
	// Experimental.
	Fqn() *string
	IngressPolicies() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesList
	IngressPoliciesInput() interface{}
	InternalValue() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpec
	SetInternalValue(val *GoogleAccessContextManagerServicePerimetersServicePerimetersSpec)
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	RestrictedServices() *[]*string
	SetRestrictedServices(val *[]*string)
	RestrictedServicesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcAccessibleServices() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesOutputReference
	VpcAccessibleServicesInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServices
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
	PutEgressPolicies(value interface{})
	PutIngressPolicies(value interface{})
	PutVpcAccessibleServices(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServices)
	ResetAccessLevels()
	ResetEgressPolicies()
	ResetIngressPolicies()
	ResetResources()
	ResetRestrictedServices()
	ResetVpcAccessibleServices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference
type jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) AccessLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) AccessLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) EgressPolicies() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesList {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesList
	_jsii_.Get(
		j,
		"egressPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) EgressPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"egressPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) IngressPolicies() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesList {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesList
	_jsii_.Get(
		j,
		"ingressPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) IngressPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingressPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) InternalValue() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpec {
	var returns *GoogleAccessContextManagerServicePerimetersServicePerimetersSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) RestrictedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) RestrictedServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) VpcAccessibleServices() GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesOutputReference {
	var returns GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServicesOutputReference
	_jsii_.Get(
		j,
		"vpcAccessibleServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) VpcAccessibleServicesInput() *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServices {
	var returns *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServices
	_jsii_.Get(
		j,
		"vpcAccessibleServicesInput",
		&returns,
	)
	return returns
}


func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference_Override(g GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleAccessContextManagerServicePerimeters.GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetAccessLevels(val *[]*string) {
	if err := j.validateSetAccessLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLevels",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetInternalValue(val *GoogleAccessContextManagerServicePerimetersServicePerimetersSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetRestrictedServices(val *[]*string) {
	if err := j.validateSetRestrictedServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedServices",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) PutEgressPolicies(value interface{}) {
	if err := g.validatePutEgressPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEgressPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) PutIngressPolicies(value interface{}) {
	if err := g.validatePutIngressPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngressPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) PutVpcAccessibleServices(value *GoogleAccessContextManagerServicePerimetersServicePerimetersSpecVpcAccessibleServices) {
	if err := g.validatePutVpcAccessibleServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccessibleServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetAccessLevels() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessLevels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetEgressPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetEgressPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetIngressPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetRestrictedServices() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ResetVpcAccessibleServices() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessibleServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimetersServicePerimetersSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

