package utils

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PrinterOptions contains options for printing
type PrinterOptions struct {
	OutputFormat        *string
	DefaultOutputFormat *string
	Caption             *string
	PopulateTable       *func(*tablewriter.Table)
	ItemsSelector       *func() interface{}
}

// NewPrinterOptions defines new printer options
func NewPrinterOptions() *PrinterOptions {
	defaultFormat := "json"

	configuredOutput := viper.GetString("output")
	if configuredOutput != "" {
		defaultFormat = configuredOutput
	}

	return (&PrinterOptions{}).WithDefaultOutput(defaultFormat)
}

// WithDefaultOutput sets a default output format if one is not provided through a flag value
func (o *PrinterOptions) WithDefaultOutput(output string) *PrinterOptions {
	o.OutputFormat = &output
	return o
}

// WithDefaultTableWriter sets a default table writer
func (o *PrinterOptions) WithDefaultTableWriter() *PrinterOptions {
	return o.WithDefaultOutput("text").WithTableWriter("n/a", func(t *tablewriter.Table) {})
}

// WithTableWriter decorates a PrinterOptions with table writer configuration
func (o *PrinterOptions) WithTableWriter(caption string, populateTable func(*tablewriter.Table)) *PrinterOptions {
	o.Caption = &caption
	o.PopulateTable = &populateTable
	return o
}

func (o *PrinterOptions) WithItemsSelector(selectItems func() interface{}) *PrinterOptions {
	o.ItemsSelector = &selectItems
	return o
}

// SupportedFormats returns the list of supported formats
func (o *PrinterOptions) SupportedFormats() []string {
	if o.PopulateTable != nil {
		return supportedListPrinterKeys
	}
	return supportedObjectPrinterKeys
}

// SupportedFormatCategories returns the list of supported formats
func (o *PrinterOptions) SupportedFormatCategories() []string {
	if o.PopulateTable != nil {
		return supportedListPrinterCategories
	}
	return supportedObjectPrinterCategories
}

// Validate asserts that the printer options are valid
func (o *PrinterOptions) Validate() error {
	if !StringInSlice(o.SupportedFormats(), *o.OutputFormat) {
		return fmt.Errorf("invalid output format: %s\nvalid format values are: %v", *o.OutputFormat, strings.Join(o.SupportedFormatCategories(), "|"))
	}
	return nil
}

// FormatCategory returns the dereferenced format category
func (o *PrinterOptions) FormatCategory() string {
	ExitIfErr(o.Validate())

	if o.PopulateTable != nil {
		return supportedListPrinterFormatMap[*o.OutputFormat]
	}
	return supportedObjectPrinterFormatMap[*o.OutputFormat]
}

// AddPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) AddPrinterFlags(c *cobra.Command) {
	if o.OutputFormat != nil {
		if o.PopulateTable != nil {
			o.addListPrinterFlags(c)
		} else {
			o.addObjectPrinterFlags(c)
		}
	}
}

// AddObjectPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addObjectPrinterFlags(c *cobra.Command) {
	if o.OutputFormat != nil {
		c.Flags().StringVarP(o.OutputFormat, "output", "o", *o.OutputFormat, fmt.Sprintf("output format. One of: %s.", strings.Join(supportedObjectPrinterCategories, "|")))
	}
}

// AddListPrinterFlags adds flags to a cobra.Command
func (o *PrinterOptions) addListPrinterFlags(c *cobra.Command) {
	if o.OutputFormat != nil {
		c.Flags().StringVarP(o.OutputFormat, "output", "o", *o.OutputFormat, fmt.Sprintf("output format. One of: %s.", strings.Join(supportedListPrinterCategories, "|")))
	}
}
