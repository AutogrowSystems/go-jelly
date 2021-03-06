package datastructs

// IDose
type iDoseShadow struct {
	State StateIDose `json:"state"`
}

// StateIDose represents the State data structure from an IntelliDose packet
type StateIDose struct {
	Reported ReportedIDose `json:"reported"`
}

// ReportedIDose represents the Reported data structure from an IntelliDose packet
type ReportedIDose struct {
	Config    ConfigIDose  `json:"config"`
	Metrics   MetricsIDose `json:"metrics"`
	Status    StatusIDose  `json:"status"`
	Source    string       `json:"source"`
	Device    string       `json:"device"`
	Timestamp int64        `json:"timestamp"`
	Connected bool         `json:"connected"`
}

// ConfigIDose represents the Config data structure from an IntelliDose packet
type ConfigIDose struct {
	Units      UnitsIDose      `json:"units"`
	Times      TimesIDose      `json:"times"`
	Functions  FunctionsIDose  `json:"functions"`
	Advanced   AdvancedIDose   `json:"advanced"`
	General    GeneralIDose    `json:"general"`
	Scheduling SchedulingIDose `json:"scheduling"`
	Reminders  RemindersIDose  `json:"reminder"`
}

// MetricsIDose represents the Metrics data structure from an IntelliDose packet
type MetricsIDose struct {
	Ec      float64 `json:"ec"`
	NutTemp float64 `json:"nut_temp"`
	PH      float64 `json:"pH"`
}

// StatusIDose represents the Status data structure from an IntelliDose packet
type StatusIDose struct {
	General   GeneralStatusIDose  `json:"general"`
	Nutrient  NutrientIDose       `json:"nutrient"`
	SetPoints SetPointsIDose      `json:"set_points"`
	Status    []StatusStatusIDose `json:"status"`
}

// GeneralStatusIDose represents the GeneralStatus data structure from an IntelliDose packet
type GeneralStatusIDose struct {
	DoseInterval        byte                    `json:"dose_interval"`
	NutrientDoseTime    byte                    `json:"nutrient_dose_time"`
	WaterOnTime         byte                    `json:"water_on_time"`
	IrrigationInterval1 IrrigationIntervalIDose `json:"irrigation_interval_1"`
	IrrigationInterval2 IrrigationIntervalIDose `json:"irrigation_interval_2"`
	IrrigationInterval3 IrrigationIntervalIDose `json:"irrigation_interval_3"`
	IrrigationInterval4 IrrigationIntervalIDose `json:"irrigation_interval_4"`
	IrrigationDuration1 int                     `json:"irrigation_duration_1"`
	IrrigationDuration2 int                     `json:"irrigation_duration_2"`
	IrrigationDuration3 int                     `json:"irrigation_duration_3"`
	IrrigationDuration4 int                     `json:"irrigation_duration_4"`
	MaxNutrientDoseTime byte                    `json:"max_nutrient_dose_time"`
	MaxPhDoseTime       byte                    `json:"max_ph_dose_time"`
	Mix1                byte                    `json:"mix_1"`
	Mix2                byte                    `json:"mix_2"`
	Mix3                byte                    `json:"mix_3"`
	Mix4                byte                    `json:"mix_4"`
	Mix5                byte                    `json:"mix_5"`
	Mix6                byte                    `json:"mix_6"`
	Mix7                byte                    `json:"mix_7"`
	Mix8                byte                    `json:"mix_8"`
	PhDoseTime          byte                    `json:"ph_dose_time"`
}

// IrrigationIntervalIDose represents the IrrigationInterval data structure from an IntelliDose packet
type IrrigationIntervalIDose struct {
	Day   int `json:"day"`
	Night int `json:"night"`
	Every int `json:"every"`
}

// NutrientIDose represents the Nutrient data structure from an IntelliDose packet
type NutrientIDose struct {
	Detent  byte         `json:"detent"`
	Ec      EcIDose      `json:"ec"`
	NutTemp NutTempIDose `json:"nut_temp"`
	Ph      PhIDose      `json:"ph"`
}

// EcIDose represents the Ec data structure from an IntelliDose packet
type EcIDose struct {
	Enabled bool    `json:"enabled"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

// NutTempIDose represents the NutTemp data structure from an IntelliDose packet
type NutTempIDose struct {
	Enabled bool    `json:"enabled"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

// PhIDose represents the Ph data structure from an IntelliDose packet
type PhIDose struct {
	Enabled bool    `json:"enabled"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

// SetPointsIDose represents the SetPoints data structure from an IntelliDose packet
type SetPointsIDose struct {
	Nutrient      float64 `json:"nutrient"`
	NutrientNight float64 `json:"nutrient_night"`
	PhDosing      string  `json:"ph_dosing"`
	Ph            float64 `json:"ph"`
}

// StatusStatusIDose represents the StatusStatus data structure from an IntelliDose packet
type StatusStatusIDose struct {
	Active   bool   `json:"active"`
	Enabled  bool   `json:"enabled"`
	ForceOn  bool   `json:"force_on"`
	Function string `json:"function"`
}

// UnitsIDose represents the Units data structure from an IntelliDose packet
type UnitsIDose struct {
	DateFormat              string `json:"date_format"`
	Temperature             string `json:"temperature"`
	Ec                      string `json:"ec"`
	TdsConversationStandart int    `json:"tds_conversation_standart"`
}

// TimesIDose represents the Times data structure from an IntelliDose packet
type TimesIDose struct {
	DayStart int `json:"day_start"`
	DayEnd   int `json:"day_end"`
}

// FunctionsIDose represents the Functions data structure from an IntelliDose packet
type FunctionsIDose struct {
	NutrientsParts     byte   `json:"nutrients_parts"`
	PhDosing           string `json:"ph_dosing"`
	IrrigationMode     string `json:"irrigation_mode"`
	IrrigationStations byte   `json:"irrigation_stations"`
	SeparatePumpOutput bool   `json:"separate_pump_output"`
	UseWater           bool   `json:"use_water"`
	ExternalAlarm      bool   `json:"external_alarm"`
	DayNightEc         bool   `json:"day_night_ec"`
	IrrigationStation1 string `json:"irrigation_station_1"`
	IrrigationStation2 string `json:"irrigation_station_2"`
	IrrigationStation3 string `json:"irrigation_station_3"`
	IrrigationStation4 string `json:"irrigation_station_4"`
	Scheduling         bool   `json:"scheduling"`
	MuteBuzzer         bool   `json:"mute_buzzer"`
}

// AdvancedIDose represents the Advanced data structure from an IntelliDose packet
type AdvancedIDose struct {
	ProportinalDosing bool   `json:"proportinal_dosing"`
	SequentialDosing  bool   `json:"sequential_dosing"`
	DisableEc         bool   `json:"disable_ec"`
	DisablePh         bool   `json:"disable_ph"`
	MntnReminderFreq  string `json:"mntn_reminder_freq"`
}

// GeneralIDose represents the General data structure from an IntelliDose packet
type GeneralIDose struct {
	DeviceName string  `json:"device_name"`
	Firmware   float64 `json:"firmware"`
	Growroom   string  `json:"growroom"`
}

// SchedulingIDose represents the scheduling data structure from an IntelliDose packet
type SchedulingIDose struct {
	LastUpdated float64 `json:"last_updated"`
	Mode        string  `json:"mode"`
}

// ReminderIDose represents the a single reminder data structure from an IntelliDose packet
type ReminderIDose struct {
	CleanECProbe float64 `json:"clean_ec_probe"`
	CleanpHProbe float64 `json:"clean_ph_electrode"`
	CheckECProbe float64 `json:"check_ec_probe"`
	CalibratePH  float64 `json:"calibrate_ph"`
	CleanFilters float64 `json:"clean_filters"`
}

// RemindersIDose represents the reminders data structure from an IntelliDose packet
type RemindersIDose struct {
	Frequency    string        `json:"frequency"`
	StartDate    float64       `json:"start_date"`
	ReminderList ReminderIDose `json:"reminder_list"`
}

// DoserHistory - consists of a slice of history points
type DoserHistory struct {
	Points []*DoserHistoryPoint `json:"points"`
}

// DoseMetricsHistory - Metrics the history point contains
type DoseMetricsHistory struct {
	EC   float64 `json:"ec"`
	PH   float64 `json:"pH"`
	Temp float64 `json:"nut_temp"`
}

// DoserHistoryPoint - defines a single history point reported for a IntelliDose
type DoserHistoryPoint struct {
	Timestamp float64            `json:"timestamp"`
	Status    Status             `json:"status"`
	Metrics   DoseMetricsHistory `json:"metrics"`
}
