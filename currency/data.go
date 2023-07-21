package currency

var currencyData = [][]string{
	{"AED", "784", "2", "United Arab Emirates dirham"},
	{"AFN", "971", "2", "Afghan afghani"},
	{"ALL", "8", "2", "Albanian lek"},
	{"AMD", "51", "2", "Armenian dram"},
	{"ANG", "532", "2", "Netherlands Antillean guilder"},
	{"AOA", "973", "2", "Angolan kwanza"},
	{"ARS", "32", "2", "Argentine peso"},
	{"AUD", "36", "2", "Australian dollar"},
	{"AWG", "533", "2", "Aruban florin"},
	{"AZN", "944", "2", "Azerbaijani manat"},
	{"BAM", "977", "2", "Bosnia and Herzegovina convertible mark"},
	{"BBD", "52", "2", "Barbados dollar"},
	{"BDT", "50", "2", "Bangladeshi taka"},
	{"BGN", "975", "2", "Bulgarian lev"},
	{"BHD", "48", "3", "Bahraini dinar"},
	{"BIF", "108", "0", "Burundian franc"},
	{"BMD", "60", "2", "Bermudian dollar (customarily known as Bermuda dollar)"},
	{"BND", "96", "2", "Brunei dollar"},
	{"BOB", "68", "2", "Boliviano"},
	{"BOV", "984", "2", "Bolivian Mvdol (funds code)"},
	{"BRL", "986", "2", "Brazilian real"},
	{"BSD", "44", "2", "Bahamian dollar"},
	{"BTN", "64", "2", "Bhutanese ngultrum"},
	{"BWP", "72", "2", "Botswana pula"},
	{"BYR", "974", "0", "Belarusian ruble"},
	{"BZD", "84", "2", "Belize dollar"},
	{"CAD", "124", "2", "Canadian dollar"},
	{"CDF", "976", "2", "Congolese franc"},
	{"CHE", "947", "2", "WIR Euro (complementary currency)"},
	{"CHF", "756", "2", "Swiss franc"},
	{"CHW", "948", "2", "WIR Franc (complementary currency)"},
	{"CLF", "990", "0", "Unidad de Fomento (funds code)"},
	{"CLP", "152", "0", "Chilean peso"},
	{"CNY", "156", "2", "Chinese yuan"},
	{"COP", "170", "2", "Colombian peso"},
	{"COU", "970", "2", "Unidad de Valor Real"},
	{"CRC", "188", "2", "Costa Rican colon"},
	{"CUC", "931", "2", "Cuban convertible peso"},
	{"CUP", "192", "2", "Cuban peso"},
	{"CVE", "132", "0", "Cape Verde escudo"},
	{"CZK", "203", "2", "Czech koruna"},
	{"DJF", "262", "0", "Djiboutian franc"},
	{"DKK", "208", "2", "Danish krone"},
	{"DOP", "214", "2", "Dominican peso"},
	{"DZD", "12", "2", "Algerian dinar"},
	{"EGP", "818", "2", "Egyptian pound"},
	{"ERN", "232", "2", "Eritrean nakfa"},
	{"ETB", "230", "2", "Ethiopian birr"},
	{"EUR", "978", "2", "Euro"},
	{"FJD", "242", "2", "Fiji dollar"},
	{"FKP", "238", "2", "Falkland Islands pound"},
	{"GBP", "826", "2", "Pound sterling"},
	{"GEL", "981", "2", "Georgian lari"},
	{"GHS", "936", "2", "Ghanaian cedi"},
	{"GIP", "292", "2", "Gibraltar pound"},
	{"GMD", "270", "2", "Gambian dalasi"},
	{"GNF", "324", "0", "Guinean franc"},
	{"GTQ", "320", "2", "Guatemalan quetzal"},
	{"GYD", "328", "2", "Guyanese dollar"},
	{"HKD", "344", "2", "Hong Kong dollar"},
	{"HNL", "340", "2", "Honduran lempira"},
	{"HRK", "191", "2", "Croatian kuna"},
	{"HTG", "332", "2", "Haitian gourde"},
	{"HUF", "348", "2", "Hungarian forint"},
	{"IDR", "360", "2", "Indonesian rupiah"},
	{"ILS", "376", "2", "Israeli new shekel"},
	{"INR", "356", "2", "Indian rupee"},
	{"IQD", "368", "3", "Iraqi dinar"},
	{"IRR", "364", "0", "Iranian rial"},
	{"ISK", "352", "0", "Icelandic króna"},
	{"JMD", "388", "2", "Jamaican dollar"},
	{"JOD", "400", "3", "Jordanian dinar"},
	{"JPY", "392", "0", "Japanese yen"},
	{"KES", "404", "2", "Kenyan shilling"},
	{"KGS", "417", "2", "Kyrgyzstani som"},
	{"KHR", "116", "2", "Cambodian riel"},
	{"KMF", "174", "0", "Comoro franc"},
	{"KPW", "408", "0", "North Korean won"},
	{"KRW", "410", "0", "South Korean won"},
	{"KWD", "414", "3", "Kuwaiti dinar"},
	{"KYD", "136", "2", "Cayman Islands dollar"},
	{"KZT", "398", "2", "Kazakhstani tenge"},
	{"LAK", "418", "0", "Lao kip"},
	{"LBP", "422", "0", "Lebanese pound"},
	{"LKR", "144", "2", "Sri Lankan rupee"},
	{"LRD", "430", "2", "Liberian dollar"},
	{"LSL", "426", "2", "Lesotho loti"},
	{"LTL", "440", "2", "Lithuanian litas"},
	{"LVL", "428", "2", "Latvian lats"},
	{"LYD", "434", "3", "Libyan dinar"},
	{"MAD", "504", "2", "Moroccan dirham"},
	{"MDL", "498", "2", "Moldovan leu"},
	// {"MGA", "969", "0.7[8]", "Malagasy ariary"},
	{"MKD", "807", "0", "Macedonian denar"},
	{"MMK", "104", "0", "Myanma kyat"},
	{"MNT", "496", "2", "Mongolian tugrik"},
	{"MOP", "446", "2", "Macanese pataca"},
	// {"MRO", "478", "0.7[8]", "Mauritanian ouguiya"},
	{"MUR", "480", "2", "Mauritian rupee"},
	{"MVR", "462", "2", "Maldivian rufiyaa"},
	{"MWK", "454", "2", "Malawian kwacha"},
	{"MXN", "484", "2", "Mexican peso"},
	{"MXV", "979", "2", "Mexican Unidad de Inversion (UDI) (funds code)"},
	{"MYR", "458", "2", "Malaysian ringgit"},
	{"MZN", "943", "2", "Mozambican metical"},
	{"NAD", "516", "2", "Namibian dollar"},
	{"NGN", "566", "2", "Nigerian naira"},
	{"NIO", "558", "2", "Nicaraguan córdoba"},
	{"NOK", "578", "2", "Norwegian krone"},
	{"NPR", "524", "2", "Nepalese rupee"},
	{"NZD", "554", "2", "New Zealand dollar"},
	{"OMR", "512", "3", "Omani rial"},
	{"PAB", "590", "2", "Panamanian balboa"},
	{"PEN", "604", "2", "Peruvian nuevo sol"},
	{"PGK", "598", "2", "Papua New Guinean kina"},
	{"PHP", "608", "2", "Philippine peso"},
	{"PKR", "586", "2", "Pakistani rupee"},
	{"PLN", "985", "2", "Polish złoty"},
	{"PYG", "600", "0", "Paraguayan guaraní"},
	{"QAR", "634", "2", "Qatari riyal"},
	{"RON", "946", "2", "Romanian new leu"},
	{"RSD", "941", "2", "Serbian dinar"},
	{"RUB", "643", "2", "Russian rouble"},
	{"RWF", "646", "0", "Rwandan franc"},
	{"SAR", "682", "2", "Saudi riyal"},
	{"SBD", "90", "2", "Solomon Islands dollar"},
	{"SCR", "690", "2", "Seychelles rupee"},
	{"SDG", "938", "2", "Sudanese pound"},
	{"SEK", "752", "2", "Swedish krona/kronor"},
	{"SGD", "702", "2", "Singapore dollar"},
	{"SHP", "654", "2", "Saint Helena pound"},
	{"SLL", "694", "0", "Sierra Leonean leone"},
	{"SOS", "706", "2", "Somali shilling"},
	{"SRD", "968", "2", "Surinamese dollar"},
	{"SSP", "728", "2", "South Sudanese pound"},
	{"STD", "678", "0", "São Tomé and Príncipe dobra"},
	{"SYP", "760", "2", "Syrian pound"},
	{"SZL", "748", "2", "Swazi lilangeni"},
	{"THB", "764", "2", "Thai baht"},
	{"TJS", "972", "2", "Tajikistani somoni"},
	{"TMT", "934", "2", "Turkmenistani manat"},
	{"TND", "788", "3", "Tunisian dinar"},
	{"TOP", "776", "2", "Tongan paʻanga"},
	{"TRY", "949", "2", "Turkish lira"},
	{"TTD", "780", "2", "Trinidad and Tobago dollar"},
	{"TWD", "901", "2", "New Taiwan dollar"},
	{"TZS", "834", "2", "Tanzanian shilling"},
	{"UAH", "980", "2", "Ukrainian hryvnia"},
	{"UGX", "800", "2", "Ugandan shilling"},
	{"USD", "840", "2", "United States dollar"},
	{"USN", "997", "2", "United States dollar (next day) (funds code)"},
	{"USS", "998", "2", "United States dollar (same day) (funds code) (one source[who?] claims it is no longer used, but it is still on the ISO 4217-MA list)"},
	{"UYI", "940", "0", "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)"},
	{"UYU", "858", "2", "Uruguayan peso"},
	{"UZS", "860", "2", "Uzbekistan som"},
	{"VEF", "937", "2", "Venezuelan bolívar fuerte"},
	{"VND", "704", "0", "Vietnamese dong"},
	{"VUV", "548", "0", "Vanuatu vatu"},
	{"WST", "882", "2", "Samoan tala"},
	{"XAF", "950", "0", "CFA franc BEAC"},
	{"XAG", "961", ".", "Silver (one troy ounce)"},
	{"XAU", "959", ".", "Gold (one troy ounce)"},
	{"XBA", "955", ".", "European Composite Unit (EURCO) (bond market unit)"},
	{"XBB", "956", ".", "European Monetary Unit (E.M.U.-6) (bond market unit)"},
	{"XBC", "957", ".", "European Unit of Account 9 (E.U.A.-9) (bond market unit)"},
	{"XBD", "958", ".", "European Unit of Account 17 (E.U.A.-17) (bond market unit)"},
	{"XCD", "951", "2", "East Caribbean dollar"},
	{"XDR", "960", ".", "Special drawing rights"},
	{"XFU", "Nil", ".", "UIC franc (special settlement currency)"},
	{"XOF", "952", "0", "CFA franc BCEAO"},
	{"XPD", "964", ".", "Palladium (one troy ounce)"},
	{"XPF", "953", "0", "CFP franc"},
	{"XPT", "962", ".", "Platinum (one troy ounce)"},
	{"XTS", "963", ".", "Code reserved for testing purposes"},
	{"XXX", "999", ".", "No currency"},
	{"YER", "886", "2", "Yemeni rial"},
	{"ZAR", "710", "2", "South African rand"},
	{"ZMW", "967", "2", "Zambian kwacha"},
}
