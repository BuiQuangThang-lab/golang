package enum

type Role int

const (
	Admin   = 1
	Staff   = 2
	Leader  = 3
	Manager = 4
)

func (d Role) ConvertRole() string {
	switch d {
	case Admin:
		return "Quản trị"
	case Staff:
		return "Nhân viên"
	case Leader:
		return "Tổ trưởng"
	case Manager:
		return "Trưởng phòng"
	default:
		return "Nhân viên"
	}
}
