{{/*
Expand the name of the chart.
*/}}
{{- define "sentinel.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "sentinel.fullname" -}}
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
{{- define "sentinel.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "sentinel.labels" -}}
helm.sh/chart: {{ include "sentinel.chart" . }}
{{ include "sentinel.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "sentinel.selectorLabels" -}}
app.kubernetes.io/name: {{ include "sentinel.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/part-of: {{ .Values.global.aegis.namespace }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "sentinel.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "sentinel.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Define image for aegis sentinel
*/}}
{{- define "sentinel.repository" -}}
{{- if eq (lower $.Values.global.baseImage) "distroless" }}
{{- .Values.global.images.sentinel.distrolessRepository }}
{{- else if eq (lower $.Values.global.baseImage) "distroless-fips" }}
{{- .Values.global.images.sentinel.distrolessFipsRepository }}
{{- else if eq (lower $.Values.global.baseImage) "photon" }}
{{- .Values.global.images.sentinel.photonRepository }}
{{- else if eq (lower $.Values.global.baseImage) "photon-fips" }}
{{- .Values.global.images.sentinel.photonFipsRepository }}
{{- else }}
{{- .Values.global.images.sentinel.distrolessRepository }}
{{- end }}
{{- end }}
