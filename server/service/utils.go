package service

import "github.com/04Akaps/Video_Chat_App/types"

func verifyPagingOption(paging *types.Paging) {
	if paging.Pagination {
		if paging.PageSize == 0 {
			paging.PageSize = 10
		}
	}
}
