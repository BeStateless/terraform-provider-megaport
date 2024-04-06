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
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/megaport/megaportgo/types"
	"github.com/megaport/terraform-provider-megaport/schema_megaport"
	"github.com/megaport/terraform-provider-megaport/terraform_utility"
)

func MegaportPorts() *schema.Resource {
	return &schema.Resource{
		Read:   dataMegaportPortsRead,
		Schema: schema_megaport.DataPortsSchema(),
	}
}

func dataMegaportPortsRead(d *schema.ResourceData, m interface{}) error {
	port := m.(*terraform_utility.MegaportClient).Port

	id, err := uuid.GenerateUUID()

	if err != nil {
		return err
	}

	ports, retrievalErr := port.GetPorts()

	if retrievalErr != nil {
		return retrievalErr
	}

	converted := make([]map[string]interface{}, len(ports))

	for i, port := range ports {
		converted[i] = tfizePort(port)
	}

	d.SetId(id)

	return d.Set("ports", &converted)
}

func tfizePort(port types.Port) map[string]interface{} {
	tf := make(map[string]interface{})
	tf["admin_locked"] = port.AdminLocked
	tf["company_name"] = port.CompanyName
	tf["create_date"] = port.CreateDate
	tf["created_by"] = port.CreatedBy
	tf["lag_id"] = port.LAGID
	tf["lag_primary"] = port.LAGPrimary
	tf["live_date"] = port.LiveDate
	tf["location_id"] = port.LocationID
	tf["locked"] = port.Locked
	tf["market_code"] = port.Market
	tf["marketplace_visibility"] = port.MarketplaceVisibility
	tf["port_name"] = port.Name
	tf["port_speed"] = port.PortSpeed
	tf["provisioning_status"] = port.ProvisioningStatus
	tf["term"] = port.ContractTermMonths
	tf["type"] = port.Type
	tf["uid"] = port.UID
	return tf
}
