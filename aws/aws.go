package aws

type Services struct {
	S3 *S3Client
}

func Init() *Services {
	return &Services{}
}

func (s *Services) InitS3() error {
	client, err := InitS3()
	if err != nil {
		return err
	}
	s.S3 = client
	return nil
}
