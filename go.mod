module github.com/megaport/terraform-provider-megaport

// For local development
//replace github.com/megaport/megaportgo => <megaportgo local directory>

go 1.13

require (
	github.com/aws/aws-sdk-go v1.44.263
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-uuid v1.0.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.26.1
	github.com/megaport/megaportgo v0.1.15-beta
)
