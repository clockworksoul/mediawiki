package main

type Module struct {
	Name        string
	Description string
	Parameters  []*Param
	Flags       []string
}

type ParamType string

const (
	String  ParamType = "string"
	Boolean ParamType = "boolean"
	Expiry  ParamType = "expiry"
	Integer ParamType = "integer"
)

type Param struct {
	Name, Description    string
	Type                 ParamType
	Deprecated, Required bool
}
