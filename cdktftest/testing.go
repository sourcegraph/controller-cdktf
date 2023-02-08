package cdktftest

import (
	"github.com/aws/jsii-runtime-go"
	hashicorpcdtkf "github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computenetwork"
	"github.com/sourcegraph/controller-cdktf/gen/google/computesubnetwork"
	"github.com/sourcegraph/controller-cdktf/gen/google/project"
	"github.com/sourcegraph/controller-cdktf/gen/google/serviceaccount"
)

var (
	// App creates an app for testing.
	NewApp = hashicorpcdtkf.Testing_App
	// Synth renders a CDKTF stack that can be provided to various assertions.
	SynthStack = hashicorpcdtkf.Testing_Synth
)

func NewStack(name string, opts *hashicorpcdtkf.TestingAppOptions) hashicorpcdtkf.TerraformStack {
	return hashicorpcdtkf.NewTerraformStack(NewApp(opts), &name)
}

func NewStackWithApp(name string, app hashicorpcdtkf.App) hashicorpcdtkf.TerraformStack {
	return hashicorpcdtkf.NewTerraformStack(app, &name)
}

// NewProject creates a new project for testing.
func NewProject(opts *project.ProjectConfig) project.Project {
	if opts == nil {
		opts = &project.ProjectConfig{
			ProjectId: jsii.String("1234"),
			Name:      jsii.String("robertlin"),
		}
	}
	return project.NewProject(
		NewStack("project", nil),
		jsii.String("project"),
		opts,
	)
}

// NewNetwork creates a new network for testing.
func NewNetwork(p project.Project, opts *computenetwork.ComputeNetworkConfig) computenetwork.ComputeNetwork {
	if opts == nil {
		opts = &computenetwork.ComputeNetworkConfig{
			Name:    jsii.String("network"),
			Project: p.ProjectId(),
		}
	}
	return computenetwork.NewComputeNetwork(
		NewStack("network", nil),
		jsii.String("network"),
		opts,
	)
}

// NewNetwork creates a new network for testing.
func NewSubNetwork(p project.Project, n computenetwork.ComputeNetwork, opts *computesubnetwork.ComputeSubnetworkConfig) computesubnetwork.ComputeSubnetwork {
	if opts == nil {
		opts = &computesubnetwork.ComputeSubnetworkConfig{
			Name:        jsii.String("subnetwork"),
			Project:     p.ProjectId(),
			Network:     n.Id(),
			IpCidrRange: jsii.String("10.0.1.0/24"),
		}
	}
	return computesubnetwork.NewComputeSubnetwork(
		NewStack("subnetwork", nil),
		jsii.String("subnetwork"),
		opts,
	)
}

func NewServiceAccount(p project.Project, opts *serviceaccount.ServiceAccountConfig) serviceaccount.ServiceAccount {
	if opts == nil {
		opts = &serviceaccount.ServiceAccountConfig{
			AccountId: jsii.String("sa"),
			Project:   p.ProjectId(),
		}
	}
	return serviceaccount.NewServiceAccount(
		NewStack("sa", nil),
		jsii.String("sa"),
		opts,
	)
}

// Assertions all accept a synthesized CDKTF stack with cdktftest.Synth as the first
// argument received.
var (
	AssertResource               = hashicorpcdtkf.Testing_ToHaveResource
	AssertResourceWithProperties = hashicorpcdtkf.Testing_ToHaveResourceWithProperties
)
