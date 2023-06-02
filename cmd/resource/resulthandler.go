package resource

import (
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/client/nerdgraph"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

// ResultHandler at a minimum provides access to the default error processing.
// If required we can provide custom processing here via composition overrides https://go.dev/doc/effective_go#embedding
type ResultHandler struct {
   // Use Go composition to access the default implementation
   model.ResultHandler
}

func NewResultHandler() (h model.ResultHandler) {
   defer func() {
      log.Debugf("(tagging) errorHandler.NewErrorHandler: exit %p", h)
   }()
   // Initialize ourself with the common core
   h = &ResultHandler{ResultHandler: nerdgraph.NewResultHandler()}
   return
}
