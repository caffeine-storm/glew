package glew

// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
// #cgo freebsd  CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib -lglfw
// #include "glew.h"
// void SetGlewExperimental(GLboolean v) {  glewExperimental = v;  }
import "C"


