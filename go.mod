module github.com/megaport/terraform-provider-megaport

// For local development
//replace github.com/megaport/megaportgo => <megaportgo local directory>
// vxc-wait-retry-with-context-clean
replace github.com/megaport/megaportgo => github.com/bestateless/megaportgo v0.1.11-stateless.0.20230522224628-211fd5db81ec

go 1.13

require (
	github.com/aws/aws-sdk-go v1.44.267
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-uuid v1.0.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.26.1
	github.com/megaport/megaportgo v0.1.15-beta
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.16.2 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.15.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.8.0 // indirect
	github.com/hashicorp/terraform-svchost v0.1.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/xlzd/gotp v0.1.0 // indirect
	github.com/zclconf/go-cty v1.13.2 // indirect
	golang.org/x/net v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
)
