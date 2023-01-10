package builder

import (
	"fmt"
	"testing"
)

func TestManufacture(t *testing.T) {
	manufacture := Manufacture{}

	homePcBuilder := &HomeEditionBuilder{}
	manufacture.SetBuilder(homePcBuilder)
	homePcBuilder.SetCPU("i3")
	homePcBuilder.SetRAM(8)
	homePcBuilder.SetOS("windows")
	homePcBuilder.SetHDD(500)
	homePcBuilder.SetGPU("Intel Graphic Card")
	homePc := homePcBuilder.GetPC()

	if homePc.cpu != "i5" {
		t.Error("Home edition pc would be i5 but found $d\n", homePc.cpu)
	}

	if homePc.gpu != "Intel graphic card" {
		t.Error("Home edition pc would be 'Intel graphic card but found $d\n", homePc.gpu)
	}

	gamePcBuilder := &GameEditionPCBuilder{}
	manufacture.SetBuilder(gamePcBuilder)
	gamePcBuilder.SetRAM(16)
	gamePcBuilder.SetOS("windows")
	gamePcBuilder.SetHDD(500)
	gamePcBuilder.SetGPU("AMD Radeon Graphic Card")
	gamePc := gamePcBuilder.GetPC()

	fmt.Printf(gamePc.os)
	// if gamePc.cpu != nil {
	// 	t.Error("Home edition pc would be i5 but found $d\n", gamePc.cpu)
	// }

}
