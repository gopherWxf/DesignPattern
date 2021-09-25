package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	laowang:=NewPerson("Wang",NewCommand(nil,nil))
	laozhang:=NewPerson("Zhang",NewCommand(&laowang,laowang.Listen))
	laofeng:=NewPerson("Feng",NewCommand(&laozhang,laozhang.Buy))
	laoding:=NewPerson("Ding",NewCommand(&laofeng,laofeng.Cook))
	laoli:=NewPerson("Li",NewCommand(&laoding,laoding.Wash))

	laoli.Talk()
}
