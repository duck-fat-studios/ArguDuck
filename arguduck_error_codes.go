package arguduck

type ArguDuckErrorString string

const (
	OK           ArguDuckErrorString = ""
	UNKNOWN_TYPE ArguDuckErrorString = "UNKNWON TYPE"
	SHORT_IN_USE ArguDuckErrorString = "SHORT ALREADY USED"
	ARG_IN_USE   ArguDuckErrorString = "ARG ALREADY USED"
)
