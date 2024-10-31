package models

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar" json:"name"`
}

type SubCategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	CategoryID uint   `gorm:"foreignKey:CategoryID" json:"category_id"`
}

var CategorySeedData = []Category{
	{ID: 1, Name: "Cărți"},
	{ID: 2, Name: "Cadouri. Totul pentru sarbatori"},
	{ID: 3, Name: "Fructe, fructe de padure & Legume"},
	{ID: 4, Name: "Culinărie (posibil și prin precomandă)"},
	{ID: 5, Name: "Produse Panificație (pâine, croissante, covrigi)"},
	{ID: 6, Name: "Mezeluri și crenvurști"},
	{ID: 7, Name: "Produse lactate"},
	{ID: 8, Name: "Dulciuri"},
	{ID: 9, Name: "Carne"},
	{ID: 10, Name: "Ouă"},
	{ID: 11, Name: "Cașcaval"},
	{ID: 12, Name: "Crupe și boboase"},
	{ID: 13, Name: "Băcănie"},
	{ID: 14, Name: "Conserve"},
	{ID: 15, Name: "Peşte"},
	{ID: 16, Name: "Produse de cofetărie"},
	{ID: 17, Name: "Ceai și cafea"},
	{ID: 18, Name: "Produse congelate"},
	{ID: 19, Name: "Nuci, fructe uscate și semințe"},
	{ID: 20, Name: "Snack-uri"},
	{ID: 21, Name: "Băuturi nealcoolice"},
	{ID: 22, Name: "Băuturi alcoolice"},
	{ID: 23, Name: "Produse chimice de uz casnic"},
	{ID: 24, Name: "Produse cosmetice"},
	{ID: 25, Name: "Igienă și îngrijire"},
	{ID: 26, Name: "Lumea copiilor"},
	{ID: 27, Name: "Papetărie"},
	{ID: 28, Name: "Hrană & Accesorii animale"},
	{ID: 29, Name: "Totul pentru CASA MODERNĂ"},
	{ID: 30, Name: "Bucătărie"},
	{ID: 31, Name: "Bunuri gospodărești"},
	{ID: 32, Name: "Îmbrăcăminte. Încălțăminte. Accesorii. Textile"},
	{ID: 33, Name: "Totul pentru masina"},
	{ID: 34, Name: "Electrocasnice. Iluminat"},
	{ID: 35, Name: "Tehnica Audio-Video"},
	{ID: 36, Name: "Vacanță. Picnic"},
	{ID: 37, Name: "Sport"},
}
var SubCategorySeedData = []SubCategory{
	// Subcategories for "Cărți"
	{ID: 1, Name: "Literatură pentru copii/adolescenți", CategoryID: 1},
	{ID: 2, Name: "Romane", CategoryID: 1},
	{ID: 3, Name: "Ficțiune", CategoryID: 1},
	{ID: 4, Name: "Literatură motivațională", CategoryID: 1},
	{ID: 5, Name: "Literatură de dezvoltare personală", CategoryID: 1},
	{ID: 6, Name: "Filosofie, istorie.", CategoryID: 1},
	{ID: 7, Name: "Memorii, jurnale, biografii.", CategoryID: 1},

	// Subcategories for "Cadouri. Totul pentru sarbatori"
	{ID: 8, Name: "Accesorii pentru sarbatori", CategoryID: 2},
	{ID: 9, Name: "Idei de cadouri", CategoryID: 2},
	{ID: 10, Name: "Ladițe decorative, cutii, pungi cadou", CategoryID: 2},

	// Subcategories for "Fructe, fructe de padure & Legume"
	{ID: 11, Name: "Fructe, fructe de padure", CategoryID: 3},
	{ID: 12, Name: "Legume", CategoryID: 3},
	{ID: 13, Name: "Salate & Verdețuri", CategoryID: 3},

	// Subcategories for "Culinărie (posibil și prin precomandă)"
	{ID: 14, Name: "Fel principal", CategoryID: 4},
	{ID: 15, Name: "Salate", CategoryID: 4},
	{ID: 16, Name: "Fast food", CategoryID: 4},
	{ID: 17, Name: "Placinte, plăcinte, pâine plate, vertutas", CategoryID: 4},
	{ID: 18, Name: "Deserturi", CategoryID: 4},

	// Subcategories for "Produse Panificație (pâine, croissante, covrigi)"
	{ID: 19, Name: "Pâine", CategoryID: 5},
	{ID: 20, Name: "Pâine uscată & expandată", CategoryID: 5},
	{ID: 21, Name: "Colaci", CategoryID: 5},
	{ID: 22, Name: "Lavaș", CategoryID: 5},
	{ID: 23, Name: "Chifle și croissante", CategoryID: 5},
	{ID: 24, Name: "Covrigi", CategoryID: 5},
	{ID: 25, Name: "Gogoși", CategoryID: 5},

	// Subcategories for "Mezeluri și crenvurști"
	{ID: 26, Name: "Parizer", CategoryID: 6},
	{ID: 27, Name: "Crenvurști & Safalade", CategoryID: 6},
	{ID: 28, Name: "Afumături", CategoryID: 6},
	{ID: 29, Name: "Şuncă", CategoryID: 6},
	{ID: 30, Name: "Salamuri crud-uscate", CategoryID: 6},
	{ID: 31, Name: "Salamuri fiert-afumate", CategoryID: 6},
	{ID: 32, Name: "Tobă, rulade", CategoryID: 6},

	// Subcategories for "Produse lactate"
	{ID: 33, Name: "Lapte", CategoryID: 7},
	{ID: 34, Name: "Chefir", CategoryID: 7},
	{ID: 35, Name: "Iaurturi", CategoryID: 7},
	{ID: 36, Name: "Smântână", CategoryID: 7},
	{ID: 37, Name: "Brânză de vaci, brânză feta", CategoryID: 7},
	{ID: 38, Name: "Frișcă", CategoryID: 7},
	{ID: 39, Name: "Lapte condensat", CategoryID: 7},
	{ID: 40, Name: "Deserturi", CategoryID: 7},
	{ID: 41, Name: "Unt", CategoryID: 7},
	{ID: 42, Name: "Margarină", CategoryID: 7},

	// Subcategories for "Dulciuri"
	{ID: 43, Name: "Bomboane de ciocolată & Praline", CategoryID: 8},
	{ID: 44, Name: "Ciocolate în cutie", CategoryID: 8},
	{ID: 45, Name: "Ciocolată tablete", CategoryID: 8},
	{ID: 46, Name: "Batoane de ciocolată, batoane de muesli", CategoryID: 8},
	{ID: 47, Name: "Caramele & drajeuri", CategoryID: 8},
	{ID: 48, Name: "Gume de mestecat & bomboane gumate", CategoryID: 8},
	{ID: 49, Name: "Biscuiți", CategoryID: 8},
	{ID: 50, Name: "Turte dulci", CategoryID: 8},
	{ID: 51, Name: "Napolitane", CategoryID: 8},
	{ID: 52, Name: "Rulade, muffin, chec, panettone", CategoryID: 8},
	{ID: 53, Name: "Blaturi tort", CategoryID: 8},
	{ID: 54, Name: "Alte dulciuri", CategoryID: 8},
	{ID: 55, Name: "Produse pentru diabetici", CategoryID: 8},
	{ID: 56, Name: "Cremă de ciocolată", CategoryID: 8},

	// Subcategories for "Carne"
	{ID: 57, Name: "Carne proaspătă", CategoryID: 9},
	{ID: 58, Name: "Carne tocată", CategoryID: 9},
	{ID: 59, Name: "Marinate", CategoryID: 9},
	{ID: 60, Name: "Cârnăciori și mici", CategoryID: 9},

	// Subcategories for "Ouă"
	{ID: 61, Name: "Ouă", CategoryID: 10},

	// Subcategories for "Cașcaval"
	{ID: 62, Name: "Brânză tare, semi-tare", CategoryID: 11},
	{ID: 63, Name: "Mozzarela", CategoryID: 11},
	{ID: 64, Name: "Cașcaval cu mucegai", CategoryID: 11},
	{ID: 65, Name: "Tartina de branza", CategoryID: 11},
	{ID: 66, Name: "Brânză, moale, procesată, portionat", CategoryID: 11},

	// Subcategories for "Crupe și boboase"
	{ID: 67, Name: "Orez", CategoryID: 12},
	{ID: 68, Name: "Hrișcă", CategoryID: 12},
	{ID: 69, Name: "Bulgur", CategoryID: 12},
	{ID: 70, Name: "Arpacaș", CategoryID: 12},
	{ID: 71, Name: "Mei", CategoryID: 12},
	{ID: 72, Name: "Crupe de gris", CategoryID: 12},
	{ID: 73, Name: "Crupe de arnaut", CategoryID: 12},
	{ID: 74, Name: "Mălai", CategoryID: 12},
	{ID: 75, Name: "Couscous", CategoryID: 12},
	{ID: 76, Name: "Grîu", CategoryID: 12},
	{ID: 77, Name: "Crupă de orz", CategoryID: 12},
	{ID: 78, Name: "Mazăre", CategoryID: 12},
	{ID: 79, Name: "Linte", CategoryID: 12},
	{ID: 80, Name: "Năut", CategoryID: 12},
	{ID: 81, Name: "Fasole", CategoryID: 12},
	{ID: 82, Name: "Amestec", CategoryID: 12},

	// Subcategories for "Băcănie"
	{ID: 83, Name: "Bucătăria orientală. Sushi", CategoryID: 13},
	{ID: 84, Name: "Zahăr", CategoryID: 13},
	{ID: 85, Name: "Sare", CategoryID: 13},
	{ID: 86, Name: "Paste", CategoryID: 13},
	{ID: 87, Name: "Făină, pesmet", CategoryID: 13},
	{ID: 88, Name: "Cereale pentru mic dejun", CategoryID: 13},
	{ID: 89, Name: "Ulei", CategoryID: 13},
	{ID: 90, Name: "Maioneză, ketchup", CategoryID: 13},
	{ID: 91, Name: "Sosuri & Dressing", CategoryID: 13},
	{ID: 92, Name: "Murături", CategoryID: 13},
	{ID: 93, Name: "Borş acru", CategoryID: 13},
	{ID: 94, Name: "Oțet", CategoryID: 13},
	{ID: 95, Name: "Alimente instant", CategoryID: 13},
	{ID: 96, Name: "Condimente și Mirodenii", CategoryID: 13},
	{ID: 97, Name: "Articole pentru copt și deserturi", CategoryID: 13},
	{ID: 98, Name: "Jeleu, Kissel", CategoryID: 13},

	// Subcategories for "Conserve"
	{ID: 99, Name: "Conserve din carne", CategoryID: 14},
	{ID: 100, Name: "Conserve de pește", CategoryID: 14},
	{ID: 101, Name: "Conserve din legume", CategoryID: 14},
	{ID: 102, Name: "Măsline", CategoryID: 14},
	{ID: 103, Name: "Pateuri vegetale", CategoryID: 14},
	{ID: 104, Name: "Ciuperci", CategoryID: 14},
	{ID: 105, Name: "Conserve din fructe", CategoryID: 14},
	{ID: 106, Name: "Miere", CategoryID: 14},

	// Subcategories for "Peşte"
	{ID: 107, Name: "Pește proaspăt", CategoryID: 15},
	{ID: 108, Name: "Fructe de mare", CategoryID: 15},
	{ID: 109, Name: "Pește afumat", CategoryID: 15},
	{ID: 110, Name: "Pește sărat", CategoryID: 15},
	{ID: 111, Name: "Pește uscat", CategoryID: 15},
	{ID: 112, Name: "Icre de pește și preparate din icre", CategoryID: 15},

	// Subcategories for "Produse de cofetărie"
	{ID: 113, Name: "Torturi", CategoryID: 16},
	{ID: 114, Name: "Prăjituri", CategoryID: 16},

	// Subcategories for "Ceai și cafea"
	{ID: 115, Name: "Ceai pachețele", CategoryID: 17},
	{ID: 116, Name: "Ceai infuzie", CategoryID: 17},
	{ID: 117, Name: "Cafea solubilă", CategoryID: 17},
	{ID: 118, Name: "Cafea măcinată", CategoryID: 17},
	{ID: 119, Name: "Cafea boabe", CategoryID: 17},
	{ID: 120, Name: "Cafea în capsule", CategoryID: 17},
	{ID: 121, Name: "Cappucinno, mixuri", CategoryID: 17},
	{ID: 122, Name: "Cacao, ciocolata caldă", CategoryID: 17},
	{ID: 123, Name: "Cicoare", CategoryID: 17},

	// Subcategories for "Produse congelate"
	{ID: 124, Name: "Congelate din pește", CategoryID: 18},
	{ID: 125, Name: "Congelate din carne", CategoryID: 18},
	{ID: 126, Name: "Legume congelate, verdeturi", CategoryID: 18},
	{ID: 127, Name: "Fructe congelate", CategoryID: 18},
	{ID: 128, Name: "Aluat congelat", CategoryID: 18},
	{ID: 129, Name: "Pizza & patiserie congelate", CategoryID: 18},
	{ID: 130, Name: "Pelmeni & Colțunași", CategoryID: 18},
	{ID: 131, Name: "Înghețată", CategoryID: 18},
	{ID: 132, Name: "Alte produse congelate", CategoryID: 18},

	// Subcategories for "Nuci, fructe uscate și semințe"
	{ID: 133, Name: "Fructe uscate", CategoryID: 19},
	{ID: 134, Name: "Nuci", CategoryID: 19},
	{ID: 135, Name: "Semințe", CategoryID: 19},

	// Subcategories for "Snack-uri"
	{ID: 136, Name: "Chipsuri", CategoryID: 20},
	{ID: 137, Name: "Sticks, crackers & snack", CategoryID: 20},
	{ID: 138, Name: "Pesmeţi", CategoryID: 20},
	{ID: 139, Name: "Popcorn", CategoryID: 20},
	{ID: 140, Name: "Semințe floarea soarelui & dovleac", CategoryID: 20},
	{ID: 141, Name: "Arahide, fistic, alte gustări", CategoryID: 20},
	{ID: 142, Name: "Gustări de pește și carne", CategoryID: 20},

	// Subcategories for "Băuturi nealcoolice"
	{ID: 143, Name: "Apă minerală", CategoryID: 21},
	{ID: 144, Name: "Băuturi răcoritoare", CategoryID: 21},
	{ID: 145, Name: "Suc & Nectar", CategoryID: 21},
	{ID: 146, Name: "Energizante", CategoryID: 21},

	// Subcategories for "Băuturi alcoolice"
	{ID: 147, Name: "Vin", CategoryID: 22},
	{ID: 148, Name: "Vin spumant", CategoryID: 22},
	{ID: 149, Name: "Divin", CategoryID: 22},
	{ID: 150, Name: "Vodcă", CategoryID: 22},
	{ID: 151, Name: "Whiskey", CategoryID: 22},
	{ID: 152, Name: "Rom. Tequila. Gin. Brandy", CategoryID: 22},
	{ID: 153, Name: "Lichior. Balsam. Vermut. Aperol", CategoryID: 22},
	{ID: 154, Name: "Bere", CategoryID: 22},
	{ID: 155, Name: "Băuturi slab alcoolice", CategoryID: 22},

	// Subcategories for "Produse chimice de uz casnic"
	{ID: 156, Name: "Detergenți de vase", CategoryID: 23},
	{ID: 157, Name: "Detergenți pentru rufe", CategoryID: 23},
	{ID: 158, Name: "Articole pentru curățenia suprafețelor", CategoryID: 23},
	{ID: 159, Name: "Produse pentru mașina de spălat vase/mașina de spălat", CategoryID: 23},
	{ID: 160, Name: "Repelente pentru insecte", CategoryID: 23},
	{ID: 161, Name: "Odorizanți", CategoryID: 23},

	// Subcategories for "Produse cosmetice"
	{ID: 162, Name: "Parfumerie", CategoryID: 24},
	{ID: 163, Name: "Machiaj", CategoryID: 24},
	{ID: 164, Name: "Creme", CategoryID: 24},
	{ID: 165, Name: "Măști cosmetice și patches", CategoryID: 24},
	{ID: 166, Name: "Produse de curățare și demachiere", CategoryID: 24},
	{ID: 167, Name: "Vopsea,tonice", CategoryID: 24},
	{ID: 168, Name: "Seria solară", CategoryID: 24},

	// Subcategories for "Igienă și îngrijire"
	{ID: 169, Name: "Săpun", CategoryID: 25},
	{ID: 170, Name: "Îngrijire  corp", CategoryID: 25},
	{ID: 171, Name: "Îngrijire păr", CategoryID: 25},
	{ID: 172, Name: "Igiena orală", CategoryID: 25},
	{ID: 173, Name: "Igiena intimă", CategoryID: 25},
	{ID: 174, Name: "Produse din bumbac", CategoryID: 25},
	{ID: 175, Name: "Igiena si cosmetica barbatilor", CategoryID: 25},
	{ID: 176, Name: "Servetele umede", CategoryID: 25},
	{ID: 177, Name: "Trusă de prim ajutor", CategoryID: 25},

	// Subcategories for "Lumea copiilor"
	{ID: 178, Name: "Alimentația copiilor", CategoryID: 26},
	{ID: 179, Name: "Produse cosmetice. Igienă. Protecție", CategoryID: 26},
	{ID: 180, Name: "Scutece și șervețele umede", CategoryID: 26},
	{ID: 181, Name: "Produse chimice de uz casnic", CategoryID: 26},
	{ID: 182, Name: "Accesorii pentru copii", CategoryID: 26},
	{ID: 183, Name: "Produse pentru mămici", CategoryID: 26},
	{ID: 184, Name: "Jucării", CategoryID: 26},
	{ID: 185, Name: "Cărți educative", CategoryID: 26},

	// Subcategories for "Papetărie"
	{ID: 186, Name: "Caiete, blocnotes, agende", CategoryID: 27},
	{ID: 187, Name: "Totul pentru desen, creativitate", CategoryID: 27},
	{ID: 188, Name: "Rechizite școlare", CategoryID: 27},

	// Subcategories for "Hrană & Accesorii animale"
	{ID: 189, Name: "Hrană pisici", CategoryID: 28},
	{ID: 190, Name: "Hrană câini", CategoryID: 28},
	{ID: 191, Name: "Alte Produse pentru animale", CategoryID: 28},
	{ID: 192, Name: "Asternut Igienic. Scutec", CategoryID: 28},

	// Subcategories for "Totul pentru CASA MODERNĂ"
	{ID: 193, Name: "Depozitare și organizarea spațiului", CategoryID: 29},
	{ID: 194, Name: "Baie", CategoryID: 29},
	{ID: 195, Name: "Decor și accesorii pentru casă", CategoryID: 29},
	{ID: 196, Name: "Lumânări", CategoryID: 29},

	// Subcategories for "Bucătărie"
	{ID: 197, Name: "Vesela pentru gatit", CategoryID: 30},
	{ID: 198, Name: "Vesela de masa", CategoryID: 30},
	{ID: 199, Name: "Vesela pentru copii", CategoryID: 30},
	{ID: 200, Name: "Accesorii pentru bucatarie", CategoryID: 30},
	{ID: 201, Name: "Vesela de unica folosinta", CategoryID: 30},
	{ID: 202, Name: "Depozitarea și organizarea spațiului", CategoryID: 30},
	{ID: 203, Name: "Termosuri, căni termice", CategoryID: 30},

	// Subcategories for "Bunuri gospodărești"
	{ID: 204, Name: "Produse din hartie", CategoryID: 31},
	{ID: 205, Name: "Folie, folie alimentară, hârtie de copt", CategoryID: 31},
	{ID: 206, Name: "Totul pentru curatenie in casa", CategoryID: 31},
	{ID: 207, Name: "Inventar curatenie", CategoryID: 31},
	{ID: 208, Name: "Mese de calcat, uscatoare de rufe, scari", CategoryID: 31},
	{ID: 209, Name: "Unelte", CategoryID: 31},
	{ID: 210, Name: "Saci menajeri", CategoryID: 31},

	// Subcategories for "Îmbrăcăminte. Încălțăminte. Accesorii. Textile"
	{ID: 211, Name: "Îmbrăcăminte", CategoryID: 32},
	{ID: 212, Name: "Incaltaminte", CategoryID: 32},
	{ID: 213, Name: "Colanți și șosete", CategoryID: 32},
	{ID: 214, Name: "Galanterie", CategoryID: 32},
	{ID: 215, Name: "Textile", CategoryID: 32},
	{ID: 216, Name: "Îngrijire și depozitare", CategoryID: 32},

	// Subcategories for "Totul pentru masina"
	{ID: 217, Name: "Curatenie auto", CategoryID: 33},
	{ID: 218, Name: "Odorizante auto", CategoryID: 33},
	{ID: 219, Name: "Accesorii pentru mașină", CategoryID: 33},

	// Subcategories for "Electrocasnice. Iluminat"
	{ID: 220, Name: "Tehnica de bucatarie", CategoryID: 34},
	{ID: 221, Name: "Tehnica pentru casa", CategoryID: 34},
	{ID: 222, Name: "Tehnică pentru frumusețe", CategoryID: 34},
	{ID: 223, Name: "Produse Electrice", CategoryID: 34},
	{ID: 224, Name: "Lămpi de masă. Noptiere, decor iluminat", CategoryID: 34},
	{ID: 225, Name: "Lanterne", CategoryID: 34},
	{ID: 226, Name: "Baterii", CategoryID: 34},

	// Subcategories for "Tehnica Audio-Video"
	{ID: 227, Name: "Căști", CategoryID: 35},
	{ID: 228, Name: "Bluetooth boxe, radio", CategoryID: 35},
	{ID: 229, Name: "Accesorii pentru echipamente audio-video", CategoryID: 35},

	// Subcategories for "Vacanță. Picnic"
	{ID: 230, Name: "Picnic", CategoryID: 36},
	{ID: 231, Name: "Vacanță pe plajă", CategoryID: 36},

	// Subcategories for "Sport"
	{ID: 232, Name: "Sport", CategoryID: 37},
}
