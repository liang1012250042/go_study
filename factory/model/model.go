package model

// 定义一个结构体
type Student struct {
	Name string
	Score float64
}

// 定义一个结构体
type person struct {
	Name string
	score float64
}

// 因为person结构体首字母小写，因此是私有变量，只能在model包里面使用
// 我们使用工厂模式来解决

func NewPerson(name string, score float64) *person{
	return &person{
		Name: name,
		score: score,
	}
}

// 如果score字段首字母是小写，则在其他包不可以直接方法，我们可以提供一个方法
func (p *person) GetScore() float64{
	return p.score;
}