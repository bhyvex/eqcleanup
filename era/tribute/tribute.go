//This script will disable tribute masters by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package tribute

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "tribute"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}

	//tribute master
	ids, err := spawngroup.GetSpawnGroupIdsByClass(db, 63)
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}

	//tribute grand master
	mids, err := spawngroup.GetSpawnGroupIdsByClass(db, 64)
	for _, id := range mids {
		ids = append(ids, id)
	}

	if len(ids) < 1 {
		fmt.Println("No ", focus, "were found to delete")
		return
	}

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")
	return
}
