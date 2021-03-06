package memory

//InMenuValues inside osu!memory
type InMenuValues struct {
	OsuStatus   uint32 `json:"state"`
	GameMode    uint32 `json:"gameMode"`
	ChatChecker int8   `json:"isChatEnabled"` //bool (1 byte)
	Bm          bm     `json:"bm"`
	Mods        modsM  `json:"mods"`
	PP          ppM    `json:"pp"`
}

//GameplayValues inside osu!memory
type GameplayValues struct {
	GameMode    int32       `json:"gameMode"`
	Score       int32       `json:"score"`
	Accuracy    float64     `json:"accuracy"`
	Combo       combo       `json:"combo"`
	Hp          hp          `json:"hp"`
	Hits        hits        `json:"hits"`
	PP          ppG         `json:"pp"`
	Leaderboard leaderboard `json:"leaderboard"`
}

type bm struct {
	Time           tim      `json:"time"`
	BeatmapID      uint32   `json:"id"`
	BeatmapSetID   uint32   `json:"set"`
	Metadata       Metadata `json:"metadata"`
	Stats          stats    `json:"stats"`
	Path           path     `json:"path"`
	HitObjectStats string   `json:"bmStats"`
	BeatmapString  string   `json:"bmInfo"`
}

type tim struct {
	FirstObj int32 `json:"firstObj"`
	PlayTime int32 `json:"current"`
	FullTime int32 `json:"full"`
	Mp3Time  int32 `json:"mp3"`
}

// Metadata Map data
type Metadata struct {
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Mapper  string `json:"mapper"`
	Version string `json:"difficulty"`
}

type stats struct {
	BeatmapAR float32 `json:"AR"`
	BeatmapCS float32 `json:"CS"`
	BeatmapOD float32 `json:"OD"`
	BeatmapHP float32 `json:"HP"`
	BeatmapSR float32 `json:"SR"`
	FullSR    float32 `json:"fullSR"`
	MemoryAR  float32 `json:"memoryAR"`
	MemoryCS  float32 `json:"memoryCS"`
	MemoryOD  float32 `json:"memoryOD"`
	MemoryHP  float32 `json:"memoryHP"`
}

type path struct {
	InnerBGPath          string `json:"full"`
	BeatmapFolderString  string `json:"folder"`
	BeatmapOsuFileString string `json:"file"`
	BGPath               string `json:"bg"`
	AudioPath            string `json:"audio"`
	FullMP3Path          string `json:"-"`
	FullDotOsu           string `json:"-"`
}

type modsM struct {
	AppliedMods int32  `json:"num"`
	PpMods      string `json:"str"`
}

type ppM struct {
	PpSS      int32     `json:"100"`
	Pp99      int32     `json:"99"`
	Pp98      int32     `json:"98"`
	Pp97      int32     `json:"97"`
	Pp96      int32     `json:"96"`
	Pp95      int32     `json:"95"`
	PpStrains []float64 `json:"strains"`
}

type combo struct {
	Current int32 `json:"current"`
	Max     int32 `json:"max"`
}

type hp struct {
	Normal float64 `json:"normal"`
	Smooth float64 `json:"smooth"`
}

type hits struct {
	H300          int16   `json:"300"`
	H100          int16   `json:"100"`
	H50           int16   `json:"50"`
	H0            int16   `json:"0"`
	HitErrorArray []int32 `json:"hitErrorArray"`
}

type ppG struct {
	Pp     int32 `json:"current"`
	PPifFC int32 `json:"fc"`
}

type dynamicAddresses struct {
	PlayContainer38    uint32
	BeatmapAddr        uint32
	LeaderBoardStruct  uint32
	IsReady            bool
	LeaderBaseSlotAddr []uint32
	LeaderSlotAddr     []uint32
}

type leaderPlayerS struct {
	Addr          uint32 `json:"-"`
	Name          string `json:"name"`
	Position      int32  `json:"position"`
	AmountOfSlots int32  `json:"amofslots"`
}

type leaderboard struct {
	OurPlayer leaderPlayerS `json:"ourplayer"`
	Slots     slotPlayerS   `json:"slots"`
}

type slotPlayerS struct {
	Name     []string `json:"name"`
	Score    []int32  `json:"score"`
	Combo    []int16  `json:"combo"`
	MaxCombo []int32  `json:"maxcombo"`
	Mods     []int32  `json:"mods"`
	H300     []int16  `json:"300"`
	H100     []int16  `json:"100"`
	H50      []int16  `json:"50"`
	H0       []int16  `json:"0"`
}

//MenuData contains raw values taken from osu! memory
var MenuData = InMenuValues{}

//GameplayData contains raw values taken from osu! memory
var GameplayData = GameplayValues{}

//DynamicAddresses are in-between pointers that lead to values
var DynamicAddresses = dynamicAddresses{}
