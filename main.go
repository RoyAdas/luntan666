package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	Author  string `json:"author"`
}

var users = map[string]User{}
var posts = []Post{}
var sessions = map[string]string{}
var nextPostID = 1

func main() {
	// 初始化账号
	users["admin"] = User{Username: "admin", Password: "123456", Role: "admin"}
	users["user1"] = User{Username: "user1", Password: "123456", Role: "user"}

	r := gin.Default()

	// 简单 CORS 支持
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// 公共路由
	r.POST("/login", login)
	r.GET("/posts", getPosts)
	r.GET("/posts/:id", getPost)

	// 受保护路由
	auth := r.Group("/")
	auth.Use(authMiddleware)
	{
		auth.GET("/user", getUser)
		auth.PUT("/user", updateUser)
		auth.POST("/posts", createPost)
		auth.PUT("/posts/:id", updatePost)
		auth.DELETE("/posts/:id", deletePost)
		auth.POST("/posts/:id/like", likePost)
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	var cred User
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	user, ok := users[cred.Username]
	if !ok || user.Password != cred.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	token := cred.Username + "-token"
	sessions[token] = cred.Username
	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}

func getUser(c *gin.Context) {
	username := c.GetString("user")
	u := users[username]
	c.JSON(http.StatusOK, gin.H{"username": u.Username, "role": u.Role})
}

func updateUser(c *gin.Context) {
	username := c.GetString("user")
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	u := users[username]
	if req.Username != "" {
		// 简单替换用户名（注意：真实项目需更复杂处理）
		delete(users, username)
		u.Username = req.Username
		username = u.Username
	}
	if req.Password != "" {
		u.Password = req.Password
	}
	users[username] = u
	c.JSON(http.StatusOK, u)
}

func getPosts(c *gin.Context) {
	search := strings.TrimSpace(c.Query("search"))
	if search == "" {
		c.JSON(http.StatusOK, posts)
		return
	}
	var res []Post
	for _, p := range posts {
		if strings.Contains(strings.ToLower(p.Title), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(p.Content), strings.ToLower(search)) {
			res = append(res, p)
		}
	}
	c.JSON(http.StatusOK, res)
}

func getPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for _, p := range posts {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
}

func createPost(c *gin.Context) {
	username := c.GetString("user")
	var req Post
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	req.ID = nextPostID
	req.Author = username
	posts = append(posts, req)
	nextPostID++
	c.JSON(http.StatusOK, req)
}

func updatePost(c *gin.Context) {
	username := c.GetString("user")
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var req Post
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	for i, p := range posts {
		if p.ID == id {
			// 仅作者或管理员可修改
			if p.Author != username && users[username].Role != "admin" {
				c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
				return
			}
			if req.Title != "" {
				posts[i].Title = req.Title
			}
			if req.Content != "" {
				posts[i].Content = req.Content
			}
			c.JSON(http.StatusOK, posts[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
}

func deletePost(c *gin.Context) {
	username := c.GetString("user")
	if users[username].Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
}

func likePost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, p := range posts {
		if p.ID == id {
			posts[i].Likes++
			c.JSON(http.StatusOK, posts[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
}

func authMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth"})
		c.Abort()
		return
	}
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		c.Abort()
		return
	}
	token := parts[1]
	username, ok := sessions[token]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Set("user", username)
	c.Next()
}
