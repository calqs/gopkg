package router

type Options struct {
	BaseURL string
}

type OptionFunc = func(*Options)

func WithBaseURL(path string) func(*Options) {
	return func(o *Options) {
		o.BaseURL = CleanPath(path)
	}
}
