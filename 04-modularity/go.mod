module modularity-demo

go 1.15

require (
	dummyModule v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.12.0
)

replace dummyModule => ../dummyModule
