package main

type Document struct {
	
}

type Machine interface {
	Fax (d Document)
	Print (d Document)
	Scan (d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Fax(d Document)  {}
func (m MultiFunctionPrinter) Print(d Document)  {}
func (m MultiFunctionPrinter) Scan(d Document)  {}
