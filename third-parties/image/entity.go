package images

type IBusiness interface {
	IsNSFW(filepath string) (bool, error)
}
