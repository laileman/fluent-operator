package output

import (
	"fmt"

	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins"
	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// Influxdb is the Influxdb output plugin, allows to ingest your records into an Influxdb.
type Influxdb struct {
	// IP address or hostname of the target InfluxDB service
	Host string `json:"host,omitempty"`
	// TCP port of the target InfluxDB service
	Port int `json:"port,omitempty"`
	// InfluxDB database name where records will be inserted
	Database string `json:"database,omitempty"`
	//InfluxDB bucket name where records will be inserted -
	//if specified, database is ignored and v2 of API is used
	Bucket string `json:"bucket,omitempty"`
	// Specifies the size of files in S3. Maximum size is 50G, minimim is 1M.
	Org string `json:"org,omitempty"`
	// The name of the tag whose value is incremented for the consecutive simultaneous events.
	SequenceTag string `json:"sequenceTag,omitempty"`
	// Optional username for HTTP Basic Authentication
	HttpUser string `json:"httpUser,omitempty"`
	// Password for user defined in HTTP_User
	HttpPasswd string `json:"httpPasswd,omitempty"`
	// Authentication token used with InfluDB v2 -
	// if specified, both HTTP_User and HTTP_Passwd are ignored
	HttpToken string `json:"httpToken,omitempty"`
	// Space separated list of keys that needs to be tagged
	TagKeys string `json:"tagKeys,omitempty"`
	// Automatically tag keys where value is string.
	// This option takes a boolean value: True/False, On/Off.
	AutoTags string `json:"autoTags,omitempty"`
	// Key of the string array optionally contained within
	// each log record that contains tag keys for that record.
	TagsListKey string `json:"tagsListKey,omitempty"`

	*plugins.TLS `json:"tls,omitempty"`
}

// Name implement Section() method
func (_ *Influxdb) Name() string {
	return "influxdb"
}

// Params implement Section() method
func (inf *Influxdb) Params(sl plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if inf.Host != "" {
		kvs.Insert("Host", inf.Host)
	}
	if inf.Port != 0 {
		kvs.Insert("Port", fmt.Sprint(inf.Port))
	}
	if inf.Database != "" {
		kvs.Insert("Database", inf.Database)
	}
	if inf.Bucket != "" {
		kvs.Insert("Bucket", inf.Bucket)
	}
	if inf.Org != "" {
		kvs.Insert("Org", inf.Org)
	}
	if inf.SequenceTag != "" {
		kvs.Insert("Sequence_Tag", inf.SequenceTag)
	}
	if inf.HttpUser != "" {
		kvs.Insert("HTTP_User", inf.HttpUser)
	}
	if inf.HttpPasswd != "" {
		kvs.Insert("HTTP_Passwd", inf.HttpPasswd)
	}
	if inf.HttpToken != "" {
		kvs.Insert("HTTP_Token", inf.HttpToken)
	}
	if inf.TagKeys != "" {
		kvs.Insert("Tag_Keys", inf.TagKeys)
	}
	if inf.AutoTags != "" {
		kvs.Insert("Auto_Tags", inf.AutoTags)
	}
	if inf.TagsListKey != "" {
		kvs.Insert("Tags_List_Key", inf.TagsListKey)
	}
	if inf.TLS != nil {
		tls, err := inf.TLS.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(tls)
	}
	return kvs, nil
}
