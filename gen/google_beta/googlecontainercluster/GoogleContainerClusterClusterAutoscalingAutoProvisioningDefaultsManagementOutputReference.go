package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference interface {
	cdktf.ComplexObject
	AutoRepair() interface{}
	SetAutoRepair(val interface{})
	AutoRepairInput() interface{}
	AutoUpgrade() interface{}
	SetAutoUpgrade(val interface{})
	AutoUpgradeInput() interface{}
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
	InternalValue() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	SetInternalValue(val *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeOptions() GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList
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
	ResetAutoRepair()
	ResetAutoUpgrade()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference
type jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoRepair() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoRepairInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InternalValue() *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement {
	var returns *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) UpgradeOptions() GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList {
	var returns GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList
	_jsii_.Get(
		j,
		"upgradeOptions",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference_Override(g GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetAutoRepair(val interface{}) {
	if err := j.validateSetAutoRepairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRepair",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetAutoUpgrade(val interface{}) {
	if err := j.validateSetAutoUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpgrade",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetInternalValue(val *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ResetAutoRepair() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRepair",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ResetAutoUpgrade() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoUpgrade",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

