package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"
	uuid2 "github.com/suyuan32/simple-admin-common/utils/uuidx"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
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
		field.String("username").Unique().
			Comment("User's login name | 登录名"),
		field.String("password").
			Comment("Password | 密码"),
		field.String("nickname").Unique().
			Comment("Nickname | 昵称"),
		field.String("description").Optional().
			Comment("The description of user | 用户的描述信息"),
		field.String("home_path").Default("/dashboard").
			Comment("The home page that the user enters after logging in | 用户登陆后进入的首页"),
		field.String("mobile").Optional().
			Comment("Mobile number | 手机号"),
		field.String("email").Optional().
			Comment("Email | 邮箱号"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("Avatar | 头像路径"),
		field.Uint64("department_id").Optional().Default(1).
			Comment("Department ID | 部门ID"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// mixins.UUIDMixin{},
		// mixins.StatusMixin{},
		// mixins.SoftDeleteMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_users"},
	}
}
