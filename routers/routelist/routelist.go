package routelist

import (
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouteList 路由列表
var RouteList *Routes

// Routes 路由
type Routes struct {
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
	// List PATH列表
	List cmap.CMap
}

func NewRoute(router *gin.Engine) *Routes {
	return &Routes{
		Router: router,
		List:   cmap.NewCMap(),
	}
}

// Group is a shortcut for router.Handle("POST", path, handle).
func (group *Routes) Group(relativePath string, handlers ...gin.HandlerFunc) *Routes {
	// 路由都是同步注册,nothing to do
	if group.RouterGroup != nil {
		relativePath = group.RouterGroup.BasePath() + relativePath
	}
	return &Routes{
		Router:      group.Router,
		RouterGroup: group.Router.Group(relativePath, handlers...),
		List:        group.List,
	}
}

// POST is a shortcut for router.Handle("POST", path, handle).
func (group *Routes) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodPost)
	return group.Router.POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (group *Routes) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodGet)
	return group.Router.GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (group *Routes) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodDelete)
	return group.Router.DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (group *Routes) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodPatch)
	return group.Router.PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (group *Routes) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodPut)
	return group.Router.PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (group *Routes) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodOptions)
	return group.Router.OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (group *Routes) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodHead)
	return group.Router.HEAD(relativePath, handlers...)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *Routes) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, "*")
	return group.Router.Any(relativePath, handlers...)
}

// StaticFile registers a single route in order to serve a single file of the local filesystem.
// router.StaticFile("favicon.ico", "./resources/favicon.ico")
func (group *Routes) StaticFile(relativePath, filepath string) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodGet)
	return group.Router.StaticFile(relativePath, filepath)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func (group *Routes) Static(relativePath, root string) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodGet)
	return group.Router.Static(relativePath, root)
}

// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (group *Routes) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	relativePath = group.RouterGroup.BasePath() + relativePath
	group.List.Set(relativePath, http.MethodGet)
	return group.Router.StaticFS(relativePath, fs)
}
