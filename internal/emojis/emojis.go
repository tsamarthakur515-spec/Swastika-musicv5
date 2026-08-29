package emojis

// Custom emoji IDs — exact copy from Python PANDAMUSIC/modules/custom_emojis.py

const (
	Butterfly     = "6001449118000487326"
	ButterflyBlue = "5999337402840127790"
	ButterflySoft = "5999175482573068600"
	ButterflyWhite = "5084613633418199991"
	FirePurple    = "5999340396432333728"
	Fire          = "6129792056589031358"
	Fire2         = "6129639980387015660"
	Lightning     = "6026367225466720832"
	Sparkles      = "6023660820544623088"
	Star          = "6129672231296440969"
	StarW1        = "6327599386746425095"
	StarW2        = "6325715141643997191"
	StarW3        = "6327637478811374847"
	StarW4        = "6327883808070701100"
	StarW5        = "6327939243213590244"
	StarW6        = "6325389759216619619"
	Check         = "6147565374289220368"
	Gear          = "6001440193058444284"
	Warning       = "5420323339723881652"
	Arrow         = "6285315214673975495"
	Arrow2        = "6269553865813660503"
	Party         = "6282977077427702833"
	Party2        = "6129731974291527294"
	Party3        = "6131660826924292492"
	Confetti      = "6023660820544623088"
	HeartBlack    = "5352918496642604333"
	HeartBroken   = "6078087767106001151"
	HeartPulse    = "5352727529511723136"
	Skull         = "6037570896766438989"
	Skull2        = "6132184924603554220"
	SkullHeadphones = "6037182932370590949"
	Devil         = "5352542184493031170"
	Mask          = "6025929233291809651"
	Shadow        = "6026236216079290036"
	Ghost         = "5253539825360843975"
	Wolf          = "6127636064610818291"
	Cat           = "6057466460886799210"
	Smirk         = "6224119292279917749"
	Spiral        = "5971837723676249096"
	Loader        = "5974235702701853774"
	SpiralTech    = "6086639764251873025"
	Dots          = "5971944878815317190"
	Maze          = "4949560993840629085"
	Bow           = "6066395745139824604"
	Baby          = "6226493198013830325"
	PlayBtn       = "5217841878262618606"
	PauseBtn      = "5307798645427424676"
	SeekBtn       = "6258123029398687278"
)

// Semantic aliases (Python CE_*)
const (
	CEPlay     = PlayBtn
	CEPause    = PauseBtn
	CESkip     = Arrow
	CEStop     = Skull
	CEBack     = SeekBtn
	CEFwd      = SeekBtn
	CEProgress = Spiral
	CEClose    = Skull2
	CEQueue    = Fire
	CEStream   = FirePurple
	CETitle    = Star
	CEDuration = Loader
	CERequest  = Butterfly
	CEPowered  = Wolf
	CEPingTitle = Lightning
	CEPingVersion = StarW1
	CEPingMS   = Sparkles
	CEPingUptime = SpiralTech
	CEPingDatabase = Gear
	CEPingDanger = Warning
)

// TG wraps custom emoji for Telegram HTML (same as Python tg_emoji).
func TG(emojiID, fallback string) string {
	if fallback == "" {
		fallback = "✨"
	}
	return `<tg-emoji emoji-id="` + emojiID + `">` + fallback + `</tg-emoji>`
}

func Play() string     { return TG(CEPlay, "▶️") }
func Pause() string    { return TG(CEPause, "⏸") }
func Skip() string     { return TG(CESkip, "⏭") }
func Stop() string     { return TG(CEStop, "⏹") }
func Queue() string    { return TG(CEQueue, "📜") }
func Stream() string   { return TG(CEStream, "📡") }
func Title() string    { return TG(CETitle, "⭐") }
func Duration() string { return TG(CEDuration, "⏱") }
func Request() string  { return TG(CERequest, "🦋") }
func CheckMark() string { return TG(Check, "✅") }
func Warn() string     { return TG(Warning, "⚠️") }
func FireE() string    { return TG(Fire, "🔥") }
func Spark() string    { return TG(Sparkles, "✨") }
func GearE() string    { return TG(Gear, "⚙️") }
func WolfE() string    { return TG(Wolf, "🐺") }
