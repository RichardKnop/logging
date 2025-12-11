package logging

type noOp struct {
}

// NewNoOp creates a no-op logger that implements LoggerInterface but does nothing.
// Useful for disabling logging. Also used with Config MinLogLevel on disabled levels.
func NewNoOp() LoggerInterface {
	return &noOp{}
}

func (n noOp) Print(i ...interface{}) {}

func (n noOp) Printf(s string, i ...interface{}) {}

func (n noOp) Println(i ...interface{}) {}

func (n noOp) Fatal(i ...interface{}) {}

func (n noOp) Fatalf(s string, i ...interface{}) {}

func (n noOp) Fatalln(i ...interface{}) {}

func (n noOp) Panic(i ...interface{}) {}

func (n noOp) Panicf(s string, i ...interface{}) {}

func (n noOp) Panicln(i ...interface{}) {}
