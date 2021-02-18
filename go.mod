module autogold

go 1.15

require (
	github.com/hexops/autogold v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/vektah/gqlparser/v2 v2.1.0
)

replace github.com/hexops/autogold => ../autogold
replace github.com/hexops/valast => ../valast