package handlers

type Item struct {
	Id    int
	Title string
	Price string
	Count string
}

var Items = []Item{{1, "Dýňovka", "70", "0"},
	{2, "Zelňčka", "70", "0"},
	{3, "Párek v rohliku", "40", "0"},
	{4, "Toast", "60", "0"},
	{5, "Videňské párky", "60", "0"},
	{6, "Tortila", "180", "0"},
	{7, "Kuřecí stehno", "200", "0"},
	{8, "Halušky", "180", "0"},
	{9, "Kuřecí plátek žampiony", "180", "0"},
	{10, "Buchta", "40", "0"},
	{11, "Muffin", "50", "0"},
	{12, "Oplatek", "30", "0"},
	{13, "Kofola 0.5", "40", "0"},
	{14, "Kofola 0.3", "30", "0"},
	{15, "Voda", "30", "0"},
	{16, "Juice 0.3", "30", "0"},
	{17, "Ledový čaj", "30", "0"},
	{28, "Ovocné Pivo", "30", "0"},
	{18, "Birel", "30", "0"},
	{19, "Pivo", "30", "0"},
	{20, "Jagermaister", "50", "0"},
	{21, "Rum", "50", "0"},
	{22, "Slivovice", "40", "0"},
	{23, "Becher", "40", "0"},
	{24, "Fernet", "40", "0"},
	{25, "Víno 0.1", "30", "0"},
	{26, "Presso", "40", "0"},
	{27, "Rozpustná/Turek", "30", "0"},
	{29, "Čaj", "30", "0"}}

type Order struct {
	Quantity int
	Product  string
	Price    int
}
