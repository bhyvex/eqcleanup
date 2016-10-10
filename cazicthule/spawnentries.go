package cazicthule

import (
	"github.com/xackery/goeq/spawn"
	//"database/sql"
)

var spawnentries []spawn.SpawnEntry = []spawn.SpawnEntry{
	{Spawngroupid: 48001, Npcid: 48001, Chance: 50},
	{Spawngroupid: 48001, Npcid: 48300, Chance: 50},
	{Chance: 100, Spawngroupid: 48038, Npcid: 48038},
	{Spawngroupid: 48040, Npcid: 48040, Chance: 100},
	{Spawngroupid: 48041, Npcid: 48041, Chance: 33},
	{Spawngroupid: 48041, Npcid: 48302, Chance: 33},
	{Spawngroupid: 48041, Npcid: 48303, Chance: 34},
	{Spawngroupid: 48049, Npcid: 48049, Chance: 50},
	{Spawngroupid: 48049, Npcid: 48305, Chance: 50},
	{Spawngroupid: 48054, Npcid: 48054, Chance: 100},
	{Chance: 100, Spawngroupid: 48057, Npcid: 48057},
	{Spawngroupid: 48058, Npcid: 48058, Chance: 100},
	{Npcid: 48059, Chance: 50, Spawngroupid: 48059},
	{Spawngroupid: 48059, Npcid: 48307, Chance: 50},
	{Spawngroupid: 48060, Npcid: 48060, Chance: 100},
	{Spawngroupid: 48062, Npcid: 48062, Chance: 100},
	{Spawngroupid: 48064, Npcid: 48064, Chance: 100},
	{Spawngroupid: 48065, Npcid: 48065, Chance: 100},
	{Chance: 100, Spawngroupid: 48066, Npcid: 48066},
	{Spawngroupid: 48067, Npcid: 48067, Chance: 100},
	{Npcid: 48069, Chance: 100, Spawngroupid: 48069},
	{Spawngroupid: 48070, Npcid: 48070, Chance: 50},
	{Spawngroupid: 48070, Npcid: 48309, Chance: 50},
	{Chance: 50, Spawngroupid: 48071, Npcid: 48071},
	{Spawngroupid: 48071, Npcid: 48311, Chance: 50},
	{Spawngroupid: 48073, Npcid: 48073, Chance: 100},
	{Spawngroupid: 48074, Npcid: 48074, Chance: 100},
	{Spawngroupid: 48077, Npcid: 48077, Chance: 100},
	{Npcid: 48079, Chance: 100, Spawngroupid: 48079},
	{Spawngroupid: 48081, Npcid: 48081, Chance: 100},
	{Spawngroupid: 48083, Npcid: 48083, Chance: 100},
	{Spawngroupid: 48084, Npcid: 48084, Chance: 100},
	{Spawngroupid: 48087, Npcid: 48087, Chance: 100},
	{Spawngroupid: 48089, Npcid: 48089, Chance: 100},
	{Spawngroupid: 48090, Npcid: 48090, Chance: 100},
	{Npcid: 48093, Chance: 100, Spawngroupid: 48093},
	{Chance: 100, Spawngroupid: 48094, Npcid: 48094},
	{Spawngroupid: 48095, Npcid: 48095, Chance: 100},
	{Npcid: 48097, Chance: 100, Spawngroupid: 48097},
	{Chance: 100, Spawngroupid: 48098, Npcid: 48098},
	{Spawngroupid: 48099, Npcid: 48099, Chance: 100},
	{Npcid: 48100, Chance: 100, Spawngroupid: 48100},
	{Spawngroupid: 48102, Npcid: 48102, Chance: 50},
	{Spawngroupid: 48102, Npcid: 48313, Chance: 50},
	{Spawngroupid: 48103, Npcid: 48103, Chance: 50},
	{Spawngroupid: 48103, Npcid: 48315, Chance: 50},
	{Spawngroupid: 48104, Npcid: 48104, Chance: 100},
	{Spawngroupid: 48106, Npcid: 48106, Chance: 100},
	{Spawngroupid: 48107, Npcid: 48107, Chance: 100},
	{Chance: 100, Spawngroupid: 48109, Npcid: 48109},
	{Spawngroupid: 48110, Npcid: 48110, Chance: 100},
	{Npcid: 48113, Chance: 50, Spawngroupid: 48113},
	{Spawngroupid: 48113, Npcid: 48317, Chance: 50},
	{Spawngroupid: 48116, Npcid: 48116, Chance: 100},
	{Spawngroupid: 48117, Npcid: 48117, Chance: 100},
	{Npcid: 48120, Chance: 100, Spawngroupid: 48120},
	{Spawngroupid: 48123, Npcid: 48123, Chance: 100},
	{Spawngroupid: 48126, Npcid: 48126, Chance: 100},
	{Spawngroupid: 48129, Npcid: 48129, Chance: 100},
	{Spawngroupid: 48131, Npcid: 48131, Chance: 100},
	{Chance: 100, Spawngroupid: 48132, Npcid: 48132},
	{Npcid: 48133, Chance: 100, Spawngroupid: 48133},
	{Chance: 100, Spawngroupid: 48135, Npcid: 48135},
	{Spawngroupid: 48138, Npcid: 48138, Chance: 100},
	{Chance: 100, Spawngroupid: 48140, Npcid: 48140},
	{Spawngroupid: 48142, Npcid: 48142, Chance: 100},
	{Spawngroupid: 48144, Npcid: 48144, Chance: 100},
	{Spawngroupid: 48146, Npcid: 48146, Chance: 100},
	{Chance: 100, Spawngroupid: 48154, Npcid: 48154},
	{Spawngroupid: 48157, Npcid: 48157, Chance: 100},
	{Spawngroupid: 48163, Npcid: 48163, Chance: 100},
	{Spawngroupid: 48164, Npcid: 48164, Chance: 100},
	{Spawngroupid: 48173, Npcid: 48173, Chance: 100},
	{Spawngroupid: 48177, Npcid: 48177, Chance: 100},
	{Spawngroupid: 48182, Npcid: 48182, Chance: 100},
	{Spawngroupid: 48187, Npcid: 48187, Chance: 100},
	{Chance: 100, Spawngroupid: 48188, Npcid: 48188},
	{Spawngroupid: 48194, Npcid: 48194, Chance: 100},
	{Spawngroupid: 48196, Npcid: 48196, Chance: 100},
	{Spawngroupid: 48197, Npcid: 48197, Chance: 100},
	{Spawngroupid: 48199, Npcid: 48199, Chance: 100},
	{Spawngroupid: 48201, Npcid: 48201, Chance: 50},
	{Spawngroupid: 48201, Npcid: 48319, Chance: 50},
	{Chance: 100, Spawngroupid: 48208, Npcid: 48208},
	{Spawngroupid: 48274, Npcid: 48274, Chance: 100},
	{Spawngroupid: 48298, Npcid: 48298, Chance: 100},
}
