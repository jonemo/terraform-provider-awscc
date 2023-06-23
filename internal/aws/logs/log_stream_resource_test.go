// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsLogStream_working(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::LogStream", "awscc_logs_log_stream", "test")
	rName := td.RandomName()

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: testAccAWSLogsLogStreamWorkingConfig(&td, rName),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func testAccAWSLogsLogStreamWorkingConfig(td *acctest.TestData, rName string) string {
	return fmt.Sprintf(`
resource awscc_logs_log_group %[2]q {
  log_group_name = %[3]q
}

resource %[1]q %[2]q {
  log_group_name  = awscc_logs_log_group.%[2]s.log_group_name
  log_stream_name = %[3]q
}
`, td.TerraformResourceType, td.ResourceLabel, rName)
}
