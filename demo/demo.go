package demo

// student 学生
type student struct {
	Name string
	age  float64
}

// Getstudent 获取
func Getstudent(n string, a float64) *student {
	return &student{
		Name: n,
		age:  a,
	}
}

func (s student) Getage() float64 {
	return s.age
}
