module github.com/wxio/acli

go 1.18

replace github.com/jpillora/opts => github.com/millergarym/opts v1.6.1

// replace github.com/jpillora/opts => ../../../millergarym/opts

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/iancoleman/strcase v0.2.0
	github.com/jpillora/md-tmpl v1.2.2
	github.com/jpillora/opts v1.2.0
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/posener/complete v1.2.2-0.20190308074557-af07aa5181b3 // indirect
)
