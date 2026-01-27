package tests

import (
	"net"
	"regexp"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinimalEC2(t *testing.T) {
	t.Parallel()
	opts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})
	defer terraform.Destroy(t, opts)
	_, err := terraform.InitAndApplyE(t, opts)
	require.NoError(t, err)

	t.Run("outputs", func(t *testing.T) {
		tests := []struct {
			name       string
			outputName string
			assertion  func(t *testing.T, val string)
		}{
			{
				name:       "instance id format",
				outputName: "instance_id",
				assertion: func(t *testing.T, val string) {
					pattern := regexp.MustCompile(`^i-[a-z0-9]{17}$`)
					assert.Regexp(t, pattern, val)
				},
			},
			{
				name:       "instance arn prefix",
				outputName: "instance_arn",
				assertion: func(t *testing.T, val string) {
					const prefix = "arn:aws:ec2:"
					assert.Contains(t, val, prefix)
				},
			},
			{
				name:       "instance public ip validity",
				outputName: "instance_public_ip",
				assertion: func(t *testing.T, val string) {
					ip := net.ParseIP(val)
					if assert.NotNil(t, ip) {
						assert.True(t, ip.IsGlobalUnicast())
						assert.NotNil(t, ip.To4())
					}
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := terraform.Output(t, opts, tt.outputName)
				tt.assertion(t, got)
			})
		}
	})
}
