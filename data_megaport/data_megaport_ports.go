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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/megaport/terraform-provider-megaport/schema_megaport"
	"github.com/megaport/terraform-provider-megaport/terraform_utility"
)

func MegaportPorts() *schema.Resource {
	return &schema.Resource{
		Read:   dataMegaportPortsRead,
		Schema: schema_megaport.ResourcePortsSchema(),
	}
}

func dataMegaportPortsRead(d *schema.ResourceData, m interface{}) error {
	port := m.(*terraform_utility.MegaportClient).Port

	portId := d.Get("port_id").(string)
	d.SetId(portId)
	ports, retrievalErr := port.GetPorts()

	if retrievalErr != nil {
		return retrievalErr
	}

	var converted []*schema.ResourceData

	for _, port := range ports {
		dd := &schema.ResourceData{}
		dd.Set("id", port.UID)
		dd.Set("uid", port.UID)
		dd.Set("port_name", port.Name)
		dd.Set("type", port.Type)
		dd.Set("provisioning_status", port.ProvisioningStatus)
		dd.Set("create_date", port.CreateDate)
		dd.Set("created_by", port.CreatedBy)
		dd.Set("port_speed", port.PortSpeed)
		dd.Set("live_date", port.LiveDate)
		dd.Set("market_code", port.Market)
		dd.Set("location_id", port.LocationID)
		dd.Set("marketplace_visibility", port.MarketplaceVisibility)
		dd.Set("company_name", port.CompanyName)
		dd.Set("term", port.ContractTermMonths)
		dd.Set("lag_primary", port.LAGPrimary)
		dd.Set("lag_id", port.LAGID)
		dd.Set("locked", port.Locked)
		dd.Set("admin_locked", port.AdminLocked)
		converted = append(converted, dd)
	}

	d.Set("id", 19)
	d.Set("ports", converted)

	return nil
}
