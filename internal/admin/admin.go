package admin

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"
)

// Ban permanently bans user from channel/supergroup.
func Ban(c *telegram.Client, chatID, userID int64) error {
	channel, err := c.GetInputChannel(chatID)
	if err != nil {
		return fmt.Errorf("resolve channel: %w", err)
	}
	user, err := c.GetInputUser(userID)
	if err != nil {
		return fmt.Errorf("resolve user: %w", err)
	}
	_, err = c.ChannelsEditBanned(&telegram.ChannelsEditBannedParams{
		Channel: channel,
		Participant: user,
		BannedRights: &telegram.ChatBannedRights{
			ViewMessages: true,
			SendMessages: true,
			SendMedia:    true,
			SendStickers: true,
			SendGifs:     true,
			SendGames:    true,
			SendInline:   true,
			EmbedLinks:   true,
			UntilDate:    0,
		},
	})
	return err
}

// Unban lifts ban.
func Unban(c *telegram.Client, chatID, userID int64) error {
	channel, err := c.GetInputChannel(chatID)
	if err != nil {
		return err
	}
	user, err := c.GetInputUser(userID)
	if err != nil {
		return err
	}
	_, err = c.ChannelsEditBanned(&telegram.ChannelsEditBannedParams{
		Channel:     channel,
		Participant: user,
		BannedRights: &telegram.ChatBannedRights{
			UntilDate: 0,
		},
	})
	return err
}

// Mute restricts sending messages (permanent until unmute).
func Mute(c *telegram.Client, chatID, userID int64) error {
	channel, err := c.GetInputChannel(chatID)
	if err != nil {
		return err
	}
	user, err := c.GetInputUser(userID)
	if err != nil {
		return err
	}
	_, err = c.ChannelsEditBanned(&telegram.ChannelsEditBannedParams{
		Channel:     channel,
		Participant: user,
		BannedRights: &telegram.ChatBannedRights{
			SendMessages: true,
			SendMedia:    true,
			SendStickers: true,
			SendGifs:     true,
			SendGames:    true,
			SendInline:   true,
			EmbedLinks:   true,
			UntilDate:    0,
		},
	})
	return err
}

// Unmute restores send rights.
func Unmute(c *telegram.Client, chatID, userID int64) error {
	return Unban(c, chatID, userID)
}

// Kick bans briefly then unbans (user can rejoin).
func Kick(c *telegram.Client, chatID, userID int64) error {
	if err := Ban(c, chatID, userID); err != nil {
		return err
	}
	return Unban(c, chatID, userID)
}

// Promote grants basic admin rights.
func Promote(c *telegram.Client, chatID, userID int64) error {
	channel, err := c.GetInputChannel(chatID)
	if err != nil {
		return err
	}
	user, err := c.GetInputUser(userID)
	if err != nil {
		return err
	}
	_, err = c.ChannelsEditAdmin(&telegram.ChannelsEditAdminParams{
		Channel: channel,
		UserID:  user,
		AdminRights: &telegram.ChatAdminRights{
			ChangeInfo:     true,
			DeleteMessages: true,
			BanUsers:       true,
			InviteUsers:    true,
			PinMessages:    true,
			ManageCall:     true,
		},
		Rank: "admin",
	})
	return err
}

// Demote removes admin rights.
func Demote(c *telegram.Client, chatID, userID int64) error {
	channel, err := c.GetInputChannel(chatID)
	if err != nil {
		return err
	}
	user, err := c.GetInputUser(userID)
	if err != nil {
		return err
	}
	_, err = c.ChannelsEditAdmin(&telegram.ChannelsEditAdminParams{
		Channel:     channel,
		UserID:      user,
		AdminRights: &telegram.ChatAdminRights{},
		Rank:        "",
	})
	return err
}

// TargetUserID from reply or fails.
func TargetFromReply(m *telegram.NewMessage) (int64, error) {
	if m.IsReply() {
		reply, err := m.GetReplyMessage()
		if err == nil && reply != nil && reply.Sender != nil {
			return reply.SenderID(), nil
		}
	}
	return 0, fmt.Errorf("reply to a user")
}
