package output

import (
	"fmt"

	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins"
	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// SimpleStorageService is the es output plugin, allows to ingest your records into an S3.
type SimpleStorageService struct {
	// The AWS region of your S3 bucket
	Region string `json:"region,omitempty"`
	// S3 Bucket name
	Bucket string `json:"bucket,omitempty"`
	// Specify the name of the time key in the output record.
	// To disable the time key just set the value to false.
	JsonDateKey string `json:"jsonDateKey,omitempty"`
	//Specify the format of the date. Supported formats are double, epoch,
	//iso8601 (eg: 2018-05-30T09:39:52.000681Z)
	//and
	//java_sql_timestamp (eg: 2018-05-30 09:39:52.000681)
	JsonDateFormat string `json:"jsonDateFormat,omitempty"`
	// Specifies the size of files in S3. Maximum size is 50G, minimim is 1M.
	TotalFileSize string `json:"totalFileSize,omitempty"`
	// The size of each 'part' for multipart uploads. Max: 50M
	UploadChunkSize string `json:"uploadChunkSize,omitempty"`
	// Whenever this amount of time has elapsed,
	// Fluent Bit will complete an upload and create a new file in S3.
	// For example, set this value to 60m and you will get a new file every hour.
	UploadTimeout string `json:"uploadTimeout,omitempty"`
	// Directory to locally buffer data before sending.
	// When multipart uploads are used, data will only be buffered
	// until the upload_chunk_size is reached.
	StoreDir string `json:"storeDir,omitempty"`
	// Format string for keys in S3. This option supports a UUID,
	// strftime time formatters, a syntax for selecting parts of
	// the Fluent log tag using a syntax inspired by the rewrite_tag filter.
	// Add $UUID in the format string to insert a random string.
	// Add $INDEX in the format string to insert an integer that increments each upload.
	// Add $TAG in the format string to insert the full log tag; add $TAG[0]
	// to insert the first part of the tag in the s3 key.
	// The tag is split into “parts” using the characters specified
	// with the s3_key_format_tag_delimiters option. Add extension directly
	// after the last piece of the format string to insert a key suffix.
	// If you want to specify a key suffix and you are in use_put_object mode,
	// you must specify $UUID as well.
	// More explanations can be found in use_put_object option.
	// See the in depth examples and tutorial in the documentation.
	S3KeyFormat string `json:"s3KeyFormat,omitempty"`
	// A series of characters which will be used to split the tag into 'parts'
	// for use with the s3_key_format option.
	// See the in depth examples and tutorial in the documentation.
	S3KeyFormatTagDelimiters string `json:"s3KeyFormatTagDelimiters,omitempty"`
	// Disables behavior where UUID string is automatically
	// appended to end of S3 key name when $UUID is not
	// provided in s3_key_format. $UUID, time formatters,
	// $TAG, and other dynamic key formatters all work
	// as expected while this feature is set to true.
	StaticFilePath *bool `json:"staticFilePath,omitempty"`
	// Use the S3 PutObject API, instead of the multipart upload API.
	// When this option is on, key extension is only available
	// when $UUID is specified in s3_key_format.
	//If $UUID is not included, a random string will be appended
	// at the end of the format string and
	// the key extension cannot be customized in this case.
	UsePutObject *bool `json:"usePutObject,omitempty"`
	// ARN of an IAM role to assume (ex. for cross account access).
	RoleArn string `json:"roleArn,omitempty"`
	// Custom endpoint for the S3 API. An endpoint can contain scheme and port.
	Endpoint string `json:"endpoint,omitempty"`
	// Custom endpoint for the STS API.
	StsEndpoint string `json:"stsEndpoint,omitempty"`
	// Predefined Canned ACL policy for S3 objects.
	CannedAcl string `json:"cannedAcl,omitempty"`
	// Compression type for S3 objects. 'gzip' is currently the only supported value.
	// The Content-Encoding HTTP Header will be set to 'gzip'.
	// Compression can be enabled when use_put_object is on.
	// If Apache Arrow support was enabled at compile time,
	// you can set 'arrow' to this option.
	Compression string `json:"compression,omitempty"`
	// A standard MIME type for the S3 object;
	// this will be set as the Content-Type HTTP header.
	ContentType string `json:"contentType,omitempty"`
	// Send the Content-MD5 header with PutObject and UploadPart requests,
	// as is required when Object Lock is enabled.
	SendContentMd5 string `json:"sendContentMd5,omitempty"`
	// Immediately retry failed requests to AWS services once.
	// This option does not affect the normal Fluent Bit retry mechanism with backoff.
	// Instead, it enables an immediate retry with no delay for networking errors,
	// which may help improve throughput when there are transient/random networking issues.
	AutoRetryRequests *bool `json:"autoRetryRequests,omitempty"`
	// By default, the whole log record will be sent to S3.
	// If you specify a key name with this option,
	// then only the value of that key will be sent to S3.
	// For example, if you are using Docker,
	// you can specify log_key log and only the log message will be sent to S3.
	LogKey string `json:"logKey,omitempty"`
	// Normally, when an upload request fails, there is a high chance for the
	// last received chunk to be swapped with a later chunk, resulting in data shuffling.
	// This feature prevents this shuffling by using a queue logic for uploads.
	PreserveDataOrdering *bool `json:"preserveDataOrdering,omitempty"`
	// Specify the storage class for S3 objects.
	// If this option is not specified,
	// objects will be stored with the default 'STANDARD' storage class.
	StorageClass string `json:"storageClass,omitempty"`
	*plugins.TLS `json:"tls,omitempty"`
}

// Name implement Section() method
func (_ *SimpleStorageService) Name() string {
	return "s3"
}

// Params implement Section() method
func (s3 *SimpleStorageService) Params(sl plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if s3.Region != "" {
		kvs.Insert("region", s3.Region)
	}
	if s3.Bucket != "" {
		kvs.Insert("bucket", s3.Bucket)
	}
	if s3.JsonDateKey != "" {
		kvs.Insert("json_date_key", s3.JsonDateKey)
	}
	if s3.JsonDateFormat != "" {
		kvs.Insert("json_date_format", s3.JsonDateFormat)
	}
	if s3.TotalFileSize != "" {
		kvs.Insert("total_file_size", s3.TotalFileSize)
	}
	if s3.UploadChunkSize != "" {
		kvs.Insert("upload_chunk_size", s3.UploadChunkSize)
	}
	if s3.UploadTimeout != "" {
		kvs.Insert("upload_timeout", s3.UploadTimeout)
	}
	if s3.StoreDir != "" {
		kvs.Insert("store_dir", s3.StoreDir)
	}
	if s3.S3KeyFormat != "" {
		kvs.Insert("s3_key_format", s3.S3KeyFormat)
	}
	if s3.S3KeyFormatTagDelimiters != "" {
		kvs.Insert("s3_key_format_tag_delimiters", s3.S3KeyFormatTagDelimiters)
	}
	if s3.StaticFilePath != nil {
		kvs.Insert("static_file_path", fmt.Sprint(*s3.StaticFilePath))
	}
	if s3.UsePutObject != nil {
		kvs.Insert("use_put_object", fmt.Sprint(*s3.UsePutObject))
	}
	if s3.RoleArn != "" {
		kvs.Insert("role_arn", s3.RoleArn)
	}
	if s3.Endpoint != "" {
		kvs.Insert("endpoint", s3.Endpoint)
	}
	if s3.StsEndpoint != "" {
		kvs.Insert("sts_endpoint", s3.StsEndpoint)
	}
	if s3.CannedAcl != "" {
		kvs.Insert("canned_acl", s3.CannedAcl)
	}
	if s3.Compression != "" {
		kvs.Insert("compression", s3.Compression)
	}
	if s3.ContentType != "" {
		kvs.Insert("content_type", s3.ContentType)
	}
	if s3.SendContentMd5 != "" {
		kvs.Insert("send_content_md5", s3.SendContentMd5)
	}
	if s3.AutoRetryRequests != nil {
		kvs.Insert("auto_retry_requests", fmt.Sprint(*s3.AutoRetryRequests))
	}
	if s3.LogKey != "" {
		kvs.Insert("log_key", s3.LogKey)
	}
	if s3.PreserveDataOrdering != nil {
		kvs.Insert("preserve_data_ordering", fmt.Sprint(*s3.PreserveDataOrdering))
	}
	if s3.StorageClass != "" {
		kvs.Insert("storage_class", s3.StorageClass)
	}
	if s3.TLS != nil {
		tls, err := s3.TLS.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(tls)
	}
	return kvs, nil
}
