package server

// Options 设置
type Options struct {
	Name  string
	Level string
	Boss  string
	Magic *Magic
}

// Option opt
type Option func(*Options)
