package options

func Parse(c Configurable, opts []Option) error {
	var err error
	for _, o := range opts {
		err = o(c) //c.SetOption(o)
		if err != nil {
			return err
		}
	}

	return nil
}
