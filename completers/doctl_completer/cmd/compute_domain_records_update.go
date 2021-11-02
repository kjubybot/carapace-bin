package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_domain_records_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a DNS record",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_domain_records_updateCmd).Standalone()
	compute_domain_records_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Type`, `Name`, `Data`, `Priority`, `Port`, `TTL`, `Weight`")
	compute_domain_records_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_domain_records_updateCmd.Flags().String("record-data", "", "Record data; varies depending on record type")
	compute_domain_records_updateCmd.Flags().Int("record-flags", 0, "An unsigned integer between 0-255 used for CAA records")
	compute_domain_records_updateCmd.Flags().Int("record-id", 0, "Record ID")
	compute_domain_records_updateCmd.Flags().String("record-name", "", "The host name, alias, or service being defined by the record")
	compute_domain_records_updateCmd.Flags().Int("record-port", 0, "The port value for an SRV record")
	compute_domain_records_updateCmd.Flags().Int("record-priority", 0, "Record priority")
	compute_domain_records_updateCmd.Flags().String("record-tag", "", "The parameter tag for CAA records. Valid values are `issue`, `issuewild`, or `iodef`")
	compute_domain_records_updateCmd.Flags().Int("record-ttl", 0, "The record's Time To Live value, in seconds")
	compute_domain_records_updateCmd.Flags().String("record-type", "", "The type of DNS record")
	compute_domain_records_updateCmd.Flags().Int("record-weight", 0, "The weight value for an SRV record")
	compute_domain_recordsCmd.AddCommand(compute_domain_records_updateCmd)

	carapace.Gen(compute_domain_records_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Type", "Name", "Data", "Priority", "Port", "TTL", "Weight").Invoke(c).Filter(c.Parts).ToA()
		}),
		"record-tag": carapace.ActionValues("issue", "issuewild", "iodef"),
	})
}
