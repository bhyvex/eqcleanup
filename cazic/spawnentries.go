package cazic

import (
	"github.com/xackery/goeq/spawn"
	//"database/sql"
)

var spawnentries []spawn.SpawnEntry = []spawn.SpawnEntry{
	spawn.SpawnEntry{Npcid: 48001, Chance: 50, Spawngroupid: 48001},
	spawn.SpawnEntry{Spawngroupid: 48001, Npcid: 48300, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48038, Npcid: 48038, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48040, Npcid: 48040, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48041, Npcid: 48041, Chance: 33},
	spawn.SpawnEntry{Spawngroupid: 48041, Npcid: 48302, Chance: 33},
	spawn.SpawnEntry{Spawngroupid: 48041, Npcid: 48303, Chance: 34},
	spawn.SpawnEntry{Spawngroupid: 48049, Npcid: 48049, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48049, Npcid: 48305, Chance: 50},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48054, Npcid: 48054},
	spawn.SpawnEntry{Spawngroupid: 48057, Npcid: 48057, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48058, Npcid: 48058, Chance: 100},
	spawn.SpawnEntry{Npcid: 48059, Chance: 50, Spawngroupid: 48059},
	spawn.SpawnEntry{Spawngroupid: 48059, Npcid: 48307, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48060, Npcid: 48060, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48062, Npcid: 48062, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48064, Npcid: 48064},
	spawn.SpawnEntry{Spawngroupid: 48065, Npcid: 48065, Chance: 100},
	spawn.SpawnEntry{Npcid: 48066, Chance: 100, Spawngroupid: 48066},
	spawn.SpawnEntry{Spawngroupid: 48067, Npcid: 48067, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48069, Npcid: 48069, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48070, Npcid: 48070, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48070, Npcid: 48309, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48071, Npcid: 48071, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48071, Npcid: 48311, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48073, Npcid: 48073, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48074, Npcid: 48074, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48077, Npcid: 48077, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48079, Npcid: 48079, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48081, Npcid: 48081, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48083, Npcid: 48083, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48084, Npcid: 48084, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48087, Npcid: 48087, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48089, Npcid: 48089, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48090, Npcid: 48090, Chance: 100},
	spawn.SpawnEntry{Npcid: 48093, Chance: 100, Spawngroupid: 48093},
	spawn.SpawnEntry{Spawngroupid: 48094, Npcid: 48094, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48095, Npcid: 48095},
	spawn.SpawnEntry{Spawngroupid: 48097, Npcid: 48097, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48098, Npcid: 48098, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48099, Npcid: 48099, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48100, Npcid: 48100, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48102, Npcid: 48102, Chance: 50},
	spawn.SpawnEntry{Chance: 50, Spawngroupid: 48102, Npcid: 48313},
	spawn.SpawnEntry{Spawngroupid: 48103, Npcid: 48103, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48103, Npcid: 48315, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48104, Npcid: 48104, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48106, Npcid: 48106, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48107, Npcid: 48107},
	spawn.SpawnEntry{Spawngroupid: 48109, Npcid: 48109, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48110, Npcid: 48110, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48113, Npcid: 48113, Chance: 50},
	spawn.SpawnEntry{Spawngroupid: 48113, Npcid: 48317, Chance: 50},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48116, Npcid: 48116},
	spawn.SpawnEntry{Npcid: 48117, Chance: 100, Spawngroupid: 48117},
	spawn.SpawnEntry{Spawngroupid: 48120, Npcid: 48120, Chance: 100},
	spawn.SpawnEntry{Npcid: 48123, Chance: 100, Spawngroupid: 48123},
	spawn.SpawnEntry{Spawngroupid: 48126, Npcid: 48126, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48129, Npcid: 48129, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48131, Npcid: 48131, Chance: 100},
	spawn.SpawnEntry{Npcid: 48132, Chance: 100, Spawngroupid: 48132},
	spawn.SpawnEntry{Spawngroupid: 48133, Npcid: 48133, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48135, Npcid: 48135, Chance: 100},
	spawn.SpawnEntry{Npcid: 48138, Chance: 100, Spawngroupid: 48138},
	spawn.SpawnEntry{Spawngroupid: 48140, Npcid: 48140, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48142, Npcid: 48142},
	spawn.SpawnEntry{Spawngroupid: 48144, Npcid: 48144, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48146, Npcid: 48146, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48154, Npcid: 48154, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48157, Npcid: 48157, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48163, Npcid: 48163, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48164, Npcid: 48164, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48173, Npcid: 48173, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48177, Npcid: 48177, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48182, Npcid: 48182},
	spawn.SpawnEntry{Spawngroupid: 48187, Npcid: 48187, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48188, Npcid: 48188, Chance: 100},
	spawn.SpawnEntry{Npcid: 48194, Chance: 100, Spawngroupid: 48194},
	spawn.SpawnEntry{Spawngroupid: 48196, Npcid: 48196, Chance: 100},
	spawn.SpawnEntry{Chance: 100, Spawngroupid: 48197, Npcid: 48197},
	spawn.SpawnEntry{Spawngroupid: 48199, Npcid: 48199, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48201, Npcid: 48201, Chance: 50},
	spawn.SpawnEntry{Chance: 50, Spawngroupid: 48201, Npcid: 48319},
	spawn.SpawnEntry{Spawngroupid: 48208, Npcid: 48208, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48274, Npcid: 48274, Chance: 100},
	spawn.SpawnEntry{Spawngroupid: 48298, Npcid: 48298, Chance: 100},
}
