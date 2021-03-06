package wipe

import (
	"database/sql"
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqdb"
)

var focus = "wipe"

type Spawndata struct {
	Npcid        int
	Spawngroupid int
	Chance       int
	Mindelay     int
	Despawntimer int
}

func Clean(args ...string) (err error) {

	totalRemoved, err := DeleteCharacters()
	if err != nil {
		return
	}
	fmt.Println("Deleted", totalRemoved, "DB entries related to", focus, "successfully.")

	return
}

func DeleteCharacters() (totalRemoved int64, err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}
	type dataStruct struct {
		tableName string
		fieldName string
		idStart   int64
	}

	tables := []dataStruct{
		{tableName: "account", fieldName: "id", idStart: 0},
		{tableName: "account_flags", fieldName: "p_accid", idStart: 0},
		{tableName: "account_ip", fieldName: "accid", idStart: 0},
		{tableName: "account_rewards", fieldName: "account_id", idStart: 0},
		{tableName: "adventure_details", fieldName: "id", idStart: 0},
		{tableName: "adventure_members", fieldName: "id", idStart: 0},
		{tableName: "adventure_stats", fieldName: "player_id", idStart: 0},
		{tableName: "buyer", fieldName: "charid", idStart: 0},
		{tableName: "char_recipe_list", fieldName: "char_id", idStart: 0},
		{tableName: "character_activities", fieldName: "charid", idStart: 0},
		{tableName: "character_alt_currency", fieldName: "char_id", idStart: 0},
		{tableName: "character_alternate_abilities", fieldName: "id", idStart: 0},
		{tableName: "character_backup", fieldName: "id", idStart: 0},
		{tableName: "character_bandolier", fieldName: "id", idStart: 0},
		{tableName: "character_bind", fieldName: "id", idStart: 0},
		{tableName: "character_buffs", fieldName: "character_id", idStart: 0},
		{tableName: "character_corpse_items", fieldName: "corpse_id", idStart: 0},
		{tableName: "character_corpses", fieldName: "id", idStart: 0},
		{tableName: "character_currency", fieldName: "id", idStart: 0},
		//dataStruct{tableName: "character_custom", fieldName: "id", idStart: 0},
		{tableName: "character_data", fieldName: "id", idStart: 0},
		{tableName: "character_disciplines", fieldName: "id", idStart: 0},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "character_inspect_messages", fieldName: "id", idStart: 0},
		{tableName: "character_item_recast", fieldName: "id", idStart: 0},
		{tableName: "character_languages", fieldName: "id", idStart: 0},
		{tableName: "character_lookup", fieldName: "id", idStart: 0},
		{tableName: "character_material", fieldName: "id", idStart: 0},
		{tableName: "character_memmed_spells", fieldName: "id", idStart: 0},
		{tableName: "character_pet_buffs", fieldName: "char_id", idStart: 0},
		{tableName: "character_pet_info", fieldName: "char_id", idStart: 0},
		{tableName: "character_pet_inventory", fieldName: "char_id", idStart: 0},
		{tableName: "character_potionbelt", fieldName: "id", idStart: 0},
		{tableName: "character_skills", fieldName: "id", idStart: 0},
		{tableName: "character_spells", fieldName: "id", idStart: 0},
		{tableName: "character_tasks", fieldName: "charid", idStart: 0},
		{tableName: "character_tribute", fieldName: "id", idStart: 0},
		{tableName: "chatchannels", fieldName: "minstatus", idStart: -1},
		{tableName: "completed_tasks", fieldName: "charid", idStart: 0},
		{tableName: "discovered_items", fieldName: "item_id", idStart: 0},
		{tableName: "eventlog", fieldName: "id", idStart: 0},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "faction_values", fieldName: "char_id", idStart: 0},
		{tableName: "friends", fieldName: "charid", idStart: 0},
		{tableName: "geq_character_currency", fieldName: "id", idStart: 0},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "gm_ips", fieldName: "account_id", idStart: 0},
		{tableName: "group_id", fieldName: "groupid", idStart: 0},
		{tableName: "group_leaders", fieldName: "gid", idStart: 0},
		{tableName: "guild_bank", fieldName: "guildid", idStart: 0},
		{tableName: "guild_members", fieldName: "char_id", idStart: 0},
		{tableName: "guild_ranks", fieldName: "guild_id", idStart: 0},
		{tableName: "guild_relations", fieldName: "guild1", idStart: 0},
		{tableName: "guilds", fieldName: "id", idStart: 0},
		{tableName: "hackers", fieldName: "id", idStart: 0},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "instance_list_player", fieldName: "id", idStart: 0},
		{tableName: "inventory", fieldName: "charid", idStart: 0},
		{tableName: "inventory_snapshots", fieldName: "charid", idStart: 0},
		{tableName: "keyring", fieldName: "char_id", idStart: 0},
		{tableName: "lfguild", fieldName: "type", idStart: -1},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "mail", fieldName: "charid", idStart: 0},
		{tableName: "character_enabledtasks", fieldName: "charid", idStart: 0},
		{tableName: "petitions", fieldName: "dib", idStart: 0},
		{tableName: "player_titlesets", fieldName: "id", idStart: 0},
		{tableName: "qs_merchant_transaction_record", fieldName: "transaction_id", idStart: 0},
		{tableName: "qs_merchant_transaction_record_entries", fieldName: "event_id", idStart: 0},
		{tableName: "qs_player_aa_rate_hourly", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_delete_record", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_delete_record_entries", fieldName: "event_id", idStart: 0},
		{tableName: "qs_player_events", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_handin_record", fieldName: "handin_id", idStart: 0},
		{tableName: "qs_player_handin_record_entries", fieldName: "event_id", idStart: 0},
		{tableName: "qs_player_move_record", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_move_record_entries", fieldName: "event_id", idStart: 0},
		{tableName: "qs_player_move_record", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_npc_kill_record", fieldName: "fight_id", idStart: 0},
		{tableName: "qs_player_npc_kill_record_entries", fieldName: "char_id", idStart: 0},
		{tableName: "qs_player_speech", fieldName: "id", idStart: 0},
		{tableName: "qs_player_trade_record", fieldName: "trade_id", idStart: 0},
		{tableName: "qs_player_trade_record_entries", fieldName: "event_id", idStart: 0},
		{tableName: "quest_globals", fieldName: "charid", idStart: 0},
		{tableName: "raid_details", fieldName: "raidid", idStart: 0},
		{tableName: "raid_leaders", fieldName: "gid", idStart: 0},
		{tableName: "raid_members", fieldName: "raidid", idStart: 0},
		{tableName: "reports", fieldName: "id", idStart: 0},
		{tableName: "sharedbank", fieldName: "acctid", idStart: 0},
		{tableName: "timers", fieldName: "char_id", idStart: 0},
		{tableName: "trader", fieldName: "char_id", idStart: 0},
		{tableName: "trader_audit", fieldName: "quantity", idStart: 0},
		{tableName: "zone_flags", fieldName: "charid", idStart: 0},
	}

	var affect int64
	for _, table := range tables {
		affect, err = DeleteQuery(fmt.Sprintf("DELETE FROM %s WHERE %s > ?", table.tableName, table.fieldName), table.idStart, fmt.Sprintf("%s delete", table.tableName))
		if err != nil {
			fmt.Errorf("Error deleting %s: %s", table.tableName, err.Error())
			return
		}
		totalRemoved += affect

		_, err = db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table.tableName))
		if err != nil {
			fmt.Errorf("Error resetting auto increment: %s", err.Error())
			return
		}
	}

	return
}

func DeleteQuery(query string, value int64, focus string) (affect int64, err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}
	var result sql.Result

	result, err = db.Exec(query, value)
	if err != nil {
		fmt.Println("Error removing ", value, "from", focus, ":", err.Error())
		return
	}

	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", focus, value, err.Error())
		return
	}
	return
}
