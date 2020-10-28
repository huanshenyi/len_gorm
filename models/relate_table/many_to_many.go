package relate_table

type Article2 struct {
	Id int
	Title string
	Content string
	Desc string
	Tags []Tag `gorm:"many2many:Article2_Tags"`
}

type Tag struct {
	Id int
	Name string
	Desc string
}
