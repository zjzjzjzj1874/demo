package _interface

type CPUDecode struct{}

func (d *CPUDecode) Open(uri string) error {
	return nil
}

func (d *CPUDecode) Close() error {
	return nil
}
