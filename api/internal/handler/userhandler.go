package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im/api/internal/logic"
	"im/api/internal/svc"
	"im/api/internal/types"
)

func userHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//id := r.URL.Query().Get("id")
		//fmt.Println(id)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		// 打印请求的 Body（假设请求的数据是JSON格式）
		fmt.Println(string(body))

		// 将请求的 Body 转回到请求对象中
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		var req types.UserReq
		//Parse 解析请求，绑定到结构体上
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err.Error())
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			fmt.Println(err.Error())
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
