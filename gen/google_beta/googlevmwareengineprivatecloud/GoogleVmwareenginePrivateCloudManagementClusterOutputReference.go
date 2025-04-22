package googlevmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareengineprivatecloud/internal"
)

type GoogleVmwareenginePrivateCloudManagementClusterOutputReference interface {
	cdktf.ComplexObject
	AutoscalingSettings() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference
	AutoscalingSettingsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	InternalValue() *GoogleVmwareenginePrivateCloudManagementCluster
	SetInternalValue(val *GoogleVmwareenginePrivateCloudManagementCluster)
	NodeTypeConfigs() GoogleVmwareenginePrivateCloudManagementClusterNodeTypeConfigsList
	NodeTypeConfigsInput() interface{}
	StretchedClusterConfig() GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference
	StretchedClusterConfigInput() *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig
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
	PutAutoscalingSettings(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings)
	PutNodeTypeConfigs(value interface{})
	PutStretchedClusterConfig(value *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig)
	ResetAutoscalingSettings()
	ResetNodeTypeConfigs()
	ResetStretchedClusterConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareenginePrivateCloudManagementClusterOutputReference
type jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) AutoscalingSettings() GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference {
	var returns GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference
	_jsii_.Get(
		j,
		"autoscalingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) AutoscalingSettingsInput() *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings
	_jsii_.Get(
		j,
		"autoscalingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) InternalValue() *GoogleVmwareenginePrivateCloudManagementCluster {
	var returns *GoogleVmwareenginePrivateCloudManagementCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) NodeTypeConfigs() GoogleVmwareenginePrivateCloudManagementClusterNodeTypeConfigsList {
	var returns GoogleVmwareenginePrivateCloudManagementClusterNodeTypeConfigsList
	_jsii_.Get(
		j,
		"nodeTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) NodeTypeConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) StretchedClusterConfig() GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference {
	var returns GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfigOutputReference
	_jsii_.Get(
		j,
		"stretchedClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) StretchedClusterConfigInput() *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig {
	var returns *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig
	_jsii_.Get(
		j,
		"stretchedClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareenginePrivateCloudManagementClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVmwareenginePrivateCloudManagementClusterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareenginePrivateCloudManagementClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVmwareenginePrivateCloudManagementClusterOutputReference_Override(g GoogleVmwareenginePrivateCloudManagementClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareenginePrivateCloud.GoogleVmwareenginePrivateCloudManagementClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetInternalValue(val *GoogleVmwareenginePrivateCloudManagementCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) PutAutoscalingSettings(value *GoogleVmwareenginePrivateCloudManagementClusterAutoscalingSettings) {
	if err := g.validatePutAutoscalingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) PutNodeTypeConfigs(value interface{}) {
	if err := g.validatePutNodeTypeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeTypeConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) PutStretchedClusterConfig(value *GoogleVmwareenginePrivateCloudManagementClusterStretchedClusterConfig) {
	if err := g.validatePutStretchedClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStretchedClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ResetAutoscalingSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ResetNodeTypeConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeTypeConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ResetStretchedClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetStretchedClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareenginePrivateCloudManagementClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

