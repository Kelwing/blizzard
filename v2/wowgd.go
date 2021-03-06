package blizzard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/wowgd"
)

// WoWAchievementCategoriesIndex returns an index of achievement categories.
func (c *Client) WoWAchievementCategoriesIndex(ctx context.Context) (*wowgd.AchievementCategoriesIndex, []byte, error) {
	var (
		dat wowgd.AchievementCategoriesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/achievement-category/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementCategory returns an achievement category by ID.
func (c *Client) WoWAchievementCategory(ctx context.Context, achievementCategoryID int) (*wowgd.AchievementCategory, []byte, error) {
	var (
		dat wowgd.AchievementCategory
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/achievement-category/%d?locale=%s", achievementCategoryID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementIndex returns an index of achievements.
func (c *Client) WoWAchievementIndex(ctx context.Context) (*wowgd.AchievementIndex, []byte, error) {
	var (
		dat wowgd.AchievementIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/achievement/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievement returns an achievement category by ID.
func (c *Client) WoWAchievement(ctx context.Context, achievementID int) (*wowgd.Achievement, []byte, error) {
	var (
		dat wowgd.Achievement
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/achievement/%d?locale=%s", achievementID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementMedia returns media for an achievement by ID.
func (c *Client) WoWAchievementMedia(ctx context.Context, achievementID int) (*wowgd.AchievementMedia, []byte, error) {
	var (
		dat wowgd.AchievementMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/achievement/%d?locale=%s", achievementID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssenceIndex returns an index of azerite essences.
func (c *Client) WoWAzeriteEssenceIndex(ctx context.Context) (*wowgd.AzeriteEssenceIndex, []byte, error) {
	var (
		dat wowgd.AzeriteEssenceIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/azerite-essence/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssence returns an azerite essence by ID.
func (c *Client) WoWAzeriteEssence(ctx context.Context, azeriteEssenceID int) (*wowgd.AzeriteEssence, []byte, error) {
	var (
		dat wowgd.AzeriteEssence
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/azerite-essence/%d?locale=%s", azeriteEssenceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssenceMedia returns media for an azerite essence by ID.
func (c *Client) WoWAzeriteEssenceMedia(ctx context.Context, azeriteEssenceID int) (*wowgd.AzeriteEssenceMedia, []byte, error) {
	var (
		dat wowgd.AzeriteEssenceMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/azerite-essence/%d?locale=%s", azeriteEssenceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWConnectedRealmIndex returns an index of connected realms
func (c *Client) WoWConnectedRealmIndex(ctx context.Context) (*wowgd.ConnectedRealmIndex, []byte, error) {
	var (
		dat wowgd.ConnectedRealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWConnectedRealm returns a single connected realm by ID
func (c *Client) WoWConnectedRealm(ctx context.Context, connectedRealmID int) (*wowgd.ConnectedRealm, []byte, error) {
	var (
		dat wowgd.ConnectedRealm
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d?locale=%s", connectedRealmID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) WoWCreatureFamiliesIndex(ctx context.Context) (*wowgd.CreatureFamiliesIndex, []byte, error) {
	var (
		dat wowgd.CreatureFamiliesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-family/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureFamily returns a creature family by ID.
func (c *Client) WoWCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowgd.CreatureFamily, []byte, error) {
	var (
		dat wowgd.CreatureFamily
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureTypesIndex returns an index of creature types.
func (c *Client) WoWCreatureTypesIndex(ctx context.Context) (*wowgd.CreatureTypesIndex, []byte, error) {
	var (
		dat wowgd.CreatureTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-type/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureType returns a creature type by ID.
func (c *Client) WoWCreatureType(ctx context.Context, creatureTypeID int) (*wowgd.CreatureType, []byte, error) {
	var (
		dat wowgd.CreatureType
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-type/%d?locale=%s", creatureTypeID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreature returns a creature type by ID.
func (c *Client) WoWCreature(ctx context.Context, creatureID int) (*wowgd.Creature, []byte, error) {
	var (
		dat wowgd.Creature
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature/%d?locale=%s", creatureID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureDisplayMedia returns media for a creature display by ID.
func (c *Client) WoWCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowgd.CreatureDisplayMedia, []byte, error) {
	var (
		dat wowgd.CreatureDisplayMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/creature-display/%d?locale=%s", creatureDisplayID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCreatureFamilyMedia returns media for a creature family by ID.
func (c *Client) WoWCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowgd.CreatureFamilyMedia, []byte, error) {
	var (
		dat wowgd.CreatureFamilyMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuildCrestComponentsIndex returns an index of guild crest media.
func (c *Client) WoWGuildCrestComponentsIndex(ctx context.Context) (*wowgd.GuildCrestComponentsIndex, []byte, error) {
	var (
		dat wowgd.GuildCrestComponentsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/guild-crest/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuildCrestBorderMedia returns media for a guild crest border by ID.
func (c *Client) WoWGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowgd.GuildCrestBorderMdedia, []byte, error) {
	var (
		dat wowgd.GuildCrestBorderMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/border/%d?locale=%s", borderID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuildCrestEmblemMedia returns media for a guild crest emblem by ID.
func (c *Client) WoWGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowgd.GuildCrestEmblemMdedia, []byte, error) {
	var (
		dat wowgd.GuildCrestEmblemMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d?locale=%s", emblemID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWItemClassesIndex returns an index of item classes.
func (c *Client) WoWItemClassesIndex(ctx context.Context) (*wowgd.ItemClassesIndex, []byte, error) {
	var (
		dat wowgd.ItemClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item-class/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWItemClass returns an item class by ID.
func (c *Client) WoWItemClass(ctx context.Context, itemClassID int) (*wowgd.ItemClass, []byte, error) {
	var (
		dat wowgd.ItemClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item-class/%d?locale=%s", itemClassID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWItemSubclass returns an item subclass by ID.
func (c *Client) WoWItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowgd.ItemSubclass, []byte, error) {
	var (
		dat wowgd.ItemSubclass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx,
		c.apiURL+fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d?locale=%s", itemClassID, itemSubclassID, c.locale),
		c.staticNamespace,
	)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWItem returns an item by ID.
func (c *Client) WoWItem(ctx context.Context, itemID int) (*wowgd.Item, []byte, error) {
	var (
		dat wowgd.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item/%d?locale=%s", itemID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWItemMedia returns media for an item by ID.
func (c *Client) WoWItemMedia(ctx context.Context, itemID int) (*wowgd.ItemMedia, []byte, error) {
	var (
		dat wowgd.ItemMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/item/%d?locale=%s", itemID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalExpansionsIndex returns an index of journal expansions.
func (c *Client) WoWJournalExpansionsIndex(ctx context.Context) (*wowgd.JournalExpansionsIndex, []byte, error) {
	var (
		dat wowgd.JournalExpansionsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-expansion/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalExpansion returns a journal expansion by ID.
func (c *Client) WoWJournalExpansion(ctx context.Context, journalExpansionID int) (*wowgd.JournalExpansion, []byte, error) {
	var (
		dat wowgd.JournalExpansion
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-expansion/%d?locale=%s",
		journalExpansionID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalEncountersIndex returns an index of journal encounters.
func (c *Client) WoWJournalEncountersIndex(ctx context.Context) (*wowgd.JournalEncountersIndex, []byte, error) {
	var (
		dat wowgd.JournalEncountersIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-encounter/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalEncounter returns a journal expansion by ID.
func (c *Client) WoWJournalEncounter(ctx context.Context, journalEncounterID int) (*wowgd.JournalEncounter, []byte, error) {
	var (
		dat wowgd.JournalEncounter
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-encounter/%d?locale=%s",
		journalEncounterID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalInstancesIndex returns an index of journal instances.
func (c *Client) WoWJournalInstancesIndex(ctx context.Context) (*wowgd.JournalInstancesIndex, []byte, error) {
	var (
		dat wowgd.JournalInstancesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-instance/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWJournalInstance returns a journal instance.
func (c *Client) WoWJournalInstance(ctx context.Context, journalInstanceID int) (*wowgd.JournalInstance, []byte, error) {
	var (
		dat wowgd.JournalInstance
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/journal-instance/%d?locale=%s",
		journalInstanceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMountIndex returns an index of mounts.
func (c *Client) WoWMountIndex(ctx context.Context) (*wowgd.MountIndex, []byte, error) {
	var (
		dat wowgd.MountIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mount/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMount returns a mount by ID.
func (c *Client) WoWMount(ctx context.Context, mountID int) (*wowgd.Mount, []byte, error) {
	var (
		dat wowgd.Mount
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mount/%d?locale=%s", mountID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneAffixIndex returns an index of Keystone affixes
func (c *Client) WoWMythicKeystoneAffixIndex(ctx context.Context) (*wowgd.MythicKeystoneAffixIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneAffixIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/keystone-affix/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneAffix returns a single connected realm by ID
func (c *Client) WoWMythicKeystoneAffix(ctx context.Context, keystoneAffixID int) (*wowgd.MythicKeystoneAffix, []byte, error) {
	var (
		dat wowgd.MythicKeystoneAffix
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/keystone-affix/%d?locale=%s", keystoneAffixID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicRaidLeaderboard returns the leaderboard for a given raid and faction
func (c *Client) WoWMythicRaidLeaderboard(ctx context.Context, raid, faction string) (*wowgd.MythicRaidLeaderboard, []byte, error) {
	var (
		dat wowgd.MythicRaidLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/leaderboard/hall-of-fame/%s/%s?locale=%s", raid, faction, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneDungeonIndex returns an index of Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneDungeonIndex(ctx context.Context) (*wowgd.MythicKeystoneDungeonIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneDungeonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/dungeon/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneDungeon returns a Mythic Keystone dungeon by ID
func (c *Client) WoWMythicKeystoneDungeon(ctx context.Context, dungeonID int) (*wowgd.MythicKeystoneDungeon, []byte, error) {
	var (
		dat wowgd.MythicKeystoneDungeon
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/dungeon/%d?locale=%s", dungeonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneIndex returns n index of links to other documents related to Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneIndex(ctx context.Context) (*wowgd.MythicKeystoneIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystonePeriodIndex returns an index of Mythic Keystone periods
func (c *Client) WoWMythicKeystonePeriodIndex(ctx context.Context) (*wowgd.MythicKeystonePeriodIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystonePeriodIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/period/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystonePeriod returns a Mythic Keystone period by ID
func (c *Client) WoWMythicKeystonePeriod(ctx context.Context, periodID int) (*wowgd.MythicKeystonePeriod, []byte, error) {
	var (
		dat wowgd.MythicKeystonePeriod
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/period/%d?locale=%s", periodID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneSeasonIndex returns an index of Mythic Keystone seasons
func (c *Client) WoWMythicKeystoneSeasonIndex(ctx context.Context) (*wowgd.MythicKeystoneSeasonIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneSeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/season/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneSeason returns a Mythic Keystone season by ID
func (c *Client) WoWMythicKeystoneSeason(ctx context.Context, seasonID int) (*wowgd.MythicKeystoneSeason, []byte, error) {
	var (
		dat wowgd.MythicKeystoneSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/season/%d?locale=%s", seasonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneLeaderboardIndex returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm
func (c *Client) WoWMythicKeystoneLeaderboardIndex(ctx context.Context, connectedRealmID int) (*wowgd.MythicKeystoneLeaderboardIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneLeaderboardIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/index?locale=%s", connectedRealmID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneLeaderboard returns a weekly Mythic Keystone Leaderboard by period
func (c *Client) WoWMythicKeystoneLeaderboard(ctx context.Context, connectedRealmID, dungeonID, period int) (*wowgd.MythicKeystoneLeaderboard, []byte, error) {
	var (
		dat wowgd.MythicKeystoneLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/%d/period/%d?locale=%s", connectedRealmID, dungeonID, period, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPetIndex returns an index of pets.
func (c *Client) WoWPetIndex(ctx context.Context) (*wowgd.PetIndex, []byte, error) {
	var (
		dat wowgd.PetIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pet/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPet returns a pet by ID.
func (c *Client) WoWPet(ctx context.Context, petID int) (*wowgd.Pet, []byte, error) {
	var (
		dat wowgd.Pet
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pet/%d?locale=%s", petID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPetMedia returns media information for a pet ID.
func (c *Client) WoWPetMedia(ctx context.Context, petID int) (*wowgd.PetMedia, []byte, error) {
	var (
		dat wowgd.PetMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/pet/%d?locale=%s", petID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPetAbilityIndex returns an index of pet abilities.
func (c *Client) WoWPetAbilityIndex(ctx context.Context) (*wowgd.PetAbilityIndex, []byte, error) {
	var (
		dat wowgd.PetAbilityIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pet-ability/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPetAbility returns a pet ability by ID.
func (c *Client) WoWPetAbility(ctx context.Context, petAbilityID int) (*wowgd.PetAbility, []byte, error) {
	var (
		dat wowgd.PetAbility
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pet-ability/%d?locale=%s", petAbilityID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPetAbilityMedia returns media information for a pet ID.
func (c *Client) WoWPetAbilityMedia(ctx context.Context, petAbilityID int) (*wowgd.PetMedia, []byte, error) {
	var (
		dat wowgd.PetMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/pet-ability/%d?locale=%s", petAbilityID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableClassesIndex returns an index of playable classes
func (c *Client) WoWPlayableClassesIndex(ctx context.Context) (*wowgd.PlayableClassesIndex, []byte, error) {
	var (
		dat wowgd.PlayableClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-class/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableClass returns a playable class by ID
func (c *Client) WoWPlayableClass(ctx context.Context, classID int) (*wowgd.PlayableClass, []byte, error) {
	var (
		dat wowgd.PlayableClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d?locale=%s", classID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableClassPvPTalentSlots returns the PvP talent slots for a playable class by ID
func (c *Client) WoWPlayableClassPvPTalentSlots(ctx context.Context, classID int) (*wowgd.PlayableClassPvPTalentSlots, []byte, error) {
	var (
		dat wowgd.PlayableClassPvPTalentSlots
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d/pvp-talent-slots?locale=%s", classID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableRacesIndex returns an index of races.
func (c *Client) WoWPlayableRacesIndex(ctx context.Context) (*wowgd.PlayableRacesIndex, []byte, error) {
	var (
		dat wowgd.PlayableRacesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-race/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableRace returns a race by ID.
func (c *Client) WoWPlayableRace(ctx context.Context, raceID int) (*wowgd.PlayableRace, []byte, error) {
	var (
		dat wowgd.PlayableRace
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-race/%d?locale=%s", raceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecializationIndex returns an index of playable specializations.
func (c *Client) WoWPlayableSpecializationIndex(ctx context.Context) (*wowgd.PlayableSpecializationIndex, []byte, error) {
	var (
		dat wowgd.PlayableSpecializationIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-specialization/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecialization returns a playable specialization by ID.
func (c *Client) WoWPlayableSpecialization(ctx context.Context, specID int) (*wowgd.PlayableSpecialization, []byte, error) {
	var (
		dat wowgd.PlayableSpecialization
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-specialization/%d?locale=%s", specID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPowerTypesIndex returns an index of power types.
func (c *Client) WoWPowerTypesIndex(ctx context.Context) (*wowgd.PowerTypesIndex, []byte, error) {
	var (
		dat wowgd.PowerTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/power-type/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPowerType returns a power type by ID.
func (c *Client) WoWPowerType(ctx context.Context, powerTypeID int) (*wowgd.PowerType, []byte, error) {
	var (
		dat wowgd.PowerType
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/power-type/%d?locale=%s", powerTypeID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPSeasonIndex returns an index of PvP seasons.
func (c *Client) WoWPvPSeasonIndex(ctx context.Context) (*wowgd.PvPSeasonIndex, []byte, error) {
	var (
		dat wowgd.PvPSeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-season/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPSeason returns a PvP season by ID.
func (c *Client) WoWPvPSeason(ctx context.Context, pvpSeasonID int) (*wowgd.PvPSeason, []byte, error) {
	var (
		dat wowgd.PvPSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-season/%d?locale=%s", pvpSeasonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPLeaderboardsIndex returns an index of PvP leaderboards for a PvP season.
func (c *Client) WoWPvPLeaderboardsIndex(ctx context.Context, pvpSeasonID int) (*wowgd.PvPLeaderboardsIndex, []byte, error) {
	var (
		dat wowgd.PvPLeaderboardsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-season/%d/pvp-leaderboard/index?locale=%s", pvpSeasonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPLeaderboard returns the PvP leaderboard of a specific PvP bracket for a PvP season.
func (c *Client) WoWPvPLeaderboard(ctx context.Context, pvpSeasonID int, pvpBracket wowgd.Bracket) (*wowgd.PvPLeaderboard, []byte, error) {
	var (
		dat wowgd.PvPLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-season/%d/pvp-leaderboard/%s?locale=%s", pvpSeasonID, pvpBracket, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPRewardsIndex returns an index of PvP rewards for a PvP season.
func (c *Client) WoWPvPRewardsIndex(ctx context.Context, pvpSeasonID int) (*wowgd.PvPRewardsIndex, []byte, error) {
	var (
		dat wowgd.PvPRewardsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-season/%d/pvp-reward/index?locale=%s", pvpSeasonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPTierMedia returns media for a PvP tier by ID.
func (c *Client) WoWPvPTierMedia(ctx context.Context, pvpTierID int) (*wowgd.PvPTierMedia, []byte, error) {
	var (
		dat wowgd.PvPTierMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/pvp-tier/%d?locale=%s", pvpTierID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPTiersIndex returns an index of PvP tiers.
func (c *Client) WoWPvPTiersIndex(ctx context.Context) (*wowgd.PvPTiersIndex, []byte, error) {
	var (
		dat wowgd.PvPTiersIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-tier/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPvPTier returns a PvP tier by ID.
func (c *Client) WoWPvPTier(ctx context.Context, pvpTierID int) (*wowgd.PvPTier, []byte, error) {
	var (
		dat wowgd.PvPTier
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-tier/%d?locale=%s", pvpTierID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWRealmIndex returns an index of realms.
func (c *Client) WoWRealmIndex(ctx context.Context) (*wowgd.RealmIndex, []byte, error) {
	var (
		dat wowgd.RealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/realm/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWRealm returns a single realm by slug or ID.
func (c *Client) WoWRealm(ctx context.Context, realmSlug string) (*wowgd.Realm, []byte, error) {
	var (
		dat wowgd.Realm
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/realm/%s?locale=%s", realmSlug, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWRegionIndex returns an index of regions.
func (c *Client) WoWRegionIndex(ctx context.Context) (*wowgd.RegionIndex, []byte, error) {
	var (
		dat wowgd.RegionIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/region/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWRegion returns a single region by ID.
func (c *Client) WoWRegion(ctx context.Context, regionID int) (*wowgd.Region, []byte, error) {
	var (
		dat wowgd.Region
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/region/%d?locale=%s", regionID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWReputationFactionsIndex returns an index of reputation factions.
func (c *Client) WoWReputationFactionsIndex(ctx context.Context) (*wowgd.ReputationFactionsIndex, []byte, error) {
	var (
		dat wowgd.ReputationFactionsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/reputation-faction/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWReputationFaction returns a single reputation faction by ID.
func (c *Client) WoWReputationFaction(ctx context.Context, reputationFactionID int) (*wowgd.ReputationFaction, []byte, error) {
	var (
		dat wowgd.ReputationFaction
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/pvp-tier/%d?locale=%s", reputationFactionID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWReputationTiersIndex returns an index of reputation tiers.
func (c *Client) WoWReputationTiersIndex(ctx context.Context) (*wowgd.ReputationTiersIndex, []byte, error) {
	var (
		dat wowgd.ReputationTiersIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/reputation-tiers/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWReputationTiers returns a single set of reputation tiers by ID.
func (c *Client) WoWReputationTiers(ctx context.Context, reputationTiersID int) (*wowgd.ReputationTiers, []byte, error) {
	var (
		dat wowgd.ReputationTiers
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/reputation-tiers/%d?locale=%s", reputationTiersID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWTitlesIndex returns an index of titles.
func (c *Client) WoWTitlesIndex(ctx context.Context) (*wowgd.TitlesIndex, []byte, error) {
	var (
		dat wowgd.TitlesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/title/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWTitle returns a title by ID.
func (c *Client) WoWTitle(ctx context.Context, titleID int) (*wowgd.Title, []byte, error) {
	var (
		dat wowgd.Title
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/title/%d?locale=%s", titleID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWToken returns the WoW Token index
func (c *Client) WoWToken(ctx context.Context) (*wowgd.Token, []byte, error) {
	var (
		dat wowgd.Token
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/token/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
