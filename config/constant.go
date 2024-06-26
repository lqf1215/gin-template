package config

const (
	LocalUseridUint  = "user_id_uint"
	LocalUseridInt64 = "user_id_int64"
	LocalToken       = "token"
	LocalAuthority   = "authority"

	AdminUseridInt64 = "admin_id_int64"   //管理员id
	AdminUsername    = "admin_user_name"  //管理员用户名
	ManageRole       = "manage_user_role" //管理系统角色
	ManageUser       = "manage_user"      //管理系统用户
)

const MESSAGE_SUCCESS = 0
const MESSAGE_FAIL = -1
const TOKEN_FAIL = -2
const OPERATION_FAIL = -3
