package model

const (
	// SIRecordType 診療行為レコードのレコード識別情報
	SIRecordType = "SI"
)

// SI 診療行為レコード
type SI struct {
	FacilityID string `dynamo:"-"` // 医療機関コード
	InvoiceYM  string `dynamo:"-"` // 請求年月
	Index      uint64 `dynamo:"-"` // インデックス
	ReceiptNo  uint32 `dynamo:"-"` // レセプト番号

	RecordType    string // レコード識別情報
	TreatmentType uint8  // 診療識別
	ChargeType    string // 負担区分
	TreatmentID   string // 診療行為コード
	Quantity      uint32 // 数量データ
	Point         uint32 // 点数
	Times         uint16 // 回数
	CommentID1    uint32 // コメントコード
	Comment1      string // 文字データ
	CommentID2    uint32 // コメントコード
	Comment2      string // 文字データ
	CommentID3    uint32 // コメントコード
	Comment3      string // 文字データ
	Day1          uint16 // 1日の情報
	Day2          uint16 // 2日の情報
	Day3          uint16 // 3日の情報
	Day4          uint16 // 4日の情報
	Day5          uint16 // 5日の情報
	Day6          uint16 // 6日の情報
	Day7          uint16 // 7日の情報
	Day8          uint16 // 8日の情報
	Day9          uint16 // 9日の情報
	Day10         uint16 // 10日の情報
	Day11         uint16 // 11日の情報
	Day12         uint16 // 12日の情報
	Day13         uint16 // 13日の情報
	Day14         uint16 // 14日の情報
	Day15         uint16 // 15日の情報
	Day16         uint16 // 16日の情報
	Day17         uint16 // 17日の情報
	Day18         uint16 // 18日の情報
	Day19         uint16 // 19日の情報
	Day20         uint16 // 20日の情報
	Day21         uint16 // 21日の情報
	Day22         uint16 // 22日の情報
	Day23         uint16 // 23日の情報
	Day24         uint16 // 24日の情報
	Day25         uint16 // 25日の情報
	Day26         uint16 // 26日の情報
	Day27         uint16 // 27日の情報
	Day28         uint16 // 28日の情報
	Day29         uint16 // 29日の情報
	Day30         uint16 // 30日の情報
	Day31         uint16 // 31日の情報
}

// GetFacilityID 医療機関コード
func (si *SI) GetFacilityID() string {
	return si.FacilityID
}

// GetInvoiceYM 請求年月
func (si *SI) GetInvoiceYM() string {
	return si.InvoiceYM
}

// GetIndex インデックス
func (si *SI) GetIndex() uint64 {
	return si.Index
}

// GetReceiptNo レセプト番号
func (si *SI) GetReceiptNo() uint32 {
	return si.ReceiptNo
}