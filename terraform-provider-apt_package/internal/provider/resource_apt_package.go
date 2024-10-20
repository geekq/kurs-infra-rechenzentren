package provider

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AptPackageResource struct{}

func (r *AptPackageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "aptpackage"
}

func (r *AptPackageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// No-op update: Nothing to do here since updates aren't supported for this resource
	// Will recreate the resource if the name or version changes
	// TODO Check for the `ForceNew` in schema,
	// is there a new way in the newest terraform framework?
}

func NewAptPackageResource() resource.Resource {
	return &AptPackageResource{}
}

// Define the schema for apt_package resource
func (r *AptPackageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Required:    true,
				Description: "Name or IP of the machine to connect to.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the package to install.",
			},
			"version": schema.StringAttribute{
				Optional:    true,
				Description: "The version of the package to install.",
			},
			"id": schema.StringAttribute{
				Computed: true, // This means Terraform will manage the ID
			},
		},
	}
}

// Create the package
func (r *AptPackageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AptPackageModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Install the package
	err := installPackage(plan.Name.ValueString(), plan.Version.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Installing Package",
			fmt.Sprintf("Could not install package %s: %s", plan.Name.ValueString(), err),
		)
		return
	}

	// Set the resource ID to the package name
	plan.ID = types.StringValue(plan.Name.ValueString())
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read the package state (check if installed)
func (r *AptPackageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AptPackageModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if the package is still installed
	if !isPackageInstalled(state.Name.ValueString()) {
		state.ID = types.StringNull()
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Delete the package
func (r *AptPackageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AptPackageModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Remove the package
	err := removePackage(state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing Package",
			fmt.Sprintf("Could not remove package %s: %s", state.Name.ValueString(), err),
		)
		return
	}

	state.ID = types.StringNull()
	resp.State.Set(ctx, state)
}

// Package installation logic
func installPackage(name, version string) error {
	var cmd *exec.Cmd
	if version != "" {
		cmd = exec.Command("sudo", "apt-get", "install", "-y", fmt.Sprintf("%s=%s", name, version))
	} else {
		cmd = exec.Command("sudo", "apt-get", "install", "-y", name)
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	// Check if the command failed
	if err != nil {
		// Convert the command to a string for context
		cmdStr := commandToString(cmd)

		// Return the enriched error
		return enrichCmdError(err, cmdStr, stderr.String())
	}

	// Return nil if no error occurred
	return nil
}

// Check if a package is installed
func isPackageInstalled(name string) bool {
	cmd := exec.Command("dpkg-query", "-W", "-f='${Status}'", name)
	output, err := cmd.Output()
	return err == nil && string(output) == "'install ok installed'"
}

// Remove the package
func removePackage(name string) error {
	cmd := exec.Command("sudo", "apt-get", "remove", "-y", name)
	return cmd.Run()
}

// enrichCmdError enriches the error with additional context (cmdStr and stderr output)
func enrichCmdError(origErr error, cmdStr, stderr string) error {
	// Add additional context to the original error and return it
	return fmt.Errorf("failed to run command: %s: %w. Stderr: %s", cmdStr, origErr, stderr)
}

// Convert exec.Command to string
func commandToString(cmd *exec.Cmd) string {
	// Join the command and arguments to form a single string
	return strings.Join(cmd.Args, " ")
}

// AptPackageModel represents the resource data structure
type AptPackageModel struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Version types.String `tfsdk:"version"`
	Server  types.String `tfsdk:"server"`
}
