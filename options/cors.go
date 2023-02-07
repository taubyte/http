package options

import "regexp"

func AllowedMethods(methods []string) Option {
	return func(s Configurable) error {
		return s.SetOption(OptionAllowedMethods{Methods: methods})
	}
}

func AllowedOrigins(regex bool, origins []string) Option {
	return func(s Configurable) error {
		if regex == false {
			return s.SetOption(OptionAllowedOrigins{
				Func: func(origin string) bool {
					for _, o := range origins {
						if origin == o {
							return true
						}
					}
					return false
				},
			})
		} else {
			_origins := make([]*regexp.Regexp, len(origins))
			for i, o := range origins {
				_origins[i] = regexp.MustCompile(o)
			}
			return s.SetOption(OptionAllowedOrigins{
				Func: func(origin string) bool {
					_origin := []byte(origin)
					for _, o := range _origins {
						if o.Match(_origin) == true {
							return true
						}
					}
					return false
				},
			})
		}
	}
}
