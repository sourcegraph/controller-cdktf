package googlecloudsecuritycomplianceframeworkdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudsecuritycomplianceframeworkdeployment/internal"
)

type GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference interface {
	cdktn.ComplexObject
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
	FolderCreationConfig() GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfigOutputReference
	FolderCreationConfigInput() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig
	SetInternalValue(val *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig)
	ProjectCreationConfig() GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference
	ProjectCreationConfigInput() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutFolderCreationConfig(value *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig)
	PutProjectCreationConfig(value *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig)
	ResetFolderCreationConfig()
	ResetProjectCreationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference
type jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) FolderCreationConfig() GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfigOutputReference {
	var returns GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfigOutputReference
	_jsii_.Get(
		j,
		"folderCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) FolderCreationConfigInput() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig {
	var returns *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig
	_jsii_.Get(
		j,
		"folderCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) InternalValue() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig {
	var returns *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ProjectCreationConfig() GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference {
	var returns GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference
	_jsii_.Get(
		j,
		"projectCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ProjectCreationConfigInput() *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig {
	var returns *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig
	_jsii_.Get(
		j,
		"projectCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudSecurityComplianceFrameworkDeployment.GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference_Override(g GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudSecurityComplianceFrameworkDeployment.GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference)SetInternalValue(val *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) PutFolderCreationConfig(value *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig) {
	if err := g.validatePutFolderCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFolderCreationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) PutProjectCreationConfig(value *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig) {
	if err := g.validatePutProjectCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProjectCreationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ResetFolderCreationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetFolderCreationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ResetProjectCreationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetProjectCreationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

