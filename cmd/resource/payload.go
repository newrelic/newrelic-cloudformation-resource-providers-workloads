package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func (p *Payload) SetIdentifier(g *string) {
   p.model.Guid = g
}

func (p *Payload) GetIdentifier() *string {
   return p.model.Guid
}

func (p *Payload) GetIdentifierKey(a model.Action) string {
   return "guid"
}

func (p *Payload) GetTagIdentifier() *string {
   return p.model.Guid
}

func (p *Payload) GetTags() map[string]string {
   return p.model.Tags
}

func (p *Payload) HasTags() bool {
   return p.model.Tags != nil
}

func (p *Payload) GetTypeName() string {
   return typeName
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

var typeName = "NewRelic::Observability::Workloads"

func (p *Payload) NewModelFromGuid(g interface{}) (m model.Model) {
   s := fmt.Sprintf("%s", g)
   return NewPayload(&Model{Guid: &s})
}

func (p *Payload) GetGraphQLFragment() *string {
   return p.model.Workload
}

func (p *Payload) GetVariables() map[string]string {
   // ACCOUNTID comes from the configuration
   // NEXTCURSOR is a _convention_
   vars := make(map[string]string)
   if p.model.Variables != nil {
      for k, v := range p.model.Variables {
         vars[k] = v
      }
   }

   if p.model.Variables == nil {
      vars = make(map[string]string)
   }

   if p.model.Guid != nil {
      vars["GUID"] = *p.model.Guid
   }

   if p.model.Workload != nil {
      vars["WORKLOAD"] = *p.model.Workload
   }

   lqf := ""
   if p.model.ListQueryFilter != nil {
      lqf = *p.model.ListQueryFilter
   }
   vars["LISTQUERYFILTER"] = lqf

   return vars
}

func (p *Payload) GetErrorKey() string {
   return "type"
}

func (p *Payload) GetCreateMutation() string {
   return `
mutation {
  workloadCreate(accountId: {{{ACCOUNTID}}}, workload: {{{WORKLOAD}}}) {
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
  workloadUpdate(guid: "{{{GUID}}}", workload: {{{WORKLOAD}}}) {
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
      ... on WorkloadEntity {
        guid
        name
      }
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
