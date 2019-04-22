// @author  dreamlu
package sessions

import (
	"github.com/gin-gonic/gin"
)

// cookie session
// router := gin.Default()
// store := cookie.NewStore([]byte("secret"))
// router.Use(sessions.Sessions("mysession", store))
func SessionCookie(c *gin.Context) Session {

	session := Default(c)
	//var count int
	//v := session.Get("count")
	//if v == nil {
	//	count = 0
	//} else {
	//	count = v.(int)
	//	count++
	//}
	//session.Set("count", count)
	//session.Save()
	//c.JSON(200, gin.H{"count": count})
	return session
}

// redis session
// router := gin.Default()
// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
// router.Use(sessions.Sessions("mysession", store))
func SessionRedis(c *gin.Context) Session{

	session := Default(c)
	//var count int
	//v := session.Get("count")
	//if v == nil {
	//	count = 0
	//} else {
	//	count = v.(int)
	//	count++
	//}
	//session.Set("count", count)
	//session.Save()
	//c.JSON(200, gin.H{"count": count})
	return session
}
