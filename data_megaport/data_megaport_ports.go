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

	//portId := d.Get("product_id").(string)
	d.SetId("my-fake-id")
	ports, retrievalErr := port.GetPorts()

	if retrievalErr != nil {
		return retrievalErr
	}

	var converted []string

	for _, port := range ports {
		converted = append(converted, port.UID)
	}

	d.Set("ports", converted)

	return nil
}
