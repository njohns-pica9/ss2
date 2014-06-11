package migdam

import (
    "encoding/json"
    "time"
)

type AssetMessage struct {
    AssetId        int64             `json:"asset_id"`
    MimeType       string            `json:"mime_type_name"`
    MimeTypeId     int64             `json:"mime_type_id"`
    IsUploaded     bool              `json:"is_uploaded"`
    IsReleased     bool              `json:"is_released"`
    UploadedBy     int               `json:"uploaded_by_user_id"`
    OwnedBy        int               `json:"owned_by_user_id"`
    ExpirationDate time.Time         `json:"expiration_date"`
    Created        time.Time         `json:"created"`
    Updated        time.Time         `json:"updated"`
    Deleted        time.Time         `json:"deleted"`
    Datapoints     []json.RawMessage `json:"datapoints"`
}

func (a *AssetEntity) MakeAssetMessage() AssetMessage {
    var message AssetMessage

    // Init Defaults
    message.AssetId = a.AssetId
    message.MimeType = a.MimeType
    message.IsUploaded = a.IsUploaded
    message.IsReleased = a.IsReleased
    message.UploadedBy = a.UploadedBy
    message.OwnedBy = a.OwnedBy
    message.Created = a.Created
    message.Updated = a.Updated

    if a.MimeTypeId.Valid {
        message.MimeTypeId = a.MimeTypeId.Int64
    }

    if a.ExpirationDate.Valid {
        message.ExpirationDate = a.ExpirationDate.Time
    }

    if a.Deleted.Valid {
        message.Deleted = a.Deleted.Time
    }

    return message
}

func (a *AssetEntity) MarshalJSON() ([]byte, error) {
    message := a.MakeAssetMessage()

    return json.Marshal(message)
}

func (a *AssetMessage) MarshalJSON() ([]byte, error) {
    return json.Marshal(a)
}
