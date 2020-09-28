package tools

type InternalPairwise struct {
	ID            string `faker:"uuid_hyphenated"`
	OurDid        string
	TheirDid      string
	TheirEndpoint string `faker:"url"`
	TheirLabel    string `faker:"name"`
	InitiatedByUs bool
	CreatedMs     int64 `faker:"unix_time"`
	ApprovedMs    int64 `faker:"unix_time"`
}
