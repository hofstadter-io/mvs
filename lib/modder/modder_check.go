package modder

import (
	"fmt"
)

func (mdr *Modder) CompareModToSum() error {
	var err error
	fmt.Println("==================")

	mod := mdr.module

	err = mod.PrintSelfDeps()
	if err != nil {
		return err
	}

	fmt.Println("==================")

	sf := mod.SumFile
	err = sf.Print()
	if err != nil {
		return err
	}

	fmt.Println("==================")
	return nil
}

func (mdr *Modder) CompareSumToMod() error {

	return nil
}

func (mdr *Modder) CompareSumToVendor() error {

	return nil
}
