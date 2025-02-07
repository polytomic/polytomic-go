// This file was auto-generated by Fern from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/polytomic/polytomic-go/core"
	time "time"
)

type CreateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}

type UpdateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}

type Webhook struct {
	CreatedAt      *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	Endpoint       *string    `json:"endpoint,omitempty" url:"endpoint,omitempty"`
	Id             *string    `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string    `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         *string    `json:"secret,omitempty" url:"secret,omitempty"`

	_rawJSON json.RawMessage
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type embed Webhook
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*w),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*w = Webhook(unmarshaler.embed)
	w.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *Webhook) MarshalJSON() ([]byte, error) {
	type embed Webhook
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*w),
		CreatedAt: core.NewOptionalDateTime(w.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (w *Webhook) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookEnvelope struct {
	Data *Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (w *WebhookEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookEnvelope(value)
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookEnvelope) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookListEnvelope struct {
	Data []*Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (w *WebhookListEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookListEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookListEnvelope(value)
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookListEnvelope) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}
