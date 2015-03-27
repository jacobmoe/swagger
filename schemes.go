package swagger

type Schemes []string

func NewSchemes() *Schemes {
	return &Schemes{}
}

func (schemes *Schemes) Add(scheme string) error {
	_, err := SCHEMEFromString(scheme)
	if err != nil {
		return err
	}
	// unique schemes
	for _, s := range *schemes {
		if s == scheme {
			return nil
		}
	}
	*schemes = append(*schemes, scheme)
	return nil
}
