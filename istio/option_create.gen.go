// Code generated by 'option-gen'. DO NOT EDIT.

package istio

import (
	prompt "github.com/c-bata/go-prompt"
)

var createOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--edit", Description: "Edit the API resource before creating"},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to create the resource"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to create the resource"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--raw", Description: "Raw URI to POST to the server.  Uses the transport specified by the kubeconfig file."},
	prompt.Suggest{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
	prompt.Suggest{Text: "--windows-line-endings", Description: "Only relevant if --edit=true. Defaults to the line ending native to your platform."},
}
