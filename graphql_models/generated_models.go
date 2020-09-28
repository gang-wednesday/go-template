// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_models

type BooleanFilter struct {
	IsTrue  *bool `json:"isTrue"`
	IsFalse *bool `json:"isFalse"`
	IsNull  *bool `json:"isNull"`
}

type ChangePasswordResponse struct {
	Ok bool `json:"ok"`
}

type Comment struct {
	ID        string `json:"id"`
	User      *User  `json:"user"`
	Post      *Post  `json:"post"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt *int   `json:"createdAt"`
	UpdatedAt *int   `json:"updatedAt"`
	DeletedAt *int   `json:"deletedAt"`
}

type CommentCreateInput struct {
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt *int   `json:"createdAt"`
	UpdatedAt *int   `json:"updatedAt"`
	DeletedAt *int   `json:"deletedAt"`
}

type CommentDeletePayload struct {
	ID string `json:"id"`
}

type CommentFilter struct {
	Search *string       `json:"search"`
	Where  *CommentWhere `json:"where"`
}

type CommentPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type CommentPayload struct {
	Comment *Comment `json:"comment"`
}

type CommentUpdateInput struct {
	UserID    *string `json:"userId"`
	PostID    *string `json:"postId"`
	Title     *string `json:"title"`
	Body      *string `json:"body"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type CommentWhere struct {
	ID        *IDFilter     `json:"id"`
	User      *UserWhere    `json:"user"`
	Post      *PostWhere    `json:"post"`
	Title     *StringFilter `json:"title"`
	Body      *StringFilter `json:"body"`
	CreatedAt *IntFilter    `json:"createdAt"`
	UpdatedAt *IntFilter    `json:"updatedAt"`
	DeletedAt *IntFilter    `json:"deletedAt"`
	Or        *CommentWhere `json:"or"`
	And       *CommentWhere `json:"and"`
}

type CommentsCreateInput struct {
	Comments []*CommentCreateInput `json:"comments"`
}

type CommentsDeletePayload struct {
	Ids []string `json:"ids"`
}

type CommentsPayload struct {
	Comments []*Comment `json:"comments"`
}

type CommentsUpdatePayload struct {
	Ok bool `json:"ok"`
}

type CompaniesCreateInput struct {
	Companies []*CompanyCreateInput `json:"companies"`
}

type CompaniesDeletePayload struct {
	Ids []string `json:"ids"`
}

type CompaniesPayload struct {
	Companies []*Company `json:"companies"`
}

type CompaniesUpdatePayload struct {
	Ok bool `json:"ok"`
}

type Company struct {
	ID        string      `json:"id"`
	Name      *string     `json:"name"`
	Active    *bool       `json:"active"`
	UpdatedAt *int        `json:"updatedAt"`
	DeletedAt *int        `json:"deletedAt"`
	CreatedAt *int        `json:"createdAt"`
	Locations []*Location `json:"locations"`
	Users     []*User     `json:"users"`
}

type CompanyCreateInput struct {
	Name      *string `json:"name"`
	Active    *bool   `json:"active"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
}

type CompanyDeletePayload struct {
	ID string `json:"id"`
}

type CompanyFilter struct {
	Search *string       `json:"search"`
	Where  *CompanyWhere `json:"where"`
}

type CompanyPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type CompanyPayload struct {
	Company *Company `json:"company"`
}

type CompanyUpdateInput struct {
	Name      *string `json:"name"`
	Active    *bool   `json:"active"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
}

type CompanyWhere struct {
	ID        *IDFilter      `json:"id"`
	Name      *StringFilter  `json:"name"`
	Active    *BooleanFilter `json:"active"`
	UpdatedAt *IntFilter     `json:"updatedAt"`
	DeletedAt *IntFilter     `json:"deletedAt"`
	CreatedAt *IntFilter     `json:"createdAt"`
	Locations *LocationWhere `json:"locations"`
	Users     *UserWhere     `json:"users"`
	Or        *CompanyWhere  `json:"or"`
	And       *CompanyWhere  `json:"and"`
}

type FloatFilter struct {
	EqualTo           *float64  `json:"equalTo"`
	NotEqualTo        *float64  `json:"notEqualTo"`
	LessThan          *float64  `json:"lessThan"`
	LessThanOrEqualTo *float64  `json:"lessThanOrEqualTo"`
	MoreThan          *float64  `json:"moreThan"`
	MoreThanOrEqualTo *float64  `json:"moreThanOrEqualTo"`
	In                []float64 `json:"in"`
	NotIn             []float64 `json:"notIn"`
}

type Follower struct {
	ID        string `json:"id"`
	Follower  *User  `json:"follower"`
	Followee  *User  `json:"followee"`
	DeletedAt *int   `json:"deletedAt"`
	UpdatedAt *int   `json:"updatedAt"`
	CreatedAt *int   `json:"createdAt"`
}

type FollowerCreateInput struct {
	FollowerID string `json:"followerId"`
	FolloweeID string `json:"followeeId"`
	DeletedAt  *int   `json:"deletedAt"`
	UpdatedAt  *int   `json:"updatedAt"`
	CreatedAt  *int   `json:"createdAt"`
}

type FollowerDeletePayload struct {
	ID string `json:"id"`
}

type FollowerFilter struct {
	Search *string        `json:"search"`
	Where  *FollowerWhere `json:"where"`
}

type FollowerPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type FollowerPayload struct {
	Follower *Follower `json:"follower"`
}

type FollowerUpdateInput struct {
	FollowerID *string `json:"followerId"`
	FolloweeID *string `json:"followeeId"`
	DeletedAt  *int    `json:"deletedAt"`
	UpdatedAt  *int    `json:"updatedAt"`
	CreatedAt  *int    `json:"createdAt"`
}

type FollowerWhere struct {
	ID        *IDFilter      `json:"id"`
	Follower  *UserWhere     `json:"follower"`
	Followee  *UserWhere     `json:"followee"`
	DeletedAt *IntFilter     `json:"deletedAt"`
	UpdatedAt *IntFilter     `json:"updatedAt"`
	CreatedAt *IntFilter     `json:"createdAt"`
	Or        *FollowerWhere `json:"or"`
	And       *FollowerWhere `json:"and"`
}

type FollowersCreateInput struct {
	Followers []*FollowerCreateInput `json:"followers"`
}

type FollowersDeletePayload struct {
	Ids []string `json:"ids"`
}

type FollowersPayload struct {
	Followers []*Follower `json:"followers"`
}

type FollowersUpdatePayload struct {
	Ok bool `json:"ok"`
}

type IDFilter struct {
	EqualTo    *string  `json:"equalTo"`
	NotEqualTo *string  `json:"notEqualTo"`
	In         []string `json:"in"`
	NotIn      []string `json:"notIn"`
}

type IntFilter struct {
	EqualTo           *int  `json:"equalTo"`
	NotEqualTo        *int  `json:"notEqualTo"`
	LessThan          *int  `json:"lessThan"`
	LessThanOrEqualTo *int  `json:"lessThanOrEqualTo"`
	MoreThan          *int  `json:"moreThan"`
	MoreThanOrEqualTo *int  `json:"moreThanOrEqualTo"`
	In                []int `json:"in"`
	NotIn             []int `json:"notIn"`
}

type Location struct {
	ID        string   `json:"id"`
	Name      *string  `json:"name"`
	Active    *bool    `json:"active"`
	Address   *string  `json:"address"`
	Company   *Company `json:"company"`
	UpdatedAt *int     `json:"updatedAt"`
	DeletedAt *int     `json:"deletedAt"`
	CreatedAt *int     `json:"createdAt"`
	Users     []*User  `json:"users"`
}

type LocationCreateInput struct {
	Name      *string `json:"name"`
	Active    *bool   `json:"active"`
	Address   *string `json:"address"`
	CompanyID string  `json:"companyId"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
}

type LocationDeletePayload struct {
	ID string `json:"id"`
}

type LocationFilter struct {
	Search *string        `json:"search"`
	Where  *LocationWhere `json:"where"`
}

type LocationPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type LocationPayload struct {
	Location *Location `json:"location"`
}

type LocationUpdateInput struct {
	Name      *string `json:"name"`
	Active    *bool   `json:"active"`
	Address   *string `json:"address"`
	CompanyID *string `json:"companyId"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
}

type LocationWhere struct {
	ID        *IDFilter      `json:"id"`
	Name      *StringFilter  `json:"name"`
	Active    *BooleanFilter `json:"active"`
	Address   *StringFilter  `json:"address"`
	Company   *CompanyWhere  `json:"company"`
	UpdatedAt *IntFilter     `json:"updatedAt"`
	DeletedAt *IntFilter     `json:"deletedAt"`
	CreatedAt *IntFilter     `json:"createdAt"`
	Users     *UserWhere     `json:"users"`
	Or        *LocationWhere `json:"or"`
	And       *LocationWhere `json:"and"`
}

type LocationsCreateInput struct {
	Locations []*LocationCreateInput `json:"locations"`
}

type LocationsDeletePayload struct {
	Ids []string `json:"ids"`
}

type LocationsPayload struct {
	Locations []*Location `json:"locations"`
}

type LocationsUpdatePayload struct {
	Ok bool `json:"ok"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type Post struct {
	ID        string     `json:"id"`
	User      *User      `json:"user"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	DeletedAt *int       `json:"deletedAt"`
	CreatedAt *int       `json:"createdAt"`
	UpdatedAt *int       `json:"updatedAt"`
	Comments  []*Comment `json:"comments"`
}

type PostCreateInput struct {
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	DeletedAt *int   `json:"deletedAt"`
	CreatedAt *int   `json:"createdAt"`
	UpdatedAt *int   `json:"updatedAt"`
}

type PostDeletePayload struct {
	ID string `json:"id"`
}

type PostFilter struct {
	Search *string    `json:"search"`
	Where  *PostWhere `json:"where"`
}

type PostPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type PostPayload struct {
	Post *Post `json:"post"`
}

type PostUpdateInput struct {
	UserID    *string `json:"userId"`
	Title     *string `json:"title"`
	Body      *string `json:"body"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
}

type PostWhere struct {
	ID        *IDFilter     `json:"id"`
	User      *UserWhere    `json:"user"`
	Title     *StringFilter `json:"title"`
	Body      *StringFilter `json:"body"`
	DeletedAt *IntFilter    `json:"deletedAt"`
	CreatedAt *IntFilter    `json:"createdAt"`
	UpdatedAt *IntFilter    `json:"updatedAt"`
	Comments  *CommentWhere `json:"comments"`
	Or        *PostWhere    `json:"or"`
	And       *PostWhere    `json:"and"`
}

type PostsCreateInput struct {
	Posts []*PostCreateInput `json:"posts"`
}

type PostsDeletePayload struct {
	Ids []string `json:"ids"`
}

type PostsPayload struct {
	Posts []*Post `json:"posts"`
}

type PostsUpdatePayload struct {
	Ok bool `json:"ok"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type Role struct {
	ID          string  `json:"id"`
	AccessLevel int     `json:"accessLevel"`
	Name        string  `json:"name"`
	UpdatedAt   *int    `json:"updatedAt"`
	DeletedAt   *int    `json:"deletedAt"`
	CreatedAt   *int    `json:"createdAt"`
	Users       []*User `json:"users"`
}

type RoleCreateInput struct {
	AccessLevel int    `json:"accessLevel"`
	Name        string `json:"name"`
	UpdatedAt   *int   `json:"updatedAt"`
	DeletedAt   *int   `json:"deletedAt"`
	CreatedAt   *int   `json:"createdAt"`
}

type RoleDeletePayload struct {
	ID string `json:"id"`
}

type RoleFilter struct {
	Search *string    `json:"search"`
	Where  *RoleWhere `json:"where"`
}

type RolePagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type RolePayload struct {
	Role *Role `json:"role"`
}

type RoleUpdateInput struct {
	AccessLevel *int    `json:"accessLevel"`
	Name        *string `json:"name"`
	UpdatedAt   *int    `json:"updatedAt"`
	DeletedAt   *int    `json:"deletedAt"`
	CreatedAt   *int    `json:"createdAt"`
}

type RoleWhere struct {
	ID          *IDFilter     `json:"id"`
	AccessLevel *IntFilter    `json:"accessLevel"`
	Name        *StringFilter `json:"name"`
	UpdatedAt   *IntFilter    `json:"updatedAt"`
	DeletedAt   *IntFilter    `json:"deletedAt"`
	CreatedAt   *IntFilter    `json:"createdAt"`
	Users       *UserWhere    `json:"users"`
	Or          *RoleWhere    `json:"or"`
	And         *RoleWhere    `json:"and"`
}

type RolesCreateInput struct {
	Roles []*RoleCreateInput `json:"roles"`
}

type RolesDeletePayload struct {
	Ids []string `json:"ids"`
}

type RolesPayload struct {
	Roles []*Role `json:"roles"`
}

type RolesUpdatePayload struct {
	Ok bool `json:"ok"`
}

type StringFilter struct {
	EqualTo            *string  `json:"equalTo"`
	NotEqualTo         *string  `json:"notEqualTo"`
	In                 []string `json:"in"`
	NotIn              []string `json:"notIn"`
	StartWith          *string  `json:"startWith"`
	NotStartWith       *string  `json:"notStartWith"`
	EndWith            *string  `json:"endWith"`
	NotEndWith         *string  `json:"notEndWith"`
	Contain            *string  `json:"contain"`
	NotContain         *string  `json:"notContain"`
	StartWithStrict    *string  `json:"startWithStrict"`
	NotStartWithStrict *string  `json:"notStartWithStrict"`
	EndWithStrict      *string  `json:"endWithStrict"`
	NotEndWithStrict   *string  `json:"notEndWithStrict"`
	ContainStrict      *string  `json:"containStrict"`
	NotContainStrict   *string  `json:"notContainStrict"`
}

type User struct {
	ID                 string      `json:"id"`
	FirstName          *string     `json:"firstName"`
	LastName           *string     `json:"lastName"`
	Username           *string     `json:"username"`
	Password           *string     `json:"password"`
	Email              *string     `json:"email"`
	Mobile             *string     `json:"mobile"`
	Phone              *string     `json:"phone"`
	Address            *string     `json:"address"`
	Active             *bool       `json:"active"`
	LastLogin          *int        `json:"lastLogin"`
	LastPasswordChange *int        `json:"lastPasswordChange"`
	Token              *string     `json:"token"`
	Role               *Role       `json:"role"`
	Company            *Company    `json:"company"`
	Location           *Location   `json:"location"`
	CreatedAt          *int        `json:"createdAt"`
	DeletedAt          *int        `json:"deletedAt"`
	UpdatedAt          *int        `json:"updatedAt"`
	Comments           []*Comment  `json:"comments"`
	FolloweeFollowers  []*Follower `json:"followeeFollowers"`
	FollowerFollowers  []*Follower `json:"followerFollowers"`
	Posts              []*Post     `json:"posts"`
}

type UserCreateInput struct {
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	Username           *string `json:"username"`
	Password           *string `json:"password"`
	Email              *string `json:"email"`
	Mobile             *string `json:"mobile"`
	Phone              *string `json:"phone"`
	Address            *string `json:"address"`
	Active             *bool   `json:"active"`
	LastLogin          *int    `json:"lastLogin"`
	LastPasswordChange *int    `json:"lastPasswordChange"`
	Token              *string `json:"token"`
	RoleID             *string `json:"roleId"`
	CompanyID          *string `json:"companyId"`
	LocationID         *string `json:"locationId"`
	CreatedAt          *int    `json:"createdAt"`
	DeletedAt          *int    `json:"deletedAt"`
	UpdatedAt          *int    `json:"updatedAt"`
}

type UserDeletePayload struct {
	ID string `json:"id"`
}

type UserFilter struct {
	Search *string    `json:"search"`
	Where  *UserWhere `json:"where"`
}

type UserPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type UserPayload struct {
	User *User `json:"user"`
}

type UserUpdateInput struct {
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	Username           *string `json:"username"`
	Password           *string `json:"password"`
	Email              *string `json:"email"`
	Mobile             *string `json:"mobile"`
	Phone              *string `json:"phone"`
	Address            *string `json:"address"`
	Active             *bool   `json:"active"`
	LastLogin          *int    `json:"lastLogin"`
	LastPasswordChange *int    `json:"lastPasswordChange"`
	Token              *string `json:"token"`
	RoleID             *string `json:"roleId"`
	CompanyID          *string `json:"companyId"`
	LocationID         *string `json:"locationId"`
	CreatedAt          *int    `json:"createdAt"`
	DeletedAt          *int    `json:"deletedAt"`
	UpdatedAt          *int    `json:"updatedAt"`
}

type UserWhere struct {
	ID                 *IDFilter      `json:"id"`
	FirstName          *StringFilter  `json:"firstName"`
	LastName           *StringFilter  `json:"lastName"`
	Username           *StringFilter  `json:"username"`
	Password           *StringFilter  `json:"password"`
	Email              *StringFilter  `json:"email"`
	Mobile             *StringFilter  `json:"mobile"`
	Phone              *StringFilter  `json:"phone"`
	Address            *StringFilter  `json:"address"`
	Active             *BooleanFilter `json:"active"`
	LastLogin          *IntFilter     `json:"lastLogin"`
	LastPasswordChange *IntFilter     `json:"lastPasswordChange"`
	Token              *StringFilter  `json:"token"`
	Role               *RoleWhere     `json:"role"`
	Company            *CompanyWhere  `json:"company"`
	Location           *LocationWhere `json:"location"`
	CreatedAt          *IntFilter     `json:"createdAt"`
	DeletedAt          *IntFilter     `json:"deletedAt"`
	UpdatedAt          *IntFilter     `json:"updatedAt"`
	Comments           *CommentWhere  `json:"comments"`
	FolloweeFollowers  *FollowerWhere `json:"followeeFollowers"`
	FollowerFollowers  *FollowerWhere `json:"followerFollowers"`
	Posts              *PostWhere     `json:"posts"`
	Or                 *UserWhere     `json:"or"`
	And                *UserWhere     `json:"and"`
}

type UsersCreateInput struct {
	Users []*UserCreateInput `json:"users"`
}

type UsersDeletePayload struct {
	Ids []string `json:"ids"`
}

type UsersPayload struct {
	Users []*User `json:"users"`
}

type UsersUpdatePayload struct {
	Ok bool `json:"ok"`
}
