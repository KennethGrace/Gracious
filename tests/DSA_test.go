package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Art-of-the-Living/gracious/base"
	"github.com/Art-of-the-Living/gracious/components"
	"github.com/Art-of-the-Living/gracious/tests/util"
	"io/ioutil"
	"os"
	"testing"
)

func TestDSABasic(t *testing.T) {
	colorDataFile, err := os.Open("data/colorA.json")
	if err != nil {
		fmt.Println(err)
	}
	defer colorDataFile.Close()
	wordDataFile, err := os.Open("data/wordA.json")
	if err != nil {
		fmt.Println(err)
	}
	defer wordDataFile.Close()
	var colorSignalData util.JsonSignalData
	var wordSignalData util.JsonSignalData
	bytes, _ := ioutil.ReadAll(colorDataFile)
	json.Unmarshal(bytes, &colorSignalData)
	bytes, _ = ioutil.ReadAll(wordDataFile)
	json.Unmarshal(bytes, &wordSignalData)
	sensorA := components.NewSensor("Color")
	sensorB := components.NewSensor("Word")
	displayCount := 16
	aSteps := 0
	aPos := 0
	sensorA.SetProcessor(func() base.DistributedSignal {
		tmp := colorSignalData.Signals[aPos].ToDistributedSignal()
		aSteps++
		if aSteps >= displayCount {
			aPos++
			aSteps = 0
			if aPos >= len(colorSignalData.Signals) {
				aPos = 0
			}
		}
		return tmp
	})
	bSteps := 0
	bPos := 0
	sensorB.SetProcessor(func() base.DistributedSignal {
		tmp := wordSignalData.Signals[bPos].ToDistributedSignal()
		bSteps++
		if bSteps >= displayCount {
			bPos++
			bSteps = 0
			if bPos >= len(wordSignalData.Signals) {
				bPos = 0
			}
		}
		return tmp
	})
	evaluatorA := components.NewEvaluator("Color")
	evaluatorB := components.NewEvaluator("Word")

	associatorA := components.NewAssociator("Color")
	associatorB := components.NewAssociator("Word")

	PsiASet := make(map[string]base.DistributedSignal)
	PsiBSet := make(map[string]base.DistributedSignal)
	XiASet := make(map[string]base.DistributedSignal)
	XiBSet := make(map[string]base.DistributedSignal)
	count := 128
	for i := 0; i < count; i++ {
		fmt.Println("THE CURRENT OBSERVATION IS", i, "OF", count)
		NuA := sensorA.Evoke()
		NuB := sensorB.Evoke()
		fmt.Println(NuA.Represent())
		fmt.Println(NuB.Represent())
		evaluatorA.Main = NuA
		evaluatorA.Associates = XiASet
		evaluatorB.Main = NuB
		evaluatorB.Associates = XiBSet
		PsiBSet["A"] = evaluatorA.Evoke()
		PsiASet["B"] = evaluatorB.Evoke()
		associatorA.Main = PsiBSet["A"]
		associatorB.Main = PsiASet["B"]
		associatorA.Associates = PsiASet
		associatorB.Associates = PsiBSet
		XiASet["A"] = associatorA.Evoke()
		XiBSet["B"] = associatorB.Evoke()
		tmp := XiASet["A"]
		fmt.Println(tmp.Represent())
		tmp = XiBSet["B"]
		fmt.Println(tmp.Represent())
	}
}