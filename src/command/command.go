package command

import (
	"encoding/json"
	"fmt"
)

type Cmds []Cmd

type Cmd struct {
	Order string
	Param []string
}

func (self *Cmds) Add(order string, param []string) {
	*self = append(*self, Cmd{Order: order, Param: param})
}

func (self *Cmds) Marshal() ([]byte, error) {
	return json.Marshal(*self)
}

func (self *Cmds) Unmarshal(data []byte) error {
	return json.Unmarshal(data, self)
}

func (self *Cmds) Print() {
	for _, cmd := range *self {
		fmt.Printf("Order:%v\n", cmd.Order)
		for _, param := range cmd.Param {
			fmt.Printf("----Param:%v\n", param)
		}
	}
}

func Test() {
	var testData = []byte(`[{"Order":"put", "Param":["param1", "Param2"]},{"Order":"get", "Param":["param3", "Param4", "Param5"]}]`)
	var cmds Cmds
	cmds.Unmarshal(testData)
	cmds.Print()

	cmds.Add("post", []string{"ssssss", "ddddddd", "tttttt", "yyyyyyy"})
	cmds.Add("check", []string{})
	b, err := cmds.Marshal()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(b)

	var cmds2 Cmds
	cmds2.Unmarshal(b)
	cmds2.Print()
}
