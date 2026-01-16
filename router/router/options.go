package router

type Options struct {
	BaseURL string
}

type OptionFunc = func(*Options)

// @deprecated: use OptionWithBaseURL instead
func WithBaseURL(path string) func(*Options) {
	return func(o *Options) {
		o.BaseURL = CleanPath(path)
	}
}

func OptionWithBaseURL(path string) OptionFunc {
	return WithBaseURL(path)
}
