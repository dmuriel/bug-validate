package tasks

import (
	"bugvalidate/app"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gobuffalo/grift/grift"
)

var _ = grift.Desc("routes", "Print out all defined routes")
var _ = grift.Add("routes", func(c *grift.Context) error {
	routes := app.New().Routes()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "METHOD\t PATH\t ALIASES\t NAME\t HANDLER")
	fmt.Fprintln(w, "------\t ----\t -------\t ----\t -------")

	for _, r := range routes {
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\t %s\n", r.Method, r.Path, strings.Join(r.Aliases, " "), r.PathName, r.HandlerName)
	}

	w.Flush()

	return nil
})
