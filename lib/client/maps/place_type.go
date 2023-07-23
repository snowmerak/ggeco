package maps

/*
accounting
airport
amusement_park
aquarium
art_gallery
atm
bakery
bank
bar
beauty_salon
bicycle_store
book_store
bowling_alley
bus_station
cafe
campground
car_dealer
car_rental
car_repair
car_wash
casino
cemetery
church
city_hall
clothing_store
convenience_store
courthouse
dentist
department_store
doctor
drugstore
electrician
electronics_store
embassy
fire_station
florist
funeral_home
furniture_store
gas_station
gym
hair_care
hardware_store
hindu_temple
home_goods_store
hospital
insurance_agency
jewelry_store
laundry
lawyer
library
light_rail_station
liquor_store
local_government_office
locksmith
lodging
meal_delivery
meal_takeaway
mosque
movie_rental
movie_theater
moving_company
museum
night_club
painter
park
parking
pet_store
pharmacy
physiotherapist
plumber
police
post_office
primary_school
real_estate_agency
restaurant
roofing_contractor
rv_park
school
secondary_school
shoe_store
shopping_mall
spa
stadium
storage
store
subway_station
supermarket
synagogue
taxi_stand
tourist_attraction
train_station
transit_station
travel_agency
university
veterinary_care
zoo

administrative_area_level_1
administrative_area_level_2
administrative_area_level_3
administrative_area_level_4
administrative_area_level_5
administrative_area_level_6
administrative_area_level_7
archipelago
colloquial_area
continent
country
establishment
finance
floor
food
general_contractor
geocode
health
intersection
landmark
locality
natural_feature
neighborhood
place_of_worship
plus_code
point_of_interest
political
post_box
postal_code
postal_code_prefix
postal_code_suffix
postal_town
premise
room
route
street_address
street_number
sublocality
sublocality_level_1
sublocality_level_2
sublocality_level_3
sublocality_level_4
sublocality_level_5
subpremise
town_square
*/
const (
	PlaceTypeAccounting            = "accounting"
	PlaceTypeAirport               = "airport"
	PlaceTypeAmusementPark         = "amusement_park"
	PlaceTypeAquarium              = "aquarium"
	PlaceTypeArtGallery            = "art_gallery"
	PlaceTypeAtm                   = "atm"
	PlaceTypeBakery                = "bakery"
	PlaceTypeBank                  = "bank"
	PlaceTypeBar                   = "bar"
	PlaceTypeBeautySalon           = "beauty_salon"
	PlaceTypeBicycleStore          = "bicycle_store"
	PlaceTypeBookStore             = "book_store"
	PlaceTypeBowlingAlley          = "bowling_alley"
	PlaceTypeBusStation            = "bus_station"
	PlaceTypeCafe                  = "cafe"
	PlaceTypeCampground            = "campground"
	PlaceTypeCarDealer             = "car_dealer"
	PlaceTypeCarRental             = "car_rental"
	PlaceTypeCarRepair             = "car_repair"
	PlaceTypeCarWash               = "car_wash"
	PlaceTypeCasino                = "casino"
	PlaceTypeCemetery              = "cemetery"
	PlaceTypeChurch                = "church"
	PlaceTypeCityHall              = "city_hall"
	PlaceTypeClothingStore         = "clothing_store"
	PlaceTypeConvenienceStore      = "convenience_store"
	PlaceTypeCourthouse            = "courthouse"
	PlaceTypeDentist               = "dentist"
	PlaceTypeDepartmentStore       = "department_store"
	PlaceTypeDoctor                = "doctor"
	PlaceTypeDrugstore             = "drugstore"
	PlaceTypeElectrician           = "electrician"
	PlaceTypeElectronicsStore      = "electronics_store"
	PlaceTypeEmbassy               = "embassy"
	PlaceTypeFireStation           = "fire_station"
	PlaceTypeFlorist               = "florist"
	PlaceTypeFuneralHome           = "funeral_home"
	PlaceTypeFurnitureStore        = "furniture_store"
	PlaceTypeGasStation            = "gas_station"
	PlaceTypeGym                   = "gym"
	PlaceTypeHairCare              = "hair_care"
	PlaceTypeHardwareStore         = "hardware_store"
	PlaceTypeHinduTemple           = "hindu_temple"
	PlaceTypeHomeGoodsStore        = "home_goods_store"
	PlaceTypeHospital              = "hospital"
	PlaceTypeInsuranceAgency       = "insurance_agency"
	PlaceTypeJewelryStore          = "jewelry_store"
	PlaceTypeLaundry               = "laundry"
	PlaceTypeLawyer                = "lawyer"
	PlaceTypeLibrary               = "library"
	PlaceTypeLightRailStation      = "light_rail_station"
	PlaceTypeLiquorStore           = "liquor_store"
	PlaceTypeLocalGovernmentOffice = "local_government_office"
	PlaceTypeLocksmith             = "locksmith"
	PlaceTypeLodging               = "lodging"
	PlaceTypeMealDelivery          = "meal_delivery"
	PlaceTypeMealTakeaway          = "meal_takeaway"
	PlaceTypeMosque                = "mosque"
	PlaceTypeMovieRental           = "movie_rental"
	PlaceTypeMovieTheater          = "movie_theater"
	PlaceTypeMovingCompany         = "moving_company"
	PlaceTypeMuseum                = "museum"
	PlaceTypeNightClub             = "night_club"
	PlaceTypePainter               = "painter"
	PlaceTypePark                  = "park"
	PlaceTypeParking               = "parking"
	PlaceTypePetStore              = "pet_store"
	PlaceTypePharmacy              = "pharmacy"
	PlaceTypePhysiotherapist       = "physiotherapist"
	PlaceTypePlumber               = "plumber"
	PlaceTypePolice                = "police"
	PlaceTypePostOffice            = "post_office"
	PlaceTypePrimarySchool         = "primary_school"
	PlaceTypeRealEstateAgency      = "real_estate_agency"
	PlaceTypeRestaurant            = "restaurant"
	PlaceTypeRoofingContractor     = "roofing_contractor"
	PlaceTypeRvPark                = "rv_park"
	PlaceTypeSchool                = "school"
	PlaceTypeSecondarySchool       = "secondary_school"
	PlaceTypeShoeStore             = "shoe_store"
	PlaceTypeShoppingMall          = "shopping_mall"
	PlaceTypeSpa                   = "spa"
	PlaceTypeStadium               = "stadium"
	PlaceTypeStorage               = "storage"
	PlaceTypeStore                 = "store"
	PlaceTypeSubwayStation         = "subway_station"
	PlaceTypeSupermarket           = "supermarket"
	PlaceTypeSynagogue             = "synagogue"
	PlaceTypeTaxiStand             = "taxi_stand"
	PlaceTypeTouristAttraction     = "tourist_attraction"
	PlaceTypeTrainStation          = "train_station"
	PlaceTypeTransitStation        = "transit_station"
	PlaceTypeTravelAgency          = "travel_agency"
	PlaceTypeUniversity            = "university"
	PlaceTypeVeterinaryCare        = "veterinary_care"
	PlaceTypeZoo                   = "zoo"

	PlaceTypeAdministrativeAreaLevel1 = "administrative_area_level_1"
	PlaceTypeAdministrativeAreaLevel2 = "administrative_area_level_2"
	PlaceTypeAdministrativeAreaLevel3 = "administrative_area_level_3"
	PlaceTypeAdministrativeAreaLevel4 = "administrative_area_level_4"
	PlaceTypeAdministrativeAreaLevel5 = "administrative_area_level_5"
	PlaceTypeColloquialArea           = "colloquial_area"
	PlaceTypeCountry                  = "country"
	PlaceTypeEstablishment            = "establishment"
	PlaceTypeFinance                  = "finance"
	PlaceTypeFloor                    = "floor"
	PlaceTypeFood                     = "food"
	PlaceTypeGeneralContractor        = "general_contractor"
	PlaceTypeGeocode                  = "geocode"
	PlaceTypeHealth                   = "health"
	PlaceTypeIntersection             = "intersection"
	PlaceTypeLocality                 = "locality"
	PlaceTypeNaturalFeature           = "natural_feature"
	PlaceTypeNeighborhood             = "neighborhood"
	PlaceTypePlaceOfWorship           = "place_of_worship"
	PlaceTypePlusCode                 = "plus_code"
	PlaceTypePointOfInterest          = "point_of_interest"
	PlaceTypePolitical                = "political"
	PlaceTypePostBox                  = "post_box"
	PlaceTypePostalCode               = "postal_code"
	PlaceTypePostalCodePrefix         = "postal_code_prefix"
	PlaceTypePostalCodeSuffix         = "postal_code_suffix"
	PlaceTypePostalTown               = "postal_town"
	PlaceTypePremise                  = "premise"
	PlaceTypeRoom                     = "room"
	PlaceTypeRoute                    = "route"
	PlaceTypeStreetAddress            = "street_address"
	PlaceTypeStreetNumber             = "street_number"
	PlaceTypeSublocality              = "sublocality"
	PlaceTypeSublocalityLevel1        = "sublocality_level_1"
	PlaceTypeSublocalityLevel2        = "sublocality_level_2"
	PlaceTypeSublocalityLevel3        = "sublocality_level_3"
	PlaceTypeSublocalityLevel4        = "sublocality_level_4"
	PlaceTypeSublocalityLevel5        = "sublocality_level_5"
	PlaceTypeSubpremise               = "subpremise"
	PlaceTypeTownSquare               = "town_square"
)

func TranslatePlaceType(source string) (result string) {
	switch source {
	case PlaceTypeAccounting:
		result = "회계사"
	case PlaceTypeAirport:
		result = "공항"
	case PlaceTypeAmusementPark:
		result = "놀이공원"
	case PlaceTypeAquarium:
		result = "수족관"
	case PlaceTypeArtGallery:
		result = "미술관"
	case PlaceTypeAtm:
		result = "ATM"
	case PlaceTypeBakery:
		result = "빵집"
	case PlaceTypeBank:
		result = "은행"
	case PlaceTypeBar:
		result = "바"
	case PlaceTypeBeautySalon:
		result = "미용실"
	case PlaceTypeBicycleStore:
		result = "자전거점"
	case PlaceTypeBookStore:
		result = "서점"
	case PlaceTypeBowlingAlley:
		result = "볼링장"
	case PlaceTypeBusStation:
		result = "버스정류장"
	case PlaceTypeCafe:
		result = "카페"
	case PlaceTypeCampground:
		result = "야영장"
	case PlaceTypeCarDealer:
		result = "자동차 딜러"
	case PlaceTypeCarRental:
		result = "렌트카"
	case PlaceTypeCarRepair:
		result = "자동차 수리"
	case PlaceTypeCarWash:
		result = "세차장"
	case PlaceTypeCasino:
		result = "카지노"
	case PlaceTypeCemetery:
		result = "묘지"
	case PlaceTypeChurch:
		result = "교회"
	case PlaceTypeCityHall:
		result = "시청"
	case PlaceTypeClothingStore:
		result = "의류점"
	case PlaceTypeConvenienceStore:
		result = "편의점"
	case PlaceTypeCourthouse:
		result = "법원"
	case PlaceTypeDentist:
		result = "치과의사"
	case PlaceTypeDepartmentStore:
		result = "백화점"
	case PlaceTypeDoctor:
		result = "의사"
	case PlaceTypeDrugstore:
		result = "약국"
	case PlaceTypeElectrician:
		result = "전기기사"
	case PlaceTypeElectronicsStore:
		result = "전자제품점"
	case PlaceTypeEmbassy:
		result = "대사관"
	case PlaceTypeFireStation:
		result = "소방서"
	case PlaceTypeFlorist:
		result = "꽃집"
	case PlaceTypeFuneralHome:
		result = "장례식장"
	case PlaceTypeFurnitureStore:
		result = "가구점"
	case PlaceTypeGasStation:
		result = "주유소"
	case PlaceTypeGym:
		result = "체육관"
	case PlaceTypeHairCare:
		result = "미용실"
	case PlaceTypeHardwareStore:
		result = "철물점"
	case PlaceTypeHinduTemple:
		result = "힌두교 사원"
	case PlaceTypeHomeGoodsStore:
		result = "가정용품점"
	case PlaceTypeHospital:
		result = "병원"
	case PlaceTypeInsuranceAgency:
		result = "보험회사"
	case PlaceTypeJewelryStore:
		result = "보석점"
	case PlaceTypeLaundry:
		result = "세탁소"
	case PlaceTypeLawyer:
		result = "변호사"
	case PlaceTypeLibrary:
		result = "도서관"
	case PlaceTypeLiquorStore:
		result = "주류점"
	case PlaceTypeLocalGovernmentOffice:
		result = "지방정부청사"
	case PlaceTypeLocksmith:
		result = "자물쇠 장수"
	case PlaceTypeLodging:
		result = "숙박"
	case PlaceTypeMealDelivery:
		result = "배달음식점"
	case PlaceTypeMealTakeaway:
		result = "포장음식점"
	case PlaceTypeMosque:
		result = "모스크"
	case PlaceTypeMovieRental:
		result = "비디오 대여점"
	case PlaceTypeMovieTheater:
		result = "영화관"
	case PlaceTypeMovingCompany:
		result = "이사회사"
	case PlaceTypeMuseum:
		result = "박물관"
	case PlaceTypeNightClub:
		result = "클럽"
	case PlaceTypePainter:
		result = "화가"
	case PlaceTypePark:
		result = "공원"
	case PlaceTypeParking:
		result = "주차장"
	case PlaceTypePetStore:
		result = "애완동물용품점"
	case PlaceTypePharmacy:
		result = "약국"
	case PlaceTypePhysiotherapist:
		result = "물리치료사"
	case PlaceTypePlumber:
		result = "배관공"
	case PlaceTypePolice:
		result = "경찰서"
	case PlaceTypePostOffice:
		result = "우체국"
	case PlaceTypeRealEstateAgency:
		result = "부동산"
	case PlaceTypeRestaurant:
		result = "음식점"
	case PlaceTypeRoofingContractor:
		result = "지붕공사업자"
	case PlaceTypeRvPark:
		result = "RV파크"
	case PlaceTypeSchool:
		result = "학교"
	case PlaceTypeShoeStore:
		result = "신발가게"
	case PlaceTypeShoppingMall:
		result = "쇼핑몰"
	case PlaceTypeSpa:
		result = "스파"
	case PlaceTypeStadium:
		result = "경기장"
	case PlaceTypeStorage:
		result = "창고"
	case PlaceTypeStore:
		result = "가게"
	case PlaceTypeSubwayStation:
		result = "지하철역"
	case PlaceTypeSupermarket:
		result = "슈퍼마켓"
	case PlaceTypeSynagogue:
		result = "유대교회"
	case PlaceTypeTaxiStand:
		result = "택시 승강장"
	case PlaceTypeTouristAttraction:
		result = "관광명소"
	case PlaceTypeTrainStation:
		result = "기차역"
	case PlaceTypeTransitStation:
		result = "환승역"
	case PlaceTypeTravelAgency:
		result = "여행사"
	case PlaceTypeVeterinaryCare:
		result = "수의사"
	case PlaceTypeZoo:
		result = "동물원"
	case PlaceTypeAdministrativeAreaLevel1:
		result = "1단계 행정구역"
	case PlaceTypeAdministrativeAreaLevel2:
		result = "2단계 행정구역"
	case PlaceTypeAdministrativeAreaLevel3:
		result = "3단계 행정구역"
	case PlaceTypeAdministrativeAreaLevel4:
		result = "4단계 행정구역"
	case PlaceTypeAdministrativeAreaLevel5:
		result = "5단계 행정구역"
	case PlaceTypeColloquialArea:
		result = "구어체 지역"
	case PlaceTypeCountry:
		result = "국가"
	case PlaceTypeEstablishment:
		result = "시설"
	case PlaceTypeFloor:
		result = "층"
	case PlaceTypeFood:
		result = "음식"
	case PlaceTypeGeneralContractor:
		result = "일반 계약자"
	case PlaceTypeGeocode:
		result = "지오코드"
	case PlaceTypeIntersection:
		result = "교차로"
	case PlaceTypeLocality:
		result = "지역"
	case PlaceTypeNaturalFeature:
		result = "자연지형"
	case PlaceTypeNeighborhood:
		result = "동네"
	case PlaceTypePolitical:
		result = "정치적"
	case PlaceTypePointOfInterest:
		result = "관심지점"
	case PlaceTypePostBox:
		result = "우편함"
	case PlaceTypePostalCode:
		result = "우편번호"
	case PlaceTypePostalCodePrefix:
		result = "우편번호 접두사"
	case PlaceTypePostalTown:
		result = "우편 도시"
	case PlaceTypePremise:
		result = "건물"
	case PlaceTypeRoom:
		result = "방"
	case PlaceTypeRoute:
		result = "경로"
	case PlaceTypeStreetAddress:
		result = "도로명 주소"
	case PlaceTypeSublocality:
		result = "세부지역"
	case PlaceTypeSublocalityLevel1:
		result = "1단계 세부지역"
	case PlaceTypeSublocalityLevel2:
		result = "2단계 세부지역"
	case PlaceTypeSublocalityLevel3:
		result = "3단계 세부지역"
	case PlaceTypeSublocalityLevel4:
		result = "4단계 세부지역"
	case PlaceTypeSublocalityLevel5:
		result = "5단계 세부지역"
	case PlaceTypeSubpremise:
		result = "세부 건물"
	case PlaceTypeTownSquare:
		result = "타운 스퀘어"
	}
	return result
}
