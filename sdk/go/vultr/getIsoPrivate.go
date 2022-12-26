// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about an ISO file uploaded to your Vultr account.
//
// ## Example Usage
//
// Get the information for a ISO file by `filename`:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.LookupIsoPrivate(ctx, &vultr.LookupIsoPrivateArgs{
//				Filters: []vultr.GetIsoPrivateFilter{
//					{
//						Name: "filename",
//						Values: []string{
//							"my-iso-filename",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIsoPrivate(ctx *pulumi.Context, args *LookupIsoPrivateArgs, opts ...pulumi.InvokeOption) (*LookupIsoPrivateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIsoPrivateResult
	err := ctx.Invoke("vultr:index/getIsoPrivate:getIsoPrivate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIsoPrivate.
type LookupIsoPrivateArgs struct {
	// Query parameters for finding ISO files.
	Filters []GetIsoPrivateFilter `pulumi:"filters"`
}

// A collection of values returned by getIsoPrivate.
type LookupIsoPrivateResult struct {
	// The date the ISO file was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The ISO file's filename.
	Filename string                `pulumi:"filename"`
	Filters  []GetIsoPrivateFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The md5 hash of the ISO file.
	Md5sum string `pulumi:"md5sum"`
	// The sha512 hash of the ISO file.
	Sha512sum string `pulumi:"sha512sum"`
	// The size of the ISO file in bytes.
	Size int `pulumi:"size"`
	// The status of the ISO file.
	Status string `pulumi:"status"`
}

func LookupIsoPrivateOutput(ctx *pulumi.Context, args LookupIsoPrivateOutputArgs, opts ...pulumi.InvokeOption) LookupIsoPrivateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIsoPrivateResult, error) {
			args := v.(LookupIsoPrivateArgs)
			r, err := LookupIsoPrivate(ctx, &args, opts...)
			var s LookupIsoPrivateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIsoPrivateResultOutput)
}

// A collection of arguments for invoking getIsoPrivate.
type LookupIsoPrivateOutputArgs struct {
	// Query parameters for finding ISO files.
	Filters GetIsoPrivateFilterArrayInput `pulumi:"filters"`
}

func (LookupIsoPrivateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsoPrivateArgs)(nil)).Elem()
}

// A collection of values returned by getIsoPrivate.
type LookupIsoPrivateResultOutput struct{ *pulumi.OutputState }

func (LookupIsoPrivateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIsoPrivateResult)(nil)).Elem()
}

func (o LookupIsoPrivateResultOutput) ToLookupIsoPrivateResultOutput() LookupIsoPrivateResultOutput {
	return o
}

func (o LookupIsoPrivateResultOutput) ToLookupIsoPrivateResultOutputWithContext(ctx context.Context) LookupIsoPrivateResultOutput {
	return o
}

// The date the ISO file was added to your Vultr account.
func (o LookupIsoPrivateResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The ISO file's filename.
func (o LookupIsoPrivateResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.Filename }).(pulumi.StringOutput)
}

func (o LookupIsoPrivateResultOutput) Filters() GetIsoPrivateFilterArrayOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) []GetIsoPrivateFilter { return v.Filters }).(GetIsoPrivateFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIsoPrivateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The md5 hash of the ISO file.
func (o LookupIsoPrivateResultOutput) Md5sum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.Md5sum }).(pulumi.StringOutput)
}

// The sha512 hash of the ISO file.
func (o LookupIsoPrivateResultOutput) Sha512sum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.Sha512sum }).(pulumi.StringOutput)
}

// The size of the ISO file in bytes.
func (o LookupIsoPrivateResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) int { return v.Size }).(pulumi.IntOutput)
}

// The status of the ISO file.
func (o LookupIsoPrivateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIsoPrivateResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIsoPrivateResultOutput{})
}
