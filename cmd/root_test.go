package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestGetStringFlag(t *testing.T) {
	tests := []struct {
		name          string
		flagName      string
		setupCmd      func() *cobra.Command
		expectedValue string
		expectError   bool
	}{
		{
			name:     "valid flag",
			flagName: "test-flag",
			setupCmd: func() *cobra.Command {
				cmd := &cobra.Command{}
				cmd.Flags().String("test-flag", "test-value", "")
				return cmd
			},
			expectedValue: "test-value",
			expectError:   false,
		},
		{
			name:     "empty flag",
			flagName: "test-flag",
			setupCmd: func() *cobra.Command {
				cmd := &cobra.Command{}
				cmd.Flags().String("test-flag", "", "")
				return cmd
			},
			expectedValue: "",
			expectError:   false,
		},
		{
			name:     "missing flag",
			flagName: "non-existent",
			setupCmd: func() *cobra.Command {
				return &cobra.Command{}
			},
			expectedValue: "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.setupCmd()
			got, err := getStringFlag(tt.flagName, cmd)
			if (err != nil) != tt.expectError {
				t.Errorf("getStringFlag() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expectedValue {
				t.Errorf("getStringFlag() = %v, expected %v", got, tt.expectedValue)
			}
		})
	}
}
