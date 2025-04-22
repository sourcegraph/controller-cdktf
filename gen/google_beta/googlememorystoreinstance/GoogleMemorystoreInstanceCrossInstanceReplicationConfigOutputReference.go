package googlememorystoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlememorystoreinstance/internal"
)

type GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InstanceRole() *string
	SetInstanceRole(val *string)
	InstanceRoleInput() *string
	InternalValue() *GoogleMemorystoreInstanceCrossInstanceReplicationConfig
	SetInternalValue(val *GoogleMemorystoreInstanceCrossInstanceReplicationConfig)
	Membership() GoogleMemorystoreInstanceCrossInstanceReplicationConfigMembershipList
	PrimaryInstance() GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference
	PrimaryInstanceInput() *GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance
	SecondaryInstances() GoogleMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList
	SecondaryInstancesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
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
	PutPrimaryInstance(value *GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance)
	PutSecondaryInstances(value interface{})
	ResetInstanceRole()
	ResetPrimaryInstance()
	ResetSecondaryInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference
type jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InstanceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InternalValue() *GoogleMemorystoreInstanceCrossInstanceReplicationConfig {
	var returns *GoogleMemorystoreInstanceCrossInstanceReplicationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Membership() GoogleMemorystoreInstanceCrossInstanceReplicationConfigMembershipList {
	var returns GoogleMemorystoreInstanceCrossInstanceReplicationConfigMembershipList
	_jsii_.Get(
		j,
		"membership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PrimaryInstance() GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference {
	var returns GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference
	_jsii_.Get(
		j,
		"primaryInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PrimaryInstanceInput() *GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance {
	var returns *GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance
	_jsii_.Get(
		j,
		"primaryInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) SecondaryInstances() GoogleMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList {
	var returns GoogleMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList
	_jsii_.Get(
		j,
		"secondaryInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) SecondaryInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewGoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference_Override(g GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMemorystoreInstance.GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetInstanceRole(val *string) {
	if err := j.validateSetInstanceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetInternalValue(val *GoogleMemorystoreInstanceCrossInstanceReplicationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PutPrimaryInstance(value *GoogleMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance) {
	if err := g.validatePutPrimaryInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrimaryInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PutSecondaryInstances(value interface{}) {
	if err := g.validatePutSecondaryInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetInstanceRole() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetPrimaryInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetSecondaryInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

