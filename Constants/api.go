package Constants

const (
	// api 前缀
	apiPrefix string = "/api"

	// API版本前缀
	versionPrefix string = "/v1"

	// 请求方法前缀
	postPrefix string = "/post"
	getPrefix  string = "/get"

	// 模块前缀
	authPrefix   string = "/auth"
	userPrefix   string = "/user"
	notePrefix   string = "/note"
	imagePrefix  string = "/picture"
	adminPrefix  string = "/admin"
	tagPrefix    string = "/tag"
	folderPrefix string = "/folder"

	// 认证模块API

	LoginApi    string = apiPrefix + versionPrefix + postPrefix + authPrefix + "/login"
	LogoutApi   string = apiPrefix + versionPrefix + postPrefix + authPrefix + "/logout"
	RefreshApi  string = apiPrefix + versionPrefix + postPrefix + authPrefix + "/refresh"
	RegisterApi string = apiPrefix + versionPrefix + postPrefix + authPrefix + "/register"

	// 用户模块API

	UserInfoApi    string = apiPrefix + versionPrefix + postPrefix + userPrefix + "/GetInfo"
	UpdatePassword string = apiPrefix + versionPrefix + postPrefix + userPrefix + "/UpdatePassword"
	UpdateIcon     string = apiPrefix + versionPrefix + postPrefix + userPrefix + "/UpdateIcon"
	UpdateUserInfo string = apiPrefix + versionPrefix + postPrefix + userPrefix + "/UpdateUserInfo"

	// 笔记模块API

	// Markdown笔记

	NewNoteApi       string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 新建笔记
	GetNoteApi       string = apiPrefix + versionPrefix + postPrefix + notePrefix + "/GetNote" // 获取笔记内容
	DelNoteApi       string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 软删除该笔记
	SaveNoteApi      string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 保存该笔记
	SetNoteNameApi   string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 重命名笔记
	SetNoteTagApi    string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 设置笔记tag
	RemoveNoteTagApi string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 移除笔记tag
	SetIsPublicApi   string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 该笔记是否不登录即可访问
	MoveNoteApi      string = apiPrefix + versionPrefix + postPrefix + notePrefix + ""         // 迁移该笔记至另一目录

	// 笔记相关内容

	// Tag

	NewNoteTagApi string = ""
	GetAllTagApi  string = apiPrefix + versionPrefix + postPrefix + tagPrefix + "/GetAllTag"
	DelNoteTagApi string = ""

	// 文件夹

	NewFolderApi     string = ""
	GetNoteFolderApi string = apiPrefix + versionPrefix + postPrefix + folderPrefix + "/GetNoteFolder"
	DelFolderApi     string = ""
	RenameFolderApi  string = ""
	MoveFolderApi    string = ""

	// 代码片段

	NewCodeApi  string = ""
	GetCodeApi  string = ""
	DelCodeApi  string = ""
	SaveCodeApi string = ""
	MoveCodeApi string = ""

	// 图床模块API

	// 管理员模块API
)
