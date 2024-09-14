package enum

type Gender int

const (
	Male Gender = iota
	Female
)

func (d Gender) ConvertGender() string {
	switch d {
	case Male:
		return "Nam"
	case Female:
		return "Nữ"
	default:
		return "Bê đê"
	}
}
