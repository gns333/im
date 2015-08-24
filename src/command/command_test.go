package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	// testing Unmarshal
	var testData = []byte(`[{"Order":"put", "Param":["param1", "param2"]},
							{"Order":"get", "Param":["param3", "param4", "param5"]}]`)
	var cmds1 Cmds
	cmds1.Unmarshal(testData)

	if len(cmds1) != 2 {
		t.Errorf("Command Unmarshal Error! (%d) != 2", len(cmds1))
	}

	if cmds1[1].Order != "get" {
		t.Errorf("Command Unmarshal Error! (%s) != get", cmds1[1].Order)
	}

	if len(cmds1[1].Param) != 3 {
		t.Errorf("Command Unmarshal Error! (%d) != 3", len(cmds1[1].Param))
	}

	if cmds1[1].Param[2] != "param5" {
		t.Errorf("Command Unmarshal Error! (%s) != param5", cmds1[1].Param[2])
	}

	// testing Add and Marshal
	var cmds2 Cmds
	cmds2.Add("post", []string{"param6", "param7", "param8", "param9"})
	cmds2.Add("check", []string{})
	cmds2.Add("take", []string{"param10", "param11"})

	b, err := cmds2.Marshal()
	if err != nil {
		t.Error("Command Marshal Error!")
	}

	var cmds3 Cmds
	cmds3.Unmarshal(b)

	if len(cmds3) != 3 {
		t.Errorf("Command Marshal Error! (%d) != 2", len(cmds3))
	}

	if cmds3[2].Order != "take" {
		t.Errorf("Command Marshal Error! (%s) != take", cmds3[2].Order)
	}

	if len(cmds3[2].Param) != 2 {
		t.Errorf("Command Marshal Error! (%d) != 2", len(cmds3[2].Param))
	}

	if cmds3[2].Param[0] != "param10" {
		t.Errorf("Command Marshal Error! (%s) != 2", cmds3[2].Param[0])
	}
}
