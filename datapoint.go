package migdam

import (
    "encoding/json"
    "time"
)

type DatapointMessage struct {
    Id            int       `json:"id"`
    AssetId       int       `json:"asset_id"`
    DatapointId   int       `json:"datapoint_id"`
    DatapointName string    `json:"datapoint_name"`
    ValidValueId  int64     `json:"valid_value_id"`
    Value         string    `json:"value"`
    Description   string    `json:"description"`
    Created       time.Time `json:"created"`
    Updated       time.Time `json:"updated"`
    Deleted       time.Time `json:"deleted"`
}

func (d *DatapointEntity) MakeDatapointMessage() DatapointMessage {
    var message DatapointMessage

    message.Id = d.Id
    message.AssetId = d.AssetId
    message.DatapointId = d.DatapointId
    message.DatapointName = d.DatapointName
    message.Created = d.Created
    message.Updated = d.Updated

    if d.ValidValueId.Valid {
        message.ValidValueId = d.ValidValueId.Int64
    }

    if d.Value.Valid {
        message.Value = d.Value.String
    }

    if d.Description.Valid {
        message.Description = d.Description.String
    }

    if d.Deleted.Valid {
        message.Deleted = d.Deleted.Time
    }

    return message
}

func (d *DatapointEntity) MarshalJSON() (json.RawMessage, error) {
    message := d.MakeDatapointMessage()

    return json.Marshal(message)
}

func (d *DatapointMessage) MarshalJSON() (json.RawMessage, error) {
    m, err := json.Marshal(d)
    if err != nil {
        return nil, err
    }

    return json.RawMessage(m), err
}
