package ecotouch

type tagDefinition struct {
	name   string
	module string
	fact   float64
	usage  int
}

var tags = map[string]tagDefinition{
	"A1": {
		name:   "temp",
		module: "outside",
		fact:   0.100000,
		usage:  1,
	},
	"A2": {
		name:   "temp_1h",
		module: "outside",
		fact:   0.100000,
		usage:  2,
	},
	"A3": {
		name:   "temp_24h",
		module: "outside",
		fact:   0.100000,
		usage:  2,
	},
	"A4": {
		name:   "temp_source_in",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A5": {
		name:   "temp_source_out",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A6": {
		name:   "temp_evap",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A7": {
		name:   "temp_suction",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A8": {
		name:   "pressure_evap",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A10": {
		name:   "temp_return_set",
		module: "main",
		fact:   0.100000,
		usage:  3,
	},
	"A11": {
		name:   "temp_return",
		module: "main",
		fact:   0.100000,
		usage:  3,
	},
	"A12": {
		name:   "temp_flow",
		module: "main",
		fact:   0.100000,
		usage:  3,
	},
	"A14": {
		name:   "temp_condens",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A15": {
		name:   "pressure_condens",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A16": {
		name:   "temp",
		module: "storage",
		fact:   0.100000,
		usage:  1,
	},
	"A17": {
		name:   "temp",
		module: "roomsensor",
		fact:   0.100000,
		usage:  1,
	},
	"A18": {
		name:   "temp_1h",
		module: "roomsensor",
		fact:   0.100000,
		usage:  2,
	},
	"A19": {
		name:   "temp",
		module: "water",
		fact:   0.100000,
		usage:  1,
	},
	"A20": {
		name:   "temp",
		module: "pool",
		fact:   0.100000,
		usage:  1,
	},
	"A21": {
		name:   "temp",
		module: "solar",
		fact:   0.100000,
		usage:  1,
	},
	"A22": {
		name:   "temp_flow",
		module: "solar",
		fact:   0.100000,
		usage:  3,
	},
	"A23": {
		name:   "position_exp_valve",
		module: "geo",
		fact:   0.100000,
		usage:  3,
	},
	"A25": {
		name:   "power_comp_el",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A26": {
		name:   "power_heating",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A27": {
		name:   "power_cooling",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A28": {
		name:   "COP_heating",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A29": {
		name:   "COP_cooling",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A30": {
		name:   "temp_return",
		module: "heating",
		fact:   0.100000,
		usage:  3,
	},
	"A31": {
		name:   "temp_set",
		module: "heating",
		fact:   0.100000,
		usage:  2,
	},
	"A32": {
		name:   "temp_set2",
		module: "heating",
		fact:   0.100000,
		usage:  2,
	},
	"A33": {
		name:   "temp_return",
		module: "cooling",
		fact:   0.100000,
		usage:  3,
	},
	"A34": {
		name:   "temp_set",
		module: "cooling",
		fact:   0.100000,
		usage:  2,
	},
	"A35": {
		name:   "temp_set2",
		module: "cooling",
		fact:   0.100000,
		usage:  2,
	},
	"A37": {
		name:   "temp_set",
		module: "water",
		fact:   0.100000,
		usage:  2,
	},
	"A38": {
		name:   "temp_set2",
		module: "water",
		fact:   0.100000,
		usage:  2,
	},
	"A40": {
		name:   "temp_set",
		module: "pool",
		fact:   0.100000,
		usage:  2,
	},
	"A41": {
		name:   "temp_set2",
		module: "pool",
		fact:   0.100000,
		usage:  2,
	},
	"A50": {
		name:   "comp_power_set",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"A51": {
		name:   "valve_percent",
		module: "comp1",
		fact:   0.100000,
		usage:  0,
	},
	"A56": {
		name:   "valve_set",
		module: "comp1",
		fact:   0.100000,
		usage:  0,
	},
	"A57": {
		name:   "valve_set2",
		module: "comp1",
		fact:   0.100000,
		usage:  0,
	},
	"A58": {
		name:   "power_percent",
		module: "comp1",
		fact:   0.100000,
		usage:  2,
	},
	"A109": {
		name:   "temp_cooling",
		module: "cooling",
		fact:   0.100000,
		usage:  3,
	},
	"A444": {
		name:   "tag_A444",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A445": {
		name:   "tag_A445",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A446": {
		name:   "tag_A446",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A447": {
		name:   "tag_A447",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A448": {
		name:   "tag_A448",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A449": {
		name:   "tag_A449",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A450": {
		name:   "tag_A450",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A451": {
		name:   "tag_A451",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A452": {
		name:   "tag_A452",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A453": {
		name:   "tag_A453",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A454": {
		name:   "tag_A454",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A455": {
		name:   "tag_A455",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A456": {
		name:   "tag_A456",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A457": {
		name:   "tag_A457",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A722": {
		name:   "tag_A722",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"A781": {
		name:   "tag_A781",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I1": {
		name:   "tag_I1",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I2": {
		name:   "tag_I2",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I3": {
		name:   "tag_I3",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I4": {
		name:   "tag_I4",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I5": {
		name:   "date_day",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I6": {
		name:   "date_month",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I7": {
		name:   "date_year",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I8": {
		name:   "time_hour",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I9": {
		name:   "time_minute",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I10": {
		name:   "OpHrs",
		module: "comp1",
		fact:   1.000000,
		usage:  3,
	},
	"I11": {
		name:   "OpHrs_tenth",
		module: "comp1",
		fact:   0.100000,
		usage:  3,
	},
	"I12": {
		name:   "tag_I12 Op?",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I13": {
		name:   "tag_I13 Op?",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I14": {
		name:   "OpHrs",
		module: "comp2",
		fact:   1.000000,
		usage:  3,
	},
	"I15": {
		name:   "tag_I15 Op?",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I16": {
		name:   "tag_I16 Op?",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I17": {
		name:   "tag_I17 Op?",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I18": {
		name:   "OpHrs_heatpump",
		module: "main",
		fact:   1.000000,
		usage:  3,
	},
	"I19": {
		name:   "OpHrs_heatpump_tenth",
		module: "main",
		fact:   0.100000,
		usage:  3,
	},
	"I20": {
		name:   "OpHrs",
		module: "geo",
		fact:   1.000000,
		usage:  3,
	},
	"I22": {
		name:   "OpHrs",
		module: "solar",
		fact:   1.000000,
		usage:  3,
	},
	"I30": {
		name:   "set_auto",
		module: "heating",
		fact:   1.000000,
		usage:  2,
	},
	"I31": {
		name:   "set_auto",
		module: "cooling",
		fact:   1.000000,
		usage:  2,
	},
	"I32": {
		name:   "set_auto",
		module: "water",
		fact:   1.000000,
		usage:  2,
	},
	"I33": {
		name:   "tag_I33",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I34": {
		name:   "tag_I34",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I35": {
		name:   "tag_I35",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I36": {
		name:   "tag_I36",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I37": {
		name:   "tag_I37",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I38": {
		name:   "tag_I38",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I39": {
		name:   "tag_I39",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I40": {
		name:   "tag_I40",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I41": {
		name:   "tag_I41",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I42": {
		name:   "tag_I42",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I51": {
		name:   "state",
		module: "main",
		fact:   1.000000,
		usage:  1,
	},
	"I52": {
		name:   "state_failure",
		module: "main",
		fact:   1.000000,
		usage:  0,
	},
	"I53": {
		name:   "state_interruption",
		module: "main",
		fact:   1.000000,
		usage:  1,
	},
	"I115": {
		name:   "state_?",
		module: "comp1",
		fact:   1.000000,
		usage:  3,
	},
	"I136": {
		name:   "state_opmode",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"I137": {
		name:   "state",
		module: "heating",
		fact:   1.000000,
		usage:  1,
	},
	"I138": {
		name:   "state",
		module: "cooling",
		fact:   1.000000,
		usage:  3,
	},
	"I139": {
		name:   "state",
		module: "water",
		fact:   1.000000,
		usage:  3,
	},
	"I263": {
		name:   "temp_set_offset",
		module: "heating",
		fact:   1.000000,
		usage:  2,
	},
	"I1458": {
		name:   "state",
		module: "comp1",
		fact:   1.000000,
		usage:  1,
	},
	"I1459": {
		name:   "tag_I1459",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"I1460": {
		name:   "tag_I1460",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"D116": {
		name:   "tag_D116",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
	"D420": {
		name:   "state_holiday",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"D671": {
		name:   "state_alarm",
		module: "main",
		fact:   1.000000,
		usage:  2,
	},
	"D696": {
		name:   "tag_D696",
		module: "unknown",
		fact:   1.000000,
		usage:  3,
	},
}

type stateDefinition struct {
	flag   int
	name   string
	module string
	usage  int
}

var stateWord = []stateDefinition{
	{flag: 1, name: "state_sourcepump", module: "main", usage: 1},         // 'Status Quellenpumpe'
	{flag: 2, name: "state_heatpump", module: "main", usage: 1},          // 'Status Heizungsumwälzpumpe'
	{flag: 4, name: "state_evd", module: "main", usage: 1},                // 'Status Freigabe Regelung EVD/Magnetventil'
	{flag: 8, name: "state_comp1", module: "main", usage: 1},            // 'Status Verdichter 1'
	{flag: 16, name: "state_comp2", module: "main", usage: 1},           // 'Status Verdichter 2'
	{flag: 32, name: "state_extheater", module: "main", usage: 1},        // 'Status externer Wärmeerzeuger'
	{flag: 64, name: "state_alarm", module: "main", usage: 1},            // 'Status Alarmausgang'
	{flag: 128, name: "state_cooling", module: "main", usage: 1},         // 'Status Motorventil Kühlbetrieb'
	{flag: 256, name: "state_water", module: "main", usage: 1},          // 'Status Motorventil Warmwasser'
	{flag: 512, name: "state_pool", module: "main", usage: 1},            // 'Status Motorventil Pool'
	{flag: 1024, name: "state_solar", module: "main", usage: 1},         // 'Status Solarbetrieb'
	{flag: 2048, name: "state_cooling4way", module: "main", usage: 1}, // 'Status 4-Wegeventil im Kältekreis'
}

var tagComments = map[string]string{
	"A1":    "'Aussentemperatur'",
	"A2":    "'Aussentemperatur Mittelwert 1h'",
	"A3":    "'Außentemperatur Mittelwert 24h'",
	"A4":    "'Quelleneintrittstemperatur'",
	"A5":    "'Quellenaustrittstemperatur'",
	"A6":    "'Verdampfungstemperatur'",
	"A7":    "'Sauggastemperatur'",
	"A8":    "'Verdampfungsdruck'",
	"A10":   "'Temperatur Rücklauf Soll'",
	"A11":   "'Temperatur Rücklauf'",
	"A12":   "'Temperatur Vorlauf'",
	"A14":   "'Kondensationstemperatur'",
	"A15":   "'Kondensationsdruck'",
	"A16":   "'Speichertemperatur'",
	"A17":   "'Raumtemperatur'",
	"A18":   "'Raumtemperatur Mittelwert 1h'",
	"A19":   "'Warmwassertemperatur'",
	"A20":   "'Pool-Temperatur'",
	"A21":   "'Solarkollektortemperatur'",
	"A22":   "'Solarkreis Vorlauftemperatur'",
	"A23":   "'Ventilöffnung el. Exp.ventil'",
	"A25":   "'elektrische Leistung Verdichter'",
	"A26":   "'abgegeb. therm. Heizleistung Wärmepumpe'",
	"A27":   "'abgegeb. therm. KälteLeistung Wärmepumpe'",
	"A28":   "'Leistungszahl Heizleistung Wärmepumpe'",
	"A29":   "'Leistungszahl WärmeLeistung Wärmepumpe'",
	"A30":   "'Aktuelle Heizkreistemperatur'",
	"A31":   "'Geforderte Temperatur im Heizbetrieb'",
	"A32":   "'Sollwertvorgabe Heizkreistemperatur'",
	"A33":   "'Aktuelle Kühlkreistemperatur'",
	"A34":   "'Geforderte Temperatur im Kühlbetrieb'",
	"A35":   "'Sollwertvorgabe Kühlbetrieb'",
	"A37":   "'Sollwert Warmwassertemperatur'",
	"A38":   "'Sollwertvorgabe Warmwassertemperatur'",
	"A40":   "'Sollwert Poolwassertemperatur'",
	"A41":   "'Sollwertvorgabe Poolwassertemperatur'",
	"A50":   "'Geforderte Verdichterleistung'",
	"A51":   "'Comp Ventil Kältekreis %'",
	"A56":   "'Comp Ventil Kältekreis Soll'",
	"A57":   "'Comp Ventil Kältekreis Soll2'",
	"A58":   "'Comp Leistung %'",
	"A109":  "'Kühltemperatur'",
	"A444":  "'A444 wk_parse.js Z.532ff = 0'",
	"A445":  "'A445 wk_parse.js Z.532ff = 0'",
	"A446":  "'A446 wk_parse.js Z.532ff = 0'",
	"A447":  "'A447 wk_parse.js Z.532ff = 0'",
	"A448":  "'A448 wk_parse.js Z.532ff = 15905'",
	"A449":  "'A449 wk_parse.js Z.532ff = 16111'",
	"A450":  "'A450 wk_parse.js Z.532ff = 15905'",
	"A451":  "'A451 wk_parse.js Z.532ff = 16111'",
	"A452":  "'A452 wk_parse.js Z.532ff = 15277'",
	"A453":  "'A453 wk_parse.js Z.532ff =-7284'",
	"A454":  "'A454 wk_parse.js Z.532ff = 15899'",
	"A455":  "'A455 wk_parse.js Z.532ff = -12333'",
	"A456":  "'A456 wk_parse.js Z.532ff = 0'",
	"A457":  "'A457 wk_parse.js Z.532ff = 0'",
	"A722":  "'A722 index.html Z.152 =12 =Monat?!'",
	"A781":  "'A781 index.html Z.152 =290'",
	"I1":    " Firmware-Version Regler value 10401 => 01.04.0110401'",
	"I2":    " Firmware-Build Regler",
	"I3":    " BIOS-Version value 620 => 06.20",
	"I4":    "'? 1282'",
	"I5":    "'Datum: Tag'",
	"I6":    "'Datum: Monat'",
	"I7":    "'Datum: Jahr'",
	"I8":    "'Uhrzeit: Stunde'",
	"I9":    "'Uhrzeit: Minute'",
	"I10":   "'Betr.std. Verdichter1'",
	"I11":   "'Betr.std. Verdichter1 1/10'",
	"I12":   "'?'",
	"I13":   "'?'",
	"I14":   "'Betr.std. Verdichter 2'",
	"I15":   "'?' I15 tenth to OpHrs_comp2 (I14) ?",
	"I16":   "'?'",
	"I17":   "'?'",
	"I18":   "'Betr.std. Hzg.umwälzpumpe'",
	"I19":   "'Betr.std. Hzg.umwälzpumpe 1/10'",
	"I20":   "'Betr.std. Quellenpumpe'",
	"I22":   "'Betr.std. Solarkreis'",
	"I30":   "'Einst.Heizung auto=1 manuell=0'",
	"I31":   "'Einst. Kuehlung auto=1 manuell=0'",
	"I32":   "'Einst. Warmwasser auto=1 manuell=0'",
	"I33":   "'? 1'",
	"I34":   "'? 1'",
	"I35":   "'? 1'",
	"I36":   "'? 0'",
	"I37":   "'? 1'",
	"I38":   "'? 1'",
	"I39":   "'? 1'",
	"I40":   "'? 1'",
	"I41":   "'? 1'",
	"I42":   "'? 1'",
	"I51":   "'Status Waermepumpenkomponenten-Bits'",
	"I52":   "'Meldg. von Ausfall-Bits F0xx'",
	"I53":   "'Meldg. Unterbrechg.-Bits I0xx'",
	"I115":  "'? CompIconstate (gr gn rt) gem. App'",
	"I136":  "'Op.mode Hzg/Hzg-Kuehlg/Kuehlg= 0/1/2'",
	"I137":  "'Heizbetrieb inakt./akt./man. (0/1/2)'",
	"I138":  "'Kuehlbetrieb inakt./akt./man. (0/1/2)'",
	"I139":  "'Warmwasser inak./akt./man. (0/1/2)'",
	"I263":  "'Temp.anpass. (-2°+ (0.5*<(wert 0..8)>)'",
	"I1458": "'CompIconstate (gr gn rt)'",
	"I1459": "'I1459 index.html Z.118 = state_compmode comp2 Mode ?",
	"I1460": "'I1460 wk_parse.jsl Z.343 comp2?' I1460 if Comp2 ?",
	"D116":  "'D116 index.html Z.118'",
	"D420":  "'Urlaubsmodus aus=0/an=1'",
	"D671":  "'Alarm anstehend ? 0=keine'",
	"D696":  "'D696 index.html Z.118'",
}
