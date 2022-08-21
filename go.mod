module github.com/davidalpert/go-git-mob

go 1.16

//replace github.com/davidalpert/go-printers => ../go-printers

require (
	github.com/apex/log v1.9.0
	github.com/davidalpert/go-printers v0.3.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.20.0
	github.com/spf13/cobra v1.5.0
)

require github.com/rogpeppe/go-internal v1.8.1 // indirect

require (
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.0
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
