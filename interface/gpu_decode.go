package _interface

type GPUDecode struct{}

func (d *GPUDecode) Open(uri string) error {
	return nil
}

func (d *GPUDecode) Close() error {
	return nil
}
