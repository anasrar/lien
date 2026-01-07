package controlerv1

import (
	"context"
	"net/http"
	"path"
	"strings"

	"github.com/anasrar/lien/internal/filesystem"
	"github.com/danielgtaylor/huma/v2"
)

var GetDirectoryEntriesOperation = huma.Operation{
	OperationID: "get-directory-entries",
	Method:      http.MethodGet,
	Path:        "/ls",
	Summary:     "Get entries from directory",
	Description: "Get entries from directory that return directories and files.",
	Tags:        []string{"filesystem"},
}

type GetDirectoryEntriesInput struct {
	Cwd string `query:"cwd" required:"true" example:"./" doc:"Path current working directory"`
}

type GetDirectoryEntriesOutput struct {
	Body struct {
		Directories []filesystem.DirectoryEntry `json:"directories" doc:"Array of directories"`
		Files       []filesystem.DirectoryEntry `json:"files" doc:"Array of files"`
	}
}

func GetDirectoryEntriesHandler(ctx context.Context, input *GetDirectoryEntriesInput) (*GetDirectoryEntriesOutput, error) {
	cwd := ctx.Value("cwd").(string)
	target := path.Join(cwd, input.Cwd)

	if !strings.HasPrefix(target, cwd) {
		return nil, huma.Error403Forbidden("Accessing parent is forbidden")
	}

	entries, err := filesystem.DirectoryEntriesFromPath(target)
	if err != nil {
		return nil, err
	}

	resp := &GetDirectoryEntriesOutput{}
	resp.Body.Directories = entries.Directories
	resp.Body.Files = entries.Files

	return resp, nil
}
