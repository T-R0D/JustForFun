package queue

// Queue represents an interface for a data structure with queue behavior.
type Queue interface {
	AppendRight(x interface{})
	
	Len() int

	PopLeft() (interface{}, error)

	String() string
}
