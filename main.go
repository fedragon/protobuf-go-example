package main

import (
	"fmt"
	semverv1 "github.com/fedragon/protobuf-go-example/semver/v1"
	typesv1 "github.com/fedragon/protobuf-go-example/types/v1"
)

func main() {
	version := semverv1.Version{
		Major: &typesv1.Change{Value: "1"},
		Minor: &typesv1.Change{Value: "0"},
		Patch: &typesv1.Change{Value: "0"},
	}

	fmt.Println(version.String())
}
