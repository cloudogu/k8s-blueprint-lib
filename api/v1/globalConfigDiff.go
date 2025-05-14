package v1

type GlobalConfigDiff []GlobalConfigEntryDiff

type GlobalConfigValueState ConfigValueState
type GlobalConfigEntryDiff struct {
	Key          string                 `json:"key"`
	Actual       GlobalConfigValueState `json:"actual"`
	Expected     GlobalConfigValueState `json:"expected"`
	NeededAction ConfigAction           `json:"neededAction"`
}
