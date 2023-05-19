package googlecloudrunv2job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2job/internal"
)

type GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference interface {
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
	InternalValue() *GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDir
	SetInternalValue(val *GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDir)
	Medium() *string
	SetMedium(val *string)
	MediumInput() *string
	SizeLimit() *string
	SetSizeLimit(val *string)
	SizeLimitInput() *string
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
	ResetMedium()
	ResetSizeLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference
type jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) InternalValue() *GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDir {
	var returns *GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDir
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) Medium() *string {
	var returns *string
	_jsii_.Get(
		j,
		"medium",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) MediumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) SizeLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) SizeLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference_Override(g GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudRunV2Job.GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetInternalValue(val *GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDir) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetMedium(val *string) {
	if err := j.validateSetMediumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"medium",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetSizeLimit(val *string) {
	if err := j.validateSetSizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ResetMedium() {
	_jsii_.InvokeVoid(
		g,
		"resetMedium",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ResetSizeLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetSizeLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2JobTemplateTemplateVolumesEmptyDirOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

