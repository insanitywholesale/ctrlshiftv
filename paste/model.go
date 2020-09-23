package paste

type Paste struct {
	Code      string `json:"code" bson:"code" msgpack:"code" xml:"code"`
	Content   string `json:"content" bson:"content" msgpack:"content" xml:"content" validate:"empty=false"`
	CreatedAt int64  `json:"created_at" bson:"created_at" msgpack:"created_at" xml:"created_at"`
}
