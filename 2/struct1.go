package main
import "fmt"

type person struct{
	name string
	age int
}

func Older(p1,p2 person)(person,int){
	if p1.age>p2.age{
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

func main(){
	var tom person
	tom.name,tom.age="Tom",18

//	bob:=person{"Bob",43}

	paul:=person{age:25,name:"Paul"}

//	tb_Older,tb_diff:=Older(tom,bob)
	tp_Older,tp_diff:=Older(tom,paul)
//	bp_Older,hp_diff:=Older(bob,paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",tom.name,paul.name,tp_Older.name,tp_diff)
}

