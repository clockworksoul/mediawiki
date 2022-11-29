package main

type Module struct {
	Name        string
	Description string
	Parameters  []*Param
	Flags       []string
}

type ParamType string

const (
	String         ParamType = "string"
	ListOfStrings  ParamType = "...string"
	Boolean        ParamType = "boolean"
	Expiry         ParamType = "expiry"
	Integer        ParamType = "integer"
	ListOfIntegers ParamType = "...int"
	Timestamp      ParamType = "timestamp"
)

type Param struct {
	Name, Description    string
	Type                 ParamType
	Deprecated, Required bool
}
