package consts

const (
	// TAG FK => Filter Key
	// Groups
	TAG    = "TAG" // Firsat Urunleri
	TAG_FK = "Tag"

	CATEGORY    = "CATEGORY"
	CATEGORY_FK = "WebCategory"

	BRAND    = "BRAND"
	BRAND_FK = "WebBrand"

	VARIANT    = "VARIANT"
	VARIANT_FK = "beden"

	PRICE    = "PRICE"
	PRICE_FK = "Price"

	GENDER    = "GENDER"
	GENDER_FK = "WebGender"

	ATTRIBUTE = "ATTRIBUTE"

	ATTRIBUTE_ARM_TYPE_FK         = "12"  // Kol tipi
	ATTRIBUTE_MATERIAL_FK         = "14"  // Materyal
	ATTRIBUTE_MOLD_FK             = "179" // Kalip
	ATRRIBUTE_COLLAR_FK           = "22"  // Yaka
	ATTRIBUTE_STYLE_FK            = "250" // Stil
	ATTRIBUTE_USAGE_AREA_FK       = "196" // Kullanim alanlari
	ATTRIBUTE_ARM_LENGTH_FK       = "191" // Kol Boyu
	ATTRIBUTE_PATTERN_FK          = "33"  // Desen
	ATTRIBUTE_TICKNESS_FK         = "178" // Kalinlik
	ATTRIBUTE_FABRIC_TYPE_FK      = "200" // Kumaş Tipi
	ATTRIBUTE_PRODUCT_TYPE_FK     = "342" // Ürün Tipi
	ATTRIBUTE_PACKAGE_INCLUDED_FK = "66"  // Paket İçeriği
	ATTRIBUTE_TEKNIK_FK           = "636" // Teknik
	ATTRIBUTE_FABRIC_FEATURE_FK   = "440" // Kumaş Teknolojisi
	ATTRIBUTE_KOLEKSIYON_FK       = "710" // Koleksiyon
	ATTRIBUTE_WEB_COLOR_FK        = "348" // Renk
)

// DISPLAY_MODE products_and_description means Category otherwise Promotion
const DISPLAY_MODE = "products_and_description"

const (
	TRENDYOL_URL = "https://www.trendyol.com/"
	LCW_URL      = "https://www.lcwaikiki.com/tr-TR/TR/"
)

const (
	DB_LCW      = "lcw_db"
	DB_TRENDYOL = "trendyol_db"
)
