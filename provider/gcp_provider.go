package provider

type gcp struct {
}

func (gcp *gcp) Type() CloudType {
	return GCP
}

func newGCP() (*gcp, error) {
	return &gcp{}, nil
}
