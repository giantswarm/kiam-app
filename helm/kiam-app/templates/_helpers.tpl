{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "labels.common" -}}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
app.kubernetes.io/name: {{ .Values.name | quote }}
app.kubernetes.io/instance: {{ .Release.Name | quote }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
giantswarm.io/service-type: "{{ .Values.serviceType }}"
helm.sh/chart: {{ include "chart" . | quote }}
{{- end -}}

{{/*
Agent selector labels
*/}}
{{- define "labels.selector.agent" -}}
app: {{ .Values.name | quote }}
component: {{ .Values.agent.name }}
{{- end -}}

{{/*
Server selector labels
*/}}
{{- define "labels.selector.server" -}}
app: {{ .Values.name | quote }}
component: {{ .Values.server.name }}
{{- end -}}

{{/*
Agent labels
*/}}
{{- define "labels.agent" -}}
{{ include "labels.common" . }}
{{ include "labels.selector.agent" . }}
{{- end -}}

{{/*
Server labels
*/}}
{{- define "labels.server" -}}
{{ include "labels.common" . }}
{{ include "labels.selector.server" . }}
{{- end -}}
