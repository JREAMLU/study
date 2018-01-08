package server

// Magic interface
type Magic interface {
	Options() Options
	Init(...Option) error
	Start(string) (string, error)
	Close(string) (string, error)
}

var (
	// DefaultMagic default
	DefaultMagic = NewMagic()
)

// NewMagic new
func NewMagic(opts ...Option) Magic {
	options := Options{
		Boss: "jream",
	}
	for _, o := range opts {
		o(&options)

	}
	return nil
}

// Init init
func Init(opts ...Option) error {
	return DefaultMagic.Init(opts...)
}

// Start start
func Start(name string) (string, error) {
	return DefaultMagic.Start(name)
}

// Close close
func Close(name string) (string, error) {
	return DefaultMagic.Close(name)
}
