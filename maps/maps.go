package maps

import "fmt"
import . "github.com/ughoavgfhw/libkq/common"

func MetadataForMap(m Map) *MapMetadata {
	switch m {
	case DayMap:
		return Day
	case NightMap:
		return Night
	case DuskMap:
		return Dusk
	case WarriorBonusMap:
		return WarriorBonus
	case SnailBonusMap:
		return SnailBonus
	}
	panic(fmt.Sprint("No metadata for map ", int(m)))
}
