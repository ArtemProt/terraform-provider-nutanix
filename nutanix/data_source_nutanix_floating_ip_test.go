package nutanix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNutanixFloatingIPDataSource_basic(t *testing.T) {
	r := randIntBetween(31, 40)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFloatingIPDataSourceConfig(r),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nutanix_floating_ip.test", "spec.0.resources.0.external_subnet_reference.kind", "subnet"),
					resource.TestCheckResourceAttr("data.nutanix_floating_ip.test", "spec.0.resources.0.vpc_reference.kind", "vpc"),
				),
			},
		},
	})
}

func testAccFloatingIPDataSourceConfig(r int) string {
	return fmt.Sprintf(`

	data "nutanix_clusters" "clusters" {}
	
	locals {
		cluster1 = [
		for cluster in data.nutanix_clusters.clusters.entities :
		cluster.metadata.uuid if cluster.service_list[0] != "PRISM_CENTRAL"
		][0]
	}
	
	resource "nutanix_subnet" "sub-test" {
		cluster_uuid = local.cluster1
		name        = "acctest-managed-%[1]d"
		description = "Description of my unit test VLAN"
		vlan_id     = %[1]d
		subnet_type = "VLAN"
		subnet_ip          = "10.250.140.0"
	  default_gateway_ip = "10.250.140.1"
	  prefix_length = 24
	  is_external = true
	  ip_config_pool_list_ranges = ["10.250.140.10 10.250.140.20"]
	}

	resource "nutanix_vpc" "test-vpc" {
		name = "acctest-vpc-%[1]d"
	  
	  
		external_subnet_reference_uuid = [
		  resource.nutanix_subnet.sub-test.id
		]
	  
		common_domain_name_server_ip_list{
				ip = "8.8.8.9"
		}
	  
		externally_routable_prefix_list{
		  ip=  "172.31.0.0"
		  prefix_length= 16
		}
	  }

	  resource "nutanix_floating_ip" "test-fip" {
		external_subnet_reference_uuid = resource.nutanix_subnet.sub-test.id
		vpc_reference_uuid= resource.nutanix_vpc.test-vpc.id
		private_ip = "10.3.3.6"
	}

	data "nutanix_floating_ip" "test"{
		floating_ip_uuid = resource.nutanix_floating_ip.test-fip.id
	}
	`, r)
}
