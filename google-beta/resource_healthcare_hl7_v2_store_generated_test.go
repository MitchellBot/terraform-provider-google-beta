// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHealthcareHl7V2Store_healthcareHl7V2StoreBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckHealthcareHl7V2StoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareHl7V2Store_healthcareHl7V2StoreBasicExample(context),
			},
		},
	})
}

func testAccHealthcareHl7V2Store_healthcareHl7V2StoreBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_hl7_v2_store" "default" {
  name    = "tf-test-example-hl7-v2-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }

  provider = google-beta
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-hl7-v2-notifications%{random_suffix}"
  provider = google-beta
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
  provider = google-beta
}
`, context)
}

func TestAccHealthcareHl7V2Store_healthcareHl7V2StoreParserConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckHealthcareHl7V2StoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareHl7V2Store_healthcareHl7V2StoreParserConfigExample(context),
			},
		},
	})
}

func testAccHealthcareHl7V2Store_healthcareHl7V2StoreParserConfigExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_hl7_v2_store" "default" {
  provider = google-beta
  name    = "tf-test-example-hl7-v2-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id

  parser_config {
    allow_null_header  = false
    segment_terminator = "Jw=="
    schema = <<EOF
{
  "schemas": [{
    "messageSchemaConfigs": {
      "ADT_A01": {
        "name": "ADT_A01",
        "minOccurs": 1,
        "maxOccurs": 1,
        "members": [{
            "segment": {
              "type": "MSH",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "EVN",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "PID",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "ZPD",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "OBX"
            }
          },
          {
            "group": {
              "name": "PROCEDURE",
              "members": [{
                  "segment": {
                    "type": "PR1",
                    "minOccurs": 1,
                    "maxOccurs": 1
                  }
                },
                {
                  "segment": {
                    "type": "ROL"
                  }
                }
              ]
            }
          },
          {
            "segment": {
              "type": "PDA",
              "maxOccurs": 1
            }
          }
        ]
      }
    }
  }],
  "types": [{
    "type": [{
        "name": "ZPD",
        "primitive": "VARIES"
      }

    ]
  }],
  "ignoreMinOccurs": true
}
EOF
  }
}

resource "google_healthcare_dataset" "dataset" {
  provider = google-beta
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func testAccCheckHealthcareHl7V2StoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_hl7_v2_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/hl7V2Stores/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("HealthcareHl7V2Store still exists at %s", url)
			}
		}

		return nil
	}
}