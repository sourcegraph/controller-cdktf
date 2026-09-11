package googledeveloperconnectinsightsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectinsightsconfig/internal"
)

type GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference interface {
	cdktn.ComplexObject
	ArtifactRegistryPackage() *string
	SetArtifactRegistryPackage(val *string)
	ArtifactRegistryPackageInput() *string
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
	InternalValue() *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry
	SetInternalValue(val *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry)
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference
type jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ArtifactRegistryPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ArtifactRegistryPackageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryPackageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) InternalValue() *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry {
	var returns *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDeveloperConnectInsightsConfig.GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference_Override(g GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDeveloperConnectInsightsConfig.GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetArtifactRegistryPackage(val *string) {
	if err := j.validateSetArtifactRegistryPackageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactRegistryPackage",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetInternalValue(val *GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistry) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectInsightsConfigArtifactConfigsGoogleArtifactRegistryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

