/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package journal

//import (
//	"testing"
//
//	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
//)
//
//func Test_printAudit(t *testing.T) {
//	type args struct {
//		correlationId string
//		entityName    string
//		method        string
//		url           string
//		spiffeid      string
//		message       string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "success_case",
//			args: args{
//				correlationId: "1234",
//				entityName:    "abcd",
//				method:        "GET",
//				url:           "http://localhost:5000/",
//				spiffeid:      "abcd1234",
//				message:       "testing audit func",
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			printAudit(tt.args.correlationId, tt.args.entityName, tt.args.method, tt.args.url, tt.args.spiffeid, tt.args.message)
//		})
//	}
//}
//
//func TestLog(t *testing.T) {
//	type args struct {
//		e Entry
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "nil_JournalEntry",
//			args: args{
//				e: Entry{
//					Entity: nil,
//				},
//			},
//		},
//		{
//			name: "Entity_default",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        "",
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretDeleteRequest",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretDeleteRequest{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretDeleteResponse",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretDeleteResponse{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretFetchRequest",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretFetchRequest{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretFetchResponse",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretFetchResponse{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretUpsertRequest",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretUpsertRequest{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretUpsertResponse",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretUpsertResponse{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretListRequest",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretListRequest{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretListResponse",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretListResponse{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_SecretEncryptedListResponse",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.SecretEncryptedListResponse{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//		{
//			name: "Entity_type_KeyInputRequest",
//			args: args{
//				e: Entry{
//					CorrelationId: "1234",
//					Entity:        reqres.KeyInputRequest{},
//					Method:        "test_method",
//					Url:           "test_url",
//					SpiffeId:      "test_spiffeid",
//					Event:         "test_event",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			Log(tt.args.e)
//		})
//	}
//}
