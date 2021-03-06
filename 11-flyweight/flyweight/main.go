package flyweight

type Shop struct {
	ID                uint32
	Username, Country string
}

// Query Types

var (
	shopIDMap = make(map[uint32]Shop)

	shopUsernameMap = make(map[string]Shop)

	// O(1) access for all shop data, O(N) memory
	shopsCountryMap = make(map[string][]Shop)
)

func GetShopByID(ID uint32) Shop {
	if shop, ok := shopIDMap[ID]; ok {
		return shop
	}
	return Shop{}
}

func GetShopByUsername(username string) Shop {
	if shop, ok := shopUsernameMap[username]; ok {
		return shop
	}
	return Shop{}
}

func GetShopsByCountry(country string) []Shop {
	if shops, ok := shopsCountryMap[country]; ok {
		return shops
	}
	return []Shop{}
}

// Flyweight Consumption Method

var (
	idShopMap = make(map[uint32]*Shop)

	usernameIdMap = make(map[string]uint32)

	countryIdsMap = make(map[string][]uint32)
)


func FindShopByID(ID uint32) *Shop {
	if shop, ok := idShopMap[ID]; ok {
		return shop
	}
	return nil
}

func FindShopByUsername(username string) *Shop {
	if shopID, ok := usernameIdMap[username]; ok {
		return FindShopByID(shopID)
	}
	return nil
}

func FindShopsByCountry(country string) []*Shop {
	if shopIDs, ok := countryIdsMap[country]; ok {
		shops := make([]*Shop, len(shopIDs))
		for i, shopID := range shopIDs {
			shop := FindShopByID(shopID)
			shops[i] = shop
		}
		return shops
	}
	return nil
}
