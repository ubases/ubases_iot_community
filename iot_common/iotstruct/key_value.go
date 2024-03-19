package iotstruct

type KeyIntValueString struct {
	Key   int64
	Value string
}

type KeyValue struct {
	Key   string
	Value string
}

type CountResult struct {
	Key   string
	Count int64
}

type DropdownItem struct {
	Code  string
	Name  string
	Count int64
}
