package hotline

import (
	"encoding/binary"
	"github.com/jhalter/mobius/concat"
)

const fieldError = 100
const fieldData = 101
const fieldUserName = 102
const fieldUserID = 103
const fieldUserIconID = 104
const fieldUserLogin = 105
const fieldUserPassword = 106
const fieldRefNum = 107
const fieldTransferSize = 108
const fieldChatOptions = 109
const fieldUserAccess = 110
const fieldUserAlias = 111
const fieldUserFlags = 112
const fieldOptions = 113
const fieldChatID = 114
const fieldChatSubject = 115
const fieldWaitingCount = 116
const fieldVersion = 160
const fieldCommunityBannerID = 161
const fieldServerName = 162
const fieldFileNameWithInfo = 200
const fieldFileName = 201
const fieldFilePath = 202
const fieldFileTypeString = 205
const fieldFileCreatorString = 206
const fieldFileSize = 207
const fieldFileCreateDate = 208
const fieldFileModifyDate = 209
const fieldFileComment = 210
const fieldFileNewName = 211
const fieldFileNewPath = 212
const fieldFileType = 213
const fieldQuotingMsg = 214 // Defined but unused in the Hotline Protocol spec
const fieldAutomaticResponse = 215
const fieldFolderItemCount = 220
const fieldUsernameWithInfo = 300
const fieldNewsArtListData = 321
const fieldNewsCatName = 322
const fieldNewsCatListData15 = 323
const fieldNewsPath = 325
const fieldNewsArtID = 326
const fieldNewsArtDataFlav = 327
const fieldNewsArtTitle = 328
const fieldNewsArtPoster = 329
const fieldNewsArtDate = 330
const fieldNewsArtPrevArt = 331
const fieldNewsArtNextArt = 332
const fieldNewsArtData = 333
const fieldNewsArtFlags = 334
const fieldNewsArtParentArt = 335
const fieldNewsArt1stChildArt = 336
const fieldNewsArtRecurseDel = 337

type Field struct {
	ID        []byte // Type of field
	FieldSize []byte // Size of the data part
	Data      []byte // Actual field content
}

type requiredField struct {
	ID     int
	minLen int
	maxLen int
}

func NewField(id uint16, data []byte) Field {
	idBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(idBytes, id)

	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(len(data)))

	return Field{
		ID:        idBytes,
		FieldSize: bs,
		Data:      data,
	}
}

func (f Field) Payload() []byte {
	return concat.Slices(f.ID, f.FieldSize, f.Data)
}

type FileNameWithInfo struct {
	Type       string // file type code
	Creator    []byte // File creator code
	FileSize   uint32 // File Size in bytes
	NameScript []byte // TODO: What is this?
	NameSize   []byte // Length of name field
	Name       string // File name
}

func (f FileNameWithInfo) Payload() []byte {
	name := []byte(f.Name)
	nameSize := make([]byte, 2)
	binary.BigEndian.PutUint16(nameSize, uint16(len(name)))

	kb := f.FileSize

	fSize := make([]byte, 4)
	binary.BigEndian.PutUint32(fSize, kb)

	return concat.Slices(
		[]byte(f.Type),
		f.Creator,
		fSize,
		[]byte{0, 0, 0, 0},
		f.NameScript,
		nameSize,
		[]byte(f.Name),
	)

}
