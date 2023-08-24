package provider

type aws struct {
}

func (aws *aws) Type() CloudType {
	return AWS
}

func newAWS() (*aws, error) {
	return &aws{}, nil
}
