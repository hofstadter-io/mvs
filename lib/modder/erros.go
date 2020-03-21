package modder

import (
	"fmt"
)

func (mdr *Modder) CheckForErrors() error {
	if len(mdr.module.Errors) > 0 {
		return fmt.Errorf("Exiting due to errors during vendoring in %s.", mdr.module.Module)
	}

	for _, dep := range mdr.depsMap {
		if len(dep.Errors) > 0 {
			return fmt.Errorf("Exiting due to errors during vendoring in %s.", dep.Module)
		}
	}

	return nil
}

func (mdr *Modder) PrintErrors() error {
	var wasError error

	if len(mdr.module.Errors) > 0 {
		wasError = fmt.Errorf("Exiting due to errors during vendoring.")
		for _, err := range mdr.module.Errors {
			fmt.Println(err)
		}
	}

	for _, dep := range mdr.depsMap {
		if len(dep.Errors) > 0 {
			if wasError != nil {
				wasError = fmt.Errorf("Exiting due to errors during vendoring.")
			}
			for _, err := range dep.Errors {
				fmt.Println(err)
			}
		}
	}

	return wasError
}

