package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/cferror"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/client/nerdgraph"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
   "strings"
)

// ErrorHandler at a minimum provides access to the default error processing.
// If required we can provide custom processing here via composition overrides, aka type embedding https://go.dev/doc/effective_go#embedding
type ErrorHandler struct {
   // Use Go composition to access the default implementation, lack of a variable name here directs go to use method forwarding
   model.ErrorHandler
   M model.Model
}

// NewErrorHandler This is all pretty magical. We return the interface so common is insulated from an implementation. Payload implements model.Model so all is good
func NewErrorHandler(p *Payload) (h model.ErrorHandler) {
   defer func() {
      log.Debugf("errorHandler.NewErrorHandler: exit %p", h)
   }()
   // Initialize ourself with the common core
   h = &ErrorHandler{ErrorHandler: nerdgraph.NewCommonErrorHandler(p), M: p}
   return
}

func (h *ErrorHandler) TypeSpecificError(data *[]byte, s string) (err error) {
   defer func() {
      log.Debugf("errorHandler.TypeSpecificError (shadowed): self: %p exit", h)
   }()
   log.Debugf("TypeSpecificError: %p enter", h)
   defer func() {
      log.Debugf("TypeSpecificError: %p returning %v", h, err)
   }()
   v, err := nerdgraph.FindKeyValue(*data, "error")
   log.Debugf("TypeSpecificError: found: %v %T", v, v)
   if err != nil {
      return
   }
   if v == nil {
      return
   }

   errorMap := make(map[string]interface{})
   h.GetErrorMap(v, errorMap)

   if errorMap == nil {
      log.Warnf("Empty errors array: %v+ %T", h, h)
      return
   }
   _type := fmt.Sprintf("%v", errorMap[h.M.GetErrorKey()])
   if strings.Contains(strings.ToLower(_type), "not_found") || strings.Contains(strings.ToLower(_type), "invalid_parameter") {
      err = fmt.Errorf("%w Not found", &cferror.NotFound{})
      return
   }
   return
}
