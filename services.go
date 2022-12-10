package services

import (
	"fmt"
)

//Service requires various methods for the rest of the command framework to function.
// A dummy service can be used if you're looking to import a particular feature absent a service.
type Service interface {
	//Just bot things, yanno
	Shutdown()              //Shuts down the service as gracefully as possible
	CmdPrefix() string      //Returns the command prefix to use on this service
	Login()     (err error) //Login to the chat service

	//Messages are the backbone of how the command framework responds to interactions and interacts with the service.
	// For services only able to process text messages, you can use Message.String() to get a preformatted text when necessary.
	// In the case of other concepts such as Discord's interaction events, use Message.Context to track the alternative type.
	MsgEdit(msg *Message)   (ret *Message, err error) //Edits any type of message
	MsgRemove(msg *Message) (err error)               //Removes a message
	MsgSend(msg *Message, ref interface{})   (ret *Message, err error) //Sends any type of message

	//Users are who can send and receive messages, and can be actioned upon through various commands.
	GetUser(serverID, userID string)                 (ret *User, err error)  //Returns the specified user
	GetUserPerms(serverID, channelID, userID string) (ret *Perms, err error) //Returns the specified user's permission map
	UserBan(user *User, reason string, rule int)     (err error)             //Bans a user for a given reason and/or rule
	UserKick(user *User, reason string, rule int)    (err error)             //Kicks a user for a given reason and/or rule

	//Servers are organizations of channels, and contain their own global settings and features.
	GetServer(serverID string) (ret *Server, err error) //Returns the specified server

	//Voice connections are crucial for things like music and voice assistants
	VoiceJoin(serverID, channelID string, muted, deafened bool) (err error) //Joins a voice channel
	VoiceLeave(serverID string)                                 (err error) //Leaves the active voice channel
}

func Error(format string, replacements ...interface{}) error {
	if len(replacements) > 0 {
		return fmt.Errorf(format, replacements)
	}
	return fmt.Errorf(format)
}

type User struct {
	ServerID string  `json:"serverID,omitempty"`
	UserID   string  `json:"userID,omitempty"`
	Username string  `json:"username,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Roles    []*Role `json:"roles,omitempty"`
}

type Role struct {
	RoleID string `json:"roleID,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Perms struct {
	Administrator bool `json:"administrator,omitempty"`
	Kick          bool `json:"kick,omitempty"`
	Ban           bool `json:"ban,omitempty"`
}
func (p *Perms) CanAdministrate() bool {
	return p.Administrator
}
func (p *Perms) CanKick() bool {
	return p.CanAdministrate() || p.Kick
}
func (p *Perms) CanBan() bool {
	return p.CanAdministrate() || p.Ban
}

type Channel struct {
	ServerID  string `json:"serverID,omitempty"`
	ChannelID string `json:"channelID,omitempty"`
}

type Server struct {
	ServerID       string        `json:"serverID,omitempty"`
	Name           string        `json:"name,omitempty"`
	Region         string        `json:"region,omitempty"`
	OwnerID        string        `json:"ownerID,omitempty"`
	DefaultChannel string        `json:"defaultChannelID,omitempty"`
	VoiceStates    []*VoiceState `json:"voiceStates,omitempty"`
}

type VoiceState struct {
	ChannelID string `json:"channelID,omitempty"`
	UserID    string `json:"userID,omitempty"`
	SessionID string `json:"sessionID,omitempty"`
	Deaf      bool   `json:"deaf,omitempty"`
	Mute      bool   `json:"mute,omitempty"`
	SelfDeaf  bool   `json:"selfDeaf,omitempty"`
	SelfMute  bool   `json:"selfMeaf,omitempty"`
}
