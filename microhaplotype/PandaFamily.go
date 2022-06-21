package microhaplotype

type Pedigree int

var SampleID = map[Pedigree]string{
	278:  "FDSW210021731-1r",
	401:  "FDSW210021757-1b",
	598:  "FDSW210021244-1r",
	678:  "FDSW202100961-1r",
	709:  "FDSW210021729-1r",
	724:  "FDSW210021763-1r",
	761:  "FDSW210021727-1b",
	788:  "FDSW210021702-1r",
	831:  "FDSW210021748-1r",
	853:  "FDSW210021764-1r",
	857:  "FDSW202100959-1r", // 857:"FDSW210021242-1r",
	858:  "FDSW202100958-1r", // 858:"FDSW210021243-1r",
	894:  "FDSW210021738-1r",
	895:  "FDSW210021736-1r",
	954:  "FDSW210021737-1r",
	958:  "FDSW210021768-1r", // 958:"FDSW210021768-1r",
	963:  "FDSW210021716-1r",
	994:  "FDSW210021746-1r",
	1013: "FDSW210021755-1r",
	1016: "FDSW210021749-1r",
	1017: "FDSW210021750-1r",
	1032: "FDSW210021704-1r",
	1033: "FDSW210021705-1r",
	1038: "FDSW210021747-1r",
	1081: "FDSW210021743-1r",
	1082: "FDSW210021742-1r",
	1122: "FDSW210021753-1r",
	1124: "FDSW210021759-1r",
	1125: "FDSW210021760-1r",
	1129: "FDSW210021754-1r",
	1224: "FDSW202143895-1r",
	555:  "FDSW210021718-1r",
	561:  "FDSW210021724-1r",
	570:  "FDSW210021701-1r",
	635:  "FDSW210021715-1r",
	731:  "FDSW210021717-1r",
	796:  "FDSW210021721-1b",
	725:  "FDSW202100960-1r", // 725:"FDSW210021238-1r",
	342:  "FDSW210021722-1r",
	540:  "FDSW210021756-1r",
	649:  "FDSW210021706-1r",
	652:  "FDSW210021761-1r",
	523:  "FDSW210021700-1b",
	584:  "FDSW210021712-1r",
	532:  "FDSW210021708-1r",
}

var Family = [][3]Pedigree{ // {Father, Mother, Child}
	{342, 278, 401},
	{342, 523, 724},
	{342, 523, 725},
	{532, 523, 857},
	{532, 523, 858},
	{532, 523, 1038},
	{540, 555, 1032},
	{540, 555, 1033},
	{342, 561, 788},
	{584, 561, 853},
	{540, 570, 1017},
	{584, 570, 1016},
	{831, 598, 963},
	{540, 635, 1224},
	{584, 635, 958},
	{584, 652, 894},
	{584, 652, 895},
	{649, 652, 954},
	{649, 652, 1122},
	{532, 678, 1129},
	{649, 709, 1013},
	{731, 725, 1081},
	{731, 725, 1082},
	{649, 761, 994},
	{532, 796, 1124},
	{532, 796, 1125},
}