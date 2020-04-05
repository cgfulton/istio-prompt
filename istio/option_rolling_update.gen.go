// Code generated by 'option-gen'. DO NOT EDIT.

package istio

import (
	prompt "github.com/c-bata/go-prompt"
)

var rollingUpdateOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--container", Description: "Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod"},
	prompt.Suggest{Text: "--deployment-label-key", Description: "The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise"},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename or URL to file to use to create the new replication controller."},
	prompt.Suggest{Text: "--filename", Description: "Filename or URL to file to use to create the new replication controller."},
	prompt.Suggest{Text: "--image", Description: "Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f"},
	prompt.Suggest{Text: "--image-pull-policy", Description: "Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--poll-interval", Description: "Time delay between polling for replication controller status after the update. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--rollback", Description: "If true, this is a request to abort an existing rollout that is partially rolled out. It effectively reverses current and next and runs a rollout"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--timeout", Description: "Max time to wait for a replication controller to update before giving up. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--update-period", Description: "Time to wait between updating pods. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}
