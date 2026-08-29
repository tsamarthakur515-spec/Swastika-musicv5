package formatters

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/emojis"
)

var smallMap = func() *strings.Replacer {
	from := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	to := "ᴀʙᴄᴅᴇғɢʜɪᴊᴋʟᴍɴᴏᴘǫʀsᴛᴜᴠᴡxʏᴢᴀʙᴄᴅᴇғɢʜɪᴊᴋʟᴍɴᴏᴘǫʀsᴛᴜᴠᴡxʏᴢ"
	args := make([]string, 0, 52*2)
	for i := 0; i < len(from); i++ {
		args = append(args, string(from[i]), string([]rune(to)[i]))
	}
	return strings.NewReplacer(args...)
}()

func Smallcaps(text string) string {
	return smallMap.Replace(text)
}

func PanelCaption(title, duration, requester, header string) string {
	if header == "" {
		header = "sᴛʀᴇᴀᴍɪɴɢ ɪɴ ᴠᴄ"
	}
	t := html.EscapeString(Smallcaps(title))
	d := html.EscapeString(Smallcaps(duration))
	req := requester
	if !strings.Contains(requester, "<a ") {
		req = html.EscapeString(Smallcaps(requester))
	}
	h := emojis.Stream() + " " + html.EscapeString(header)
	return fmt.Sprintf(
		`<blockquote expandable>%s

%s %s : %s
%s %s : %s
%s %s : %s

%s ᴘᴏᴡᴇʀᴇᴅ ʙʏ : sᴡᴀsᴛɪᴋᴀ
%s ʏᴛ ᴍᴜsɪᴄ ᴀᴘɪ ᴘᴏᴡᴇʀᴇᴅ ʙʏ : ᴀʀᴜʏᴛ ᴀᴘɪ</blockquote>`,
		h,
		emojis.Title(), Smallcaps("title"), t,
		emojis.Duration(), Smallcaps("duration"), d,
		emojis.Request(), Smallcaps("request by"), req,
		emojis.WolfE(),
		emojis.Spark(),
	)
}

func QueueCaption(position int, title, duration, requester string) string {
	header := fmt.Sprintf("%s #%d", Smallcaps("added to queue"), position)
	t := html.EscapeString(Smallcaps(title))
	d := html.EscapeString(Smallcaps(duration))
	req := requester
	if !strings.Contains(requester, "<a ") {
		req = html.EscapeString(Smallcaps(requester))
	}
	return fmt.Sprintf(
		`<blockquote expandable>%s %s

%s %s : %s
%s %s : %s
%s %s : %s</blockquote>`,
		emojis.Queue(), html.EscapeString(header),
		emojis.Title(), Smallcaps("title"), t,
		emojis.Duration(), Smallcaps("duration"), d,
		emojis.Request(), Smallcaps("request by"), req,
	)
}

func Duration(seconds int) string {
	if seconds < 0 {
		return "0:00"
	}
	m := seconds / 60
	s := seconds % 60
	if m >= 60 {
		return fmt.Sprintf("%d:%02d:%02d", m/60, m%60, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}

func Uptime(start time.Time) string {
	d := time.Since(start)
	return fmt.Sprintf("%dh %dm %ds", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60)
}
