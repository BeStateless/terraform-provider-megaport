// Copyright 2020 Megaport Pty Ltd
//
// Licensed under the Mozilla Public License, Version 2.0 (the
// "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//       https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema_megaport

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataPartnerPortSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connect_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"product_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"location_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"company_uid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"speed": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"diversity_zone": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
	}
}
