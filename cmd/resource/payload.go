package resource

import (
   "fmt"
   "github.com/newrelic-experimental/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func NewPayload(m *Model) *Payload {
   return &Payload{
      model:  m,
      models: make([]interface{}, 0),
   }
}

func (p *Payload) GetResourceModel() interface{} {
   return p.model
}

func (p *Payload) GetResourceModels() []interface{} {
   log.Debugf("GetResourceModels: returning %+v", p.models)
   return p.models
}

func (p *Payload) AppendToResourceModels(m model.Model) {
   p.models = append(p.models, m.GetResourceModel())
}

//
// These are API specific, must be configured per API
//

var typeName = "NewRelic::CloudFormation::Workloads"

func (p *Payload) NewModelFromGuid(g interface{}) (m model.Model) {
   s := fmt.Sprintf("%s", g)
   return NewPayload(&Model{Guid: &s})
}

func (p *Payload) GetGraphQLFragment() *string {
   return p.model.Workload
}

func (p *Payload) SetGuid(g *string) {
   p.model.Guid = g
   log.Debugf("SetGuid: %s", *p.model.Guid)
}

func (p *Payload) GetGuid() *string {
   return p.model.Guid
}

func (p *Payload) GetCreateMutation() string {
   return `
mutation {
  workloadCreate(accountId: {{{ACCOUNTID}}}, {{{WORKLOAD}}}) {
    guid
  }
}
`
}

func (p *Payload) GetDeleteMutation() string {
   return `
mutation {
  workloadDelete(guid: "{{{GUID}}}") {
    guid
  }
}
`
}

func (p *Payload) GetUpdateMutation() string {
   return `
mutation {
  workloadUpdate(guid: "{{{GUID}}}", {{{WORKLOAD}}}) {
    guid
  }
}
`
}

func (p *Payload) GetReadQuery() string {
   return `
{
  actor {
    entity(guid: "{{{GUID}}}") {
        guid
        name
    }
  }
}
`
}

func (p *Payload) GetListQuery() string {
   return `
{
  actor {
    entitySearch(queryBuilder: {type: WORKLOAD}) {
      count
      results {
        nextCursor
        entities {
            guid
            name
        }
      }
    }
  }
}
`
}

func (p *Payload) GetListQueryNextCursor() string {
   return `
{
  actor {
    entitySearch(queryBuilder: {type: WORKLOAD}) {
      count
      results(cursor: "{{{NEXTCURSOR}}}") {
        nextCursor
        entities {
            guid
            name
        }
      }
    }
  }
}
`
}

func (p *Payload) GetGuidKey() string {
   return "guid"
}

func (p *Payload) GetVariables() map[string]string {
   // ACCOUNTID comes from the configuration
   // NEXTCURSOR is a _convention_

   if p.model.Variables == nil {
      p.model.Variables = make(map[string]string)
   }

   if p.model.Guid != nil {
      p.model.Variables["GUID"] = *p.model.Guid
   }

   if p.model.Workload != nil {
      p.model.Variables["WORKLOAD"] = *p.model.Workload
   }

   lqf := ""
   if p.model.ListQueryFilter != nil {
      lqf = *p.model.ListQueryFilter
   }
   p.model.Variables["LISTQUERYFILTER"] = lqf

   return p.model.Variables
}

func (p *Payload) GetErrorKey() string {
   return "type"
}

func (p *Payload) GetResultKey(a model.Action) string {
   return p.GetGuidKey()
}

func (p *Payload) NeedsPropagationDelay(a model.Action) bool {
   return true
}
