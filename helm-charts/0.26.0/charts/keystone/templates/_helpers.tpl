# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

{{/*
Expand the name of the chart.
*/}}
{{- define "keystone.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "keystone.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "keystone.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "keystone.labels" -}}
helm.sh/chart: {{ include "keystone.chart" . }}
{{ include "keystone.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "keystone.selectorLabels" -}}
app.kubernetes.io/name: {{ include "keystone.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/part-of: {{ .Values.global.vsecm.namespace }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "keystone.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "keystone.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Define image for VSecM Keystone
*/}}
{{- define "keystone.repository" -}}
{{- if eq (lower $.Values.global.baseImage) "distroless" }}
{{- .Values.global.images.keystone.distrolessRepository }}
{{- else if eq (lower $.Values.global.baseImage) "distroless-fips" }}
{{- .Values.global.images.keystone.distrolessFipsRepository }}
{{- else }}
{{- .Values.global.images.keystone.distrolessRepository }}
{{- end }}
{{- end }}
