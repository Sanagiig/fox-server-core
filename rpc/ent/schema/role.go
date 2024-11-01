package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"
	uuid2 "github.com/suyuan32/simple-admin-common/utils/uuidx"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid2.NewUUID).Comment("UUID"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("Create Time | 创建日期"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Update Time | 修改日期"),
		field.String("name").
			Comment("Role name | 角色名"),
		field.String("code").
			Comment("Role code for permission control in front end | 角色码，用于前端权限控制"),
		field.String("default_router").Default("dashboard").
			Comment("Default menu : dashboard | 默认登录页面"),
		field.String("remark").Default("").
			Comment("Remark | 备注"),
		field.Uint32("sort").Default(0).
			Comment("Order number | 排序编号"),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("roles"),
	}
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique(),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_roles"},
	}
}
