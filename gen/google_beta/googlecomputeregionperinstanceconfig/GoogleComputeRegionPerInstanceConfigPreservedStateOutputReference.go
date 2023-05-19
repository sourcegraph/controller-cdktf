package googlecomputeregionperinstanceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionperinstanceconfig/internal"
)

type GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference interface {
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
	Disk() GoogleComputeRegionPerInstanceConfigPreservedStateDiskList
	DiskInput() interface{}
	ExternalIp() GoogleComputeRegionPerInstanceConfigPreservedStateExternalIpList
	ExternalIpInput() interface{}
	// Experimental.
	Fqn() *string
	InternalIp() GoogleComputeRegionPerInstanceConfigPreservedStateInternalIpList
	InternalIpInput() interface{}
	InternalValue() *GoogleComputeRegionPerInstanceConfigPreservedState
	SetInternalValue(val *GoogleComputeRegionPerInstanceConfigPreservedState)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
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
	PutDisk(value interface{})
	PutExternalIp(value interface{})
	PutInternalIp(value interface{})
	ResetDisk()
	ResetExternalIp()
	ResetInternalIp()
	ResetMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference
type jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) Disk() GoogleComputeRegionPerInstanceConfigPreservedStateDiskList {
	var returns GoogleComputeRegionPerInstanceConfigPreservedStateDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ExternalIp() GoogleComputeRegionPerInstanceConfigPreservedStateExternalIpList {
	var returns GoogleComputeRegionPerInstanceConfigPreservedStateExternalIpList
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ExternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) InternalIp() GoogleComputeRegionPerInstanceConfigPreservedStateInternalIpList {
	var returns GoogleComputeRegionPerInstanceConfigPreservedStateInternalIpList
	_jsii_.Get(
		j,
		"internalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) InternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) InternalValue() *GoogleComputeRegionPerInstanceConfigPreservedState {
	var returns *GoogleComputeRegionPerInstanceConfigPreservedState
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionPerInstanceConfigPreservedStateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionPerInstanceConfigPreservedStateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionPerInstanceConfig.GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionPerInstanceConfigPreservedStateOutputReference_Override(g GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionPerInstanceConfig.GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetInternalValue(val *GoogleComputeRegionPerInstanceConfigPreservedState) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) PutDisk(value interface{}) {
	if err := g.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) PutExternalIp(value interface{}) {
	if err := g.validatePutExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) PutInternalIp(value interface{}) {
	if err := g.validatePutInternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ResetDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ResetExternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ResetInternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionPerInstanceConfigPreservedStateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

