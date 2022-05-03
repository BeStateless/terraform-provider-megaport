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

package data_megaport

import (
	"context"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/megaport/terraform-provider-megaport/schema_megaport"
	"github.com/megaport/terraform-provider-megaport/terraform_utility"
)

func MegaportLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataMegaportLocationsRead,
		Schema:      schema_megaport.DataLocationsSchema(),
	}
}

func dataMegaportLocationsRead(c context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	errors := make([]diag.Diagnostic, 0)
	location := m.(*terraform_utility.MegaportClient).Location
	locationsResponse, err := location.GetAllLocations()
	if err != nil {
		errors = append(errors, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed to get megaport locations list",
			Detail:        err.Error(),
			AttributePath: make([]cty.PathStep, 0),
		})
		return errors
	}

	locations := make([]map[string]interface{}, len(locationsResponse))

	for i, loc := range locationsResponse {
		location := make(map[string]interface{})
		location["address"] = loc.Address
		location["country"] = loc.Country
		location["has_mcr"] = loc.VRouterAvailable
		location["id"] = loc.ID
		location["latitude"] = loc.Latitude
		location["live_date"] = loc.LiveDate
		location["longitude"] = loc.Longitude
		location["market"] = loc.Market
		location["metro"] = loc.Metro
		location["name"] = loc.Name
		location["site_code"] = loc.SiteCode
		location["status"] = loc.Status
		locations[i] = location
	}

	id, err := uuid.GenerateUUID()

	if err != nil {
		errors = append(errors, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       err.Error(),
			Detail:        "Failed to generate UUID as Id",
			AttributePath: make([]cty.PathStep, 0),
		})
		return errors
	}
	d.Set("locations", &locations)
	d.SetId(id)
	return errors
}
